package netconfig

import (
	"github.com/google/nftables"
)

type Firewall struct {
	conn *nftables.Conn
}

func NewFirewall() (*Firewall, error) {
	fw := &Firewall{}

	var err error
	fw.conn, err = nftables.New()
	if err != nil {
		return nil, err
	}

	return fw, nil
}

func (fw *Firewall) apply() error {
	fw.conn.FlushRuleset()

	table4 := &nftables.Table{
		Family: nftables.TableFamilyIPv4,
		Name:   "filter",
	}

	table6 := &nftables.Table{
		Family: nftables.TableFamilyIPv6,
		Name:   "filter",
	}

	for _, table := range []*nftables.Table{table4, table6} {
		fw.applyTableFilter(table)
	}

	nat4 := &nftables.Table{
		Family: nftables.TableFamilyIPv4,
		Name:   "nat",
	}

	for _, table := range []*nftables.Table{nat4} {
		fw.applyTableNAT(table)
	}

	return fw.conn.Flush()
}

func (fw *Firewall) applyTableFilter(table *nftables.Table) {
	filter := fw.conn.AddTable(table)

	fw.conn.AddChain(&nftables.Chain{
		Name:     "input",
		Hooknum:  nftables.ChainHookInput,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeFilter,
	})

	fw.conn.AddChain(&nftables.Chain{
		Name:     "forward",
		Hooknum:  nftables.ChainHookForward,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeFilter,
	})

	fw.conn.AddChain(&nftables.Chain{
		Name:     "output",
		Hooknum:  nftables.ChainHookOutput,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeFilter,
	})

}

func (fw *Firewall) applyTableNAT(table *nftables.Table) {
	filter := fw.conn.AddTable(table)

	fw.conn.AddChain(&nftables.Chain{
		Name:     "prerouting",
		Hooknum:  nftables.ChainHookPrerouting,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeNAT,
	})

	fw.conn.AddChain(&nftables.Chain{
		Name:     "postrouting",
		Hooknum:  nftables.ChainHookPostrouting,
		Priority: nftables.ChainPriorityNATSource,
		Table:    filter,
		Type:     nftables.ChainTypeNAT,
	})
}
