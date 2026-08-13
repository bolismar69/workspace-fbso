# Jira-Confluence Toolkit (jctk)

Toolkit reutilizável para integrar roadmaps Waterfall com Jira e Confluence via REST API.

## Estrutura

```
.specs/scripts/
├── jctk/                       # Módulos reutilizáveis
│   ├── config.py               # Auth, site URL, helpers
│   ├── jira/
│   │   ├── projects.py         # create_project, get_project
│   │   ├── issues.py           # create_epic, create_story, create_task, create_subtask
│   │   └── links.py            # link_issues, add_remote_links
│   └── confluence/
│       └── pages.py            # create_page, upload_markdown_file, add_jira_footer
├── projects/                   # Configurações por projeto
│   └── shield/                 # Projeto SHIELD (piloto)
│       ├── config.json         # Conexão e identificação
│       ├── structure.json      # Hierarquia Epics→Stories→Tasks
│       ├── doc-mapping.json    # Mapeamento docs→Confluence→Jira
│       └── create-shield.py    # Orquestrador
└── README.md
```

## Uso Rápido

```python
# Adicionar .specs/scripts ao PYTHONPATH
import sys
sys.path.insert(0, '/home/bolismar/work/workspace-fbso/.specs/scripts')

from jctk.jira.issues import create_epic, create_story, create_task
from jctk.jira.links import link_issues, add_remote_links
from jctk.confluence.pages import upload_markdown_file, create_page

# Criar Epic
epic = create_epic("PRJSHIELD", "F1 — Negócios & Discovery",
    description="📚 Docs: https://...",
    labels=["waterfall", "shield"])

# Criar Story filha do Epic
story = create_story("PRJSHIELD", "D1 — Infraestrutura",
    parent_epic_key=epic.get("key"),
    description="📚 Docs: https://...")

# Criar Task
task = create_task("PRJSHIELD", "[D1] 1.1.1 — DOKS + Istio",
    parent_key=epic.get("key"))  # Task filha do Epic

# Vincular Task ao Story
link_issues("PRJSHIELD-9", "PRJSHIELD-17", "Relates")

# Adicionar link do Confluence no Jira
add_remote_links("PRJSHIELD-9", [
    ("Deployment Plan v2.0", "https://bolismar69.atlassian.net/wiki/...")
])

# Subir documento para o Confluence
upload_markdown_file(
    "/path/to/doc.md", "NEPF", "01 — Project Charter", parent_id=4096001
)
```

## Criando um Novo Projeto

1. Copie `projects/shield/` para `projects/<seu-projeto>/`
2. Edite `config.json` com project_key, confluence_space, etc.
3. Edite `structure.json` com a hierarquia Epics→Stories→Tasks
4. Edite `doc-mapping.json` com o mapeamento docs→Confluence→Jira
5. Execute `python3 -m projects.<seu-projeto>.create-shield all`

## Pré-requisitos

- Token de API Atlassian salvo em `/tmp/jira-token.txt`
- Python 3.8+
- `curl` disponível no PATH
