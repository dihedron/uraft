package cluster

import "github.com/dihedron/uraft/unmarshal"

type Peer struct {
	ID      string
	Address string
}

func (p *Peer) UnmarshalFlag(value string) error {
	return unmarshal.FromFlag(value, p)
}
