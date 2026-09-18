package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Aycode01/soroban-mock-go/pkg/config"
	"github.com/Aycode01/soroban-mock-go/pkg/contract"
	"github.com/Aycode01/soroban-mock-go/pkg/mock"
	"github.com/Aycode01/soroban-mock-go/pkg/rpc"
)

func main() {
	configPath := flag.String("config", "", "Path to mock configuration YAML file (required)")
	simulateTx := flag.String("simulate", "", "Optional transaction JSON to simulate")
	applyTx := flag.String("apply", "", "Optional transaction JSON to apply")
	serve := flag.Bool("serve", false, "Start JSON-RPC server")
	addr := flag.String("addr", ":8080", "Address to listen on for RPC server")
	flag.Parse()

	if *configPath == "" {
		fmt.Fprintln(os.Stderr, "error: --config flag is required")
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	engine := mock.NewEngine(cfg)

	if *serve {
		log.Printf("Starting mock RPC server on %s", *addr)
		server := rpc.NewServer(engine)
		if err := http.ListenAndServe(*addr, server); err != nil {
			log.Fatalf("server error: %v", err)
		}
		return
	}

	var output interface{}
	if *simulateTx != "" {
		result, err := contract.Simulate(engine, *simulateTx)
		if err != nil {
			log.Fatalf("simulation error: %v", err)
		}
		if err := json.Unmarshal([]byte(result), &output); err != nil {
			output = result
		}
		printOutput(output)
		os.Exit(0)
	} else if *applyTx != "" {
		result, err := contract.Apply(engine, *applyTx)
		if err != nil {
			log.Fatalf("apply error: %v", err)
		}
		if err := json.Unmarshal([]byte(result), &output); err != nil {
			output = result
		}
		printOutput(output)
		
		fmt.Println("\nNew State:")
		printOutput(engine.DumpState())
	} else {
		output = engine.DumpState()
		printOutput(output)
	}
}

func printOutput(output interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(output); err != nil {
		log.Fatalf("failed to encode output: %v", err)
	}
}
