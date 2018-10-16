package service

type LXC struct {
	provider LXCServiceProvider

	VMID         int    `n:"vmid"`
	Name         string `n:"name"`
	Description  string `n:"description"`
	Status       string `n:"status"`
	ConfigDigest string `n:"digest"`
	LXCConfig
}

type LXCConfig struct {
	Architecture   string               `n:"arch" d:"amd64"`
	OSType         string               `n:"ostype"`
	CPU            int                  `n:"cores"`
	CPULimit       float64              `n:"cpulimit" d:"0.0"`
	CPUUnits       int                  `n:"cpuunits" d:"1024"`
	MemoryTotal    int                  `n:"memory"`
	MemorySwap     int                  `n:"swap"`
	RootMountPoint LXCMountPoint        `n:"rootfs" t:"kv"`
	MountPoints    LXCMountPointDict    `n:"mp" t:"kvdict" min:"0" max:"9"`
	NetworkDevices LXCNetworkDeviceDict `n:"net" t:"kvdict" min:"0" max:"9"`
	DNSServers     string               `n:"nameserver"`
	DNSDomain      string               `n:"searchdomain"`
	StartOnBoot    bool                 `n:"onboot" d:"0"`
	LockType       string               `n:"lock"`
	TTYCount       int                  `n:"tty" d:"2"`
	HasConsole     bool                 `n:"console" d:"1"`
	ConsoleMode    string               `n:"cmode" d:"tty"`
	IsProtected    bool                 `n:"protection" d:"0"`
	IsTemplate     bool                 `n:"template" d:"0"`
	IsUnprivileged bool                 `n:"unprivileged" d:"0"`
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
	Trunks      []int  `n:"trunks" d:"" s:";"`
	HasFirewall bool   `n:"firewall"`
}

type LXCNetworkDeviceDict = map[int]*LXCNetworkDevice

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
