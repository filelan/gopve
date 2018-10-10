package service

import (
	"errors"
	"strconv"

	"github.com/xabinapal/gopve/internal"
)

type QEMUServiceProvider interface {
	List() (*QEMUList, error)
	Get(int) (*QEMU, error)
	Start(int) error
	Stop(int) error
	Reset(int) error
	Shutdown(int) error
	Suspend(int) error
	Resume(int) error
	Create() error
	Clone(vmid int, opts *VMCreateOptions) (*Task, error)
	Update(int, *QEMUConfig) error
	Delete(int) (*Task, error)
}

type QEMUService struct {
	client *internal.Client
	node   *Node
}

type QEMUServiceFactoryProvider interface {
	Create(*Node) QEMUServiceProvider
}

type QEMUServiceFactory struct {
	client    *internal.Client
	providers map[string]QEMUServiceProvider
}

func NewQEMUServiceFactoryProvider(c *internal.Client) QEMUServiceFactoryProvider {
	return &QEMUServiceFactory{
		client:    c,
		providers: make(map[string]QEMUServiceProvider),
	}
}

func (factory *QEMUServiceFactory) Create(node *Node) QEMUServiceProvider {
	provider, ok := factory.providers[node.Node]
	if !ok {
		provider = &QEMUService{
			client: factory.client,
			node:   node,
		}

		factory.providers[node.Node] = provider
	}

	return provider
}

func (s *QEMUService) List() (*QEMUList, error) {
	data, err := s.client.Get("nodes/" + s.node.Node + "/qemu")
	if err != nil {
		return nil, err
	}

	var res QEMUList
	for _, qemu := range internal.NewJArray(data) {
		val := internal.NewJObject(qemu)
		vmid, _ := strconv.Atoi(val.GetString("vmid"))
		row := &QEMU{
			provider: s,

			VMID:   vmid,
			Name:   val.GetString("name"),
			Status: val.GetString("status"),
			QEMUConfig: QEMUConfig{
				CPU:           val.GetInt("cpus"),
				MemoryTotal:   val.GetInt("maxmem"),
				MemoryMinimum: val.GetIntDefault("balloon_min", 0),
			},
		}

		if row.MemoryMinimum == 0 {
			row.MemoryMinimum = row.MemoryTotal
			row.MemoryBallooning = false
		} else {
			row.MemoryBallooning = true
		}

		res = append(res, row)
	}

	return &res, nil
}

func (s *QEMUService) Get(vmid int) (*QEMU, error) {
	dataConfig, err := s.client.Get("nodes/" + s.node.Node + "/qemu/" + strconv.Itoa(vmid) + "/config")
	if err != nil {
		return nil, err
	}

	dataStatus, err := s.client.Get("nodes/" + s.node.Node + "/qemu/" + strconv.Itoa(vmid) + "/status/current")
	if err != nil {
		return nil, err
	}

	valConfig := internal.NewJObject(dataConfig)
	valStatus := internal.NewJObject(dataStatus)

	res := &QEMU{
		provider: s,

		VMID:   vmid,
		Name:   valStatus.GetString("name"),
		Status: valStatus.GetString("status"),
		QEMUConfig: QEMUConfig{
			OSType:        valConfig.GetString("ostype"),
			CPUSockets:    valConfig.GetInt("sockets"),
			CPUCores:      valConfig.GetInt("cores"),
			CPULimit:      valConfig.GetIntDefault("cpulimit", QEMUDefaultCPULimit),
			CPUUnits:      valConfig.GetIntDefault("cpuunits", QEMUDefaultCPUUnits),
			MemoryTotal:   valConfig.GetInt("memory"),
			MemoryMinimum: valConfig.GetInt("balloon"),
			IsNUMAAware:   valConfig.GetBool("numa"),
		},
	}

	res.CPU = res.CPUSockets * res.CPUCores
	if res.MemoryMinimum == 0 {
		res.MemoryMinimum = res.MemoryTotal
		res.MemoryBallooning = false
	} else {
		res.MemoryBallooning = true
	}

	return res, nil
}

func (s *QEMUService) power(vmid int, command string) error {
	_, err := s.client.Post("nodes/"+s.node.Node+"/qemu/"+strconv.Itoa(vmid)+"/status/"+command, nil)
	return err
}

func (s *QEMUService) Start(vmid int) error {
	return s.power(vmid, "start")
}

func (s *QEMUService) Stop(vmid int) error {
	return s.power(vmid, "stop")
}

func (s *QEMUService) Reset(vmid int) error {
	return s.power(vmid, "reset")
}

func (s *QEMUService) Shutdown(vmid int) error {
	return s.power(vmid, "shutdown")
}

func (s *QEMUService) Suspend(vmid int) error {
	return s.power(vmid, "suspend")
}

func (s *QEMUService) Resume(vmid int) error {
	return s.power(vmid, "resume")
}

func (s *QEMUService) Create() error {
	return errors.New("Not yet implemented")
}

func (s *QEMUService) Clone(vmid int, opts *VMCreateOptions) (*Task, error) {
	form := internal.StructToForm(opts, []string{"vm_c_n", "vm_n", "c_n", "n"})
	task, err := s.client.Post("nodes/"+s.node.Node+"/qemu/"+strconv.Itoa(vmid)+"/clone", form)
	if err != nil {
		return nil, err
	}

	return &Task{provider: s.node.Task, upid: task.(string)}, nil
}

func (s *QEMUService) Update(vmid int, cfg *QEMUConfig) error {
	norm := *cfg
	if !norm.MemoryBallooning {
		norm.MemoryMinimum = 0
	}

	form := internal.StructToForm(norm, []string{"n"})
	_, err := s.client.Put("nodes/"+s.node.Node+"/qemu/"+strconv.Itoa(vmid)+"/config", form)
	return err
}

func (s *QEMUService) Delete(vmid int) (*Task, error) {
	task, err := s.client.Delete("nodes/"+s.node.Node+"/qemu/"+strconv.Itoa(vmid), nil)
	if err != nil {
		return nil, err
	}

	return &Task{provider: s.node.Task, upid: task.(string)}, nil
}
