package cluster

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

const (
	retainSnapshotCount = 2
	raftTimeout         = 10 * time.Second
)

type Cluster struct {
	id        string
	directory string
	address   string
	peers     []Peer
	raft      *raft.Raft
}

// New creates a new Cluster, applying all the provided functional options.
func New(id string, fsm raft.FSM, options ...Option) (*Cluster, error) {
	c := &Cluster{
		id:    id,
		peers: []Peer{},
	}
	for _, option := range options {
		option(c)
	}

	err := os.MkdirAll(c.directory, 0700)
	if err != nil {
		return nil, fmt.Errorf("error creating raft state directory '%s': %w", c.directory, err)
	}

	// setup Raft configuration
	configuration := raft.DefaultConfig()
	configuration.LocalID = raft.ServerID(c.id)

	// Setup Raft communication.
	addr, err := net.ResolveTCPAddr("tcp", c.address)
	if err != nil {
		return nil, fmt.Errorf("error resolving bind address '%s': %w", c.address, err)
	}
	transport, err := raft.NewTCPTransport(c.address, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("error creating TCP transport: %w", err)
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(c.directory, retainSnapshotCount, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("error creating file snapshot store: %w", err)
	}

	// create the BoltDB instance for both log store and stable store
	boltDB, err := raftboltdb.NewBoltStore(filepath.Join(c.directory, "raft.db"))
	if err != nil {
		return nil, fmt.Errorf("error creating new Bolt store: %w", err)
	}

	// instantiate the Raft system
	c.raft, err = raft.NewRaft(configuration, fsm, boltDB, boltDB, snapshots, transport)
	if err != nil {
		return nil, fmt.Errorf("error creating new raft cluster: %w", err)
	}

	servers := []raft.Server{
		{
			ID:      raft.ServerID(c.id),
			Address: transport.LocalAddr(),
		},
	}

	if len(c.peers) > 0 {
		for _, peer := range c.peers {
			servers = append(servers, raft.Server{
				ID:      raft.ServerID(peer.ID),
				Address: raft.ServerAddress(peer.Address.String()),
			})
		}
	}

	cluster := raft.Configuration{
		Servers: servers,
	}
	if err := c.raft.BootstrapCluster(cluster).Error(); err != nil {
		// maybe already bootstrapped??
	}

	return c, err
}
