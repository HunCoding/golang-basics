package raft

import (
	"log"
	"math/rand"
	"net/rpc"
	"sync"
	"time"
)

type Node struct {
	mu sync.Mutex

	config *Config

	// Persistent state
	currentTerm int
	votedFor    string
	log         []LogEntry

	// Volatile state
	state       NodeState
	commitIndex int
	lastApplied int

	// Leader state
	nextIndex  map[string]int
	matchIndex map[string]int

	// Channels
	applyCh chan interface{}
	stopCh  chan struct{}
}

// NewNode creates a new Raft node
func NewNode(config *Config) *Node {
	if config == nil {
		config = DefaultConfig()
	}

	node := &Node{
		config:     config,
		state:      Follower,
		nextIndex:  make(map[string]int),
		matchIndex: make(map[string]int),
		applyCh:    make(chan interface{}, 100),
		stopCh:     make(chan struct{}),
	}

	// Initialize log with a dummy entry at index 0
	node.log = append(node.log, LogEntry{Term: 0})

	return node
}

// Start begins the Raft node's main loop
func (n *Node) Start() {
	go n.runElectionTimer()
}

// Stop stops the Raft node
func (n *Node) Stop() {
	close(n.stopCh)
}

// runElectionTimer runs the election timer
func (n *Node) runElectionTimer() {
	for {
		timeout := n.config.ElectionTimeout + time.Duration(rand.Int63n(int64(n.config.ElectionTimeout)))
		log.Printf("Node %s waiting for %v before next election check", n.config.ID, timeout)

		select {
		case <-time.After(timeout):
			log.Printf("Node %s about to acquire lock in runElectionTimer", n.config.ID)
			n.mu.Lock()
			log.Printf("Node %s acquired lock in runElectionTimer", n.config.ID)
			if n.state != Leader {
				log.Printf("Node %s election timeout, current state: %v", n.config.ID, n.state)
				go n.startElection()
			}
			log.Printf("Node %s about to release lock in runElectionTimer", n.config.ID)
			n.mu.Unlock()
			log.Printf("Node %s released lock in runElectionTimer", n.config.ID)
		case <-n.stopCh:
			return
		}
	}
}

// startElection starts a new election
func (n *Node) startElection() {
	log.Printf("Node %s ENTERED startElection", n.config.ID)

	n.mu.Lock()
	log.Printf("Node %s acquired lock in startElection", n.config.ID)

	if n.state == Leader {
		log.Printf("Node %s is already leader, skipping election", n.config.ID)
		n.mu.Unlock()
		log.Printf("Node %s EXITING startElection", n.config.ID)
		return
	}

	log.Printf("Node %s about to set state to Candidate", n.config.ID)
	n.state = Candidate
	log.Printf("Node %s set state to Candidate", n.config.ID)

	n.currentTerm++
	log.Printf("Node %s incremented currentTerm to %d", n.config.ID, n.currentTerm)

	n.votedFor = n.config.ID
	log.Printf("Node %s set votedFor to %s", n.config.ID, n.votedFor)

	currentTerm := n.currentTerm
	log.Printf("Node %s set currentTerm local var", n.config.ID)

	if len(n.log) == 0 {
		log.Printf("Node %s WARNING: log is empty!", n.config.ID)
	}

	log.Printf("Node %s about to unlock mutex", n.config.ID)
	n.mu.Unlock()
	log.Printf("Node %s unlocked mutex", n.config.ID)

	log.Printf("Node %s starting election for term %d", n.config.ID, currentTerm)

	// Request votes from all peers
	votes := 1 // Vote for self
	votesCh := make(chan bool, len(n.config.Peers))
	responses := 0
	responsesCh := make(chan struct{}, len(n.config.Peers))

	// Send vote requests to all peers
	for _, peer := range n.config.Peers {
		log.Printf("Node %s sending RequestVote to %s", n.config.ID, peer)
		go func(peer string) {
			args := &RequestVoteArgs{
				Term:         currentTerm,
				CandidateID:  n.config.ID,
				LastLogIndex: len(n.log) - 1,
				LastLogTerm:  n.log[len(n.log)-1].Term,
			}
			reply := &RequestVoteReply{}

			if err := n.sendRequestVote(peer, args, reply); err == nil {
				n.mu.Lock()
				defer n.mu.Unlock()

				// Check if we're still a candidate for the same term
				if n.state != Candidate || n.currentTerm != currentTerm {
					return
				}

				responses++
				responsesCh <- struct{}{}

				if reply.VoteGranted {
					log.Printf("Node %s received vote from %s for term %d", n.config.ID, peer, currentTerm)
					votesCh <- true
				} else if reply.Term > currentTerm {
					log.Printf("Node %s discovered higher term %d from %s, becoming follower", n.config.ID, reply.Term, peer)
					n.becomeFollower(reply.Term)
				} else {
					log.Printf("Node %s did not receive vote from %s for term %d", n.config.ID, peer, currentTerm)
				}
			} else {
				log.Printf("Node %s failed to get vote from %s: %v", n.config.ID, peer, err)
				responses++
				responsesCh <- struct{}{}
			}
		}(peer)
	}

	// Wait for votes or timeout
	timeout := time.After(n.config.ElectionTimeout)
	for {
		select {
		case <-votesCh:
			votes++
			log.Printf("Node %s now has %d votes for term %d", n.config.ID, votes, currentTerm)
			if votes > (len(n.config.Peers)+1)/2 {
				log.Printf("Node %s received majority of votes (%d) for term %d", n.config.ID, votes, currentTerm)
				n.mu.Lock()
				if n.state == Candidate && n.currentTerm == currentTerm {
					n.becomeLeader()
				}
				n.mu.Unlock()
				log.Printf("Node %s EXITING startElection (became leader)", n.config.ID)
				return
			}
		case <-responsesCh:
			if responses == len(n.config.Peers) {
				log.Printf("Node %s received all responses but only has %d votes for term %d",
					n.config.ID, votes, currentTerm)
				log.Printf("Node %s EXITING startElection (all responses)", n.config.ID)
				return
			}
		case <-timeout:
			log.Printf("Node %s election timeout for term %d (received %d/%d responses)",
				n.config.ID, currentTerm, responses, len(n.config.Peers))
			log.Printf("Node %s EXITING startElection (timeout)", n.config.ID)
			return
		}
	}
}

