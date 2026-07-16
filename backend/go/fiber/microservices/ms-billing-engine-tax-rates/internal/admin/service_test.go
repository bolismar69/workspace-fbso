// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/admin/service_test.go
package admin

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

// mockAdminRepo implementa AdminRepository para testes.
type mockAdminRepo struct {
	rules     []IvaDualRuleOutput
	upsertFn  func(ctx context.Context, input IvaDualRuleInput, changedBy string) (*IvaDualRuleOutput, error)
}

func (m *mockAdminRepo) UpsertIvaDualRule(ctx context.Context, input IvaDualRuleInput, changedBy string) (*IvaDualRuleOutput, error) {
	if m.upsertFn != nil {
		return m.upsertFn(ctx, input, changedBy)
	}
	out := &IvaDualRuleOutput{
		ID:                   1,
		NCM:                  input.NCM,
		UFDestino:            input.UFDestino,
		MunicipioDestinoIBGE: input.MunicipioDestinoIBGE,
		AliquotaCBS:          input.AliquotaCBS,
		AliquotaIBSEstadual:  input.AliquotaIBSEstadual,
		AliquotaIBSMunicipal: input.AliquotaIBSMunicipal,
		AliquotaIS:           input.AliquotaIS,
		InicioValidade:       time.Now(),
	}
	m.rules = append(m.rules, *out)
	return out, nil
}

func (m *mockAdminRepo) ListIvaDualRules(ctx context.Context, filter ListRulesFilter) ([]IvaDualRuleOutput, error) {
	var result []IvaDualRuleOutput
	for _, r := range m.rules {
		if filter.NCM != "" && r.NCM != filter.NCM {
			continue
		}
		if filter.UF != "" && r.UFDestino != filter.UF {
			continue
		}
		result = append(result, r)
	}
	return result, nil
}

type mockCacheInvalidator struct {
	invalidatedKeys []string
}

func (m *mockCacheInvalidator) InvalidateIvaDualCache(ctx context.Context, ncm, uf, municipioIBGE string) error {
	m.invalidatedKeys = append(m.invalidatedKeys, ncm+":"+uf+":"+municipioIBGE)
	return nil
}

func dec(s string) decimal.Decimal { return decimal.RequireFromString(s) }

// ─── TST-001.01: PUT com payload válido insere nova regra ──────────────────

func TestAdmin_UpsertValid(t *testing.T) {
	repo := &mockAdminRepo{}
	cache := &mockCacheInvalidator{}
	svc := NewAdminTaxService(repo, cache)

	out, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "84713012", UFDestino: "RJ", MunicipioDestinoIBGE: "3304557",
		AliquotaCBS: dec("9"), AliquotaIBSEstadual: dec("5"),
		AliquotaIBSMunicipal: dec("2.5"), AliquotaIS: dec("1.5"),
	}, "user@taxnexus")
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if out.ID != 1 {
		t.Errorf("ID = %d, esperado 1", out.ID)
	}
	if out.NCM != "84713012" {
		t.Errorf("NCM = %s, esperado 84713012", out.NCM)
	}

	// Cache deve ter sido invalidado
	if len(cache.invalidatedKeys) != 1 {
		t.Errorf("cache não foi invalidado: %d keys", len(cache.invalidatedKeys))
	}
}

// ─── TST-001.02: PUT com mesmo NCM+UF+IBGE atualiza (upsert) ──────────────

func TestAdmin_UpsertUpdate(t *testing.T) {
	var capturedInput IvaDualRuleInput
	repo := &mockAdminRepo{
		upsertFn: func(ctx context.Context, input IvaDualRuleInput, changedBy string) (*IvaDualRuleOutput, error) {
			capturedInput = input
			return &IvaDualRuleOutput{ID: 2, NCM: input.NCM, UFDestino: input.UFDestino}, nil
		},
	}
	cache := &mockCacheInvalidator{}
	svc := NewAdminTaxService(repo, cache)

	_, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "84713012", UFDestino: "RJ",
		AliquotaCBS: dec("10"), AliquotaIBSEstadual: dec("6"),
		AliquotaIBSMunicipal: dec("3"), AliquotaIS: dec("2"),
	}, "user@taxnexus")
	if err != nil {
		t.Fatalf("erro: %v", err)
	}

	if capturedInput.AliquotaCBS.String() != "10" {
		t.Errorf("CBS = %s, esperado 10", capturedInput.AliquotaCBS)
	}
}

// ─── TST-001.03: Alíquota CBS > 100% → erro de validação ───────────────────

func TestAdmin_AliquotaAcima100(t *testing.T) {
	svc := NewAdminTaxService(&mockAdminRepo{}, &mockCacheInvalidator{})

	_, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "84713012", UFDestino: "RJ",
		AliquotaCBS:          dec("150"),
		AliquotaIBSEstadual:  dec("5"),
		AliquotaIBSMunicipal: dec("2.5"),
	}, "user")
	if err == nil {
		t.Fatal("esperado erro para CBS > 100%")
	}
}

// ─── TST-001.04: NCM inválido (não numérico) → 400 ─────────────────────────

