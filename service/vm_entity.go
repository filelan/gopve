package service

type VMCreateOptions struct {
	VMID        int    `n:"vmid" c_n:"newid"`
	Node        string `n:"node" c_n:"target" i:"default"`
	Storage     string `n:"storage" i:"default"`
	Name        string `vm_n:"name" ct_n:"hostname" i:"default"`
	Description string `n:"description" i:"default"`
	Pool        string `n:"pool" i:"default"`
	FullClone   bool   `n:"full" i:"default"`
	Snapshot    string `n:"snapname" i:"default"`
}
