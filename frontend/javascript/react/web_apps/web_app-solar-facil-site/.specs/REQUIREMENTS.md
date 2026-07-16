# Requisitos Observados: Solar Fácil Site

> **Especificação de engenharia reversa** — todos os requisitos fundamentados no comportamento real do código.
> Formato: **EARS** (Easy Approach to Requirements Syntax).
> Gerado por `/spec-miner` em 2026-07-05. Revisado para pt-BR em 2026-07-06.
> Conteúdo de produto integrado de `PRODUCT.md`.

---

## 1. Módulo Calculadora

### OBS-CALC-001: Cálculo de Economia do Consumidor
Enquanto o usuário fornece uma conta de luz mensal válida (R$ 50–5000), quando o usuário clica em "Calcular", o sistema deve calcular `economia = Math.round(conta × 0.12 × 100) / 100` e sugerir um plano compatível pela faixa de capacidade.
> **Código**: `src/lib/calculator.ts:47-77`

### OBS-CALC-002: Cálculo de Ganho do Fornecedor
Enquanto o usuário fornece um excedente mensal válido (50–10000 kWh), quando o usuário clica em "Simular", o sistema deve calcular `ganho = Math.round(excedente × 0.40 × 100) / 100` e exibir a receita mensal estimada.
> **Código**: `src/lib/calculator.ts:84-113`

### OBS-CALC-003: Tratamento de Outlier — Consumidor
Quando o valor informado pelo consumidor for abaixo de R$ 50 ou acima de R$ 5.000, o sistema deve retornar `isOutlier: true` com uma mensagem contextual e sem sugestão de plano.
> **Código**: `src/lib/calculator.ts:48-66`

### OBS-CALC-004: Tratamento de Outlier — Fornecedor
Quando o valor informado pelo fornecedor for abaixo de 50 kWh ou acima de 10.000 kWh, o sistema deve retornar `isOutlier: true` com uma mensagem contextual direcionando para contato comercial.
> **Código**: `src/lib/calculator.ts:85-103`

### OBS-CALC-005: Tabela de Decisão de Planos
Quando `suggestPlan()` é chamada, o sistema deve associar a conta mensal a um plano usando a faixa de capacidade conforme definido na DT-001:
- 100–200 kWh → Basic
- 200–350 kWh → Special
- 350–600 kWh → Premium
> **Código**: `src/lib/calculator.ts:33-40`

### OBS-CALC-006: Validação de Entrada
Quando o campo da calculadora estiver vazio, não-numérico ou ≤ 0, o sistema deve exibir o erro "Informe um valor para calcular" e impedir o cálculo.
> **Código**: `src/hooks/useCalculator.ts:32-38`

### OBS-CALC-007: Suporte a Teclado
Quando a tecla Enter for pressionada no campo da calculadora, o sistema deve disparar o cálculo.
> **Código**: `src/components/home/ConsumerCalculator.tsx:44`

### OBS-CALC-008: Persistência do Resultado
Enquanto um resultado estiver exibido (`hasCalculated: true`), o sistema deve mostrar um botão "Limpar" (reset) que limpa o valor, resultado e estado de erro.
> **Código**: `src/hooks/useCalculator.ts:63-65`

### OBS-CALC-009: Analytics da Calculadora
Quando um cálculo produzir um resultado válido, o sistema deve disparar um evento GA4 `calculator_use` com `persona`, `input_value`, `result` e opcionalmente `plan_suggested`.
> **Código**: `src/hooks/useCalculator.ts:46-59`

### OBS-CALC-010: Formatação Monetária
O sistema deve formatar todos os valores monetários usando `Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' })`.
> **Código**: `src/lib/calculator.ts:20-23`

---

## 2. Módulo Formulário de Contato

### OBS-CON-001: Validação de Nome
Quando o campo nome for enviado, o sistema deve exigir no mínimo 2 caracteres não-espaço.
> **Código**: `src/lib/validation.ts:6-13`

### OBS-CON-002: Validação de E-mail
Quando o campo e-mail for enviado, o sistema deve validar contra o padrão `/^[^\s@]+@[^\s@]+\.[^\s@]+$/` e exigir valor não-vazio.
> **Código**: `src/lib/validation.ts:16-24`

