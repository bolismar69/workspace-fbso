// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/token/memory_store.go
package token

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// MemoryTokenStore implementa TokenStore em memória (para testes unitários).
type MemoryTokenStore struct {
	mu     sync.RWMutex
	tokens map[string]*TaxToken
}

// NewMemoryTokenStore cria um store em memória.
func NewMemoryTokenStore() *MemoryTokenStore {
	return &MemoryTokenStore{
		tokens: make(map[string]*TaxToken),
	}
}

func (s *MemoryTokenStore) Save(ctx context.Context, token *TaxToken) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tokens[token.ID]; exists {
		return fmt.Errorf("token %s já existe", token.ID)
	}

	clone := *token
	s.tokens[token.ID] = &clone
	return nil
}

func (s *MemoryTokenStore) FindByID(ctx context.Context, id string) (*TaxToken, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, exists := s.tokens[id]
	if !exists {
		return nil, fmt.Errorf("token %s não encontrado", id)
	}

	clone := *t
	return &clone, nil
}

func (s *MemoryTokenStore) DeleteExpired(ctx context.Context) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var deleted int64
	now := time.Now()
	for id, token := range s.tokens {
		if now.After(token.ExpiresAt) {
			delete(s.tokens, id)
			deleted++
		}
	}
	return deleted, nil
}

// Count retorna o número de tokens armazenados (útil para testes).
func (s *MemoryTokenStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.tokens)
}
