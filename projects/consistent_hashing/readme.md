# Consistent Hashing Implementation in Go

## Overview
This project implements consistent hashing, a distributed hashing scheme that enables efficient data partitioning across multiple database shards. I developed this implementation as a learning exercise to deepen my understanding of both Go programming concepts and distributed systems principles.

## What is Consistent Hashing?
Consistent hashing solves a fundamental problem in distributed databases: how to distribute data across multiple servers while minimizing redistribution when the number of servers changes. In traditional hash-based distribution (hash(key) % n, where n is the number of servers), adding or removing a server would require redistributing almost all the data. Consistent hashing reduces this redistribution to k/n of the keys, where k is the number of keys and n is the number of servers.

## Implementation Details

### Core Components

1. **Hash Ring**: 
   - Implemented as a `ring` struct that maintains:
     - A mapping of server IDs to virtual nodes
     - References to database instances
     - Configurable replication factor
     - Custom hash function support

2. **Virtual Nodes (vNodes)**:
   - Each server is represented by multiple virtual nodes on the hash ring
   - Improves load distribution and balance
   - Configurable number of vNodes per server via replication factor

3. **Database**:
   - Simulates a banking application with customer records
   - Each record contains:
     - Customer information (First Name, Last Name)
     - Account balances (Savings, Checking, 401k)
     - Pre-calculated hash position

### Key Features

- **Dynamic Server Management**: Support for adding and removing database servers
- **Automatic Rebalancing**: Data redistribution when topology changes
- **Thread-Safe Operations**: Concurrent record additions and modifications
- **Configurable Hashing**: Supports custom hash functions with MD5 as default
- **Data Consistency**: Ensures data integrity during rebalancing operations

## Go Programming Concepts Utilized

This project helped me learn and implement various Go-specific features and patterns:

1. **Concurrency Primitives**:
   - **Goroutines**: Lightweight threads for concurrent operations
   - **WaitGroups**: Synchronization primitive for coordinating multiple goroutines
   - **Mutex**: Thread synchronization for safe access to shared resources

2. **Pointer Operations**:
   - Extensive use of pointer receivers for methods
   - Pointer-based data structures for efficient memory management
   - Pass-by-reference semantics for large data structures

3. **Interfaces and Types**:
   - Custom type definitions for domain objects
   - Function type definitions (e.g., `HashFunc`)
   - Struct composition for complex data structures

4. **Error Handling**:
   - Idiomatic Go error handling patterns
   - Error propagation through the call stack
   - Custom error types with meaningful messages

## Testing

The implementation includes comprehensive tests that verify:
- Hash function idempotency
- Ring creation and management
- Database operations
- Concurrent record additions
- Data rebalancing logic
- Edge cases and error conditions

The test suite demonstrates Go's built-in testing framework usage and concurrent testing patterns.

## Learning Outcomes

This project provided hands-on experience with:
1. Distributed systems concepts and algorithms
2. Go's concurrency model and primitives
3. Thread-safe data structure implementation
4. Test-driven development in Go
5. Error handling patterns and best practices

## Future Improvements

Potential enhancements could include:
- Network layer implementation for actual distributed operation
- Persistent storage integration
- Additional rebalancing strategies
- Enhanced monitoring and metrics
- Fault tolerance mechanisms

## Usage

To use this implementation:

```go
// Create a new hash ring with replication factor of 3
ring := NewRing(3, nil)

// Create and add a database server
db := &database{serverid: 1}
ring.AddDB(db)

// Add records
record := NewRecord("John", "Doe", 10000.00, 5000.00, 100000.00, nil)
ring.addRecord(record)
```

## Testing

Run the tests using:

```bash
go test ./...
```
