package service

import (
	"strings"
)

type QEMU struct {
	provider QEMUServiceProvider

	VMID         int    `n:"vmid"`
	Name         string `n:"name"`
	Description  string `n:"description"`
	Status       string `n:"status"`
	ConfigDigest string `n:"digest"`
	QEMUConfig
	QEMUCloudInitConfig
}

type QEMUConfig struct {
	MachineType                  string                `n:"machine"`
	Architecture                 QEMUArchitecture      `n:"cpu" t:"kv"`
	OSType                       string                `n:"ostype"`
	BootArguments                string                `n:"args"`
	CPU                          int                   `h:"true"`
	CPUSockets                   int                   `n:"sockets"`
	CPUCores                     int                   `n:"cores"`
	CPULimit                     float64               `n:"cpulimit" d:"0.0"`
	CPUUnits                     int                   `n:"cpuunits" d:"1024"`
	MemoryTotal                  int                   `n:"memory"`
	MemoryBallooning             bool                  `h:"true"`
	MemoryMinimum                int                   `n:"balloon"`
	MemoryShares                 int                   `n:"shares" d:"1000"`
	HugePages                    string                `n:"hugepages"`
	SCSIController               string                `n:"scsihw"`
	EFIDisk                      QEMUEFIVolumeDevice   `n:"efidisk0" t:"kv"`
	IDEVolumes                   QEMUVolumeDeviceDict  `n:"ide" t:"kvdict" min:"0" max:"3"`
	SATAVolumes                  QEMUVolumeDeviceDict  `n:"sata" t:"kvdict" min:"0" max:"5"`
	SCSIVolumes                  QEMUVolumeDeviceDict  `n:"scsi" t:"kvdict" min:"0" max:"13"`
	VIRTIOVolumes                QEMUVolumeDeviceDict  `n:"virtio" t:"kvdict" min:"0" max:"15"`
	NetworkDevices               QEMUNetworkDeviceDict `n:"net" t:"kvdict" min:"0" max:"9"`
	SerialDevices                map[int]string        `n:"serial" t:"dict" min:"0" max:"3"`
	BiosType                     string                `n:"bios" d:"seabios"`
	KeyboardLayout               string                `n:"keyboard"`
	StartOnBoot                  bool                  `n:"onboot" d:"0"`
	RestartOnCrash               bool                  `n:"autostart" d:"0"`
	FreezeAtBoot                 bool                  `n:"freeze" d:"0"`
	LockType                     string                `n:"lock"`
	BootOrder                    []string              `n:"boot" d:"cdn" h:"true"`
	BootDisk                     string                `n:"bootdisk"`
	HotPlug                      []string              `n:"hotplug" d:"network,disk,usb" s:","`
	MigrateMaxDowntime           float64               `n:"migrate_downtime" d:"0.1"`
	MigrateMaxMBps               int                   `n:"migrate_speed" d:"0"`
	HasAutoStart                 bool                  `n:"autostart"`
	HasQEMUAgent                 bool                  `n:"agent" d:"0"`
	HasKVMHardwareVirtualization bool                  `n:"kvm" d:"1"`
	IsRebootAllowed              bool                  `n:"reboot"`
	IsProtected                  bool                  `n:"protection"`
	IsACPIEnabled                bool                  `n:"acpi" d:"1"`
	IsNUMAAware                  bool                  `n:"numa"`
	IsRTCLocalTime               bool                  `n:"localtime"`
}

func (e *QEMUConfig) CPUHelper() int {
	return e.CPUSockets * e.CPUCores
}

func (e *QEMUConfig) MemoryBallooningHelper() bool {
	return e.MemoryMinimum != 0
}

func (e *QEMUConfig) BootOrderHelper(v string) []string {
	var val = make([]string, 0)
	for _, v := range v {
		switch v {
		case 'a':
			val = append(val, "floppy")
		case 'c':
			val = append(val, "disk")
		case 'd':
			val = append(val, "cdrom")
		case 'n':
			val = append(val, "network")
		}
	}

	return val
}

func (e *QEMUConfig) UnmarshalHelper() {
	if !e.MemoryBallooning {
		e.MemoryMinimum = e.MemoryTotal
	}
}

type QEMUArchitecture struct {
	Type   string   `n:",cputype"`
	Flags  []string `n:"flags" d:"" s:";"`
	Hidden bool     `n:"hidden"`
}

