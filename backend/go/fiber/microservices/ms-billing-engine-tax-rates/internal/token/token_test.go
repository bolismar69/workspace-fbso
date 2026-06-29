// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/token/token_test.go
package token

import (
	"context"
	"database/sql"
	"sync"
	"testing"
	"time"

	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

// mockTaxRepo implementa repository.TaxRepository para testes.
type mockTaxRepo struct {
	ivaDualRule *repository.IvaDualRule
	err         error
}

func (m *mockTaxRepo) GetIvaDualRule(ctx context.Context, ncm, ufDestino, municipioIBGE string) (*repository.IvaDualRule, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.ivaDualRule, nil
}

// Stubs para os demais métodos da interface (não usados nos testes de token).
func (m *mockTaxRepo) GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error) {
	return 0, nil
}
func (m *mockTaxRepo) GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error) {
	return 0, nil
}
func (m *mockTaxRepo) GetFederalTaxRule(ctx context.Context, regimeTributario, cstPIS, cstCOFINS string) (*repository.FederalTaxRule, error) {
	return nil, nil
}
func (m *mockTaxRepo) GetICMSRule(ctx context.Context, orig, dest string) (*repository.ICMSRule, error) {
	return nil, nil
}
func (m *mockTaxRepo) GetEquivalence(ctx context.Context, CSOSN string, tipoOperacao string) (*repository.TaxEquivalence, error) {
	return nil, nil
}
func (m *mockTaxRepo) GetSimplesFaixa(ctx context.Context, anexo string, rbt12 decimal.Decimal) (*repository.SimplesFaixa, error) {
	return nil, nil
}
func (m *mockTaxRepo) GetProductException(ctx context.Context, ncmFull, ncmGroup, ufDestino string, regimeTributarioDestino string) (*repository.ProductException, error) {
	return nil, nil
}
func (m *mockTaxRepo) GetIPIRegra(ctx context.Context, NCM string, ExIPI string, CrtEmitente string, TipoOperacaoFiscal string, PerfilComprador string, UFDestino string, ZonaEspecial bool, DataOperacao string) (*repository.IPIRegra, error) {
	return nil, nil
}
func (m *mockTaxRepo) GetNCMSeletivo(ctx context.Context, ncm string) (*repository.NCMSeletivoRule, error) {
	return nil, nil
}

// validRule retorna uma regra IVA Dual válida para testes.
func validRule() *repository.IvaDualRule {
	return &repository.IvaDualRule{
		NCM:                  "84713012",
		UFDestino:            "RJ",
		MunicipioDestinoIBGE: sql.NullString{String: "3304557", Valid: true},
		AliquotaCBS:          decimal.NewFromFloat(9.0),
		AliquotaIBSEstadual:  decimal.NewFromFloat(5.0),
		AliquotaIBSMunicipal: decimal.NewFromFloat(2.5),
		AliquotaIS:           decimal.NewFromFloat(1.5),
	}
}

func newTestService(repo *mockTaxRepo) (*TokenService, *MemoryTokenStore) {
	store := NewMemoryTokenStore()
	svc := &TokenService{
		store:      store,
		repo:       repo,
		ttlMinutes: 60,
	}
	return svc, store
}

// ─── TST-002.01: Geração de token com dados válidos ────────────────────────

func TestToken_GenerateValid(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, _ := newTestService(repo)

	tok, err := svc.Generate(context.Background(), TokenGenerateRequest{
		NCM: "84713012", UFDestino: "RJ", UFOrigem: "SP", MunicipioIBGE: "3304557",
	})
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if tok.ID == "" {
		t.Error("ID do token não pode ser vazio")
	}
	if tok.AliquotaCBS.String() != "9" {
		t.Errorf("CBS = %s, esperado 9", tok.AliquotaCBS)
	}
	if tok.AliquotaIBSEstadual.String() != "5" {
		t.Errorf("IBS Estadual = %s, esperado 5", tok.AliquotaIBSEstadual)
	}
	if tok.ExpiresAt.Before(time.Now()) {
		t.Error("ExpiresAt deve ser no futuro")
	}
	// TTL de 60 min deve resultar em ~60 min no futuro
	expectedExpiry := time.Now().Add(60 * time.Minute)
	if tok.ExpiresAt.Sub(expectedExpiry).Abs() > 2*time.Second {
		t.Errorf("TTL errado: expires_at=%v, esperado ~%v", tok.ExpiresAt, expectedExpiry)
	}
}

// ─── TST-002.02: Cálculo com token_id válido usa aliquotas do token ────────

