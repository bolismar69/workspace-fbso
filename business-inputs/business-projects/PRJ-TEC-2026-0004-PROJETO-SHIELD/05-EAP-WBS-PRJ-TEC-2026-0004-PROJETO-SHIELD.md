# EAP/WBS — Estrutura Analítica de Projeto: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 02-BRD |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. EAP Gráfica (3 níveis)

```
1. Plataforma Shield
├── 1.1 Infraestrutura e Ambiente Base
│   ├── 1.1.1 Provisionamento do Cluster (DOKS)
│   ├── 1.1.2 Configuração de Rede e Segurança (Istio mTLS)
│   ├── 1.1.3 API Gateway (Kong)
│   └── 1.1.4 Pipeline de Entrega (CI/CD + GitOps)
├── 1.2 Motor de Identidade por Cliente
│   ├── 1.2.1 Provisionamento Multi-Cliente (Keycloak Realms)
│   ├── 1.2.2 Configuração de Temas e Fluxos OIDC
│   └── 1.2.3 Integração com Proteção de Borda (Cloudflare)
├── 1.3 Portal de Acesso (Auth-BFF)
│   ├── 1.3.1 Reconhecimento de Cliente por Domínio
│   ├── 1.3.2 Fluxo de Login e Logout
│   ├── 1.3.3 Gestão de Sessão e Cookies Protegidos
│   └── 1.3.4 Cache de Mapeamento e Chaves
├── 1.4 Camada de Isolamento de Dados
│   ├── 1.4.1 Modelagem Multi-Cliente e Políticas de Isolamento
│   └── 1.4.2 Pool de Conexões e Otimização
├── 1.5 Observabilidade
│   ├── 1.5.1 Métricas e Dashboards
│   ├── 1.5.2 Agregação de Logs
│   └── 1.5.3 Rastreamento Distribuído
├── 1.6 Testes e Homologação
│   ├── 1.6.1 Testes de Isolamento entre Clientes
│   ├── 1.6.2 Testes de Carga e Escala
│   ├── 1.6.3 Testes de Segurança
│   └── 1.6.4 Homologação e Aceite
└── 1.7 Entrega e Documentação
    ├── 1.7.1 Documentação de API
    ├── 1.7.2 Manuais Operacionais
    └── 1.7.3 Liberação para Produção (Go-Live)
```

## 2. Dicionário da EAP

