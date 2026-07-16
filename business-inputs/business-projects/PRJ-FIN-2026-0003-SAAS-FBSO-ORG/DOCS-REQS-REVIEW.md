# Docs Review: BRD + Epics + Features
## Revisor: Caveman | Data: 15/07/2026 | Escopo: 02-BUSINESS-REQUIREMENTS.md, 03-EPICS.md, 04-FEATURES.md

---

## Resumo

1119 linhas em 3 documentos. Boa estrutura geral. 22 achados: 4 contradições cross-doc, 5 desalinhamentos com charter revisado, 13 issues pontuais.

---

## 1. Problemas Cross-Documento (Afetam 2+ documentos)

### ❌ CROSS-01: RBAC — 3 ou 4 papéis?

| Documento | O que diz |
|-----------|-----------|
| **Charter (rev 1.1)** | MVP: 3 papéis (Admin, Gerente, Operador). Auditor documentado, fase futura. |
| **BRD L178** | "atribuir papéis (Admin, Gerente, Operador, Auditor)" — 4 papéis |
| **Epics L258** | "papéis de acesso: Admin do Tenant, Gerente de Unidade, Operador, Auditor" — 4 papéis |
| **Features L274 (RN10-01)** | Tabela lista 4 colunas: Admin Tenant, Gerente BU, Operador BU, Auditor |

**Severidade: Alta.** Charter é doc primário. BRD/Epics/Features desalinhados. `Fix:` Unificar para 3 papéis MVP + nota "Auditor previsto para fase futura" nos 3 documentos.

---

### ❌ CROSS-02: BR-B05 (Catálogo) — Should Have ou Must Have?

| Documento | Prioridade |
|-----------|-----------|
| **BRD L213** | Should Have |
| **Features L16** | Must Have |
| **Charter D7** | Entrega obrigatória (equivale a Must Have) |

**Severidade: Média.** `Fix:` Subir BR-B05 para Must Have no BRD. Catálogo é entrega D7 do charter.

---

### ❌ CROSS-03: App Switcher — componente visual ou menu?

| Documento | Descrição |
|-----------|-----------|
| **Charter (rev 1.1) D5** | "Menu de navegação entre módulos... expande para App Switcher visual quando houver 2+ produtos" |
| **Features US-049 (L421)** | "App Switcher visível mesmo com um único módulo... exibe nome do módulo ativo sem dropdown" |

**Severidade: Baixa.** Charter diz "menu", Features diz "App Switcher sem dropdown". Ambos descrevem mesma coisa com nomes diferentes. `Fix:` Padronizar terminologia. Sugestão: "Seletor de Módulo" na fase atual, "App Switcher" na futura.

---

### ❌ CROSS-04: Premissa A1 (visão de produto madura) — contradição com risco

| Documento | O que diz |
|-----------|-----------|
| **Charter (rev 1.1) A2** | "Visão de produto possui direção estratégica definida, mas reconhece-se que requisitos detalhados ainda estão em maturação" |
| **BRD L234 (A1)** | "A visão de produto multi-módulo está madura o suficiente para orientar o design do Core" — mantém redação antiga |

**Severidade: Média.** Charter já corrigiu essa premissa. BRD ainda tem versão otimista original. `Fix:` Atualizar A1 no BRD para alinhar com charter revisado.

---

## 2. BRD (02-BUSINESS-REQUIREMENTS.md)

### L18: 🔴 Frase única de 4 linhas (Problem Statement)
Charter revisado quebrou em 3 frases. BRD ainda tem versão longa original.
`Fix:` Quebrar igual charter: "Sem essa fundação, [...] Resultado: [...]"

### L35, L63: 🔴 Meta "imediato" sem baseline numérico
C2 e C3 usam "Ativação imediata via portal" e "Imediato (via portal)". Charter revisado define baseline (~2 dias úteis) e target (≤ 5 minutos).
`Fix:` Adicionar baseline e target numérico. Igual charter.

### L275: 🟡 R5 diz "3-4 papéis" — ambíguo
Mitigação do risco RBAC usa intervalo "3-4 papéis essenciais". Charter definiu 3.
`Fix:` "Começar com 3 papéis essenciais (Admin, Gerente, Operador)."

