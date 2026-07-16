// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/supplier/store.go
package supplier

import (
	"context"
	"fmt"
	"sync"
)

// MemorySupplierStore implementa SupplierStore em memória para testes.
type MemorySupplierStore struct {
	mu        sync.RWMutex
	suppliers map[string]*SupplierFiscal
}

func NewMemorySupplierStore() *MemorySupplierStore {
	return &MemorySupplierStore{suppliers: make(map[string]*SupplierFiscal)}
}

func (s *MemorySupplierStore) Save(ctx context.Context, sup *SupplierFiscal) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.suppliers[sup.CNPJ]; exists {
		return fmt.Errorf("fornecedor %s já cadastrado", sup.CNPJ)
	}
	clone := *sup
	s.suppliers[sup.CNPJ] = &clone
	return nil
}

func (s *MemorySupplierStore) FindByCNPJ(ctx context.Context, cnpj string) (*SupplierFiscal, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sup, exists := s.suppliers[cnpj]
	if !exists {
		return nil, fmt.Errorf("fornecedor %s não encontrado", cnpj)
	}
	clone := *sup
	return &clone, nil
}

func (s *MemorySupplierStore) Update(ctx context.Context, cnpj string, sup *SupplierFiscal) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.suppliers[cnpj]; !exists {
		return fmt.Errorf("fornecedor %s não encontrado", cnpj)
	}
	clone := *sup
	s.suppliers[cnpj] = &clone
	return nil
}

// Count retorna o número de fornecedores armazenados (útil para testes).
func (s *MemorySupplierStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.suppliers)
}
