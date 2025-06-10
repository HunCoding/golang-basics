package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"consensus/pkg/api"
	"consensus/pkg/raft"
	"consensus/pkg/store"
)

func main() {
	// Parse command line flags
	nodeID := flag.String("id", "", "Node ID")
	httpAddr := flag.String("http", ":8080", "HTTP server address")
	rpcAddr := flag.String("rpc", ":9090", "RPC server address")
	peers := flag.String("peers", "", "Comma-separated list of peer addresses")
	flag.Parse()

	if *nodeID == "" {
		log.Fatal("Node ID is required")
	}

	// Create Raft node
	config := raft.DefaultConfig()
	config.ID = *nodeID

	// Parse peers
	if *peers != "" {
		config.Peers = strings.Split(*peers, ",")
	}

	raftNode := raft.NewNode(config)

	// Start RPC server
	if err := raftNode.StartRPCServer(*rpcAddr); err != nil {
		log.Fatal(err)
	}

	raftNode.Start()

	// Create distributed store
	store := store.NewStore(raftNode)

	// Create and start HTTP server
	server := api.NewServer(store, raftNode)
	go func() {
		if err := server.Start(*httpAddr); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// Cleanup
	raftNode.Stop()
}
