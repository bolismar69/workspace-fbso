// path: backend/go/libs/go-native/taxnexus-billing-core-lib/repository/entities.go
package repository

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

// ICMSConfig consolida os dados para o motor de cálculo
type ICMSConfig struct {
	Id                    int32
	UFOrigem              string
	UFDestino             string
	AliquotaInterna       decimal.Decimal
	AliquotaInterestadual decimal.Decimal
	CSTPadrao             string
	ReducaoBase           decimal.Decimal
	PercentualFCP         decimal.Decimal
	MVAPadrao             decimal.Decimal
	PossuiProtocoloST     bool
	InicioValidade        time.Time
	FinalValidade         *time.Time
	CriadoEm              time.Time
	AtualizadoEm          time.Time
}

type FederalTaxRule struct {
	Id               int32
	RegimeTributario string
	CSTPIS           string
	CSTCOFINS        string
	AliquotaPIS      decimal.Decimal
	AliquotaCOFINS   decimal.Decimal
	ExcluiICMSBase   bool
	InicioValidade   time.Time
	FinalValidade    *time.Time
	CriadoEm         time.Time
	AtualizadoEm     time.Time
}

type ICMSRule struct {
	Id                    int32
	UFOrigem              string
	UFDestino             string
	AliquotaInterna       decimal.Decimal
	AliquotaInterestadual decimal.Decimal
	CSTPadrao             string
	ReducaoBase           decimal.Decimal
	PercentualFCP         decimal.Decimal
	MVAPadrao             decimal.Decimal
	PossuiProtocoloST     bool
	InicioValidade        time.Time
	FinalValidade         *time.Time
	CriadoEm              time.Time
	AtualizadoEm          time.Time
}

type TaxEquivalence struct {
	Id             int32
	CSOSN          string
	CSTEquivalente string
	PermiteCredito bool
	Descricao      string
	InicioValidade time.Time
	FinalValidade  *time.Time
	CriadoEm       time.Time
	AtualizadoEm   time.Time
}

type SimplesFaixa struct {
	Id             int32
	Anexo          string
	Faixa          int32
	FaturamentoMin decimal.Decimal
	FaturamentoMax decimal.Decimal
	AliqNominal    decimal.Decimal
	ValorDeduzir   decimal.Decimal
	PercIcmsAnexo  decimal.Decimal
	InicioValidade time.Time
	FinalValidade  *time.Time
	CriadoEm       time.Time
	AtualizadoEm   time.Time
}

type ProductException struct {
	Id                     int32
	NCM                    string
	UFDestino              string
	CSTPIS                 string
	CSTCOFINS              string
	CSTICMS                string
	AliquotaPISUnitario    decimal.Decimal
	AliquotaCOFINSUnitario decimal.Decimal
	MVAST                  decimal.Decimal
	CEST                   string
	PossuiProtocoloST      bool
	CSOSN                  string
	AliquotaInternaDestino decimal.Decimal // Novo campo
	AliquotaInterestadual  decimal.Decimal // Novo campo
	PercentualFCP          decimal.Decimal // Novo campo
	ReducaoBase            decimal.Decimal // Novo campo
	InicioValidade         time.Time
	FinalValidade          *time.Time
	CriadoEm               time.Time
	AtualizadoEm           time.Time
}

// ANALISANDO PRIMEIRO ESBOCO DA ENTIDADE REGRAS_IPI
type IPIRegra struct {
	ID                 int             `db:"id"`
	NCM                string          `db:"ncm"`
	ExIPI              string          `db:"ex_ipi"`
	CrtEmitente        string          `db:"crt_emitente"`
	TipoOperacaoFiscal string          `db:"tipo_operacao_fiscal"`
	PerfilComprador    string          `db:"perfil_comprador"`
	UFDestino          string          `db:"uf_destino"`
	ZonaEspecial       string          `db:"zona_especial"`
	AliquotaIPI        decimal.Decimal `db:"aliquota_ipi"`
	ValorPautaIPI      decimal.Decimal `db:"valor_pauta_ipi"`
	CSTIPI             string          `db:"cst_ipi"`
	CEnq               string          `db:"c_enq"`
	PossuiDesoneracao  bool            `db:"possui_desoneracao"`
	MotivoDesoneracao  sql.NullString  `db:"motivo_desoneracao"`
	InicioValidade     time.Time       `db:"inicio_validade"`
	FinalValidade      *time.Time      `db:"final_validade"` // Pointer para lidar com NULL
	CriadoEm           time.Time       `db:"criado_em"`
	AtualizadoEm       time.Time       `db:"atualizado_em"`
}

// NCMSeletivoRule representa uma entrada na tabela ncm_seletivo.
// Define quais NCMs estao sujeitos ao Imposto Seletivo (IS) e qual aliquota
// se aplica por categoria de produto.
// Ref: PROCEDURE-FIN-00001 SOP-003, RULES-CATALOG-FIN-00001 BR-TAX-INF-005
type NCMSeletivoRule struct {
	NCM            string          `db:"ncm"`
	Categoria      string          `db:"categoria"`
	AliquotaIS     decimal.Decimal `db:"aliquota_is"`
	Descricao      sql.NullString  `db:"descricao"`
	InicioValidade time.Time       `db:"inicio_validade"`
	FinalValidade  *time.Time      `db:"final_validade"`
}

type IvaDualRule struct {
	NCM                  string          `db:"ncm"`
	UFDestino            string          `db:"uf_destino"`
	MunicipioDestinoIBGE sql.NullString  `db:"municipio_destino_ibge"`
	AliquotaCBS          decimal.Decimal `db:"aliquota_cbs"`
	AliquotaIBSEstadual  decimal.Decimal `db:"aliquota_ibs_estadual"`
	AliquotaIBSMunicipal decimal.Decimal `db:"aliquota_ibs_municipal"`
	PercentualReducao    decimal.Decimal `db:"percentual_reducao"`
	IsImpostoSeletivo    bool            `db:"is_imposto_seletivo"`
	AliquotaIS           decimal.Decimal `db:"aliquota_is"`
	InicioValidade       time.Time       `db:"inicio_validade"`
	FinalValidade        *time.Time      `db:"final_validade"`
}
