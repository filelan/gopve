package service

import (
	"errors"
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
	return nil, errors.New("Not yet implemented")
}

func (s *TaskService) Get(upid string) (*Task, error) {
	return nil, errors.New("Not yet implemented")
}

func (s *TaskService) Wait(upid string) error {
	ch := make(chan error, 1)
	go func() {
		for {
			data, err := s.client.Get("nodes/" + s.node.Node + "/tasks/" + upid + "/status")
			if err != nil {
				ch <- err
				return
			}

			val := data.(map[string]interface{})
			if val["status"].(string) == "stopped" {
				exitStatus := val["exitstatus"].(string)
				if exitStatus != "OK" {
					ch <- errors.New(exitStatus)
					return
				}

				ch <- nil
				return
			}

			<-time.After(time.Second)
		}
	}()

	res := <-ch
	close(ch)
	return res
}
