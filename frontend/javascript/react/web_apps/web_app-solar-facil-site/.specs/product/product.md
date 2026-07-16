# Product — Solar Fácil Site

> Descrição do produto e visão (Diátaxis: Explanation).
> Gerado por `documentation-writer` em 2026-07-08.
> Integra conteúdo do `.specs/PRODUCT.md` original (2026-07-05).

---

## 1. Registro

**Marca**: Solar Fácil
**Site**: https://www.solarfacil.com.br
**Segmento**: Energia Solar Compartilhada (ANEEL)

---

## 2. Usuários (Personas)

| Persona | Quem são | Necessidade Principal |
|---|---|---|
| **Consumidores** | Pessoas físicas e pequenas empresas. Não especialistas em energia. Acessam via mobile, horário comercial, após conta de luz alta. | "Quero reduzir minha conta de luz sem instalar painéis solares." |
| **Fornecedores/Produtores** | Quem já tem painéis solares e gera excedente. Perfil mais técnico. | "Quero rentabilizar meu excedente de energia." |
| **Cooperativas** | Parceiros B2B que operam distribuição de energia. | "Preciso de credibilidade e informações sobre integração." |

---

## 3. Propósito do Produto

A **Solar Fácil** conecta produtores de energia solar excedente com consumidores que querem desconto na conta de luz, via cooperativas regulamentadas pela ANEEL.

O site é a **porta de entrada**: educa sobre energia compartilhada, calcula a economia potencial, apresenta os planos e converte visitantes em leads qualificados.

**Critério de sucesso**: o visitante entender em 30 segundos que pode economizar sem investir nada, usar a calculadora, e iniciar contato.

---

## 4. Personalidade da Marca

**Confiança · Sustentabilidade · Inovação · Amigável · Prático**

Uma marca que fala a língua das pessoas. Não é corporativa nem burocrática — é próxima, direta e transparente.

**Tom de voz**: "A gente resolve." — Informativo sem ser professoral, entusiasmado sem ser vendedor, técnico sem ser incompreensível.

---

## 5. Anti-Referências

- ❌ **Startup tech fria** — Sem dark mode futurista, neon, glassmorphism
- ❌ **Site governamental** — Sem bandeiras, brasões, azul imperial
- ❌ **Site genérico de energia** — Sem clichês verdes, ícones de folha/gota

---

## 6. Princípios de Design

1. **"Energia que aproxima"** — Qualquer pessoa entende em segundos
2. **"Transparência radical"** — Números claros, sem asteriscos
3. **"Brasil real, sem clichês"** — Brasileiro sem bandeira ou tropicalismo
4. **"Confiança ancorada em fatos"** — Prova social, ANEEL, depoimentos reais
5. **"Simplicidade que respeita"** — Complexidade regulatória nos bastidores

---

## 7. Acessibilidade & Inclusão

- WCAG AA como padrão mínimo
- Contraste ≥ 4.5:1 (texto), ≥ 3:1 (texto grande)
- Navegação completa por teclado
- Suporte a leitores de tela
- `prefers-reduced-motion: reduce` respeitado
- pt-BR primário, arquitetura preparada para `/es/`

---

## 8. Contexto Técnico

| Item | Detalhe |
|---|---|
| Stack | Next.js 16 + React 19 + Tailwind CSS v4 + TypeScript |
| Hospedagem | Standalone (DigitalOcean ou similar) |
| Design Tokens | CSS custom properties `--color-solar-*` |
| Fonte | Inter (sans-serif) |
| Cor Primária | `#1E5631` (Verde Terra) |

---

Última atualização: 2026-07-08
