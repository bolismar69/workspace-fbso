# Cálculo de IRPF, Tarefas de Implementação

> Lista de tarefas para reimplementar o endpoint de cálculo de IRPF com base no comportamento do legado.

## Pré-requisitos
- [ ] Lib `taxnexus-individual-core-lib` integrada (models, repository, cache) 🟢
- [ ] MS de INSS acessível via URL de ambiente 🟢
- [ ] Banco de dados com tabelas `tax_definitions` e `tax_configs` populadas 🟢

## Tarefas

- [ ] T-01, Implementar Handler HTTP para captura de RequestID e Parsing de Body
  - Origem no legado: `handlers/tax_handler.go:21`
  - Critério de pronto: Endpoint responde 400 para JSON inválido e 200 para válido.
  - Confiança: 🟢

- [ ] T-02, Implementar carregamento de tabelas de alíquotas por TaxCode e Data
  - Origem no legado: `services/calculation_service.go:101`
  - Critério de pronto: Dados carregados via repository com base no `TaxCode="IRPF"`.
  - Confiança: 🟢

- [ ] T-03, Executar cálculos 'Completa' e 'Simplificada' em paralelo via Goroutines
  - Origem no legado: `services/calculation_service.go:108`
  - Critério de pronto: O handler aguarda ambos os resultados antes de responder.
  - Confiança: 🟢

- [ ] T-04, Implementar integração com microserviço de INSS (modelo completo)
  - Origem no legado: `services/calculation_service.go:159`
  - Critério de pronto: Falha no INSS não interrompe o cálculo principal (apenas Warn).
  - Confiança: 🟢

- [ ] T-05, Aplicar lógica de desconto simplificado (20% limitado ao teto)
  - Origem no legado: `services/calculation_service.go:225`
  - Critério de pronto: Cálculo respeita `simplified_discount_monthly_limit` das configs.
  - Confiança: 🟢

- [ ] T-06, Implementar regra de transição da Reforma 2026
  - Origem no legado: `services/calculation_service.go:275`
  - Critério de pronto: Isenção total ou redução adicional aplicada se data >= 2026-01-01.
  - Confiança: 🟢

- [ ] T-07, Lógica de recomendação do melhor modelo
  - Origem no legado: `services/calculation_service.go:132`
  - Critério de pronto: Campo `IsRecommended: true` atribuído ao modelo com menor imposto.
  - Confiança: 🟢

## Tarefas de Teste

- [ ] TT-01, Teste unitário do cálculo progressivo (faixas de alíquota)
- [ ] TT-02, Teste de integração com MS de INSS (simulado via WireMock)
- [ ] TT-03, Teste de carga para validar concorrência das goroutines
- [ ] TT-04, Verificação de regressão para as regras de 2026 (cenário de isenção vs redução)

## Ordem Sugerida
1. Iniciar pelo handler e carregamento de dados (T-01, T-02).
2. Implementar a estrutura paralela (T-03).
3. Desenvolver os núcleos de cálculo (T-04, T-05, T-06).
4. Finalizar com a recomendação (T-07).

## Lacunas Pendentes (🔴)
- Nenhuma lacuna bloqueante identificada. 🟢
