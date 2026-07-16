# Cálculo de IRPF, Design Técnico

> Especificação técnica do endpoint de cálculo de imposto de renda. Foca na implementação paralela e lógica de recomendação.

## Interface

### Endpoint HTTP 🟢

| Método | Caminho | Entrada | Saída | Status codes |
|--------|---------|---------|-------|--------------|
| POST | `/api/v1/calculate/irpf` | `models.UniversalTaxRequest` | `map[string]models.TaxResponse` | 200, 400, 500 |

**Campos de Entrada (Sinais em `taxnexus-individual-core-lib/models`):**
- `TaxCode`: "IRPF" 🟢
- `GrossIncome`: Renda bruta 🟢
- `CalculationType`: "monthly" ou "annual" 🟢
- `ReferenceDate`: Data para vigência de tabelas 🟢
- `Inputs`: Lista de deduções (pension_amount, education_expenses, health_expenses, pgbl_contribution, dependents_qty) 🟢

**Campos de Saída:**
- Mapa com chaves `completa` e `simplificada`.
- Cada objeto contém `TaxAmount`, `BaseValue`, `TotalDeductionAmount`, `EffectiveRate`, `DeductionDetails`, `IsRecommended`. 🟢

## Fluxo Principal

1. **Extração de Contexto:** Obtém `X-Request-ID` e injeta no contexto Go (`tax_handler.go:23`). 🟢
2. **Carregamento de Configurações:** Busca tabelas e parâmetros no repositório (`calculation_service.go:101`). 🟢
3. **Processamento Paralelo:** Dispara duas goroutines para calcular os modelos (`calculation_service.go:108-117`):
   - **Modelo Completo:** Inclui integração com INSS externo e deduções detalhadas.
   - **Modelo Simplificado:** Aplica desconto padrão de 20% limitado ao teto.
4. **Consolidação:** Aguarda os resultados via channel (`resChan`). 🟢
5. **Recomendação:** Compara os `TaxAmount` e marca o menor como `IsRecommended: true` (`calculation_service.go:132-140`). 🟢
6. **Resposta:** Retorna JSON com ambos os cálculos. 🟢

## Fluxos Alternativos

- **Falha no INSS:** Se o microserviço de INSS falhar, o sistema registra um aviso (Warn) e prossegue com o cálculo do modelo completo sem a dedução previdenciária (`calculation_service.go:159`). 🟢
- **Data Omitida:** Se `ReferenceDate` for zero, assume a data atual (`tax_handler.go:42`). 🟢
- **Erro de Parsing:** Retorna HTTP 400 se o JSON de entrada estiver malformado (`tax_handler.go:37`). 🟢

## Dependências

- **INSSClient:** Chama MS externo de INSS via REST (`inss_client.go`). Timeout de 5 segundos configurado. 🟢
- **TaxRepository:** Acesso ao PostgreSQL/Redis para tabelas de alíquotas e configs (`calculation_service.go:34`). 🟢

## Decisões de Design Identificadas

| Decisão | Evidência no código | Confiança |
|---------|---------------------|-----------|
| Cálculos em paralelo via Goroutines | `calculation_service.go:108` | 🟢 |
| Resiliência a falhas de integração (INSS) | `calculation_service.go:159` | 🟢 |
| Rastreabilidade via RequestID propagado no Context | `tax_handler.go:23` | 🟢 |
| Redução progressiva da Reforma 2026 | `calculation_service.go:275-296` | 🟢 |
| Timeout de 5s para integração INSS | `inss_client.go:19` | 🟢 |

## Estado Interno
A unit é **stateless**. Todo o processamento depende dos inputs da requisição e das tabelas carregadas do banco de dados na hora do cálculo. 🟢

## Observabilidade

- **Logs Estruturados:** Uso de `slog` com injeção de `trace_id` (`calculation_service.go:48`). 🟢
- **Métricas de Duração:** Log de info ao final do cálculo com `duration_ms` (`calculation_service.go:142`). 🟢

## Riscos e Lacunas

- 🟡 **Inferência:** Assume-se que o Redis é usado para cache de tabelas de alíquotas para otimizar o carregamento recorrente, mas a lógica de cache está encapsulada na lib `repository`.

