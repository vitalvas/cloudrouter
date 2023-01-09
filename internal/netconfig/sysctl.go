package netconfig

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var defaultSysctl = map[string]string{
	"net.ipv4.tcp_congestion_control": "bbr",
	"net.ipv4.ip_forward":             "1",
	"net.ipv4.conf.default.rp_filter": "1",
	"net.ipv4.conf.all.rp_filter":     "1",

	"net.ipv6.conf.all.forwarding": "1",

	"net.netfilter.nf_conntrack_max":                     "2097152",
	"net.netfilter.nf_conntrack_acct":                    "1",
	"net.netfilter.nf_conntrack_checksum":                "0",
	"net.netfilter.nf_conntrack_tcp_timeout_established": "86400",
	"net.netfilter.nf_conntrack_udp_timeout":             "60",
	"net.netfilter.nf_conntrack_udp_timeout_stream":      "300",
}

func sysctlPathFromKey(key string) string {
	return filepath.Join("/proc/sys", strings.ReplaceAll(key, ".", "/"))
}

func sysctlGet(key string) (string, error) {
	path := sysctlPathFromKey(key)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return "", os.ErrNotExist
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func sysctlSet(key, value string) error {
	return os.WriteFile(sysctlPathFromKey(key), []byte(value), 0644)
}

func (nc *NetConfig) applySysctl() error {
	for key, val := range defaultSysctl {
		current, err := sysctlGet(key)
		if err != nil {
			if err == os.ErrNotExist {
				continue
			} else {
				nc.log.Fatal(err)
			}
		}

		if val != current {
			nc.log.Println("changing", key, "from", current, "to", val)
			sysctlSet(key, val)
		}
	}

	return nil
}
