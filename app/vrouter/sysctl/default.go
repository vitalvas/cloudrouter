package sysctl

var defaults = map[string]string{
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

// TODO: dynamic calculate net.netfilter.nf_conntrack_buckets + net.netfilter.nf_conntrack_max

func Params(kv map[string]string) map[string]string {
	data := make(map[string]string)

	for key, value := range defaults {
		data[key] = value
	}

	if kv != nil {
		for key, value := range kv {
			data[key] = value
		}
	}

	return data
}
