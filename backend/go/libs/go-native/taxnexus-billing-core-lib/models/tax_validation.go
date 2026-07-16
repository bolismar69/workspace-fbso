// path: backend/go/libs/go-native/taxnexus-billing-core-lib/models/tax_validation.go
package models

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

// ValidationError represents a structured validation error for a single field.
type ValidationError struct {
	Field         string   `json:"field"`
	Code          string   `json:"code"`
	Message       string   `json:"message"`
	AllowedValues []string `json:"allowed_values,omitempty"`
}

// ValidationErrors is a slice of ValidationError that also implements the error interface.
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	msgs := make([]string, len(ve))
	for i, e := range ve {
		msgs[i] = e.Field + ": " + e.Message
	}
	return strings.Join(msgs, "; ")
}

// allowedKeysOf returns a sorted slice of string representations of all keys in a set map.
func allowedKeysOf[K ~string](m map[K]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	return keys
}

var (
	validatorOnce sync.Once
	validatorInst *validator.Validate
)

func getPayloadValidator() *validator.Validate {
	validatorOnce.Do(func() {
		v := validator.New()

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" || name == "" {
				return fld.Name
			}
			return name
		})

		_ = v.RegisterValidation("dec_gt_zero", func(fl validator.FieldLevel) bool {
			dec, ok := fl.Field().Interface().(decimal.Decimal)
			if !ok {
				return false
			}
			return dec.IsPositive()
		})

		_ = v.RegisterValidation("tipo_operacao_fiscal", func(fl validator.FieldLevel) bool {
			_, ok := validTiposOperacao[NormalizeTipoOperacao(fl.Field().String())]
			return ok
		})

		_ = v.RegisterValidation("crt_emitente", func(fl validator.FieldLevel) bool {
			_, ok := validCRTEmitentes[NormalizeCRTEmitente(fl.Field().String())]
			return ok
		})

		_ = v.RegisterValidation("indicador_presenca", func(fl validator.FieldLevel) bool {
			_, ok := validIndicadoresPresenca[NormalizeIndicadorPresenca(fl.Field().String())]
			return ok
		})

		_ = v.RegisterValidation("natureza_operacao", func(fl validator.FieldLevel) bool {
			_, ok := validNaturezasOperacao[NormalizeNaturezaOperacao(fl.Field().String())]
			return ok
		})

		validatorInst = v
	})

	return validatorInst
}

func trimStructPrefix(namespace string) string {
	parts := strings.SplitN(namespace, ".", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return namespace
}

func allowedValuesForTag(tag string) []string {
	switch tag {
	case "tipo_operacao_fiscal":
		return allowedKeysOf(validTiposOperacao)
	case "crt_emitente":
		return allowedKeysOf(validCRTEmitentes)
	case "indicador_presenca":
		return allowedKeysOf(validIndicadoresPresenca)
	case "natureza_operacao":
		return allowedKeysOf(validNaturezasOperacao)
	default:
		return nil
	}
}

func mapValidationTagToCode(fieldPath string, fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "REQUIRED"
	case "len":
		return "INVALID_FORMAT"
	case "min":
		if fieldPath == "itens" {
			return "REQUIRED"
		}
		return "INVALID_VALUE"
	default:
		return "INVALID_VALUE"
	}
}

func mapValidationTagToMessage(fieldPath string, fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "campo obrigatório"
	case "len":
		if fieldPath == "localizacao_origem.uf" || fieldPath == "localizacao_destino.uf" {
			return "deve ter exatamente 2 caracteres (ex: SP, RJ)"
		}
		return "formato inválido"
	case "min":
		if fieldPath == "itens" {
			return "ao menos um item é obrigatório"
		}
		return "valor inválido"
	case "dec_gt_zero":
		return "deve ser maior que zero"
	default:
		return fmt.Sprintf("valor inválido: %q", fe.Value())
	}
}

func collectStructValidationErrors(value interface{}) ValidationErrors {
	var errs ValidationErrors
	err := getPayloadValidator().Struct(value)
	if err == nil {
		return errs
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	seen := map[string]struct{}{}
	for _, fe := range validationErrs {
		fieldPath := trimStructPrefix(fe.Namespace())
		key := fieldPath + "|" + fe.Tag()
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}

		errs = append(errs, ValidationError{
			Field:         fieldPath,
			Code:          mapValidationTagToCode(fieldPath, fe),
			Message:       mapValidationTagToMessage(fieldPath, fe),
			AllowedValues: allowedValuesForTag(fe.Tag()),
		})
	}

	return errs
}

func validateDocumentoDetalhes(d *DocumentoFiscalEntrada) ValidationErrors {
	var errs ValidationErrors
	for _, det := range d.DetalhesDocumentoFiscal {
		if KeyDocumentoInfos(strings.ToUpper(strings.TrimSpace(det.Key))) == KeyDocumentoInfosZonaEspecial {
			valStr := fmt.Sprintf("%v", det.Value)
			if valStr != "" && !IsZonaEspecial(valStr) {
				errs = append(errs, ValidationError{
					Field:         fmt.Sprintf("detalhes.%s", KeyDocumentoInfosZonaEspecial),
					Code:          "INVALID_VALUE",
					Message:       fmt.Sprintf("valor inválido: %q", valStr),
					AllowedValues: []string{"ZFM", "ALC"},
				})
			}
		}
	}
	return errs
}

