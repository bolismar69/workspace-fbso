# CI/CD Java (monorepo) — documentação dos workflows

## Escopo

Este documento descreve como o pipeline de CI/CD Java funciona no monorepo, incluindo:
- Workflows participantes e como se integram
- Scripts participantes
- Contratos de input/output (JSON)
- Regras de validação do inventory (governança)
- Passos de execução relevantes

Arquivos principais:
- Trigger/orquestração: [.github/workflows/ci-cd-trigger-java.yml](../ci-cd-trigger-java.yml)
- Detector (reusable): [.github/workflows/global-callables-detection-solutions-changed.yml](../global-callables-detection-solutions-changed.yml)
- Build Java (reusable): [.github/workflows/global-callables-java-build-solutions.yml](../global-callables-java-build-solutions.yml)
- Script de detecção/enriquecimento: [.github/scripts/detect_solutions_changed.py](../../scripts/detect_solutions_changed.py)
- Inventory (governança): [architecture/governance/config/manager-solutions.json](../../../architecture/governance/config/manager-solutions.json)
- Registries de governança (blueprints/merge):
  - Framework: [architecture/governance/config/global-frameworks.json](../../../architecture/governance/config/global-frameworks.json)
  - Type (stackType): [architecture/governance/config/global-types-solution.json](../../../architecture/governance/config/global-types-solution.json)
  - Solution: [architecture/governance/config/global-solutions.json](../../../architecture/governance/config/global-solutions.json)
- Env compartilhado Docker: [.github/docker.env](../../docker.env)
- Env compartilhado Docker Java: [.github/docker.java.env](../../docker.java.env)

---

## Visão geral do fluxo (call-chain)

1. **Trigger**: [.github/workflows/ci-cd-trigger-java.yml](../ci-cd-trigger-java.yml) (push/PR em `backend/java/**` ou `workflow_dispatch`)
2. **Detecção de solutions alteradas** (reusable):
   - Chama [.github/workflows/global-callables-detection-solutions-changed.yml](../global-callables-detection-solutions-changed.yml)
   - Executa o script [.github/scripts/detect_solutions_changed.py](../../scripts/detect_solutions_changed.py)
   - O script detecta paths alterados entre SHAs e **enriquece** os objetos com dados do inventory [architecture/governance/config/manager-solutions.json](../../../architecture/governance/config/manager-solutions.json)
3. **Preparação da matrix Java**:
   - Ainda dentro do trigger, o job `prepare_services` valida o contrato (campos obrigatórios)
   - Qualquer solution com inventory incompleto gera **WARNING** e é **removida da matrix**
4. **Build Java** (reusable):
   - Trigger chama [.github/workflows/global-callables-java-build-solutions.yml](../global-callables-java-build-solutions.yml)
   - Esse workflow faz `setup-java`, Maven build, resolve blueprint (merge de registries de governança) e executa Docker build/push
   - Carrega variáveis compartilhadas via [.github/docker.env](../../docker.env) e [.github/docker.java.env](../../docker.java.env)

---

## Schema dos registries de governança

Schemas atuais esperados:
- [architecture/governance/config/global-frameworks.json](../../../architecture/governance/config/global-frameworks.json)
  - chave: `globalFramework` (array)
- [architecture/governance/config/global-types-solution.json](../../../architecture/governance/config/global-types-solution.json)
  - chave: `globalTypeSolution` (array)
- [architecture/governance/config/global-solutions.json](../../../architecture/governance/config/global-solutions.json)
  - chave: `globalSolution` (array)

---

## Contrato do inventory (manager-solutions.json)

O inventory em [architecture/governance/config/manager-solutions.json](../../../architecture/governance/config/manager-solutions.json) é a fonte de verdade.

Para solutions Java, os atributos abaixo são tratados como **obrigatórios** no trigger:
- `stack`
- `platform`
- `platformVersion`
- `platformDistributor`
- `framework`
- `type`
- `name`
- `status`
- `path`

