# ESTIMATE-RECEIPT.md — Guia de Recebimento de Estimativas
## Sourcing & Factory Bidding — Fase 4 — Bloco B

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | ESTIMATE-RECEIPT-v1.0 |
| **Versão** | 1.0 — Discovery-Level |
| **Data** | 03 de agosto de 2026 |
| **Modo** | `discovery` |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 03/08/2026 |

---

## 1. Procedimento de Recebimento

Quando uma fábrica enviar sua estimativa, o time operacional deve:

1. **Salvar o arquivo CSV** na pasta `estimates/` com o nome padrão:
   ```
   estimates/ESTIMATION-SCHEMA-{NOME-DA-FABRICA}.csv
   ```

2. **Verificar a integridade básica do arquivo:**
   - [ ] Arquivo CSV abre corretamente (sem corrupção)
   - [ ] Separador `;` (ponto e vírgula)
   - [ ] Header com 16 colunas conforme schema discovery
   - [ ] 4 linhas de dados (EP-0001 a EP-0004)
   - [ ] Colunas numéricas contêm números (não texto)
   - [ ] `total_horas = horas_dev + horas_qa + horas_arch + horas_devops + horas_gestao`
   - [ ] `time_estimado_pessoas` e `valor_estimado` preenchidos

3. **Registrar o recebimento** na tabela de controle abaixo

4. **NÃO modificar** o arquivo original da fábrica

---

## 2. Controle de Recebimento

| # | Fábrica | Arquivo | Data Receb. | Integridade OK? | Encaminhado para Validação? |
|---|---------|---------|:-----------:|:---------------:|:---------------------------:|
| 1 | | | | | |
| 2 | | | | | |
| 3 | | | | | |

---

## 3. Estrutura de Arquivos

```
estimates/
├── ESTIMATION-SCHEMA-{FABRICA-1}.csv    ← arquivo original da fábrica
├── ESTIMATION-SCHEMA-{FABRICA-2}.csv
├── ESTIMATION-SCHEMA-{FABRICA-3}.csv
├── ESTIMATE-VALIDATION-{FABRICA-1}.md    ← gerado na Fase 5
├── ESTIMATE-VALIDATION-{FABRICA-2}.md
└── ESTIMATE-VALIDATION-{FABRICA-3}.md
```

---

## 4. Critérios de Aceite para Validação

Para que uma estimativa seja encaminhada para a Fase 5 (Validação DTA), ela deve passar na verificação de integridade básica (Seção 1, passo 2). Arquivos que falharem na integridade devem ser devolvidos à fábrica para correção.

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 03/08/2026 | Criação inicial: guia de recebimento de estimativas | PMO |

---

🤖 *Sourcing & Factory Bidding — Fase 4. Guia operacional para recebimento de estimativas.*
