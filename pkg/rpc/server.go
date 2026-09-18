package rpc

import (
	"encoding/json"
	"net/http"

	"github.com/Aycode01/soroban-mock-go/pkg/contract"
	"github.com/Aycode01/soroban-mock-go/pkg/mock"
)

// Request represents a JSON-RPC 2.0 request.
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

// Response represents a JSON-RPC 2.0 response.
type Response struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

// Error represents a JSON-RPC 2.0 error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Server implements a JSON-RPC 2.0 HTTP server.
type Server struct {
	engine *mock.Engine
}

// NewServer creates a new RPC server.
func NewServer(e *mock.Engine) *Server {
	return &Server{engine: e}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.writeError(w, nil, -32700, "Parse error")
		return
	}

	if req.JSONRPC != "2.0" {
		s.writeError(w, req.ID, -32600, "Invalid Request")
		return
	}

	var result any
	var err *Error

	switch req.Method {
	case "simulateTransaction":
		result, err = s.handleTransaction(req.Params, contract.Simulate)
	case "sendTransaction":
		result, err = s.handleTransaction(req.Params, contract.Apply)
	case "getAccount":
		result, err = s.handleGetAccount(req.Params)
	case "getLedgerEntries":
		result, err = s.handleGetLedgerEntries(req.Params)
	default:
		err = &Error{Code: -32601, Message: "Method not found"}
	}

	if err != nil {
		s.writeError(w, req.ID, err.Code, err.Message)
		return
	}

	resp := Response{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result:  result,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) handleTransaction(params json.RawMessage, fn func(*mock.Engine, string) (string, error)) (any, *Error) {
	var tx string
	if err := json.Unmarshal(params, &tx); err == nil {
		// if params is string (JSON)
	} else {
		tx = string(params)
	}

	resStr, err := fn(s.engine, tx)
	if err != nil {
		return nil, &Error{Code: -32000, Message: err.Error()}
	}
	var res map[string]any
	_ = json.Unmarshal([]byte(resStr), &res)
	return res, nil
}

func (s *Server) handleGetAccount(params json.RawMessage) (any, *Error) {
	var p struct {
		Address string `json:"address"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return nil, &Error{Code: -32602, Message: "Invalid params"}
	}
	bal, ok := s.engine.GetAccountBalance(p.Address)
	if !ok {
		return nil, &Error{Code: -32001, Message: "Account not found"}
	}
	return map[string]any{"balance": bal}, nil
}

func (s *Server) handleGetLedgerEntries(params json.RawMessage) (any, *Error) {
	var p struct {
		ContractID string   `json:"contract_id"`
		Keys       []string `json:"keys"`
	}
	if err := json.Unmarshal(params, &p); err != nil {
		return nil, &Error{Code: -32602, Message: "Invalid params"}
	}
	storage, ok := s.engine.GetContractStorage(p.ContractID)
	if !ok {
		return nil, &Error{Code: -32002, Message: "Contract not found"}
	}
	entries := make(map[string]string)
	for _, k := range p.Keys {
		if v, ok := storage[k]; ok {
			entries[k] = v
		}
	}
	return entries, nil
}

func (s *Server) writeError(w http.ResponseWriter, id json.RawMessage, code int, msg string) {
	resp := Response{
		JSONRPC: "2.0",
		ID:      id,
		Error:   &Error{Code: code, Message: msg},
	}
	_ = json.NewEncoder(w).Encode(resp)
}
