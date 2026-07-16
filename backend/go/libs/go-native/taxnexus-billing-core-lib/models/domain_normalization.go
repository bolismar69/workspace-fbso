package models

import "strings"

func normalizeDomainValue(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func NormalizeCRTEmitente(value string) CRTEmitente {
	return CRTEmitente(normalizeDomainValue(value))
}

func NormalizeTipoOperacao(value string) TipoOperacaoFiscal {
	return TipoOperacaoFiscal(normalizeDomainValue(value))
}

func NormalizeIndicadorPresenca(value string) IndicadorPresenca {
	return IndicadorPresenca(normalizeDomainValue(value))
}

func NormalizeCSTPISCOFINS(value string) CSTPISCOFINS {
	return CSTPISCOFINS(normalizeDomainValue(value))
}

func NormalizeCSTICMS(value string) CSTICMS {
	return CSTICMS(normalizeDomainValue(value))
}

func NormalizeNaturezaOperacao(value string) NaturezaOperacao {
	return NaturezaOperacao(normalizeDomainValue(value))
}

func NormalizeZonaEspecial(value string) ZonaEspecial {
	switch normalizeDomainValue(value) {
	case string(ZonaEspecialZFM):
		return ZonaEspecialZFM
	case string(ZonaEspecialALC):
		return ZonaEspecialALC
	default:
		return ZonaEspecialNenhuma
	}
}

func IsZonaEspecial(value string) bool {
	normalized := normalizeDomainValue(value)
	if normalized == string(ZonaEspecialZFM) || normalized == string(ZonaEspecialALC) {
		return true
	}
	return normalized == "TRUE" || normalized == "SIM"
}
