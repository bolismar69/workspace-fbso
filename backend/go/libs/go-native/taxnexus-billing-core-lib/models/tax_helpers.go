// path: backend/go/libs/go-native/taxnexus-billing-core-lib/models/tax_helpers.go
package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func normalizeDetalheKey(key interface{}) string {
	switch value := key.(type) {
	case KeyDocumentoInfos:
		return strings.ToUpper(strings.TrimSpace(string(value)))
	case string:
		return strings.ToUpper(strings.TrimSpace(value))
	case fmt.Stringer:
		return strings.ToUpper(strings.TrimSpace(value.String()))
	case nil:
		return ""
	default:
		return strings.ToUpper(strings.TrimSpace(fmt.Sprint(value)))
	}
}

func detalheKeyCandidates(key interface{}) []string {
	normalized := normalizeDetalheKey(key)
	if normalized == "" {
		return nil
	}

	candidates := []string{normalized}
	current := normalized
	prefixes := []string{
		"ITEM_",
		"DOCUMENTO_",
		"SUBSTITUIR_",
		"DETALHE_",
		"DESONERACAO_",
		"PISCOFINS_",
		"ST_",
	}

	for {
		trimmed := current
		for _, prefix := range prefixes {
			if strings.HasPrefix(trimmed, prefix) {
				trimmed = strings.TrimPrefix(trimmed, prefix)
				break
			}
		}
		if trimmed == current {
			break
		}
		candidates = append(candidates, trimmed)
		current = trimmed
	}

	return candidates
}

func getMapValue(m map[string]interface{}, key interface{}) (interface{}, bool) {
	for _, candidate := range detalheKeyCandidates(key) {
		if value, ok := m[candidate]; ok {
			return value, true
		}
	}
	return nil, false
}

func toDecimal(value interface{}) (decimal.Decimal, bool) {
	switch typed := value.(type) {
	case decimal.Decimal:
		return typed, true
	case *decimal.Decimal:
		if typed == nil {
			return decimal.Zero, false
		}
		return *typed, true
	case json.Number:
		converted, err := decimal.NewFromString(typed.String())
		return converted, err == nil
	case string:
		converted, err := decimal.NewFromString(strings.TrimSpace(typed))
		return converted, err == nil
	case int:
		return decimal.NewFromInt(int64(typed)), true
	case int8:
		return decimal.NewFromInt(int64(typed)), true
	case int16:
		return decimal.NewFromInt(int64(typed)), true
	case int32:
		return decimal.NewFromInt(int64(typed)), true
	case int64:
		return decimal.NewFromInt(typed), true
	case uint:
		return decimal.NewFromInt(int64(typed)), true
	case uint8:
		return decimal.NewFromInt(int64(typed)), true
	case uint16:
		return decimal.NewFromInt(int64(typed)), true
	case uint32:
		return decimal.NewFromInt(int64(typed)), true
	case uint64:
		converted, err := decimal.NewFromString(fmt.Sprintf("%d", typed))
		return converted, err == nil
	case float32:
		return decimal.NewFromFloat32(typed), true
	case float64:
		return decimal.NewFromFloat(typed), true
	default:
		return decimal.Zero, false
	}
}

func cloneAnyMap(source map[string]interface{}) map[string]interface{} {
	if len(source) == 0 {
		return map[string]interface{}{}
	}
	cloned := make(map[string]interface{}, len(source))
	for key, value := range source {
		cloned[key] = value
	}
	return cloned
}

func cloneDecimalMap(source map[string]decimal.Decimal) map[string]decimal.Decimal {
	if len(source) == 0 {
		return map[string]decimal.Decimal{}
	}
	cloned := make(map[string]decimal.Decimal, len(source))
	for key, value := range source {
		cloned[key] = value
	}
	return cloned
}

func NewDetalhesNormalizados(detalhes []Detalhe) DetalhesNormalizados {
	normalized := DetalhesNormalizados{
		Todos:    make(map[string]interface{}, len(detalhes)),
		Decimais: make(map[string]decimal.Decimal),
	}
	for _, detalhe := range detalhes {
		normalized.Set(detalhe.Key, detalhe.Value)
	}
	return normalized
}

func (d *DetalhesNormalizados) Set(key interface{}, value interface{}) {
	if d.Todos == nil {
		d.Todos = make(map[string]interface{})
	}
	if d.Decimais == nil {
		d.Decimais = make(map[string]decimal.Decimal)
	}

	normalizedKey := normalizeDetalheKey(key)
	if normalizedKey == "" {
		return
	}

	d.Todos[normalizedKey] = value
	if converted, ok := toDecimal(value); ok {
		d.Decimais[normalizedKey] = converted
	} else {
		delete(d.Decimais, normalizedKey)
	}
}

