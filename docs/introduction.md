# Introduction

`soroban-mock-go` is a local development tool designed for testing Soroban integration code (such as wallet backends, indexers, or frontend RPC calls) without running a real Soroban node or hitting testnet. This removes the slow setup time and the need for funded testnet accounts.

This tool does NOT execute Soroban WASM contract bytecode. Instead, it mocks storage reads and writes, account balance operations at the state level, and the JSON-RPC surface a client would call.

To learn more about how the in-memory engine handles transactions and RPC endpoints, see [How it works](architecture.md).
