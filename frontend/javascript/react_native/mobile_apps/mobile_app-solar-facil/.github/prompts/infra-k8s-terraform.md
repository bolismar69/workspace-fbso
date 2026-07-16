# Infraestrutura K8s e Terraform

## Descrição

Guia para configurar e manter a infraestrutura do projeto usando Kubernetes e Terraform.

## Prompt

- Você está gerenciando a infraestrutura do Solar Facil com Kubernetes e Terraform.

### Contexto da infraestrutura:

- Cluster gerenciado (EKS/GKE/AKS)
- Deploy via Helm Charts e Ingress NGINX
- TLS via cert-manager + Let's Encrypt
- Banco de dados PostgreSQL provisionado
- Configuração via Terraform (módulos por ambiente)

### Práticas recomendadas:

- Evitar hardcoding (usar variáveis no Terraform)
- Separar ambientes (dev/staging/prod) com workspaces
- Armazenar estado do Terraform com backend remoto (S3, GCS)
- Usar secrets seguros (sealed-secrets, external-secrets)
- Monitorar com Prometheus/Grafana + alertas

## Tags

- infrastructure
- kubernetes
- terraform
- cloud
