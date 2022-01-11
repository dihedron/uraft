package command

import "fmt"

type Options struct {

	// Directory is the directory for Raft cluster state storage.
	Directory string `short:"d" long:"directory" description:"Raft cluster state storage directory." optional:"yes" default:"./state"`

	// Address is the intra-cluster bind address fro Raft communications.
	Address string `short:"a" long:"address" description:"Raft intra-cluster address." optional:"yes" default:":8001"`

	// Join specified whether the node should join a cluster.
	Peers []string `short:"p" long:"peer" description:"The address of a peer node in the cluster to join" optional:"yes"`
}

func (cmd *Options) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("no node id specified")
	}
	return nil
}
