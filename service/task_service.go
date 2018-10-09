package service

import (
	"time"

	"github.com/xabinapal/gopve/internal"
)

type TaskServiceProvider interface {
	List() (*TaskList, error)
	Get(string) (*Task, error)
	Wait(string) error
}

type TaskService struct {
	client *internal.Client
	node   *Node
}

type TaskServiceFactoryProvider interface {
	Create(*Node) TaskServiceProvider
}

type TaskServiceFactory struct {
	client    *internal.Client
	providers map[string]TaskServiceProvider
}

func NewTaskServiceFactoryProvider(c *internal.Client) TaskServiceFactoryProvider {
	return &TaskServiceFactory{
		client:    c,
		providers: make(map[string]TaskServiceProvider),
	}
}

func (factory *TaskServiceFactory) Create(node *Node) TaskServiceProvider {
	provider, ok := factory.providers[node.Node]
	if !ok {
		provider = &TaskService{
			client: factory.client,
			node:   node,
		}

		factory.providers[node.Node] = provider
	}

	return provider
}

func (s *TaskService) List() (*TaskList, error) {
	data, err := s.client.Get("nodes/" + s.node.Node + "/tasks")
	if err != nil {
		return nil, err
	}

	var res TaskList
	for _, task := range data.([]interface{}) {
		val := task.(map[string]interface{})
		row := &Task{
			provider: s,
			upid:     val["upid"].(string),
		}

		res = append(res, row)
	}

	return &res, nil
}

func (s *TaskService) Get(upid string) (*Task, error) {
	data, err := s.client.Get("nodes/" + s.node.Node + "/tasks/" + upid + "/status")
	if err != nil {
		return nil, err
	}

	val := data.(map[string]interface{})
	res := &Task{
		provider: s,
		filled:   true,
		upid:     upid,
		taskType: val["type"].(string),
		status:   val["status"].(string),
	}

	exitStatus, ok := val["exitstatus"]
	if ok {
		res.exitStatus = exitStatus.(string)
	}

	return res, nil
}

func (s *TaskService) Wait(upid string) error {
	ch := make(chan error, 1)
	go func() {
		defer close(ch)
		for {
			data, err := s.client.Get("nodes/" + s.node.Node + "/tasks/" + upid + "/status")
			if err != nil {
				ch <- err
				return
			}

			val := data.(map[string]interface{})
			if val["status"].(string) == "stopped" {
				ch <- nil
				return
			}

			<-time.After(time.Second)
		}
	}()

	res := <-ch
	return res
}
