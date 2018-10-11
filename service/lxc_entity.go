package service

type LXC struct {
	provider LXCServiceProvider

	VMID        int
	Name        string
	Description string
	Status      string
	LXCConfig
}

type LXCList = []*LXC

type LXCConfig struct {
	Architecture   string               `n:"arch"`
	OSType         string               `n:"ostype"`
	CPU            int                  `n:"cores"`
	CPULimit       float64              `n:"cpulimit"`
	CPUUnits       int                  `n:"cpuunits"`
	MemoryTotal    int                  `n:"memory"`
	MemorySwap     int                  `n:"swap"`
	RootMountPoint LXCMountPoint        `n:"rootfs"`
	MountPoints    LXCMountPointDict    `n:"mp"`
	NetworkDevices LXCNetworkDeviceDict `n:"net"`
	StartOnBoot    bool                 `n:"onboot"`
	TTYCount       int                  `n:"tty"`
	HasConsole     bool                 `n:"console"`
	ConsoleMode    string               `n:"cmode"`
	IsProtected    bool                 `n:"protection"`
	IsTemplate     bool                 `n:"template"`
	IsUnprivileged bool                 `n:"unprivileged"`
}

type LXCMountPoint struct {
	Volume         string `n:",volume"`
	MountPoint     string `n:"mp" d:"/"`
	Size           string `n:"size"`
	HasACL         bool   `n:"acl" d:"false"`
	HasBackup      bool   `n:"backup" d:"false"`
	HasQuota       bool   `n:"quota" d:"false"`
	HasReplication bool   `n:"replicate" d:"true"`
	IsReadOnly     bool   `n:"ro" d:"false"`
	IsShared       bool   `n:"shared" d:"false"`
}

type LXCMountPointDict = map[int]*LXCMountPoint

type LXCNetworkDevice struct {
	Type        string `n:"type"`
	Name        string `n:",name"`
	MACAddress  string `n:"hwaddr"`
	Bridge      string `n:"bridge"`
	IPAddress   string `n:"ip"`
	Gateway     string `n:"gw"`
	IPv6Address string `n:"ip6"`
	GatewayIPv6 string `n:"gw6"`
	MTU         int    `n:"mtu"`
	Rate        int    `n:"rate"`
	VLANTag     int    `n:"tag"`
	Trunks      []int  `n:"trunks"`
	HasFirewall bool   `n:"firewall"`
}

type LXCNetworkDeviceDict = map[int]*LXCNetworkDevice

const (
	LXCDefaultCPULimit      = 0.0
	LXCDefaultCPUUnits      = 1024
	LXCMinimumMountPoint    = 0
	LXCMaximumMountPoint    = 9
	LXCMinimumNetworkDevice = 0
	LXCMaximumNetworkDevice = 9
)

// LXC Architectures
const (
	AMD64 = "amd64"
	I386  = "i386"
)

// LXC OS Types
const (
	Unmanaged = "unmanaged"
	Alpine    = "alpine"
	Debian    = "debian"
	Ubuntu    = "ubuntu"
	CentOS    = "centos"
	Fedora    = "fedora"
	OpenSUSE  = "opensuse"
	ArchLinux = "archlinux"
	Gentoo    = "gentoo"
)

func (e *LXC) Start() error {
	return e.provider.Start(e.VMID)
}

func (e *LXC) Stop() error {
	return e.provider.Stop(e.VMID)
}

func (e *LXC) Reset() error {
	return e.provider.Reset(e.VMID)
}

func (e *LXC) Shutdown() error {
	return e.provider.Shutdown(e.VMID)
}

func (e *LXC) Suspend() error {
	return e.provider.Suspend(e.VMID)
}

func (e *LXC) Resume() error {
	return e.provider.Resume(e.VMID)
}

func (e *LXC) Clone(opts *VMCreateOptions) (*Task, error) {
	return e.provider.Clone(e.VMID, opts)
}

func (e *LXC) Update(cfg *LXCConfig) error {
	return e.provider.Update(e.VMID, cfg)
}

func (e *LXC) Delete() (*Task, error) {
	return e.provider.Delete(e.VMID)
}