func validateItemDetalhes(i *ItemDocumentoFiscalEntrada, prefix string) ValidationErrors {
	var errs ValidationErrors
	for _, det := range i.DetalhesItemDocumentoFiscal {
		key := KeyDocumentoInfos(strings.ToUpper(strings.TrimSpace(det.Key)))
		valStr := fmt.Sprintf("%v", det.Value)
		if valStr == "" {
			continue
		}

		switch key {
		case KeyDocumentoInfosItemPISCOFINSCSTPIS:
			if _, ok := validCSTPISCOFINS[NormalizeCSTPISCOFINS(valStr)]; !ok {
				errs = append(errs, ValidationError{
					Field:         prefix + fmt.Sprintf("detalhes.%s", KeyDocumentoInfosItemPISCOFINSCSTPIS),
					Code:          "INVALID_VALUE",
					Message:       fmt.Sprintf("CST PIS inválido: %q", valStr),
					AllowedValues: allowedKeysOf(validCSTPISCOFINS),
				})
			}
		case KeyDocumentoInfosItemPISCOFINSCSTCOFINS:
			if _, ok := validCSTPISCOFINS[NormalizeCSTPISCOFINS(valStr)]; !ok {
				errs = append(errs, ValidationError{
					Field:         prefix + fmt.Sprintf("detalhes.%s", KeyDocumentoInfosItemPISCOFINSCSTCOFINS),
					Code:          "INVALID_VALUE",
					Message:       fmt.Sprintf("CST COFINS inválido: %q", valStr),
					AllowedValues: allowedKeysOf(validCSTPISCOFINS),
				})
			}
		case KeyDocumentoInfosItemSubstituirCSTICMS:
			if _, ok := validCSTICMS[NormalizeCSTICMS(valStr)]; !ok {
				errs = append(errs, ValidationError{
					Field:         prefix + fmt.Sprintf("detalhes.%s", KeyDocumentoInfosItemSubstituirCSTICMS),
					Code:          "INVALID_VALUE",
					Message:       fmt.Sprintf("CST ICMS inválido: %q", valStr),
					AllowedValues: allowedKeysOf(validCSTICMS),
				})
			}
		}
	}
	return errs
}

// Valid-set maps for P1 one-of validation.
var validCRTEmitentes = map[CRTEmitente]struct{}{
	CRTEmitenteSimples:        {},
	CRTEmitenteLucroReal:      {},
	CRTEmitenteLucroPresumido: {},
	CRTEmitenteMEI:            {},
}

var validTiposOperacao = map[TipoOperacaoFiscal]struct{}{
	TipoOperacaoFiscalSaida:         {},
	TipoOperacaoFiscalEntrada:       {},
	TipoOperacaoFiscalTransferencia: {},
}

var validIndicadoresPresenca = map[IndicadorPresenca]struct{}{
	IndicadorPresencaCodigo0: {},
	IndicadorPresencaCodigo1: {},
	IndicadorPresencaCodigo2: {},
	IndicadorPresencaCodigo3: {},
	IndicadorPresencaCodigo4: {},
	IndicadorPresencaCodigo5: {},
	IndicadorPresencaCodigo6: {},
	IndicadorPresencaCodigo7: {},
	IndicadorPresencaCodigo8: {},
	IndicadorPresencaCodigo9: {},
}

var validNaturezasOperacao = map[NaturezaOperacao]struct{}{
	NaturezaOperacaoRevenda:          {},
	NaturezaOperacaoConsumo:          {},
	NaturezaOperacaoAtivoImobiliario: {},
	NaturezaOperacaoServico:          {},
	NaturezaOperacaoOutros:           {},
}

var validCSTPISCOFINS = map[CSTPISCOFINS]struct{}{
	CSTPISCOFINS01: {},
	CSTPISCOFINS02: {},
	CSTPISCOFINS03: {},
	CSTPISCOFINS04: {},
	CSTPISCOFINS05: {},
	CSTPISCOFINS06: {},
	CSTPISCOFINS49: {},
	CSTPISCOFINS99: {},
}

var validCSTICMS = map[CSTICMS]struct{}{
	CSTICMS000: {},
	CSTICMS010: {},
}

// Validate checks DocumentoFiscalEntrada for required fields, format rules,
// and P1 one-of domain constraints. Returns a nil (empty) slice when valid.
func (d *DocumentoFiscalEntrada) Validate() ValidationErrors {
	err := collectStructValidationErrors(d)
	err = append(err, validateDocumentoDetalhes(d)...)
	for idx := range d.Itens {
		err = append(err, validateItemDetalhes(&d.Itens[idx], fmt.Sprintf("itens[%d].", idx))...)
	}

	if len(err) == 0 {
		return nil
	}

	return err
}

// Validate checks ItemDocumentoFiscalEntrada for required fields and P1 one-of domain constraints.
func (i *ItemDocumentoFiscalEntrada) Validate() ValidationErrors {
	err := collectStructValidationErrors(i)
	err = append(err, validateItemDetalhes(i, "")...)
	if len(err) == 0 {
		return nil
	}
	return err
}
