package command

import (
	"fmt"
	"os"

	"github.com/dihedron/uraft/cache"
	"github.com/dihedron/uraft/cluster"
	"github.com/dihedron/uraft/logging"
)

type Options struct {
	// Address is the intra-cluster bind address for Raft communications.
	Address cluster.Address `short:"a" long:"address" description:"Raft intra-cluster address." optional:"yes" default:"localhost:8001"`
	// Join specified whether the node should join a cluster.
	Peers []cluster.Peer `short:"p" long:"peer" description:"The address of a peer node in the cluster to join" optional:"yes"`
	// State is the directory for Raft cluster state storage.
	State string `short:"s" long:"state" description:"Raft cluster state storage directory." optional:"yes" default:"./state"`
}

func (cmd *Options) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("no node id specified")
	}
	fmt.Printf("starting a node at '%s' (state in directory '%s'), with peers %+v\n", cmd.Address, cmd.State, cmd.Peers)

	logger := logging.NewConsoleLogger(os.Stdout)

	fsm := cache.New()

	c, err := cluster.New(
		args[0],
		fsm,
		cluster.WithDirectory(cmd.State),
		cluster.WithBindAddress(cmd.Address.String()),
		cluster.WithPeers(cmd.Peers...),
		cluster.WithLogger(logger),
	)
	if err != nil {
		return fmt.Errorf("error creating new cluster: %w", err)
	}
	c.Test()

	return nil
}
