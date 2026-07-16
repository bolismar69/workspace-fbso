// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/legacy/telecom.go
package legacy

import (
	"log/slog"

	"taxnexus-billing-core-lib/models"
)

// NaturezaServicoTelecom classifica um serviço para fins de FUST/FUNTTEL.
//
// Classificações:
//   - SCM: Serviço de Comunicação Multimídia (banda larga, dados)
//   - STFC: Serviço Telefônico Fixo Comutado (telefonia fixa)
//   - SVA: Serviço de Valor Adicionado (streaming, antivírus, suporte técnico)
//
// Apenas SCM e STFC são sujeitos a FUST (Lei 9.998/2000) e FUNTTEL (Lei 10.052/2000).
// SVA não sofre incidência destas contribuições.
//
// Fonte: BR-TAX-INF-007, POLICE-FIN-00001 §4.3 (FUST/FUNTTEL)
const (
	NaturezaSCM   = "SCM"
	NaturezaSTFC  = "STFC"
	NaturezaSVA   = "SVA"
)

// TelecomClassifier determina se um item de serviço está sujeito a FUST/FUNTTEL.
type TelecomClassifier struct{}

// NewTelecomClassifier cria um novo classificador de serviços de telecom.
func NewTelecomClassifier() *TelecomClassifier {
	return &TelecomClassifier{}
}

// IsTelecomService verifica se a natureza do serviço é SCM ou STFC
// (serviços de telecom propriamente ditos, sujeitos a FUST/FUNTTEL).
func (c *TelecomClassifier) IsTelecomService(natureza string) bool {
	return natureza == NaturezaSCM || natureza == NaturezaSTFC
}

// IsSVA verifica se a natureza do serviço é SVA
// (Serviço de Valor Adicionado — não sujeito a FUST/FUNTTEL).
func (c *TelecomClassifier) IsSVA(natureza string) bool {
	return natureza == NaturezaSVA
}

// Classify extrai a natureza do serviço dos detalhes do item.
// Retorna a string da natureza e um bool indicando se foi encontrada.
func (c *TelecomClassifier) Classify(item models.ItemDocumentoFiscalEntrada) (string, bool) {
	attrs := item.ToMap()
	natureza := models.GetString(attrs, models.KeyDocumentoInfos("NATUREZA_SERVICO"))
	if natureza == "" {
		return "", false
	}
	return natureza, true
}

// MustCalculateFUST verifica todas as condições para incidência de FUST:
//  1. Item tem natureza de serviço definida
//  2. Natureza é SCM ou STFC (não SVA)
func (c *TelecomClassifier) MustCalculateFUST(item models.ItemDocumentoFiscalEntrada) bool {
	natureza, found := c.Classify(item)
	if !found {
		slog.Debug("FUST/FUNTTEL: item sem natureza de serviço definida, assumindo SVA (não incide)",
			"SKU", item.SKU,
		)
		return false
	}

	if c.IsSVA(natureza) {
		slog.Debug("FUST/FUNTTEL: item SVA, não incide",
			"SKU", item.SKU,
			"natureza", natureza,
		)
		return false
	}

	if !c.IsTelecomService(natureza) {
		slog.Warn("FUST/FUNTTEL: natureza de serviço desconhecida, não incide",
			"SKU", item.SKU,
			"natureza", natureza,
		)
		return false
	}

	return true
}
