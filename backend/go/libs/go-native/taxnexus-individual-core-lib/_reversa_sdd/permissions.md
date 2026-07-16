# Permissões & Controle de Acesso — taxnexus-individual-core-lib

> Gerado pelo **Detetive** (Reversa) em 2026-06-10 · `doc_level = completo`
> Confiança: 🟢 CONFIRMADO · 🟡 INFERIDO · 🔴 LACUNA

## Veredito: não há RBAC/ACL nesta biblioteca 🟢

`taxnexus-individual-core-lib` é uma **biblioteca de acesso a dados** (não um serviço com endpoints nem fronteira de autenticação). Não existe:
- Conceito de usuário, papel ou sessão no código. 🟢
- Verificação de permissão em nenhuma função do `repository`. 🟢
- Middleware de autenticação/autorização (não há camada HTTP). 🟢

O controle de acesso real acontece em **duas camadas externas a este recorte**:

| Camada de controle | Mecanismo | Onde | Confiança |
|--------------------|-----------|------|-----------|
| **Aplicação consumidora** | Autenticação/autorização do serviço de cálculo que importa esta lib (camada ausente, L1) | fora do repositório | 🔴 LACUNA |
| **Banco de dados (PostgreSQL)** | Credenciais do `connString` (`DATABASE_URL`) — usuário do banco com privilégios sobre o schema `individual_tax_rates` | infraestrutura | 🟡 |
| **Redis** | Endereço sem auth (`REDIS_ADDR`) — **sem senha/TLS** no código | `cache/redis.go` | 🟢 |

---

## Matriz de acesso a dados (nível de banco) 🟡

A única "permissão" observável é o que o usuário do banco precisa para a biblioteca funcionar. Inferido pelas operações executadas:

| Recurso (schema `individual_tax_rates`) | Operação exigida pela lib | Função | Confiança |
|------------------------------------------|---------------------------|--------|-----------|
| `tax_definitions` | `SELECT` | `GetTaxRulesForPeriod` (JOIN) | 🟢 |
| `tax_rules_history` | `SELECT` | `GetTaxRulesForPeriod` | 🟢 |
| `tax_configs` | `SELECT` | `GetConfig`, `GetTableConfigs` | 🟢 |
| `tax_calculation_log` | — (nenhuma operação neste recorte) | — | 🔴 |

> 🟢 Esta biblioteca é **somente leitura**: não há `INSERT`/`UPDATE`/`DELETE` em nenhuma função. O usuário do banco poderia, em tese, ter apenas privilégio `SELECT` sobre as três tabelas lidas. Escrita (manutenção das faixas, gravação de log) pertence a outras camadas. 🟡

---

## Sinais de segurança relevantes (para o Architect / Reviewer)

| # | Observação | Risco | Confiança |
|---|------------|-------|-----------|
| P1 | Conexão Redis **sem senha e sem TLS** (`redis.Options{Addr: addr}`) | Aceitável em rede interna isolada; risco se Redis for exposto | 🟢 |
| P2 | `DATABASE_URL` de exemplo usa `sslmode=disable` (README) | Sem TLS no banco — aceitável só em ambiente local | 🟢 (exemplo) 🟡 (prod) |
| P3 | Credenciais via variável de ambiente (`DATABASE_URL`, `REDIS_ADDR`) | Padrão 12-factor; gestão de segredo é responsabilidade da infra | 🟢 |
| P4 | Todas as queries usam **parâmetros posicionais** (`$1..$3`) — sem concatenação de string | Imune a SQL injection | 🟢 |

---

## Lacuna 🔴
- **D8 🔴** — Modelo de autorização da aplicação consumidora (quem pode solicitar cálculos, multi-tenant?) está fora deste repositório. Validar com o time se há requisito de isolamento por cliente/tenant que deva refletir no acesso aos dados.

> Como não há RBAC central ao sistema, este artefato documenta a **postura de segurança de acesso a dados** em vez de uma matriz papel×permissão. Em `doc_level = essencial` ele teria sido omitido.
