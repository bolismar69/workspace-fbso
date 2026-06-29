// Package ibsclient implementa o cliente de consulta de aliquotas IBS
// junto a API do Comite Gestor do IBS, com cache Redis e circuit breaker
// conforme especificado em:
//
//	PROCEDURE-FIN-00001 SOP-002, SOP-014
//	RULES-CATALOG-FIN-00001 BR-TAX-ACT-001, BR-TAX-CONS-009
//
// Arquitetura:
//
//	IBSRateFetcher (interface)
//	  ├── HTTPIBSClient        — chamada HTTP a API do Comite Gestor
//	  ├── CachedIBSClient      — decorator com cache Redis (TTL 24h)
//	  ├── CircuitBreakerIBSClient — decorator com circuit breaker
//	  └── FallbackIBSClient    — tenta API, fallback DB (GetIvaDualRule)
//
// Enquanto a API do Comite Gestor nao e publicada (Gap G2 — LC 214/2025),
// o FallbackIBSClient utiliza GetIvaDualRule() como fonte de aliquotas IBS.
//
// Thread-safe: o cache Redis ja e thread-safe; o circuit breaker tambem.
package ibsclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"ms-billing-engine-tax-rates/internal/circuitbreaker"
	"taxnexus-billing-core-lib/repository"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

// IBSRate representa a aliquota IBS para um municipio, retornada pela API
// do Comite Gestor ou pelo fallback de banco de dados.
type IBSRate struct {
	IBGECode        string          `json:"ibge_code"`
	AliquotaEstadual decimal.Decimal `json:"aliquota_estadual"`
	AliquotaMunicipal decimal.Decimal `json:"aliquota_municipal"`
	VigenciaInicio   string          `json:"vigencia_inicio"`
	Fonte            string          `json:"fonte"` // CACHE, API_COMITE_GESTOR, FALLBACK_CIRCUIT_OPEN, FALLBACK_DB
}

// IBSRateFetcher define a interface para obtencao de aliquotas IBS.
// Permite trocar a fonte de dados (HTTP, cache, DB fallback)
// sem alterar as calculadoras consumidoras.
type IBSRateFetcher interface {
	// FetchRate obtem a aliquota IBS para um codigo IBGE de municipio.
	// Retorna nil se nao houver aliquota configurada.
	FetchRate(ctx context.Context, ibgeCode string) (*IBSRate, error)
}

// HTTPIBSClient implementa IBSRateFetcher via chamada HTTP a API do
// Comite Gestor do IBS.
//
// Endpoint: GET {BaseURL}/api/v1/rates?ibge_code={code}
//
// NOTA: Enquanto a API do Comite Gestor nao e publicada (Gap G2),
// esta implementacao nao pode ser usada em producao.
// Use FallbackIBSClient para obter aliquotas do banco de dados.
type HTTPIBSClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewHTTPIBSClient cria um cliente HTTP para a API do Comite Gestor IBS.
// baseURL: URL base da API (ex: "https://api.comitegestoribs.gov.br")
func NewHTTPIBSClient(baseURL string) *HTTPIBSClient {
	return &HTTPIBSClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// IBSRateResponse e a estrutura de resposta da API do Comite Gestor.
type IBSRateResponse struct {
	IBGECode         string  `json:"ibge_code"`
	AliquotaEstadual float64 `json:"aliquota_estadual"`
	AliquotaMunicipal float64 `json:"aliquota_municipal"`
	VigenciaInicio   string  `json:"vigencia_inicio"`
}

// FetchRate implementa IBSRateFetcher via HTTP.
func (c *HTTPIBSClient) FetchRate(ctx context.Context, ibgeCode string) (*IBSRate, error) {
	url := fmt.Sprintf("%s/api/v1/rates?ibge_code=%s", c.baseURL, ibgeCode)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar request IBS: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro na chamada HTTP IBS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API IBS retornou status %d: %s", resp.StatusCode, string(body))
	}

	var apiResp IBSRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta IBS: %w", err)
	}

	return &IBSRate{
		IBGECode:         apiResp.IBGECode,
		AliquotaEstadual: decimal.NewFromFloat(apiResp.AliquotaEstadual),
		AliquotaMunicipal: decimal.NewFromFloat(apiResp.AliquotaMunicipal),
		VigenciaInicio:   apiResp.VigenciaInicio,
		Fonte:            "API_COMITE_GESTOR",
	}, nil
}