### OBS-CON-003: Validação de Telefone
Onde telefone for informado, o sistema deve validar exatamente 10 ou 11 dígitos (DDD + número). O sistema deve tratar telefone como opcional — campo vazio passa na validação.
> **Código**: `src/lib/validation.ts:27-35`

### OBS-CON-004: Perfil Obrigatório
Quando o formulário for enviado, o sistema deve exigir uma seleção de perfil não-vazia (`consumidor`, `fornecedor` ou `cooperativa`).
> **Código**: `src/lib/validation.ts:38-43`

### OBS-CON-005: Limite de Mensagem
Onde uma mensagem for fornecida, o sistema deve impor o máximo de 1000 caracteres.
> **Código**: `src/lib/validation.ts:45-49`

### OBS-CON-006: Portão de Tempo Anti-Spam
Quando o formulário for enviado dentro de 3 segundos do carregamento da página (`Date.now() - pageLoadTime < 3000`), o sistema deve retornar silenciosamente sucesso sem enviar os dados, para bloquear bots.
> **Código**: `src/hooks/useContactForm.ts:48-53`

### OBS-CON-007: Campo Honeypot
O sistema deve incluir um campo oculto `website` (tabIndex -1, posicionado fora da tela) que deve permanecer vazio para envio válido.
> **Código**: `src/components/contact/ContactForm.tsx:47-53`, `src/hooks/useContactForm.ts:66`

### OBS-CON-008: Envio do Formulário
Quando o formulário passar na validação e verificações anti-spam, o sistema deve enviar POST como `FormData` para `NEXT_PUBLIC_FORM_ENDPOINT` e exibir tela de sucesso ao receber HTTP 2xx.
> **Código**: `src/hooks/useContactForm.ts:58-84`

### OBS-CON-009: Fallback de Erro de Envio
Quando a requisição POST falhar, o sistema deve exibir uma mensagem de erro com link direto para WhatsApp como alternativa.
> **Código**: `src/hooks/useContactForm.ts:78-82`

### OBS-CON-010: Tela de Sucesso
Quando o envio for bem-sucedido, o sistema deve exibir uma tela de sucesso com ícone de checkmark, mensagem de agradecimento e CTA para baixar o aplicativo.
> **Código**: `src/components/contact/SuccessScreen.tsx`

### OBS-CON-011: Limpeza de Erro ao Alterar Campo
Quando qualquer valor de campo do formulário for alterado, o sistema deve limpar a mensagem de erro daquele campo e o erro de envio.
> **Código**: `src/hooks/useContactForm.ts:30-33`

### OBS-CON-012: Contexto da Jornada via Calculadora
Quando os parâmetros de URL `plano`, `economia`, `excedente` ou `ganho` estiverem presentes, o sistema deve pré-preencher o perfil e exibir um banner JourneySummary mostrando os resultados da calculadora.
> **Código**: `src/components/contact/ContactForm.tsx:14-24`, `src/components/contact/JourneySummary.tsx`

### OBS-CON-013: Grupo de Rádio de Perfil
O sistema deve exibir a seleção de perfil como um grupo de 3 chips de rádio: Consumidor, Fornecedor, Cooperativa — cada um com ícone e texto distintos.
> **Código**: `src/components/contact/ContactForm.tsx:87-118`

---

## 3. Módulo Planos

### OBS-PLN-001: Dados de Plano Estáticos
O sistema deve servir dados de planos a partir da constante em código `PLANS` (não do JSON mock) com um delay assíncrono simulado de 500ms.
> **Código**: `src/services/servicePlans.ts:6-17`

### OBS-PLN-002: Três Níveis de Plano
O sistema deve oferecer exatamente três planos: Basic (R$150/mês), Special (R$250/mês), Premium (R$400/mês), com Special marcado como destaque.
> **Código**: `src/lib/constants.ts:6-47`

### OBS-PLN-003: Exibição de Card de Plano
O sistema deve exibir cards de plano com nome, preço, capacidade, recursos e um botão CTA que leva a `/contato?perfil=consumidor&plano={nomePlano}`.
> **Código**: `src/components/plans/PlanCard.tsx`

### OBS-PLN-004: Plano em Destaque
O plano com `highlight: true` deve exibir um selo "Mais Popular", acento de anel verde e ênfase visual scale(1.02).
> **Código**: `src/components/plans/PlanCard.tsx:17-29`