### L326-332: 🟡 Orçamento total "A definir"
Charter revisado adicionou seção de orçamento. BRD tem placeholder.
`Fix:` Remover "A definir", colocar estimativa ou referenciar planilha complementar.

### L302: 🔵 Métrica "Registros de auditoria sem falhas" duplica BR-NFR03
Seção 11 repete NFR de auditabilidade como métrica de qualidade. Redundante.
`Fix:` Consolidar. NFR já cobre; métrica pode referenciar.

### Faltas no BRD:
- 🟡 Sem menção a ambiente de homologação (charter seção 10)
- 🟡 Sem plano de comunicação (charter seção 11)
- 🟡 Sem Definição de Pronto (charter seção 13)

---

## 3. EPICS (03-EPICS.md)

### L46: 🟡 "EP-01 e EP-02 evoluem em paralelo" — contradiz milestones
Cronograma de épicos diz paralelismo inicial. Charter M2→M3 são sequenciais.
`Fix:` Verificar se paralelismo é real. Se sim, atualizar charter. Se não, corrigir Epics.

### L226, L258, L302: 🔴 4 papéis RBAC (ver CROSS-01)
Persona "Auditor" listada, 4 papéis definidos, mas charter MVP tem 3.
`Fix:` Mover Auditor para seção "Fora do Escopo" ou marcar como "Fase Futura".

### L423: 🟡 Persona "Contador (futuro)" — multi-tenant
Epics menciona contador que "gerencia múltiplos clientes" (visão multi-empresa). Fora do escopo desta fase.
`Fix:` Adicionar nota explícita: "Funcionalidade multi-tenant para contadores NÃO está no escopo desta fase."

### L142 (EP-02 Jornada 1): 🔵 Fluxo "venda consultiva" fora do escopo
Jornada menciona "Vendedor fecha contrato... gera Ordem de Serviço". Charter diz que não há comercialização nesta fase.
`Fix:` Ressalvar: "Quando a comercialização for ativada em fase futura..."

### L83 (EP-01): 🟡 "Exporta visão resumida para apresentação"
Jornada 2 menciona exportação. Fora do escopo do EP-01 (L113: "Exportação de relatórios em PDF/Excel — funcionalidade futura").
`Fix:` Remover menção a exportação da jornada ou adicionar nota de ressalva.

---

## 4. FEATURES (04-FEATURES.md)

### L274 (RN10-01): 🔴 Tabela de permissões inclui 4 papéis (ver CROSS-01)
Tabela lista coluna "Auditor" com permissões "Ver" em várias linhas.
`Fix:` Remover coluna Auditor ou marcá-la como "[Fase Futura]".

### L421 (US-049): 🟡 App Switcher com 1 módulo (ver CROSS-03)
Desalinhamento terminológico com charter revisado.
`Fix:` Renomear para "Seletor de Módulo" ou alinhar descrição com charter.

### L16 (F04-03): 🟡 Dashboard do Cliente é Should Have
Dashboard do cliente como Should Have em Features, mas charter D5 lista "Portal do Cliente" como entrega Must Have.
`Fix:` Esclarecer: D5 cobre autenticação + onboarding + menu. Dashboard do cliente (F04-03) é bônus se tempo permitir. Manter Should Have, adicionar nota.

### L332 (RN12-02): 🔵 Nome do módulo ativo no topo
Regra diz que usuário vê "nome do módulo ativo no topo, ao lado do logo". Na Fase 0, só existe "FBSO Platform" como módulo.
`Fix:` OK como está. Apenas verificar consistência com US-049.

### L426 (RN16-02): 🟡 Módulo placeholder "FBSO Platform"
Regra define módulo placeholder. Charter não menciona esse nome.
`Fix:` Documentar nome do módulo placeholder no charter ou glossário.

### Faltas nas Features:
- 🔵 Sem referência a DoD (charter seção 13) para critérios de conclusão de US
- 🔵 User stories não referenciam ambiente de homologação como critério de aceitação implícito

---

## 5. Matriz de Consistência: Charter (rev 1.1) × Docs

