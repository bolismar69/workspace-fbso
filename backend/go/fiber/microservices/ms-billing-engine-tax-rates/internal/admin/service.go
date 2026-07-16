// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/admin/service.go
package admin

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

// CacheInvalidator define operação de invalidação de cache Redis.
type CacheInvalidator interface {
	InvalidateIvaDualCache(ctx context.Context, ncm, uf, municipioIBGE string) error
}

// RedisCacheInvalidator implementa CacheInvalidator usando Redis.
type RedisCacheInvalidator struct {
	rdb *redis.Client
}

// NewRedisCacheInvalidator cria um invalidator de cache via Redis.
func NewRedisCacheInvalidator(rdb *redis.Client) *RedisCacheInvalidator {
	return &RedisCacheInvalidator{rdb: rdb}
}

func (c *RedisCacheInvalidator) InvalidateIvaDualCache(ctx context.Context, ncm, uf, municipioIBGE string) error {
	// Formato da chave: tax:iva:<ncm>:<uf>:<municipio>
	key := fmt.Sprintf("tax:iva:%s:%s:%s", ncm, uf, municipioIBGE)
	if municipioIBGE == "" {
		key = fmt.Sprintf("tax:iva:%s:%s:*", ncm, uf)
	}

	if municipioIBGE != "" {
		if err := c.rdb.Del(ctx, key).Err(); err != nil {
			return fmt.Errorf("falha ao invalidar cache %s: %w", key, err)
		}
	} else {
		// Se não tem municipio, invalida todas as keys com prefixo
		pattern := fmt.Sprintf("tax:iva:%s:%s:*", ncm, uf)
		keys, err := c.rdb.Keys(ctx, pattern).Result()
		if err != nil {
			return fmt.Errorf("falha ao buscar keys para invalidar: %w", err)
		}
		if len(keys) > 0 {
			if err := c.rdb.Del(ctx, keys...).Err(); err != nil {
				return fmt.Errorf("falha ao invalidar %d keys: %w", len(keys), err)
			}
		}
	}

	slog.Info("Cache Redis invalidado", "key", key)
	return nil
}

// AdminTaxService gerencia operações administrativas sobre alíquotas (BR-02).
type AdminTaxService struct {
	repo     AdminRepository
	cache    CacheInvalidator
}

// NewAdminTaxService cria um novo AdminTaxService.
func NewAdminTaxService(repo AdminRepository, cache CacheInvalidator) *AdminTaxService {
	return &AdminTaxService{repo: repo, cache: cache}
}

// UpsertRule valida e insere/atualiza uma regra IVA Dual.
//
// Validações:
//   - NCM: 8 dígitos numéricos
//   - UF: 2 letras maiúsculas
//   - Alíquotas: [0, 100]
func (s *AdminTaxService) UpsertRule(ctx context.Context, input IvaDualRuleInput, changedBy string) (*IvaDualRuleOutput, error) {
	if err := validateRuleInput(input); err != nil {
		return nil, err
	}

	if input.InicioValidade.IsZero() {
		input.InicioValidade = time.Now()
	}

	out, err := s.repo.UpsertIvaDualRule(ctx, input, changedBy)
	if err != nil {
		return nil, fmt.Errorf("falha ao salvar regra: %w", err)
	}

	// Invalida cache Redis para que próxima consulta use a nova regra
	if err := s.cache.InvalidateIvaDualCache(ctx, input.NCM, input.UFDestino, input.MunicipioDestinoIBGE); err != nil {
		slog.Warn("Falha ao invalidar cache Redis após upsert", "error", err)
	}

	slog.Info("Regra IVA Dual atualizada",
		"id", out.ID,
		"ncm", out.NCM,
		"uf_destino", out.UFDestino,
		"municipio_ibge", out.MunicipioDestinoIBGE,
		"changed_by", changedBy,
	)

	return out, nil
}

// ListRules lista regras IVA Dual com filtros opcionais.
func (s *AdminTaxService) ListRules(ctx context.Context, filter ListRulesFilter) ([]IvaDualRuleOutput, error) {
	return s.repo.ListIvaDualRules(ctx, filter)
}

// ─── Validações ────────────────────────────────────────────────────────────

func validateRuleInput(input IvaDualRuleInput) error {
	// NCM: 8 dígitos numéricos
	if len(input.NCM) != 8 {
		return fmt.Errorf("NCM deve ter 8 dígitos, recebido %q (%d caracteres)", input.NCM, len(input.NCM))
	}
	for _, c := range input.NCM {
		if c < '0' || c > '9' {
			return fmt.Errorf("NCM deve conter apenas dígitos, recebido %q", input.NCM)
		}
	}

	// UF: 2 letras maiúsculas
	if len(input.UFDestino) != 2 {
		return fmt.Errorf("UF deve ter 2 caracteres, recebido %q", input.UFDestino)
	}
	uf := strings.ToUpper(input.UFDestino)
	if uf[0] < 'A' || uf[0] > 'Z' || uf[1] < 'A' || uf[1] > 'Z' {
		return fmt.Errorf("UF deve conter apenas letras, recebido %q", input.UFDestino)
	}
	input.UFDestino = uf

	// Alíquotas: [0, 100]
	if err := validateAliquota("CBS", input.AliquotaCBS); err != nil {
		return err
	}
	if err := validateAliquota("IBS Estadual", input.AliquotaIBSEstadual); err != nil {
		return err
	}
	if err := validateAliquota("IBS Municipal", input.AliquotaIBSMunicipal); err != nil {
		return err
	}
	if err := validateAliquota("IS", input.AliquotaIS); err != nil {
		return err
	}
	if err := validateAliquota("Redução", input.PercentualReducao); err != nil {
		return err
	}

	return nil
}

func validateAliquota(nome string, valor decimal.Decimal) error {
	if valor.LessThan(decimal.Zero) {
		return fmt.Errorf("%s não pode ser negativa: %s", nome, valor)
	}
	if valor.GreaterThan(decimal.NewFromInt(100)) {
		return fmt.Errorf("%s não pode exceder 100%%: %s", nome, valor)
	}
	return nil
}