### OBS-PLN-005: Tabela Comparativa (Desktop)
O sistema deve exibir planos como uma tabela HTML responsiva com linhas de recursos em viewport ≥1024px.
> **Código**: `src/components/plans/PlansComparisonTable.tsx:62-116`

### OBS-PLN-006: Pilha de Cards (Mobile)
O sistema deve exibir planos como uma pilha vertical de cards abaixo de 1024px de viewport.
> **Código**: `src/components/plans/PlansComparisonTable.tsx:17-59`

### OBS-PLN-007: CTA de Fornecedor
Onde a página for `/planos`, o sistema deve exibir um card ProviderHighlight oferecendo R$ 0,40/kWh para produtores solares.
> **Código**: `src/components/plans/ProviderHighlight.tsx`, `src/app/planos/page.tsx:30-32`

### OBS-PLN-008: Variantes de Card de Plano
O sistema deve suportar duas variantes de PlanCard:
- `compact`: exibe os primeiros 3 recursos (homepage)
- `full`: exibe todos os recursos (não utilizado atualmente)
> **Código**: `src/components/plans/PlanCard.tsx:11`

---

## 4. Navegação e Layout

### OBS-NAV-001: Cabeçalho Fixo
O sistema deve renderizar um cabeçalho fixo de 64px com fundo branco/80%, backdrop-blur e borda inferior.
> **Código**: `src/components/layout/Header.tsx:22-23`

### OBS-NAV-002: Indicação de Link Ativo
Quando o caminho atual corresponder ao href de um link de navegação, o sistema deve estilizá-lo com `text-solar-primary` e `font-semibold`.
> **Código**: `src/components/layout/Header.tsx:14-19`

### OBS-NAV-003: Menu Mobile
Quando o botão hambúrguer for clicado, o sistema deve exibir um overlay full-screen com painel deslizante da direita contendo links de navegação + CTA. O sistema deve fechar ao pressionar Escape, clicar no botão fechar ou clicar no backdrop.
> **Código**: `src/components/layout/MobileMenu.tsx`

### OBS-NAV-004: Bloqueio de Scroll do Body
Enquanto o menu mobile estiver aberto, o sistema deve definir `document.body.style.overflow = 'hidden'` e restaurar ao fechar/desmontar.
> **Código**: `src/components/layout/MobileMenu.tsx:22-29`

### OBS-NAV-005: Navegação Breadcrumb
O sistema deve exibir breadcrumbs nas subpáginas (`/planos`, `/contato`) com `aria-label="Breadcrumb"` estruturado e separadores de chevron.
> **Código**: `src/components/shared/Breadcrumb.tsx`

### OBS-NAV-006: Estrutura do Rodapé
O sistema deve renderizar um rodapé de 4 colunas com Marca, Navegação, Links das Lojas de App e Contato/Redes Sociais — mais uma barra inferior com selo ANEEL e links legais.
> **Código**: `src/components/layout/Footer.tsx`

### OBS-NAV-007: Link Pular para Conteúdo
O sistema deve incluir um link de salto visualmente oculto que se torna visível ao receber foco, posicionado no topo-esquerda com z-100.
> **Código**: `src/app/layout.tsx:45-50`

### OBS-NAV-008: Margem de Scroll
O sistema deve aplicar `scroll-margin-top: 80px` a todos os elementos `section[id]` para compensar o cabeçalho fixo.
> **Código**: `src/app/globals.css:57-59`

### OBS-NAV-009: Scroll Suave
O sistema deve usar `scroll-behavior: smooth` no nível HTML e `scrollIntoView({ behavior: 'smooth' })` programático para navegação por clique em CTA.
> **Código**: `src/app/globals.css:53-54`, `src/components/home/HeroSection.tsx:9`

---

## 5. SEO e Dados Estruturados

### OBS-SEO-001: Metadados
O sistema deve exportar `Metadata` no nível da página de cada rota com title, description, tags OpenGraph (title, description, url, siteName, locale, type) e diretivas de robots.
> **Código**: `src/app/layout.tsx:14-35`, `src/app/planos/page.tsx:8-12`, `src/app/contato/page.tsx:10-13`

### OBS-SEO-002: Dados Estruturados JSON-LD
O sistema deve injetar um script JSON-LD Schema.org `Organization` com nome, url, descrição, logo, sameAs (Instagram, LinkedIn) e contactPoint.
> **Código**: `src/components/shared/JsonLd.tsx`

