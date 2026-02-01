# Kestra - ETL

Pipelines de ETL/ELT e rotinas de dados.

## Estrutura sugerida por pipeline

- `<pipeline>/flows/`
- `<pipeline>/scripts/` (SQL, Python, etc.)

## Regras

- Preferir namespaces `fbso.etl.*`
- Scripts versionados e referenciados pelo flow (evitar lógica grande “inline” no YAML)
- Backfills/reprocessamentos devem ser flows explícitos, versionados
