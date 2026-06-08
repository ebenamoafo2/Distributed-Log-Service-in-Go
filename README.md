This project walks through the full lifecycle of a real-world distributed system: from a bare-bones storage engine and Protocol Buffer-defined data structures, to a gRPC-powered networked service with TLS security, all the way to a fully distributed, replicated cluster deployed on Kubernetes.
What's inside:

Custom commit log with segment and index management
gRPC client/server with mutual TLS authentication
Service discovery with automatic cluster membership
Distributed replication powered by the Raft consensus algorithm
Topology-aware load balancing
Observability via structured logs, metrics, and traces
CLI for service configuration and management
Cloud deployment with Kubernetes and a custom Operator

Built to explore the core ideas behind systems like Kafka, etcd, and CockroachDB — from first principles, in Go.
