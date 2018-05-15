package ipcheck

import (
	"errors"
	"net"
)

type ipTypeId uint8

const (
	public ipTypeId = iota
	private
	loopback
	multicast
	linklocal
	six2four
	documentation
	reserved
)

type IpType struct {
	Id   ipTypeId
	Name string
}

func (ip IpType) String() string {
	return ip.Name
}

// These should really be consts, but...
var Public = IpType{public, "public"}
var Private = IpType{private, "private"}
var Loopback = IpType{loopback, "loopback"}
var Multicast = IpType{multicast, "multicast"}
var LinkLocal = IpType{linklocal, "linklocal"}
var Six2Four = IpType{six2four, "6to4"}
var Documentation = IpType{documentation, "documetation"}
var Reserved = IpType{reserved, "reserved"}

type IpError error

var Malformed = IpError(errors.New("malformed"))

var _, linkLocal, _ = net.ParseCIDR("169.254.0.0/16")
var _, private10, _ = net.ParseCIDR("10.0.0.0/8")
var _, private172, _ = net.ParseCIDR("172.16.0.0/12")
var _, private192, _ = net.ParseCIDR("192.168.0.0/16")
var _, private100, _ = net.ParseCIDR("100.64.0.0/10")      // https://tools.ietf.org/html/rfc6598
var _, private192Inter, _ = net.ParseCIDR("198.18.0.0/15") // https://tools.ietf.org/html/rfc2544
var _, private198, _ = net.ParseCIDR("198.18.0.0/15")      // https://tools.ietf.org/html/rfc2544
var _, loopBack0, _ = net.ParseCIDR("0.0.0.0/8")
var _, sixToFour, _ = net.ParseCIDR("192.88.99.0/24")
var _, doc192, _ = net.ParseCIDR("192.0.2.0/24") // https://tools.ietf.org/html/rfc5737
var _, doc198, _ = net.ParseCIDR("198.51.100.0/24")
var _, doc203, _ = net.ParseCIDR("203.0.113.0/24")
var _, reservedNet, _ = net.ParseCIDR("240.0.0.0/4") // https://tools.ietf.org/html/rfc6890

// isPrivate, a missing method for net.IP
func isPrivate(ip net.IP) bool {
	switch {
	case private10.Contains(ip):
		return true
	case private100.Contains(ip):
		return true
	case private172.Contains(ip):
		return true
	case private192.Contains(ip):
		return true
	case private198.Contains(ip):
		return true
	default:
		return false
	}
}

// isDocumentation, a missing method for net.IP
func isDocumentation(ip net.IP) bool {
	switch {
	case doc192.Contains(ip):
		return true
	case doc198.Contains(ip):
		return true
	case doc203.Contains(ip):
		return true
	default:
		return false
	}
}

func Type(s string) (r IpType, err error) {
	err = nil
	ip := net.ParseIP(s)
	switch {
	case ip == nil:
		err = errors.New("malformed")
	case ip.IsLoopback():
		r = Loopback
	case loopBack0.Contains(ip):
		r = Loopback
	case ip.IsMulticast():
		r = Multicast
	case linkLocal.Contains(ip):
		r = LinkLocal
	case isPrivate(ip):
		r = Private
	case sixToFour.Contains(ip):
		r = Six2Four
	case isDocumentation(ip):
		r = Documentation
	case reservedNet.Contains(ip):
		r = Reserved
	default:
		r = Public
	}
	return
}
