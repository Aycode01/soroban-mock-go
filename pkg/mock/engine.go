package mock

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Aycode01/soroban-mock-go/pkg/config"
)

type Operation struct {
	Type       string `json:"type"` // "set_storage" | "delete_storage" | "transfer"
	ContractID string `json:"contract_id,omitempty"`
	Key        string `json:"key,omitempty"`
	Value      string `json:"value,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	Amount     int64  `json:"amount,omitempty"`
}

type Transaction struct {
	Operations []Operation `json:"operations"`
}

type OperationResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type TransactionResult struct {
	Status     string            `json:"status"`
	Operations []OperationResult `json:"operations"`
}

type Engine struct {
	mu        sync.RWMutex
	contracts map[string]map[string]string
	accounts  map[string]int64
}

func NewEngine(cfg *config.MockConfig) *Engine {
	e := &Engine{
		contracts: make(map[string]map[string]string),
		accounts:  make(map[string]int64),
	}
	for _, c := range cfg.Contracts {
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

func (e *Engine) GetContractStorage(contractID string) (map[string]string, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	storage, ok := e.contracts[contractID]
	return storage, ok
}

func (e *Engine) GetAccountBalance(address string) (int64, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	bal, ok := e.accounts[address]
	return bal, ok
}

func (e *Engine) cloneState() (map[string]map[string]string, map[string]int64) {
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
	return contractsCopy, accountsCopy
}

func applyOps(contracts map[string]map[string]string, accounts map[string]int64, ops []Operation) ([]OperationResult, bool) {
	results := make([]OperationResult, len(ops))
	allSuccess := true

	for i, op := range ops {
		switch op.Type {
		case "set_storage":
			if _, ok := contracts[op.ContractID]; !ok {
				results[i] = OperationResult{Success: false, Error: "contract does not exist"}
				allSuccess = false
				break
			}
			contracts[op.ContractID][op.Key] = op.Value
			results[i] = OperationResult{Success: true}
		case "delete_storage":
			if _, ok := contracts[op.ContractID]; !ok {
				results[i] = OperationResult{Success: false, Error: "contract does not exist"}
				allSuccess = false
				break
			}
			if _, ok := contracts[op.ContractID][op.Key]; !ok {
				results[i] = OperationResult{Success: false, Error: "key does not exist"}
				allSuccess = false
				break
			}
			delete(contracts[op.ContractID], op.Key)
			results[i] = OperationResult{Success: true}
		case "transfer":
			balFrom, okFrom := accounts[op.From]
			_, okTo := accounts[op.To]
			if !okFrom || !okTo {
				results[i] = OperationResult{Success: false, Error: "account does not exist"}
				allSuccess = false
				break
			}
			if op.Amount <= 0 {
				results[i] = OperationResult{Success: false, Error: "amount must be positive"}
				allSuccess = false
				break
			}
			if balFrom < op.Amount {
				results[i] = OperationResult{Success: false, Error: "insufficient balance"}
				allSuccess = false
				break
			}
			accounts[op.From] -= op.Amount
			accounts[op.To] += op.Amount
			results[i] = OperationResult{Success: true}
		default:
			results[i] = OperationResult{Success: false, Error: "unsupported operation type"}
			allSuccess = false
		}
	}
	return results, allSuccess
}

func (e *Engine) Simulate(txJSON string) (*TransactionResult, error) {
	var tx Transaction
	if err := json.Unmarshal([]byte(txJSON), &tx); err != nil {
		return nil, fmt.Errorf("invalid transaction JSON: %w", err)
	}

	e.mu.RLock()
	contracts, accounts := e.cloneState()
	e.mu.RUnlock()

	results, allSuccess := applyOps(contracts, accounts, tx.Operations)

	status := "success"
	if !allSuccess {
		status = "error"
	}

	return &TransactionResult{
		Status:     status,
		Operations: results,
	}, nil
}

func (e *Engine) Apply(txJSON string) (*TransactionResult, error) {
	var tx Transaction
	if err := json.Unmarshal([]byte(txJSON), &tx); err != nil {
		return nil, fmt.Errorf("invalid transaction JSON: %w", err)
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	contracts, accounts := e.cloneState()
	results, allSuccess := applyOps(contracts, accounts, tx.Operations)

	status := "success"
	if !allSuccess {
		status = "error"
	} else {
		e.contracts = contracts
		e.accounts = accounts
	}

	return &TransactionResult{
		Status:     status,
		Operations: results,
	}, nil
}

func (e *Engine) DumpState() map[string]any {
	e.mu.RLock()
	defer e.mu.RUnlock()
	c, a := e.cloneState()
	return map[string]any{
		"contracts": c,
		"accounts":  a,
	}
}