// becomeLeader transitions the node to leader state
func (n *Node) becomeLeader() {
	if n.state == Leader {
		log.Printf("Node %s is already leader for term %d", n.config.ID, n.currentTerm)
		return
	}

	n.state = Leader
	log.Printf("Node %s became leader for term %d", n.config.ID, n.currentTerm)

	// Initialize leader state
	for _, peer := range n.config.Peers {
		n.nextIndex[peer] = len(n.log)
		n.matchIndex[peer] = 0
	}

	// Start sending heartbeats immediately
	go n.sendHeartbeats()
}

// becomeFollower transitions the node to follower state
func (n *Node) becomeFollower(term int) {
	if n.state == Follower && n.currentTerm == term {
		return
	}

	n.state = Follower
	n.currentTerm = term
	n.votedFor = ""
	log.Printf("Node %s became follower for term %d", n.config.ID, n.currentTerm)
}

// sendHeartbeats sends heartbeats to all peers
func (n *Node) sendHeartbeats() {
	for {
		n.mu.Lock()
		if n.state != Leader {
			n.mu.Unlock()
			return
		}
		currentTerm := n.currentTerm
		n.mu.Unlock()

		// Send heartbeats to all peers
		for _, peer := range n.config.Peers {
			go func(peer string) {
				n.mu.Lock()
				args := &AppendEntriesArgs{
					Term:         currentTerm,
					LeaderID:     n.config.ID,
					PrevLogIndex: n.nextIndex[peer] - 1,
					PrevLogTerm:  n.log[n.nextIndex[peer]-1].Term,
					Entries:      n.log[n.nextIndex[peer]:],
					LeaderCommit: n.commitIndex,
				}
				n.mu.Unlock()

				reply := &AppendEntriesReply{}
				if err := n.sendAppendEntries(peer, args, reply); err == nil {
					n.mu.Lock()
					defer n.mu.Unlock()

					if n.state != Leader || n.currentTerm != currentTerm {
						return
					}

					if reply.Success {
						n.matchIndex[peer] = args.PrevLogIndex + len(args.Entries)
						n.nextIndex[peer] = n.matchIndex[peer] + 1
						n.updateCommitIndex()
					} else if reply.Term > currentTerm {
						log.Printf("Node %s discovered higher term %d from %s, becoming follower", n.config.ID, reply.Term, peer)
						n.becomeFollower(reply.Term)
					} else {
						n.nextIndex[peer] = max(1, n.nextIndex[peer]-1)
					}
				}
			}(peer)
		}

		time.Sleep(n.config.HeartbeatInterval)
	}
}

// updateCommitIndex updates the commit index
func (n *Node) updateCommitIndex() {
	for i := n.commitIndex + 1; i < len(n.log); i++ {
		if n.log[i].Term == n.currentTerm {
			count := 1
			for _, peer := range n.config.Peers {
				if n.matchIndex[peer] >= i {
					count++
				}
			}
			if count > len(n.config.Peers)/2 {
				n.commitIndex = i
			}
		}
	}
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// sendRequestVote sends a RequestVote RPC to a peer
func (n *Node) sendRequestVote(peer string, args *RequestVoteArgs, reply *RequestVoteReply) error {
	client, err := rpc.Dial("tcp", peer)
	if err != nil {
		return err
	}
	defer client.Close()

	return client.Call("RPCServer.RequestVote", args, reply)
}

// sendAppendEntries sends an AppendEntries RPC to a peer
func (n *Node) sendAppendEntries(peer string, args *AppendEntriesArgs, reply *AppendEntriesReply) error {
	client, err := rpc.Dial("tcp", peer)
	if err != nil {
		return err
	}
	defer client.Close()

	return client.Call("RPCServer.AppendEntries", args, reply)
}

// GetApplyCh returns the channel for applying committed entries
func (n *Node) GetApplyCh() <-chan interface{} {
	return n.applyCh
}

// GetState returns the current state and term of the node
func (n *Node) GetState() (NodeState, int) {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.state, n.currentTerm
}
