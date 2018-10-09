package service

type VMCreateOptions struct {
	VMID        int    `n:"vmid" c_n:"newid" i:"f"`
	Node        string `n:"node" c_n:"target"`
	Storage     string `n:"storage"`
	Name        string `n:"name" ct_n:"hostname"`
	Description string `n:"description"`
	Pool        string `n:"pool"`
	Snapshot    string `n:"snapname"`
}
