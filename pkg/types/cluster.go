package types

type ClusterService interface {
	HA() HighAvailabilityService
}

type Cluster interface {
}
