package raft

import (
	"log"
	"net"
	"net/rpc"
)

// RPCServer represents the RPC server for a Raft node
type RPCServer struct {
	node *Node
}

// RequestVote handles the RequestVote RPC
func (s *RPCServer) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) error {
	s.node.mu.Lock()
	defer s.node.mu.Unlock()

	reply.Term = s.node.currentTerm
	reply.VoteGranted = false

	// If the request is from an old term, reject it
	if args.Term < s.node.currentTerm {
		log.Printf("Node %s rejecting vote request from %s for term %d (current term: %d)",
			s.node.config.ID, args.CandidateID, args.Term, s.node.currentTerm)
		return nil
	}

	// If we discover a higher term, become follower
	if args.Term > s.node.currentTerm {
		log.Printf("Node %s discovered higher term %d from %s, becoming follower",
			s.node.config.ID, args.Term, args.CandidateID)
		s.node.becomeFollower(args.Term)
	}

	// Check if we can vote for this candidate
	if s.node.votedFor == "" || s.node.votedFor == args.CandidateID {
		// Check if candidate's log is at least as up-to-date as ours
		lastLogIndex := len(s.node.log) - 1
		lastLogTerm := s.node.log[lastLogIndex].Term

		if args.LastLogTerm > lastLogTerm ||
			(args.LastLogTerm == lastLogTerm && args.LastLogIndex >= lastLogIndex) {
			s.node.votedFor = args.CandidateID
			reply.VoteGranted = true
			log.Printf("Node %s granting vote to %s for term %d",
				s.node.config.ID, args.CandidateID, args.Term)
		} else {
			log.Printf("Node %s rejecting vote for %s: log not up-to-date",
				s.node.config.ID, args.CandidateID)
		}
	} else {
		log.Printf("Node %s already voted for %s in term %d",
			s.node.config.ID, s.node.votedFor, s.node.currentTerm)
	}

	return nil
}

// AppendEntries handles the AppendEntries RPC
func (s *RPCServer) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) error {
	s.node.mu.Lock()
	defer s.node.mu.Unlock()

	reply.Term = s.node.currentTerm
	reply.Success = false

	// If the request is from an old term, reject it
	if args.Term < s.node.currentTerm {
		log.Printf("Node %s rejecting append entries from %s for term %d (current term: %d)",
			s.node.config.ID, args.LeaderID, args.Term, s.node.currentTerm)
		return nil
	}

	// If we discover a higher term, become follower
	if args.Term > s.node.currentTerm {
		log.Printf("Node %s discovered higher term %d from %s, becoming follower",
			s.node.config.ID, args.Term, args.LeaderID)
		s.node.becomeFollower(args.Term)
	}

	// Reset election timer
	s.node.votedFor = ""

	// Check if log is consistent
	if args.PrevLogIndex >= len(s.node.log) {
		log.Printf("Node %s rejecting append entries: prevLogIndex %d >= log length %d",
			s.node.config.ID, args.PrevLogIndex, len(s.node.log))
		return nil
	}

	if args.PrevLogIndex > 0 && s.node.log[args.PrevLogIndex].Term != args.PrevLogTerm {
		log.Printf("Node %s rejecting append entries: term mismatch at index %d",
			s.node.config.ID, args.PrevLogIndex)
		return nil
	}

	// Append new entries
	if len(args.Entries) > 0 {
		s.node.log = s.node.log[:args.PrevLogIndex+1]
		s.node.log = append(s.node.log, args.Entries...)
		log.Printf("Node %s appended %d entries from leader %s",
			s.node.config.ID, len(args.Entries), args.LeaderID)
	}

	// Update commit index
	if args.LeaderCommit > s.node.commitIndex {
		s.node.commitIndex = min(args.LeaderCommit, len(s.node.log)-1)
		log.Printf("Node %s updated commit index to %d",
			s.node.config.ID, s.node.commitIndex)
	}

	reply.Success = true
	return nil
}

// StartRPCServer starts the RPC server for the node
func (n *Node) StartRPCServer(addr string) error {
	server := &RPCServer{node: n}
	rpc.Register(server)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v", err)
				continue
			}
			go rpc.ServeConn(conn)
		}
	}()

	return nil
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
