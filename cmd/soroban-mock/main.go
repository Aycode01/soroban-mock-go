package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "os"

    "github.com/Aycode01/soroban-mock-go/pkg/config"
    "github.com/Aycode01/soroban-mock-go/pkg/contract"
    "github.com/Aycode01/soroban-mock-go/pkg/mock"
)

func main() {
    // Define command‑line flags.
    configPath := flag.String("config", "", "Path to mock configuration YAML file (required)")
    simulateTx := flag.String("simulate", "", "Optional transaction JSON to simulate")
    flag.Parse()

    if *configPath == "" {
        fmt.Fprintln(os.Stderr, "error: --config flag is required")
        flag.Usage()
        os.Exit(1)
    }

    // Load configuration.
    cfg, err := config.LoadConfig(*configPath)
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    // Initialize mock engine.
    engine := mock.NewEngine(cfg)

    var output interface{}
    if *simulateTx != "" {
        // Perform transaction simulation.
        result, err := contract.Simulate(engine, *simulateTx)
        if err != nil {
            log.Fatalf("simulation error: %v", err)
        }
        // The Simulate helper returns a JSON string; unmarshal for pretty printing.
        if err := json.Unmarshal([]byte(result), &output); err != nil {
            // Fallback to raw string if unmarshalling fails.
            output = result
        }
    } else {
        // No simulation requested – dump the full mock state.
        output = engine.DumpState()
    }

    // Marshal to indented JSON and write to stdout.
    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", "  ")
    if err := enc.Encode(output); err != nil {
        log.Fatalf("failed to encode output: %v", err)
    }
}
