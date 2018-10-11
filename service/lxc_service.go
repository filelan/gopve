package service

import (
	"errors"
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
	Create() (*Task, error)
	Clone(int, *VMCreateOptions) (*Task, error)
	Update(int, *LXCConfig) error
	Delete(int) (*Task, error)
}

type LXCService struct {
	client *internal.Client
	node   *Node
}

type LXCServiceFactoryProvider interface {
	Create(*Node) LXCServiceProvider
}

type LXCServiceFactory struct {
	client    *internal.Client
	providers map[string]LXCServiceProvider
}

func NewLXCServiceFactoryProvider(c *internal.Client) LXCServiceFactoryProvider {
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
	for _, lxc := range data.(internal.JArray) {
		val := lxc.(internal.JObject)
		row := &LXC{
			provider: s,

			VMID:   internal.AsJInt(val, "vmid"),
			Name:   internal.JString(val, "name"),
			Status: internal.JString(val, "status"),
			LXCConfig: LXCConfig{
				CPU:         internal.JInt(val, "cpus"),
				MemoryTotal: internal.JInt(val, "maxmem"),
				MemorySwap:  internal.JInt(val, "maxswap"),
			},
		}

		res = append(res, row)
	}

	return &res, nil
}

func (s *LXCService) Get(vmid int) (*LXC, error) {
	dataStatus, err := s.client.Get("nodes/" + s.node.Node + "/lxc/" + strconv.Itoa(vmid) + "/status/current")
	if err != nil {
		return nil, err
	}

	dataConfig, err := s.client.Get("nodes/" + s.node.Node + "/lxc/" + strconv.Itoa(vmid) + "/config")
	if err != nil {
		return nil, err
	}

	valStatus := dataStatus.(internal.JObject)
	valConfig := dataConfig.(internal.JObject)

	res := &LXC{
		provider: s,

		VMID:        internal.AsJInt(valStatus, "vmid"),
		Name:        internal.JString(valStatus, "name"),
		Description: internal.JString(valConfig, "description"),
		Status:      internal.JString(valStatus, "status"),
		LXCConfig: LXCConfig{
			Architecture: internal.JString(valConfig, "arch"),
			OSType:       internal.JString(valConfig, "ostype"),
			CPU:          internal.JInt(valConfig, "cores"),
			CPULimit:     internal.JFloatDefault(valConfig, "cpulimit", LXCDefaultCPULimit),
			CPUUnits:     internal.JIntDefault(valConfig, "cpuunits", LXCDefaultCPUUnits),
			MemoryTotal:  internal.JInt(valConfig, "memory"),
			MemorySwap:   internal.JInt(valConfig, "swap"),
		},
	}

	rootMountPoint := internal.JString(valConfig, "rootfs")
	internal.KVToStruct(rootMountPoint, &res.LXCConfig.RootMountPoint)

	res.LXCConfig.MountPoints = make(LXCMountPointDict)
	for i := LXCMinimumMountPoint; i <= LXCMaximumMountPoint; i++ {
		mountPoint := internal.JStringDefault(valConfig, "mp" + strconv.Itoa(i), "")
		if mountPoint != "" {
			res.LXCConfig.MountPoints[i] = &LXCMountPoint{}
			internal.KVToStruct(mountPoint, res.LXCConfig.MountPoints[i])
		}
	}

	res.LXCConfig.NetworkDevices = make(LXCNetworkDeviceDict)
	for i := LXCMinimumNetworkDevice; i <= LXCMaximumNetworkDevice; i++ {
		networkDevice := internal.JStringDefault(valConfig, "net" + strconv.Itoa(i), "")
		if networkDevice != "" {
			res.LXCConfig.NetworkDevices[i] = &LXCNetworkDevice{}
			internal.KVToStruct(networkDevice, res.LXCConfig.NetworkDevices[i])
		}
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

func (s *LXCService) Create() (*Task, error) {
	return nil, errors.New("Not yet implemented")
}

func (s *LXCService) Clone(vmid int, opts *VMCreateOptions) (*Task, error) {
	form := internal.StructToForm(opts, []string{"ct_c_n", "ct_n", "c_n", "n"})
	task, err := s.client.Post("nodes/"+s.node.Node+"/lxc/"+strconv.Itoa(vmid)+"/clone", form)
	if err != nil {
		return nil, err
	}

	return &Task{provider: s.node.Task, upid: task.(string)}, nil
}

func (s *LXCService) Update(vmid int, cfg *LXCConfig) error {
	form := internal.StructToForm(cfg, []string{"n"})
	_, err := s.client.Put("nodes/"+s.node.Node+"/lxc/"+strconv.Itoa(vmid)+"/config", form)
	return err
}

func (s *LXCService) Delete(vmid int) (*Task, error) {
	task, err := s.client.Delete("nodes/"+s.node.Node+"/lxc/"+strconv.Itoa(vmid), nil)
	if err != nil {
		return nil, err
	}

	return &Task{provider: s.node.Task, upid: task.(string)}, nil
}
