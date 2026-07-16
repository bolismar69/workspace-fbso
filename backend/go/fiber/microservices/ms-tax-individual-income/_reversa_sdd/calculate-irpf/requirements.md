# Requirements: Cálculo de IRPF

> Identificador: `001-calculate-irpf`
> Data: `2026-06-08`
> Pasta da extração reversa: `_reversa_sdd/`
> Confidência: 🟢 CONFIRMADO, 🟡 INFERIDO, 🔴 LACUNA / DÚVIDA

## 1. Resumo executivo

O endpoint de Cálculo de IRPF processa dados financeiros de uma pessoa física para determinar o imposto de renda devido. Ele realiza cálculos simultâneos nos modelos **Completo** (com deduções detalhadas) e **Simplificado** (desconto padrão), recomendando automaticamente a opção mais vantajosa para o contribuinte, garantindo conformidade com as regras vigentes e as regras de transição da reforma de 2026.

## 2. Contexto a partir do legado

| Fonte | Trecho relevante | Confidência |
|-------|------------------|-------------|
| `_reversa_sdd/architecture.md#integrações-e-protocolos` | Endpoint POST /api/v1/calculate/irpf e integração INSS | 🟢 |
| `_reversa_sdd/domain.md#regras-de-negócio-principais` | Regras de seleção de modelo, limites de dedução e transição 2026 | 🟢 |
| `_reversa_sdd/code-analysis.md#services` | Implementação paralela via goroutines e lógica de recomendação | 🟢 |

## 3. Personas e cenários de uso

| Persona | Objetivo | Cenário-chave |
|---------|----------|---------------|
| Contribuinte | Calcular o imposto devido e escolher o melhor modelo de declaração. | Enviar renda e gastos para obter o valor do IRPF e a recomendação do modelo. |

## 4. Regras de negócio novas ou alteradas

1. **RN-01: Comparação de Modelos** 🟢
   - Origem no legado: `_reversa_sdd/domain.md#1-seleção-do-modelo-de-cálculo`
   - Tipo: confirmada
2. **RN-02: Limites de Dedução Progressiva** 🟢
   - Origem no legado: `_reversa_sdd/domain.md#2-limites-de-dedução-modelo-completo`
   - Tipo: confirmada
3. **RN-03: Reforma Tributária 2026** 🟢
   - Origem no legado: `_reversa_sdd/domain.md#5-mecanismo-de-transição-reforma-2026`
   - Tipo: confirmada

## 5. Requisitos Funcionais

| ID | Requisito | Prioridade | Critério de aceite | Confidência |
|----|-----------|------------|--------------------|-------------|
| RF-01 | Calcular IRPF Modelo Completo | Must | Subtrair deduções (INSS, dependentes, educação, saúde, PGBL) da renda bruta. | 🟢 |
| RF-02 | Calcular IRPF Modelo Simplificado | Must | Aplicar 20% de desconto sobre renda bruta, respeitando o teto configurado. | 🟢 |
| RF-03 | Obter INSS Externo | Must | Chamar o microserviço de INSS e usar o retorno como dedução no modelo completo. | 🟢 |
| RF-04 | Recomendar Melhor Opção | Must | Marcar o resultado com menor valor de imposto como `IsRecommended`. | 🟢 |
| RF-05 | Aplicar Transição 2026 | Must | Zerar imposto ou aplicar redução adicional se data >= 2026. | 🟢 |

## 6. Requisitos Não Funcionais

| Tipo | Requisito | Evidência ou justificativa | Confidência |
|------|-----------|----------------------------|-------------|
| Desempenho | Cálculos em paralelo | Uso de goroutines no `calculation_service.go:104`. | 🟢 |
| Observabilidade | Rastreabilidade (Trace ID) | Injeção de `requestid` no contexto em `tax_handler.go:23`. | 🟢 |
| Resiliência | Fallback INSS | Continua cálculo se MS de INSS falhar (`calculation_service.go:142`). | 🟢 |

## 7. Critérios de Aceitação

```gherkin
Cenário: Cálculo bem-sucedido com recomendação
  Dado que o contribuinte informa renda bruta de 5000.00
  E gastos de educação de 200.00
  Quando o cálculo é solicitado para data 2025-06-08
  Então o sistema deve retornar os modelos completo e simplificado
  E um dos modelos deve estar marcado como IsRecommended: true

Cenário: Aplicação de Reforma 2026
  Dado que a data de referência é 2026-02-01
  E a BaseValue é inferior ao floor de isenção
  Quando o cálculo é processado
  Então o TaxAmount final deve ser zero
  E deve constar o detalhe 'reforma_2026_isencao_total'
```

## 8. Prioridade MoSCoW

| Item | MoSCoW | Justificativa |
|------|--------|---------------|
| RF-01, RF-02 | Must | Core do microserviço. |
| RF-03 | Must | Necessário para cálculo real do modelo completo. |
| RF-05 | Should | Crítico para vigência futura, mas segue regras específicas de transição. |

## 9. Esclarecimentos

> Nenhuma sessão de dúvidas registrada ainda. Rode `/reversa-clarify` quando houver `[DÚVIDA]` pendente.

## 10. Lacunas

- 🟢 **CONFIRMADO** — A responsabilidade pela validação de procedência e tipo dos gastos com Saúde/Educação é do sistema de origem; este microserviço atua estritamente como motor de cálculo.
- 🟢 **CONFIRMADO** — Os Fatores A e B da transição 2026 são parâmetros externos carregados dinamicamente via `GetTableConfigs`, garantindo que a lógica de cálculo seja independente da gestão dos valores legislativos.

## 11. Histórico de alterações

| Data | Alteração | Autor |
|------|-----------|-------|
| 2026-06-08 | Versão inicial gerada por `reversa-writer` | reversa |
