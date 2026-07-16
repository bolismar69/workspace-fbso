// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/cmd/test_engine/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	// "taxnexus-billing-core-lib/cache"
	// "taxnexus-billing-core-lib/db"
	"taxnexus-billing-core-lib/cache"
	"taxnexus-billing-core-lib/db"
	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"

	"ms-billing-engine-tax-rates/internal/calculator"
	"ms-billing-engine-tax-rates/internal/domain"
	"ms-billing-engine-tax-rates/internal/legacy"

	"github.com/shopspring/decimal"
)

func main() {
	// 1. Configuração do Logger (Nível Debug para ver a injeção de valores)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	// 2. Instanciar Calculadoras Legais
	// Para este teste, vamos usar o repositório real, mas poderia ser um mock
	// // 	// repo := repository.NewMockTaxRepository()
	dbURL := os.Getenv("DATABASE_URL")
	redisAddr := os.Getenv("REDIS_ADDR")

	// 3. Inicializar Conexões
	// Ajuste a string de conexão conforme seu ambiente local
	pgPool, err := db.ConnectPostgres(dbURL)
	if err != nil {
		slog.Error("Erro ao conectar no Postgres", "err", err)
		os.Exit(1)
	}

	// 4. Inicializar Cache Redis
	rdb := cache.ConnectRedis(redisAddr)

	// 5. Montar Repositório com Cache
	pgRepo := repository.NewPostgresTaxRepository(pgPool)
	cachedRepo := repository.NewCachedTaxRepository(pgRepo, rdb)

	// 6. Inicialização das Calculadoras Legais
	ipiCalc := legacy.NewIPICalculator(cachedRepo)
	icmsCalc := legacy.NewICMSCalculator(cachedRepo)
	pisCofinsCalc := legacy.NewPISCofinsCalculator(cachedRepo)

	// 7. Montagem do Engine com ordem garantida:
	//   Fase 1 (pré, sequencial): IPI
	//   Fase 2 (paralela):        ICMS e PIS/COFINS (já encontram ITEM_IPI_VALOR nos detalhes)
	engine := calculator.BillingEngineOrdered(
		[]domain.TaxCalculator{ipiCalc},
		calculator.LegacyAdapter(icmsCalc),
		calculator.LegacyAdapter(pisCofinsCalc, icmsCalc),
	)

	// 8. Payload de Teste: Venda Interestadual (RJ -> SP)
	input := models.DocumentoFiscalEntrada{
		CorrelacaoID:       "12345",
		DocumentoID:        "NF-12345",
		DataOperacao:       time.Now(),
		TipoOperacaoFiscal: string(models.TipoOperacaoFiscalSaida),
		CRTEmitente:        string(models.CRTEmitenteLucroReal),
		LocalizacaoOrigem:  models.LocalizacaoFiscal{UF: "RJ"},
		LocalizacaoDestino: models.LocalizacaoFiscal{UF: "SP"},
		Itens: []models.ItemDocumentoFiscalEntrada{
			{
				SKU:           "PROD-001",
				ValorUnitario: decimal.NewFromFloat(100000.00),
				NCM:           "40111000",
				Quantidade:    decimal.NewFromInt(1),
				DetalhesItemDocumentoFiscal: []models.Detalhe{
					{Key: string(models.KeyDocumentoInfosItemIPIAliquota), Value: "10.00"}, // 10% de IPI (ad valorem)
					{Key: string(models.KeyDocumentoInfosItemIPICST), Value: "50"},
					{Key: string(models.KeyDocumentoInfosItemIPICEnq), Value: "999"},
					{Key: string(models.KeyDocumentoInfosItemSubstituirCSTICMS), Value: "000"},
					{Key: string(models.KeyDocumentoInfosItemSubstituirCSTPIS), Value: string(models.CSTPISCOFINS01)},
					{Key: string(models.KeyDocumentoInfosItemSubstituirCSTCOFINS), Value: string(models.CSTPISCOFINS01)},
					{Key: string(models.KeyDocumentoInfosValorFrete), Value: decimal.NewFromFloat(50.00)},
					{Key: string(models.KeyDocumentoInfosValorExclusaoICMS), Value: decimal.NewFromFloat(5000.00)},        // Valor alto para testar a exclusão na base do PIS/COFINS
					{Key: string(models.KeyDocumentoInfosItemSubstituirMVAPercentual), Value: decimal.NewFromFloat(40.0)}, // Para testar o cálculo do ICMS-ST
				},
			},
		},
	}

	// 9. Processamento
	ctx := context.Background()
	response, err := engine.Process(ctx, input)
	if err != nil {
		slog.Error("Erro no processamento do engine", "error", err)
		os.Exit(1)
	}

	// 10. Exibir Resultado
	slog.Info("Cálculo concluído", "response", response)

	// 11. Output do Resultado somente em JSON
	jsonResult, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		slog.Error("Erro ao serializar response para JSON", "err", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonResult))
}
