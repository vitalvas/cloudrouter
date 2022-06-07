package netconfig

import (
	"github.com/google/nftables"
)

type Firewall struct {
	conn *nftables.Conn
}

func NewFirewall() (*Firewall, error) {
	this := &Firewall{}

	var err error
	this.conn, err = nftables.New()
	if err != nil {
		return nil, err
	}

	return this, nil
}

func (this *Firewall) apply() error {
	this.conn.FlushRuleset()

	table4 := &nftables.Table{
		Family: nftables.TableFamilyIPv4,
		Name:   "filter",
	}

	table6 := &nftables.Table{
		Family: nftables.TableFamilyIPv6,
		Name:   "filter",
	}

	for _, table := range []*nftables.Table{table4, table6} {
		this.applyTableFilter(table)
	}

	nat4 := &nftables.Table{
		Family: nftables.TableFamilyIPv4,
		Name:   "nat",
	}

	for _, table := range []*nftables.Table{nat4} {
		this.applyTableNAT(table)
	}

	return this.conn.Flush()
}

func (this *Firewall) applyTableFilter(table *nftables.Table) {
	filter := this.conn.AddTable(table)

	this.conn.AddChain(&nftables.Chain{
		Name:     "input",
		Hooknum:  nftables.ChainHookInput,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeFilter,
	})

	this.conn.AddChain(&nftables.Chain{
		Name:     "forward",
		Hooknum:  nftables.ChainHookForward,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeFilter,
	})

	this.conn.AddChain(&nftables.Chain{
		Name:     "output",
		Hooknum:  nftables.ChainHookOutput,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeFilter,
	})

}

func (this *Firewall) applyTableNAT(table *nftables.Table) {
	filter := this.conn.AddTable(table)

	this.conn.AddChain(&nftables.Chain{
		Name:     "prerouting",
		Hooknum:  nftables.ChainHookPrerouting,
		Priority: nftables.ChainPriorityFilter,
		Table:    filter,
		Type:     nftables.ChainTypeNAT,
	})

	this.conn.AddChain(&nftables.Chain{
		Name:     "postrouting",
		Hooknum:  nftables.ChainHookPostrouting,
		Priority: nftables.ChainPriorityNATSource,
		Table:    filter,
		Type:     nftables.ChainTypeNAT,
	})
}
