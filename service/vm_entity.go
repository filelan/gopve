package service

type VMCreateOptions struct {
	VMID        int    `n:"vmid" c_n:"newid"`
	Node        string `n:"node" c_n:"target" i:"true"`
	Storage     string `n:"storage" i:"true"`
	Name        string `n:"name" ct_n:"hostname" i:"true"`
	Description string `n:"description" i:"true"`
	Pool        string `n:"pool" i:"true"`
	Snapshot    string `n:"snapname" i:"true"`
}
