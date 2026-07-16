# Product

## Register

brand

## Users

**Consumidores** — Pessoas físicas e pequenas empresas que querem reduzir a conta de luz sem instalar painéis solares. Não são especialistas em energia; buscam economia sem complexidade. Acessam o site principalmente via mobile, durante o horário comercial, muitas vezes após receberem uma conta de luz alta.

**Fornecedores/Produtores** — Quem já tem painéis solares instalados e gera excedente. Querem rentabilizar esse excedente de forma simples e legalizada. Perfil mais técnico que o consumidor, mas ainda valorizam clareza e praticidade.

**Cooperativas** — Parceiros B2B que operam a distribuição de energia. Precisam de credibilidade e informações sobre integração com a plataforma.

## Product Purpose

A Solar Fácil conecta produtores de energia solar excedente com consumidores que querem desconto na conta de luz, via cooperativas regulamentadas pela ANEEL. O site é a porta de entrada: educa sobre energia compartilhada, calcula a economia potencial, apresenta os planos e converte visitantes em leads qualificados.

**Sucesso é:** o visitante entender em 30 segundos que pode economizar sem investir nada, usar a calculadora, e iniciar contato.

## Brand Personality

**Confiança · Sustentabilidade · Inovação · Amigável · Prático**

Uma marca que fala a língua das pessoas. Não é corporativa nem burocrática — é próxima, direta e transparente. Sustentabilidade não é bandeira verde; é consequência natural de um modelo de negócio inteligente. Inovação não é buzzword tech; é fazer algo complexo (regulação ANEEL, distribuição de energia) parecer simples.

**Tom de voz:** "A gente resolve." — Informativo sem ser professoral, entusiasmado sem ser vendedor, técnico sem ser incompreensível.

**Referência visual:** [Portal Solar](https://www.portalsolar.com.br/energia-solar) — abordável mas profissional, com hierarquia clara de conteúdo, confiança ancorada em autoridade de mercado, e CTAs integrados naturalmente ao fluxo de leitura.

## Anti-references

- **Nada de startup tech fria** — Sem dark mode futurista, neon, glassmorphism, ou estética "SaaS dashboard". Não somos uma ferramenta de software.
- **Nada de governo/estatal** — Sem bandeiras do Brasil, brasões, azul imperial, ou estética de site de prefeitura. Credibilidade ≠ burocracia visual.
- **Nada de site genérico de energia** — Sem clichês verdes com folhas, gradientes ecológicos, ícones de gota d'água e lâmpada, ou o azul-marinho-cinza das utilities tradicionais.

## Design Principles

1. **"Energia que aproxima"** — Toda pessoa, independente de conhecimento técnico, entende o que fazemos em segundos. A linguagem visual e verbal elimina atritos, não os cria.

2. **"Transparência radical"** — Números são claros e verificáveis. A calculadora mostra exatamente a economia. Nada de letras miúdas, asteriscos, ou promessas vagas. O que aparece na tela é o que o usuário recebe.

3. **"Brasil real, sem clichês"** — A identidade é brasileira porque o público é brasileiro, mas sem apelar para verde-amarelo, bandeiras, ou tropicalismo forçado. A brasilidade está na linguagem, nas referências culturais, nos valores em reais — não na paleta.

4. **"Confiança ancorada em fatos"** — Prova social com números reais, chancela ANEEL visível, depoimentos verificáveis. Cada afirmação no site pode ser sustentada por um dado, uma norma, ou um cliente real.

5. **"Simplicidade que respeita"** — O modelo de negócio é inovador e regulado (ANEEL), mas o site não transforma isso em complexidade para o usuário. A complexidade fica nos bastidores; a interface entrega clareza.

## Accessibility & Inclusion

- **WCAG AA** como padrão mínimo em todas as páginas
- Contraste ≥ 4.5:1 para texto corrido, ≥ 3:1 para texto grande
- Navegação completa por teclado (focus rings visíveis, tab order lógica)
- Suporte a leitores de tela (ARIA labels, landmarks semânticas, alt text descritivo)
- `prefers-reduced-motion: reduce` respeitado em todas as animações
- Testes com os 3 tipos mais comuns de daltonismo (deuteranopia, protanopia, tritanopia)
- **Idiomas:** pt-BR (primário) com arquitetura preparada para espanhol (internacionalização futura — rotas com prefixo `/es/`)

## Technical Context

- **Stack:** Next.js 16 + React 19 + Tailwind CSS v4 + TypeScript
- **Hospedagem:** Vercel (Fluid Compute)
- **Design tokens:** CSS custom properties com prefixo `--color-solar-*` em `src/app/globals.css`
- **Font:** Inter (sans-serif)
- **Cores atuais:** Primary #1E5631 (verde escuro), Secondary #A5C9CA (teal/verde-água), Background #ffffff, Background alt #ffffbf (amarelo claro)
