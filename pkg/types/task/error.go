package task

import "github.com/xabinapal/gopve/pkg/types/errors"

const (
	ErrInvalidUPID = errors.ClientError("500 - invalid task upid!")
	ErrInvalidKind = errors.ClientError("500 - invalid task kind!")
	ErrInvalidID   = errors.ClientError("500 - invalid task id!")
)
