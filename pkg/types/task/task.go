package task

type Task interface {
	Node() string
	UPID() string
	Wait() error
}
