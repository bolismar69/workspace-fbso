# Matriz de Indicadores de Negócio (KPIs) para Reporte Executivo (CFO)

```text
[ AUTONOMIA OPERACIONAL ] ──► [ GOVERNANÇA E COMPLIANCE ] ──► [ EFICIÊNCIA E SATISFAÇÃO ]
    Independência do Time           Rastreabilidade Fiscal           Agilidade nos Ajustes
    de Finanças                     e Controles Internos            e Experiência do Usuário
```

------------------------------
## 1. Dimensão: Autonomia Operacional

### KPI A1: Índice de Autonomia do Time Fiscal

* Objetivo de Negócio: Medir a redução da dependência do time técnico para ajustes fiscais de rotina (Vinculado ao Objetivo "Autonomia Operacional do Time de Finanças" do Charter) [INDEX].
* Fórmula de Cálculo: (Chamados técnicos para ajuste de alíquotas no mês atual / Média mensal de chamados nos 6 meses anteriores ao projeto) × 100
* Meta Recomendada: Redução ≥ 80% em relação à linha de base pré-projeto, a partir do 3º mês pós-implantação da Entrega 1 (Portal: Gestão Básica de Alíquotas).
* Frequência de Reporte: Mensal.

### KPI A2: Tempo Médio de Efetivação de Ajustes Fiscais

* Objetivo de Negócio: Medir a agilidade do time de Finanças na resposta a mudanças regulatórias ou necessidades de negócio (Vinculado ao Objetivo "Autonomia Operacional" do Charter) [INDEX].
* Fórmula de Cálculo: Tempo médio (em horas) entre a identificação da necessidade de ajuste de uma alíquota e sua efetivação no portal, medido da abertura do registro de alteração até a confirmação de vigência.
* Meta Recomendada: ≤ 4 horas para ajustes simples (alíquota única), ≤ 24 horas para ajustes que requerem aprovação em fluxo de duas etapas.
* Frequência de Reporte: Mensal.

------------------------------
## 2. Dimensão: Governança e Compliance Fiscal

### KPI G1: Cobertura de Trilha de Auditoria

* Objetivo de Negócio: Garantir que 100% das alterações em tabelas fiscais possuam registro completo de auditoria (Vinculado ao Objetivo "Governança e Rastreabilidade Societária" do Charter) [INDEX].
* Fórmula de Cálculo: (Número de alterações com registro completo de auditoria / Total de alterações realizadas no portal) × 100
* Meta Recomendada: 100% — toda alteração deve gerar trilha automaticamente; desvios configuram falha de conformidade.
* Frequência de Reporte: Mensal (auditoria amostral pela Controladoria).

### KPI G2: Índice de Conflitos Fiscais Prevenidos

* Objetivo de Negócio: Medir a eficácia das validações automáticas do portal em impedir configurações fiscais inválidas (Vinculado ao Objetivo "Blindagem contra Erros de Configuração Fiscal" do Charter) [INDEX].
* Fórmula de Cálculo: (Tentativas de cadastro bloqueadas por validações do portal / Total de tentativas de cadastro) × 100
* Meta Recomendada: Indicador de monitoramento — o valor absoluto de bloqueios deve reduzir ao longo do tempo, indicando maturidade do time fiscal. Bloqueios > 20% no primeiro mês são esperados; após 6 meses, devem ser < 5%.
* Frequência de Reporte: Mensal.

### KPI G3: Conformidade em Auditorias Fiscais

* Objetivo de Negócio: Assegurar que as alíquotas praticadas e registradas no portal estejam em conformidade com as publicações oficiais dos entes tributantes (Vinculado ao Critério de Sucesso "Precisão das Tabelas Fiscais" do Charter) [INDEX].
* Fórmula de Cálculo: (Número de alíquotas com divergência entre portal e publicação oficial / Total de alíquotas vigentes) — meta é zero divergências não justificadas.
* Meta Recomendada: Zero divergências não documentadas com justificativa de negócio aprovada pelo Comitê Fiscal.
* Frequência de Reporte: Trimestral (reconciliação programada).

------------------------------
## 3. Dimensão: Eficiência Operacional e Satisfação

### KPI E1: Índice de Cobertura Geográfica de Alíquotas

