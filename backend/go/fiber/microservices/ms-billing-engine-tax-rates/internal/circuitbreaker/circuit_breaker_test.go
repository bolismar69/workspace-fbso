package circuitbreaker

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestCircuitBreaker_Closed_Success verifica que chamadas em estado CLOSED
// passam normalmente quando a funcao subjacente nao falha.
func TestCircuitBreaker_Closed_Success(t *testing.T) {
	cb := New("test", DefaultConfig())

	callCount := 0
	fn := func() (interface{}, error) {
		callCount++
		return "ok", nil
	}

	for i := 0; i < 5; i++ {
		result, err := cb.Execute(context.Background(), fn)
		if err != nil {
			t.Fatalf("chamada %d: erro inesperado: %v", i, err)
		}
		if result != "ok" {
			t.Errorf("chamada %d: resultado = %v, esperado 'ok'", i, result)
		}
	}

	if callCount != 5 {
		t.Errorf("fn executada %d vezes, esperado 5", callCount)
	}
	if cb.State() != StateClosed {
		t.Errorf("estado = %s, esperado CLOSED", cb.State())
	}
}

// TestCircuitBreaker_OpensAfterFailures verifica que o circuito abre apos
// N falhas consecutivas (default: 3).
func TestCircuitBreaker_OpensAfterFailures(t *testing.T) {
	cb := New("test", DefaultConfig())
	testErr := errors.New("service unavailable")

	// 2 falhas — ainda nao deve abrir
	for i := 0; i < 2; i++ {
		_, err := cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
		if err != testErr {
			t.Fatalf("falha %d: deveria ter retornado testErr", i+1)
		}
	}
	if cb.State() != StateClosed {
		t.Errorf("apos 2 falhas, estado = %s, esperado CLOSED", cb.State())
	}

	// 3a falha — deve abrir
	_, err := cb.Execute(context.Background(), func() (interface{}, error) {
		return nil, testErr
	})
	if err != testErr {
		t.Fatal("deveria ter retornado testErr na 3a falha")
	}
	if cb.State() != StateOpen {
		t.Errorf("apos 3 falhas, estado = %s, esperado OPEN", cb.State())
	}
}

// TestCircuitBreaker_RejectsWhenOpen verifica que chamadas sao rejeitadas
// imediatamente quando o circuito esta OPEN.
func TestCircuitBreaker_RejectsWhenOpen(t *testing.T) {
	cb := New("test", DefaultConfig())
	testErr := errors.New("fail")

	// Abre o circuito com 3 falhas
	for i := 0; i < 3; i++ {
		cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
	}

	if cb.State() != StateOpen {
		t.Fatalf("circuito deveria estar OPEN, mas esta %s", cb.State())
	}

	// Chamada subsequente deve ser rejeitada sem executar fn
	fnCalled := false
	_, err := cb.Execute(context.Background(), func() (interface{}, error) {
		fnCalled = true
		return "never", nil
	})

	if _, ok := err.(ErrCircuitOpen); !ok {
		t.Errorf("erro = %v, esperado ErrCircuitOpen", err)
	}
	if fnCalled {
		t.Error("fn nao deveria ter sido chamada com circuito OPEN")
	}
}

// TestCircuitBreaker_HalfOpenToClosed verifica que o circuito fecha
// apos uma chamada bem-sucedida em HALF_OPEN.
func TestCircuitBreaker_HalfOpenToClosed(t *testing.T) {
	// Usa um timeout curto para HALF_OPEN nos testes
	cfg := Config{
		FailureThreshold: 2,
		FailureWindow:    60 * time.Second,
		HalfOpenTimeout:  10 * time.Millisecond, // curto para teste
	}
	cb := New("test", cfg)
	testErr := errors.New("fail")

	// Abre o circuito com 2 falhas
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
	}
	if cb.State() != StateOpen {
		t.Fatalf("estado = %s, esperado OPEN", cb.State())
	}

	// Espera o timeout de HALF_OPEN
	time.Sleep(20 * time.Millisecond)

	// Chamada bem-sucedida em HALF_OPEN → deve fechar
	result, err := cb.Execute(context.Background(), func() (interface{}, error) {
		return "recovered", nil
	})
	if err != nil {
		t.Fatalf("erro inesperado em HALF_OPEN: %v", err)
	}
	if result != "recovered" {
		t.Errorf("resultado = %v, esperado 'recovered'", result)
	}
	if cb.State() != StateClosed {
		t.Errorf("apos sucesso em HALF_OPEN, estado = %s, esperado CLOSED", cb.State())
	}
}

// TestCircuitBreaker_HalfOpenBackToOpen verifica que se a chamada de teste
// em HALF_OPEN falhar, o circuito volta a OPEN.
func TestCircuitBreaker_HalfOpenBackToOpen(t *testing.T) {
	cfg := Config{
		FailureThreshold: 2,
		FailureWindow:    60 * time.Second,
		HalfOpenTimeout:  10 * time.Millisecond,
	}
	cb := New("test", cfg)
	testErr := errors.New("fail")

	// Abre o circuito
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
	}

	// Espera HALF_OPEN
	time.Sleep(20 * time.Millisecond)

	// Chamada de teste falha → volta a OPEN
	_, err := cb.Execute(context.Background(), func() (interface{}, error) {
		return nil, testErr
	})
	if err != testErr {
		t.Fatal("deveria ter falhado em HALF_OPEN")
	}
	if cb.State() != StateOpen {
		t.Errorf("apos falha em HALF_OPEN, estado = %s, esperado OPEN", cb.State())
	}
}

// TestCircuitBreaker_FailureWindowReset verifica que falhas fora da janela
// nao acumulam para abrir o circuito.
func TestCircuitBreaker_FailureWindowReset(t *testing.T) {
	cfg := Config{
		FailureThreshold: 3,
		FailureWindow:    50 * time.Millisecond, // janela curta
		HalfOpenTimeout:  5 * time.Minute,
	}
	cb := New("test", cfg)
	testErr := errors.New("fail")

	// 2 falhas — ainda na janela
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
	}

	// Espera a janela expirar
	time.Sleep(60 * time.Millisecond)

	// Mais 2 falhas — contador deve ter resetado
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
	}

	// Circuito ainda deve estar CLOSED (falhas nao acumularam)
	if cb.State() != StateClosed {
		t.Errorf("falhas fora da janela nao deveriam abrir circuito; estado = %s", cb.State())
	}
}

// TestCircuitBreaker_SuccessResetsCounter verifica que um sucesso reseta
// o contador de falhas.
func TestCircuitBreaker_SuccessResetsCounter(t *testing.T) {
	cb := New("test", DefaultConfig())
	testErr := errors.New("fail")

	// 2 falhas
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
	}

	// 1 sucesso — reseta contador
	cb.Execute(context.Background(), func() (interface{}, error) {
		return "ok", nil
	})

	// Mais 2 falhas — nao deve abrir (contador resetou)
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func() (interface{}, error) {
			return nil, testErr
		})
	}

	if cb.State() != StateClosed {
		t.Errorf("sucesso deveria resetar contador; estado = %s, esperado CLOSED", cb.State())
	}
}