func (d DetalhesNormalizados) Merge(other DetalhesNormalizados) DetalhesNormalizados {
	merged := DetalhesNormalizados{
		Todos:    cloneAnyMap(d.Todos),
		Decimais: cloneDecimalMap(d.Decimais),
	}
	for key, value := range other.Todos {
		merged.Todos[key] = value
	}
	for key, value := range other.Decimais {
		merged.Decimais[key] = value
	}
	return merged
}

func (d *DocumentoFiscalEntrada) PrepararMapasDetalhes() {
	if d == nil {
		return
	}

	d.MapaDetalhes = NewDetalhesNormalizados(d.DetalhesDocumentoFiscal)
	for index := range d.Itens {
		d.Itens[index].PrepararMapasDetalhes(d.MapaDetalhes)
	}
}

func (d *DocumentoFiscalEntrada) ToMap() map[string]interface{} {
	if d == nil {
		return map[string]interface{}{}
	}
	if d.MapaDetalhes.Todos == nil {
		d.PrepararMapasDetalhes()
	}
	return cloneAnyMap(d.MapaDetalhes.Todos)
}

func (d *DocumentoFiscalEntrada) ToDecimalMap() map[string]decimal.Decimal {
	if d == nil {
		return map[string]decimal.Decimal{}
	}
	if d.MapaDetalhes.Todos == nil {
		d.PrepararMapasDetalhes()
	}
	return cloneDecimalMap(d.MapaDetalhes.Decimais)
}

func (d *DocumentoFiscalEntrada) AddDetalhe(key interface{}, value interface{}) {
	if d == nil {
		return
	}
	stringKey := normalizeDetalheKey(key)
	if stringKey == "" {
		return
	}
	d.DetalhesDocumentoFiscal = append(d.DetalhesDocumentoFiscal, Detalhe{Key: stringKey, Value: value})
	d.MapaDetalhes.Set(stringKey, value)
	for index := range d.Itens {
		d.Itens[index].MapaDetalhes = d.MapaDetalhes.Merge(NewDetalhesNormalizados(d.Itens[index].DetalhesItemDocumentoFiscal))
	}
}

func (i *ItemDocumentoFiscalEntrada) PrepararMapasDetalhes(base ...DetalhesNormalizados) {
	if i == nil {
		return
	}

	merged := DetalhesNormalizados{}
	for _, currentBase := range base {
		merged = merged.Merge(currentBase)
	}
	merged = merged.Merge(NewDetalhesNormalizados(i.DetalhesItemDocumentoFiscal))
	i.MapaDetalhes = merged
}

func (i *ItemDocumentoFiscalEntrada) ToMap() map[string]interface{} {
	if i == nil {
		return map[string]interface{}{}
	}
	if i.MapaDetalhes.Todos == nil {
		i.PrepararMapasDetalhes()
	}
	return cloneAnyMap(i.MapaDetalhes.Todos)
}

func (i *ItemDocumentoFiscalEntrada) ToDecimalMap() map[string]decimal.Decimal {
	if i == nil {
		return map[string]decimal.Decimal{}
	}
	if i.MapaDetalhes.Todos == nil {
		i.PrepararMapasDetalhes()
	}
	return cloneDecimalMap(i.MapaDetalhes.Decimais)
}

func (i *ItemDocumentoFiscalEntrada) AddDetalhe(key interface{}, value interface{}) {
	if i == nil {
		return
	}
	stringKey := normalizeDetalheKey(key)
	if stringKey == "" {
		return
	}
	i.DetalhesItemDocumentoFiscal = append(i.DetalhesItemDocumentoFiscal, Detalhe{Key: stringKey, Value: value})
	i.MapaDetalhes.Set(stringKey, value)
}

func GetFloat(m map[string]interface{}, key interface{}) float64 {
	val, ok := getMapValue(m, key)
	if !ok {
		return 0
	}
	switch v := val.(type) {
	case decimal.Decimal:
		floatValue, _ := v.Float64()
		return floatValue
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	default:
		return 0
	}
}

func GetString(m map[string]interface{}, key interface{}) string {
	val, ok := getMapValue(m, key)
	if !ok {
		return ""
	}
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

func GetDecimal(m map[string]interface{}, key interface{}) decimal.Decimal {
	val, ok := getMapValue(m, key)
	if !ok {
		return decimal.Zero
	}
	if converted, ok := toDecimal(val); ok {
		return converted
	}
	return decimal.Zero
}
