# soroban-mock-go

A lightweight local development tool to mock Soroban contract state and RPC calls.

**Status:** Mocks storage/balance operations and JSON-RPC `simulateTransaction`/`sendTransaction`/`getAccount`/`getLedgerEntries` calls using an in-memory state engine. This does *not* execute full Soroban WASM bytecode.

## Installation

```bash
# Option 1: Install via go install
go install github.com/Aycode01/soroban-mock-go/cmd/soroban-mock@latest

# Option 2: Clone and build
git clone https://github.com/Aycode01/soroban-mock-go.git
cd soroban-mock-go
go build -o soroban-mock ./cmd/soroban-mock
```

## Usage

Create a configuration file (e.g., `mock_config.yaml`) defining the initial state:

```yaml
contracts:
  - id: "CA123"
    storage:
      "COUNTER": "1"
accounts:
  - address: "GABC"
    balance: 10000
```

### CLI Mode

You can apply a transaction directly via the CLI to see the resulting state:

```bash
soroban-mock --config mock_config.yaml --apply '{"operations":[{"type":"transfer","from":"GABC","to":"GDEF","amount":50}]}'
```

Output:
```json
{
  "status": "success",
  "operations": [
    {
      "success": true
    }
  ]
}

New State:
{
  "accounts": {
    "GABC": 9950,
    "GDEF": 50
  },
  "contracts": {
    "CA123": {
      "COUNTER": "1"
    }
  }
}
```

### RPC Server Mode

Start the JSON-RPC server on port 8080:

```bash
soroban-mock --config mock_config.yaml --serve --addr ":8080"
```

Then you can send standard JSON-RPC 2.0 requests:

```bash
curl -X POST http://localhost:8080 -d '{"jsonrpc":"2.0","id":1,"method":"getAccount","params":{"address":"GABC"}}'
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to run tests, lint, and submit PRs.

## Maintainer

<!-- TODO: add your contact -->
