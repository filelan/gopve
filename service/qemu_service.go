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
	for _, qemu := range data.(internal.JArray) {
		val := qemu.(internal.JObject)
		row := &QEMU{
			provider: s,

			VMID:   internal.AsJInt(val, "vmid"),
			Name:   internal.JString(val, "name"),
			Status: internal.JString(val, "status"),
			QEMUConfig: QEMUConfig{
				CPU:           internal.JInt(val, "cpus"),
				MemoryTotal:   internal.JInt(val, "maxmem"),
				MemoryMinimum: internal.JIntDefault(val, "balloon_min", 0),
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

	valConfig := dataConfig.(internal.JObject)
	valStatus := dataStatus.(internal.JObject)

	res := &QEMU{
		provider: s,

		VMID:   internal.AsJInt(valStatus, "vmid"),
		Name:   internal.JString(valStatus, "name"),
		Status: internal.JString(valStatus, "status"),
		QEMUConfig: QEMUConfig{
			OSType:        internal.JString(valConfig, "ostype"),
			CPUSockets:    internal.JInt(valConfig, "sockets"),
			CPUCores:      internal.JInt(valConfig, "cores"),
			CPULimit:      internal.JIntDefault(valConfig, "cpulimit", QEMUDefaultCPULimit),
			CPUUnits:      internal.JIntDefault(valConfig, "cpuunits", QEMUDefaultCPUUnits),
			MemoryTotal:   internal.JInt(valConfig, "memory"),
			MemoryMinimum: internal.JInt(valConfig, "balloon"),
			IsNUMAAware:   internal.JBoolean(valConfig, "numa"),
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
