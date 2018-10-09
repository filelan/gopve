package service

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/xabinapal/gopve/internal"
)

type LXCServiceProvider interface {
	List() (*LXCList, error)
	Get(int) (*LXC, error)
	Start(int) error
	Stop(int) error
	Reset(int) error
	Shutdown(int) error
	Suspend(int) error
	Resume(int) error
	Create() error
	Update(int, *LXCConfig) error
	Delete() error
	Clone(int, bool, *VMCreateOptions) error
}

type LXCService struct {
	client *internal.Client
	node   *Node
}

type LXCServiceProviderFactory interface {
	Create(*Node) LXCServiceProvider
}

type LXCServiceFactory struct {
	client    *internal.Client
	providers map[string]LXCServiceProvider
}

func NewLXCServiceProviderFactory(c *internal.Client) LXCServiceProviderFactory {
	return &LXCServiceFactory{
		client:    c,
		providers: make(map[string]LXCServiceProvider),
	}
}

func (factory *LXCServiceFactory) Create(node *Node) LXCServiceProvider {
	provider, ok := factory.providers[node.Node]
	if !ok {
		provider = &LXCService{
			client: factory.client,
			node:   node,
		}

		factory.providers[node.Node] = provider
	}

	return provider
}

func (s *LXCService) List() (*LXCList, error) {
	data, err := s.client.Get("nodes/" + s.node.Node + "/lxc")
	if err != nil {
		return nil, err
	}

	var res LXCList
	for _, lxc := range data.([]interface{}) {
		val := lxc.(map[string]interface{})
		vmid, _ := strconv.Atoi(val["vmid"].(string))
		row := &LXC{
			provider: s,

			VMID:   vmid,
			Name:   val["name"].(string),
			Status: val["status"].(string),
			LXCConfig: LXCConfig{
				CPU:         int(val["cpus"].(float64)),
				MemoryTotal: int(val["maxmem"].(float64)),
				MemorySwap:  int(val["maxswap"].(float64)),
			},
		}

		res = append(res, row)
	}

	return &res, nil
}

func (s *LXCService) Get(vmid int) (*LXC, error) {
	dataConfig, err := s.client.Get("nodes/" + s.node.Node + "/lxc/" + strconv.Itoa(vmid) + "/config")
	if err != nil {
		return nil, err
	}

	dataStatus, err := s.client.Get("nodes/" + s.node.Node + "/lxc/" + strconv.Itoa(vmid) + "/status/current")
	if err != nil {
		return nil, err
	}

	valConfig := dataConfig.(map[string]interface{})
	valStatus := dataStatus.(map[string]interface{})

	res := &LXC{
		provider: s,

		VMID:   vmid,
		Name:   valStatus["name"].(string),
		Status: valStatus["status"].(string),
		LXCConfig: LXCConfig{
			CPU:         int(valConfig["cores"].(float64)),
			MemoryTotal: int(valConfig["memory"].(float64)),
			MemorySwap:  int(valConfig["swap"].(float64)),
		},
	}

	cpuLimit, ok := valConfig["cpulimit"]
	if ok {
		cpuLimit, _ := strconv.Atoi(cpuLimit.(string))
		res.CPULimit = cpuLimit
	} else {
		res.CPULimit = LXC_DEFAULT_CPU_LIMIT
	}

	cpuUnits, ok := valConfig["cpuunits"]
	if ok {
		res.CPUUnits = int(cpuUnits.(float64))
	} else {
		res.CPUUnits = LXC_DEFAULT_CPU_UNITS
	}

	return res, nil
}

func (s *LXCService) power(vmid int, command string) error {
	_, err := s.client.Post("nodes/"+s.node.Node+"/lxc/"+strconv.Itoa(vmid)+"/status/"+command, nil)
	return err
}

func (s *LXCService) Start(vmid int) error {
	return s.power(vmid, "start")
}

func (s *LXCService) Stop(vmid int) error {
	return s.power(vmid, "stop")
}

func (s *LXCService) Reset(vmid int) error {
	return s.power(vmid, "reset")
}

func (s *LXCService) Shutdown(vmid int) error {
	return s.power(vmid, "shutdown")
}

func (s *LXCService) Suspend(vmid int) error {
	return s.power(vmid, "suspend")
}

func (s *LXCService) Resume(vmid int) error {
	return s.power(vmid, "resume")
}

func (s *LXCService) Create() error {
	return errors.New("Not yet implemented")
}

func (s *LXCService) Update(vmid int, cfg *LXCConfig) error {
	return errors.New("Not yet implemented")
}

func (s *LXCService) Delete() error {
	return errors.New("Not yet implemented")
}

func (s *LXCService) Clone(vmid int, full bool, opts *VMCreateOptions) error {
	form := url.Values{}
	form.Set("full", internal.BoolToForm(full))
	internal.AddStructToForm(&form, opts, []string{"ct_c_n", "ct_n", "c_n", "n"})

	_, err := s.client.Post("nodes/"+s.node.Node+"/lxc/"+strconv.Itoa(vmid)+"/clone", form)
	if err != nil {
		return err
	}

	return nil
}
