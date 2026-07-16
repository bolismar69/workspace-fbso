# Requirements — Solar Fácil Site

> Especificação de requisitos funcionais e não-funcionais (EARS + MoSCoW).
> Gerado por `documentation-writer` em 2026-07-08.
> Base: `.specs/REQUIREMENTS.md` original (2026-07-06), análise do código-fonte.

---

## 1. Requisitos Funcionais (RF)

### Módulo: Calculadora (10 RFs)

| ID | Requisito | Prioridade | Status |
|---|---|---|---|
| RF-CAL-001 | O sistema DEVE calcular a economia mensal do consumidor como 12% do valor da conta | Must | ✅ |
| RF-CAL-002 | O sistema DEVE sugerir um plano baseado na faixa de consumo (DT-001) | Must | ✅ |
| RF-CAL-003 | O sistema DEVE formatar valores em Real brasileiro (R$) | Must | ✅ |
| RF-CAL-004 | O sistema DEVE exibir mensagem para valores abaixo de R$ 50 | Must | ✅ |
| RF-CAL-005 | O sistema DEVE exibir mensagem para valores acima de R$ 5.000 | Must | ✅ |
| RF-CAL-006 | O sistema DEVE calcular o ganho do fornecedor como kWh × R$ 0,40 | Must | ✅ |
| RF-CAL-007 | O sistema DEVE exibir mensagem para menos de 50 kWh | Must | ✅ |
| RF-CAL-008 | O sistema DEVE exibir mensagem para mais de 10.000 kWh | Must | ✅ |
| RF-CAL-009 | O sistema DEVE permitir alternar entre calculadora de consumidor e fornecedor | Should | ✅ |
| RF-CAL-010 | O sistema DEVE rastrear uso da calculadora no GA4 | Should | ✅ |

### Módulo: Formulário de Contato (13 RFs)

| ID | Requisito | Prioridade | Status |
|---|---|---|---|
| RF-CON-001 | O sistema DEVE validar nome (obrigatório, mín 2 chars) | Must | ✅ |
| RF-CON-002 | O sistema DEVE validar email (obrigatório, formato válido) | Must | ✅ |
| RF-CON-003 | O sistema DEVE validar telefone (opcional, 10-11 dígitos) | Must | ✅ |
| RF-CON-004 | O sistema DEVE validar perfil (obrigatório) | Must | ✅ |
| RF-CON-005 | O sistema DEVE validar mensagem (opcional, máx 1000 chars) | Must | ✅ |
| RF-CON-006 | O sistema DEVE ter proteção anti-spam (honeypot + timer 3s) | Must | ✅ |
| RF-CON-007 | O sistema DEVE exibir mensagens de erro em português | Must | ✅ |
| RF-CON-008 | O sistema DEVE enviar formulário via POST multipart/form-data | Must | ✅ |
| RF-CON-009 | O sistema DEVE exibir tela de sucesso após envio | Must | ✅ |
| RF-CON-010 | O sistema DEVE sugerir WhatsApp como alternativa em caso de falha | Should | ✅ |
| RF-CON-011 | O sistema DEVE preservar contexto da calculadora no formulário | Should | ✅ |
| RF-CON-012 | O sistema DEVE ter validação server-side | Should | ❌ |
| RF-CON-013 | O endpoint do formulário DEVE ser configurável via variável de ambiente | Must | ✅ |

### Módulo: Planos (8 RFs)

| ID | Requisito | Prioridade | Status |
|---|---|---|---|
| RF-PLN-001 | O sistema DEVE exibir 3 planos (Basic, Special, Premium) | Must | ✅ |
| RF-PLN-002 | O sistema DEVE destacar o plano Special como "Mais Popular" | Should | ✅ |
| RF-PLN-003 | O sistema DEVE exibir preço, capacidade e features de cada plano | Must | ✅ |
| RF-PLN-004 | O sistema DEVE ter tabela comparativa na página /planos | Should | ✅ |
| RF-PLN-005 | O sistema DEVE ter seção de planos na homepage | Should | ✅ |
| RF-PLN-006 | O sistema DEVE ter FAQ com 6 perguntas | Should | ✅ |
| RF-PLN-007 | O sistema DEVE ter accordion para FAQ (expandir/fechar) | Should | ✅ |
| RF-PLN-008 | O sistema DEVE carregar planos com loading skeleton | Should | ✅ |

