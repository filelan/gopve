package node

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xabinapal/gopve/pkg/request"
)

type getTimeResponseJSON struct {
	LocalTime int64  `json:"localtime"`
	UTCTime   int64  `json:"time"`
	Timezone  string `json:"timezone"`
}

func newTimeWithTimezone(timestamp int64, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}

	t := time.Unix(timestamp, 0).In(loc)

	return t, nil
}

func (node *Node) getTime() (*getTimeResponseJSON, error) {
	var res getTimeResponseJSON

	err := node.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/time", node.Name()), nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (node *Node) GetTime(local bool) (time.Time, error) {
	res, err := node.getTime()
	if err != nil {
		return time.Time{}, err
	}

	timezone := "UTC"
	if local {
		timezone = res.Timezone
	}

	return newTimeWithTimezone(res.UTCTime, timezone)
}

func (node *Node) GetTimezone() (*time.Location, error) {
	res, err := node.getTime()
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation(res.Timezone)
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func (node *Node) SetTimezone(timezone *time.Location) error {
	err := node.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/time", node.Name()), request.Values{
		"timezone": {timezone.String()},
	}, nil)

	return err
}
