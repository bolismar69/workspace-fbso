package legacy

import (
	"context"
	"time"

	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

type mockTaxRepository struct {
	federalRule    *repository.FederalTaxRule
	federalRuleErr error

	icmsRule    *repository.ICMSRule
	icmsRuleErr error

	ipiRegra    *repository.IPIRegra
	ipiRegraErr error

	productException    *repository.ProductException
	productExceptionErr error

	equivalence    *repository.TaxEquivalence
	equivalenceErr error

	simplesFaixa    *repository.SimplesFaixa
	simplesFaixaErr error

	ncmSeletivoRule    *repository.NCMSeletivoRule
	ncmSeletivoRuleErr error
}

func (m *mockTaxRepository) GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error) {
	return 0, nil
}

func (m *mockTaxRepository) GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error) {
	return 0, nil
}

func (m *mockTaxRepository) GetFederalTaxRule(ctx context.Context, regimeTributario, cstPIS, cstCOFINS string) (*repository.FederalTaxRule, error) {
	return m.federalRule, m.federalRuleErr
}

func (m *mockTaxRepository) GetICMSRule(ctx context.Context, orig, dest string) (*repository.ICMSRule, error) {
	return m.icmsRule, m.icmsRuleErr
}

func (m *mockTaxRepository) GetEquivalence(ctx context.Context, CSOSN string, tipoOperacao string) (*repository.TaxEquivalence, error) {
	return m.equivalence, m.equivalenceErr
}

func (m *mockTaxRepository) GetSimplesFaixa(ctx context.Context, anexo string, rbt12 decimal.Decimal) (*repository.SimplesFaixa, error) {
	return m.simplesFaixa, m.simplesFaixaErr
}

func (m *mockTaxRepository) GetProductException(ctx context.Context, ncmFull, ncmGroup, ufDestino string, regimeTributarioDestino string) (*repository.ProductException, error) {
	return m.productException, m.productExceptionErr
}

func (m *mockTaxRepository) GetIPIRegra(ctx context.Context, NCM string, ExIPI string, CrtEmitente string, TipoOperacaoFiscal string, PerfilComprador string, UFDestino string, ZonaEspecial bool, DataOperacao string) (*repository.IPIRegra, error) {
	return m.ipiRegra, m.ipiRegraErr
}

func (m *mockTaxRepository) GetIvaDualRule(ctx context.Context, ncm, ufDestino, municipioIBGE string) (*repository.IvaDualRule, error) {
	return nil, nil
}

func (m *mockTaxRepository) GetNCMSeletivo(ctx context.Context, ncm string) (*repository.NCMSeletivoRule, error) {
	return m.ncmSeletivoRule, m.ncmSeletivoRuleErr
}

func (m *mockTaxRepository) GetCSTReforma(ctx context.Context, flags repository.CSTFlags) (*repository.CSTReforma, error) {
	return nil, nil
}

func newMockWithFederalRule(rule *repository.FederalTaxRule) *mockTaxRepository {
	return &mockTaxRepository{federalRule: rule}
}

func newMockWithFederalRuleError(err error) *mockTaxRepository {
	return &mockTaxRepository{federalRuleErr: err, federalRule: nil}
}

func defaultFederalRule(pisAliquota, cofinsAliquota float64, excluiICMS bool) *repository.FederalTaxRule {
	return &repository.FederalTaxRule{
		Id:               1,
		RegimeTributario: "3",
		CSTPIS:           "01",
		CSTCOFINS:        "01",
		AliquotaPIS:      decimal.NewFromFloat(pisAliquota),
		AliquotaCOFINS:   decimal.NewFromFloat(cofinsAliquota),
		ExcluiICMSBase:   excluiICMS,
		InicioValidade:   time.Now().AddDate(-1, 0, 0),
	}
}
