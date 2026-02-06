# GitHub Actions (local) - FBSO

Este diretório contém **composite actions locais** para reutilizar “primitivos” do pipeline (detectar/validar/enriquecer/rotear services) em novos workflows, sem precisar duplicar YAML.

## Ações disponíveis

### `detect-solutions-changed`
Detecta solutions modificadas entre dois SHAs e emite:
- `solutions`: JSON array string de objetos `{path}`
- `paths`: JSON array string de paths

Uso:
```yaml
- uses: ./.github/actions/detect-solutions-changed
  id: detect
  with:
    stack: backend
    base_sha: ${{ github.event.before }}
    head_sha: ${{ github.sha }}
    python-version: ${{ vars.CI_PYTHON_VERSION || '3.11' }}

- run: echo "${{ steps.detect.outputs.paths }}"
```

### `validate-input-services`
Valida o *layout* do payload `services` (ex.: workflow_dispatch/workflow_call) e emite `solutions/paths`.

### `validate-identified-solutions`
Valida `services` contra o inventory `architecture/governance/config/manager-solutions-inventory.json` e retorna apenas `active` (full objects).

### `apply-global-types-specification`
Preenche `specification` ausente/vazia usando `architecture/governance/config/global-types-solution.json` e emite `solutions/paths`.

### `route-services-by-platform`
Roteia um JSON array string de services em 4 outputs:
- `java_services`, `go_services`, `python_services`, `csharp_services`

## Observações
- Todas as actions baseadas em Python aceitam `python-version` (default `3.11`).
- `skip-checkout` existe para evitar checkout duplicado quando o caller já fez `actions/checkout@v4`.
