package firewall

import (
	"encoding/json"
	"fmt"
)

type Macro uint

const (
	MacroNone Macro = iota
	// Proxmox Mail Gateway web interface
	MacroPMG
	// Telnet over SSL
	MacroTelnets
	// OpenPGP HTTP keyserver protocol traffic
	MacroHKP
	// Internet Relay Chat traffic
	MacroIRC
	// RFC 868 Time protocol
	MacroTime
	// Digital Audio Access Protocol traffic (iTunes, Rythmbox daemons)
	MacroDAAP
	// BitTorrent traffic for BitTorrent 3.1 and earlier
	MacroBitTorrent
	// Amanda Backup
	MacroAmanda
	// Traceroute (for up to 30 hops) traffic
	MacroTrcrt
	// Subversion server (svnserve)
	MacroSVN
	// Concurrent Versions System pserver traffic
	MacroCVS
	// Layer 2 Tunneling Protocol traffic
	MacroL2TP
	// Hypertext Transfer Protocol (WWW)
	MacroHTTP
	// Microsoft SQL Server
	MacroMSSQL
	// IPsec traffic and Nat-Traversal
	MacroIPsecnat
	// HP Jetdirect printing
	MacroJetdirect
	// Mail message submission traffic
	MacroSubmission
	// Squid web proxy traffic
	MacroSquid
	// IPIP capsulation traffic
	MacroIPIP
	// Microsoft Notification Protocol
	MacroMSNP
	// Simple Mail Transfer Protocol
	MacroSMTP
	// Auth (identd) traffic
	MacroAuth
	// Secure Lightweight Directory Access Protocol traffic
	MacroLDAPS
	// SixXS IPv6 Deployment and Tunnel Broker
	MacroSixXS
	// Distributed Compiler service
	MacroDistcc
	// Internet Message Access Protocol over SSL
	MacroIMAPS
	// Point-to-Point Tunneling Protocol
	MacroPPtP
	// POP3 traffic
	MacroPOP3
	// SANE network scanning
	MacroSANE
	// OpenVPN traffic
	MacroOpenVPN
	// Forwarded DHCP traffic
	MacroDHCPfwd
	// MySQL server
	MacroMySQL
	// BIND remote management protocol
	MacroRNDC
	// NNTP traffic (Usenet).
	MacroNNTP
	// BitTorrent traffic for BitTorrent 3.2 and later
	MacroBitTorrent32
	// IPsec authentication (AH) traffic
	MacroIPsecah
	// ICMP echo request
	MacroPing
	// Distributed Checksum Clearinghouse spam filtering mechanism
	MacroDCC
	// GNUnet secure peer-to-peer networking traffic
	MacroGNUnet
	// File Transfer Protocol
	MacroFTP
	// Border Gateway Protocol traffic
	MacroBGP
	// Simple Network Management Protocol
	MacroSNMP
	// Microsoft Remote Desktop Protocol traffic
	MacroRDP
	// Finger protocol (RFC 742)
	MacroFinger
	// Symantec PCAnywere (tm)
	MacroPCA
	// Internet Cache Protocol V2 (Squid) traffic
	MacroICPV2
	// Telnet traffic
	MacroTelnet
	// Ceph Storage Cluster traffic (Ceph Monitors, OSD & MDS Deamons)
	MacroCeph
	// IPsec traffic
	MacroIPsec
	// Microsoft SMB traffic
	MacroSMB
	// Web Cache/Proxy traffic (port 8080)
	MacroWebcache
	// VNC traffic from Vncservers to Vncviewers in listen mode
	MacroVNCL
	// Network Time Protocol (ntpd)
	MacroNTP
	// Domain Name System traffic (upd and tcp)
	MacroDNS
	// Mail traffic (SMTP, SMTPS, Submission)
	MacroMail
	// DHCPv6 traffic
	MacroDHCPv6
	// Citrix/ICA traffic (ICA, ICA Browser, CGP)
	MacroCitrix
	// Line Printer protocol printing
	MacroPrinter
	// WWW traffic (HTTP and HTTPS)
	MacroWeb
	// Hypertext Transfer Protocol (WWW) over SSL
	MacroHTTPS
	// Webmin traffic
	MacroWebmin
	// PostgreSQL server
	MacroPostgreSQL
	// AOL Instant Messenger traffic
	MacroICQ
	// Trivial File Transfer Protocol traffic
	MacroTFTP
	// Encrypted POP3 traffic
	MacroPOP3S
	// Internet Message Access Protocol
	MacroIMAP
	// Generic Routing Encapsulation tunneling protocol
	MacroGRE
	// Lightweight Directory Access Protocol traffic
	MacroLDAP
	// IPv6 neighbor solicitation, neighbor and router advertisement
	MacroNeighborDiscovery
	// Razor Antispam System
	MacroRazor
	// OSPF multicast traffic
	MacroOSPF
	// Secure shell traffic
	MacroSSH
	// Git distributed revision control traffic
	MacroGit
	// Rsync server
	MacroRsync
	// Encrypted Simple Mail Transfer Protocol
	MacroSMTPS
	// Munin networked resource monitoring traffic
	MacroMunin
	// Remote time retrieval (rdate)
	MacroRdate
	// Routing Information Protocol (bidirectional)
	MacroRIP
	// Syslog protocol (RFC 5424) traffic
	MacroSyslog
	// Samba Web Administration Tool
	MacroSMBswat
	// Whois (nicname, RFC 3912) traffic
	MacroWhois
	// Spam Assassin SPAMD traffic
	MacroSPAMD
	// VNC traffic for VNC display's 0 - 99
	MacroVNC
	// Multicast DNS
	MacroMDNS
	// Encrypted NNTP traffic (Usenet)
	MacroNNTPS
)

