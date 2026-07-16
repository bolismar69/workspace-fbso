package reforma

import (
	"context"
	"testing"
	"time"

	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"github.com/shopspring/decimal"
)

type mockRepo struct {
	ivaDualRule      *repository.IvaDualRule
	ivaDualRuleErr   error
	cstReformaRule   *repository.CSTReforma
	cstReformaRuleErr error
}

func (m *mockRepo) GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error) { return 0, nil }
func (m *mockRepo) GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error)    { return 0, nil }
func (m *mockRepo) GetFederalTaxRule(ctx context.Context, regimeTributario, cstPIS, cstCOFINS string) (*repository.FederalTaxRule, error) {
	return nil, nil
}
func (m *mockRepo) GetICMSRule(ctx context.Context, orig, dest string) (*repository.ICMSRule, error) {
	return nil, nil
}
func (m *mockRepo) GetEquivalence(ctx context.Context, CSOSN string, tipoOperacao string) (*repository.TaxEquivalence, error) {
	return nil, nil
}
func (m *mockRepo) GetSimplesFaixa(ctx context.Context, anexo string, rbt12 decimal.Decimal) (*repository.SimplesFaixa, error) {
	return nil, nil
}
func (m *mockRepo) GetProductException(ctx context.Context, ncmFull, ncmGroup, ufDestino string, regimeTributarioDestino string) (*repository.ProductException, error) {
	return nil, nil
}
func (m *mockRepo) GetIPIRegra(ctx context.Context, NCM string, ExIPI string, CrtEmitente string, TipoOperacaoFiscal string, PerfilComprador string, UFDestino string, ZonaEspecial bool, DataOperacao string) (*repository.IPIRegra, error) {
	return nil, nil
}
func (m *mockRepo) GetIvaDualRule(ctx context.Context, ncm, ufDestino, municipioIBGE string) (*repository.IvaDualRule, error) {
	return m.ivaDualRule, m.ivaDualRuleErr
}
func (m *mockRepo) GetNCMSeletivo(ctx context.Context, ncm string) (*repository.NCMSeletivoRule, error) {
	return nil, nil
}
func (m *mockRepo) GetCSTReforma(ctx context.Context, flags repository.CSTFlags) (*repository.CSTReforma, error) {
	if m.cstReformaRuleErr != nil {
		return nil, m.cstReformaRuleErr
	}
	return m.cstReformaRule, nil
}

func inputPadrao(ncm, ufDestino, municipio string) models.DocumentoFiscalEntrada {
	return models.DocumentoFiscalEntrada{
		CorrelacaoID:       "test-001",
		DocumentoID:        "DOC-001",
		DataOperacao:       time.Now(),
		TipoOperacaoFiscal: "ENTRADA",
		CFOP:               "5102",
		CRTEmitente:        "3",
		LocalizacaoDestino: models.LocalizacaoFiscal{
			UF:        ufDestino,
			Municipio: municipio,
		},
		Itens: []models.ItemDocumentoFiscalEntrada{
			{
				SKU:           "SKU-001",
				NCM:           ncm,
				Quantidade:    decimal.NewFromInt(10),
				ValorUnitario: decimal.NewFromFloat(100.0),
			},
		},
	}
}

func TestReformaCalculator_CBS_IBS_AliquotasNormais(t *testing.T) {
	repo := &mockRepo{
		ivaDualRule: &repository.IvaDualRule{
			NCM:                  "84713019",
			UFDestino:            "SP",
			AliquotaCBS:          decimal.NewFromFloat(8.8),
			AliquotaIBSEstadual:  decimal.NewFromFloat(8.8),
			AliquotaIBSMunicipal: decimal.NewFromFloat(2.5),
			PercentualReducao:    decimal.Zero,
			IsImpostoSeletivo:    false,
			InicioValidade:       time.Now().AddDate(-1, 0, 0),
		},
	}

	calc := NewReformaCalculator(repo)
	input := inputPadrao("84713019", "SP", "3550308")
	itens, err := calc.Calculate(context.Background(), input)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if len(itens) == 0 || len(itens[0].Tributos) == 0 {
		t.Fatal("esperava tributos calculados, mas veio vazio")
	}

	baseEsperada := decimal.NewFromInt(1000)
	var valorCBS, valorIBS decimal.Decimal
	for _, trib := range itens[0].Tributos {
		if !trib.BaseCalculo.Equal(baseEsperada) {
			t.Errorf("BaseCalculo esperado %s, obtido %s para %s", baseEsperada, trib.BaseCalculo, trib.Tributo)
		}
		switch trib.Tributo {
		case "CBS":
			valorCBS = trib.Valor
		case "IBS":
			valorIBS = trib.Valor
		}
	}

	if !valorCBS.Equal(decimal.NewFromFloat(88.0)) {
		t.Errorf("CBS esperado 88, obtido %s", valorCBS)
	}
	if !valorIBS.Equal(decimal.NewFromFloat(113.0)) {
		t.Errorf("IBS esperado 113, obtido %s", valorIBS)
	}
}

