# FEATURE - FEAT-EP-0004-0002: Onboarding Guiado de Primeiro Acesso

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0004-0002 — Onboarding Guiado de Primeiro Acesso |
| **Épico** | [EP-0004 — Experiência do Cliente e Autoatendimento](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | **Anterior:** [FEAT-EP-0004-0001 — Autenticação](../FEAT-EP-0004-0001-autenticacao-recuperacao-senha.md) | **Próximo:** [FEAT-EP-0004-0003 — Dashboard do Cliente](../FEAT-EP-0004-0003-dashboard-cliente.md)

**Requisitos BRD Vinculados:** [BR-06](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Portal do Cliente), [BR-07](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Onboarding Guiado)

---

## Objetivo de Negócio
Conduzir o cliente por um fluxo simples e guiado no primeiro acesso, garantindo que ele configure o essencial para começar a usar a plataforma sem precisar de ajuda do suporte.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-040 | Como **Cliente no primeiro acesso**, quero ser recebido por um fluxo guiado de onboarding que me conduza passo a passo pelas configurações iniciais obrigatórias | • Ao detectar primeiro login, sistema inicia automaticamente o onboarding • Barra de progresso visível (Passo 1 de 4, Passo 2 de 4...) • Não é possível pular etapas obrigatórias • Cliente pode salvar e continuar depois (retoma de onde parou) |
| US-041 | Como **Cliente no onboarding**, quero confirmar e complementar meus dados cadastrais (razão social, nome fantasia, segmento) para garantir que as informações estão corretas | • Dados pré-preenchidos com informações fornecidas pelo time FBSO.ORG • Cliente confirma ou edita cada campo • Avançar para o próximo passo salva os dados |
| US-042 | Como **Cliente no onboarding**, quero cadastrar minha primeira Unidade de Negócio (CNPJ matriz, regime tributário, endereço) para começar a operar na plataforma | • Formulário de cadastro de Unidade de Negócio integrado ao fluxo • Campos: CNPJ, razão social, regime tributário (Simples, Lucro Real, Lucro Presumido), endereço • Validação de CNPJ informa se formato é válido (validação de dígitos) |
| US-043 | Como **Cliente no onboarding**, quero visualizar um resumo do meu plano contratado (nome do plano, módulos incluídos, valor) para entender o que está disponível para mim | • Card com informações do plano: nome, módulos incluídos (ícones e nomes), valor mensal • Informação de que novos módulos podem ser contratados futuramente • Botão "Começar a usar" para finalizar o onboarding |
| US-044 | Como **Cliente**, quero ser recebido com uma tela de boas-vindas após concluir o onboarding, com orientações sobre os próximos passos | • Tela de boas-vindas com: saudação personalizada, resumo do que foi configurado • Sugestões de próximos passos: "Convide seu time", "Cadastre seus produtos" • Botão "Ir para o Portal" que leva ao dashboard do cliente |

## Regras de Negócio

- **RN-FEAT-EP-0004-0002-0001:** Onboarding é obrigatório no primeiro acesso; não pode ser pulado
- **RN-FEAT-EP-0004-0002-0002:** Primeira Unidade de Negócio cadastrada no onboarding é automaticamente definida como Matriz
- **RN-FEAT-EP-0004-0002-0003:** Onboarding só é considerado concluído quando todos os passos obrigatórios são finalizados
- **RN-FEAT-EP-0004-0002-0004:** Tenant só muda para status "Ativo" após conclusão do onboarding

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-06** | Portal do Cliente com Autenticação | [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) / J1: Primeiro acesso e onboarding | **FEAT-EP-0004-0002** — Onboarding Guiado |
| **BR-07** | Onboarding Guiado de Primeiro Acesso | [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) / J1: Primeiro acesso e onboarding | **FEAT-EP-0004-0002** — Onboarding Guiado |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