### OBS-SEO-003: robots.txt
O sistema deve gerar `robots.txt` permitindo todos os agentes de usuário, apontando para o sitemap em `{SITE_URL}/sitemap.xml`.
> **Código**: `src/app/robots.ts`

### OBS-SEO-004: Sitemap
O sistema deve gerar um sitemap XML com 3 URLs: `/` (prioridade 1.0), `/planos` (0.8), `/contato` (0.8), todos com `changeFrequency` mensal.
> **Código**: `src/app/sitemap.ts`

---

## 6. Analytics

### OBS-ANA-001: Carregamento Condicional do GA4
Onde `NEXT_PUBLIC_GA_ID` estiver configurado e não for o placeholder `'G-XXXXXXXXXX'`, o sistema deve injetar o script Google Analytics 4 com carregamento assíncrono e rastreamento de page_view.
> **Código**: `src/components/shared/AnalyticsProvider.tsx:4-5`

### OBS-ANA-002: Rastreamento de Clique em CTA
Quando um botão CTA do hero for clicado, o sistema deve disparar um evento GA4 `cta_click` com `cta_type` (`'consumidor' | 'fornecedor'`) e `location` (`'hero' | 'final_cta'`).
> **Código**: `src/lib/analytics.ts:53-55`, `src/components/home/HeroSection.tsx:8`

### OBS-ANA-003: Rastreamento de Uso da Calculadora
Quando uma calculadora produzir um resultado, o sistema deve disparar um evento GA4 `calculator_use`.
> **Código**: `src/hooks/useCalculator.ts:46-59`

### OBS-ANA-004: Rastreamento de Abertura de FAQ
(Definido mas ainda não conectado ao componente) O sistema deve suportar eventos GA4 `faq_open` com `question_index`.
> **Código**: `src/lib/analytics.ts:61-63`

### OBS-ANA-005: Rastreamento de Captura de Lead
(Definido mas ainda não conectado) O sistema deve suportar eventos GA4 `lead_capture` com `persona` e `has_plan`.
> **Código**: `src/lib/analytics.ts:65-67`

---

## 7. Prova Social e Conteúdo

### OBS-PRF-001: Métricas da Seção de Prova
O sistema deve exibir 3 métricas de consumidor (500+ usuários, 4.8★ NPS, 12% desconto médio) com destaque dinâmico opcional baseado no parâmetro de URL `perfil`.
> **Código**: `src/components/home/ProofSection.tsx:9-12,25-39`

### OBS-PRF-002: Exibição Condicional de Métricas de Fornecedor
Quando o parâmetro de URL `perfil=fornecedor` estiver presente, o sistema deve exibir adicionalmente 3 métricas de fornecedor (3 cooperativas, 1.000 kWh, R$0,40/kWh).
> **Código**: `src/components/home/ProofSection.tsx:42-54`

### OBS-PRF-003: Selo de Conformidade ANEEL
O sistema deve exibir um selo de conformidade regulatória ANEEL no rodapé e na ProofSection, citando "RN 687/2015".
> **Código**: `src/components/home/ProofSection.tsx:57-66`, `src/components/layout/Footer.tsx:57-59`

### OBS-PRF-004: Timeline Como Funciona
O sistema deve exibir 3 passos como uma timeline conectada: horizontal (desktop ≥768px) com uma linha de conexão, vertical (mobile <768px) com segmentos de linha verticais.
> **Código**: `src/components/home/HowItWorksSection.tsx:19-94`

### OBS-PRF-005: Accordion de FAQ
O sistema deve exibir FAQs como um accordion onde apenas um item fica aberto por vez, com animação de rotação do chevron e atributos de acessibilidade `aria-expanded`/`aria-controls`.
> **Código**: `src/components/plans/FaqAccordion.tsx`

### OBS-PRF-006: Dica de Contato no FAQ
Quando o usuário abrir 3+ itens de FAQ únicos sem clicar em um CTA, o sistema deve exibir uma dica: "Não encontrou o que procurava? Fale com a gente →".
> **Código**: `src/hooks/useFaqAccordion.ts:11-26`, `src/components/plans/FaqAccordion.tsx:57-63`

---

## 8. Sistema de Design

