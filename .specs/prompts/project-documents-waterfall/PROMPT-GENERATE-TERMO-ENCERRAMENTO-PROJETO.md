# PROMPT: GERADOR DE TERMO DE ENCERRAMENTO DE PROJETO
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Gerente de Projetos Sênior e Redator Técnico de Contratos e Propostas. Sua missão é criar o **Termo de Encerramento de Projeto** que formaliza o fechamento do projeto: confirma as entregas contra o Charter, registra o aceite final do Sponsor, consolida as lições aprendidas e libera formalmente o time e os recursos.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Lista: `[{PROJECT_COMPLETE_PATH_NAME}/001-PROJECT-CHARTER-{PROJECT_ID_NAME}.md, {PROJECT_COMPLETE_PATH_NAME}/105-TERMO-ACEITE-{PROJECT_ID_NAME}.md, {PROJECT_COMPLETE_PATH_NAME}/110-LICOES-APRENDIDAS-{PROJECT_ID_NAME}.md]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["senior-pm", "contract-and-proposal-writer"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** 001-PROJECT-CHARTER (objetivos, entregas e critérios de sucesso), 105-TERMO-ACEITE (entregas formalmente aceitas) e 110-LICOES-APRENDIDAS (lições consolidadas)
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback (7 Seções)

```
# Termo de Encerramento de Projeto: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 001-PROJECT-CHARTER, 105-TERMO-ACEITE, 110-LICOES-APRENDIDAS |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## O que é um Termo de Encerramento de Projeto?

O **Termo de Encerramento de Projeto** é o documento de governança que formaliza o fechamento do projeto no pipeline WATERFALL. Ele confirma que todas as entregas contratadas no Charter foram concluídas e aceitas, registra o aceite final do Sponsor, consolida as lições aprendidas, e libera formalmente o time, os recursos e as responsabilidades administrativas e financeiras.

### Conexão com o Pipeline

- **UPSTREAM:** Consome os objetivos e entregas do 001-PROJECT-CHARTER, as entregas formalmente aceitas do 105-TERMO-ACEITE e as lições consolidadas da 110-LICOES-APRENDIDAS
- **DOWNSTREAM:** Documento final do pipeline — encerra o ciclo de vida do projeto e habilita o handover para operação

---

## 1. Sumário Executivo de Encerramento

| Campo | Detalhe |
|-------|---------|
| **Objetivo do Projeto (Charter)** | {objetivo conforme 001-PROJECT-CHARTER} |
| **Resultado Alcançado** | {resumo do que foi entregue} |
| **Data de Encerramento** | {DATA DE ENCERRAMENTO} |
| **Motivo do Encerramento** | Conclusão / Cancelamento / Substituição / Outro |
| **Situação Final do Escopo** | {escopo entregue conforme baseline / desvios aprovados} |

---

## 2. Confirmação de Entregas vs Charter

| Entrega (Charter Seção 4) | Critério de Aceitação | Aceite (105) | Evidência | Concluído? |
|---------------------------|------------------------|--------------|-----------|------------|
| {Entrega 1} | {critério} | {data do aceite} | {evidência} | ✅ Sim / ⚠️ Parcial / ❌ Não |
| {Entrega 2} | {critério} | {data do aceite} | {evidência} | ✅ Sim / ⚠️ Parcial / ❌ Não |

### Pendências Não Resolvidas

| Pendência | Impacto | Plano de Ação | Responsável |
|-----------|---------|---------------|-------------|
| {se vazio: "NENHUMA — todas as entregas do Charter foram aceitas ✅"} | | | |

---

## 3. Aceite Final do Sponsor

Declaro, na qualidade de **Sponsor do Projeto**, que recebi e aceito formalmente as entregas deste projeto, conforme confirmado no 105-TERMO-ACEITE, e que o projeto está apto ao encerramento.

| Papel | Nome | Assinatura | Data |
|-------|------|------------|------|
| **Sponsor** | {nome} | {assinatura} | {data} |
| **Product Owner** | {nome} | {assinatura} | {data} |
| **Gestor do Projeto** | {nome} | {assinatura} | {data} |

---

## 4. Handover para Operação

### 4.1 Itens Transferidos

| Item | Destinatário | Data da Transferência | Documentação de Apoio |
|------|--------------|------------------------|------------------------|
| {Sistema/Infra} | {Time de Operação} | {data} | 095-MANUAL-OPERACAO, 100-MANUAL-USUARIO, 090-PLANO-TESTES |
| {Ambiente} | {Time de Infra} | {data} | {documentos de apoio} |

### 4.2 Responsabilidades Pós-Entrega

| Responsabilidade | Time Responsável | SLA / Acordo de Nível de Serviço |
|------------------|------------------|-----------------------------------|
| {Suporte nível 1} | {Time} | {SLA} |
| {Suporte nível 2} | {Time} | {SLA} |
| {Manutenção evolutiva} | {Time} | {acordo} |

---

## 5. Lições Aprendidas Consolidadas

Fonte: `110-LICOES-APRENDIDAS-{PROJECT_ID_NAME}.md` (documento completo de lições).

| Lição | Categoria | O que funcionou / O que falhou | Ação Futura Recomendada |
|-------|-----------|--------------------------------|--------------------------|
| {L1} | {Processo / Pessoas / Tecnologia} | {descrição} | {recomendação} |
| {L2} | {Processo / Pessoas / Tecnologia} | {descrição} | {recomendação} |

> **NOTA:** O registro completo de lições está no documento 110-LICOES-APRENDIDAS; este termo consolida apenas os destaques relevantes ao encerramento.

---

## 6. Liberação Formal do Time

Declaro que o time do projeto está formalmente **liberado** de suas atribuições no projeto, a partir de {DATA DE ENCERRAMENTO}, ficando à disposição da organização para novas alocações.

| Perfil | Recursos Liberados | Data | Observação |
|--------|--------------------|------|------------|
| {Perfil/Time} | {pessoas ou carga horária} | {data} | {observação} |

---

## 7. Encerramento Administrativo e Financeiro

| Item | Situação | Observação |
|------|----------|------------|
| **Contratos e fornecedores** | {encerrados / em encerramento} | {observação} |
| **Orçamento executado** | {valor} vs {orçamento planejado} | {desvio e justificativa} |
| **Recursos físicos e licenças** | {devolvidos / cancelados} | {observação} |
| **Documentação do projeto** | {arquivada no repositório} | {local} |
| **Auditoria / conformidade** | {concluída} | {observação} |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Baseline inicial do termo de encerramento | Gestão do Projeto |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se as 7 seções estiverem completas, todas as entregas do Charter confirmadas contra o 105-TERMO-ACEITE, e o aceite final do Sponsor registrado.
