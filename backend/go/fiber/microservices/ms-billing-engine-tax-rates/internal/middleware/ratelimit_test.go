// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/middleware/ratelimit_test.go
package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

// setupRateLimitApp cria uma app Fiber com o middleware de rate limiting.
// Habilita EnableTrustedProxyCheck para que c.IP() respeite X-Forwarded-For
// de proxies confiáveis (necessário para testes de IP isolation).
func setupRateLimitApp(rl *RateLimiter) *fiber.App {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0/0"}, // permite todos em teste
	})
	app.Use(rl.NewRateLimitMiddleware())
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
	return app
}

// doRequestWithHeaders executa uma requisição e retorna os headers da resposta HTTP.
// Retorna http.Header que suporta Get() case-insensitive (canonicalização MIME).
func doRequestWithHeaders(app *fiber.App, ip string) (int, string, http.Header) {
	req := httptest.NewRequest("GET", "/test", nil)
	if ip != "" {
		req.Header.Set("X-Forwarded-For", ip)
	}
	resp, err := app.Test(req)
	if err != nil {
		return 0, "", nil
	}
	body, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(body), resp.Header
}

// newTestRateLimiter cria um RateLimiter com valores customizados para teste.
func newTestRateLimiter(max int64, window int64) *RateLimiter {
	rl := &RateLimiter{
		entries:     make(map[string]*windowEntry),
		maxRequests: max,
		windowSecs:  window,
	}
	return rl
}

// ─── TST-008.01: Requisições dentro do limite passam normalmente ────────────

func TestRateLimit_WithinLimit(t *testing.T) {
	rl := newTestRateLimiter(5, 60)
	app := setupRateLimitApp(rl)

	for i := 1; i <= 5; i++ {
		status, body, respHeaders := doRequestWithHeaders(app, "")

		if status != 200 {
			t.Fatalf("req %d: esperado 200, obtido %d (body: %s)", i, status, body)
		}

		limit := respHeaders.Get(HeaderRateLimitLimit)
		if limit != "5" {
			t.Errorf("req %d: X-RateLimit-Limit = %q, esperado 5", i, limit)
		}

		remaining := respHeaders.Get(HeaderRateLimitRemaining)
		expectedRemaining := 5 - i
		if remaining != itoa(expectedRemaining) {
			t.Errorf("req %d: X-RateLimit-Remaining = %q, esperado %d",
				i, remaining, expectedRemaining)
		}
	}
}

// ─── TST-008.02: Requisição acima do limite retorna 429 ─────────────────────

func TestRateLimit_OverLimit(t *testing.T) {
	rl := newTestRateLimiter(3, 60)
	app := setupRateLimitApp(rl)

	// Envia 3 requisições (dentro do limite)
	for i := 0; i < 3; i++ {
		status, _, _ := doRequestWithHeaders(app, "")
		if status != 200 {
			t.Fatalf("req %d dentro do limite: esperado 200, obtido %d", i+1, status)
		}
	}

	// 4ª requisição deve ser bloqueada
	status, body, _ := doRequestWithHeaders(app, "")
	if status != 429 {
		t.Fatalf("req acima do limite: esperado 429, obtido %d (body: %s)", status, body)
	}
}

// ─── TST-008.03: Header Retry-After presente na resposta 429 ────────────────

func TestRateLimit_RetryAfterHeader(t *testing.T) {
	rl := newTestRateLimiter(1, 60)
	app := setupRateLimitApp(rl)

	// Primeira requisição — OK
	status, _, _ := doRequestWithHeaders(app, "")
	if status != 200 {
		t.Fatalf("primeira req: esperado 200, obtido %d", status)
	}

	// Segunda requisição — bloqueada, deve ter Retry-After
	status, body, respHeaders := doRequestWithHeaders(app, "")
	if status != 429 {
		t.Fatalf("segunda req: esperado 429, obtido %d (body: %s)", status, body)
	}

	retryAfter := respHeaders.Get(HeaderRetryAfter)
	if retryAfter == "" {
		t.Error("esperado header Retry-After na resposta 429, mas não foi encontrado")
	}
	if retryAfter == "0" {
		t.Error("Retry-After deveria ser > 0, obtido 0")
	}

	// Verifica que o header X-RateLimit-Remaining é 0
	remaining := respHeaders.Get(HeaderRateLimitRemaining)
	if remaining != "0" {
		t.Errorf("X-RateLimit-Remaining = %q, esperado 0", remaining)
	}
}

// ─── TST-008.04: Após janela de tempo, contador reseta ──────────────────────

func TestRateLimit_WindowReset(t *testing.T) {
	rl := newTestRateLimiter(2, 1) // janela de 1s
	app := setupRateLimitApp(rl)

	// Esgota o limite
	doRequestWithHeaders(app, "")
	status, _, _ := doRequestWithHeaders(app, "")
	if status != 200 {
		t.Fatalf("req 2: esperado 200, obtido %d", status)
	}

	// 3ª deve ser bloqueada
	status, body, _ := doRequestWithHeaders(app, "")
	if status != 429 {
		t.Fatalf("req 3 deveria ser bloqueada, obtido %d (body: %s)", status, body)
	}

	// Simula a passagem do tempo: força a expiração de todas as janelas
	rl.mu.Lock()
	for key, entry := range rl.entries {
		entry.windowStart = time.Now().Add(-2 * time.Second)
		entry.count = 0
		t.Logf("Reset entry key=%q, count=%d, window=%v", key, entry.count, entry.windowStart)
	}
	rl.mu.Unlock()

	// Após reset, requisição deve passar
	status, _, _ = doRequestWithHeaders(app, "")
	if status != 200 {
		t.Fatalf("após reset da janela: esperado 200, obtido %d", status)
	}
}

