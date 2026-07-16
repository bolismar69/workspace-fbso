// Package circuitbreaker implementa o padrao Circuit Breaker para protecao
// de chamadas a APIs externas, conforme especificado em:
//
//	PROCEDURE-FIN-00001 SOP-014
//	RULES-CATALOG-FIN-00001 BR-TAX-ACT-001
//
// Estados do Circuit Breaker:
//
//	CLOSED     — operacao normal, chamadas passam direto
//	OPEN       — circuito aberto apos N falhas consecutivas, chamadas sao
//	             rejeitadas imediatamente (fast-fail)
//	HALF_OPEN  — apos timeout, permite uma chamada de teste; se falhar,
//	             volta a OPEN; se sucesso, volta a CLOSED
//
// Parametros (configuraveis):
//
//	failureThreshold — numero de falhas consecutivas para abrir (default: 3)
//	failureWindow    — janela de tempo para contar falhas (default: 60s)
//	halfOpenTimeout  — tempo ate tentar HALF_OPEN apos abrir (default: 5min)
//
// Thread-safe: utiliza sync.Mutex para proteger o estado interno.
package circuitbreaker

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

// State representa o estado atual do circuit breaker.
type State string

const (
	StateClosed   State = "CLOSED"
	StateOpen     State = "OPEN"
	StateHalfOpen State = "HALF_OPEN"
)

// Config contem a configuracao do circuit breaker.
type Config struct {
	// FailureThreshold e o numero de falhas consecutivas para abrir o circuito.
	FailureThreshold int

	// FailureWindow e a janela de tempo (em segundos) para contar falhas.
	FailureWindow time.Duration

	// HalfOpenTimeout e o tempo que o circuito permanece OPEN antes de
	// tentar HALF_OPEN.
	HalfOpenTimeout time.Duration
}

// DefaultConfig retorna a configuracao padrao conforme SOP-014:
//
//	3 falhas em 60s → OPEN
//	5 minutos → HALF_OPEN
func DefaultConfig() Config {
	return Config{
		FailureThreshold: 3,
		FailureWindow:    60 * time.Second,
		HalfOpenTimeout:  5 * time.Minute,
	}
}

// CircuitBreaker protege chamadas a servicos externos contra falhas em cascata.
type CircuitBreaker struct {
	config Config
	state  State

	mu sync.Mutex

	failureCount    int
	lastFailureTime time.Time
	openedAt        time.Time

	// name identifica este circuit breaker nos logs (ex: "ibs-api")
	name string
}

// New cria um novo CircuitBreaker com a configuracao especificada.
func New(name string, config Config) *CircuitBreaker {
	return &CircuitBreaker{
		config: config,
		state:  StateClosed,
		name:   name,
	}
}

// State retorna o estado atual do circuit breaker (thread-safe).
func (cb *CircuitBreaker) State() State {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

// Execute executa a funcao fn protegida pelo circuit breaker.
//
// Retorna o resultado de fn e um error. Se o circuito estiver OPEN,
// retorna ErrCircuitOpen imediatamente sem executar fn.
//
// Em estado HALF_OPEN, permite exatamente uma chamada de teste.
// Se falhar, volta a OPEN. Se passar, volta a CLOSED.
func (cb *CircuitBreaker) Execute(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	cb.mu.Lock()

	switch cb.state {
	case StateOpen:
		// Verifica se ja passou o timeout para tentar HALF_OPEN
		if time.Since(cb.openedAt) >= cb.config.HalfOpenTimeout {
			slog.Info("Circuit Breaker: transitando OPEN → HALF_OPEN",
				"name", cb.name,
				"opened_for", time.Since(cb.openedAt),
			)
			cb.state = StateHalfOpen
			cb.mu.Unlock()
			return cb.executeHalfOpen(ctx, fn)
		}

		cb.mu.Unlock()
		slog.Warn("Circuit Breaker: chamada rejeitada (circuito OPEN)",
			"name", cb.name,
			"opened_at", cb.openedAt,
		)
		return nil, ErrCircuitOpen{}

	case StateHalfOpen:
		cb.mu.Unlock()
		return cb.executeHalfOpen(ctx, fn)

	default: // StateClosed
		cb.mu.Unlock()
		return cb.executeClosed(ctx, fn)
	}
}

// executeClosed executa a funcao em estado CLOSED. Se falhar, registra a falha
// e verifica se deve abrir o circuito.
func (cb *CircuitBreaker) executeClosed(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	result, err := fn()
	if err != nil {
		cb.mu.Lock()
		cb.recordFailure()
		cb.mu.Unlock()
		return nil, err
	}

	// Sucesso — reseta contador de falhas
	cb.mu.Lock()
	cb.failureCount = 0
	cb.mu.Unlock()

	return result, nil
}

// executeHalfOpen executa a funcao em estado HALF_OPEN.
// Se falhar, volta a OPEN. Se passar, volta a CLOSED.
func (cb *CircuitBreaker) executeHalfOpen(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	result, err := fn()
	if err != nil {
		cb.mu.Lock()
		slog.Warn("Circuit Breaker: teste HALF_OPEN falhou, voltando a OPEN",
			"name", cb.name,
			"error", err,
		)
		cb.state = StateOpen
		cb.openedAt = time.Now()
		cb.failureCount = cb.config.FailureThreshold // mantem aberto
		cb.mu.Unlock()
		return nil, err
	}

	// Sucesso — fecha o circuito
	cb.mu.Lock()
	slog.Info("Circuit Breaker: teste HALF_OPEN ok, fechando circuito",
		"name", cb.name,
	)
	cb.state = StateClosed
	cb.failureCount = 0
	cb.mu.Unlock()

	return result, nil
}

// recordFailure registra uma falha e verifica se deve abrir o circuito.
// Deve ser chamada com o mutex ja adquirido.
func (cb *CircuitBreaker) recordFailure() {
	now := time.Now()

	// Reseta contador se a ultima falha foi fora da janela
	if cb.failureCount > 0 && now.Sub(cb.lastFailureTime) > cb.config.FailureWindow {
		cb.failureCount = 0
	}

	cb.failureCount++
	cb.lastFailureTime = now

	if cb.failureCount >= cb.config.FailureThreshold {
		slog.Warn("Circuit Breaker: limite de falhas atingido, abrindo circuito",
			"name", cb.name,
			"failures", cb.failureCount,
			"threshold", cb.config.FailureThreshold,
		)
		cb.state = StateOpen
		cb.openedAt = now
	}
}

// ErrCircuitOpen e retornado quando o circuit breaker esta OPEN
// e rejeita a chamada imediatamente.
type ErrCircuitOpen struct{}

func (e ErrCircuitOpen) Error() string {
	return "circuit breaker is OPEN — chamada rejeitada por fast-fail"
}
