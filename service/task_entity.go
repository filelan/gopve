package service

type Task struct {
	provider TaskServiceProvider
	filled   bool

	upid       string
	taskType   string
	status     string
	exitStatus string
}

type TaskList []*Task

func (e *Task) UPID() (string, error) {
	return e.upid, nil
}

func (e *Task) Type() (string, error) {
	if !e.filled {
		if err := e.Update(); err != nil {
			return "", err
		}
	}
	return e.taskType, nil
}

func (e *Task) Status() (string, error) {
	if !e.filled {
		if err := e.Update(); err != nil {
			return "", err
		}
	}
	return e.status, nil
}

func (e *Task) ExitStatus() (string, error) {
	if !e.filled {
		if err := e.Update(); err != nil {
			return "", err
		}
	}
	return e.exitStatus, nil
}

func (e *Task) Update() error {
	new, err := e.provider.Get(e.upid)
	if err == nil {
		e.filled = true
		e.taskType = new.taskType
		e.status = new.status
		e.exitStatus = new.exitStatus
	}
	return err
}

func (e *Task) Wait() {
	err := e.provider.Wait(e.upid)
	if err == nil {
		e.Update()
	}
}
