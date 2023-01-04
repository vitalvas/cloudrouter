package netconfig

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/vishvananda/netlink"
	"github.com/vitalvas/cloudrouter/lib/general"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type WireguardInterfaces struct {
	Interfaces []WireguardInterface `json:"interfaces"`
}

type Wireguard struct {
	client *wgctrl.Client
	cfg    WireguardInterfaces
}

type WireguardInterface struct {
	Name       string          `json:"name"`
	PrivateKey string          `json:"private_key"`
	Port       int             `json:"port"`
	Peers      []WireguardPeer `json:"peers"`
}

type WireguardPeer struct {
	PublicKey                   string        `json:"public_key"`
	Endpoint                    string        `json:"endpoint"`
	AllowedIPs                  []string      `json:"allowed_ips"`
	PresharedKey                string        `json:"preshared_key"`
	PersistentKeepaliveInterval time.Duration `json:"persistent_keepalive_interval"`
}

func NewWireguard() (*Wireguard, error) {
	wg := &Wireguard{}

	var err error

	wg.client, err = wgctrl.New()
	if err != nil {
		return nil, err
	}

	return wg, nil
}

func (wg *Wireguard) Close() {
	wg.client.Close()
}

func (wg *Wireguard) apply() error {
	b, err := os.ReadFile(filepath.Join(general.ConfigDir, "netconfig_wireguard.json"))
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if err := json.Unmarshal(b, &wg.cfg); err != nil {
			return fmt.Errorf("cannot unmarshal config: %w", err)
		}
	}

	var link netlink.Link
	for _, iface := range wg.cfg.Interfaces {
		if err := wg.createInterface(iface.Name); err != nil {
			return err
		}

		link, err = netlink.LinkByName(iface.Name)
		if err != nil {
			return fmt.Errorf("cannot read link: %w", err)
		}

		if err := netlink.LinkSetUp(link); err != nil {
			return fmt.Errorf("cannot set link up: %w", err)
		}

		if err := wg.syncInterface(iface); err != nil {
			return err
		}
	}

	return wg.clean()
}

func (wg *Wireguard) createInterface(name string) error {
	_, err := netlink.LinkByName(name)
	// skip if link exists
	if err == nil {
		return nil
	}

	if _, ok := err.(netlink.LinkNotFoundError); !ok {
		return fmt.Errorf("cannot read link: %w", err)
	}

	wgLink := &netlink.GenericLink{
		LinkAttrs: netlink.LinkAttrs{
			Name: name,
		},
		LinkType: "wireguard",
	}

	if err := netlink.LinkAdd(wgLink); err != nil {
		return fmt.Errorf("cannot create link: %w", err)
	}

	// Needs some sleeping to wait for interface creating.
	time.Sleep(2 * time.Second)

	return nil
}

func (wg *Wireguard) syncInterface(iface WireguardInterface) error {
	cfg := wgtypes.Config{
		ReplacePeers: true,
	}

	key, err := wgtypes.ParseKey(iface.PrivateKey)
	if err != nil {
		return fmt.Errorf("cannot parse private key: %w", err)
	}

	cfg.PrivateKey = &key

	if iface.Port > 0 {
		cfg.ListenPort = &iface.Port
	}

	for _, peer := range iface.Peers {
		b, err := base64.StdEncoding.DecodeString(peer.PublicKey)
		if err != nil {
			return fmt.Errorf("cannot decode public key: %w", err)
		}

		publicKey, err := wgtypes.NewKey(b)
		if err != nil {
			return fmt.Errorf("cannot parse public key: %w", err)
		}

		var addr *net.UDPAddr
		if peer.Endpoint != "" {
			addr, err = net.ResolveUDPAddr("udp", peer.Endpoint)
			if err != nil {
				return fmt.Errorf("cannot resolve endpoint: %w", err)
			}
		}

		var ips []net.IPNet
		for _, ip := range peer.AllowedIPs {
			_, ipnet, err := net.ParseCIDR(ip)
			if err != nil {
				return fmt.Errorf("cannot prase allowed-ips: %w", err)
			}
			ips = append(ips, *ipnet)
		}

		var presharedKey wgtypes.Key
		if peer.PresharedKey != "" {
			b, err := base64.StdEncoding.DecodeString(peer.PublicKey)
			if err != nil {
				return fmt.Errorf("cannot decode preshred key: %w", err)
			}

			presharedKey, err = wgtypes.NewKey(b)
			if err != nil {
				return fmt.Errorf("cannot parse public key: %w", err)
			}
		}

		peerCfg := wgtypes.PeerConfig{
			PublicKey:         publicKey,
			Endpoint:          addr,
			ReplaceAllowedIPs: true,
			AllowedIPs:        ips,
			PresharedKey:      &presharedKey,
		}

		if peer.PersistentKeepaliveInterval > time.Second {
			peerCfg.PersistentKeepaliveInterval = &peer.PersistentKeepaliveInterval
		}

		cfg.Peers = append(cfg.Peers, peerCfg)
	}

	if err := wg.client.ConfigureDevice(iface.Name, cfg); err != nil {
		return fmt.Errorf("cannot configure device: %w", err)
	}

	return nil
}

func (wg *Wireguard) clean() error {
	devs, err := wg.client.Devices()
	if err != nil {
		log.Fatal(err)
	}

	var unknownDevices []string

	for _, dev := range devs {
		exists := false
		for _, row := range wg.cfg.Interfaces {
			if dev.Name == row.Name {
				exists = true
			}
		}

		if !exists {
			unknownDevices = append(unknownDevices, dev.Name)
		}
	}

	for _, row := range unknownDevices {
		log.Println("deleting", row)

		link, err := netlink.LinkByName(row)
		if err != nil {
			return fmt.Errorf("cannot get link: %w", err)
		}

		if err := netlink.LinkDel(link); err != nil {
			return fmt.Errorf("cannot delete link: %w", err)
		}
	}

	return nil
}
