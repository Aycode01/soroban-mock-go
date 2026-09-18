package mock

import (
	"testing"

	"github.com/Aycode01/soroban-mock-go/pkg/config"
)

func TestEngineOperations(t *testing.T) {
	cfg := &config.MockConfig{
		Contracts: []config.ContractConfig{
			{ID: "C1", Storage: map[string]string{"k1": "v1"}},
		},
		Accounts: []config.AccountConfig{
			{Address: "A1", Balance: 100},
			{Address: "A2", Balance: 50},
		},
	}
	e := NewEngine(cfg)

	// test set_storage
	res, err := e.Apply(`{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"}]}`)
	if err != nil || res.Status != "success" {
		t.Fatalf("set_storage failed: %v", res)
	}

	// test set_storage missing contract
	res, err = e.Apply(`{"operations":[{"type":"set_storage","contract_id":"C2","key":"k2","value":"v2"}]}`)
	if err != nil || res.Status != "error" {
		t.Fatalf("set_storage missing contract should fail")
	}

	// test delete_storage
	res, err = e.Apply(`{"operations":[{"type":"delete_storage","contract_id":"C1","key":"k1"}]}`)
	if err != nil || res.Status != "success" {
		t.Fatalf("delete_storage failed: %v", res)
	}

	// test delete_storage missing key
	res, err = e.Apply(`{"operations":[{"type":"delete_storage","contract_id":"C1","key":"k3"}]}`)
	if err != nil || res.Status != "error" {
		t.Fatalf("delete_storage missing key should fail")
	}

	// test transfer
	res, err = e.Apply(`{"operations":[{"type":"transfer","from":"A1","to":"A2","amount":10}]}`)
	if err != nil || res.Status != "success" {
		t.Fatalf("transfer failed: %v", res)
	}
	if b, _ := e.GetAccountBalance("A1"); b != 90 {
		t.Fatalf("A1 balance wrong: %d", b)
	}
	if b, _ := e.GetAccountBalance("A2"); b != 60 {
		t.Fatalf("A2 balance wrong: %d", b)
	}

	// test transfer negative
	res, err = e.Apply(`{"operations":[{"type":"transfer","from":"A1","to":"A2","amount":-10}]}`)
	if err != nil || res.Status != "error" {
		t.Fatalf("transfer negative amount should fail")
	}
}

func TestEngineSimulate(t *testing.T) {
	cfg := &config.MockConfig{
		Contracts: []config.ContractConfig{
			{ID: "C1", Storage: map[string]string{"k1": "v1"}},
		},
	}
	e := NewEngine(cfg)

	// Simulate
	res, err := e.Simulate(`{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"}]}`)
	if err != nil || res.Status != "success" {
		t.Fatalf("simulate failed: %v", res)
	}

	// Check it didn't mutate
	if _, ok := e.contracts["C1"]["k2"]; ok {
		t.Fatalf("simulate mutated state")
	}
}

func TestEngineAtomicity(t *testing.T) {
	cfg := &config.MockConfig{
		Contracts: []config.ContractConfig{
			{ID: "C1", Storage: map[string]string{"k1": "v1"}},
		},
	}
	e := NewEngine(cfg)

	// one success, one fail
	res, err := e.Apply(`{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"},{"type":"delete_storage","contract_id":"C1","key":"k3"}]}`)
	if err != nil || res.Status != "error" {
		t.Fatalf("apply should fail")
	}

	// Check atomicity (k2 should not be set)
	if _, ok := e.contracts["C1"]["k2"]; ok {
		t.Fatalf("apply was not atomic")
	}
}
