#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)"

KESTRA_URL="${KESTRA_URL:-}"
KESTRA_API_FLOWS_ENDPOINT="${KESTRA_API_FLOWS_ENDPOINT:-/api/v1/flows}"

if [[ -z "$KESTRA_URL" ]]; then
  echo "KESTRA_URL é obrigatório (ex.: https://kestra.suaempresa.com)" >&2
  exit 2
fi

API_URL="${KESTRA_URL%/}${KESTRA_API_FLOWS_ENDPOINT}"

AUTH_ARGS=()
if [[ -n "${KESTRA_API_TOKEN:-}" ]]; then
  AUTH_ARGS+=( -H "Authorization: Bearer ${KESTRA_API_TOKEN}" )
elif [[ -n "${KESTRA_BASIC_AUTH:-}" ]]; then
  AUTH_ARGS+=( -u "${KESTRA_BASIC_AUTH}" )
fi

mapfile -t FLOW_FILES < <(find "$ROOT_DIR/orchestration/kestra" -type f \( -name "*.yml" -o -name "*.yaml" \) -path "*/flows/*" | sort)

if [[ ${#FLOW_FILES[@]} -eq 0 ]]; then
  echo "Nenhum flow encontrado em orchestration/kestra/**/flows/*.yml" >&2
  exit 0
fi

echo "Publicando ${#FLOW_FILES[@]} flow(s) em: $API_URL"

FAILED=0
for f in "${FLOW_FILES[@]}"; do
  rel="${f#$ROOT_DIR/}"
  echo "- $rel"

  # Publicação via API (endpoint parametrizável). O Kestra deve validar o YAML e responder 2xx em sucesso.
  # Caso seu ambiente exija um método diferente (PUT vs POST), ajuste KESTRA_API_FLOWS_ENDPOINT ou o comando abaixo.
  http_code=$(curl -sS -o /tmp/kestra_publish.out -w "%{http_code}" \
    -X POST "$API_URL" \
    -H "Content-Type: application/yaml" \
    "${AUTH_ARGS[@]}" \
    --data-binary "@$f" || true)

  if [[ "$http_code" != 2* ]]; then
    echo "  ERRO: HTTP $http_code ao publicar $rel" >&2
    sed -n '1,200p' /tmp/kestra_publish.out >&2 || true
    FAILED=1
  fi

done

if [[ "$FAILED" -ne 0 ]]; then
  echo "Falha ao publicar um ou mais flows." >&2
  exit 1
fi

echo "Publicação concluída com sucesso."