func (obj Macro) Marshal() (string, error) {
	switch obj {
	case MacroNone:
		return "", nil
	case MacroPMG:
		return "PMG", nil
	case MacroTelnets:
		return "Telnets", nil
	case MacroHKP:
		return "HKP", nil
	case MacroIRC:
		return "IRC", nil
	case MacroTime:
		return "Time", nil
	case MacroDAAP:
		return "DAAP", nil
	case MacroBitTorrent:
		return "BitTorrent", nil
	case MacroAmanda:
		return "Amanda", nil
	case MacroTrcrt:
		return "Trcrt", nil
	case MacroSVN:
		return "SVN", nil
	case MacroCVS:
		return "CVS", nil
	case MacroL2TP:
		return "L2TP", nil
	case MacroHTTP:
		return "HTTP", nil
	case MacroMSSQL:
		return "MSSQL", nil
	case MacroIPsecnat:
		return "IPsecnat", nil
	case MacroJetdirect:
		return "Jetdirect", nil
	case MacroSubmission:
		return "Submission", nil
	case MacroSquid:
		return "Squid", nil
	case MacroIPIP:
		return "IPIP", nil
	case MacroMSNP:
		return "MSNP", nil
	case MacroSMTP:
		return "SMTP", nil
	case MacroAuth:
		return "Auth", nil
	case MacroLDAPS:
		return "LDAPS", nil
	case MacroSixXS:
		return "SixXS", nil
	case MacroDistcc:
		return "Distcc", nil
	case MacroIMAPS:
		return "IMAPS", nil
	case MacroPPtP:
		return "PPtP", nil
	case MacroPOP3:
		return "POP3", nil
	case MacroSANE:
		return "SANE", nil
	case MacroOpenVPN:
		return "OpenVPN", nil
	case MacroDHCPfwd:
		return "DHCPfwd", nil
	case MacroMySQL:
		return "MySQL", nil
	case MacroRNDC:
		return "RNDC", nil
	case MacroNNTP:
		return "NNTP", nil
	case MacroBitTorrent32:
		return "BitTorrent32", nil
	case MacroIPsecah:
		return "IPsecah", nil
	case MacroPing:
		return "Ping", nil
	case MacroDCC:
		return "DCC", nil
	case MacroGNUnet:
		return "GNUnet", nil
	case MacroFTP:
		return "FTP", nil
	case MacroBGP:
		return "BGP", nil
	case MacroSNMP:
		return "SNMP", nil
	case MacroRDP:
		return "RDP", nil
	case MacroFinger:
		return "Finger", nil
	case MacroPCA:
		return "PCA", nil
	case MacroICPV2:
		return "ICPV2", nil
	case MacroTelnet:
		return "Telnet", nil
	case MacroCeph:
		return "Ceph", nil
	case MacroIPsec:
		return "IPsec", nil
	case MacroSMB:
		return "SMB", nil
	case MacroWebcache:
		return "Webcache", nil
	case MacroVNCL:
		return "VNCL", nil
	case MacroNTP:
		return "NTP", nil
	case MacroDNS:
		return "DNS", nil
	case MacroMail:
		return "Mail", nil
	case MacroDHCPv6:
		return "DHCPv6", nil
	case MacroCitrix:
		return "Citrix", nil
	case MacroPrinter:
		return "Printer", nil
	case MacroWeb:
		return "Web", nil
	case MacroHTTPS:
		return "HTTPS", nil
	case MacroWebmin:
		return "Webmin", nil
	case MacroPostgreSQL:
		return "PostgreSQL", nil
	case MacroICQ:
		return "ICQ", nil
	case MacroTFTP:
		return "TFTP", nil
	case MacroPOP3S:
		return "POP3S", nil
	case MacroIMAP:
		return "IMAP", nil
	case MacroGRE:
		return "GRE", nil
	case MacroLDAP:
		return "LDAP", nil
	case MacroNeighborDiscovery:
		return "NeighborDiscovery", nil
	case MacroRazor:
		return "Razor", nil
	case MacroOSPF:
		return "OSPF", nil
	case MacroSSH:
		return "SSH", nil
	case MacroGit:
		return "Git", nil
	case MacroRsync:
		return "Rsync", nil
	case MacroSMTPS:
		return "SMTPS", nil
	case MacroMunin:
		return "Munin", nil
	case MacroRdate:
		return "Rdate", nil
	case MacroRIP:
		return "RIP", nil
	case MacroSyslog:
		return "Syslog", nil
	case MacroSMBswat:
		return "SMBswat", nil
	case MacroWhois:
		return "Whois", nil
	case MacroSPAMD:
		return "SPAMD", nil
	case MacroVNC:
		return "VNC", nil
	case MacroMDNS:
		return "MDNS", nil
	case MacroNNTPS:
		return "NNTPS", nil

	default:
		return "", fmt.Errorf("unknown firewall macro")
	}
}