func TestReformaCalculator_Reducao60(t *testing.T) {
	repo := &mockRepo{
		ivaDualRule: &repository.IvaDualRule{
			NCM:                  "84713019",
			UFDestino:            "SP",
			AliquotaCBS:          decimal.NewFromFloat(8.8),
			AliquotaIBSEstadual:  decimal.NewFromFloat(8.8),
			AliquotaIBSMunicipal: decimal.NewFromFloat(2.5),
			PercentualReducao:    decimal.NewFromFloat(60.0),
			IsImpostoSeletivo:    false,
			InicioValidade:       time.Now().AddDate(-1, 0, 0),
		},
	}

	calc := NewReformaCalculator(repo)
	input := inputPadrao("84713019", "SP", "3550308")
	itens, err := calc.Calculate(context.Background(), input)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if len(itens) == 0 || len(itens[0].Tributos) == 0 {
		t.Fatal("esperava tributos com reducao")
	}

	var valorCBS, valorIBS decimal.Decimal
	for _, trib := range itens[0].Tributos {
		switch trib.Tributo {
		case "CBS":
			valorCBS = trib.Valor
		case "IBS":
			valorIBS = trib.Valor
		}
	}

	if !valorCBS.Equal(decimal.NewFromFloat(35.2)) {
		t.Errorf("CBS com reducao 60%% esperado 35.2, obtido %s", valorCBS)
	}
	if !valorIBS.Equal(decimal.NewFromFloat(45.2)) {
		t.Errorf("IBS com reducao 60%% esperado 45.2, obtido %s", valorIBS)
	}
}

func TestReformaCalculator_Isento100(t *testing.T) {
	repo := &mockRepo{
		ivaDualRule: &repository.IvaDualRule{
			NCM:                  "84713019",
			UFDestino:            "SP",
			AliquotaCBS:          decimal.NewFromFloat(8.8),
			AliquotaIBSEstadual:  decimal.NewFromFloat(8.8),
			AliquotaIBSMunicipal: decimal.NewFromFloat(2.5),
			PercentualReducao:    decimal.NewFromFloat(100.0),
			IsImpostoSeletivo:    false,
			InicioValidade:       time.Now().AddDate(-1, 0, 0),
		},
	}

	calc := NewReformaCalculator(repo)
	input := inputPadrao("84713019", "SP", "3550308")
	itens, err := calc.Calculate(context.Background(), input)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	for _, trib := range itens[0].Tributos {
		if trib.Tributo == "CBS" || trib.Tributo == "IBS" {
			t.Errorf("Item isento nao deve gerar %s, mas gerou valor %s", trib.Tributo, trib.Valor)
		}
	}
}

