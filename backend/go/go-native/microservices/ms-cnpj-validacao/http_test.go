package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cnpj/validate?value=04.252.011/0001-10", nil)
	w := httptest.NewRecorder()
	validateHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp cnpjValidationResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Valid != true {
		t.Fatalf("expected valid true")
	}
	if resp.Normalized == nil || *resp.Normalized != "04252011000110" {
		t.Fatalf("expected normalized digits, got %v", resp.Normalized)
	}
}

func TestValidatePost(t *testing.T) {
	body := []byte(`{"cnpj":"04.252.011/0001-11"}`)
	req := httptest.NewRequest(http.MethodPost, "/cnpj/validate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	validateHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp cnpjValidationResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Valid != false {
		t.Fatalf("expected valid false")
	}
	if resp.Normalized == nil || *resp.Normalized != "04252011000111" {
		t.Fatalf("expected normalized digits, got %v", resp.Normalized)
	}
}
