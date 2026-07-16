// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/simulation/service.go
package simulation

import (
	"context"
	"fmt"
	"log/slog"

	"ms-billing-engine-tax-rates/internal/phase"
	"taxnexus-billing-core-lib/models"

	"github.com/shopspring/decimal"
)

// SimulationRequest representa a requisição de simulação de margem (GAP-003).
// Aceita o mesmo payload do /calculate, acrescido de custo por item e
// múltiplos destinos para comparação de cenários pelo time comercial.
type SimulationRequest struct {
	// Itens da simulação (SKU, NCM, valor, quantidade, custo opcional)
	Itens []SimulationItem `json:"itens"`

	// UF de origem (default: SP)
	UFOrigem string `json:"uf_origem"`

	// Destinos para comparação (múltiplas UFs)
	Destinos []Destino `json:"destinos"`
}

// SimulationItem representa um item na simulação com custo associado.
type SimulationItem struct {
	SKU           string          `json:"sku"`
	NCM           string          `json:"ncm"`
	Quantidade    decimal.Decimal `json:"quantidade"`
	ValorUnitario decimal.Decimal `json:"valor_unitario"`
	Custo         decimal.Decimal `json:"custo"` // custo de aquisição/produção para margem
}

// Destino representa uma UF/município para simulação comparativa.
type Destino struct {
	UF            string `json:"uf"`
	MunicipioIBGE string `json:"municipio_ibge,omitempty"`
}

// SimulationResponse representa o resultado da simulação.
type SimulationResponse struct {
	MargemLiquida  decimal.Decimal `json:"margem_liquida"`
	AlertaMargem   bool            `json:"alerta_margem"`
	ImpactoPorUF   []ImpactoUF     `json:"impacto_por_uf"`
}

// ImpactoUF representa o impacto dos tributos para um destino específico.
type ImpactoUF struct {
	UF              string          `json:"uf"`
	MunicipioIBGE   string          `json:"municipio_ibge,omitempty"`
	ValorTotal      decimal.Decimal `json:"valor_total"`
	ValorLiquido    decimal.Decimal `json:"valor_liquido"`
	TotalImpostos   decimal.Decimal `json:"total_impostos"`
	MargemLiquida   decimal.Decimal `json:"margem_liquida"`
	AlertaMargem    bool            `json:"alerta_margem"`
	Warning         string          `json:"warning,omitempty"`
}

// TaxEngine é a interface que o SimulationService usa para calcular tributos.
// Implementada por calculator.BillingEngineStruct.ProcessWithPhase.
type TaxEngine interface {
	ProcessWithPhase(ctx context.Context, input models.DocumentoFiscalEntrada, filter phase.CalculatorFilter) (models.DocumentoFiscalSaida, error)
}

// SimulationService executa simulações de margem para o time comercial (BR-05).
type SimulationService struct {
	engine TaxEngine
}

// NewSimulationService cria um novo SimulationService.
func NewSimulationService(engine TaxEngine) *SimulationService {
	return &SimulationService{engine: engine}
}