// TestReformaCalculator_CBS_IBS_SemIS verifica que o ReformaCalculator
// calcula apenas CBS e IBS (IS foi extraído para ISFilter — F-006).
func TestReformaCalculator_CBS_IBS_SemIS(t *testing.T) {
	repo := &mockRepo{
		ivaDualRule: &repository.IvaDualRule{
			NCM:                  "24022000",
			UFDestino:            "SP",
			AliquotaCBS:          decimal.NewFromFloat(8.8),
			AliquotaIBSEstadual:  decimal.NewFromFloat(8.8),
			AliquotaIBSMunicipal: decimal.NewFromFloat(2.5),
			PercentualReducao:    decimal.Zero,
			IsImpostoSeletivo:    true,
			AliquotaIS:           decimal.NewFromFloat(50.0),
			InicioValidade:       time.Now().AddDate(-1, 0, 0),
		},
	}

	calc := NewReformaCalculator(repo)
	input := inputPadrao("24022000", "SP", "3550308")
	itens, err := calc.Calculate(context.Background(), input)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	temCBS := false
	temIBS := false
	temIS := false
	for _, trib := range itens[0].Tributos {
		switch trib.Tributo {
		case "CBS":
			temCBS = true
		case "IBS":
			temIBS = true
		case "IS":
			temIS = true
		}
	}

	if !temCBS {
		t.Error("faltou CBS")
	}
	if !temIBS {
		t.Error("faltou IBS")
	}
	if temIS {
		t.Error("IS nao deveria ser calculado pelo ReformaCalculator (movido para ISFilter Fase 0)")
	}
}

func TestReformaCalculator_RuleNotFound(t *testing.T) {
	repo := &mockRepo{
		ivaDualRule: nil,
	}

	calc := NewReformaCalculator(repo)
	input := inputPadrao("99999999", "SP", "3550308")
	itens, err := calc.Calculate(context.Background(), input)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if len(itens[0].Tributos) != 0 {
		t.Errorf("NCM sem regra deve retornar zero tributos, mas retornou %d", len(itens[0].Tributos))
	}
}

func TestReformaCalculator_MultiplosItens(t *testing.T) {
	repo := &mockRepo{
		ivaDualRule: &repository.IvaDualRule{
			NCM:                  "84713019",
			UFDestino:            "SP",
			AliquotaCBS:          decimal.NewFromFloat(8.8),
			AliquotaIBSEstadual:  decimal.NewFromFloat(8.8),
			AliquotaIBSMunicipal: decimal.NewFromFloat(2.5),
			PercentualReducao:    decimal.Zero,
			IsImpostoSeletivo:    false,
			InicioValidade:       time.Now().AddDate(-1, 0, 0),
		},
	}

	calc := NewReformaCalculator(repo)
	input := inputPadrao("84713019", "SP", "3550308")
	input.Itens = append(input.Itens, models.ItemDocumentoFiscalEntrada{
		SKU:           "SKU-002",
		NCM:           "84713019",
		Quantidade:    decimal.NewFromInt(5),
		ValorUnitario: decimal.NewFromFloat(200.0),
	})

	itens, err := calc.Calculate(context.Background(), input)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if len(itens) != 2 {
		t.Fatalf("esperava 2 itens, obteve %d", len(itens))
	}

	for i, it := range itens {
		if len(it.Tributos) == 0 {
			t.Errorf("item %d ficou sem tributos", i)
		}
		if it.SKU == "" {
			t.Errorf("item %d sem SKU", i)
		}
	}
}

func TestReformaCalculator_MunicipioEspecificoPrecedeEstadual(t *testing.T) {
	repo := &mockRepo{
		ivaDualRule: &repository.IvaDualRule{
			NCM:                  "84713019",
			UFDestino:            "SP",
			AliquotaCBS:          decimal.NewFromFloat(8.8),
			AliquotaIBSEstadual:  decimal.NewFromFloat(10.0),
			AliquotaIBSMunicipal: decimal.NewFromFloat(3.0),
			PercentualReducao:    decimal.Zero,
			IsImpostoSeletivo:    false,
			InicioValidade:       time.Now().AddDate(-1, 0, 0),
		},
	}

	calc := NewReformaCalculator(repo)
	input := inputPadrao("84713019", "SP", "3550308")
	itens, err := calc.Calculate(context.Background(), input)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	for _, trib := range itens[0].Tributos {
		if trib.Tributo == "IBS" {
			ibsTotalEsperado := decimal.NewFromFloat(130.0)
			if !trib.Valor.Equal(ibsTotalEsperado) {
				t.Errorf("IBS municipio especifico esperado %s, obtido %s", ibsTotalEsperado, trib.Valor)
			}
		}
	}
}