// CachedIBSClient e um decorator que adiciona cache Redis ao IBSRateFetcher.
//
// Chave de cache: ibs:rate:{ibge_code}:{date}
// TTL: 24 horas (aliquotas IBS mudam no maximo diariamente)
//
// Ref: BR-TAX-CONS-009
type CachedIBSClient struct {
	inner IBSRateFetcher
	rdb   *redis.Client
	ttl   time.Duration
}

// NewCachedIBSClient cria um cliente IBS com cache Redis.
func NewCachedIBSClient(inner IBSRateFetcher, rdb *redis.Client) *CachedIBSClient {
	return &CachedIBSClient{
		inner: inner,
		rdb:   rdb,
		ttl:   24 * time.Hour,
	}
}

// FetchRate implementa IBSRateFetcher com cache Redis.
func (c *CachedIBSClient) FetchRate(ctx context.Context, ibgeCode string) (*IBSRate, error) {
	today := time.Now().Format("2006-01-02")
	cacheKey := fmt.Sprintf("ibs:rate:%s:%s", ibgeCode, today)

	// Tenta cache primeiro
	if val, err := c.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var rate IBSRate
		if err := json.Unmarshal([]byte(val), &rate); err == nil {
			rate.Fonte = "CACHE"
			slog.Debug("IBS Client: cache hit",
				"ibge_code", ibgeCode,
				"cache_key", cacheKey,
			)
			return &rate, nil
		}
	}

	// Cache miss — busca na fonte original
	rate, err := c.inner.FetchRate(ctx, ibgeCode)
	if err != nil {
		return nil, err
	}

	if rate != nil {
		data, _ := json.Marshal(rate)
		c.rdb.Set(ctx, cacheKey, data, c.ttl)

		slog.Debug("IBS Client: cache miss, armazenado",
			"ibge_code", ibgeCode,
			"cache_key", cacheKey,
			"fonte", rate.Fonte,
		)
	}

	return rate, nil
}

// CircuitBreakerIBSClient e um decorator que adiciona circuit breaker
// ao IBSRateFetcher, protegendo contra falhas na API do Comite Gestor.
//
// Comportamento:
//   - CLOSED: chamadas passam normalmente
//   - OPEN (≥3 falhas em 60s): usa cache expirado como fallback
//   - HALF_OPEN (apos 5min): permite uma chamada de teste
//
// Ref: BR-TAX-ACT-001, SOP-014
type CircuitBreakerIBSClient struct {
	inner IBSRateFetcher
	cb    *circuitbreaker.CircuitBreaker
	rdb   *redis.Client
}

// NewCircuitBreakerIBSClient cria um cliente IBS com circuit breaker.
func NewCircuitBreakerIBSClient(inner IBSRateFetcher, rdb *redis.Client) *CircuitBreakerIBSClient {
	return &CircuitBreakerIBSClient{
		inner: inner,
		cb:    circuitbreaker.New("ibs-api", circuitbreaker.DefaultConfig()),
		rdb:   rdb,
	}
}

