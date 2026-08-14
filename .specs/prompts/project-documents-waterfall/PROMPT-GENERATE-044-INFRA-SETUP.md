# PROMPT: GERADOR DE INFRASTRUCTURE & CLOUD DESIGN SETUP (044-INFRA-SETUP / IDD)
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Arquiteto de Infraestrutura e Cloud Sênior, especializado em topologia de rede, provisionamento IaC e resiliência, no contexto da metodologia WATERFALL.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 030-SAD, 035-HLD, 040-LLD]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["cloud-architect", "terraform-specialist", "senior-devops"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** o 035-HLD (topologia, componentes, integradores), o 040-LLD (contratos e requisitos de infra por componente) e o 030-SAD (decisões de arquitetura) — a infraestrutura materializa a topologia desenhada
3. **ORDEM DA ESTEIRA F3:** este documento executa como 4º passo da esteira (`040-LLD → 042-DATA-SETUP → 043-SEC-SETUP → 044-INFRA-SETUP → ...`)
4. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
5. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
6. Use o prefixo padronizado: **IDD-NN** (componentes de infraestrutura e cloud)
7. Aplique a tabela VOCABULÁRIO WATERFALL abaixo em todo o documento
8. Ao final, retorne `{DOC_PATH}` confirmando a criação

## VOCABULÁRIO WATERFALL (obrigatório — não usar vocabulário ágil)

| Termo ágil (PROIBIDO) | Equivalente WATERFALL (usar) |
|---|---|
| Epic | Pacote de trabalho da EAP (060-EAP-WBS) |
| Feature | Funcionalidade `FEAT-NN` (010-FRD) |
| User Story | Caso de Uso `UC-NN` (010-FRD) |
| Definition of Ready (DoR) | GATE de COMPLIANCE do documento de origem |
| Sprint | Ciclo de entrega (FASE 5 — EXECUÇÃO E CONSTRUÇÃO) |
| Product Backlog | 088-PRODUCT-BACKLOG-LIST (FASE 4) |

## Template de Fallback (7 Seções)

```
# Infrastructure & Cloud Design Setup (IDD): {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 030-SAD, 035-HLD, 040-LLD |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## Infrastructure & Cloud Design Setup (IDD)

O **IDD** é a especificação de infraestrutura e cloud do projeto: topologia física, provisionamento como código, resiliência, rede e custos. Ele materializa a topologia do 035-HLD em recursos concretos de nuvem/on-premise.

### O que contém

- **Topologia de Infraestrutura (IDD-NN):** VPC/rede, compute, storage e serviços gerenciados
- **Provisionamento:** IaC (Terraform) com módulos e state
- **Escalabilidade e Disponibilidade:** AZs, autoscaling e estratégia de DR
- **Rede e Conectividade:** DNS, conectores e integradores
- **Custos:** SKUs, sizing e estimativa mensal

### Conexão com o Pipeline

- **UPSTREAM:** Consome topologia do 035-HLD, requisitos de infra do 040-LLD e decisões do 030-SAD
- **DOWNSTREAM:** Alimenta 041-DEVOPS-SETUP (IaC no pipeline), 065-CRONOGRAMA-GANTT (provisionamento), 070-ORCAMENTO (custos cloud), 087-PLANO-CI-CD-AMBIENTES, 088-PRODUCT-BACKLOG-LIST e 100-MANUAIS-OPERACIONAIS

---

## 1. Topologia de Infraestrutura (IDD-NN)

| ID | Componente | Serviço | Região/AZ | Sizing | Origem (035-HLD) |
|----|------------|---------|-----------|--------|-------------------|
| IDD-01 | {VPC/rede} | {serviço cloud} | {região} | {tamanho} | {componente do HLD} |
| IDD-02 | {compute} | {K8s/VM/FaaS} | ... | ... | ... |
| IDD-03 | {storage} | {RDS/S3/Redis} | ... | ... | ... |

```mermaid
flowchart LR
    Internet["Internet"] --> WAF["WAF"] --> LB["Load Balancer"] --> K8s["Kubernetes Cluster"]
    K8s --> DB["Managed Database"]
    K8s --> Cache["Redis"]
```

---

## 2. Provisionamento (IaC)

| Item | Especificação |
|------|---------------|
| Tooling | {Terraform/Pulumi/CloudFormation} |
| Estrutura de Módulos | {módulos por domínio: network, compute, data} |
| State Management | {backend remoto, locking} |
| Drift Detection | {ferramenta e frequência} |

---

## 3. Escalabilidade e Disponibilidade

| Item | Especificação |
|------|---------------|
| Alta Disponibilidade | {AZs, réplicas, failover} |
| Autoscaling | {regras por componente} |
| Disaster Recovery | {RPO/RTO, estratégia de restauração} |

---

## 4. Rede e Conectividade

| Item | Especificação |
|------|---------------|
| DNS | {registros, gestão} |
| Integradores | {conectores, filas, eventos — do 035-HLD} |
| Segurança de Rede | {segurança alinhada ao 043-SEC-SETUP} |

---

## 5. Custos

| Componente (IDD-NN) | SKU | Estimativa Mensal | Observações |
|---------------------|-----|-------------------|-------------|
| IDD-01 | {SKU} | R$ {valor} | {premissas} |

> **NOTA:** O 070-ORCAMENTO consolida estes custos com o esforço do time.

---

## 6. Rastreabilidade

| Componente IDD | Origem (030/035/040) | Consumidores Previstos | Status |
|----------------|----------------------|------------------------|--------|
| IDD-01 | {componente do 035-HLD} | 041, 065, 070, 087, 100 | ✅ Vinculado |

> **REGRA DE OURO:** Nenhum componente de infraestrutura pode existir sem lastro na topologia do HLD (035) ou no design do LLD (040).

---

## 7. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Criação inicial a partir de SAD/HLD/LLD | Time de Engenharia |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se as 7 seções estiverem completas, toda a topologia do HLD tiver componente IDD correspondente, o diagrama cobrir o fluxo de rede principal, os custos forem estimados para todos os componentes, e a rastreabilidade não tiver órfãos.
