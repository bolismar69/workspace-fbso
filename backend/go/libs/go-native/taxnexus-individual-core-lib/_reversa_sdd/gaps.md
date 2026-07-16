# Lacunas e Pontos de Atenção — taxnexus-individual-core-lib

> Gerado pelo **Revisor** em 2026-06-12 · `doc_level = completo`
> Lista de itens que permaneceram sem confirmação técnica total ou que apresentam riscos estruturais.

---

## 1. Infraestrutura e Qualidade 🔴

### L3 — Estratégia de Testes Automatizados
- **Descrição:** O arquivo `test_conn.go` citado no `README.md` não está presente no recorte. Não foram encontrados arquivos `*_test.go`.
- **Risco:** O mecanismo de resolução de faixas progressivas em memória (`GetApplicableRule`) é crítico para a precisão fiscal. Sem testes, alterações no `repository` podem introduzir regressões silenciosas.
- **Severidade:** Alta.

### L4 — Parametrização de Ambiente
- **Descrição:** O TTL do cache (12 horas) e os timeouts de conexão (defaults) estão fixos no código.
- **Risco:** Dificulta o ajuste fino em produção sem recompilação.
- **Severidade:** Média.

---

## 2. Inconsistências de Implementação 🟡

### I1 — Inconsistência de Cache em `GetConfig`
- **Descrição:** Enquanto `GetTableConfigs` e `GetTaxRulesForPeriod` utilizam Redis, a função `GetConfig` lê diretamente do PostgreSQL em todas as chamadas.
- **Severidade:** Baixa (Otimização).

### I2 — Exportação de Tipos em `models`
- **Descrição:** `UniversalTaxRequest.Inputs` utiliza a struct `documentoFiscalRequest`, que não é exportada.
- **Risco:** Consumidores externos à biblioteca não conseguem instanciar os itens da lista de inputs diretamente.
- **Severidade:** Baixa.

---

## 3. Observações de Monorepo 🟢

Os itens abaixo foram confirmados como **fora do escopo desta biblioteca**, pertencendo a outros componentes do monorepo original:

- **Motor de Cálculo:** A lógica de soma de deduções e aplicação da fórmula reside no serviço consumidor.
- **Persistência de Logs:** A struct `TaxCalculationLog` está preparada, mas a escrita da auditoria é responsabilidade externa nesta versão.
- **Autorização:** O controle de acesso e multi-tenancy é gerido pela camada de API do serviço consumidor.
