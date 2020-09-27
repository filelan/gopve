package storage

import "github.com/xabinapal/gopve/pkg/types/errors"

const (
	ErrInvalidKind = errors.ClientError("unsupported storage type")
)

var ErrMissingProperty = errors.NewKeyedClientError("500 - missing property!", nil)
