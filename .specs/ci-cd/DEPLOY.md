# 🚀 Guia de Deploy e Automação (CI/CD)

Este documento define o fluxo de entrega contínua, infraestrutura e automação do projeto.

🤖 **Instrução para Agentes de IA:** Sempre que você criar uma nova feature, alterar pacotes/dependências ou modificar a infraestrutura, valide se as mudanças respeitam este fluxo e atualize os scripts de CI/CD se necessário.

---

## 🏗️ Fluxo de Branches e Ambientes

O projeto adota o modelo de entrega baseado em gatilhos de Git:

| Ambiente | Branch Gatilho | Tipo de Deploy | Objetivo |
| :--- | :--- | :--- | :--- |
| **Development** | `develop` | Automático (CI/CD) | Testes integrados e homologação interna. |
| **Staging** | `release/*` | Automático (CI/CD) | Validação final pré-produção (UAT). |
| **Production** | `main` | Automático via Tag | Versão estável aberta ao usuário final. |

---

## 🔒 Integração de Segurança no CI/CD (`SECURITY.md`)

O pipeline de CI/CD bloqueia o deploy caso o código não atenda aos critérios do **`docs/SECURITY.md`**. As etapas automáticas abaixo rodam obrigatoriamente a cada Pull Request e antes de qualquer deploy:

1. **SAST (Static Application Security Testing):** Varredura estática de código para capturar falhas antes do build (ex: senhas expostas, SQL injection).
2. **Dependency Audit:** Bloqueio imediato se houver pacotes com vulnerabilidades conhecidas instalados.
3. **Secrets Detection:** Varredura automática no Git para garantir que nenhum arquivo `.env` ou token privado foi commitado por engano.

---

## 📝 Checklist de CI/CD para Desenvolvimento (SDD)

Use este checklist para validar se a funcionalidade atual está pronta para o pipeline de deploy.

### 1. Variáveis de Ambiente e Configuração
- [ ] **Ambiente Isolado:** Toda nova chave de API ou variável criada foi documentada no arquivo `.env.example`?
- [ ] **Configuração na Nuvem:** As novas variáveis já foram cadastradas secretamente no painel do provedor de CI/CD (GitHub Secrets, GitLab Variables)?
- [ ] **Zero Hardcoded:** Foi validado que o código usa exclusivamente o gerenciador de ambiente do sistema e respeita as regras de dados sensíveis do `SECURITY.md`?

### 2. Validação e Qualidade (Passos do Pipeline)
- [ ] **Build Check:** O comando de build (`npm run build` / `docker build`) roda localmente sem avisos (warnings) ou erros?
- [ ] **Suíte de Testes:** Os testes unitários e de integração (definidos no seu `TEST_PLAN`) rodam e passam com 100% de sucesso?
- [ ] **Linter & Formatter:** O código passa na validação estática de estilo e boas práticas?

### 3. Banco de Dados e Migrations
- [ ] **Backward Compatibility:** Se houver alteração no banco de dados (migrations), ela é retrocompatível? (Não quebra o sistema atual que está rodando em produção enquanto o deploy acontece).
- [ ] **Script de Rollback:** Existe um plano ou script automático de reversão caso a migração do banco falhe no meio do deploy?

### 4. Empacotamento e Infraestrutura
- [ ] **Docker / Container:** O `Dockerfile` (se aplicável) foi atualizado caso uma nova dependência de sistema operacional seja necessária?
- [ ] **Health Check:** O endpoint de monitoramento de saúde da aplicação (`/health` ou `/status`) está respondendo corretamente para garantir que o orquestrador saiba se o deploy deu certo?

---

## 🛠️ Comandos Úteis do Pipeline (Execução Local)

Simule os passos que o servidor de CI/CD executará para garantir que o seu Pull Request não quebre o pipeline:

```bash
# 1. Instalação limpa de dependências de produção
npm ci

# 2. Auditoria de Segurança automatizada (Cruza dados com o SECURITY.md)
npm audit --audit-level=high

# 3. Execução dos testes definidos por feature
npm run test:coverage

# 4. Compilação/Build do projeto
npm run build
```
*(Nota para IA: Adapte os comandos acima caso a stack técnica descrita no `SYSTEM_PROMPT.md` mude).*

---

## 🚨 Plano de Rollback (Em caso de falha em Produção)

Se o pipeline completar o deploy mas o monitoramento indicar erros graves (Status 5XX) em produção:
1. O orquestrador reverterá automaticamente o tráfego para o container da versão anterior (*Blue/Green Deployment*).
2. Se o banco já tiver rodado a migration e falhado, execute imediatamente o script de rollback documentado nas tarefas da feature correspondente.
