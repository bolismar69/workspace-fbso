# 🛡️ Política e Checklist de Segurança do Projeto

Este documento define os padrões obrigatórios de segurança para este repositório.

🤖 **Instrução para Agentes de IA:** Você deve ler este arquivo antes de implementar qualquer funcionalidade, rota de API, alteração em banco de dados ou refatoração. Nenhuma entrega será aceita se violar as regras abaixo.

---

## 🛑 Regras de Ouro (Inegociáveis)

1. **Princípio do Menor Privilégio:** Por padrão, todo novo recurso, rota ou endpoint é **privado e bloqueado**, a menos que a `SPEC.md` da funcionalidade explicitamente exija acesso público.
2. **Zero Hardcoded Secrets:** Nunca insira chaves de API, senhas, tokens ou strings de conexão diretamente no código. Use variáveis de ambiente (`process.env`, `.env`).
3. **Não confie no Input do Usuário:** Todo dado vindo do cliente (query params, body, headers, cookies) é hostil e deve ser sanitizado e validado antes do processamento.

---

## 📝 Checklist de Segurança para Desenvolvimento (SDD)

Use este checklist para validar o código gerado antes de marcar as tarefas como concluídas.

### 1. Autenticação e Autorização
- [ ] **Middleware de Sessão:** A rota possui validação de token/sessão ativa?
- [ ] **RBAC (Role-Based Access Control):** Foi validado se o usuário atual tem o nível de permissão correto (ex: `admin`, `user`, `manager`) para executar a ação?
- [ ] **IDOR (Insecure Direct Object References):** Ao buscar um recurso (ex: `/api/orders/123`), o sistema valida se o ID `123` pertence de fato ao usuário autenticado?

### 2. Proteção de Dados e Privacidade
- [ ] **Criptografia em Repouso:** Senhas e dados altamente sensíveis são hasheados usando algoritmos seguros (ex: `bcrypt`, `argon2`) antes de irem para o banco?
- [ ] **Vazamento em Logs:** O sistema foi configurado para **nunca** registrar dados sensíveis (senhas, cartões, tokens, CPFs) nos logs de erro ou de auditoria?
- [ ] **Campos Mascarados:** Dados sensíveis retornados na API pública estão mascarados ou omitidos (ex: retornar `****.****.1234` para cartões)?

### 3. Validação e Sanitização de Entradas
- [ ] **Schema Validation:** O payload da requisição foi validado usando uma biblioteca de tipagem/schema (ex: `Zod`, `Joi`, `Yup`, `Pydantic`)?
- [ ] **SQL/NoSQL Injection:** Todas as queries de banco de dados usam *Parameterized Queries* (ORM/Query Builders)? Nenhuma query foi construída usando concatenação manual de strings.
- [ ] **XSS (Cross-Site Scripting):** Inputs de texto que aceitam HTML ou Markdown passam por um processo de sanitização de strings?

### 4. Proteção de Infraestrutura e API
- [ ] **Rate Limiting:** A nova rota exposta possui alguma proteção contra ataques de força bruta ou Denial of Service (DoS)?
- [ ] **CORS (Cross-Origin Resource Sharing):** As permissões de CORS estão restritas apenas aos domínios autorizados (evitar `Origin: *` em produção)?
- [ ] **Tratamento de Erros Seguro:** Mensagens de erro de exceções internas (ex: stack traces de banco de dados) foram ocultadas do usuário final e substituídas por mensagens genéricas?

---

## 🛠️ Ferramentas de Verificação Automatizada

Antes de abrir um Pull Request, garanta que os seguintes comandos rodem sem reportar vulnerabilidades:

* **Análise de Dependências:** `npm audit` / `pip audit` / `cargo audit` (conforme a stack).
* **Análise Estática (SAST):** Garanta que o Linter do projeto esteja configurado com regras de segurança ativas.

---

## 🚨 Como Reportar uma Vulnerabilidade (Para Humanos)

- Se você encontrar uma falha de segurança real neste sistema, **não abra uma Issue pública**. 
- Por favor, envie um e-mail detalhado com o passo a passo para reproduzir o problema para: **org-fbso@gmail.com**. Responderemos em até 48 horas com o plano de mitigação.
