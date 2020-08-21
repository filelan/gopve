package utils

type PVETask struct {
	node string
	upid string
	ch   <-chan error
}

func NewPVETask(node string, upid string, ch <-chan error) *PVETask {
	return &PVETask{
		node: node,
		upid: upid,
		ch:   ch,
	}
}

func (t *PVETask) Node() string {
	return t.node
}

func (t *PVETask) UPID() string {
	return t.upid
}

func (t *PVETask) Wait() error {
	return <-t.ch
}
