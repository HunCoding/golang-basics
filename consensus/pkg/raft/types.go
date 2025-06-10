package raft

import (
	"time"
)

// NodeState represents the current state of a Raft node
type NodeState int

const (
	Follower NodeState = iota
	Candidate
	Leader
)

// LogEntry represents a single entry in the Raft log
type LogEntry struct {
	Term    int
	Command interface{}
}

// RequestVoteArgs represents the arguments for RequestVote RPC
type RequestVoteArgs struct {
	Term         int
	CandidateID  string
	LastLogIndex int
	LastLogTerm  int
}

// RequestVoteReply represents the reply for RequestVote RPC
type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

// AppendEntriesArgs represents the arguments for AppendEntries RPC
type AppendEntriesArgs struct {
	Term         int
	LeaderID     string
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}

// AppendEntriesReply represents the reply for AppendEntries RPC
type AppendEntriesReply struct {
	Term    int
	Success bool
}

// Config holds the configuration for a Raft node
type Config struct {
	ID                string
	Peers             []string
	ElectionTimeout   time.Duration
	HeartbeatInterval time.Duration
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		ElectionTimeout:   time.Millisecond * 300,
		HeartbeatInterval: time.Millisecond * 100,
	}
}
