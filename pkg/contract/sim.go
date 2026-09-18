package contract

import (
    "encoding/json"
    "fmt"
    "github.com/Aycode01/soroban-mock-go/pkg/mock"
)

// Simulate is a thin wrapper around the mock Engine's SimulateTransaction.
// It receives a transaction JSON string and returns a JSON‑encoded result.
func Simulate(e *mock.Engine, txJSON string) (string, error) {
    result, err := e.SimulateTransaction(txJSON)
    if err != nil {
        return "", err
    }
    b, err := json.MarshalIndent(result, "", "  ")
    if err != nil {
        return "", fmt.Errorf("failed to marshal simulation result: %w", err)
    }
    return string(b), nil
}
