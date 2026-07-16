---
title: "Requisitos — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["requirements", "ears", "moscow", "reference"]
---

# Especificação de Requisitos — TaxNexus Portal

## 1. Requisitos Funcionais (EARS)

### RF-01: Autenticação por CNPJ

**OBS-AUTH-001 — Login**
```
When o contribuinte submete CNPJ com 14+ dígitos no formulário de login,
the system shall liberar acesso ao simulador tributário.
```
📍 `src/App.tsx:9-14`

**OBS-AUTH-002 — CNPJ Inválido**
```
When o contribuinte submete CNPJ com menos de 14 dígitos,
the system shall exibir alerta "Por favor, insira um CNPJ válido."
```
📍 `src/App.tsx:10-14`

**OBS-AUTH-003 — Logout**
```
When o contribuinte clica em "Sair / Trocar CNPJ",
the system shall retornar à tela de login e limpar o estado de autenticação.
```
📍 `src/App.tsx:47`

---

### RF-02: Simulação Tributária

**OBS-SIM-001 — Parâmetros de Entrada**
```
The system shall enviar ao backend: CNPJ, NCM, código IBGE do município e saldo remanescente.
```
📍 `src/hooks/useTaxService.ts:43-52`, `src/components/TaxSimulator.tsx:31-36`

**OBS-SIM-002 — Endpoint de Cálculo**
```
The system shall chamar POST http://localhost:8080/v1/tax/calculate com Content-Type application/json.
```
📍 `src/hooks/useTaxService.ts:46-52`

**OBS-SIM-003 — Resposta do Cálculo**
```
When a API retorna status 200,
the system shall exibir cards comparativos (Legado vs. Reforma) e gráfico de transição.
```
📍 `src/components/TaxSimulator.tsx:38-54`

**OBS-SIM-004 — Botão Desabilitado**
```
While não há cidade selecionada ou requisição está em andamento,
the system shall desabilitar o botão "SIMULAR REFORMA TRIBUTÁRIA".
```
📍 `src/components/TaxSimulator.tsx:96`

**OBS-SIM-005 — Loading State**
```
While a requisição está em andamento,
the system shall exibir texto "Processando..." no botão de simulação.
```
📍 `src/components/TaxSimulator.tsx:99`

---

### RF-03: Visualização de Dados

**OBS-VIZ-001 — Cards Comparativos**
```
When dados do cálculo estão disponíveis,
the system shall exibir dois cards lado a lado:
- Sistema Atual (Legado): PIS, COFINS, ICMS, IPI
- Pós-Reforma (IVA Dual): CBS, IBS, Imposto Seletivo, IPVA+ITCMD
```
📍 `src/components/TaxSimulator.tsx:106-127`

**OBS-VIZ-002 — Gráfico de Transição**
```
When dados do cálculo estão disponíveis,
the system shall exibir gráfico de barras empilhadas Recharts com:
- 2026 (Transição): valores CBS + IBS
- 2027 (Projeção): valores CBS*1.02 + IBS*1.02
```
📍 `src/components/TaxSimulator.tsx:130-154`

**OBS-VIZ-003 — Formatação de Moeda**
```
The system shall formatar todos os valores monetários em BRL (pt-BR) usando toLocaleString.
```
📍 `src/components/TaxSimulator.tsx:110-113`

---

### RF-04: Seletores de Localidade

**OBS-LOC-001 — Estados Disponíveis**
```
The system shall disponibilizar dropdown com SP (35), RJ (33) e AM (13).
```
📍 `src/components/TaxSimulator.tsx:18-22`

**OBS-LOC-002 — Cidades por Estado**
```
When um estado é selecionado,
the system shall carregar dropdown de cidades correspondente (apenas capitais: SP→São Paulo, RJ→Rio de Janeiro, AM→Manaus).
```
📍 `src/components/TaxSimulator.tsx:24-28`

**OBS-LOC-003 — Cidade Desabilitada sem Estado**
```
While nenhum estado está selecionado,
the system shall desabilitar o dropdown de cidade.
```
📍 `src/components/TaxSimulator.tsx:76`

---

## 2. Requisitos Não-Funcionais

### RNF-01: Performance

| ID | Requisito | Evidência |
|---|---|---|
| RNF-PERF-001 | Dev server Vite com HMR na porta 5173 | `package.json:7` |
| RNF-PERF-002 | Build de produção: `tsc -b && vite build` | `package.json:8` |
| RNF-PERF-003 | Docker multi-stage para imagem final enxuta (Nginx alpine) | `Dockerfile:1,20` |

### RNF-02: Segurança

| ID | Requisito | Evidência |
|---|---|---|
| RNF-SEC-001 | Pseudo-auth client-side (CNPJ length >= 14) — sem verificação backend | `App.tsx:10` |
| RNF-SEC-002 | Sem autenticação na chamada à API | `useTaxService.ts:46-52` |
| RNF-SEC-003 | Nginx configurado sem `server_tokens` (compliance parcial) | `Dockerfile:22` |

### RNF-03: Deploy

| ID | Requisito | Evidência |
|---|---|---|
| RNF-DEP-001 | Container expõe porta 5173 | `Dockerfile:32`, `nginx.conf:2` |
| RNF-DEP-002 | SPA fallback configurado (`try_files $uri /index.html`) | `nginx.conf:5` |
| RNF-DEP-003 | Build com `npm ci` (lock file determinístico) | `Dockerfile:11` |

### RNF-04: Qualidade de Código

| ID | Requisito | Evidência |
|---|---|---|
| RNF-QUAL-001 | TypeScript strict mode | `tsconfig.app.json` |
| RNF-QUAL-002 | ESLint com regras React Hooks + typescript-eslint | `eslint.config.js` |
| RNF-QUAL-003 | Sem testes automatizados | Ausência de `*.test.ts` ou `*.spec.ts` |

---

## 3. MoSCoW Prioritization

### Must Have (MVP — Implementado)
- ✅ RF-01: Login por CNPJ
- ✅ RF-02: Simulação tributária com chamada à API
- ✅ RF-03: Cards comparativos + gráfico
- ✅ RF-04: Seletores de UF/cidade

### Should Have (Próximo)
- 🔄 Autenticação real (JWT/OAuth)
- 🔄 URL do backend configurável por ambiente
- 🔄 Tratamento de erros com feedback ao usuário
- 🔄 Timeout e AbortController nas chamadas fetch

### Could Have (Futuro)
- 📋 Histórico de simulações por CNPJ
- 📋 Comparação lado a lado de múltiplos cenários
- 📋 Exportação PDF/CSV do resultado
- 📋 Testes unitários e de integração

### Won't Have (Agora)
- ❌ Multi-tenancy
- ❌ White-label
- ❌ Integração com ERPs
- ❌ Recomendações por IA

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
