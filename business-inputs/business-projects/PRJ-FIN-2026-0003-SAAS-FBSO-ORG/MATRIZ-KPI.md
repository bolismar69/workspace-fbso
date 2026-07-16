# Matriz de Indicadores de Negócio (KPIs) para Reporte Executivo (Diretoria FBSO.ORG)

```text
[ ADOÇÃO E AUTONOMIA ] ──► [ OPERAÇÃO E GOVERNANÇA ] ──► [ SATISFAÇÃO E QUALIDADE ]
   Independência do Cliente        Eficiência Administrativa        Percepção de Valor
   no Autoatendimento              e Controle de Acessos            e Experiência do Usuário
```

------------------------------
## 1. Dimensão: Adoção e Autonomia do Cliente

### KPI A1: Taxa de Onboarding Autônomo

* **Objetivo de Negócio:** Medir a capacidade dos clientes de completar o fluxo de ativação sem intervenção do time FBSO.ORG (Vinculado ao Critério de Sucesso C3 do Project Charter) [INDEX].
* **Fórmula de Cálculo:** (Número de clientes que concluíram o onboarding sem abertura de chamado de suporte / Total de clientes que iniciaram o onboarding no período) × 100
* **Meta Recomendada:** ≥ 80% a partir do 3º mês pós-lançamento (M7 + 90 dias).
* **Frequência de Reporte:** Mensal.

### KPI A2: Tempo Médio de Onboarding

* **Objetivo de Negócio:** Medir a fluidez da experiência de primeiro acesso do cliente (Vinculado ao Critério de Sucesso C3 do Project Charter) [INDEX].
* **Fórmula de Cálculo:** Tempo médio (em minutos) entre o primeiro login do cliente e a conclusão do último passo do onboarding (tela de boas-vindas).
* **Meta Recomendada:** ≤ 10 minutos.
* **Frequência de Reporte:** Mensal.

### KPI A3: Taxa de Abandono de Onboarding

* **Objetivo de Negócio:** Identificar barreiras no fluxo de ativação que fazem o cliente desistir antes de concluir (Vinculado ao requisito BR-B03 do BRD) [INDEX].
* **Fórmula de Cálculo:** (Número de clientes que iniciaram mas não concluíram o onboarding em 7 dias / Total de clientes que iniciaram o onboarding) × 100
* **Meta Recomendada:** ≤ 15% de abandono.
* **Frequência de Reporte:** Mensal.

------------------------------
## 2. Dimensão: Operação e Governança Administrativa

### KPI O1: Tempo de Ativação de Nova Conta

* **Objetivo de Negócio:** Medir a agilidade do processo de ativação de cliente, da decisão comercial à conta operacional (Vinculado ao Critério de Sucesso C2 do Project Charter) [INDEX].
* **Fórmula de Cálculo:** Tempo médio (em minutos) entre a criação do Tenant pelo time administrativo e o disparo do e-mail de boas-vindas ao cliente.
* **Meta Recomendada:** ≤ 2 minutos (acionamento imediato via portal).
* **Frequência de Reporte:** Mensal.

### KPI O2: Cobertura de Trilha de Auditoria Administrativa

* **Objetivo de Negócio:** Garantir que 100% das ações administrativas possuam registro completo de auditoria (Vinculado aos requisitos BR-A02, BR-A04, BR-A05 do BRD) [INDEX].
* **Fórmula de Cálculo:** (Número de ações administrativas com registro completo de auditoria / Total de ações administrativas realizadas) × 100
* **Meta Recomendada:** 100% — toda ação de ativação, suspensão, alteração de plano e mudança de permissão deve gerar trilha automaticamente.
* **Frequência de Reporte:** Mensal (auditoria amostral).

### KPI O3: Incidentes de Vazamento entre Unidades de Negócio

* **Objetivo de Negócio:** Assegurar isolamento total de dados entre filiais de um mesmo cliente (Vinculado ao Critério de Sucesso C5 do Project Charter) [INDEX].
* **Fórmula de Cálculo:** Número absoluto de incidentes reportados de acesso indevido a dados de Unidade de Negócio não autorizada.
* **Meta Recomendada:** Zero incidentes.
* **Frequência de Reporte:** Reporte imediato em caso de incidente; sumário mensal com total acumulado.

### KPI O4: Tempo de Bloqueio de Acesso

* **Objetivo de Negócio:** Medir a eficácia da suspensão de acesso — quanto tempo entre a decisão administrativa e o bloqueio efetivo do usuário (Vinculado ao requisito BR-A02 do BRD) [INDEX].
* **Fórmula de Cálculo:** Tempo médio (em minutos) entre a alteração de status do tenant para "Suspenso" e o bloqueio efetivo do último usuário logado.
* **Meta Recomendada:** ≤ 5 minutos.
* **Frequência de Reporte:** Por evento, com sumário mensal.

------------------------------
## 3. Dimensão: Satisfação e Qualidade

### KPI S1: Satisfação do Time Interno (NPS)

