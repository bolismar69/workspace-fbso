---
title: "Requisitos — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
diataxis: "Reference"
---

# Especificação de Requisitos — Solar Fácil

## 1. Requisitos Funcionais (RF)

### RF01 — Cadastro de Associado
**EARS:** WHEN the user preenche o formulário de cadastro AND submete dados válidos THEN o sistema DEVE salvar o associado no SQLite local.

**Critérios de Aceitação:**
- [ ] Formulário com campos obrigatórios: nome, email, CPF/CNPJ, senha, tipoPessoa, tipoAssociado
- [ ] Validação de CPF/CNPJ (algoritmo específico)
- [ ] Campos condicionais: PF (dataNascimento, nomeSocial), PJ (dataAbertura, razaoSocial, nomeFantasia)
- [ ] Campos condicionais: Beneficiado (concessionaria, consumoMedio, planoDesejado), Fornecedor (potenciaInstalada, disponibilidade, tipoConexao)
- [ ] Aceite de termos obrigatório
- [ ] CPF/CNPJ único (constraint SQLite)

**MoSCoW:** Must Have

### RF02 — Login
**EARS:** WHEN the user informa CPF/CNPJ e senha válidos THEN o sistema DEVE autenticar e redirecionar para a home.

**Critérios de Aceitação:**
- [ ] Busca por CPF/CNPJ + senha no SQLite
- [ ] Estado de autenticação mantido em AuthContext
- [ ] Tab de login muda para logout quando autenticado
- [ ] Mensagem de erro para credenciais inválidas

**MoSCoW:** Must Have

### RF03 — Logout
**EARS:** WHEN the user toca na tab de logout THEN o sistema DEVE limpar o estado de autenticação.

**MoSCoW:** Must Have

### RF04 — Listagem de Associados
**EARS:** WHEN the user acessa a tela de lista THEN o sistema DEVE exibir todos os associados cadastrados.

**MoSCoW:** Must Have

### RF05 — Visualização de Movimentações
**EARS:** WHEN the user acessa a tela de movimentações THEN o sistema DEVE exibir as movimentações mensais do associado logado.

**Critérios de Aceitação:**
- [ ] Cards com: valor total, energia recebida (kWh), tarifa unitária, valor cobrado, valor economizado, percentual economizado
- [ ] Status de pagamento com cor de fundo (Pago=verde, Pendente=amarelo)
- [ ] Gráfico de economia (Victory)

**MoSCoW:** Must Have

### RF06 — Catálogo de Planos
**EARS:** WHEN the user acessa a tela de planos THEN o sistema DEVE exibir os planos comerciais disponíveis.

**MoSCoW:** Should Have

### RF07 — FAQ
**EARS:** WHEN the user acessa a tela de FAQ THEN o sistema DEVE exibir perguntas frequentes com accordion expansível.

**MoSCoW:** Should Have

### RF08 — Conteúdo Institucional (Saiba Mais)
**EARS:** WHEN the user acessa a tela Saiba Mais THEN o sistema DEVE exibir conteúdo institucional sobre energia solar.

**MoSCoW:** Could Have

### RF09 — CRUD de Movimentações
**EARS:** WHEN the user precisa gerenciar movimentações THEN o sistema DEVE permitir criar, editar e excluir registros.

**MoSCoW:** Should Have

### RF10 — Persistência de Sessão
**EARS:** WHEN the user fecha e reabre o app THEN o sistema DEVE restaurar a sessão se o usuário estava logado.

**MoSCoW:** Should Have — [TODO] Não implementado

## 2. Requisitos Não-Funcionais (RNF)

### RNF01 — Offline-First
**Descrição:** O app DEVE funcionar completamente sem conexão de internet.
**Validação:** Todas as funcionalidades disponíveis em modo avião.
**MoSCoW:** Must Have

### RNF02 — Performance de Listas
**Descrição:** Listas com mais de 50 itens DEVEM usar virtualização (FlatList/FlashList).
**Validação:** Scroll com 100+ itens mantém 60fps.
**MoSCoW:** Should Have — [TODO] Ainda usa ScrollView

### RNF03 — Acessibilidade Básica
**Descrição:** Componentes interativos DEVEM ter accessibilityLabel e accessibilityRole.
**Validação:** VoiceOver/TalkBack navegam corretamente por todas as telas.
**MoSCoW:** Should Have — [TODO] Não implementado

### RNF04 — Segurança de Dados
**Descrição:** Senhas NÃO DEVEM ser armazenadas em plain text.
**Validação:** Verificar hash no SQLite.
**MoSCoW:** Must Have — [TODO] Implementar hash

### RNF05 — iOS e Android
**Descrição:** O app DEVE funcionar em iOS 15+ e Android 7.0+.
**Validação:** Testar em iPhone SE e Pixel 6a.
**MoSCoW:** Must Have

### RNF06 — Tamanho de Touch Targets
**Descrição:** Elementos interativos DEVEM ter no mínimo 44×44pt.
**Validação:** Accessibility Inspector (iOS) / Scanner (Android).
**MoSCoW:** Should Have

### RNF07 — Suporte a Temas
**Descrição:** O app DEVE oferecer tema claro e escuro, seguindo a preferência do sistema.
**Validação:** Alternar tema no sistema → app reflete mudança.
**MoSCoW:** Must Have

## 3. Matriz MoSCoW

| Prioridade | RFs |
|---|---|
| **Must Have** | RF01 (Cadastro), RF02 (Login), RF03 (Logout), RF04 (Lista), RF05 (Movimentações), RNF01 (Offline), RNF04 (Segurança), RNF05 (Cross-platform), RNF07 (Temas) |
| **Should Have** | RF06 (Planos), RF07 (FAQ), RF09 (CRUD Mov), RF10 (Persistência), RNF02 (Performance), RNF03 (A11y), RNF06 (Touch) |
| **Could Have** | RF08 (Saiba Mais) |
