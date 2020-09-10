package firewall

import (
	"encoding/json"
	"fmt"
)

type Protocol uint

const (
	ProtocolNone Protocol = iota
	// Internet Control Message Protocol
	ProtocolICMP
	// Internet Group Management
	ProtocolIGMP
	// gateway-gateway protocol
	ProtocolGGP
	// IP encapsulated in IP
	ProtocolIPEncap
	// ST datagram mode
	ProtocolST
	// Transmission Control Protocol
	ProtocolTCP
	// exterior gateway protocol
	ProtocolEGP
	// any private interior gateway (Cisco)
	ProtocolIGP
	// PARC universal packet protocol
	ProtocolPUP
	// User Datagram Protocol
	ProtocolUDP
	// host monitoring protocol
	ProtocolHMP
	// Xerox NS IDP
	ProtocolXNSIDP
	// reliable datagram" protocol
	ProtocolRDP
	// ISO Transport Protocol class 4 [RFC905]
	ProtocolISOTP4
	// Datagram Congestion Control Prot. [RFC4340]
	ProtocolDCCP
	// Xpress Transfer Protocol
	ProtocolXTP
	// Datagram Delivery Protocol
	ProtocolDDP
	// IDPR Control Message Transport
	ProtocolIDPRCMTP
	// Internet Protocol, version 6
	ProtocolIPv6
	// Routing Header for IPv6
	ProtocolIPv6Route
	// Fragment Header for IPv6
	ProtocolIPv6Frag
	// Inter-Domain Routing Protocol
	ProtocolIDRP
	// Reservation Protocol
	ProtocolRSVP
	// General Routing Encapsulation
	ProtocolGRE
	// Encap Security Payload [RFC2406]
	ProtocolESP
	// Authentication Header [RFC2402]
	ProtocolAH
	// SKIP
	ProtocolSKIP
	// ICMP for IPv6
	ProtocolIPv6ICMP
	// No Next Header for IPv6
	ProtocolIPv6NoNXT
	// Destination Options for IPv6
	ProtocolIPv6Opts
	// Versatile Message Transport
	ProtocolVMTP
	// Enhanced Interior Routing Protocol (Cisco)
	ProtocolEIGRP
	// Open Shortest Path First IGP
	ProtocolOSPF
	// AX.25 frames
	ProtocolAX25
	// IP-within-IP Encapsulation Protocol
	ProtocolIPIP
	// Ethernet-within-IP Encapsulation [RFC3378]
	ProtocolEtherIP
	// Yet Another IP encapsulation [RFC1241]
	ProtocolEncap
	// Protocol Independent Multicast
	ProtocolPIM
	// IP Payload Compression Protocol
	ProtocolIPComp
	// Virtual Router Redundancy Protocol [RFC5798]
	ProtocolVRRP
	// Layer Two Tunneling Protocol [RFC2661]
	ProtocolL2TP
	// IS-IS over IPv4
	ProtocolISIS
	// Stream Control Transmission Protocol
	ProtocolSCTP
	// Fibre Channel
	ProtocolFC
	// Mobility Support for IPv6 [RFC3775]
	ProtocolMobilityHeader
	// UDP-Lite [RFC3828]
	ProtocolUDPLite
	// MPLS-in-IP [RFC4023]
	ProtocolMPLSInIP
	// Host Identity Protocol
	ProtocolHIP
	// Shim6 Protocol [RFC5533]
	ProtocolSHIM6
	// Wrapped Encapsulating Security Payload
	ProtocolWESP
	// Robust Header Compression
	ProtocolROHC
)

func (obj Protocol) Marshal() (string, error) {
	switch obj {
	case ProtocolNone:
		return "", nil
	case ProtocolICMP:
		return "icmp", nil
	case ProtocolIGMP:
		return "igmp", nil
	case ProtocolGGP:
		return "ggp", nil
	case ProtocolIPEncap:
		return "ipencap", nil
	case ProtocolST:
		return "st", nil
	case ProtocolTCP:
		return "tcp", nil
	case ProtocolEGP:
		return "egp", nil
	case ProtocolIGP:
		return "igp", nil
	case ProtocolPUP:
		return "pup", nil
	case ProtocolUDP:
		return "udp", nil
	case ProtocolHMP:
		return "hmp", nil
	case ProtocolXNSIDP:
		return "xns-idp", nil
	case ProtocolRDP:
		return "rdp", nil
	case ProtocolISOTP4:
		return "iso-tp4", nil
	case ProtocolDCCP:
		return "dccp", nil
	case ProtocolXTP:
		return "xtp", nil
	case ProtocolDDP:
		return "ddp", nil
	case ProtocolIDPRCMTP:
		return "idpr-cmtp", nil
	case ProtocolIPv6:
		return "ipv6", nil
	case ProtocolIPv6Route:
		return "ipv6-route", nil
	case ProtocolIPv6Frag:
		return "ipv6-frag", nil
	case ProtocolIDRP:
		return "idrp", nil
	case ProtocolRSVP:
		return "rsvp", nil
	case ProtocolGRE:
		return "gre", nil
	case ProtocolESP:
		return "esp", nil
	case ProtocolAH:
		return "ah", nil
	case ProtocolSKIP:
		return "skip", nil
	case ProtocolIPv6ICMP:
		return "ipv6-icmp", nil
	case ProtocolIPv6NoNXT:
		return "ipv6-nonxt", nil
	case ProtocolIPv6Opts:
		return "ipv6-opts", nil
	case ProtocolVMTP:
		return "vmtp", nil
	case ProtocolEIGRP:
		return "eigrp", nil
	case ProtocolOSPF:
		return "ospf", nil
	case ProtocolAX25:
		return "ax.25", nil
	case ProtocolIPIP:
		return "ipip", nil
	case ProtocolEtherIP:
		return "etherip", nil
	case ProtocolEncap:
		return "encap", nil
	case ProtocolPIM:
		return "pim", nil
	case ProtocolIPComp:
		return "ipcomp", nil
	case ProtocolVRRP:
		return "vrrp", nil
	case ProtocolL2TP:
		return "l2tp", nil
	case ProtocolISIS:
		return "isis", nil
	case ProtocolSCTP:
		return "sctp", nil
	case ProtocolFC:
		return "fc", nil
	case ProtocolMobilityHeader:
		return "mobility-header", nil
	case ProtocolUDPLite:
		return "udplite", nil
	case ProtocolMPLSInIP:
		return "mpls-in-ip", nil
	case ProtocolHIP:
		return "hip", nil
	case ProtocolSHIM6:
		return "shim6", nil
	case ProtocolWESP:
		return "wesp", nil
	case ProtocolROHC:
		return "rohc", nil
	default:
		return "", fmt.Errorf("unknown firewall protocol")
	}
}

