package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

type getSyslogJSON struct {
	LineNumber int    `json:"n"`
	Contents   string `json:"t"`
}

func (node *Node) GetSyslog(opts types.NodeGetSyslogOptions) (types.LogEntries, error) {
	var form utils.RequestValues

	form.ConditionalAddUint("start", opts.LineStart, opts.LineStart != 0)
	form.ConditionalAddUint("limit", opts.LineLimit, opts.LineLimit != 0)
	form.ConditionalAddTime("since", opts.Since, !opts.Since.IsZero())
	form.ConditionalAddTime("until", opts.Until, !opts.Until.IsZero())
	form.ConditionalAddString("service", opts.Service, opts.Service != "")

	var res []getSyslogJSON
	if err := node.svc.Client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/syslog", node.name), form, &res); err != nil {
		return nil, err
	}

	entries := make(types.LogEntries)
	for _, entry := range res {
		entries[entry.LineNumber] = entry.Contents
	}

	return entries, nil
}
