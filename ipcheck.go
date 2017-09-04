// ipcheck <ip> -> category
// Given an IP, describe it as either public, private, loopback, multicast, etc.
package main

import (
	"fmt"
	"net"
	"os"
)

var _, linkLocal, _ = net.ParseCIDR("169.254.0.0/16")
var _, private10, _ = net.ParseCIDR("10.0.0.0/8")
var _, private172, _ = net.ParseCIDR("172.16.0.0/12")
var _, private192, _ = net.ParseCIDR("192.168.0.0/16")
var _, private100, _ = net.ParseCIDR("100.64.0.0/10") // https://tools.ietf.org/html/rfc6598
var _, private198, _ = net.ParseCIDR("198.18.0.0/15") // https://tools.ietf.org/html/rfc2544
var _, loopBack0, _ = net.ParseCIDR("0.0.0.0/8")
var _, sixToFour, _ = net.ParseCIDR("192.88.99.0/24")
var _, doc192, _ = net.ParseCIDR("192.0.2.0/24") // https://tools.ietf.org/html/rfc5737
var _, doc198, _ = net.ParseCIDR("198.51.100.0/24")
var _, doc203, _ = net.ParseCIDR("203.0.113.0/24")
var _, reserved, _ = net.ParseCIDR("240.0.0.0/4") // https://tools.ietf.org/html/rfc6890

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

// IpCheck takes an IPv4 as as string and returns one of the following types.
// malformed, public, private, loopback, multicast
// linklocal, six2four, documentation, reserved
func IpCheck(s string) string {
	ip := net.ParseIP(s)
	switch {
	case ip == nil:
		return "malformed"
	case ip.IsLoopback():
		return "loopback"
	case loopBack0.Contains(ip):
		return "loopback"
	case ip.IsMulticast():
		return "multicast"
	case linkLocal.Contains(ip):
		return "linklocal"
	case isPrivate(ip):
		return "private"
	case sixToFour.Contains(ip):
		return "six2four"
	case isDocumentation(ip):
		return "documentation"
	case reserved.Contains(ip):
		return "reserved"
	default:
		return "public"
	}
}

func main() {
	fmt.Println(IpCheck(os.Args[1]))
}
