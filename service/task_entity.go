package service

type Task struct {
	provider TaskServiceProvider

	UPID       string `n:"upid"`
	TaskType   string `n:"type"`
	Status     string `n:"status"`
	ExitStatus string `n:"exitstatus"`
}

type TaskList []*Task

func (e *Task) Update() error {
	new, err := e.provider.Get(e.UPID)
	if err == nil {
		e.TaskType = new.TaskType
		e.Status = new.Status
		e.ExitStatus = new.ExitStatus
	}
	return err
}

func (e *Task) Wait() {
	err := e.provider.Wait(e.UPID)
	if err == nil {
		e.Update()
	}
}
