# Project Charter: PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Produto** | Plataforma Shield — Identidade e Acesso Centralizada |
| **Data de Elaboração** | 07/08/2026 |
| **Versão** | 1.0 — Documento Inicial (WATERFALL v2.0) |
| **Patrocinador** | Diretoria de Tecnologia — FBSO.ORG |
| **Metodologia** | WATERFALL |
| **Status** | COMPLIANCE |

---

## O que é um Project Charter?

O **Project Charter (Termo de Abertura do Projeto)** é o documento fundador que autoriza formalmente a existência de um projeto. No nosso pipeline WATERFALL, ele é o **start do projeto** — o primeiro artefato criado, quando muitas informações ainda não existem.

### O que o Charter CONTÉM (e o que NÃO contém)

| ✅ Contém | ❌ NÃO Contém |
|---|---|
| Problema de negócio e justificativa | Datas absolutas de entrega |
| Escopo macro (In/Out) | Cronograma detalhado |
| Entregas em linguagem de negócio | Especificações técnicas |
| Critérios de aceitação de negócio | Arquitetura ou stack técnica |
| Stakeholders e RACI macro | Orçamento detalhado por recurso |
| Orçamento estimado (Budget/Limite) | Plano de testes |
| Marcos com referências temporais de negócio | Casos de uso |

### Conexão com o Pipeline

- **DOWNSTREAM:** Alimenta TODOS os 21 documentos posteriores como UPSTREAM_DOCS raiz
- **002-STAKEHOLDER-MAP:** A Seção 5 deste Charter é uma versão simplificada — o registro completo está no Stakeholder Map
- **GATE UPSTREAM (ROM ±50%):** Fornece as justificativas de negócio para o Comitê de Governança

---

### 1. Declaração do Problema (Problem Statement)

A FBSO.ORG opera um portfólio de produtos digitais para o setor educacional — Gestão Escolar, Comunidades de Ensino, Portal da Reforma. Cada produto gerencia seus próprios usuários e senhas de forma independente. Isso gera:

- **Risco de vazamento de dados entre clientes:** sem uma camada centralizada de segurança, dados de uma escola podem ser acessados indevidamente por outra. Um incidente desse tipo compromete a reputação da empresa e pode gerar sanções regulatórias (LGPD).
- **Experiência fragmentada para os clientes:** cada escola passa por um processo manual e demorado de configuração de acessos, sem um portal unificado.
- **Custo crescente de manutenção:** cada produto duplica funcionalidades de login, recuperação de senha e controle de permissões.
- **Impossibilidade de crescer o portfólio:** lançar um novo produto significa implementar autenticação do zero novamente — atrasando o time-to-market.

**O custo de não resolver:** risco regulatório (LGPD), perda de competitividade, e inviabilidade de escalar o ecossistema de produtos para dezenas ou centenas de escolas.

---

### 2. Propósito do Projeto (Project Purpose)

O **Projeto Shield** cria a **Plataforma de Identidade e Segurança da FBSO.ORG** — um produto interno que será a porta de entrada única para todos os sistemas da empresa.

O que ele entrega para o negócio:

- **Reconhecimento automático do cliente:** um diretor acessa o sistema da sua escola e a plataforma já sabe quem ele é e para onde direcioná-lo, sem qualquer configuração manual.
- **Isolamento total entre clientes:** os dados da Escola A são completamente inacessíveis para a Escola B. Cada cliente opera em seu próprio ambiente seguro.
- **Login protegido contra ataques:** as credenciais dos usuários jamais ficam expostas no navegador, eliminando os riscos mais comuns de roubo de sessão (XSS/CSRF).
- **Velocidade de resposta:** a validação de identidade acontece em menos de 15 milissegundos — imperceptível para o usuário final.
- **Porta única para todos os produtos:** qualquer novo sistema da FBSO.ORG usará a mesma plataforma de identidade.

#### 2.1 Visão de Longo Prazo

- **Onboarding self-service:** uma nova escola poderá ser provisionada na plataforma em até 4 horas (hoje são dias).
- **Login unificado (SSO):** um usuário fará login uma vez e acessará todos os produtos FBSO.
- **Futuro:** integração com identidades que as escolas já usam (Google for Education, Microsoft 365 for Education, GOV.BR).

---

### 3. Escopo (Scope)

#### 3.1 Dentro do Escopo (In Scope)

1. **Plataforma de Acesso Unificada** — Conjunto de funcionalidades de autenticação disponível para todos os produtos do ecossistema.

#### 3.2 Fora do Escopo (Out of Scope)

- Desenvolvimento de funcionalidades de negócio dos produtos que consumirão a plataforma
- Integração com provedores externos de identidade (Google, Microsoft, GOV.BR) — fase futura
- Desenvolvimento de interfaces de usuário das aplicações cliente

---

### 4. Entregas (Deliverables) & Critérios de Aceitação

> **REGRA: NUNCA incluir Data-Alvo nesta seção.** Datas absolutas não existem no momento da criação do Charter.

| # | Entrega | Critérios de Aceitação de Negócio |
|---|---|---|
| D1 | Plataforma Única de Acesso | Nova platafora Única de Acesso totalmente funcional |
| D2 | Todos os sistemas da empresa integrados com a nova Plataforma Única de Acesso | Todos os sistemas da empresa conectadas na nova Plataforma Única de Acesso totalmente funcionais |

