# CLI Reference

The `soroban-mock` binary exposes a command-line interface for manipulating and querying the mock state.

## Flags

- `--config` (string): Path to the mock configuration YAML file (required).
- `--simulate` (string): Optional transaction JSON string to perform a dry-run simulation. Returns the result without mutating state.
- `--apply` (string): Optional transaction JSON string to apply directly to the state. Mutates state on success.
- `--serve` (boolean, default `false`): Start the JSON-RPC server instead of running a one-shot command.
- `--addr` (string, default `":8080"`): Address to listen on when the RPC server is running.

## Examples

### Apply a Transaction

```bash
soroban-mock --config mock_config.yaml --apply '{"operations":[{"type":"transfer","from":"GABCDEF123456","to":"GHIJKL789012","amount":50}]}'
```

**Output:**
```json
{
  "status": "success",
  "operations": [
    {
      "success": true,
      "error": ""
    }
  ]
}

New State:
{
  "accounts": {
    "GABCDEF123456": 999950,
    "GHIJKL789012": 500050
  },
  "contracts": {
    "C0987654321": {
      "counter": "42"
    },
    "C1234567890": {
      "key1": "value1",
      "key2": "value2"
    }
  }
}
```
