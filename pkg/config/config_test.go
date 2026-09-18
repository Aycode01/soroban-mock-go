package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMockConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  MockConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: MockConfig{
				Contracts: []ContractConfig{
					{ID: "C1", Storage: map[string]string{"k1": "v1"}},
				},
				Accounts: []AccountConfig{
					{Address: "A1", Balance: 100},
				},
			},
			wantErr: false,
		},
		{
			name: "empty contract id",
			config: MockConfig{
				Contracts: []ContractConfig{
					{ID: "", Storage: map[string]string{"k1": "v1"}},
				},
			},
			wantErr: true,
		},
		{
			name: "duplicate contract id",
			config: MockConfig{
				Contracts: []ContractConfig{
					{ID: "C1", Storage: map[string]string{"k1": "v1"}},
					{ID: "C1", Storage: map[string]string{"k2": "v2"}},
				},
			},
			wantErr: true,
		},
		{
			name: "empty storage key",
			config: MockConfig{
				Contracts: []ContractConfig{
					{ID: "C1", Storage: map[string]string{"": "v1"}},
				},
			},
			wantErr: true,
		},
		{
			name: "empty account address",
			config: MockConfig{
				Accounts: []AccountConfig{
					{Address: "", Balance: 100},
				},
			},
			wantErr: true,
		},
		{
			name: "duplicate account address",
			config: MockConfig{
				Accounts: []AccountConfig{
					{Address: "A1", Balance: 100},
					{Address: "A1", Balance: 200},
				},
			},
			wantErr: true,
		},
		{
			name: "negative balance",
			config: MockConfig{
				Accounts: []AccountConfig{
					{Address: "A1", Balance: -100},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("MockConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	tempDir := t.TempDir()

	validConfig := filepath.Join(tempDir, "valid.yaml")
	os.WriteFile(validConfig, []byte(`
contracts:
  - id: C1
    storage:
      k1: v1
accounts:
  - address: A1
    balance: 100
`), 0644)

	invalidConfig := filepath.Join(tempDir, "invalid.yaml")
	os.WriteFile(invalidConfig, []byte(`
contracts:
  - id: C1
accounts:
  - address: A1
    balance: -100
`), 0644)

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "valid config",
			path:    validConfig,
			wantErr: false,
		},
		{
			name:    "invalid config",
			path:    invalidConfig,
			wantErr: true,
		},
		{
			name:    "missing file",
			path:    filepath.Join(tempDir, "missing.yaml"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadConfig(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