E é recomendado também ter:
- `docker.argCompilationMode` (atualmente logado como WARNING quando ausente, mas não bloqueia por padrão)

Exemplo (inventory item):
```json
{
  "stack": "backend",
  "platform": "java",
  "platformVersion": "21",
  "platformDistributor": "temurin",
  "framework": "spring",
  "type": "queue_workers",
  "name": "ms-transaction-processor-queue",
  "status": "active",
  "path": "backend/java/spring/queue_workers/ms-transaction-processor-queue",
  "docker": { "argCompilationMode": "jvm" }
}
```

---

## Outputs do detector (workflow_call)

O reusable detector [.github/workflows/global-callables-detection-solutions-changed.yml](../global-callables-detection-solutions-changed.yml) retorna:

- `outputs.solutions`: JSON array de objetos (solutions) **enriquecidos pelo inventory**
- `outputs.paths`: JSON array de paths alterados

Exemplo de `solutions`:
```json
[
  {
    "path": "backend/java/spring/microservices/ms-example",
    "stack": "backend",
    "platform": "java",
    "framework": "spring",
    "type": "microservices",
    "name": "ms-example",
    "status": "active",
    "platformVersion": "21",
    "platformDistributor": "temurin",
    "docker": { "argCompilationMode": "jvm" }
  }
]
```

---

## Validação “hard” no trigger (WARNING + exclusão da matrix)

No trigger [.github/workflows/ci-cd-trigger-java.yml](../ci-cd-trigger-java.yml), o job `prepare_services`:

- Lê `needs.detect_solutions.outputs.solutions`
- Para cada solution:
  - Valida campos obrigatórios
  - Se faltarem campos ou estiverem `null`/vazio:
    - Loga `core.warning` com prefixo `[inventory-incomplete]`
    - **Não inclui** a solution na matrix de build
  - Filtra apenas `stack=backend` e `platform=java`

---

## Formato da matrix passada para o build reusable

O build reusable [.github/workflows/global-callables-java-build-solutions.yml](../global-callables-java-build-solutions.yml) recebe `inputs.services` (string JSON array), onde cada item deve conter:

Obrigatórios (para Java):
- `path`
- `goals`
- `platformVersion`
- `platformDistributor`
- `docker.argCompilationMode` (recomendado; dependendo da política, pode virar obrigatório)

Exemplo:
```json
[
  {
    "path": "backend/java/spring/microservices/ms-example",
    "goals": "clean install",
    "platformVersion": "21",
    "platformDistributor": "temurin",
    "docker": { "argCompilationMode": "jvm" }
  }
]
```

---

## `setup-java` e `platformDistributor`

No build reusable, o `actions/setup-java@v4` usa:
- `distribution: matrix.platformDistributor`
- `java-version: matrix.platformVersion`

Padronização adotada:
- `platformDistributor` deve estar no formato aceito pelo `setup-java` (ex.: `temurin`, `corretto`, `zulu`, `microsoft`).

---

## Variáveis Docker compartilhadas

O build reusable carrega:
- [.github/docker.env](../../docker.env)
- [.github/docker.java.env](../../docker.java.env)

Esses arquivos definem defaults como registry/org/tag prefix e nomes padrão de imagens builder/runtime.

---

## Pontos de atenção / troubleshooting

1. Matrix vazia:
   - Se todas as solutions detectadas estiverem incompletas no inventory, o trigger loga WARNING e não chama o build.
2. Inventory incompleto:
   - Procure por warnings `[inventory-incomplete]` no job `prepare_services` do trigger.
3. Erro de parse do `solutions`:
   - O trigger tenta parsear o JSON; se falhar, assume `[]` e o build não roda.

---

## Onde evoluir depois

- Tornar `docker.argCompilationMode` obrigatório (fail em vez de warning), se o processo exigir 100% de padronização.
- Adicionar uma checagem “lint de inventory” (job dedicado) para validar o JSON do inventory em PRs.