* **Objetivo de Negócio:** Medir a percepção de valor e usabilidade do portal administrativo junto ao time FBSO.ORG (Vinculado ao Critério de Sucesso C6 do Project Charter) [INDEX].
* **Fórmula de Cálculo:** Net Promoter Score (NPS) calculado a partir da pergunta: "Em uma escala de 0 a 10, quanto você recomendaria o Portal Administrativo da FBSO Platform para um colega de outra empresa de SaaS?" NPS = % Promotores (9-10) − % Detratores (0-6).
* **Meta Recomendada:** NPS ≥ 50.
* **Frequência de Reporte:** Trimestral (após 90 dias de uso de cada entrega).

### KPI S2: Satisfação do Cliente com o Onboarding

* **Objetivo de Negócio:** Medir a qualidade da primeira experiência do cliente na plataforma [INDEX].
* **Fórmula de Cálculo:** Nota média (1-5) atribuída pelo cliente ao final do onboarding em pesquisa opcional: "Como foi sua experiência de ativação?"
* **Meta Recomendada:** Nota média ≥ 4,0 / 5,0.
* **Frequência de Reporte:** Mensal.

### KPI S3: Chamados de Suporte por Cliente Novo

* **Objetivo de Negócio:** Medir indiretamente a clareza do portal e a qualidade do onboarding — quanto menor o número de chamados, mais autônomo é o cliente [INDEX].
* **Fórmula de Cálculo:** Número médio de chamados de suporte por cliente nos primeiros 30 dias após ativação. Contabilizam-se todos os níveis de chamado (N1, N2, N3). Chamados de natureza não-relacionada ao portal (ex: dúvidas fiscais sobre módulos futuros) são excluídos da métrica.
* **Meta Recomendada:** ≤ 1 chamado por cliente nos primeiros 30 dias.
* **Frequência de Reporte:** Mensal.

------------------------------
## 4. Dimensão: Prontidão para o Futuro

### KPI P1: Cobertura de Catálogo Pré-Mapeamento

* **Objetivo de Negócio:** Medir o percentual de clientes que já possuem catálogo de produtos cadastrado, indicando prontidão para ativação futura do Tributali-Engine (Vinculado ao Critério de Sucesso C4 do Project Charter) [INDEX].
* **Fórmula de Cálculo:** (Número de tenants com pelo menos 1 produto cadastrado no catálogo / Total de tenants ativos) × 100
* **Meta Recomendada:** Indicador de monitoramento. Meta ≥ 50% ao final do 3º mês pós-lançamento.
* **Frequência de Reporte:** Mensal.

### KPI P2: Tempo para Ativar Primeiro Módulo-Produto

* **Objetivo de Negócio:** Medir a efetividade da arquitetura de produto — quanto tempo para acoplar um novo módulo sobre o Core (Vinculado ao Critério de Sucesso C4 do Project Charter) [INDEX].
* **Fórmula de Cálculo:** Tempo estimado (em sprints) entre a decisão de iniciar o desenvolvimento do Tributali-Engine e a integração funcional do módulo ao Core.
* **Meta Recomendada:** ≤ 1 sprint de desenvolvimento para iniciar a integração.
* **Frequência de Reporte:** Medição única na fase seguinte do programa.

------------------------------
## 5. Template de Dashboard Executivo (Visão Mensal para a Diretoria)

| Categoria | Indicador de Negócio | Status Atual | Meta | Tendência | Impacto no Negócio |
|---|---|---|---|---|---|
| Adoção | A1: Onboarding Autônomo | — | ≥ 80% | — | Redução de custo de suporte |
| Adoção | A2: Tempo Médio de Onboarding | — | ≤ 10 min | — | Experiência do cliente |
| Adoção | A3: Abandono de Onboarding | — | ≤ 15% | — | Perda de clientes no funil de ativação |
| Operação | O1: Tempo de Ativação de Conta | — | ≤ 2 min | — | Agilidade comercial |
| Operação | O2: Cobertura de Auditoria | — | 100% | — | Conformidade e rastreabilidade |
| Operação | O3: Incidentes de Vazamento | — | Zero | — | Risco legal e reputacional |
| Operação | O4: Tempo de Bloqueio de Acesso | — | ≤ 5 min | — | Segurança operacional |
| Satisfação | S1: NPS Time Interno | — | NPS ≥ 50 | — | Retenção e engajamento do time |
| Satisfação | S2: Satisfação com Onboarding | — | ≥ 4,0/5,0 | — | Primeira impressão do cliente |
| Satisfação | S3: Chamados por Cliente Novo | — | ≤ 1 | — | Autonomia do cliente |
| Prontidão | P1: Cobertura de Catálogo | — | ≥ 50% | — | Base para ativação do Tributali-Engine |
| Prontidão | P2: Tempo p/ Ativar Módulo | — | ≤ 1 sprint | — | Agilidade de expansão do produto |

------------------------------

## 6. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-13 | Criação inicial: 4 dimensões (Adoção, Operação, Satisfação, Prontidão), 10 KPIs | Time de Negócios |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: finance-metrics-quickref, agile-ba-practices.*
🔍 *Revisado pelo skill caveman-review em 15/07/2026. Ajustes aplicados: S1 unificado para NPS (sem conflito com Likert), referência BRD corrigida (BR-B03), dashboard template completo com 10 KPIs, metodologia S3 detalhada.*