### Módulo: Navegação (9 RFs)

| ID | Requisito | Prioridade | Status |
|---|---|---|---|
| RF-NAV-001 | O sistema DEVE ter 3 rotas: /, /planos, /contato | Must | ✅ |
| RF-NAV-002 | O sistema DEVE ter Header fixo com logo e links | Must | ✅ |
| RF-NAV-003 | O sistema DEVE ter menu mobile (hamburger) | Must | ✅ |
| RF-NAV-004 | O sistema DEVE ter Footer com links e redes sociais | Must | ✅ |
| RF-NAV-005 | O sistema DEVE ter Breadcrumb nas páginas internas | Could | ✅ |

### Módulo: SEO (4 RFs)

| ID | Requisito | Prioridade | Status |
|---|---|---|---|
| RF-SEO-001 | O sistema DEVE gerar robots.txt dinâmico | Must | ✅ |
| RF-SEO-002 | O sistema DEVE gerar sitemap.xml dinâmico | Must | ✅ |
| RF-SEO-003 | O sistema DEVE ter metadata (title, description, OG) em todas as páginas | Must | ✅ |
| RF-SEO-004 | O sistema DEVE ter JSON-LD structured data | Should | ✅ |

### Módulo: Analytics (5 RFs)

| ID | Requisito | Prioridade | Status |
|---|---|---|---|
| RF-ACT-001 | O sistema DEVE carregar Google Analytics 4 condicionalmente | Should | ✅ |
| RF-ACT-002 | O sistema DEVE rastrear cliques em CTA | Should | ✅ |
| RF-ACT-003 | O sistema DEVE rastrear uso da calculadora | Should | ✅ |
| RF-ACT-004 | O sistema DEVE rastrear abertura de FAQ | Could | ❌ |
| RF-ACT-005 | O sistema DEVE rastrear captura de lead | Should | ❌ |

---

## 2. Requisitos Não-Funcionais (RNF)

| ID | Requisito | Categoria | Prioridade | Status |
|---|---|---|---|---|
| RNF-PERF-001 | LCP < 2.5s em 4G | Performance | Must | ✅ (SSR) |
| RNF-PERF-002 | Bundle JS < 200KB (gzip) | Performance | Should | ✅ (leve) |
| RNF-SEC-001 | Headers de segurança (X-Frame-Options, X-Content-Type-Options) | Segurança | Must | ✅ |
| RNF-SEC-002 | Content-Security-Policy header | Segurança | Should | ❌ |
| RNF-SEC-003 | HSTS header | Segurança | Should | ❌ |
| RNF-ACC-001 | WCAG 2.1 AA | Acessibilidade | Should | 🟡 |
| RNF-ACC-002 | Contraste ≥ 4.5:1 | Acessibilidade | Must | ✅ |
| RNF-RESP-001 | Funcionar em 375px–1920px | Responsividade | Must | ✅ |
| RNF-DEP-001 | Output standalone | Deploy | Must | ✅ |
| RNF-I18N-001 | Suporte a pt-BR (primário) | i18n | Must | ✅ |
| RNF-I18N-002 | Arquitetura preparada para /es/ | i18n | Could | ❌ |

---

## 3. MoSCoW Summary

| Prioridade | Count |
|---|---|
| **Must** (obrigatório) | 18 |
| **Should** (importante) | 14 |
| **Could** (desejável) | 3 |
| **Won't** (não agora) | 0 |
| **Total** | **35** |

---

Última atualização: 2026-07-08
