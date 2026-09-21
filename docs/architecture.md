# Architecture

## The Engine

The `Engine` struct holds the current state of mock contract storage and account balances in memory. Access to the state is protected by a `sync.RWMutex` to ensure thread-safety, enabling concurrent operations such as RPC requests.

## Modes of Operation

- **Simulate:** Invoked via the `--simulate` flag on the CLI or `simulateTransaction` over RPC. It takes a complete snapshot clone of the in-memory state, applies operations to the snapshot, and returns the result without mutating the live state.
- **Apply:** Invoked via the `--apply` flag or `sendTransaction` over RPC. It applies operations directly to the live engine.

## Atomicity

Transactions are atomic. During an apply step, the engine first clones the live state. It then iterates through all requested operations using the `applyOps` function. If any single operation fails validation (e.g. insufficient balance during a transfer, or targeting a missing contract), the entire transaction immediately halts and rolls back. None of the operations within the failing transaction are committed to the live state.

## RPC Server

The RPC server wraps the mock `Engine`. It exposes standard JSON-RPC 2.0 endpoints. Incoming `simulateTransaction`, `sendTransaction`, `getAccount`, and `getLedgerEntries` calls are dispatched to a single shared `Engine` instance which is mutex-guarded for the life of the server process.