| Item do Charter | BRD | Epics | Features | Status |
|-----------------|-----|-------|----------|--------|
| RBAC 3 papéis MVP | ❌ (4) | ❌ (4) | ❌ (4) | Desalinhado |
| C2 com baseline numérico | ❌ ("imediato") | ✅ (KPIs OK) | ✅ (regras OK) | Parcial |
| D3/D7 artefatos concretos | ❌ (genérico) | ✅ | ⚠️ (parcial) | Parcial |
| Ambiente de homologação | ❌ (ausente) | ❌ (ausente) | ❌ (ausente) | Ausente |
| Plano de comunicação | ❌ (ausente) | ❌ (ausente) | ❌ (ausente) | Ausente |
| DoD | ❌ (ausente) | ❌ (ausente) | ❌ (ausente) | Ausente |
| Orçamento | ❌ ("A definir") | — | — | Desalinhado |
| Premissa A2 (visão madura) | ❌ (antiga) | — | — | Desalinhado |
| App Switcher × Menu | ✅ | ✅ | ⚠️ (termo) | Parcial |
| Out of Scope (migração) | ✅ | ✅ | ✅ | OK |
| Out of Scope (treinamento) | ✅ (NFR04) | ✅ | ✅ | OK |

---

## 6. Ações Recomendadas (Priorizadas)

### Alta (contradições que afetam escopo)

1. **Unificar RBAC para 3 papéis MVP** em BRD (L178), Epics (L258, L226), Features (L274-RN10-01). Adicionar "Auditor = fase futura" como nota.
2. **Adicionar baseline/target numérico em C2** no BRD (L35, L63). Copiar do charter: "baseline ~2 dias úteis, target ≤ 5 minutos."
3. **Alinhar premissa A1/A2** no BRD (L234) com charter revisado: "visão tem direção definida, requisitos em maturação."

### Média (desalinhamentos estruturais)

4. **Subir BR-B05 para Must Have** no BRD (L213). Catálogo é entrega D7.
5. **Resolver paralelismo EP-01/EP-02** (Epics L46): ou charter ganha paralelismo, ou Epics remove.
6. **Adicionar seções faltantes** em BRD: ambiente de homologação, plano de comunicação, DoD (igual charter).
7. **Remover "A definir" do orçamento** no BRD (L332). Referenciar planilha complementar.
8. **Adicionar nota de ressalva** na persona "Contador" (Epics L423): multi-tenant fora do escopo.
9. **Padronizar terminologia App Switcher** (Features L421 vs Charter D5): "Seletor de Módulo" na fase atual.

### Baixa (cosméticos e documentação)

10. **Quebrar frase longa** no BRD L18 (Problem Statement).
11. **Adicionar ressalva "fase futura"** na jornada de venda consultiva (Epics L142).
12. **Remover menção a exportação** da jornada EP-01 (Epics L83) ou adicionar nota.
13. **Documentar "FBSO Platform"** como nome do módulo placeholder no glossário do charter.
14. **Referenciar DoD** nas Features como checklist implícito de conclusão de US.

---

## Resumo: Nota Caveman

| Dimensão | BRD | Epics | Features |
|----------|-----|-------|----------|
| Alinhamento com Charter rev 1.1 | ⚠️ 6 divergências | ⚠️ 4 divergências | ⚠️ 4 divergências |
| Consistência interna | ★★★★☆ | ★★★★☆ | ★★★★★ |
| Cobertura de escopo | ★★★★★ | ★★★★★ | ★★★★★ |
| Clareza de critérios | ★★★☆☆ | ★★★★☆ | ★★★★☆ |
| Rastreabilidade | ★★★★★ | ★★★★★ | ★★★★★ |

**Nota geral do conjunto: 3.5/5** — Documentos bem estruturados, cobertura completa, mas 6 desalinhamentos com charter revisado (versão 1.1) precisam de correção. Maioria dos problemas é herança da versão 1.0 do charter — atualizar charter sem cascatear para BRD/Epics/Features gerou inconsistências.

---

> Revisão caveman: 22 achados em 3 documentos. 4 contradições cross-doc. 5 desalinhamentos com charter 1.1. 14 ações, 3 de alta prioridade.
