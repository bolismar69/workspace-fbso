# Manuais de Usuário: Plataforma Shield
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 01-PROJECT-CHARTER, 03-SRS |
| **Solução Técnica** | ms-shield-identity-auth |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. Getting Started

### Para Times de Produto (Consumidores da Plataforma)

A Plataforma Shield é a camada de identidade centralizada da FBSO.ORG. Para integrar seu produto:

1. **Registre seu produto** — solicite ao Product Owner o cadastro da sua aplicação como cliente OIDC
2. **Configure redirect URIs** — informe as URLs de callback do seu frontend
3. **Implemente o fluxo de login** — redirecione usuários não autenticados para `/auth/login?redirect_uri=SEU_CALLBACK`
4. **Consuma o perfil** — após autenticação, chame `/auth/me` para obter dados do usuário
5. **Implemente o logout** — redirecione para `/auth/logout`

### Para Administradores de Cliente (Escola/Universidade)

1. Acesse `https://[seu-dominio].fbso.org`
2. A plataforma reconhece automaticamente sua instituição
3. Faça login com as credenciais fornecidas
4. Acesse os produtos contratados

## 2. Feature Walkthrough

| Feature (SRS) | Como usar |
|--------------|----------|
| F-01 — Reconhecimento | Acesse o domínio da sua escola. Nenhuma ação manual necessária |
| F-02 — Login | Ao acessar qualquer produto, você será redirecionado para a tela de login da sua instituição |
| F-03 — Perfil | Seus dados (nome, e-mail, papel) estão disponíveis em todos os produtos |
| F-04 — Segurança | Seus dados são isolados — ninguém de outra escola pode acessá-los |
| F-07 — Logout | Ao sair, você será desconectado de todos os produtos simultaneamente |

## 3. Step-by-Step Guides

### Integração de Novo Produto

| Passo | Ação | Responsável |
|-------|------|------------|
| 1 | Solicitar client_id e client_secret ao time Shield | Time de Produto |
| 2 | Configurar redirect_uri no Keycloak | IAM Specialist |
| 3 | Implementar redirecionamento `/auth/login?redirect_uri=...` no frontend | Dev Frontend |
| 4 | Implementar consumo de `/auth/me` para perfil | Dev Frontend |
| 5 | Implementar chamada a `/auth/logout` no botão "Sair" | Dev Frontend |
| 6 | Testar fluxo completo em staging | QA + Time de Produto |
| 7 | Solicitar liberação para produção | Time de Produto → PO |

### Primeiro Acesso de Usuário Final

| Passo | Ação |
|-------|------|
| 1 | Acesse o endereço da sua escola no navegador |
| 2 | Você será redirecionado para a tela de login |
| 3 | Informe o usuário e senha fornecidos pela sua instituição |
| 4 | Após o primeiro login, altere sua senha |
| 5 | Pronto — você tem acesso a todos os produtos contratados |

## 4. FAQ

| Pergunta | Resposta |
|----------|---------|
| Preciso criar uma conta para cada produto? | Não. Uma única conta dá acesso a todos os produtos contratados pela sua escola |
| Esqueci minha senha. O que fazer? | Clique em "Esqueci minha senha" na tela de login. Um link de recuperação será enviado para seu e-mail |
| Posso acessar de casa? | Sim. O acesso é via internet, de qualquer lugar |
| Meus dados estão seguros? | Sim. Cada escola tem seu ambiente isolado. Ninguém de fora da sua instituição acessa seus dados |
| O que acontece se eu ficar muito tempo sem usar? | Após 30 minutos de inatividade, você precisará fazer login novamente |

## 5. Troubleshooting

| Problema | Causa Provável | Solução |
|----------|--------------|---------|
| "Domínio não reconhecido" | Sua escola não está cadastrada na plataforma | Entre em contato com o suporte FBSO |
| "Sessão expirada" | Você ficou mais de 30min sem usar o sistema | Faça login novamente |
| Tela branca após login | Erro de conexão com o produto | Atualize a página (F5). Se persistir, contate o suporte |
| "Acesso negado" | Seu perfil não tem permissão para acessar este recurso | Solicite acesso ao administrador da sua escola |

---

**[STATUS: SUCESSO]** — Manual com 5 seções: getting started, walkthrough, step-by-step, FAQ (5 itens), troubleshooting (4 cenários).
