// path: backend/go/fiber/microservices/ms-tax-individual-income/services/inss_client.go
package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"taxnexus-individual-core-lib/models"
	"time"
)

type INSSClient struct {
	BaseURL string
	HTTPClient *http.Client
}

func NewINSSClient(baseURL string) *INSSClient {
	return &INSSClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *INSSClient) FetchINSS(ctx context.Context, req models.UniversalTaxRequest) (models.TaxResponse, error) {
	// Reutilizamos o modelo UniversalTaxRequest para a chamada
	body, _ := json.Marshal(req)
	
	httpReq, _ := http.NewRequestWithContext(ctx, "POST", c.BaseURL+"/api/v1/calculate/inss", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	
	// Repassa o Trace ID para manter a rastreabilidade entre serviços
	if tid, ok := ctx.Value("requestid").(string); ok {
		httpReq.Header.Set("X-Request-ID", tid)
	}

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return models.TaxResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.TaxResponse{}, fmt.Errorf("inss service returned status: %d", resp.StatusCode)
	}

	var taxRes models.TaxResponse
	if err := json.NewDecoder(resp.Body).Decode(&taxRes); err != nil {
		return models.TaxResponse{}, err
	}

	return taxRes, nil
}
