# APIBench

> Understanding **why REST, GraphQL, and gRPC exist** by implementing the same backend three different ways.

## Overview

Every backend engineer eventually encounters REST, GraphQL, and gRPC.

Most articles explain **what** they are.

Few explain **why all three continue to exist.**

APIBench is a small experimental project built to answer that question.

Instead of comparing different applications, the same message queue is implemented using **REST**, **GraphQL**, and **gRPC**, while keeping the underlying business logic identical. This allows the communication layer to be isolated so the trade-offs between each API style become much easier to understand.

The goal isn't to prove that one technology is "better."

The goal is to understand **where each one fits in modern software systems.**

---

## Motivation

Questions that inspired this project:

- Why does Stripe expose REST APIs?
- Why does Shopify heavily use GraphQL?
- Why does Google build internal systems with gRPC?
- If gRPC is faster, why isn't every API built with it?
- Why do all three technologies still exist?

Rather than reading documentation or benchmark charts, I wanted to answer these questions by building the same system three different ways.

---

## Project Architecture

```
                     Client
                        │
        ┌───────────────┼───────────────┐
        │               │               │
        ▼               ▼               ▼
      REST          GraphQL           gRPC
        │               │               │
        └───────────────┼───────────────┘
                        │
              Message Queue Service
                        │
                In-Memory Repository
```

All three implementations expose the same operations:

- Publish a message
- Poll the next message
- Acknowledge a message
- Reject a message

The only thing that changes is **how clients communicate with the service.**

---

## Experiments

This repository includes several small experiments to understand the practical differences between the three API styles.

### Field Selection (GraphQL)

Demonstrates how GraphQL allows clients to request only the fields they need instead of receiving an entire resource.

Example:

REST

```json
{
  "id": "...",
  "payload": "...",
  "status": "...",
  "createdAt": "..."
}
```

GraphQL

```graphql
query {
  message(id: "msg-1") {
    status
  }
}
```

---

### Serialization Comparison

Compares JSON serialization against Protocol Buffers to understand why gRPC produces smaller payloads.

Example result:

| Format | Payload Size |
|---------|-------------:|
| REST JSON | 180 bytes |
| GraphQL JSON | 201 bytes |
| Protocol Buffers | 139 bytes |

---

### Latency Benchmark

Each implementation is benchmarked under the same workload.

Example:

| API | Avg Latency |
|-----|------------:|
| REST | ~319 μs |
| GraphQL | ~325 μs |
| gRPC | ~92 μs |

The purpose isn't to declare a winner, but to understand the design goals behind each technology.

---

## Technologies

- Go
- Gin
- GraphQL
- gRPC
- Protocol Buffers

---

## Repository Structure

```
api/
    rest/
    graphql/
    grpc/

benchmark/
    benchmarking utilities

cmd/
    server entrypoints

experiments/
    field selection
    payload comparison
    serialization
    latency benchmarks

internal/
    queue
    repository
    service
    models
```

---

## What I Learned

Building the same system three different ways completely changed how I think about APIs.

REST isn't "old."

GraphQL isn't "REST but better."

gRPC isn't "the fastest API."

Each one optimizes for a different communication problem.

- REST prioritizes simplicity and developer experience.
- GraphQL prioritizes flexible data fetching.
- gRPC prioritizes efficient service-to-service communication.

Modern production systems often use all three.

---

## Future Improvements

- Streaming benchmarks using gRPC
- Authentication and authorization
- Persistent storage
- Docker Compose setup
- Distributed message queue
- Kubernetes deployment
- Additional benchmarking scenarios

---

## Blog

This project accompanies my article:

> **Understanding REST, GraphQL, and gRPC: Why Modern Software Still Needs All Three**

The article explains the reasoning behind each implementation and how these technologies are used together in modern production systems.

---

## License

MIT