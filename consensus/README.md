# Distributed Key-Value Store with Raft Consensus

This project implements a distributed key-value store using the Raft consensus algorithm in Go. It demonstrates how to build a fault-tolerant distributed system that maintains consistency across multiple nodes.

## Features

- Raft consensus algorithm implementation
- Distributed key-value store
- HTTP API for interacting with the store
- Leader election and log replication
- Fault tolerance and consistency guarantees

## Project Structure

```
.
├── cmd/
│   └── server/         # Main server application
├── pkg/
│   ├── api/           # HTTP API implementation
│   ├── raft/          # Raft consensus implementation
│   └── store/         # Key-value store implementation
└── README.md
```

## Building

```bash
go build -o server ./cmd/server
```

## Running

To run a single node:

```bash
./server -id node1 -http :8080
```

To run multiple nodes (in different terminals):

```bash
# Terminal 1
./server -id node1 -http :8080 -peers node2,node3

# Terminal 2
./server -id node2 -http :8081 -peers node1,node3

# Terminal 3
./server -id node3 -http :8082 -peers node1,node2
```

## API Endpoints

### Get Value
```bash
curl "http://localhost:8080/get?key=mykey"
```

### Set Value
```bash
curl -X POST http://localhost:8080/set \
  -H "Content-Type: application/json" \
  -d '{"key": "mykey", "value": "myvalue"}'
```

### Delete Value
```bash
curl -X DELETE "http://localhost:8080/delete?key=mykey"
```

## Implementation Details

### Raft Consensus

The implementation follows the Raft consensus algorithm as described in the paper "In Search of an Understandable Consensus Algorithm" by Diego Ongaro and John Ousterhout. Key features include:

- Leader election
- Log replication
- Safety guarantees
- Membership changes

### Key-Value Store

The key-value store is built on top of the Raft consensus layer, providing:

- Strong consistency
- Fault tolerance
- Simple HTTP API
- Concurrent access

## Contributing

Feel free to submit issues and enhancement requests! 