func TestToken_ValidateValid(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, store := newTestService(repo)

	// Gera token
	tok, err := svc.Generate(context.Background(), TokenGenerateRequest{
		NCM: "84713012", UFDestino: "RJ", UFOrigem: "SP", MunicipioIBGE: "3304557",
	})
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	// Valida o token
	validated, err := svc.Validate(context.Background(), tok.ID)
	if err != nil {
		t.Fatalf("token deveria ser válido: %v", err)
	}

	if validated.AliquotaCBS.String() != "9" {
		t.Errorf("token CBS = %s, esperado 9", validated.AliquotaCBS)
	}

	// Verifica persistência
	if store.Count() != 1 {
		t.Errorf("store deveria ter 1 token, tem %d", store.Count())
	}
}

// ─── TST-002.03: Token expirado retorna ErrTokenExpired ─────────────────────

func TestToken_Expired(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, store := newTestService(repo)

	tok, err := svc.Generate(context.Background(), TokenGenerateRequest{
		NCM: "84713012", UFDestino: "RJ", UFOrigem: "SP",
	})
	if err != nil {
		t.Fatalf("erro ao gerar token: %v", err)
	}

	// Força expiração manipulando o store
	store.mu.Lock()
	stored, _ := store.tokens[tok.ID]
	stored.ExpiresAt = time.Now().Add(-1 * time.Hour) // expirou há 1h
	store.mu.Unlock()

	_, err = svc.Validate(context.Background(), tok.ID)
	if err == nil {
		t.Fatal("esperado erro de token expirado")
	}
	if _, ok := err.(ErrTokenExpired); !ok {
		t.Errorf("esperado ErrTokenExpired, obtido %T: %v", err, err)
	}
}

// ─── TST-002.04: Token não encontrado retorna erro ──────────────────────────

func TestToken_NotFound(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, _ := newTestService(repo)

	_, err := svc.Validate(context.Background(), "uuid-inexistente")
	if err == nil {
		t.Fatal("esperado erro para token inexistente")
	}
}

// ─── TST-002.05: GET token status — válido e expirado ────────────────────────

func TestToken_Status(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, store := newTestService(repo)

	tok, _ := svc.Generate(context.Background(), TokenGenerateRequest{
		NCM: "84713012", UFDestino: "RJ", UFOrigem: "SP",
	})

	// Token válido
	st, err := svc.Status(context.Background(), tok.ID)
	if err != nil {
		t.Fatalf("status erro: %v", err)
	}
	if st.Status != "valido" {
		t.Errorf("status = %q, esperado 'valido'", st.Status)
	}

	// Força expiração
	store.mu.Lock()
	store.tokens[tok.ID].ExpiresAt = time.Now().Add(-1 * time.Hour)
	store.mu.Unlock()

	st, err = svc.Status(context.Background(), tok.ID)
	if err != nil {
		t.Fatalf("status não deve falhar para token expirado: %v", err)
	}
	if st.Status != "expirado" {
		t.Errorf("status = %q, esperado 'expirado'", st.Status)
	}
}

// ─── TST-002.06: Idempotência — duas gerações criam tokens diferentes ───────

func TestToken_Idempotency(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, _ := newTestService(repo)

	req := TokenGenerateRequest{NCM: "84713012", UFDestino: "RJ", UFOrigem: "SP", MunicipioIBGE: "3304557"}

	tok1, _ := svc.Generate(context.Background(), req)
	tok2, _ := svc.Generate(context.Background(), req)

	if tok1.ID == tok2.ID {
		t.Error("dois tokens devem ter IDs diferentes (idempotência por geração)")
	}

	// Mas as aliquotas devem ser as mesmas (mesma regra)
	if !tok1.AliquotaCBS.Equal(tok2.AliquotaCBS) {
		t.Error("aliquotas CBS devem ser iguais entre tokens do mesmo NCM/UF")
	}
}

// ─── TST-002.07: Concorrência — múltiplas validações simultâneas ────────────

func TestToken_Concurrency(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, _ := newTestService(repo)

	tok, _ := svc.Generate(context.Background(), TokenGenerateRequest{
		NCM: "84713012", UFDestino: "RJ", UFOrigem: "SP",
	})

	var wg sync.WaitGroup
	errs := make(chan error, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := svc.Validate(context.Background(), tok.ID)
			if err != nil {
				errs <- err
			}
		}()
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		t.Errorf("validação concorrente falhou: %v", err)
	}
}

// ─── TST-002.08: Status de token inexistente retorna erro ───────────────────

func TestToken_StatusNotFound(t *testing.T) {
	repo := &mockTaxRepo{ivaDualRule: validRule()}
	svc, _ := newTestService(repo)

	_, err := svc.Status(context.Background(), "uuid-que-nao-existe")
	if err == nil {
		t.Fatal("esperado erro para status de token inexistente")
	}
}
