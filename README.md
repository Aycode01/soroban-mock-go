# soroban-mock-go

A lightweight developer tool for mocking Soroban RPC calls and contract state locally.

## Overview

`soroban-mock-go` provides a simple way to spin up a mock RPC environment for Soroban smart contract development. It reads a declarative YAML configuration describing contract IDs, storage key‑values, and mock account balances, then serves simulated RPC responses. This enables fast, deterministic testing of contract interactions without needing a full network or Horizon node.

## Features

- **YAML‑driven configuration** – define contracts, storage, and accounts in a clear, version‑controlled file.
- **CLI interface** – run simulations directly from the command line.
- **Extensible engine** – query contract state, account balances, and simulate transaction outcomes.
- **JSON output** – convenient for piping into other tools or test harnesses.

## Installation

```bash
# Clone the repo
git clone https://github.com/secondaccount/project/soroban-mock-go.git
cd soroban-mock-go

# Build the binary
go build -o soroban-mock ./cmd/soroban-mock
```

## Quickstart

1. Create a configuration file (see `examples/mock_config.yaml`).
2. Run the tool:

```bash
./soroban-mock --config examples/mock_config.yaml
```

The tool will load the configuration and output a JSON representation of the mock state and a sample transaction simulation.

## Usage

```bash
soroban-mock --config <path-to-config.yaml> [--simulate "<tx-json>"]
```

- `--config` – path to the mock configuration file (required).
- `--simulate` – optional JSON describing a transaction to simulate; if omitted the tool prints the loaded mock state.

## Example Output

```json
{
  "contracts": {
    "C123": {
      "storage": {
        "key1": "value1",
        "key2": "value2"
      }
    }
  },
  "accounts": {
    "GABC": 1000000,
    "GDEF": 500000
  }
}
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch.
3. Ensure code passes `go vet` and `golint`.
4. Open a Pull Request targeting the `main` branch.

All contributions are made under the auspices of **Drips Wave** – a community of Stellar developers.

## License

MIT – see the [LICENSE](LICENSE) file for details.
