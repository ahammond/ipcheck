package main

import (
	"fmt"
	"net"
	"os"
)

type IpType uint8

const (
	malformed IpType = iota
	public
	private
	loopback
	multicast
	linklocal
	six2four
	documentation
	reserved
)

var name = map[IpType]string{
	malformed:     "malformed",
	public:        "public",
	private:       "private",
	loopback:      "loopback",
	multicast:     "multicast",
	linklocal:     "linklocal",
	six2four:      "6to4",
	documentation: "documentation",
	reserved:      "reserved",
}

var _, LinkLocal, _ = net.ParseCIDR("169.254.0.0/16")
var _, Private10, _ = net.ParseCIDR("10.0.0.0/8")
var _, Private172, _ = net.ParseCIDR("172.16.0.0/12")
var _, Private192, _ = net.ParseCIDR("192.168.0.0/16")
var _, Private100, _ = net.ParseCIDR("100.64.0.0/10")      // https://tools.ietf.org/html/rfc6598
var _, Private192Inter, _ = net.ParseCIDR("198.18.0.0/15") // https://tools.ietf.org/html/rfc2544
var _, LoopBack0, _ = net.ParseCIDR("0.0.0.0/8")
var _, SixToFour, _ = net.ParseCIDR("192.88.99.0/24")
var _, Doc192, _ = net.ParseCIDR("192.0.2.0/24") // https://tools.ietf.org/html/rfc5737
var _, Doc198, _ = net.ParseCIDR("198.51.100.0/24")
var _, Doc203, _ = net.ParseCIDR("203.0.113.0/24")
var _, Reserved, _ = net.ParseCIDR("240.0.0.0/4") // https://tools.ietf.org/html/rfc6890

func Categorize(s string) IpType {
	ip := net.ParseIP(s)
	switch {
	case ip == nil:
		return malformed
	case ip.IsLoopback():
		return loopback
	case LoopBack0.Contains(ip):
		return loopback
	case ip.IsMulticast():
		return multicast
	case LinkLocal.Contains(ip):
		return linklocal
	case Private10.Contains(ip):
		return private
	case Private100.Contains(ip):
		return private
	case Private172.Contains(ip):
		return private
	case Private192.Contains(ip):
		return private
	case Private192Inter.Contains(ip):
		return private
	case SixToFour.Contains(ip):
		return six2four
	case Doc192.Contains(ip):
		return documentation
	case Doc198.Contains(ip):
		return documentation
	case Doc203.Contains(ip):
		return documentation
	case Reserved.Contains(ip):
		return reserved
	default:
		return public
	}
}

func Describe(ip IpType) string {
	return name[ip]
}

func main() {
	ip := Categorize(os.Args[1])
	fmt.Println(Describe(ip))
}
