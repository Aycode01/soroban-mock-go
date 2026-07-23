package config

import (
    "io/ioutil"
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

// LoadConfig reads a YAML file at the given path and unmarshals it into a MockConfig.
func LoadConfig(path string) (*MockConfig, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var cfg MockConfig
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
