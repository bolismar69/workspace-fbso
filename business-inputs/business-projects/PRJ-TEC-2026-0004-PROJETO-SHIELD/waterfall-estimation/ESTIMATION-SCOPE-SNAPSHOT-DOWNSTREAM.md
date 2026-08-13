# SCOPE SNAPSHOT — DOWNSTREAM/REFINEMENT: PRJ-TEC-2026-0004-PROJETO-SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Estimativa Vinculada** | WATERFALL-ESTIMATION-DOWNSTREAM-PERT.md v1.0 |
| **Data de Congelamento** | 04/08/2026 |
| **Versão** | 1.0 |
| **Modo** | DOWNSTREAM/REFINEMENT |

---

## SCOPE SNAPSHOT — DOWNSTREAM/REFINEMENT
O termo Scope Snapshot — Downstream/Refinement refere-se no contexto de gestão de projetos, engenharia de software ou arquitetura de sistema sao congelamento e desdobramento do escopo para etapas posteriores (downstream) e seu respectivo detalhamento/refinamento (refinement)

### O Que Significa Cada Parte
- **Scope Snapshot (Retrato do Escopo):** É uma foto ou um registro oficial do estado atual do escopo em um dado momento. Ele define limites, entregáveis e requisitos fixados para evitar mudanças descontroladas.
- **Downstream (Etapas Posteriores):** Refere-se às fases seguintes do fluxo de trabalho ou cadeia de valor que vão receber esse escopo já aprovado (por exemplo, equipes de desenvolvimento, testes ou implantação).
- **Refinement (Refinamento):** É o processo de quebrar itens grandes em partes menores, esclarecer dúvidas técnicas e adicionar o nível de detalhe necessário para a execução prática.

### Aplicação Prática
- **Congelamento de Fase:** O escopo é validado e registrado (snapshot).
- **Transição (Hand-off):** Os dados detalhados são repassados para as equipes de ponta frente (downstream).
- **Ajustes Contínuos:** O refinamento garante que os requisitos mantenham-se claros conforme o projeto avança.

---

### 1. Pacotes EAP Estimados (22 pacotes — 929,2h PERT)

| ID EAP | Pacote de Trabalho | Grupo | Fase WATERFALL | Σ E (h) |
|--------|-------------------|-------|---------------|---------|
| 1.1.1 | Provisionamento do Cluster (DOKS) | 1.1 Infraestrutura | Design | 56,7 |
| 1.1.2 | Rede e Segurança (Istio mTLS) | 1.1 Infraestrutura | Design | 46,7 |
| 1.1.3 | API Gateway (Kong) | 1.1 Infraestrutura | Design | 43,0 |
| 1.1.4 | Pipeline CI/CD + GitOps | 1.1 Infraestrutura | Design | 28,3 |
| 1.2.1 | Provisionamento Multi-Cliente (Keycloak) | 1.2 Motor de Identidade | Design | 76,0 |
| 1.2.2 | Temas e Fluxos OIDC | 1.2 Motor de Identidade | Design | 37,3 |
| 1.2.3 | Integração Cloudflare | 1.2 Motor de Identidade | Design | 19,7 |
| 1.3.1 | Reconhecimento de Domínio | 1.3 Portal de Acesso | Desenvolvimento | 45,7 |
| 1.3.2 | Login e Logout | 1.3 Portal de Acesso | Desenvolvimento | 79,3 |
| 1.3.3 | Sessão e Cookies | 1.3 Portal de Acesso | Desenvolvimento | 47,0 |
| 1.3.4 | Cache (Host→Realm, JWKS) | 1.3 Portal de Acesso | Desenvolvimento | 24,0 |
| 1.4.1 | Modelagem e Isolamento RLS | 1.4 Isolamento de Dados | Desenvolvimento | 57,7 |
| 1.4.2 | Pool de Conexões | 1.4 Isolamento de Dados | Desenvolvimento | 12,3 |
| 1.5.1 | Métricas e Dashboards | 1.5 Observabilidade | Desenvolvimento | 38,7 |
| 1.5.2 | Agregação de Logs | 1.5 Observabilidade | Desenvolvimento | 20,0 |
| 1.5.3 | Rastreamento Distribuído | 1.5 Observabilidade | Desenvolvimento | 20,3 |
| 1.6.1 | Testes de Isolamento | 1.6 Testes | Testes | 52,7 |
| 1.6.2 | Testes de Carga | 1.6 Testes | Testes | 41,3 |
| 1.6.3 | Testes de Segurança | 1.6 Testes | Testes | 57,3 |
| 1.6.4 | Homologação e Aceite | 1.6 Testes | Testes | 36,7 |
| 1.7.1 | Documentação de API | 1.7 Entrega | Deploy | 12,3 |
| 1.7.2 | Manuais Operacionais | 1.7 Entrega | Deploy | 16,7 |
| 1.7.3 | Go-Live | 1.7 Entrega | Deploy | 25,7 |

### 2. Exclusões Explícitas

| Item | Motivo | Fonte |
|------|--------|-------|
| MS de negócio | Apenas integração | 01-Charter §3.1 |
| Migração de dados | Fora do escopo | 01-Charter §3 |
| Portal admin | Não previsto | 01-Charter §3.2 |

### 3. Independência

> Esta estimativa PERT foi calculada exclusivamente a partir de 03-SRS, 04-RTM, 07-LLD e 11-EAP/WBS. NENHUM valor do ROM upstream foi usado como baseline. O escopo congelado reflete exclusivamente os 22 pacotes EAP/WBS do projeto.

### 4. Versões dos Documentos Fonte

| Documento | Versão | Data | Status |
|-----------|--------|------|--------|
| 03-SRS | v1.0 | 03/08/2026 | COMPLIANCE |
| 04-RTM | v1.0 | 03/08/2026 | COMPLIANCE |
| 07-LLD | v3.0 | 03/08/2026 | COMPLIANCE |
| 11-EAP/WBS | v1.0 | 03/08/2026 | COMPLIANCE |

### 5. Premissas de Escopo

| Premissa | Impacto se Inválida |
|----------|---------------------|
| 22 pacotes EAP cobrem 100% do escopo | Escopo adicional requer reestimativa |
| Stack 100% baseline corporativa | Desvio requer justificativa + replanejamento |
| Equipe dedicada 4 pessoas × 6h/dia | Se shared, prazo estica proporcionalmente |
