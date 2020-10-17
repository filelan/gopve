package qemu

type NetworkModel string

const (
	NetworkModelIntelE1000     NetworkModel = "e1000"
	NetworkModelVirtIO         NetworkModel = "VirtIO"
	NetworkModelRealtekRTL8139 NetworkModel = "rtl8139"
	NetworkModelVMwareVMXNET3  NetworkModel = "vmxnet3"
)
