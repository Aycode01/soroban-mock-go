package mock

import (
    "encoding/json"
    "fmt"
    "sync"

    "github.com/Aycode01/soroban-mock-go/pkg/config"
)

// Engine is the core mock engine that holds contract and account state.
type Engine struct {
    mu        sync.RWMutex
    contracts map[string]map[string]string // contractID -> storage key/value
    accounts  map[string]int64            // address -> balance
}

// NewEngine creates a new Engine instance by loading the provided config.
func NewEngine(cfg *config.MockConfig) *Engine {
    e := &Engine{
        contracts: make(map[string]map[string]string),
        accounts:  make(map[string]int64),
    }
    for _, c := range cfg.Contracts {
        // Ensure storage map is non-nil.
        storage := make(map[string]string)
        for k, v := range c.Storage {
            storage[k] = v
        }
        e.contracts[c.ID] = storage
    }
    for _, a := range cfg.Accounts {
        e.accounts[a.Address] = a.Balance
    }
    return e
}

// GetContractStorage returns the storage map for a given contract ID.
func (e *Engine) GetContractStorage(contractID string) (map[string]string, bool) {
    e.mu.RLock()
    defer e.mu.RUnlock()
    storage, ok := e.contracts[contractID]
    return storage, ok
}

// GetAccountBalance returns the balance for a given account address.
func (e *Engine) GetAccountBalance(address string) (int64, bool) {
    e.mu.RLock()
    defer e.mu.RUnlock()
    bal, ok := e.accounts[address]
    return bal, ok
}

// SimulateTransaction is a placeholder that demonstrates how a transaction could be simulated.
// In a real implementation this would interpret the transaction, modify state, and return a result.
func (e *Engine) SimulateTransaction(txJSON string) (map[string]any, error) {
    // For now, just echo back the input with a success flag.
    var tx map[string]any
    if err := json.Unmarshal([]byte(txJSON), &tx); err != nil {
        return nil, fmt.Errorf("invalid transaction JSON: %w", err)
    }
    // TODO: implement real simulation logic.
    result := map[string]any{
        "status": "success",
        "tx":    tx,
    }
    return result, nil
}

// DumpState returns the full mock state as a JSON‑serializable structure.
func (e *Engine) DumpState() map[string]any {
    e.mu.RLock()
    defer e.mu.RUnlock()
    // Deep copy to avoid exposing internal maps.
    contractsCopy := make(map[string]map[string]string)
    for cid, storage := range e.contracts {
        sCopy := make(map[string]string)
        for k, v := range storage {
            sCopy[k] = v
        }
        contractsCopy[cid] = sCopy
    }
    accountsCopy := make(map[string]int64)
    for addr, bal := range e.accounts {
        accountsCopy[addr] = bal
    }
    return map[string]any{
        "contracts": contractsCopy,
        "accounts":  accountsCopy,
    }
}
