# Requirements Traceability Matrix (RTM): PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 02-BRD, 03-SRS |
| **Data de Elaboração** | 03/08/2026 |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## Matriz de Rastreabilidade Bidirecional

| Objetivo Charter | Requisito de Negócio (BRD) | Requisito Funcional (SRS) | Feature (SRS) | Status |
|:---|:---|:---|:---|:---|
| C1 — Segurança entre Clientes | REQ-01 — Reconhecimento automático do cliente | FR-01 — Extração de domínio | F-01 — Reconhecimento de Cliente | ✅ |
| C1 — Segurança entre Clientes | REQ-02 — Isolamento total entre clientes | FR-08 — Filtro de isolamento | F-04 — Isolamento de Dados | ✅ |
| C1 — Segurança entre Clientes | — | FR-12 — Suspensão de cliente | F-08 — Suspensão de Cliente | ✅ |
| C2 — Proteção de Credenciais | REQ-03 — Proteção de credenciais | FR-03 — Início de fluxo | F-02 — Login Protegido | ✅ |
| C2 — Proteção de Credenciais | REQ-03 | FR-04 — Troca e armazenamento seguro | F-02 — Login Protegido | ✅ |
| C2 — Proteção de Credenciais | — | FR-07 — Logout completo | F-03 — Portal de Sessão | ✅ |
| C3 — Velocidade de Resposta | REQ-04 — Portal padronizado | FR-03, FR-04, FR-05 | F-02, F-03 | ✅ |
| C3 — Velocidade de Resposta | REQ-05 — Resposta < 15ms | FR-05 — Consulta de perfil | F-03 — Portal de Sessão | ✅ |
| C4 — Capacidade de Atender Picos | REQ-06 — Suporte a picos | FR-10 — Métricas operacionais | F-06, F-09 — Monitoramento, Escalabilidade | ✅ |
| C4 — Capacidade | REQ-09 — Adaptação automática | FR-10 | F-09 — Escalabilidade | ✅ |
| C5 — Cobertura de Testes de Segurança | REQ-02, REQ-03 | FR-08, FR-04 | F-04, F-02 | ✅ |
| C6 — Tempo de Ativação | REQ-08 — Ativação < 4h | FR-02, FR-11 | F-01, F-07 — Ativação de Cliente | ✅ |
| C7 — Disponibilidade 99.9% | REQ-06, REQ-09 | FR-06, FR-10 | F-03, F-06, F-09 | ✅ |
| C8 — Rastreabilidade de Acessos | REQ-07 — Registro de auditoria | FR-09 — Registro de eventos | F-05 — Auditoria de Acessos | ✅ |
| — | REQ-10 — Experiência consistente | FR-05 — Consulta de perfil | F-03 — Portal de Sessão | ✅ |

---

## Cobertura Forward (Charter → BRD → SRS)

| Critério Charter | Tem REQ no BRD? | Tem FR no SRS? | Cobertura |
|:---|---:|:---|:---|
| C1 — Segurança entre Clientes | ✅ (REQ-01, REQ-02) | FR-01, FR-08, FR-12 | 100% |
| C2 — Proteção de Credenciais | ✅ (REQ-03) | FR-03, FR-04, FR-07 | 100% |
| C3 — Velocidade de Resposta | ✅ (REQ-04, REQ-05) | FR-03, FR-04, FR-05 | 100% |
| C4 — Capacidade de Atender Picos | ✅ (REQ-06, REQ-09) | FR-10 | 100% |
| C5 — Cobertura de Testes | ✅ (REQ-02, REQ-03) | FR-04, FR-08 | 100% |
| C6 — Tempo de Ativação | ✅ (REQ-08) | FR-02, FR-11 | 100% |
| C7 — Disponibilidade 99.9% | ✅ (REQ-06, REQ-09) | FR-06, FR-10 | 100% |
| C8 — Rastreabilidade | ✅ (REQ-07) | FR-09 | 100% |

**Forward Coverage: 8/8 (100%)** — Todos os objetivos do Charter possuem requisitos de negócio e funcionais vinculados.

---

## Cobertura Backward (SRS → BRD → Charter)

| FR (SRS) | Vinculado a REQ (BRD)? | Vinculado a OBJ (Charter)? | Status |
|:---|---:|---:|:---|
| FR-01 — Extração de domínio | ✅ REQ-01 | ✅ C1 | ✅ |
| FR-02 — Cache de mapeamento | ✅ REQ-01, REQ-08 | ✅ C1, C6 | ✅ |
| FR-03 — Início de fluxo | ✅ REQ-04 | ✅ C2, C3 | ✅ |
| FR-04 — Troca e armazenamento | ✅ REQ-03, REQ-04 | ✅ C2, C3, C5 | ✅ |
| FR-05 — Consulta de perfil | ✅ REQ-04, REQ-05, REQ-10 | ✅ C3 | ✅ |
| FR-06 — Renovação silenciosa | ✅ REQ-04 | ✅ C3, C7 | ✅ |
| FR-07 — Logout completo | ✅ REQ-04 | ✅ C2 | ✅ |
| FR-08 — Filtro de isolamento | ✅ REQ-02 | ✅ C1, C5 | ✅ |
| FR-09 — Registro de auditoria | ✅ REQ-07 | ✅ C8 | ✅ |
| FR-10 — Métricas operacionais | ✅ REQ-09 | ✅ C4, C7 | ✅ |
| FR-11 — Ativação de cliente | ✅ REQ-08 | ✅ C6 | ✅ |
| FR-12 — Suspensão de cliente | ✅ BR-08 | ✅ C1 | ✅ |

**Backward Coverage: 12/12 (100%)** — Todos os requisitos funcionais estão vinculados a requisitos de negócio e objetivos do Charter.

| Requisito BRD | Coberto por FR? | Status |
|:---|---:|:---|
| REQ-01 — Reconhecimento automático | ✅ FR-01, FR-02 | ✅ |
| REQ-02 — Isolamento total | ✅ FR-08 | ✅ |
| REQ-03 — Proteção de credenciais | ✅ FR-03, FR-04 | ✅ |
| REQ-04 — Portal padronizado | ✅ FR-03 a FR-07 | ✅ |
| REQ-05 — Resposta < 15ms | ✅ FR-05 | ✅ |
| REQ-06 — Suporte a picos | ✅ FR-10 | ✅ |
| REQ-07 — Registro de auditoria | ✅ FR-09 | ✅ |
| REQ-08 — Ativação < 4h | ✅ FR-02, FR-11 | ✅ |
| REQ-09 — Adaptação automática | ✅ FR-10 | ✅ |
| REQ-10 — Experiência consistente | ✅ FR-05 | ✅ |

**BRD Coverage: 10/10 (100%)**

---

## Sumário de Conformidade

| Métrica | Resultado |
|:---|---:|
| Objetivos Charter com cobertura FR | 8/8 (100%) |
| Requisitos BRD com cobertura FR | 10/10 (100%) |
| FRs vinculados a REQs | 12/12 (100%) |
| Órfãos (FR sem REQ/OBJ) | 0 |
| Órfãos (REQ sem OBJ) | 0 |
| Órfãos (OBJ sem cobertura) | 0 |
| **Status de Auditoria** | **✅ 100% Alinhado (Pass)** |

---

**[STATUS: SUCESSO]** — Matriz de rastreabilidade completa. Zero órfãos. 100% cobertura bidirecional.
