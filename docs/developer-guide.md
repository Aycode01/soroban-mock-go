# Developer Guide

## Local Setup

We use standard Go tooling, exactly as run in our CI workflow:

```bash
# Verify modules
go mod verify

# Linting (requires golangci-lint)
golangci-lint run

# Build the project
go build ./...

# Run the test suite with race detector
go test -race ./...
```

## Project Layout

- `pkg/config`: Handles loading, parsing, and validating YAML initial state configurations.
- `pkg/mock`: Houses the in-memory state engine, cloning logic, and transaction operations.
- `pkg/contract`: Wraps state operations and coordinates JSON serialization for simulate/apply flows.
- `pkg/rpc`: Implements the JSON-RPC 2.0 HTTP server bridging client requests to the state engine.
- `cmd/soroban-mock`: The entrypoint binary providing the CLI flags and execution pathways.

## Adding a New Operation Type

To add a new mock operation (e.g. conditional storage checks), follow these steps:

1. Locate the `applyOps` function in `pkg/mock/engine.go`.
2. Add a new `case` block inside the `switch op.Type` statement.
3. Write the required validations. If validation fails, set `results[i] = OperationResult{Success: false, Error: "your error message"}`, mark `allSuccess = false`, and call `break`.
4. If validation passes, apply the changes to the cloned state and set `results[i] = OperationResult{Success: true}`.
5. Add test coverage in `pkg/mock/engine_test.go` ensuring both the success case and the exact validation failure cases work as expected.
