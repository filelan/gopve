package types

type Task interface {
	Node() string
	UPID() string
	Wait() error
}

type TaskStatus string

const (
	TaskRunning TaskStatus = "running"
	TaskStopped TaskStatus = "stopped"
)

func (ts TaskStatus) IsValid() bool {
	return true
}
