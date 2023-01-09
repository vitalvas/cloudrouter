package general

import (
	"log"
	"net"
)

var (
	privateNets   []net.IPNet
	ipv6LinkLocal net.IPNet
)

var PrivateNetworks = []string{
	// loopback: https://tools.ietf.org/html/rfc3330#section-2
	"127.0.0.0/8",
	// loopback: https://tools.ietf.org/html/rfc3513#section-2.4
	"::1/128",

	// reserved: https://tools.ietf.org/html/rfc1918#section-3
	"10.0.0.0/8",
	"172.16.0.0/12",
	"192.168.0.0/16",
	"192.0.2.0/24",

	// reserved: https://datatracker.ietf.org/doc/html/rfc6598#section-7
	"100.64.0.0/10",

	// reserved: https://tools.ietf.org/html/rfc4193#section-3.1
	"fc00::/7",

	// link-local: https://tools.ietf.org/html/rfc3927#section-1.2
	"169.254.0.0/16",
	// link-local: https://tools.ietf.org/html/rfc4291#section-2.4
	"fe80::/10",
}

func init() {
	privateNets = make([]net.IPNet, len(PrivateNetworks))
	for idx, s := range PrivateNetworks {
		_, net, err := net.ParseCIDR(s)
		if err != nil {
			log.Fatal(err)
		}
		privateNets[idx] = *net

		if s == "fe80::/10" {
			ipv6LinkLocal = *net
		}
	}
}

func isPrivate(iface string, ipaddr net.IP) bool {
	for _, n := range privateNets {
		if n.Contains(ipaddr) {
			return true
		}
	}

	return false
}

func interfaceAddrs(keep func(string, net.IP) bool) ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var hosts []string

	for _, i := range ifaces {
		if i.Flags&net.FlagUp != net.FlagUp {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, a := range addrs {
			ipaddr, _, err := net.ParseCIDR(a.String())
			if err != nil {
				return nil, err
			}

			if !keep(i.Name, ipaddr) {
				continue
			}

			host := ipaddr.String()
			if ipv6LinkLocal.Contains(ipaddr) {
				host = host + "%" + i.Name
			}

			hosts = append(hosts, host)
		}
	}

	return hosts, nil
}

func PrivateInterfaceAddrs() ([]string, error) {
	return interfaceAddrs(isPrivate)
}

func PublicInterfaceAddrs() ([]string, error) {
	return interfaceAddrs(func(iface string, addr net.IP) bool {
		return !isPrivate(iface, addr)
	})
}
