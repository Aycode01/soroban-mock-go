package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aycode01/soroban-mock-go/pkg/config"
	"github.com/Aycode01/soroban-mock-go/pkg/mock"
)

func TestRPCServer(t *testing.T) {
	cfg := &config.MockConfig{
		Contracts: []config.ContractConfig{
			{ID: "C1", Storage: map[string]string{"k1": "v1"}},
		},
		Accounts: []config.AccountConfig{
			{Address: "A1", Balance: 100},
		},
	}
	e := mock.NewEngine(cfg)
	server := NewServer(e)

	tests := []struct {
		name       string
		method     string
		reqBody    string
		wantStatus int
		wantResult bool
		wantError  bool
	}{
		{
			name:       "getAccount success",
			method:     http.MethodPost,
			reqBody:    `{"jsonrpc":"2.0","id":1,"method":"getAccount","params":{"address":"A1"}}`,
			wantStatus: http.StatusOK,
			wantResult: true,
			wantError:  false,
		},
		{
			name:       "getAccount not found",
			method:     http.MethodPost,
			reqBody:    `{"jsonrpc":"2.0","id":2,"method":"getAccount","params":{"address":"A2"}}`,
			wantStatus: http.StatusOK,
			wantResult: false,
			wantError:  true,
		},
		{
			name:       "simulateTransaction success",
			method:     http.MethodPost,
			reqBody:    `{"jsonrpc":"2.0","id":3,"method":"simulateTransaction","params":{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"}]}}`,
			wantStatus: http.StatusOK,
			wantResult: true,
			wantError:  false,
		},
		{
			name:       "sendTransaction success",
			method:     http.MethodPost,
			reqBody:    `{"jsonrpc":"2.0","id":4,"method":"sendTransaction","params":{"operations":[{"type":"set_storage","contract_id":"C1","key":"k2","value":"v2"}]}}`,
			wantStatus: http.StatusOK,
			wantResult: true,
			wantError:  false,
		},
		{
			name:       "getLedgerEntries success",
			method:     http.MethodPost,
			reqBody:    `{"jsonrpc":"2.0","id":5,"method":"getLedgerEntries","params":{"contract_id":"C1","keys":["k1","k2"]}}`,
			wantStatus: http.StatusOK,
			wantResult: true,
			wantError:  false,
		},
		{
			name:       "invalid method",
			method:     http.MethodGet,
			reqBody:    "",
			wantStatus: http.StatusMethodNotAllowed,
			wantResult: false,
			wantError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, "/", bytes.NewBufferString(tt.reqBody))
			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("expected status %d got %d", tt.wantStatus, rr.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var resp Response
				_ = json.Unmarshal(rr.Body.Bytes(), &resp)
				if tt.wantResult && resp.Result == nil {
					t.Errorf("expected result, got nil")
				}
				if tt.wantError && resp.Error == nil {
					t.Errorf("expected error, got nil")
				}
			}
		})
	}
}