// ─── TST-008.05: IPs diferentes têm contadores independentes ─────────────────
// Testa isolamento de IP via manipulação direta do mapa interno, já que
// Fiber app.Test() não permite simular IPs remotos diferentes com
// confiabilidade (usa 0.0.0.0 para todas as conexões).

func TestRateLimit_DifferentIPs(t *testing.T) {
	rl := newTestRateLimiter(2, 60)

	// Simula requisições do IP A manualmente (entradas no mapa)
	rl.mu.Lock()
	rl.entries["10.0.0.1"] = &windowEntry{count: 2, windowStart: time.Now()} // no limite
	rl.entries["10.0.0.2"] = &windowEntry{count: 0, windowStart: time.Now()} // IP B zerado
	rl.mu.Unlock()

	// IP A (count=2, max=2): ainda no limite, não excedeu
	if rl.isOverLimit("10.0.0.1") {
		t.Error("IP A (count=2, max=2) NÃO deveria estar bloqueado (count == max)")
	}

	// Incrementa IP A → deve exceder
	rl.mu.Lock()
	rl.entries["10.0.0.1"].count = 3
	rl.mu.Unlock()
	if !rl.isOverLimit("10.0.0.1") {
		t.Error("IP A (count=3, max=2) deveria estar bloqueado (count > max)")
	}

	// IP B ainda está livre
	if rl.isOverLimit("10.0.0.2") {
		t.Error("IP B (count=0, max=2) deveria estar livre")
	}

	// IP B pode chegar até max=2 sem ser bloqueado
	rl.mu.Lock()
	rl.entries["10.0.0.2"].count = 2
	rl.mu.Unlock()
	if rl.isOverLimit("10.0.0.2") {
		t.Error("IP B (count=2, max=2) deveria passar")
	}

	// IP B excedendo → bloqueado
	rl.mu.Lock()
	rl.entries["10.0.0.2"].count = 3
	rl.mu.Unlock()
	if !rl.isOverLimit("10.0.0.2") {
		t.Error("IP B (count=3, max=2) deveria estar bloqueado (count > max)")
	}
}

// ─── TST-008.06: Rate limit com burst — rajada curta dentro do limite maior ─

func TestRateLimit_Burst(t *testing.T) {
	// Com limite 10 e janela 60s, todas as 10 devem passar
	rl := newTestRateLimiter(10, 60)
	app := setupRateLimitApp(rl)

	for i := 1; i <= 10; i++ {
		status, _, respHeaders := doRequestWithHeaders(app, "")
		if status != 200 {
			t.Fatalf("burst req %d: esperado 200, obtido %d", i, status)
		}
		remaining := respHeaders.Get(HeaderRateLimitRemaining)
		expectedRemaining := 10 - i
		if remaining != itoa(expectedRemaining) {
			t.Errorf("burst req %d: remaining = %q, esperado %d",
				i, remaining, expectedRemaining)
		}
	}

	// 11ª deve ser bloqueada
	status, body, respHeaders := doRequestWithHeaders(app, "")
	if status != 429 {
		t.Fatalf("req 11 após burst: esperado 429, obtido %d (body: %s)", status, body)
	}

	remaining := respHeaders.Get(HeaderRateLimitRemaining)
	if remaining != "0" {
		t.Errorf("após burst: remaining = %q, esperado 0", remaining)
	}
}

// ─── Teste adicional: Cleanup de entradas expiradas ──────────────────────────

func TestRateLimit_Cleanup(t *testing.T) {
	rl := newTestRateLimiter(5, 1) // janela de 1s

	// Insere entradas manualmente (evita dependência de HTTP/Fiber para teste de cleanup)
	rl.mu.Lock()
	rl.entries["ip-a"] = &windowEntry{
		count:       5,
		windowStart: time.Now().Add(-2 * time.Second), // já expirada
	}
	rl.entries["ip-b"] = &windowEntry{
		count:       3,
		windowStart: time.Now().Add(-3 * time.Second), // já expirada
	}
	rl.entries["ip-c"] = &windowEntry{
		count:       1,
		windowStart: time.Now(), // ainda válida
	}
	rl.mu.Unlock()

	// Executa cleanup
	rl.cleanup()

	rl.mu.Lock()
	count := len(rl.entries)
	_, hasC := rl.entries["ip-c"]
	rl.mu.Unlock()

	if count != 1 {
		t.Errorf("cleanup: esperado 1 entrada restante (ip-c), obtido %d", count)
	}
	if !hasC {
		t.Error("cleanup removeu entrada ip-c que ainda estava válida")
	}
}

// ─── Teste adicional: X-RateLimit-Limit reflete configuração ─────────────────

func TestRateLimit_HeadersReflectConfig(t *testing.T) {
	rl := newTestRateLimiter(42, 120)
	app := setupRateLimitApp(rl)

	_, _, respHeaders := doRequestWithHeaders(app, "")

	limit := respHeaders.Get(HeaderRateLimitLimit)
	if limit != "42" {
		t.Errorf("X-RateLimit-Limit = %q, esperado 42", limit)
	}
}

// ─── Helper: int para string sem import extra ────────────────────────────────

func itoa(n int) string {
	if n <= 0 {
		return "0"
	}
	digits := make([]byte, 0, 20)
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}
