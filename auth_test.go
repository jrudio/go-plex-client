package plex

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTAuthFlow(t *testing.T) {
	// Mock Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v2/pins.json":
			// Verify POST and body has JWK
			if r.Method != http.MethodPost {
				t.Errorf("Expected POST, got %s", r.Method)
			}
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("Failed to decode body: %v", err)
			}
			if body["jwk"] == nil {
				t.Error("Expected JWK in body")
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(PinResponse{
				ID:   123,
				Code: "ABCD",
			})
		case "/api/v2/pins/123.json":
			// Verify deviceJWT param
			query := r.URL.Query()
			deviceJWT := query.Get("deviceJWT")
			if deviceJWT == "" {
				t.Error("Expected deviceJWT query param")
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(PinResponse{
				ID:        123,
				Code:      "ABCD",
				AuthToken: "test-auth-token",
			})
		case "/api/v2/auth/nonce":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"nonce": "test-nonce"})
		case "/api/v2/auth/token":
			var body map[string]string
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("Failed to decode token body: %v", err)
			}
			if body["jwt"] == "" {
				t.Error("Expected jwt in body")
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"auth_token": "refreshed-token"})
		default:
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
			http.Error(w, "not found", http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Override URL
	originalURL := plexURL
	plexURL = server.URL
	defer func() { plexURL = originalURL }()

	// Test GenerateKeyPair
	kp, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	// Test RequestPIN
	headers := DefaultHeaders()
	pin, err := RequestPIN(headers, kp)
	if err != nil {
		t.Fatalf("RequestPIN failed: %v", err)
	}
	if pin.Code != "ABCD" {
		t.Errorf("Expected PIN code ABCD, got %s", pin.Code)
	}

	// Test CheckPIN
	pinResp, err := CheckPIN(pin.ID, "client-id", kp)
	if err != nil {
		t.Fatalf("CheckPIN failed: %v", err)
	}
	if pinResp.AuthToken != "test-auth-token" {
		t.Errorf("Expected auth token test-auth-token, got %s", pinResp.AuthToken)
	}

	// Test RefreshToken
	p := &Plex{
		Headers: headers,
	}

	// Need to initialize Header ClientIdentifier as it's used in RefreshToken
	p.Headers.ClientIdentifier = "client-id"

	if err := p.RefreshToken(kp); err != nil {
		t.Fatalf("RefreshToken failed: %v", err)
	}

	if p.Token != "refreshed-token" {
		t.Errorf("Expected refreshed token, got %s", p.Token)
	}
}
