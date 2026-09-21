# RPC Reference

The JSON-RPC server mocks endpoints matching standard Soroban integration patterns. All methods accept standard JSON-RPC 2.0 requests and return the same envelope.

## Methods

### `simulateTransaction`
Performs a dry run of the requested operations without modifying state.
- **Params:** `{"operations": [{"type": "set_storage", "contract_id": "C1", "key": "k2", "value": "v2"}]}`
- **Result:** `{"status": "success", "operations": [{"success": true, "error": ""}]}`

**Example:**
```json
{"jsonrpc":"2.0","id":3,"method":"simulateTransaction","params":{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"}]}}
```

### `sendTransaction`
Applies requested operations atomically to the engine state.
- **Params:** `{"operations": [{"type": "set_storage", "contract_id": "C1", "key": "k2", "value": "v2"}]}`
- **Result:** `{"status": "success", "operations": [{"success": true, "error": ""}]}`

**Example:**
```json
{"jsonrpc":"2.0","id":4,"method":"sendTransaction","params":{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"}]}}
```

### `getAccount`
Fetches a mock account's balance.
- **Params:** `{"address": "A1"}`
- **Result:** `{"balance": 100}`
- **Error:** `{"code": -32001, "message": "Account not found"}`

**Example:**
```json
{"jsonrpc":"2.0","id":1,"method":"getAccount","params":{"address":"A1"}}
```

### `getLedgerEntries`
Fetches storage values for a given contract.
- **Params:** `{"contract_id": "C1", "keys": ["k1", "k2"]}`
- **Result:** `{"k1": "v1", "k2": "v2"}`
- **Error:** `{"code": -32002, "message": "Contract not found"}`

**Example:**
```json
{"jsonrpc":"2.0","id":5,"method":"getLedgerEntries","params":{"contract_id":"C1","keys":["k1","k2"]}}
```