---

### 5. Partes Interessadas e Matriz RACI (Stakeholders & RACI)
> O registro completo e detalhado dos stakeholders está no documento 🔗 `002-STAKEHOLDER-MAP-PRJ-TEC-2026-0004-PROJETO-SHIELD.md`

---

### 6. Critérios de Sucesso (Success Criteria)

| # | Critério | Como Medimos | Meta |
|---|---|---|---|
| C1 | Segurança entre Clientes | Tentativas de acessar dados de outro cliente | 100% bloqueadas |
| C2 | Proteção de Credenciais | Verificação de que credenciais não aparecem em lugar nenhum fora do ambiente seguro | 100% dos acessos protegidos |
| C3 | Velocidade de Resposta | Tempo para validar identidade do usuário | Abaixo de 15ms |
| C4 | Capacidade de Atender Picos | Simular o horário de entrada de todas as escolas simultaneamente | Sem falhas no sistema |
| C5 | Cobertura a Ataques Cibernéticos | Principais categorias de ataque cobertas | 100% dos cenários |
| C6 | Tempo para Adicionar Novo Cliente | Processo completo de ativação de uma nova escola | Abaixo de 4 horas |
| C7 | Disponibilidade da Plataforma | Tempo em que o sistema de login fica fora do ar | Máximo 0.1% do tempo (99.9%) |
| C8 | Rastreabilidade de Acessos | Registro de cada tentativa de acesso (bem-sucedida ou não) | 100% dos eventos |

---

### 7. Premissas (Assumptions)

1. Não iremos gerar nenhum impacto nos processos operacionais de nossos clientes.
2. As aplicações corporativas serão adaptadas para consumir a nova plataforma de login.

---

### 8. Restrições (Constraints)

- **Prazo de Negócio:** Solução deve estar operacional antes da expansão para novas escolas no próximo ciclo letivo
- **Orçamento:** Limitado ao aprovado pela Diretoria de Tecnologia (ver Seção 11.1)
- **LGPD:** Precisaremos seguir todas as diretrizes da LGPD.

---

### 9. Riscos de Alto Nível (High-Level Risks)

| Risco | Probabilidade | Impacto | Mitigação |
|---|---|---|---|
| Parada total de todos os sistemas comerciais | Baixa | Alto | Projeto só irá para produção com 100% de acertividade |

---

### 10. Marcos do Projeto (Project Milestones)

> **REGRA:** A coluna "Referência Temporal" usa datas contextuais de negócio, NUNCA datas absolutas (dd/mm/aaaa).

| Marco | Referência Temporal | Critério de Conclusão |
|---|---|---|
| M1: Kickoff | Dia 0 | Aprovação de inicio do projeto pelo Sponsor e Stakeholders |
| M2: Finalização dos Documentos Funcionais | Primeira quinzena | Documentos (001-PROJECT-CHARTER, BRD, FRD, 015-RTM-FASE-1) entregues |
| M3: Estimativa High-Level e Decisão Go/No-Go | Segunda quinzena | Estimativa High Level (ROM ±50%) entregue, e, decisão de Go para o projeto, e aprovação de Reserva de Budget |
| M4: Estimativa detalhada, Orçamento e Cronograma | Terceira quinzena | Entrega da estimativa detalhada, e, efetivação de aprovação Real do Budget |
| M5: Desenvolvimento | Conforme Cronograma  | Entrega da estimativa detalhada, e, efetivação de aprovação Real do Budget |
| M6: Projeto Entregue e Homologação Concluída | Conforme Cronograma  | Projeto Concluido e Homologado por todas partes interessadas |
| M7: Go-Live | Conforme Cronograma | Plataforma em produção, termo de aceite assinado |

---

### 11. Orçamento Estimado (Estimated Budget)

| Categoria | Estimativa |
|---|---|
| **Valor máximo de investimento pretendido (Budget/Limite)** | ~R$ 100.000 |

---

### 12. Plano de Comunicação (Communication Plan)
> Vide no documento 🔗 `002-STAKEHOLDER-MAP-PRJ-TEC-2026-0004-PROJETO-SHIELD.md`.

---

### 13. Governança (Governance)

- **Comitê de Projeto:** Reúne-se quinzenalmente para aprovar mudanças de escopo, orçamento e riscos.
- **Gestão de Projeto:** Responsável por cronograma, dependências, riscos, custos e relatórios de status. Remove bloqueios da equipe.
- **Produto:** Responsável por definir e priorizar funcionalidades, validar critérios de aceite, e ser a voz do cliente interno.
- **Controle de Mudanças:** Alterações de escopo, prazo ou orçamento acima de 10% exigem aprovação formal do Comitê.
- **Portal de Segurança:** Nenhuma liberação para produção sem revisão de segurança aprovada pelo Comitê.

---

### 14. Aprovações (Approvals)

| Nome | Papel | Data | Assinatura |
|---|---|---|---|
| `<Sponsor>` | Diretoria de Tecnologia — FBSO.ORG | Pendente | Pendente |
| `<Product Owner>` | Product Owner — Plataforma Shield | Pendente | Pendente |
| `<PMO>` | PMO Corporativo | Pendente | Pendente |
