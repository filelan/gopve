package cluster

import "github.com/xabinapal/gopve/pkg/types/errors"

const (
	ErrNotInCluster = errors.ClientError(
		"500 - node is not in a cluster, no join info available!",
	)
	ErrAlreadyInCluster = errors.ClientError(
		"500 - node is already in a cluster!",
	)
)
