package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/node"
)

type getSyslogJSON struct {
	LineNumber int    `json:"n"`
	Contents   string `json:"t"`
}

func (n *Node) GetSyslog(
	opts node.GetSyslogOptions,
) (node.SyslogEntries, error) {
	var form request.Values

	form.ConditionalAddUint("start", opts.LineStart, opts.LineStart != 0)
	form.ConditionalAddUint("limit", opts.LineLimit, opts.LineLimit != 0)
	form.ConditionalAddTime("since", opts.Since, !opts.Since.IsZero())
	form.ConditionalAddTime("until", opts.Until, !opts.Until.IsZero())
	form.ConditionalAddString("service", opts.Service, opts.Service != "")

	var res []getSyslogJSON
	if err := n.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/syslog", n.name), form, &res); err != nil {
		return nil, err
	}

	entries := make(node.SyslogEntries)
	for _, entry := range res {
		entries[uint(entry.LineNumber)] = entry.Contents
	}

	return entries, nil
}
