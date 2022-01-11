package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dihedron/uraft/command"
	"github.com/jessevdk/go-flags"
)

// Command line defaults
const (
	DefaultHTTPAddr = ":11000"
	DefaultRaftAddr = ":12000"
)

// Command line parameters
var inmem bool
var httpAddr string
var raftAddr string
var joinAddr string
var nodeID string

func init() {
	flag.BoolVar(&inmem, "inmem", false, "Use in-memory storage for Raft")
	flag.StringVar(&httpAddr, "haddr", DefaultHTTPAddr, "Set the HTTP bind address")
	flag.StringVar(&raftAddr, "raddr", DefaultRaftAddr, "Set Raft bind address")
	flag.StringVar(&joinAddr, "join", "", "Set join address, if any")
	flag.StringVar(&nodeID, "id", "", "Node ID")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <raft-data-path> \n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	var (
		options command.Options
		parser  = flags.NewParser(&options, flags.Default)
		args    []string
		err     error
	)
	if args, err = parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		default:
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
	if err = options.Execute(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

/*
func main2() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		os.Exit(1)
	}

	// Ensure Raft storage exists.
	raftDir := flag.Arg(0)
	if raftDir == "" {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		os.Exit(1)
	}
	os.MkdirAll(raftDir, 0700)

	s := store.New(inmem)
	s.RaftDir = raftDir
	s.RaftBind = raftAddr
	if err := s.Open(joinAddr == "", nodeID); err != nil {
		log.Fatalf("failed to open store: %s", err.Error())
	}

	// If join was specified, make the join request.
	if joinAddr != "" {
		if err := join(joinAddr, raftAddr, nodeID); err != nil {
			log.Fatalf("failed to join node at %s: %s", joinAddr, err.Error())
		}
	}

	log.Println("hraftd started successfully")

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	log.Println("hraftd exiting")
}

func join(joinAddr, raftAddr, nodeID string) error {
	b, err := json.Marshal(map[string]string{"addr": raftAddr, "id": nodeID})
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/join", joinAddr), "application-type/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
*/
