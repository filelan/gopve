package vm

import "github.com/xabinapal/gopve/pkg/types/errors"

const (
	ErrNotFound = errors.ClientError("404 - virtual machine not found!")

	ErrNoSnapshot         = errors.ClientError("500 - snapshot not found!")
	ErrRootParentSnapshot = errors.ClientError(
		"500 - snapshot has no parent!",
	)
	ErrUpdateCurrentSnapshot = errors.ClientError(
		"500 - can't set properties of current snapshot!",
	)
	ErrDeleteCurrentSnapshot = errors.ClientError(
		"500 - can't delete current snapshot!",
	)
)
