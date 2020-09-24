package pool

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/pool"
)

type getResponseJSON struct {
	Name        string `json:"poolid"`
	Description string `json:"comment"`
	Members     []struct {
		ID   string          `json:"id"`
		Type pool.MemberKind `json:"type"`
	} `json:"members"`
}

func (res getResponseJSON) Map(svc *Service, name string, full bool) (pool.Pool, error) {
	if full {
		var members []pool.PoolMember

		for _, m := range res.Members {
			switch m.Type {
			case pool.MemberKindVirtualMachine:
				member, err := NewPoolMemberVirtualMachine(svc, m.ID)
				if err != nil {
					return nil, err
				}

				members = append(members, member)
			case pool.MemberKindStorage:
				member, err := NewPoolMemberStorage(svc, m.ID)
				if err != nil {
					return nil, err
				}

				members = append(members, member)
			default:
				return nil, fmt.Errorf("unsupported pool member type")
			}
		}

		return NewFullPool(svc, name, res.Description, members), nil
	}

	return NewPool(svc, name, res.Description), nil
}

func (svc *Service) List() ([]pool.Pool, error) {
	var res []getResponseJSON
	if err := svc.client.Request(http.MethodGet, "pools", nil, &res); err != nil {
		return nil, err
	}

	pools := make([]pool.Pool, len(res))

	for i, pool := range res {
		out, err := pool.Map(svc, pool.Name, false)
		if err != nil {
			return nil, err
		}

		pools[i] = out
	}

	return pools, nil
}

func (svc *Service) Get(name string) (pool.Pool, error) {
	var res getResponseJSON
	if err := svc.client.Request(http.MethodGet, fmt.Sprintf("pools/%s", name), nil, &res); err != nil {
		return nil, err
	}

	return res.Map(svc, name, true)
}
