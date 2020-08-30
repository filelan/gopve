package task

type Status string

const (
	StatusRunning Status = "running"
	StatusStopped Status = "stopped"
)

func (ts Status) IsValid() bool {
	switch ts {
	case StatusRunning, StatusStopped:
		return true

	default:
		return false
	}
}