| ID | Pacote de Trabalho | Descrição | Responsável | Critério de Aceitação | Estimativa |
|----|-------------------|-----------|-------------|----------------------|------------|
| 1.1.1 | Provisionamento do Cluster | Criar cluster Kubernetes gerenciado na DigitalOcean com Terraform | DevOps | Cluster operacional, acessível, com nodes provisionados | 24h |
| 1.1.2 | Rede e Segurança (mTLS) | Configurar Istio Service Mesh com mTLS entre serviços | DevOps | Comunicação entre pods criptografada e autenticada | 16h |
| 1.1.3 | API Gateway | Instalar e configurar Kong com rotas de autenticação | DevOps | Rotas `/auth/*` e `/api/*` funcionais com validação de token | 16h |
| 1.1.4 | Pipeline CI/CD + GitOps | Configurar GitHub Actions + Argo CD para deploy automatizado | DevOps | Push na main → build → deploy automático no cluster | 8h |
| 1.2.1 | Provisionamento Multi-Cliente | Criar estrutura de ambientes isolados (Realms) no Keycloak | IAM Specialist | Realm template criado; 1 cliente provisionado e validado | 40h |
| 1.2.2 | Temas e Fluxos OIDC | Configurar fluxo OAuth2/PKCE e tema visual padrão | IAM Specialist | Fluxo OIDC funcional; tema aplicado; cookies HttpOnly | 24h |
| 1.2.3 | Integração Cloudflare | Configurar DNS, WAF e header de identificação de cliente | DevOps + IAM | Header `X-Tenant-Host` injetado e validado | 8h |
| 1.3.1 | Reconhecimento de Domínio | Implementar endpoint que captura domínio e consulta cache | Dev Backend | Domínio mapeado → redireciona; não mapeado → erro padronizado | 24h |
| 1.3.2 | Login e Logout | Implementar fluxos completos de autenticação e encerramento | Dev Backend | Login/logout funcional com PKCE; cookies HttpOnly/Secure/SameSite | 40h |
| 1.3.3 | Sessão e Cookies | Implementar gestão de sessão com renovação silenciosa | Dev Backend | Sessão renovada automaticamente; perfil em <15ms | 24h |
| 1.3.4 | Cache | Implementar cache de mapeamento Host→Realm e JWKS | Dev Backend | Cache responde em <5ms; invalidação sob demanda | 16h |
| 1.4.1 | Modelagem e Isolamento | Criar políticas de isolamento por cliente no banco de dados | DBA | Query com cliente A retorna 0 linhas para dados do cliente B | 32h |
| 1.4.2 | Pool de Conexões | Configurar gerenciador de pool de conexões | DBA | Conexões reutilizadas; sem exaustão sob carga | 8h |
| 1.5.1 | Métricas e Dashboards | Configurar coleta de métricas e dashboards operacionais | DevOps | Métricas de latência, erro e sessões ativas visíveis | 16h |
| 1.5.2 | Logs | Configurar agregação de logs estruturados | DevOps | Logs com correlation_id e tenant_id; zero PII | 8h |
| 1.5.3 | Tracing | Configurar rastreamento distribuído ponta a ponta | DevOps | Spans visíveis: Cloudflare→Kong→BFF→Keycloak | 8h |
| 1.6.1 | Testes de Isolamento | Testar acesso cruzado entre clientes | QA | 100% das tentativas cross-tenant bloqueadas | 24h |
| 1.6.2 | Testes de Carga | Simular picos de acesso e validar escalabilidade | QA | 200+ req/s sem erros 5xx; KEDA escala pods | 20h |
| 1.6.3 | Testes de Segurança | Validar proteção de credenciais e OWASP Top 10 | QA + IAM | Cookies inacessíveis via JS; zero vulnerabilidades críticas | 24h |
| 1.6.4 | Homologação | Validar todos os fluxos e critérios de aceite | QA + PO | Checklist de aceite 100% aprovado | 16h |
| 1.7.1 | Documentação de API | Gerar especificação OpenAPI dos endpoints de autenticação | Dev Backend | OpenAPI 3.0 publicado e acessível | 8h |
| 1.7.2 | Manuais Operacionais | Criar runbooks de operação e recuperação | DevOps + Tech Lead | Runbooks validados em simulação | 8h |
| 1.7.3 | Go-Live | Executar liberação para produção e monitorar | DevOps + PM | Plataforma em produção; termo de aceite assinado | 8h |

## 3. Matriz EAP × Entregas do Charter

| Pacote EAP | Entrega Charter | Status |
|------------|----------------|--------|
| 1.1 (Infraestrutura) | D1 — Ambiente de Produção | ✅ Vinculado |
| 1.2 (Motor de Identidade) | D2 — Motor de Identidade por Cliente | ✅ Vinculado |
| 1.3 (Portal de Acesso) | D3 — Portal de Acesso | ✅ Vinculado |
| 1.4 (Isolamento de Dados) | D4 — Camada de Isolamento | ✅ Vinculado |
| 1.5 (Observabilidade) | D5 — Monitoramento e Alertas | ✅ Vinculado |
| 1.6 (Testes) | D6 — Homologação de Segurança | ✅ Vinculado |
| 1.7 (Entrega) | D7 — Liberação para Uso | ✅ Vinculado |

## 4. Matriz EAP × Requisitos BRD

| Pacote EAP | Requisito BRD Vinculado | Cobertura |
|------------|------------------------|-----------|
| 1.3.1 Reconhecimento de Domínio | REQ-01 — Reconhecimento automático | ✅ |
| 1.4.1 Isolamento | REQ-02 — Isolamento total | ✅ |
| 1.3.2 Login/Logout + 1.3.3 Sessão | REQ-03 — Proteção de credenciais | ✅ |
| 1.3 (todo o Portal) | REQ-04 — Portal padronizado | ✅ |
| 1.3.4 Cache + 1.3.3 Sessão | REQ-05 — Resposta <15ms | ✅ |
| 1.6.2 Carga + 1.1 Infra | REQ-06 — Suporte a picos | ✅ |
| 1.5.2 Logs | REQ-07 — Registro de auditoria | ✅ |
| 1.2 + 1.3.1 | REQ-08 — Ativação <4h | ✅ |
| 1.6.2 Carga + 1.1.4 Pipeline | REQ-09 — Adaptação automática | ✅ |
| 1.3 (Portal) | REQ-10 — Experiência consistente | ✅ |

**Cobertura EAP × BRD: 10/10 (100%)**

---

**[STATUS: SUCESSO]** — EAP com 7 entregas, 23 pacotes de trabalho, 3 níveis de decomposição. Cobertura total das entregas e requisitos.
