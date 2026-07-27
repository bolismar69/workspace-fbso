# FEATURE - FEAT-EP-0002-0001: Cadastro e Ativação de Contas de Clientes

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0002-0001 — Cadastro e Ativação de Contas de Clientes |
| **Épico** | [EP-0002 — Gestão de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | **Próximo:** [FEAT-EP-0002-0002 — Gestão de Status](../FEAT-EP-0002-0002-gestao-status-tenant.md)

**Requisitos BRD Vinculados:** [BR-02](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Ativação e Gestão de Contas

---

## Objetivo de Negócio
Permitir que o time interno crie contas de clientes de forma estruturada, com todos os dados corporativos necessários, gerando automaticamente o convite de acesso ao portal.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-008 | Como **Administrador FBSO**, quero criar uma nova conta de cliente (Tenant) preenchendo razão social, nome fantasia e segmento de mercado para registrar o cliente na plataforma | • Formulário com campos obrigatórios: razão social, segmento • Campos opcionais: nome fantasia, observações • Ao salvar, Tenant é criado com status "Pendente Onboarding" • Sistema gera link único de ativação para o cliente |
| US-009 | Como **Administrador FBSO**, quero que o sistema envie automaticamente um e-mail de boas-vindas ao cliente com o link de ativação da conta após a criação do Tenant | • E-mail disparado automaticamente após criação do Tenant • E-mail contém link único e instruções de primeiro acesso • Link expira em 7 dias (renovável pelo administrador) |
| US-010 | Como **Administrador FBSO**, quero editar os dados cadastrais de um cliente (razão social, nome fantasia, segmento) para manter as informações sempre atualizadas | • Tela de edição acessível a partir da lista de contas • Campos editáveis: razão social, nome fantasia, segmento, observações • Alterações registradas no histórico de auditoria |
| US-011 | Como **Administrador FBSO**, quero reenviar o e-mail de ativação caso o cliente não tenha recebido ou o link tenha expirado | • Botão "Reenviar Convite" na tela de detalhes do Tenant • Disponível apenas para contas com status "Pendente Onboarding" • Gera novo link e novo prazo de 7 dias |

## Regras de Negócio

- **RN-FEAT-EP-0002-0001-0001:** Toda criação de Tenant gera registro de auditoria com: administrador responsável, data/hora, dados iniciais
- **RN-FEAT-EP-0002-0001-0002:** Razão social é validada como obrigatória; sistema alerta se já existir Tenant ativo com mesma razão social
- **RN-FEAT-EP-0002-0001-0003:** Link de ativação é único e de uso único por Tenant

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-02** | Ativação e Gestão de Contas | [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) / J1: Ativação de novo cliente | **FEAT-EP-0002-0001** — Cadastro e Ativação de Contas |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