* Objetivo de Negócio: Monitorar a completude da base de alíquotas de IBS por município, garantindo que todos os municípios onde a companhia opera tenham alíquotas cadastradas (Vinculado ao Objetivo "Prontidão para o Período Híbrido" do Charter) [INDEX].
* Fórmula de Cálculo: (Municípios com operação comercial que possuem alíquota de IBS cadastrada / Total de municípios com operação comercial ativa) × 100
* Meta Recomendada: 100% dos municípios com faturamento ativo nos últimos 12 meses.
* Frequência de Reporte: Mensal.

### KPI E2: Satisfação do Time de Finanças (NPS Interno)

* Objetivo de Negócio: Medir a percepção de valor e usabilidade do portal junto aos usuários finais — analistas fiscais, contadores e controller (Vinculado ao Critério de Sucesso "Satisfação do Time de Finanças" do Charter) [INDEX].
* Fórmula de Cálculo: Net Promoter Score (NPS) calculado a partir da pergunta: "Em uma escala de 0 a 10, quanto você recomendaria o Portal de Gestão Tributária para um colega de outra empresa?"
* Meta Recomendada: NPS ≥ 70 (zona de excelência), com pontuação SUS (System Usability Scale) ≥ 75.
* Frequência de Reporte: Trimestral (após 90 dias de uso de cada entrega, com medição final consolidada após a Entrega 4).

### KPI E3: Taxa de Adoção do Portal

* Objetivo de Negócio: Garantir que o portal se torne efetivamente a ferramenta primária de gestão tributária, substituindo planilhas e processos manuais (Vinculado ao Critério de Sucesso "Adoção como Ferramenta Oficial" do Charter) [INDEX].
* Fórmula de Cálculo: (Operações de gestão de alíquotas realizadas exclusivamente no portal / Total de operações de gestão de alíquotas identificadas) × 100
* Meta Recomendada: ≥ 95% das operações administrativas fiscais originadas no portal após 6 meses do go-live da Entrega 4 (Portal Completo).
* Frequência de Reporte: Mensal.

------------------------------
## Template de Dashboard Executivo (Visão Mensal para o CFO)

| Categoria | Indicador de Negócio | Status Atual | Meta | Tendência | Impacto no Negócio |
|---|---|---|---|---|---|
| Autonomia | A1: Redução de Chamados Técnicos | — | ≥ 80% | — | Agilidade fiscal e liberação de capacidade técnica |
| Autonomia | A2: Tempo Médio de Ajuste Fiscal | — | ≤ 4h | — | Velocidade de resposta a mudanças regulatórias |
| Governança | G1: Cobertura de Trilha de Auditoria | — | 100% | — | Defesa em fiscalizações e conformidade societária |
| Governança | G3: Divergências Portal × Publicação Oficial | — | Zero | — | Risco de autuações e erros de precificação |
| Eficiência | E1: Cobertura Geográfica (IBS) | — | 100% | — | Cobertura de operações interestaduais |
| Satisfação | E2: NPS Interno (Time Fiscal) | — | ≥ 70 | — | Engajamento e retenção do time de Finanças |

------------------------------

# 99. LISTA DE BIBLIOTECAS

- Dependencies:

```bash
npm install @radix-ui/react-accordion @radix-ui/react-checkbox @radix-ui/react-dialog @radix-ui/react-dropdown-menu @radix-ui/react-hover-card @radix-ui/react-label @radix-ui/react-navigation-menu @radix-ui/react-popover @radix-ui/react-radio-group @radix-ui/react-select @radix-ui/react-slider @radix-ui/react-slot @radix-ui/react-switch @radix-ui/react-tabs @radix-ui/react-toast @radix-ui/react-tooltip @internationalized/date date-fns@3.6.0 react-day-picker@8.10.1 recharts @react-aria/datepicker @react-stately/datepicker
```

- Tailwind CSSv3.4.0+ and @tailwindcss/formsv0.5.10:

```bash
npm install tailwindcss@3.4.17 -D @tailwindcss/forms
```

- site com componentes tailwind:

```web

https://blocks.tremor.so/getting-started

```



---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: finance-metrics-quickref, agile-ba-practices.*
