# Godis

A Redis server implementation in Go.

## Description

Godis is a lightweight Redis-like server written in Go. It implements the Redis Serialization Protocol (RESP) and provides a TCP server that can process Redis commands.

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/godis.git

# Navigate to the project directory
cd godis

# Build the project
go build
```

## Usage

```bash
# Run the server with default settings (127.0.0.1:8080)
./godis

# Run the server with custom host and port
./godis -h 0.0.0.0 -p 6379
```

You can connect to the server using any Redis client, or use a simple TCP client like netcat:

```bash
nc 127.0.0.1 8080
```

## Project Status

This project is in early development. The basic server structure is in place, but many features are not yet implemented.

## Todo List

- [ ] Fix bugs in the RESP decoder:
  - [x] Fix DecodeCommands function which currently always returns nil
  - [ ] Fix panics
- [x] Implement async server based on eventloop
- [ ] Add support for windows based on IOCP
- [x] Implement RESP encoder for proper response formatting
- [x] Implement command evaluation logic
- [x] Add data storage functionality
- [x] Implement basic Redis commands:
  - [x] GET
  - [x] SET
  - [x] DEL
  - [x] TTL
  - [x] INCR
  - [x] DECR
  - [x] EXPIRE
- [ ] Improve Naive Implementation of Delete & Expiry
- [ ] Add support for Eviction Policies
    - [ ] LRU
    - [ ] LFU
    - [ ] Approximate LRU
    - [ ] Fixed
- [ ] Add config file and refactor argument flags
- [ ] Add support for data types/structures:
  - [x] Strings
  - [ ] Lists
  - [ ] Sorted Sets
  - [ ] Geo Hash
  - [ ] HyperLogLog
- [ ] Improve Naive Implementation of Delete & Expiry
- [ ] Improve error handling
- [ ] Add comprehensive tests
- [ ] Add benchmarking
- [ ] Add Redis CLI using Charm / Bubble Tea
- [ ] Add persistence options
- [ ] Add transactions
- [ ] Improve documentation


## Contributing

Contributions are welcome! Please feel free to submit issues.