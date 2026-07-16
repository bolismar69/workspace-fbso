# Product

## Register

product

## Users

- **Controller Tributário**: Simula carga tributária para diferentes NCMs e jurisdições no departamento fiscal
- **CFO / Diretor Financeiro**: Projeta impacto da reforma no fluxo de caixa para planejamento estratégico
- **Consultor Tributário**: Gera cenários comparativos para múltiplos clientes
- **Desenvolvedor Integrador**: Consome API de cálculo para integrar em ERPs

## Product Purpose

O **TaxNexus Portal (TaaS — Tax as a Service)** permite a empresas brasileiras simularem o impacto financeiro da Reforma Tributária 2026. Compara o sistema atual (PIS, COFINS, ICMS, ISS, IPI) com o novo IVA Dual (CBS, IBS, Imposto Seletivo) via visualização lado a lado com gráficos. MVP Fase 0 com autenticação simplificada por CNPJ e simulação por NCM/localidade.

## Brand Personality

**Técnico, Confiável, Direto**. A interface deve transmitir precisão fiscal sem parecer burocrática. Linguagem de negócio tributário, não jargão técnico de desenvolvimento. Confiança através de clareza numérica, não através de animações ou decoração.

## Anti-references

- Não deve parecer um dashboard genérico de analytics (evitar cards idênticos em grid, hero metrics, glassmorphism)
- Não deve usar linguagem informal ou tom de startup de consumo
- Não deve ter paleta "SaaS cream" (bege/quente como padrão)
- Não deve ser excessivamente animado ou "delightful" — é uma ferramenta fiscal, não um produto de consumo
- Evitar dark mode roxo/neon e glassmorphism

## Design Principles

1. **Precisão acima de decoração**: Números são o produto. Formatação monetária correta (R$, pt-BR) é prioridade máxima
2. **Caminho mínimo para o resultado**: CNPJ → localidade → NCM → resultado. Sem distrações no fluxo principal
3. **Comparação visual clara**: Sistema atual vs. reforma sempre visíveis lado a lado
4. **Confiança via transparência**: Toda simulação tem ID rastreável, dados claros, sem "mágica"

## Accessibility & Inclusion

- WCAG 2.1 Nível AA como alvo mínimo
- Alto contraste para dados numéricos (4.5:1 mínimo em texto de corpo)
- Navegação por teclado completa no formulário de simulação
- Labels claros em todos os campos de formulário
- Suporte a leitores de tela nos resultados numéricos
