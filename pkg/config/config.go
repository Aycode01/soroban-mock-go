package config

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v3"
)

// ContractConfig describes a single contract's mock state.
type ContractConfig struct {
    ID      string            `yaml:"id"`
    Storage map[string]string `yaml:"storage"`
}

// AccountConfig describes a mock account's balance.
type AccountConfig struct {
    Address string `yaml:"address"`
    Balance int64  `yaml:"balance"`
}

// MockConfig is the top‑level configuration structure parsed from YAML.
type MockConfig struct {
    Contracts []ContractConfig `yaml:"contracts"`
    Accounts  []AccountConfig  `yaml:"accounts"`
}

// Validate checks the configuration for invalid states.
func (cfg *MockConfig) Validate() error {
    seenContracts := make(map[string]bool)
    for i, contract := range cfg.Contracts {
        if contract.ID == "" {
            return fmt.Errorf("contract at index %d has empty ID", i)
        }
        if seenContracts[contract.ID] {
            return fmt.Errorf("duplicate contract ID: %s", contract.ID)
        }
        seenContracts[contract.ID] = true

        for k := range contract.Storage {
            if k == "" {
                return fmt.Errorf("contract %s has empty storage key", contract.ID)
            }
        }
    }

    seenAccounts := make(map[string]bool)
    for i, account := range cfg.Accounts {
        if account.Address == "" {
            return fmt.Errorf("account at index %d has empty address", i)
        }
        if seenAccounts[account.Address] {
            return fmt.Errorf("duplicate account address: %s", account.Address)
        }
        seenAccounts[account.Address] = true

        if account.Balance < 0 {
            return fmt.Errorf("account %s has negative balance: %d", account.Address, account.Balance)
        }
    }

    return nil
}

// LoadConfig reads a YAML file at the given path and unmarshals it into a MockConfig.
func LoadConfig(path string) (*MockConfig, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    var cfg MockConfig
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %w", err)
    }
    if err := cfg.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &cfg, nil
}
