package service

type LXC struct {
	provider LXCServiceProvider

	VMID        int
	Name        string
	Description string
	Status      string
	LXCConfig
}

type LXCConfig struct {
	Architecture   string             `n:"arch"`
	OSType         string             `n:"ostype"`
	CPU            int                `n:"cores"`
	CPULimit       int                `n:"cpulimit"`
	CPUUnits       int                `n:"cpuunits"`
	MemoryTotal    int                `n:"memory"`
	MemorySwap     int                `n:"swap"`
	RootMountPoint LXCMountPoint      `n:"rootfs"`
	MountPoints    []LXCMountPoint    `n:"mp"`
	NetworkDevices []LXCNetworkDevice `n:"net"`
	StartOnBoot    bool               `n:"onboot"`
	TTYCount       int                `n:"tty"`
	HasConsole     bool               `n:"console"`
	ConsoleMode    string             `n:"cmode"`
	IsProtected    bool               `n:"protection"`
	IsTemplate     bool               `n:"template"`
	IsUnprivileged bool               `n:"unprivileged"`
}

type LXCList []*LXC

type LXCMountPoint struct {
	Volume         string `n:"volume"`
	MountPoint     string `n:"mp"`
	Size           string `n:"size"`
	HasACL         bool   `n:"acl"`
	HasBackup      bool   `n:"backup"`
	HasQuota       bool   `n:"quota"`
	HasReplication bool   `n:"replicate"`
	IsReadOnly     bool   `n:"ro"`
	IsShared       bool   `n:"shared"`
}

type LXCNetworkDevice struct {
	Name        string `n:"name"`
	Bridge      string `n:"bridge"`
	Firewall    bool   `n:"firewall"`
	Gateway     string `n:"gw"`
	GatewayIPv6 string `n:"gw6"`
	MACAddress  string `n:"hwaddr"`
	IPAddress   string `n:"ip"`
	IPv6Address string `n:"ip6"`
	MTU         int    `n:"mtu"`
	Rate        string `n:"rate"`
	VLANTag     int    `n:"tag"`
	Trunks      string `n:"trunks"`
	Type        string `n:"type"`
}

const (
	LXCDefaultCPULimit = 0
	LXCDefaultCPUUnits = 1000
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
