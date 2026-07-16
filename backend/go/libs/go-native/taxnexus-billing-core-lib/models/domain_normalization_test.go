package models

import "testing"

func TestNormalizeCRTEmitente(t *testing.T) {
	if got := NormalizeCRTEmitente(" simples "); got != CRTEmitenteSimples {
		t.Fatalf("expected SIMPLES, got %q", got)
	}
}

func TestNormalizeTipoOperacao(t *testing.T) {
	if got := NormalizeTipoOperacao(" saida "); got != TipoOperacaoFiscalSaida {
		t.Fatalf("expected SAIDA, got %q", got)
	}
}

func TestIsZonaEspecial(t *testing.T) {
	for _, value := range []string{"zfm", "ALC", "true", "SIM"} {
		if !IsZonaEspecial(value) {
			t.Fatalf("expected %q to be treated as zona especial", value)
		}
	}

	if IsZonaEspecial("nao") {
		t.Fatalf("expected non special marker to return false")
	}
}
