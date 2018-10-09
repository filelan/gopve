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
	Update() error
	Delete() error
	Clone() error
}

type QEMUService struct {
	client *internal.Client
	node   *Node
}

type QEMUServiceProviderFactory interface {
	Create(*Node) QEMUServiceProvider
}

type QEMUServiceFactory struct {
	client    *internal.Client
	providers map[string]QEMUServiceProvider
}

func NewQEMUServiceProviderFactory(c *internal.Client) QEMUServiceProviderFactory {
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
	for _, qemu := range data.([]interface{}) {
		val := qemu.(map[string]interface{})
		vmid, _ := strconv.Atoi(val["vmid"].(string))
		row := &QEMU{
			provider: s,

			VMID:   vmid,
			Name:   val["name"].(string),
			Status: val["status"].(string),
			QEMUConfig: QEMUConfig{
				CPU:         int(val["cpus"].(float64)),
				MemoryTotal: int(val["maxmem"].(float64)),
			},
		}

		ballooning, ok := val["balloon_min"]
		if ok {
			row.MemoryMinimum = int(ballooning.(float64))
			row.MemoryBallooning = true
		} else {
			row.MemoryMinimum = row.MemoryTotal
			row.MemoryBallooning = false
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

	valConfig := dataConfig.(map[string]interface{})
	valStatus := dataStatus.(map[string]interface{})

	res := &QEMU{
		provider: s,

		VMID:   vmid,
		Name:   valStatus["name"].(string),
		Status: valStatus["status"].(string),
		QEMUConfig: QEMUConfig{
			CPUSockets:  int(valConfig["sockets"].(float64)),
			CPUCores:    int(valConfig["cores"].(float64)),
			MemoryTotal: int(valConfig["memory"].(float64)),
		},
	}

	res.CPU = res.CPUSockets * res.CPUCores

	cpuLimit, ok := valConfig["cpulimit"]
	if ok {
		cpuLimit, _ := strconv.Atoi(cpuLimit.(string))
		res.CPULimit = cpuLimit
	} else {
		res.CPULimit = QEMU_DEFAULT_CPU_LIMIT
	}

	cpuUnits, ok := valConfig["cpuunits"]
	if ok {
		res.CPUUnits = int(cpuUnits.(float64))
	} else {
		res.CPUUnits = QEMU_DEFAULT_CPU_UNITS
	}

	ballooning := int(valConfig["balloon"].(float64))
	if ballooning == 0 {
		res.MemoryMinimum = res.MemoryTotal
		res.MemoryBallooning = false
	} else {
		res.MemoryMinimum = ballooning
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

func (s *QEMUService) Update() error {
	return errors.New("Not yet implemented")
}

func (s *QEMUService) Delete() error {
	return errors.New("Not yet implemented")
}

func (s *QEMUService) Clone() error {
	return errors.New("Not yet implemented")
}