### OBS-DS-001: Tokens de Cor
O sistema deve definir todas as cores como propriedades CSS customizadas prefixadas `--color-solar-*` via Tailwind v4 `@theme inline` em `globals.css`.
> **Código**: `src/app/globals.css:8-36`

### OBS-DS-002: Família Tipográfica Única
O sistema deve usar Inter como a única família tipográfica, carregada via `next/font/google` com `subsets: ['latin']` e exposta como variável CSS `--font-inter`.
> **Código**: `src/app/layout.tsx:9-12`, `src/app/globals.css:34-35`

### OBS-DS-003: Movimento Reduzido
O sistema deve respeitar `prefers-reduced-motion: reduce` definindo durações de animação/transição para 0.01ms e desabilitando scroll suave.
> **Código**: `src/app/globals.css:73-82`

### OBS-DS-004: Classes de Animação
Onde `prefers-reduced-motion: no-preference`, o sistema deve fornecer classes utilitárias `.animate-fade-in` (0.4s fade+translateY) e `.animate-slide-in` (0.3s translateX).
> **Código**: `src/app/globals.css:85-93`

### OBS-DS-005: Sistema de Botões
O sistema deve fornecer 3 variantes de botão (primary, secondary, outline) × 3 tamanhos (sm, md, lg), com padding, border-radius, anéis de focus-visible consistentes e prop `href` opcional para renderização como link.
> **Código**: `src/components/shared/Button.tsx`

### OBS-DS-006: Ritmo de Seções
O sistema deve alternar fundos de seção usando `SectionWrapper` com `bg="white"` (padrão) e `bg="alt"` (`--color-solar-bg-alt`, o amarelo claro).
> **Código**: `src/components/shared/SectionWrapper.tsx`

---

## 9. Segurança

### OBS-SEC-001: Headers de Segurança
O sistema deve servir `X-Frame-Options: DENY`, `X-Content-Type-Options: nosniff` e `Referrer-Policy: strict-origin-when-cross-origin` em todas as rotas.
> **Código**: `next.config.ts:13-24`

### OBS-SEC-002: Anti-Spam do Formulário
O sistema deve implementar anti-spam duplo: portão de tempo mínimo de 3 segundos + campo honeypot oculto. Bots que preenchem o honeypot ou enviam rápido demais recebem um falso sucesso sem transmissão de dados.
> **Código**: `src/hooks/useContactForm.ts:8-9,48-53,65-66`

### OBS-SEC-003: Sanitização de Entrada
O sistema deve validar todas as entradas do formulário no lado do servidor... [NÃO OBSERVADO — não existe validação server-side]. Toda validação é executada apenas no lado do cliente.

---

## 10. Estados de Carregamento e Resiliência

### OBS-LD-001: Suspense da Seção de Prova
O sistema deve envolver ProofSection em uma boundary React `<Suspense>` com fallback de skeleton que preserva o layout (3 skeletons de card sobre fundo verde).
> **Código**: `src/app/page.tsx:11-26,33-35`

### OBS-LD-002: Suspense do Formulário de Contato
O sistema deve envolver ContactForm em uma boundary `<Suspense>` com skeletons de input como fallback.
> **Código**: `src/app/contato/page.tsx:16-25,40-44`

### OBS-LD-003: Resiliência na Busca de Dados
Quando a busca de dados assíncrona falhar, o sistema deve capturar o erro, registrar no console e retornar um array vazio em vez de quebrar.
> **Código**: Todos os serviços em `src/services/`

### OBS-LD-004: Delay Simulado nos Serviços
O sistema deve simular latência de rede com um `setTimeout` de 500ms em todas as funções de serviço.
> **Código**: ex., `src/services/servicePlans.ts:9-11`

---

## 11. Preparação para Internacionalização

### OBS-I18N-001: Atributo de Idioma
O sistema deve definir `lang="pt-BR"` no elemento HTML.
> **Código**: `src/app/layout.tsx:43`

### OBS-I18N-002: Formatação Consciente de Locale
O sistema deve formatar todos os valores monetários usando o locale `pt-BR`.
> **Código**: `src/lib/calculator.ts:20-23`

### OBS-I18N-003: Arquitetura i18n Futura
(Declarado em PRODUCT.md mas não implementado) O sistema deve ser arquitetado para suportar o prefixo de rota `/es/` para internacionalização em espanhol.
> **Fonte**: `.specs/PRODUCT.md:57`