func (obj *Protocol) Unmarshal(s string) error {
	switch s {
	case "":
		*obj = ProtocolNone
	case "icmp":
	case "1":
		*obj = ProtocolICMP
	case "igmp":
	case "2":
		*obj = ProtocolIGMP
	case "ggp":
	case "3":
		*obj = ProtocolGGP
	case "ipencap":
	case "4":
		*obj = ProtocolIPEncap
	case "st":
	case "5":
		*obj = ProtocolST
	case "tcp":
	case "6":
		*obj = ProtocolTCP
	case "egp":
	case "8":
		*obj = ProtocolEGP
	case "igp":
	case "9":
		*obj = ProtocolIGP
	case "pup":
	case "12":
		*obj = ProtocolPUP
	case "udp":
	case "17":
		*obj = ProtocolUDP
	case "hmp":
	case "20":
		*obj = ProtocolHMP
	case "xns-idp":
	case "22":
		*obj = ProtocolXNSIDP
	case "rdp":
	case "27":
		*obj = ProtocolRDP
	case "iso-tp4":
	case "29":
		*obj = ProtocolISOTP4
	case "dccp":
	case "33":
		*obj = ProtocolDCCP
	case "xtp":
	case "36":
		*obj = ProtocolXTP
	case "ddp":
	case "37":
		*obj = ProtocolDDP
	case "idpr-cmtp":
	case "38":
		*obj = ProtocolIDPRCMTP
	case "ipv6":
	case "41":
		*obj = ProtocolIPv6
	case "ipv6-route":
	case "43":
		*obj = ProtocolIPv6Route
	case "ipv6-frag":
	case "44":
		*obj = ProtocolIPv6Frag
	case "idrp":
	case "45":
		*obj = ProtocolIDRP
	case "rsvp":
	case "46":
		*obj = ProtocolRSVP
	case "gre":
	case "47":
		*obj = ProtocolGRE
	case "esp":
	case "50":
		*obj = ProtocolESP
	case "ah":
	case "51":
		*obj = ProtocolAH
	case "skip":
	case "57":
		*obj = ProtocolSKIP
	case "ipv6-icmp":
	case "58":
		*obj = ProtocolIPv6ICMP
	case "ipv6-nonxt":
	case "59":
		*obj = ProtocolIPv6NoNXT
	case "ipv6-opts":
	case "60":
		*obj = ProtocolIPv6Opts
	case "vmtp":
	case "81":
		*obj = ProtocolVMTP
	case "eigrp":
	case "88":
		*obj = ProtocolEIGRP
	case "ospf":
	case "89":
		*obj = ProtocolOSPF
	case "ax.25":
	case "93":
		*obj = ProtocolAX25
	case "ipip":
	case "94":
		*obj = ProtocolIPIP
	case "etherip":
	case "97":
		*obj = ProtocolEtherIP
	case "encap":
	case "98":
		*obj = ProtocolEncap
	case "pim":
	case "103":
		*obj = ProtocolPIM
	case "ipcomp":
	case "108":
		*obj = ProtocolIPComp
	case "vrrp":
	case "112":
		*obj = ProtocolVRRP
	case "l2tp":
	case "115":
		*obj = ProtocolL2TP
	case "isis":
	case "124":
		*obj = ProtocolISIS
	case "sctp":
	case "132":
		*obj = ProtocolSCTP
	case "fc":
	case "133":
		*obj = ProtocolFC
	case "mobility-header":
	case "135":
		*obj = ProtocolMobilityHeader
	case "udplite":
	case "136":
		*obj = ProtocolUDPLite
	case "mpls-in-ip":
	case "137":
		*obj = ProtocolMPLSInIP
	case "hip":
	case "139":
		*obj = ProtocolHIP
	case "shim6":
	case "140":
		*obj = ProtocolSHIM6
	case "wesp":
	case "141":
		*obj = ProtocolWESP
	case "rohc":
	case "142":
		*obj = ProtocolROHC
	default:
		return fmt.Errorf("can't unmarshal firewall protocol %s", s)
	}

	return nil
}

func (obj *Protocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
