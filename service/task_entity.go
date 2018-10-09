package service

type Task struct {
	provider TaskServiceProvider

	UPID       string
	Type       string
	Status     string
	ExitStatus string
}

type TaskList []*Task

func (e *Task) Wait() {
	e.provider.Wait(e.UPID)
}