func (obj *Macro) Unmarshal(s string) error {
	switch s {
	case "":
		*obj = MacroNone
	case "PMG":
		*obj = MacroPMG
	case "Telnets":
		*obj = MacroTelnets
	case "HKP":
		*obj = MacroHKP
	case "IRC":
		*obj = MacroIRC
	case "Time":
		*obj = MacroTime
	case "DAAP":
		*obj = MacroDAAP
	case "BitTorrent":
		*obj = MacroBitTorrent
	case "Amanda":
		*obj = MacroAmanda
	case "Trcrt":
		*obj = MacroTrcrt
	case "SVN":
		*obj = MacroSVN
	case "CVS":
		*obj = MacroCVS
	case "L2TP":
		*obj = MacroL2TP
	case "HTTP":
		*obj = MacroHTTP
	case "MSSQL":
		*obj = MacroMSSQL
	case "IPsecnat":
		*obj = MacroIPsecnat
	case "Jetdirect":
		*obj = MacroJetdirect
	case "Submission":
		*obj = MacroSubmission
	case "Squid":
		*obj = MacroSquid
	case "IPIP":
		*obj = MacroIPIP
	case "MSNP":
		*obj = MacroMSNP
	case "SMTP":
		*obj = MacroSMTP
	case "Auth":
		*obj = MacroAuth
	case "LDAPS":
		*obj = MacroLDAPS
	case "SixXS":
		*obj = MacroSixXS
	case "Distcc":
		*obj = MacroDistcc
	case "IMAPS":
		*obj = MacroIMAPS
	case "PPtP":
		*obj = MacroPPtP
	case "POP3":
		*obj = MacroPOP3
	case "SANE":
		*obj = MacroSANE
	case "OpenVPN":
		*obj = MacroOpenVPN
	case "DHCPfwd":
		*obj = MacroDHCPfwd
	case "MySQL":
		*obj = MacroMySQL
	case "RNDC":
		*obj = MacroRNDC
	case "NNTP":
		*obj = MacroNNTP
	case "BitTorrent32":
		*obj = MacroBitTorrent32
	case "IPsecah":
		*obj = MacroIPsecah
	case "Ping":
		*obj = MacroPing
	case "DCC":
		*obj = MacroDCC
	case "GNUnet":
		*obj = MacroGNUnet
	case "FTP":
		*obj = MacroFTP
	case "BGP":
		*obj = MacroBGP
	case "SNMP":
		*obj = MacroSNMP
	case "RDP":
		*obj = MacroRDP
	case "Finger":
		*obj = MacroFinger
	case "PCA":
		*obj = MacroPCA
	case "ICPV2":
		*obj = MacroICPV2
	case "Telnet":
		*obj = MacroTelnet
	case "Ceph":
		*obj = MacroCeph
	case "IPsec":
		*obj = MacroIPsec
	case "SMB":
		*obj = MacroSMB
	case "Webcache":
		*obj = MacroWebcache
	case "VNCL":
		*obj = MacroVNCL
	case "NTP":
		*obj = MacroNTP
	case "DNS":
		*obj = MacroDNS
	case "Mail":
		*obj = MacroMail
	case "DHCPv6":
		*obj = MacroDHCPv6
	case "Citrix":
		*obj = MacroCitrix
	case "Printer":
		*obj = MacroPrinter
	case "Web":
		*obj = MacroWeb
	case "HTTPS":
		*obj = MacroHTTPS
	case "Webmin":
		*obj = MacroWebmin
	case "PostgreSQL":
		*obj = MacroPostgreSQL
	case "ICQ":
		*obj = MacroICQ
	case "TFTP":
		*obj = MacroTFTP
	case "POP3S":
		*obj = MacroPOP3S
	case "IMAP":
		*obj = MacroIMAP
	case "GRE":
		*obj = MacroGRE
	case "LDAP":
		*obj = MacroLDAP
	case "NeighborDiscovery":
		*obj = MacroNeighborDiscovery
	case "Razor":
		*obj = MacroRazor
	case "OSPF":
		*obj = MacroOSPF
	case "SSH":
		*obj = MacroSSH
	case "Git":
		*obj = MacroGit
	case "Rsync":
		*obj = MacroRsync
	case "SMTPS":
		*obj = MacroSMTPS
	case "Munin":
		*obj = MacroMunin
	case "Rdate":
		*obj = MacroRdate
	case "RIP":
		*obj = MacroRIP
	case "Syslog":
		*obj = MacroSyslog
	case "SMBswat":
		*obj = MacroSMBswat
	case "Whois":
		*obj = MacroWhois
	case "SPAMD":
		*obj = MacroSPAMD
	case "VNC":
		*obj = MacroVNC
	case "MDNS":
		*obj = MacroMDNS
	case "NNTPS":
		*obj = MacroNNTPS
	default:
		return fmt.Errorf("can't unmarshal firewall macro %s", s)
	}

	return nil
}

func (obj *Macro) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
