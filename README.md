# FlashKV ⚡

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](https://github.com/Hari08-dev/flashkv/pulls)
[![Architecture](https://img.shields.io/badge/Architecture-In--Memory-orange?style=for-the-badge)](https://github.com/Hari08-dev/flashkv)

**FlashKV** is a high-performance, lightweight, concurrent in-memory Key-Value store written in Go. Designed with low latency and clean modular architecture in mind, FlashKV features thread-safe data structures, non-blocking TTL key expiration routines, a real-time Pub/Sub engine, JSON snapshot persistence, and an interactive CLI client.

---

## 🌟 Key Features

- **🚀 Concurrent & Thread-Safe**: Built from the ground up using Go's `sync.RWMutex` primitives to ensure lock safety under high concurrency.
- **📦 Rich Data Types**: Full support for **Strings**, **Lists**, **Hashes**, and **Sets**.
- **⏱️ Automated TTL & Key Expiration**: Supports key-level TTLs with a hybrid passive eviction model (on access) and active background ticker sweep goroutine.
- **📢 Real-Time Pub/Sub**: Light-weight TCP event broadcasting channel system for message streaming across connected clients.
- **💾 Snapshot Persistence**: On-demand JSON snapshot export (`SAVE`) and import (`LOAD`) to disk.
- **💻 Interactive CLI Client**: Dedicated command-line interface with ANSI terminal formatting and instant command execution.
- **🔌 Zero External Dependencies**: Powered entirely by the Go Standard Library (with `google/uuid` for client connection session tracking).

---

## 🏗️ Architecture Overview

FlashKV runs as an asynchronous TCP server listening by default on port `:6379`. The internal architecture cleanly decouples connection management, command parsing, data storage, and background processing routines:

```text
               +-----------------------------------+
               |  TCP Clients / CLI Interactive    |
               +-----------------------------------+
                                 |
                                 v
               +-----------------------------------+
               |    Client Session Manager         |
               | (UUID tracking & TCP connection)  |
               +-----------------------------------+
                                 |
                                 v
               +-----------------------------------+
               |     Command Lexer & Parser        |
               +-----------------------------------+
                 /               |               \
                v                v                v
      +------------------+  +---------+  +------------------+
      | In-Memory Store  |  | Pub/Sub |  | Persistence Engine|
      | (Strings/Lists/  |  | Engine  |  | (JSON Snapshots) |
      |  Hashes/Sets)    |  +---------+  +------------------+
      +------------------+
               ^
               | (Background Sweep)
      +------------------+
      |  TTL Cleaner     |
      |    Goroutine     |
      +------------------+
```

---

## 🚀 Quick Start

### Prerequisites

- [Go 1.20+](https://golang.org/doc/install) installed on your machine.

### Installation

Clone the repository:

```bash
git clone https://github.com/Hari08-dev/flashkv.git
cd flashkv
```

### Running the Server

Start the FlashKV server on default port `6379`:

```bash
go run cmd/server/main.go
```

_Output:_

```text
2026/08/21 12:00:00 FlashKV server starting on port 6379...
2026/08/21 12:00:00 Listening on :6379
```

### Running the Interactive CLI Client

In a separate terminal, launch the FlashKV CLI:

```bash
go run cmd/client/main.go
```

_Output:_

```text
Connected to FlashKV 🚀
Type EXIT to quit.

flashkv> PING
PONG
flashkv> SET title FlashKV
OK
flashkv> GET title
FlashKV
```

---

## 📖 Complete Command Reference

### 🔹 Basic & String Operations

| Command  | Usage               | Description                                                        | Example          |
| :------- | :------------------ | :----------------------------------------------------------------- | :--------------- |
| `PING`   | `PING`              | Health check endpoint. Returns `PONG`.                             | `PING`           |
| `SET`    | `SET <key> <value>` | Sets a string value for a key.                                     | `SET name Alice` |
| `GET`    | `GET <key>`         | Retrieves the string value of a key. Returns `(nil)` if not found. | `GET name`       |
| `DEL`    | `DEL <key>`         | Deletes a key and its associated value.                            | `DEL name`       |
| `EXISTS` | `EXISTS <key>`      | Checks if a key exists (`1`) or not (`0`).                         | `EXISTS name`    |
| `KEYS`   | `KEYS`              | Returns all active keys in the store.                              | `KEYS`           |

---

### 🔹 Key Expiration & TTL

| Command  | Usage                    | Description                                                                       | Example                |
| :------- | :----------------------- | :-------------------------------------------------------------------------------- | :--------------------- |
| `EXPIRE` | `EXPIRE <key> <seconds>` | Sets a time-to-live (in seconds) on a key.                                        | `EXPIRE session:1 300` |
| `TTL`    | `TTL <key>`              | Gets remaining TTL in seconds. Returns `-1` if no TTL, `-2` if key doesn't exist. | `TTL session:1`        |

---

### 🔹 List Operations

| Command  | Usage                         | Description                                                | Example                    |
| :------- | :---------------------------- | :--------------------------------------------------------- | :------------------------- |
| `LPUSH`  | `LPUSH <key> <value>`         | Prepends a value to the head of a list.                    | `LPUSH tasks "write docs"` |
| `RPUSH`  | `RPUSH <key> <value>`         | Appends a value to the tail of a list.                     | `RPUSH tasks "run tests"`  |
| `LPOP`   | `LPOP <key>`                  | Removes and returns the first element of a list.           | `LPOP tasks`               |
| `RPOP`   | `RPOP <key>`                  | Removes and returns the last element of a list.            | `RPOP tasks`               |
| `LLEN`   | `LLEN <key>`                  | Returns the length of the list.                            | `LLEN tasks`               |
| `LRANGE` | `LRANGE <key> <start> <stop>` | Returns a subset of elements within range `[start, stop]`. | `LRANGE tasks 0 -1`        |

---

### 🔹 Hash Operations

| Command   | Usage                        | Description                                        | Example                  |
| :-------- | :--------------------------- | :------------------------------------------------- | :----------------------- |
| `HSET`    | `HSET <key> <field> <value>` | Sets a field-value pair in a hash stored at key.   | `HSET user:100 name Bob` |
| `HGET`    | `HGET <key> <field>`         | Retrieves the value associated with field in hash. | `HGET user:100 name`     |
| `HDEL`    | `HDEL <key> <field>`         | Removes a field from a hash.                       | `HDEL user:100 name`     |
| `HEXISTS` | `HEXISTS <key> <field>`      | Checks if field exists in a hash (`1` or `0`).     | `HEXISTS user:100 name`  |
| `HGETALL` | `HGETALL <key>`              | Returns all fields and values of a hash.           | `HGETALL user:100`       |
| `HLEN`    | `HLEN <key>`                 | Returns the number of fields in a hash.            | `HLEN user:100`          |

---

### 🔹 Set Operations

| Command     | Usage                      | Description                                       | Example                 |
| :---------- | :------------------------- | :------------------------------------------------ | :---------------------- |
| `SADD`      | `SADD <key> <member>`      | Adds a member to a set.                           | `SADD tags golang`      |
| `SREM`      | `SREM <key> <member>`      | Removes a member from a set.                      | `SREM tags golang`      |
| `SISMEMBER` | `SISMEMBER <key> <member>` | Checks if member exists in set (`1` or `0`).      | `SISMEMBER tags golang` |
| `SMEMBERS`  | `SMEMBERS <key>`           | Returns all members of a set.                     | `SMEMBERS tags`         |
| `SCARD`     | `SCARD <key>`              | Returns the set cardinality (number of elements). | `SCARD tags`            |

---

### 🔹 Pub/Sub Messaging

| Command     | Usage                         | Description                                             | Example                      |
| :---------- | :---------------------------- | :------------------------------------------------------ | :--------------------------- |
| `SUBSCRIBE` | `SUBSCRIBE <channel>`         | Subscribes the current client connection to a channel.  | `SUBSCRIBE news`             |
| `PUBLISH`   | `PUBLISH <channel> <message>` | Posts a message to a channel. Returns subscriber count. | `PUBLISH news "Hello World"` |

---

### 🔹 Persistence

| Command | Usage  | Description                                                | Example |
| :------ | :----- | :--------------------------------------------------------- | :------ |
| `SAVE`  | `SAVE` | Persists all memory state to disk (`snapshots/dump.json`). | `SAVE`  |
| `LOAD`  | `LOAD` | Restores memory state from disk (`snapshots/dump.json`).   | `LOAD`  |

---

## 🛠️ Usage Examples

### 1. Hash & Expiration Workflow

```text
flashkv> HSET user:42 name "Hari"
OK
flashkv> HSET user:42 role "Developer"
OK
flashkv> HGETALL user:42
name: Hari
role: Developer
flashkv> EXPIRE user:42 60
1
flashkv> TTL user:42
58
```

### 2. Pub/Sub Demonstration

**Client 1 (Subscriber):**

```text
flashkv> SUBSCRIBE alerts
Subscribed to alerts
[alerts] High CPU usage detected!
```

**Client 2 (Publisher):**

```text
flashkv> PUBLISH alerts "High CPU usage detected!"
1
```

---

## 📂 Project Structure

```text
flashkv/
├── cmd/
│   ├── server/          # Main server entrypoint (TCP Server listener)
│   └── client/          # Interactive CLI client terminal application
├── internal/
│   ├── client/          # Client connection model and active session manager
│   ├── commands/        # Command handler implementations (SET, GET, HSET, etc.)
│   ├── datastore/       # Concurrent in-memory storage engine & core data types
│   ├── parser/          # Text protocol lexer, token parser, & command router
│   ├── persistence/     # JSON Snapshot Save/Load engine
│   ├── pubsub/          # In-memory pub/sub broker & connection subscriber registry
│   └── ttl/             # Background TTL cleanup ticker worker
├── snapshots/           # Directory storing JSON database dumps (dump.json)
├── go.mod               # Go module descriptor
└── README.md            # Project documentation
```

---

## 📌 Roadmap & Upcoming Features

The following features are actively under development or planned for upcoming releases:

- [ ] **Create/Get Semantics Refactor**: Refactor `GetList`, `GetHash`, and `GetSet` semantics in the datastore layer to cleanly separate creation-on-write from non-mutating lookups.
- [ ] **`UNSUBSCRIBE` Command**: Complete the Pub/Sub command set with explicit channel unsubscription support.
- [ ] **Auto-Save & Auto-Load**: Implement automatic snapshot recovery on startup and configurable periodic / shutdown auto-saving.
- [ ] **Graceful Shutdown**: OS signal handling (`SIGINT`, `SIGTERM`) for zero-data-loss server termination and clean connection tear-down.
- [ ] **LRU Eviction Policy**: Configurable max-memory bounds and Least Recently Used (LRU) key eviction under memory pressure.
- [ ] **Unit & Integration Test Suite**: Comprehensive testing coverage across concurrency handlers, TTL cleanup, Pub/Sub broadcasting, and command parsers.
- [ ] **Performance Benchmarks**: Benchmark suite (`go test -bench`) for measuring ops/sec, throughput, and memory allocations.
- [ ] **Docker & Docker Compose**: Official containerized setup (`Dockerfile` & `docker-compose.yml`) for seamless deployment.
- [ ] **Sorted Sets (`ZSET`)** _(Optional)_: Skip-list or balanced-tree backed sorted sets for score-based ranking queries (`ZADD`, `ZRANGE`, `ZSCORE`).

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check the [issues page](https://github.com/Hari08-dev/flashkv/issues) or submit a Pull Request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

<p align="center">
  Crafted with ❤️ by <a href="https://github.com/Hari08-dev">Hari08-dev</a>
</p>
