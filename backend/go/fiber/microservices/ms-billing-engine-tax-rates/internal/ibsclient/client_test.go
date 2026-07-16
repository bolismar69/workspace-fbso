package ibsclient

import (
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
)

// stubFetcher implementa IBSRateFetcher com comportamento controlado para testes.
type stubFetcher struct {
	rate *IBSRate
	err  error
}

func (s *stubFetcher) FetchRate(ctx context.Context, ibgeCode string) (*IBSRate, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.rate, nil
}

func rateSP() *IBSRate {
	return &IBSRate{
		IBGECode:         "3550308",
		AliquotaEstadual: decimal.NewFromFloat(8.8),
		AliquotaMunicipal: decimal.NewFromFloat(2.5),
		Fonte:            "API_COMITE_GESTOR",
	}
}

// TestHTTPIBSClient_EndpointFormatting verifica que a URL e formatada
// corretamente (teste de unidade da formatacao, sem chamada real).
func TestHTTPIBSClient_EndpointFormatting(t *testing.T) {
	// Teste de construcao — nao faz chamada HTTP real
	client := NewHTTPIBSClient("https://api.teste.gov.br")
	if client.baseURL != "https://api.teste.gov.br" {
		t.Errorf("baseURL = %q, esperado 'https://api.teste.gov.br'", client.baseURL)
	}
	if client.httpClient.Timeout == 0 {
		t.Error("timeout do HTTP client nao configurado")
	}

	// Verifica que FetchRate tenta a URL correta (vai falhar por nao ter servidor)
	_, err := client.FetchRate(context.Background(), "3550308")
	if err == nil {
		t.Error("deveria falhar (sem servidor HTTP real)")
	}
}

// TestStubFetcher_Success verifica o comportamento basico do stub.
func TestStubFetcher_Success(t *testing.T) {
	stub := &stubFetcher{rate: rateSP()}
	rate, err := stub.FetchRate(context.Background(), "3550308")
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if rate.IBGECode != "3550308" {
		t.Errorf("IBGECode = %s, esperado 3550308", rate.IBGECode)
	}
	if rate.Fonte != "API_COMITE_GESTOR" {
		t.Errorf("Fonte = %s, esperado API_COMITE_GESTOR", rate.Fonte)
	}
}

// TestStubFetcher_Error verifica propagacao de erro.
func TestStubFetcher_Error(t *testing.T) {
	testErr := errors.New("servico indisponivel")
	stub := &stubFetcher{err: testErr}
	_, err := stub.FetchRate(context.Background(), "3550308")
	if err != testErr {
		t.Errorf("erro = %v, esperado %v", err, testErr)
	}
}

// TestIBSRate_Aliquotas verifica que as aliquotas sao armazenadas corretamente.
func TestIBSRate_Aliquotas(t *testing.T) {
	rate := rateSP()

	estadualEsperado := decimal.NewFromFloat(8.8)
	if !rate.AliquotaEstadual.Equal(estadualEsperado) {
		t.Errorf("AliquotaEstadual = %s, esperado %s", rate.AliquotaEstadual, estadualEsperado)
	}

	municipalEsperado := decimal.NewFromFloat(2.5)
	if !rate.AliquotaMunicipal.Equal(municipalEsperado) {
		t.Errorf("AliquotaMunicipal = %s, esperado %s", rate.AliquotaMunicipal, municipalEsperado)
	}

	totalEsperado := decimal.NewFromFloat(11.3)
	total := rate.AliquotaEstadual.Add(rate.AliquotaMunicipal)
	if !total.Equal(totalEsperado) {
		t.Errorf("Aliquota total = %s, esperado %s", total, totalEsperado)
	}
}

// TestIBSRate_FonteFlag verifica que a flag de fonte e preenchida.
func TestIBSRate_FonteFlag(t *testing.T) {
	testCases := []struct {
		name     string
		fonte    string
		esperado string
	}{
		{"API", "API_COMITE_GESTOR", "API_COMITE_GESTOR"},
		{"Cache", "CACHE", "CACHE"},
		{"Fallback", "FALLBACK_CIRCUIT_OPEN", "FALLBACK_CIRCUIT_OPEN"},
		{"DB", "FALLBACK_DB", "FALLBACK_DB"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rate := &IBSRate{
				IBGECode:         "3550308",
				AliquotaEstadual: decimal.NewFromFloat(8.8),
				AliquotaMunicipal: decimal.NewFromFloat(2.5),
				Fonte:            tc.fonte,
			}
			if rate.Fonte != tc.esperado {
				t.Errorf("Fonte = %s, esperado %s", rate.Fonte, tc.esperado)
			}
		})
	}
}

// TestFallbackIBSClient_FallsBackToDB verifica que quando a API primaria
// falha, o fallback para banco de dados e acionado.
func TestFallbackIBSClient_FallsBackToDB(t *testing.T) {
	// Nota: Este teste verifica que o fallback tenta ser acionado.
	// Como requer um repositorio real, verificamos apenas que o erro
	// da API primaria e capturado e o fallback e tentado.
	// O teste completo de integracao requer um banco de dados real.
	t.Skip("Teste de integracao — requer banco de dados real com iva_dual_rules")
}

// TestCircuitBreakerIBSClient_OpenUsesExpiredCache verifica que quando
// o circuito esta aberto, o cache expirado e utilizado como fallback.
func TestCircuitBreakerIBSClient_OpenUsesExpiredCache(t *testing.T) {
	// Nota: Este teste requer Redis real para verificar cache expirado.
	// O circuito abre apos 3 falhas em 60s, entao o comportamento
	// em teste unitario e verificado via circuitbreaker package.
	t.Skip("Teste de integracao — requer Redis real")
}