type QEMUVolumeDevice struct {
	ModelName                 string `n:"model"`
	ModelSerialNumber         string `n:"serial"`
	Volume                    string `n:",file"`
	Format                    string `n:"format"`
	Media                     string `n:"media" d:"disk"`
	Size                      string `n:"size"`
	CacheType                 string `n:"cache" default:"none"`
	ReadErrorAction           string `n:"rerror"`
	WriteErrorAction          string `n:"werror"`
	AIOType                   string `n:"aio"`
	MaxRWSpeed                int    `n:"bps"`
	MaxRWBurstSpeedSeconds    int    `n:"bps_max_length"`
	MaxRWIOPS                 int    `n:"iops"`
	MaxRWUnthrottledIOPS      int    `n:"iops_max"`
	MaxRWBurstIOPS            int    `n:"iops_max_length"`
	MaxReadSpeed              int    `n:"bps_rd"`
	MaxReadBurstSpeedSeconds  int    `n:"bps_rd_max_length"`
	MaxReadIOPS               int    `n:"iops_rd"`
	MaxReadUnthrottledIOPS    int    `n:"iops_rd_max"`
	MaxReadBurstIOPS          int    `n:"iops_rd_max_length"`
	MaxWriteSpeed             int    `n:"bps_wr"`
	MaxWriteBurstSpeedSeconds int    `n:"bps_wr_max_length"`
	MaxWriteIOPS              int    `n:"iops_wr"`
	MaxWriteUnthrottledIOPS   int    `n:"iops_wr_max"`
	MaxWriteBurstIOPS         int    `n:"iops_wr_max_length"`
	GeometryMode              string `n:"trans"`
	CylinderCount             int    `n:"cyls"`
	SectorCount               int    `n:"secs"`
	HeadCount                 int    `n:"heads"`
	DetectZeroes              bool   `n:"detect_zeroes" d:"0"`
	PassTrimRequests          bool   `n:"discard" h:"true"`
	HasBackup                 bool   `n:"backup" d:"0"`
	HasReplication            bool   `n:"replicate" d:"1"`
	IsShared                  bool   `n:"shared" d:"0"`
	IsSnapshotMode            bool   `n:"snapshot" d:"0"`
}

var QEMUVolumeTypes = []string{"cloop", "cow", "qcow", "qcow2", "qed", "raw", "vmdk"}

func (e *QEMUVolumeDevice) PassTrimRequestsHelper(v string) bool {
	return v == "on"
}

func (e *QEMUVolumeDevice) UnmarshalHelper() {
	if e.Format == "" {
		for _, v := range QEMUVolumeTypes {
			if strings.HasSuffix(e.Volume, "."+v) {
				e.Format = v
				break
			}
		}

		if e.Format == "" {
			e.Format = "raw"
		}
	}
}

type QEMUVolumeDeviceDict = map[int]*QEMUVolumeDevice

type QEMUEFIVolumeDevice struct {
	Volume string `n:",file"`
	Format string `n:"format"`
	Size   string `n:"size"`
}

type QEMUEFIVolumeDeviceDict = map[int]*QEMUEFIVolumeDevice

type QEMUNetworkDevice struct {
	Model       string `n:",model"`
	MACAddress  string `n:"macaddr"`
	Bridge      string `n:"bridge"`
	Queues      int    `n:"queues"`
	Rate        int    `n:"rate"`
	VLANTag     int    `n:"tag"`
	Trunks      []int  `n:"trunks" d:"" s:";"`
	HasFirewall bool   `n:"firewall"`
	IsLinkDown  bool   `n:"link_down"`
}

type QEMUNetworkDeviceDict = map[int]*QEMUNetworkDevice

type QEMUCloudInitConfig struct {
	Format       string                   `n:"citype"`
	User         string                   `n:"ciuser"`
	Password     string                   `n:"cipassword"`
	IPConfig     QEMUCloudInitNetworkDict `n:"ipconfig"`
	NameServer   string                   `n:"nameserver"`
	SearchDomain string                   `n:"searchdomain"`
	SSHKeys      []string                 `n:"sshkeys" s:"\n"`
}

type QEMUCloudInitNetwork struct {
	IPAddress   string `n:"ip" d:"dhcp"`
	Gateway     string `n:"gw"`
	IPv6Address string `n:"ip6" d:"dhcp"`
	GatewayIPv6 string `n:"gw6"`
}

type QEMUCloudInitNetworkDict = map[int]*QEMUCloudInitNetwork

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
