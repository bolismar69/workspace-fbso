# Kestra - CI/CD scripts

Scripts utilitários para publicação de flows/ativos do Kestra a partir de CI/CD.

## publish_flows.sh

Publica arquivos `.yml/.yaml` encontrados em `orchestration/kestra/**/flows/`.

### Variáveis de ambiente

- `KESTRA_URL` (obrigatória) — exemplo: `https://kestra.suaempresa.com`
- `KESTRA_API_FLOWS_ENDPOINT` (opcional) — default: `/api/v1/flows`

Autenticação (use **uma** das opções):

- **Bearer token**: `KESTRA_API_TOKEN`
- **Basic auth**: `KESTRA_BASIC_AUTH` no formato `user:pass`

### Exemplo (local)

```bash
export KESTRA_URL="http://localhost:8080"
export KESTRA_API_TOKEN="..."
./devops/scripts/kestra/publish_flows.sh
```
