package service

import (
	"errors"
	"strconv"

	"github.com/xabinapal/gopve/internal"
)

type LXCServiceProvider interface {
	List() (*LXCList, error)
	Get(int) (*LXC, error)
	Create() error
	Update() error
	Delete() error
	Clone() error
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
		ctid, _ := strconv.Atoi(val["vmid"].(string))
		row := &LXC{
			provider: s,

			CTID:        ctid,
			Name:        val["name"].(string),
			Status:      val["status"].(string),
			CPU:         int(val["cpus"].(float64)),
			MemoryTotal: int(val["maxmem"].(float64)),
			MemorySwap:  int(val["maxswap"].(float64)),
		}

		res = append(res, row)
	}

	return &res, nil
}

func (s *LXCService) Get(id int) (*LXC, error) {
	data, err := s.client.Get("nodes/" + s.node.Node + "/lxc/" + strconv.Itoa(id) + "/status/current")
	if err != nil {
		return nil, err
	}

	val := data.(map[string]interface{})
	ctid, _ := strconv.Atoi(val["vmid"].(string))
	res := &LXC{
		provider: s,

		CTID:        ctid,
		Name:        val["name"].(string),
		Status:      val["status"].(string),
		CPU:         int(val["cpus"].(float64)),
		MemoryTotal: int(val["maxmem"].(float64)),
		MemorySwap:  int(val["maxswap"].(float64)),
	}

	return res, nil
}

func (s *LXCService) Create() error {
	return errors.New("Not yet implemented")
}

func (s *LXCService) Update() error {
	return errors.New("Not yet implemented")
}

func (s *LXCService) Delete() error {
	return errors.New("Not yet implemented")
}

func (s *LXCService) Clone() error {
	return errors.New("Not yet implemented")
}
