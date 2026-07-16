// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/cmd/api/main.go
package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	// Importando sua Lib e Pacotes Internos
	"ms-billing-engine-tax-rates/internal/calculator"
	"ms-billing-engine-tax-rates/internal/ibsclient"
	"ms-billing-engine-tax-rates/internal/legacy"
	"ms-billing-engine-tax-rates/internal/middleware"
	"ms-billing-engine-tax-rates/internal/phase"
	"ms-billing-engine-tax-rates/internal/reforma"
	"ms-billing-engine-tax-rates/internal/admin"
	"ms-billing-engine-tax-rates/internal/credit"
	"ms-billing-engine-tax-rates/internal/simulation"
	"ms-billing-engine-tax-rates/internal/supplier"
	"ms-billing-engine-tax-rates/internal/token"
	"taxnexus-billing-core-lib/cache"
	"taxnexus-billing-core-lib/db"
	"taxnexus-billing-core-lib/models"
	"taxnexus-billing-core-lib/repository"
)

func main() {
	// 1. Configuração do Log Nativo (slog) em JSON conforme seu padrão
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(handler))

	slog.Info("Iniciando TaxNexus Billing Engine", "version", "1.0.0")

	// 2. Inicializar OpenTelemetry (W3C Trace Context)
	middleware.InitTracing()

	// 3. Inicializar Conexoes (Postgres e Redis)
	dbURL := os.Getenv("DATABASE_URL")
	redisAddr := os.Getenv("REDIS_ADDR")

	pgPool, err := db.ConnectPostgres(dbURL)
	if err != nil {
		slog.Error("Falha crítica no Postgres", "err", err)
		os.Exit(1)
	}
	rdb := cache.ConnectRedis(redisAddr)

	// 4. Camada de Dados: Repositorio com Cache
	taxRepo := repository.NewPostgresTaxRepository(pgPool)
	cachedRepo := repository.NewCachedTaxRepository(taxRepo, rdb)

	// 4.1 IBS Client com Circuit Breaker e Cache (F-007)
	ibsAPIBaseURL := os.Getenv("IBS_API_BASE_URL")
	if ibsAPIBaseURL == "" {
		ibsAPIBaseURL = "https://api.comitegestoribs.gov.br"
	}
	ibsHTTPClient := ibsclient.NewHTTPIBSClient(ibsAPIBaseURL)
	ibsCachedClient := ibsclient.NewCachedIBSClient(ibsHTTPClient, rdb)
	ibsCBClient := ibsclient.NewCircuitBreakerIBSClient(ibsCachedClient, rdb)
	_ = ibsclient.NewFallbackIBSClient(ibsCBClient, cachedRepo)

	slog.Info("IBS Client inicializado com circuit breaker e cache Redis (F-007)",
		"api_base_url", ibsAPIBaseURL,
		"cache_ttl", "24h",
		"circuit_breaker", "3 falhas em 60s → OPEN, HALF_OPEN apos 5min",
		"fallback", "GetIvaDualRule (banco de dados)",
	)

	// 5. Instanciar Calculadoras (Injecao de Dependencia)
	isFilter := legacy.NewISFilter(cachedRepo)
	ipiCalc := legacy.NewIPICalculator(cachedRepo)
	cbsCalc := reforma.NewCBSCalculator(cachedRepo)
	icmsCalc := legacy.NewICMSCalculator(cachedRepo)
	ibsCalc := reforma.NewIBSCalculator(cachedRepo)
	issCalc := legacy.NewISSCalculator()
	pisCofinsCalc := legacy.NewPISCofinsCalculator(cachedRepo)
	fustCalc := legacy.NewFUSTCalculator()
	funttelCalc := legacy.NewFUNTTELCalculator()

	// 6. Phase Resolver & Tax Selector (Reforma Tributaria — F-005)
	phaseResolver := phase.NewPhaseResolver()
	taxSelector := phase.NewTaxSelector(phaseResolver)

	// 7. Configurar a Billing Engine com pipeline SOP-013 (C-001)
	engine := calculator.BillingEnginePhased(
		calculator.Phase("IS", calculator.Sequential, isFilter),
		calculator.Phase("IPI", calculator.Sequential, ipiCalc),
		calculator.Phase("CBS", calculator.Sequential, cbsCalc),
		calculator.Phase("ICMS", calculator.Sequential, calculator.LegacyAdapter(icmsCalc)),
		calculator.Phase("IBS+ISS+PISCOFINS", calculator.Parallel,
			ibsCalc,
			issCalc,
			calculator.LegacyAdapter(pisCofinsCalc),
		),
		calculator.Phase("FUST", calculator.Sequential, fustCalc),
		calculator.Phase("FUNTTEL", calculator.Sequential, funttelCalc),
	)
	slog.Info("Engine pipeline SOP-013 (C-001) inicializado",
		"phases", 7,
		"pipeline", "IS→IPI→CBS→ICMS→(IBS+ISS+PISCOFINS)→FUST→FUNTTEL",
	)

	// 8. Configuracao do Servidor Fiber
	app := fiber.New(fiber.Config{
		AppName: "TaxNexus API v1",
	})

	app.Use(recover.New())
	app.Use(middleware.NewRequestIDMiddleware())
	app.Use(middleware.NewAuthMiddleware())

	// Rate Limiting — sliding window por IP + token (GAP-008)
	rateLimiter := middleware.NewRateLimiter()
	rateLimiter.StartCleanup(120 * time.Second)
	app.Use(rateLimiter.NewRateLimitMiddleware())

	app.Use(logger.New())

	metrics := middleware.NewMetricsCollector()
	app.Use(middleware.NewMetricsMiddleware(metrics))

	// 8.5 TokenService — Congelamento de Alíquotas (GAP-002, BR-06)
	tokenStore := token.NewMemoryTokenStore()
	tokenSvc := token.NewTokenService(tokenStore, cachedRepo)

	// 8.6 SimulationService — Simulação de Margem (GAP-003, BR-05)
	simSvc := simulation.NewSimulationService(engine)

	// 8.7 AdminTaxService — Gestão de Alíquotas (GAP-001, BR-02)
	adminRepo := admin.NewPostgresAdminRepository(pgPool)
	adminCache := admin.NewRedisCacheInvalidator(rdb)
	adminSvc := admin.NewAdminTaxService(adminRepo, adminCache)

	// 8.8 SupplierService — Qualificação Fiscal (GAP-007, BR-08)
	supplierStore := supplier.NewMemorySupplierStore()
	supplierSvc := supplier.NewValidationService(supplierStore)
	supplierChecker := supplier.NewSupplierCheckerAdapter(supplierSvc)

	// 8.9 CreditEngine — Créditos na Entrada (GAP-005, BR-08)
	creditEngine := credit.NewCreditEngine(engine, supplierChecker, cachedRepo)

	// 9. API Versioning — Grupo /v1/ (GAP-009)
	v1 := app.Group("/v1")

	// Health Check — liveness probe
	healthzHandler := func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	}
	v1.Get("/healthz", healthzHandler)
	app.Get("/healthz", withDeprecation(healthzHandler))

	// Health Check — readiness probe (PG + Redis)
	healthHandler := func(c *fiber.Ctx) error {
		healthy := true
		checks := make(map[string]string)

		if err := pgPool.Ping(c.Context()); err != nil {
			slog.Warn("postgres health check failed", "err", err)
			checks["postgres"] = "unhealthy"
			healthy = false
		} else {
			checks["postgres"] = "healthy"
		}

		if err := rdb.Ping(c.Context()).Err(); err != nil {
			slog.Warn("redis health check failed", "err", err)
			checks["redis"] = "unhealthy"
			healthy = false
		} else {
			checks["redis"] = "healthy"
		}

		status := fiber.StatusOK
		if !healthy {
			status = fiber.StatusServiceUnavailable
		}

		return c.Status(status).JSON(fiber.Map{
			"status": map[bool]string{true: "ok", false: "degraded"}[healthy],
			"checks": checks,
		})
	}
	v1.Get("/health", healthHandler)
	app.Get("/health", withDeprecation(healthHandler))

	// Métricas Prometheus — protegidas por role check (admin/monitoring)
	// Configurável via METRICS_REQUIRE_AUTH (default: true)
	metricsHandler := metrics.Handler()
	metricsGuard := func(c *fiber.Ctx) error {
		if os.Getenv("METRICS_REQUIRE_AUTH") != "false" {
			if !middleware.HasRole(c, "admin", "monitoring") {
				return c.Status(403).JSON(fiber.Map{
					"error": "acesso_negado",
					"message": "Métricas requerem role admin/monitoring. Configure METRICS_REQUIRE_AUTH=false para desabilitar.",
				})
			}
		}
		return metricsHandler(c)
	}
	v1.Get("/metrics", metricsGuard)
	app.Get("/metrics", metricsGuard)

	// 10. Rota de Calculo com Phase-Aware Processing (F-005) + TaxToken (GAP-002)
	calculateHandler := func(c *fiber.Ctx) error {
		start := time.Now()

		// Parse do JSON com suporte a token_id opcional (GAP-002)
		var payload struct {
			models.DocumentoFiscalEntrada
			TokenID string `json:"token_id"`
		}
		if err := c.BodyParser(&payload); err != nil {
			slog.Error("Erro ao processar JSON de entrada", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Payload inválido",
			})
		}
		input := payload.DocumentoFiscalEntrada

		// Token Fiscal: validar e injetar aliquotas congeladas (GAP-002)
		if payload.TokenID != "" {
			tok, err := tokenSvc.Validate(c.Context(), payload.TokenID)
			if err != nil {
				if _, ok := err.(token.ErrTokenExpired); ok {
					return c.Status(fiber.StatusConflict).JSON(fiber.Map{
						"error":   "token_expirado",
						"message": "token expirado — renegocie",
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   "token_invalido",
					"message": err.Error(),
				})
			}
			for i := range input.Itens {
				input.Itens[i].AddDetalhe("TOKEN_CBS_ALIQUOTA", tok.AliquotaCBS)
				input.Itens[i].AddDetalhe("TOKEN_IBS_ESTADUAL_ALIQUOTA", tok.AliquotaIBSEstadual)
				input.Itens[i].AddDetalhe("TOKEN_IBS_MUNICIPAL_ALIQUOTA", tok.AliquotaIBSMunicipal)
				input.Itens[i].AddDetalhe("TOKEN_IS_ALIQUOTA", tok.AliquotaIS)
			}
			slog.Info("Token fiscal aplicado ao cálculo",
				"token_id", tok.ID, "ncm", tok.NCM,
				"expires_at", tok.ExpiresAt,
			)
		}

		filter := taxSelector.Filter(input.DataOperacao)

		slog.Info("Requisição recebida",
			"DocumentoID", input.DocumentoID,
			"data_operacao", input.DataOperacao,
			"phase", filter.Phase,
			"itens_count", len(input.Itens),
			"token_id", payload.TokenID,
			"trace_id", middleware.TraceID(c),
			"request_id", middleware.SpanID(c),
		)

		response, err := engine.ProcessWithPhase(c.Context(), input, filter)
		if err != nil {
			slog.Error("Erro durante o cálculo tributário", "DocumentoID", input.DocumentoID, "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Falha interna no processamento dos impostos",
			})
		}

		latency := time.Since(start)
		slog.Info("Cálculo finalizado",
			"IDTransaction", response.IDTransaction,
			"DocumentoID", input.DocumentoID,
			"phase", filter.Phase,
			"total_nota", response.TotalNota,
			"total_impostos", response.TotalImpostos,
			"tempo_ms", latency.Milliseconds(),
			"trace_id", middleware.TraceID(c),
			"request_id", middleware.SpanID(c),
		)

		return c.Status(fiber.StatusOK).JSON(response)
	}
	v1.Post("/calculate", calculateHandler)
	app.Post("/calculate", withDeprecation(calculateHandler))

	// 11. Token Fiscal (GAP-002 — BR-06 Garantia de Preço Ofertado)
	v1.Post("/token/generate", func(c *fiber.Ctx) error {
		var req struct {
			NCM           string `json:"ncm"`
			UFOrigem      string `json:"uf_origem"`
			UFDestino     string `json:"uf_destino"`
			MunicipioIBGE string `json:"municipio_ibge"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Payload inválido"})
		}
		if req.NCM == "" || req.UFDestino == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ncm e uf_destino são obrigatórios"})
		}
		tok, err := tokenSvc.Generate(c.Context(), token.TokenGenerateRequest{
			NCM: req.NCM, UFOrigem: req.UFOrigem, UFDestino: req.UFDestino, MunicipioIBGE: req.MunicipioIBGE,
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "falha_ao_gerar_token", "message": err.Error()})
		}
		return c.Status(201).JSON(tok)
	})

	// POST /v1/simulate — Simulação de margem (GAP-003, BR-05)
	v1.Post("/simulate", func(c *fiber.Ctx) error {
		var req simulation.SimulationRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Payload inválido"})
		}
		resp, err := simSvc.Simulate(c.Context(), req)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "simulacao_invalida", "message": err.Error()})
		}
		return c.JSON(resp)
	})

	// Admin: PUT /v1/admin/tax-rates/iva-dual (GAP-001, BR-02)
	v1.Put("/admin/tax-rates/iva-dual", func(c *fiber.Ctx) error {
		if !middleware.HasRole(c, "admin", "fiscal") {
			return c.Status(403).JSON(fiber.Map{"error": "acesso_negado", "message": "Permissão insuficiente"})
		}
		changedBy := middleware.GetUserName(c)
		if changedBy == "" {
			changedBy = middleware.GetUserID(c)
		}
		if changedBy == "" {
			slog.Warn("Admin endpoint acessado sem identidade no JWT")
			return c.Status(401).JSON(fiber.Map{"error": "nao_autenticado", "message": "Identidade não encontrada no token"})
		}
		var input admin.IvaDualRuleInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Payload inválido"})
		}
		out, err := adminSvc.UpsertRule(c.Context(), input, changedBy)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "validacao_falhou", "message": err.Error()})
		}
		return c.Status(201).JSON(out)
	})

	// Admin: GET /v1/admin/tax-rates/iva-dual (GAP-001, BR-02)
	v1.Get("/admin/tax-rates/iva-dual", func(c *fiber.Ctx) error {
		if !middleware.HasRole(c, "admin", "fiscal") {
			return c.Status(403).JSON(fiber.Map{"error": "acesso_negado", "message": "Permissão insuficiente"})
		}
		filter := admin.ListRulesFilter{
			NCM:          c.Query("ncm"),
			UF:           c.Query("uf"),
			AtivasApenas: c.Query("ativas_apenas", "true") == "true",
		}
		rules, err := adminSvc.ListRules(c.Context(), filter)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "erro_consulta", "message": err.Error()})
		}
		return c.JSON(rules)
	})

	// Supplier: POST /v1/supplier/validate (GAP-007, BR-08)
	v1.Post("/supplier/validate", func(c *fiber.Ctx) error {
		if !middleware.HasRole(c, "admin", "fiscal") {
			return c.Status(403).JSON(fiber.Map{"error": "acesso_negado"})
		}
		var req supplier.ValidateRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Payload inválido"})
		}
		sup, err := supplierSvc.Validate(c.Context(), req)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "validacao_fornecedor", "message": err.Error()})
		}
		return c.Status(201).JSON(sup)
	})

	// Supplier: GET /v1/supplier/:cnpj (GAP-007)
	v1.Get("/supplier/:cnpj", func(c *fiber.Ctx) error {
		if !middleware.HasRole(c, "admin", "fiscal", "credit") {
			return c.Status(403).JSON(fiber.Map{"error": "acesso_negado"})
		}
		cnpj := c.Params("cnpj")
		sup, err := supplierSvc.GetByCNPJ(c.Context(), cnpj)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "fornecedor_nao_encontrado", "message": err.Error()})
		}
		return c.JSON(sup)
	})

	// Supplier: PUT /v1/supplier/:cnpj (GAP-007)
	v1.Put("/supplier/:cnpj", func(c *fiber.Ctx) error {
		if !middleware.HasRole(c, "admin", "fiscal") {
			return c.Status(403).JSON(fiber.Map{"error": "acesso_negado"})
		}
		cnpj := c.Params("cnpj")
		var req supplier.ValidateRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Payload inválido"})
		}
		req.CNPJ = cnpj
		sup, err := supplierSvc.Update(c.Context(), cnpj, req)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "fornecedor_nao_encontrado", "message": err.Error()})
		}
		return c.JSON(sup)
	})

	// Credit: POST /v1/credit/calculate (GAP-005, BR-08)
	v1.Post("/credit/calculate", func(c *fiber.Ctx) error {
		if !middleware.HasRole(c, "admin", "fiscal", "credit") {
			return c.Status(403).JSON(fiber.Map{"error": "acesso_negado"})
		}
		var req credit.CreditCalculationRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Payload inválido"})
		}
		resp, err := creditEngine.Calculate(c.Context(), req)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "falha_credito", "message": err.Error()})
		}
		return c.JSON(resp)
	})

	// Credit: GET /v1/credit/summary (GAP-005, BR-08)
	v1.Get("/credit/summary", func(c *fiber.Ctx) error {
		if !middleware.HasRole(c, "admin", "fiscal") {
			return c.Status(403).JSON(fiber.Map{"error": "acesso_negado"})
		}
		periodo := c.Query("periodo", "")
		resp, err := creditEngine.Summary(c.Context(), periodo)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "falha_summary", "message": err.Error()})
		}
		return c.JSON(resp)
	})

	v1.Get("/token/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		status, err := tokenSvc.Status(c.Context(), id)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "token_nao_encontrado", "message": err.Error()})
		}
		return c.JSON(status)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	if port[0] != ':' {
		port = ":" + port
	}
	slog.Info("Servidor iniciando", "porta", port)
	if err := app.Listen(port); err != nil {
		slog.Error("Erro ao iniciar servidor", "error", err)
		os.Exit(1)
	}
}

// withDeprecation adiciona headers de deprecação à rota legada (GAP-009).
func withDeprecation(handler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Deprecation", "true")
		c.Set("Sunset", "Sat, 01 Jan 2028 00:00:00 GMT")
		c.Set("Link", "</v1"+c.Path()+">; rel=\"successor-version\"")
		return handler(c)
	}
}
