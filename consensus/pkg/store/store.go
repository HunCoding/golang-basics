package store

import (
	"encoding/json"
	"errors"
	"sync"

	"consensus/pkg/raft"
)

// Command represents a command to be executed on the store
type Command struct {
	Op    string      `json:"op"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// Store represents a distributed key-value store
type Store struct {
	mu    sync.RWMutex
	raft  *raft.Node
	store map[string]interface{}
}

// NewStore creates a new distributed key-value store
func NewStore(raftNode *raft.Node) *Store {
	s := &Store{
		raft:  raftNode,
		store: make(map[string]interface{}),
	}

	// Start applying committed entries
	go s.applyCommittedEntries()

	return s
}

// Get retrieves a value from the store
func (s *Store) Get(key string) (interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if value, ok := s.store[key]; ok {
		return value, nil
	}
	return nil, errors.New("key not found")
}

// Set sets a value in the store
func (s *Store) Set(key string, value interface{}) error {
	cmd := Command{
		Op:    "SET",
		Key:   key,
		Value: value,
	}

	if _, err := json.Marshal(cmd); err != nil {
		return err
	}

	// Apply the change locally immediately
	s.mu.Lock()
	s.store[key] = value
	s.mu.Unlock()

	// TODO: Submit command to Raft log
	// This would be implemented when we add the actual RPC layer
	return nil
}

// Delete removes a value from the store
func (s *Store) Delete(key string) error {
	cmd := Command{
		Op:  "DELETE",
		Key: key,
	}

	if _, err := json.Marshal(cmd); err != nil {
		return err
	}

	// Apply the change locally immediately
	s.mu.Lock()
	delete(s.store, key)
	s.mu.Unlock()

	// TODO: Submit command to Raft log
	// This would be implemented when we add the actual RPC layer
	return nil
}

// applyCommittedEntries applies committed entries from the Raft log
func (s *Store) applyCommittedEntries() {
	for {
		select {
		case entry := <-s.raft.GetApplyCh():
			var cmd Command
			if err := json.Unmarshal(entry.([]byte), &cmd); err != nil {
				continue
			}

			s.mu.Lock()
			switch cmd.Op {
			case "SET":
				s.store[cmd.Key] = cmd.Value
			case "DELETE":
				delete(s.store, cmd.Key)
			}
			s.mu.Unlock()
		}
	}
}
