// path: backend/go/libs/go-native/taxnexus-billing-core-lib/repository/cached_tax_repository.go
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type cachedTaxRepository struct {
	realRepo TaxRepository
	rdb      *redis.Client
	ttl      time.Duration
}

func NewCachedTaxRepository(real TaxRepository, rdb *redis.Client) TaxRepository {
	return &cachedTaxRepository{
		realRepo: real,
		rdb:      rdb,
		ttl:      24 * time.Hour, // Alíquotas mudam pouco, 24h é seguro
	}
}

func (r *cachedTaxRepository) GetFederalTaxRule(ctx context.Context, regimeTributario, cstPIS, cstCOFINS string) (*FederalTaxRule, error) {
	cacheKey := fmt.Sprintf("tax:federal:%s:%s:%s", regimeTributario, cstPIS, cstCOFINS)

	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var rule FederalTaxRule
		if err := json.Unmarshal([]byte(val), &rule); err == nil {
			return &rule, nil
		}
	}

	rule, err := r.realRepo.GetFederalTaxRule(ctx, regimeTributario, cstPIS, cstCOFINS)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(rule)
	r.rdb.Set(ctx, cacheKey, data, r.ttl)

	return rule, nil
}

// GetICMSRule com Cache
func (r *cachedTaxRepository) GetICMSRule(ctx context.Context, orig, dest string) (*ICMSRule, error) {
	cacheKey := fmt.Sprintf("tax:icms:%s:%s", orig, dest)

	// Tenta Redis
	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var rule ICMSRule
		if err := json.Unmarshal([]byte(val), &rule); err == nil {
			return &rule, nil
		}
	}

	// Se falhar ou não existir, vai no Postgres
	rule, err := r.realRepo.GetICMSRule(ctx, orig, dest)
	if err != nil {
		return nil, err
	}

	// Salva no Redis para a próxima
	data, _ := json.Marshal(rule)
	r.rdb.Set(ctx, cacheKey, data, r.ttl)

	return rule, nil
}

// GetEquivalence com Cache (Muito importante para o Simples Nacional)
func (r *cachedTaxRepository) GetEquivalence(ctx context.Context, CSOSN string, tipoOperacao string) (*TaxEquivalence, error) {
	cacheKey := fmt.Sprintf("tax:equiv:%s:%s", CSOSN, tipoOperacao)

	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var eq TaxEquivalence
		if err := json.Unmarshal([]byte(val), &eq); err == nil {
			return &eq, nil
		}
	}

	eq, err := r.realRepo.GetEquivalence(ctx, CSOSN, tipoOperacao)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(eq)
	r.rdb.Set(ctx, cacheKey, data, r.ttl)

	return eq, nil
}

// Implementar GetSimplesFaixa seguindo o mesmo padrão...
func (r *cachedTaxRepository) GetSimplesFaixa(ctx context.Context, anexo string, rbt12 decimal.Decimal) (*SimplesFaixa, error) {
	// Para o Simples, a chave de cache pode ser por anexo e faixa de valor (ex: arredondado)
	// ou simplesmente delegar ao banco se o RBT12 variar muito por cliente.
	return r.realRepo.GetSimplesFaixa(ctx, anexo, rbt12)
}

func (r *cachedTaxRepository) GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error) {
	return r.realRepo.GetRateByNCM(ctx, ncm, uf)
}

func (r *cachedTaxRepository) GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error) {
	return r.realRepo.GetIBSRate(ctx, municipioIBGE)
}

func (r *cachedTaxRepository) GetProductException(ctx context.Context, ncmFull, ncmGroup, ufDestino string, regimeTributarioDestino string) (*ProductException, error) {
	return r.realRepo.GetProductException(ctx, ncmFull, ncmGroup, ufDestino, regimeTributarioDestino)
}

func (r *cachedTaxRepository) GetIPIRegra(ctx context.Context, NCM string, ExIPI string, CrtEmitente string, TipoOperacaoFiscal string, PerfilComprador string, UFDestino string, ZonaEspecial bool, DataOperacao string) (*IPIRegra, error) {
	cacheKey := fmt.Sprintf("tax:ipi:%s:%s:%s:%s:%s:%s:%t:%s", NCM, ExIPI, CrtEmitente, TipoOperacaoFiscal, PerfilComprador, UFDestino, ZonaEspecial, DataOperacao)

	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var regra IPIRegra
		if err := json.Unmarshal([]byte(val), &regra); err == nil {
			return &regra, nil
		}
	}

	regra, err := r.realRepo.GetIPIRegra(ctx, NCM, ExIPI, CrtEmitente, TipoOperacaoFiscal, PerfilComprador, UFDestino, ZonaEspecial, DataOperacao)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(regra)
	r.rdb.Set(ctx, cacheKey, data, r.ttl)

	return regra, nil
}

func (r *cachedTaxRepository) GetIvaDualRule(ctx context.Context, ncm, ufDestino, municipioIBGE string) (*IvaDualRule, error) {
	cacheKey := fmt.Sprintf("tax:iva:%s:%s:%s", ncm, ufDestino, municipioIBGE)

	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var rule IvaDualRule
		if err := json.Unmarshal([]byte(val), &rule); err == nil {
			return &rule, nil
		}
	}

	rule, err := r.realRepo.GetIvaDualRule(ctx, ncm, ufDestino, municipioIBGE)
	if err != nil {
		return nil, err
	}

	if rule != nil {
		data, _ := json.Marshal(rule)
		r.rdb.Set(ctx, cacheKey, data, r.ttl)
	}

	return rule, nil
}

// GetNCMSeletivo com cache Redis.
// A tabela ncm_seletivo muda raramente — cache de 24h reduz latencia
// significativamente no pre-filtro do IS.
func (r *cachedTaxRepository) GetNCMSeletivo(ctx context.Context, ncm string) (*NCMSeletivoRule, error) {
	cacheKey := fmt.Sprintf("tax:ncm_seletivo:%s", ncm)

	if val, err := r.rdb.Get(ctx, cacheKey).Result(); err == nil {
		var rule NCMSeletivoRule
		if err := json.Unmarshal([]byte(val), &rule); err == nil {
			return &rule, nil
		}
	}

	rule, err := r.realRepo.GetNCMSeletivo(ctx, ncm)
	if err != nil {
		return nil, err
	}

	if rule != nil {
		data, _ := json.Marshal(rule)
		r.rdb.Set(ctx, cacheKey, data, r.ttl)
	}

	return rule, nil
}
