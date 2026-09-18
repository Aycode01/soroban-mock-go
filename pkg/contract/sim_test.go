package contract

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Aycode01/soroban-mock-go/pkg/config"
	"github.com/Aycode01/soroban-mock-go/pkg/mock"
)

func TestSimulateAndApply(t *testing.T) {
	cfg := &config.MockConfig{
		Contracts: []config.ContractConfig{
			{ID: "C1", Storage: map[string]string{"k1": "v1"}},
		},
	}
	e := mock.NewEngine(cfg)
	tx := `{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"}]}`

	// Simulate
	resStr, err := Simulate(e, tx)
	if err != nil {
		t.Fatalf("Simulate failed: %v", err)
	}

	var simRes mock.TransactionResult
	if err := json.Unmarshal([]byte(resStr), &simRes); err != nil {
		t.Fatalf("failed to unmarshal simulate result: %v", err)
	}
	if simRes.Status != "success" {
		t.Fatalf("expected success")
	}

	// Apply
	resStr, err = Apply(e, tx)
	if err != nil {
		t.Fatalf("Apply failed: %v", err)
	}
	var appRes mock.TransactionResult
	if err := json.Unmarshal([]byte(resStr), &appRes); err != nil {
		t.Fatalf("failed to unmarshal apply result: %v", err)
	}
	if appRes.Status != "success" {
		t.Fatalf("expected success")
	}

	if !strings.Contains(resStr, `"status": "success"`) {
		t.Fatalf("marshaled result missing status")
	}
}