// Simulate executa a simulação de margem para um ou mais destinos.
//
// Fluxo:
//  1. Para cada destino, monta um DocumentoFiscalEntrada
//  2. Chama o motor de cálculo (ProcessWithPhase) com fase PRE_REFORMA
//  3. Calcula margem_liquida = (valor_liquido − custo) / valor_liquido
//  4. Agrega resultados por UF
//
// Se o destino não tiver alíquota IBS cadastrada, o campo warning é preenchido
// e o IBS não é calculado (comportamento normal do engine).
func (s *SimulationService) Simulate(ctx context.Context, req SimulationRequest) (*SimulationResponse, error) {
	if len(req.Itens) == 0 {
		return nil, fmt.Errorf("pelo menos um item é obrigatório")
	}
	if len(req.Destinos) == 0 {
		return nil, fmt.Errorf("pelo menos um destino é obrigatório")
	}
	if req.UFOrigem == "" {
		req.UFOrigem = "SP"
	}

	defaultFilter := phase.CalculatorFilter{
		Phase:           "PRE_REFORMA",
		IPIActive:       true,
		PISCOFINSActive: true,
		ICMSActive:      true,
		ISSActive:       true,
		CBSActive:       true,
		IBSActive:       true,
		ISActive:        true,
	}

	resp := &SimulationResponse{
		ImpactoPorUF: make([]ImpactoUF, 0, len(req.Destinos)),
	}

	totalValor := decimal.Zero
	totalCusto := decimal.Zero
	totalLiquido := decimal.Zero

	for _, destino := range req.Destinos {
		// Monta input para este destino
		input := models.DocumentoFiscalEntrada{
			CorrelacaoID:       "sim",
			DocumentoID:        fmt.Sprintf("sim-%s-%s", destino.UF, destino.MunicipioIBGE),
			TipoOperacaoFiscal: "SAIDA",
			CRTEmitente:        "3",
			LocalizacaoOrigem:  models.LocalizacaoFiscal{UF: req.UFOrigem},
			LocalizacaoDestino: models.LocalizacaoFiscal{UF: destino.UF, Municipio: destino.MunicipioIBGE},
			Itens:              make([]models.ItemDocumentoFiscalEntrada, len(req.Itens)),
		}

		valorDestino := decimal.Zero
		custoDestino := decimal.Zero
		for i, item := range req.Itens {
			input.Itens[i] = models.ItemDocumentoFiscalEntrada{
				SKU:           item.SKU,
				NCM:           item.NCM,
				Quantidade:    item.Quantidade,
				ValorUnitario: item.ValorUnitario,
			}
			valorDestino = valorDestino.Add(item.ValorUnitario.Mul(item.Quantidade))
			custoDestino = custoDestino.Add(item.Custo)
		}

		// Executa cálculo
		output, err := s.engine.ProcessWithPhase(ctx, input, defaultFilter)
		if err != nil {
			slog.Warn("Erro no cálculo da simulação",
				"uf", destino.UF,
				"error", err,
			)
			continue
		}

		// Extrai valor_liquido e total_impostos
		liquidoDestino := decimal.Zero
		impostosDestino := output.TotalImpostos
		if len(output.Itens) > 0 {
			liquidoDestino = output.Itens[0].ValorLiquido
		}

		// Calcula margem: (valor_liquido − custo) / valor_liquido
		margem := decimal.Zero
		if liquidoDestino.GreaterThan(decimal.Zero) {
			margem = liquidoDestino.Sub(custoDestino).Div(liquidoDestino).Mul(decimal.NewFromInt(100))
		}

		alerta := margem.LessThan(decimal.Zero)
		warning := ""
		if impostosDestino.IsZero() && destino.MunicipioIBGE != "" {
			warning = "aliquota_ibs_nao_encontrada"
		}

		impacto := ImpactoUF{
			UF:            destino.UF,
			MunicipioIBGE: destino.MunicipioIBGE,
			ValorTotal:    valorDestino.Round(2),
			ValorLiquido:  liquidoDestino.Round(2),
			TotalImpostos: impostosDestino.Round(2),
			MargemLiquida: margem.Round(2),
			AlertaMargem:  alerta,
			Warning:       warning,
		}
		resp.ImpactoPorUF = append(resp.ImpactoPorUF, impacto)

		totalValor = totalValor.Add(valorDestino)
		totalCusto = totalCusto.Add(custoDestino)
		totalLiquido = totalLiquido.Add(liquidoDestino)
	}

	// Margem líquida consolidada
	if totalLiquido.GreaterThan(decimal.Zero) {
		resp.MargemLiquida = totalLiquido.Sub(totalCusto).Div(totalLiquido).Mul(decimal.NewFromInt(100)).Round(2)
	}
	resp.AlertaMargem = resp.MargemLiquida.LessThan(decimal.Zero)

	slog.Info("Simulação concluída",
		"destinos", len(resp.ImpactoPorUF),
		"margem_liquida", resp.MargemLiquida,
	)

	return resp, nil
}
