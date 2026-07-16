// path: backend/go/libs/go-native/taxnexus-billing-core-lib/repository/contracts.go
package repository

import (
	"context"

	"github.com/shopspring/decimal"
)

type TaxRepository interface {
	GetRateByNCM(ctx context.Context, ncm string, uf string) (float64, error)
	GetIBSRate(ctx context.Context, municipioIBGE string) (float64, error)
	GetFederalTaxRule(ctx context.Context, regimeTributario, cstPIS, cstCOFINS string) (*FederalTaxRule, error)
	GetICMSRule(ctx context.Context, orig, dest string) (*ICMSRule, error)
	GetEquivalence(ctx context.Context, CSOSN string, tipoOperacao string) (*TaxEquivalence, error)
	GetSimplesFaixa(ctx context.Context, anexo string, rbt12 decimal.Decimal) (*SimplesFaixa, error)
	GetProductException(ctx context.Context, ncmFull, ncmGroup, ufDestino string, regimeTributarioDestino string) (*ProductException, error)
	GetIPIRegra(ctx context.Context, NCM string, ExIPI string, CrtEmitente string, TipoOperacaoFiscal string, PerfilComprador string, UFDestino string, ZonaEspecial bool, DataOperacao string) (*IPIRegra, error)
	GetIvaDualRule(ctx context.Context, ncm, ufDestino, municipioIBGE string) (*IvaDualRule, error)
	GetNCMSeletivo(ctx context.Context, ncm string) (*NCMSeletivoRule, error)
}