func TestAdmin_NCMInvalido(t *testing.T) {
	svc := NewAdminTaxService(&mockAdminRepo{}, &mockCacheInvalidator{})

	_, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "ABCD1234", UFDestino: "RJ",
		AliquotaCBS: dec("9"), AliquotaIBSEstadual: dec("5"),
		AliquotaIBSMunicipal: dec("2.5"),
	}, "user")
	if err == nil {
		t.Fatal("esperado erro para NCM não numérico")
	}
}

// ─── TST-001.05: GET com filtro NCM retorna apenas regras daquele NCM ──────

func TestAdmin_ListByNCM(t *testing.T) {
	repo := &mockAdminRepo{
		rules: []IvaDualRuleOutput{
			{NCM: "84713012", UFDestino: "RJ"},
			{NCM: "84713012", UFDestino: "SP"},
			{NCM: "22030000", UFDestino: "RJ"},
		},
	}
	svc := NewAdminTaxService(repo, &mockCacheInvalidator{})

	rules, err := svc.ListRules(context.Background(), ListRulesFilter{NCM: "84713012"})
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if len(rules) != 2 {
		t.Errorf("esperado 2 regras para NCM 84713012, obtido %d", len(rules))
	}
}

// ─── TST-001.06: GET com filtro UF retorna apenas regras daquele estado ────

func TestAdmin_ListByUF(t *testing.T) {
	repo := &mockAdminRepo{
		rules: []IvaDualRuleOutput{
			{NCM: "84713012", UFDestino: "RJ"},
			{NCM: "84713012", UFDestino: "SP"},
		},
	}
	svc := NewAdminTaxService(repo, &mockCacheInvalidator{})

	rules, err := svc.ListRules(context.Background(), ListRulesFilter{UF: "SP"})
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if len(rules) != 1 {
		t.Errorf("esperado 1 regra para UF=SP, obtido %d", len(rules))
	}
	if rules[0].UFDestino != "SP" {
		t.Errorf("UF = %s, esperado SP", rules[0].UFDestino)
	}
}

// ─── TST-001.07: Auditoria — changed_by é propagado ────────────────────────

func TestAdmin_AuditoriaChangedBy(t *testing.T) {
	var capturedBy string
	repo := &mockAdminRepo{
		upsertFn: func(ctx context.Context, input IvaDualRuleInput, changedBy string) (*IvaDualRuleOutput, error) {
			capturedBy = changedBy
			return &IvaDualRuleOutput{ID: 3}, nil
		},
	}
	svc := NewAdminTaxService(repo, &mockCacheInvalidator{})

	_, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "84713012", UFDestino: "RJ",
		AliquotaCBS: dec("9"), AliquotaIBSEstadual: dec("5"),
		AliquotaIBSMunicipal: dec("2.5"),
	}, "fiscal.ana@taxnexus")
	if err != nil {
		t.Fatalf("erro: %v", err)
	}
	if capturedBy != "fiscal.ana@taxnexus" {
		t.Errorf("changed_by = %q, esperado 'fiscal.ana@taxnexus'", capturedBy)
	}
}

// ─── TST-001.08: Cache invalidado após PUT ──────────────────────────────────

func TestAdmin_CacheInvalidated(t *testing.T) {
	cache := &mockCacheInvalidator{}
	svc := NewAdminTaxService(&mockAdminRepo{}, cache)

	_, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "84713012", UFDestino: "RJ", MunicipioDestinoIBGE: "3304557",
		AliquotaCBS: dec("9"), AliquotaIBSEstadual: dec("5"),
		AliquotaIBSMunicipal: dec("2.5"), AliquotaIS: dec("1.5"),
	}, "user")
	if err != nil {
		t.Fatalf("erro: %v", err)
	}

	if len(cache.invalidatedKeys) == 0 {
		t.Error("cache não foi invalidado após upsert")
	}
	t.Logf("Keys invalidadas: %v", cache.invalidatedKeys)
}

// ─── Testes de validação adicionais ────────────────────────────────────────

func TestAdmin_AliquotaNegativa(t *testing.T) {
	svc := NewAdminTaxService(&mockAdminRepo{}, &mockCacheInvalidator{})
	_, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "84713012", UFDestino: "RJ",
		AliquotaCBS: dec("-5"), AliquotaIBSEstadual: dec("5"),
		AliquotaIBSMunicipal: dec("2.5"),
	}, "user")
	if err == nil {
		t.Fatal("esperado erro para alíquota negativa")
	}
}

func TestAdmin_UFInvalida(t *testing.T) {
	svc := NewAdminTaxService(&mockAdminRepo{}, &mockCacheInvalidator{})
	_, err := svc.UpsertRule(context.Background(), IvaDualRuleInput{
		NCM: "84713012", UFDestino: "XX1",
		AliquotaCBS: dec("9"), AliquotaIBSEstadual: dec("5"),
		AliquotaIBSMunicipal: dec("2.5"),
	}, "user")
	if err == nil {
		t.Fatal("esperado erro para UF inválida")
	}
}
