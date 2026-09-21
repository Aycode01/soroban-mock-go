# Config Reference

The state engine is initialized using a YAML configuration file describing mock contracts and account balances.

## YAML Schema

```yaml
contracts:
  - id: string
    storage:
      string: string
accounts:
  - address: string
    balance: int64
```

## Validation Rules

The `Validate()` function enforces the following constraints during load:

- **Contracts:**
  - `id` cannot be empty.
  - `id` must be unique across all contracts (no duplicates).
  - Storage keys cannot be empty.
- **Accounts:**
  - `address` cannot be empty.
  - `address` must be unique across all accounts (no duplicates).
  - `balance` cannot be negative.

## Example Configuration

```yaml
# Sample mock configuration for soroban-mock-go

contracts:
  - id: "C1234567890"
    storage:
      "key1": "value1"
      "key2": "value2"
  - id: "C0987654321"
    storage:
      "counter": "42"

accounts:
  - address: "GABCDEF123456"
    balance: 1000000
  - address: "GHIJKL789012"
    balance: 500000
```