// FetchRate implementa IBSRateFetcher com circuit breaker.
// Quando o circuito esta OPEN, tenta usar cache Redis expirado como fallback.
func (c *CircuitBreakerIBSClient) FetchRate(ctx context.Context, ibgeCode string) (*IBSRate, error) {
	result, err := c.cb.Execute(ctx, func() (interface{}, error) {
		rate, err := c.inner.FetchRate(ctx, ibgeCode)
		if err != nil {
			return nil, err
		}
		return rate, nil
	})

	if err != nil {
		// Circuito OPEN — tenta cache expirado como fallback
		if _, isCircuitOpen := err.(circuitbreaker.ErrCircuitOpen); isCircuitOpen {
			slog.Warn("IBS Client: circuit breaker OPEN, tentando cache expirado",
				"ibge_code", ibgeCode,
			)

			if cachedRate := c.getExpiredCache(ctx, ibgeCode); cachedRate != nil {
				cachedRate.Fonte = "FALLBACK_CIRCUIT_OPEN"
				return cachedRate, nil
			}

			return nil, fmt.Errorf("IBS indisponivel: circuito OPEN e sem cache fallback para %s", ibgeCode)
		}
		return nil, err
	}

	rate, ok := result.(*IBSRate)
	if !ok || rate == nil {
		return nil, nil
	}

	return rate, nil
}

// getExpiredCache busca no Redis mesmo que a chave tenha expirado.
// Estrategia de fallback quando a API externa esta indisponivel.
// Tenta varias chaves antigas (ultimos 7 dias).
func (c *CircuitBreakerIBSClient) getExpiredCache(ctx context.Context, ibgeCode string) *IBSRate {
	// Tenta os ultimos 7 dias de cache
	for daysBack := 1; daysBack <= 7; daysBack++ {
		date := time.Now().AddDate(0, 0, -daysBack).Format("2006-01-02")
		cacheKey := fmt.Sprintf("ibs:rate:%s:%s", ibgeCode, date)

		val, err := c.rdb.Get(ctx, cacheKey).Result()
		if err != nil {
			continue
		}

		var rate IBSRate
		if err := json.Unmarshal([]byte(val), &rate); err != nil {
			continue
		}

		slog.Info("IBS Client: usando cache expirado como fallback",
			"ibge_code", ibgeCode,
			"cache_date", date,
		)
		return &rate
	}

	return nil
}

// FallbackIBSClient implementa IBSRateFetcher com fallback para o banco de dados.
//
// Ordem de precedencia:
//  1. API do Comite Gestor (via cached + circuit breaker)
//  2. Banco de dados (GetIvaDualRule) como fallback
//
// Este e o cliente recomendado para uso em producao enquanto a API do
// Comite Gestor nao esta publicada (Gap G2).
type FallbackIBSClient struct {
	primary IBSRateFetcher // cached + circuit breaker → API
	repo    repository.TaxRepository
}

// NewFallbackIBSClient cria um cliente IBS com fallback para banco de dados.
func NewFallbackIBSClient(primary IBSRateFetcher, repo repository.TaxRepository) *FallbackIBSClient {
	return &FallbackIBSClient{
		primary: primary,
		repo:    repo,
	}
}

// FetchRate implementa IBSRateFetcher com fallback DB.
func (c *FallbackIBSClient) FetchRate(ctx context.Context, ibgeCode string) (*IBSRate, error) {
	// Tenta API primeiro
	rate, err := c.primary.FetchRate(ctx, ibgeCode)
	if err == nil && rate != nil {
		return rate, nil
	}

	if err != nil {
		slog.Warn("IBS Client: API indisponivel, usando fallback banco de dados",
			"ibge_code", ibgeCode,
			"error", err,
		)
	}

	// Fallback: busca do banco de dados via IvaDualRule
	// Usa um NCM coringa e UF destino para obter aliquotas IBS do municipio
	rule, dbErr := c.repo.GetIvaDualRule(ctx, "*", "", ibgeCode)
	if dbErr != nil || rule == nil {
		return nil, fmt.Errorf("IBS indisponivel: API falhou e fallback DB sem resultado para %s", ibgeCode)
	}

	return &IBSRate{
		IBGECode:         ibgeCode,
		AliquotaEstadual: rule.AliquotaIBSEstadual,
		AliquotaMunicipal: rule.AliquotaIBSMunicipal,
		Fonte:            "FALLBACK_DB",
	}, nil
}
