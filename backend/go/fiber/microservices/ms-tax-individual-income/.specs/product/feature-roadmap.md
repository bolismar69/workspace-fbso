# Feature Roadmap — ms-tax-individual-income

## 🗺️ Evolução Programada do Motor de Cálculo

- [ ] **Módulo de Resiliência INSS (Circuit Breaker):** Tratar falhas no cliente HTTP externo de maneira segura com políticas de retry.
- [x] **Integração Nativa de Cache:** Implementar o barramento do Redis mapeado na arquitetura para regras fiscais voláteis. _(2026-06-20)_
- [ ] **Suporte Completo Multilíngue para Erros:** Suportar mensagens parametrizáveis de internacionalização.
- [x] **Health Check Endpoint:** Expor `/health`, `/healthz` e `/api/v1/health` para probes de orquestrador (K8s, Docker). _(2026-06-20)_
- [ ] **Autenticação/ Autorização:** Adicionar middleware de segurança ao endpoint (API key, JWT ou mTLS).
- [ ] **Métricas Prometheus:** Exportar métricas de latência, taxa de erro, e falhas de integração.
- [ ] **Suíte de Testes Automatizados:** Cobertura de testes unitários e de integração para cenários fiscais.

## 🖥️ Natureza do Serviço

Este microserviço (`ms-tax-individual-income`) opera estritamente na camada de backend, expondo contratos RESTful sem componentes de interface visual direta (Stateless Engine).
