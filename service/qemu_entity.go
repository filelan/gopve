package service

import (
	"strings"

	"github.com/xabinapal/gopve/internal"
)

type QEMU struct {
	provider QEMUServiceProvider

	VMID   int
	Name   string
	Status string
	QEMUConfig
}

type QEMUList = []*QEMU

type QEMUConfig struct {
	Architecture     QEMUArchitecture      `n:"cpu"`
	OSType           string                `n:"ostype"`
	CPU              int                   `i:"always"`
	CPUSockets       int                   `n:"sockets"`
	CPUCores         int                   `n:"cores"`
	CPULimit         float64               `n:"cpulimit"`
	CPUUnits         int                   `n:"cpuunits"`
	MemoryTotal      int                   `n:"memory"`
	MemoryMinimum    int                   `n:"balloon"`
	MemoryBallooning bool                  `i:"always"`
	IDEVolumes       QEMUVolumeDeviceDict  `n:"ide"`
	SATAVolumes      QEMUVolumeDeviceDict  `n:"sata"`
	SCSIVolumes      QEMUVolumeDeviceDict  `n:"scsi"`
	VIRTIOVolumes    QEMUVolumeDeviceDict  `n:"virtio"`
	NetworkDevices   QEMUNetworkDeviceDict `n:"net"`
	HasQEMUAgent     bool                  `n:"agent"`
	IsNUMAAware      bool                  `n:"numa"`
}

type QEMUArchitecture struct {
	Type   string   `n:",cputype"`
	Flags  []string `n:"flags"`
	Hidden bool     `n:"hidden"`
}

type QEMUVolumeDevice struct {
	Volume         string `n:",file"`
	Format         string `n:"format"`
	Media          string `n:"media" d:"disk"`
	Size           string `n:"size"`
	HasBackup      bool   `n:"backup" d:"false"`
	HasReplication bool   `n:"replicate" d:"true"`
	IsShared       bool   `n:"shared" d:"false"`
}

type QEMUVolumeDeviceDict = map[int]*QEMUVolumeDevice

type QEMUNetworkDevice struct {
	Model       string `n:",model"`
	MACAddress  string `n:"macaddr"`
	Bridge      string `n:"bridge"`
	Queues      int    `n:"queues"`
	Rate        int    `n:"rate"`
	VLANTag     int    `n:"tag"`
	Trunks      []int  `n:"trunks"`
	HasFirewall bool   `n:"firewall"`
	IsLinkDown  bool   `n:"link_down"`
}

type QEMUNetworkDeviceDict = map[int]*QEMUNetworkDevice

type QEMUCloudInitConfig struct {
}

const (
	QEMUDefaultArchitecture = "kvm64"
	QEMUDefaultCPULimit      = 0.0
	QEMUDefaultCPUUnits      = 1024
	QEMUMinimumIDEDevice     = 0
	QEMUMaximumIDEDevice     = 3
	QEMUMinimumSATADevice    = 0
	QEMUMaximumSATADevice    = 5
	QEMUMinimumSCSIDevice    = 0
	QEMUMaximumSCSIDevice    = 13
	QEMUMinimumVIRTIODevice  = 0
	QEMUMaximumVIRTIODevice  = 15
	QEMUMinimumNetworkDevice = 0
	QEMUMaximumNetworkDevice = 9
)

// QEMU OS Types
const (
	Other         = "other"
	WindowsXP     = "wxp"
	Windows2000   = "w2k"
	Windows2003   = "w2k3"
	Windows2008   = "w2k8"
	WindowsVista  = "wvista"
	Windows7      = "win7"
	Windows8      = "win8"
	Windows2012   = "win8"
	Windows2012r2 = "win8"
	Windows10     = "win10"
	Windows2016   = "win10"
	Linux24       = "l24"
	Linux26       = "l26"
	Linux3x       = "l26"
	Linux4x       = "l26"
	Solaris       = "solaris"
	OpenSolaris   = "solaris"
	OpenIndiania  = "solaris"
)

var QEMUVolumeTypes = []string{"cloop", "cow", "qcow", "qcow2", "qed", "raw", "vmdk"}

func qemuVolumeTypeHelper(meta internal.StructMetaDict) {
	if !meta["format"].IsSet {
		volume := meta["file"].Get().(string)
		for _, v := range QEMUVolumeTypes {
			if strings.HasSuffix(volume, "."+v) {
				meta["format"].Set(v)
				return
			}
		}
		meta["format"].Set("raw")
	}
}

func (e QEMU) Start() error {
	return e.provider.Start(e.VMID)
}

func (e *QEMU) Stop() error {
	return e.provider.Stop(e.VMID)
}

func (e *QEMU) Reset() error {
	return e.provider.Reset(e.VMID)
}

func (e *QEMU) Shutdown() error {
	return e.provider.Shutdown(e.VMID)
}

func (e *QEMU) Suspend() error {
	return e.provider.Suspend(e.VMID)
}

func (e *QEMU) Resume() error {
	return e.provider.Resume(e.VMID)
}
