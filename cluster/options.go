package cluster

import "github.com/dihedron/uraft/logging"

// Option is the type for functional options.
type Option func(*Cluster)

// WithDirectory specifies the directory where the Raft cluster
// state is stored.
func WithDirectory(path string) Option {
	return func(c *Cluster) {
		if path != "" {
			c.directory = path
		}
	}
}

// WithBindAddress specifies the bind address used for intra-cluster
// communications.
func WithBindAddress(address string) Option {
	return func(c *Cluster) {
		if address != "" {
			c.address = address
		}
	}
}

// WithPeer specifies a peer to contact to join the cluster.
func WithPeer(peer Peer) Option {
	return func(c *Cluster) {
		c.peers = append(c.peers, peer)
	}
}

// WithPeers specifies the peers to contact to join the cluster.
func WithPeers(peers ...Peer) Option {
	return func(c *Cluster) {
		c.peers = append(c.peers, peers...)
	}
}

// WithLogger specifies a logger.
func WithLogger(logger logging.Logger) Option {
	return func(c *Cluster) {
		c.peers = append(c.peers, peer)
	}
}
