# FACTORY-NOTIFICATION-CAPGEMINI — Notificação de Resultado

- **Fábrica:** Capgemini
- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG — FBSO Platform
- **Modo:** Full — Projeto Completo (62 US · 18 Features)
- **Data:** 31/07/2026
- **Status:** 🔴 NÃO APROVADA — REALINHAMENTO SOLICITADO

---

## 1. Resultado da Avaliação

Prezada equipe Capgemini,

Agradecemos o envio da estimativa para o projeto FBSO Platform. Após análise técnica pela equipe de arquitetura da FBSO.ORG, sua estimativa **não foi aprovada** nesta rodada. Solicitamos **realinhamento e reenvio** conforme orientações abaixo.

### Estimativa Apresentada

| Métrica | Valor |
|:---|---:|
| Total de Horas | 11,680h |
| QA | 11.0% (mínimo exigido: 25%) |
| Arquitetura/SRE | 11.0% (mínimo exigido: 5%) |
| Prazo Declarado | 3 meses |
| Time Estimado | 15 pessoas |

---

## 2. Não-Conformidades Identificadas

- **QA:** QA (qualidade) abaixo do mínimo de 25% exigido. É necessário redistribuir horas para garantir cobertura adequada de testes unitários, integração, E2E e performance.
- **Prazo:** Divergência entre prazo declarado e horas×time. O prazo informado é inconsistente com o volume de horas e tamanho do time proposto.
- **PIB:** Proximidade à Baseline Interna (PIB) abaixo do threshold de 0.50. A estimativa está significativamente distante da referência técnica interna da FBSO.ORG, baseada em PERT bottom-up validado por duas metodologias independentes.

---


---

## 3. Observações Adicionais (Análise Retrospectiva PIB)

Além das não-conformidades específicas da Capgemini listadas acima, a análise retrospectiva com a métrica PIB (Proximidade à Baseline Interna) identificou problemas **sistêmicos** que afetaram múltiplas fábricas:

### 3.1 Comentários Genéricos

A coluna `comentarios` do schema CSV foi preenchida com texto genérico: *"seguimos especificamente o material reportado"*. O RFQ solicita que esta coluna contenha o **racional detalhado da estimativa**: metodologia utilizada, premissas adotadas por épico, e justificativas para os números apresentados. Textos genéricos impedem a FBSO.ORG de avaliar a qualidade do raciocínio técnico por trás da estimativa.

**Ação solicitada:** Preencher `comentarios` com no mínimo 200 caracteres por linha, detalhando a metodologia de estimativa para cada épico.

### 3.2 Estimativa Plana (Flat Estimate)

Sua estimativa apresenta **pouca ou nenhuma diferenciação de esforço entre os épicos**, apesar de os épicos terem complexidades e quantidades de User Stories significativamente diferentes:

| Épico | US | Complexidade | Sua Estimativa |
|:---|---:|:---|:---|
| EP-0001 Portal Admin | 7 | Média | Mesmo valor dos demais |
| EP-0002 Clientes | 16 | Alta | Mesmo valor |
| EP-0003 RBAC | 16 | Alta | Mesmo valor |
| EP-0004 Portal Cliente | 23 | Média | Mesmo valor |

É **improvável** que épicos com 7, 16, 16 e 23 User Stories tenham exatamente o mesmo esforço. Isso sugere que a estimativa foi feita por rateio uniforme, sem análise individual do escopo de cada épico.

**Ação solicitada:** Diferenciar o esforço entre épicos conforme a complexidade e quantidade de User Stories. Justificar a variação (ou ausência dela) na coluna `comentarios`.

### 3.3 Independência do Processo

A FBSO.ORG identificou que sua estimativa apresenta valores **idênticos** aos de outras duas fábricas participantes em todas as 18 colunas do schema (horas_dev, horas_qa, horas_arch, horas_devops, horas_gestao, complexidade, e comentários para todos os 4 épicos). A probabilidade de 3 fábricas independentes chegarem aos mesmos valores é efetivamente zero.

Possíveis explicações incluem: (a) uso de template compartilhado sem revisão individual, (b) coordenação entre participantes, ou (c) preenchimento automatizado sem análise. Qualquer uma dessas situações compromete a independência do processo de bidding.

**Ação solicitada:** Explicação formal sobre a identidade dos valores. A Capgemini deve demonstrar que sua estimativa foi produzida de forma independente e baseada em análise própria do escopo.

---

## 3. Orientações para Realinhamento

Solicitamos que a Capgemini:

1. **Redistribua as horas** garantindo proporcionalidade:
   - `horas_qa` ≥ 25% do `total_horas` (percentual, NÃO valor fixo)
   - `horas_arch` ≥ 5% do `total_horas` (percentual, NÃO valor fixo)
   - `total_horas` = `horas_dev` + `horas_qa` + `horas_arch` + `horas_devops` + `horas_gestao`

2. **Diferencie o esforço entre épicos** conforme complexidade e quantidade de User Stories. Épicos com 23 US não devem ter o mesmo esforço que épicos com 7 US.

3. **Revise o prazo declarado** para consistência com a fórmula `total_horas / (time × 160h)`. A divergência deve ser ≤ 50%.

4. **Preencha a coluna `comentarios`** com o racional detalhado da estimativa (mínimo 200 caracteres por linha): metodologia utilizada, premissas adotadas por épico, e justificativas para os números. Textos genéricos como "seguimos o material reportado" NÃO serão aceitos.

5. **Reenvie o CSV** preenchido até o novo prazo que será comunicado em seguida.

---

## 4. Próximos Passos

1. A FBSO.ORG reabrirá o RFQ com instruções atualizadas
2. Novo prazo de entrega será comunicado em breve
3. As estimativas realinhadas passarão por nova validação DTA + PIB
4. Dúvidas técnicas podem ser enviadas para o canal de comunicação do RFQ

---

## 5. Confidencialidade

Este documento contém informações confidenciais sobre o processo de seleção. Os dados aqui apresentados referem-se exclusivamente à estimativa da Capgemini e aos critérios objetivos de avaliação. Nenhuma informação sobre outras fábricas participantes é divulgada.

---

🤖 *Notificação gerada automaticamente — Fase 7 do Sourcing & Factory Bidding (Full Mode). Documento confidencial.*
