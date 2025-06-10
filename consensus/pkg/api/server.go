package api

import (
	"encoding/json"
	"log"
	"net/http"

	"consensus/pkg/raft"
	"consensus/pkg/store"
)

// Server represents the HTTP API server
type Server struct {
	store *store.Store
	raft  *raft.Node
}

// NewServer creates a new HTTP API server
func NewServer(store *store.Store, raftNode *raft.Node) *Server {
	return &Server{
		store: store,
		raft:  raftNode,
	}
}

// Start starts the HTTP server
func (s *Server) Start(addr string) error {
	http.HandleFunc("/get", s.handleGet)
	http.HandleFunc("/set", s.handleSet)
	http.HandleFunc("/delete", s.handleDelete)
	http.HandleFunc("/state", s.handleState)

	log.Printf("Starting HTTP server on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// handleState handles requests to get the node's state
func (s *Server) handleState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	state, term := s.raft.GetState()

	// Converter o estado para string
	var stateStr string
	switch state {
	case raft.Follower:
		stateStr = "Follower"
	case raft.Candidate:
		stateStr = "Candidate"
	case raft.Leader:
		stateStr = "Leader"
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"state": stateStr,
		"term":  term,
	})
}

// handleGet handles GET requests
func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	value, err := s.store.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"key":   key,
		"value": value,
	})
}

// handleSet handles SET requests
func (s *Server) handleSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	if err := s.store.Set(req.Key, req.Value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// handleDelete handles DELETE requests
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	if err := s.store.Delete(key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
