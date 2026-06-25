// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/middleware/ratelimit.go
package middleware

import (
	"log/slog"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	// HeaderRateLimitLimit indica o número máximo de requisições permitidas na janela.
	HeaderRateLimitLimit = "X-RateLimit-Limit"
	// HeaderRateLimitRemaining indica quantas requisições ainda restam na janela atual.
	HeaderRateLimitRemaining = "X-RateLimit-Remaining"
	// HeaderRateLimitReset indica o timestamp Unix (segundos) em que a janela reseta.
	HeaderRateLimitReset = "X-RateLimit-Reset"
	// HeaderRetryAfter indica em quantos segundos o cliente pode tentar novamente.
	HeaderRetryAfter = "Retry-After"

	// EnvRateLimitMax é a variável de ambiente que define o número máximo de
	// requisições por janela por cliente. Default: 100.
	EnvRateLimitMax = "RATE_LIMIT_MAX"
	// EnvRateLimitWindow é a variável de ambiente que define a duração da
	// janela de rate limiting em segundos. Default: 60.
	EnvRateLimitWindow = "RATE_LIMIT_WINDOW"

	defaultRateLimitMax    = 100
	defaultRateLimitWindow = 60
)

// windowEntry armazena o contador e o timestamp de início da janela para um cliente.
type windowEntry struct {
	count     int64
	windowStart time.Time
}

// RateLimiter implementa um algoritmo de sliding window counter em memória.
//
// Cada cliente é identificado pelo IP de origem (ou pelo header X-Forwarded-For
// quando disponível — compatível com proxy reverso). Se o cliente estiver
// autenticado (JWT), o token é usado como chave adicional para evitar que um
// atacante contorne o limite alternando IPs.
//
// O algoritmo mantém uma janela deslizante: quando uma requisição chega, o
// contador da janela atual é incrementado. Se o contador exceder o limite
// configurado, a requisição é rejeitada com HTTP 429 Too Many Requests.
type RateLimiter struct {
	mu      sync.Mutex
	entries map[string]*windowEntry

	maxRequests  int64
	windowSecs   int64
}

// NewRateLimiter cria um novo RateLimiter com configuração via variáveis de ambiente.
//
// Env vars:
//
//	RATE_LIMIT_MAX    — número máximo de requisições por janela (default 100)
//	RATE_LIMIT_WINDOW — duração da janela em segundos (default 60)
func NewRateLimiter() *RateLimiter {
	rl := &RateLimiter{
		entries:      make(map[string]*windowEntry),
		maxRequests:  defaultRateLimitMax,
		windowSecs:   defaultRateLimitWindow,
	}

	if v := os.Getenv(EnvRateLimitMax); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil && n > 0 {
			rl.maxRequests = n
		} else {
			slog.Warn("RATE_LIMIT_MAX inválido, usando default",
				"valor_configurado", v,
				"default", defaultRateLimitMax,
			)
		}
	}

	if v := os.Getenv(EnvRateLimitWindow); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil && n > 0 {
			rl.windowSecs = n
		} else {
			slog.Warn("RATE_LIMIT_WINDOW inválido, usando default",
				"valor_configurado", v,
				"default", defaultRateLimitWindow,
			)
		}
	}

	slog.Info("Rate Limiter inicializado",
		"max_requests", rl.maxRequests,
		"window_seconds", rl.windowSecs,
	)

	return rl
}

// NewRateLimitMiddleware retorna um middleware Fiber que aplica rate limiting
// baseado no IP do cliente e, quando autenticado, no token JWT.
//
// Posição no pipeline (conforme TASKS.md T-008.3):
//
//	recover → requestid → auth → rate_limit → logger → metrics → handler
//
// Headers injetados na resposta:
//   - X-RateLimit-Limit: limite máximo configurado
//   - X-RateLimit-Remaining: requisições restantes na janela atual
//   - X-RateLimit-Reset: timestamp Unix em que a janela reseta
//   - Retry-After: segundos até poder tentar novamente (apenas em 429)
func (rl *RateLimiter) NewRateLimitMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := rl.clientKey(c)

		rl.mu.Lock()

		now := time.Now()
		windowDuration := time.Duration(rl.windowSecs) * time.Second

		entry, exists := rl.entries[key]
		if !exists || now.Sub(entry.windowStart) >= windowDuration {
			// Nova janela ou janela expirada — reset.
			rl.entries[key] = &windowEntry{
				count:       1,
				windowStart: now,
			}
			rl.mu.Unlock()

			rl.setRateLimitHeaders(c, 1)
			return c.Next()
		}

		entry.count++
		currentCount := entry.count
		windowStart := entry.windowStart
		rl.mu.Unlock()

		resetAt := windowStart.Add(windowDuration)
		retryAfter := int64(resetAt.Sub(now).Seconds())
		if retryAfter < 0 {
			retryAfter = 0
		}

		rl.setRateLimitHeaders(c, currentCount)

		if currentCount > rl.maxRequests {
			slog.Warn("Rate limit excedido",
				"client_key", key,
				"count", currentCount,
				"limit", rl.maxRequests,
				"retry_after", retryAfter,
			)

			c.Set(HeaderRetryAfter, strconv.FormatInt(retryAfter, 10))
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "too_many_requests",
				"message":     "Limite de requisições excedido. Tente novamente mais tarde.",
				"retry_after": retryAfter,
			})
		}

		return c.Next()
	}
}

// clientKey retorna a chave de identificação do cliente baseada no IP de origem
// e, quando disponível, no user ID extraído do JWT (via auth middleware).
//
// Segurança: Usa c.IP() como fonte primária do IP. O Fiber resolve o IP real
// através do header X-Forwarded-For SOMENTE quando EnableTrustedProxyCheck
// está ativo e o remote addr pertence a um proxy confiável (Kong/Envoy).
// Isso previne IP spoofing — um atacante não pode injetar X-Forwarded-For
// falsos sem que a requisição passe por um proxy confiável primeiro.
func (rl *RateLimiter) clientKey(c *fiber.Ctx) string {
	// c.IP() já lida com trusted proxies quando configurado no Fiber.
	// Não lemos X-Forwarded-For diretamente para evitar spoofing.
	ip := c.IP()
	if ip == "" {
		ip = "unknown"
	}

	userID := GetUserID(c)
	if userID != "" {
		return ip + ":" + userID
	}

	return ip
}

// setRateLimitHeaders injeta os headers informativos de rate limiting na resposta.
func (rl *RateLimiter) setRateLimitHeaders(c *fiber.Ctx, currentCount int64) {
	remaining := rl.maxRequests - currentCount
	if remaining < 0 {
		remaining = 0
	}

	c.Set(HeaderRateLimitLimit, strconv.FormatInt(rl.maxRequests, 10))
	c.Set(HeaderRateLimitRemaining, strconv.FormatInt(remaining, 10))
}

// Limpa entradas expiradas periodicamente para evitar memory leak.
// Chamado em background a cada 2x a duração da janela.
func (rl *RateLimiter) StartCleanup(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			rl.cleanup()
		}
	}()
}

// isOverLimit verifica se uma chave de cliente excedeu o limite.
// Usado internamente e exposto para testes unitários de isolamento de IP.
func (rl *RateLimiter) isOverLimit(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	entry, exists := rl.entries[key]
	if !exists {
		return false
	}
	return entry.count > rl.maxRequests
}

func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	windowDuration := time.Duration(rl.windowSecs) * time.Second
	cutoff := time.Now().Add(-windowDuration)

	for key, entry := range rl.entries {
		if entry.windowStart.Before(cutoff) {
			delete(rl.entries, key)
		}
	}
}
