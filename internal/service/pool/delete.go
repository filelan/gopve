package pool

import (
	"fmt"
	"net/http"
)

func (svc *Service) Delete(name string) error {
	return svc.client.Request(http.MethodDelete, fmt.Sprintf("pools/%s", name), nil, nil)
}
