# Kestra - Automation

Automações operacionais (rotinas, reconciliações, tarefas administrativas, jobs de manutenção).

## Estrutura sugerida por automação

- `<automacao>/flows/`
- `<automacao>/scripts/`

## Regras

- Preferir namespaces `fbso.automation.*`
- Logs/observabilidade e alertas devem ser configurados no runtime (devops), mas referenciados na doc do projeto
