# PROMPT: PORTÃO DE VALIDAÇÃO DE PLANO DE GESTÃO DE MUDANÇAS
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Auditor de Governança de Projetos. Valide APENAS o arquivo em `DOC_PATH`.

## Checklist de Compliance

1. **Cabeçalho:** Projeto, Documentos Base (001-PROJECT-CHARTER, 060-EAP-WBS, 080-PLANO-RISCOS), Data, Versão, Metodologia preenchidos? Status "Em análise" ou "Em revisão"?
2. **Formulário de Solicitação de Mudança:** Seção 1 presente com campos para identificação (ID sequencial `CCR-{NN}`), tipo, descrição, justificativa e impacto estimado? (LACUNA = FAIL)
3. **Composição do CCB:** Seção 2 define papéis, membros, voto e responsabilidades? Presidente, representantes de negócio e técnico presentes? (LACUNA = FAIL)
4. **Classificação de Impacto:** Seção 3 define critérios objetivos e mensuráveis para ALTO/MÉDIO/BAIXO (percentuais, prazos, marcos)?
5. **Fluxo de Aprovação:** Seção 4 descreve o workflow completo com decisão de aprovação/rejeição por nível de classificação e prazos de decisão? (LACUNA = FAIL)
6. **Registro de Mudanças:** Seção 5 (Change Log) presente com rastreabilidade completa (CCR, classificação, status, aprovador, decisão, itens impactados)?
7. **Interação com Riscos:** O plano considera a relação entre mudanças e o inventário de riscos do 080-PLANO-RISCOS? (LACUNA = FAIL)

Retorne `{PASS}` e `[STATUS: Em revisão]` ou `{FAIL, VIOLATIONS: [...]}`.
