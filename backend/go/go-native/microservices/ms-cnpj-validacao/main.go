package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type cnpjValidationRequest struct {
	CNPJ *string `json:"cnpj"`
}

type cnpjValidationResponse struct {
	Input      *string `json:"input"`
	Normalized *string `json:"normalized"`
	Valid      bool    `json:"valid"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cnpj/validate", validateHandler)

	addr := "0.0.0.0:" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, withJSONHeaders(mux)); err != nil {
		log.Fatal(err)
	}
}

func withJSONHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var input *string
		if values, ok := r.URL.Query()["value"]; ok {
			value := ""
			if len(values) > 0 {
				value = values[0]
			}
			input = &value
		}
		resp := validateValue(input)
		writeJSON(w, http.StatusOK, resp)
	case http.MethodPost:
		req, err := decodeRequestBody(r)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
			return
		}

		var input *string
		if req != nil {
			input = req.CNPJ
		}
		resp := validateValue(input)
		writeJSON(w, http.StatusOK, resp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func decodeRequestBody(r *http.Request) (*cnpjValidationRequest, error) {
	if r.Body == nil {
		return nil, nil
	}
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)

	var req cnpjValidationRequest
	if err := dec.Decode(&req); err != nil {
		// Treat empty body like null request (similar to Quarkus passing null)
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	// Disallow trailing non-whitespace content.
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		if err == nil {
			return nil, errors.New("multiple json values")
		}
		return nil, err
	}

	return &req, nil
}

func validateValue(input *string) cnpjValidationResponse {
	normalized := normalizeCNPJ(input)
	valid := isValidCNPJ(input)
	return cnpjValidationResponse{
		Input:      input,
		Normalized: normalized,
		Valid:      valid,
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
