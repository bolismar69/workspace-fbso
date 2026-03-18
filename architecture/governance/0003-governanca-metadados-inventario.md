### 📝 ADR 0003: Governança de Metadados e Inventário de Soluções

## 📅 Data

2026-03-12

## 💡 Status

Aceito

## 🎯 Contexto

Com o crescimento do monorepo e a convivência de múltiplas stacks (Java, Go, Python, .NET), a rastreabilidade e a padronização de deploy e conformidade tornaram-se complexas. Precisamos de uma "Fonte Única da Verdade" (Single Source of Truth) que descreva:

* Onde cada solução reside e qual seu blueprint.
* Quais são as dependências entre serviços, orquestradores e bancos de dados.
* Quais os requisitos de conformidade e segurança para cada módulo do **Tax-as-a-Service (TaaS)**.

## ✅ Decisão

Adotar o arquivo `architecture/governance/config/manager-solutions-inventory.json` como o **Manifesto de Governança Central**.

1. **Atributo Âncora (`systemNamespace`):** Toda solução deve pertencer a um `systemNamespace`. Este atributo define o isolamento lógico e físico (Namespace no Kubernetes, Projetos no ArgoCD e Prefixo no Kestra).
2. **Vínculos de Dependência (`dependsOn`):** É obrigatório declarar dependências de serviços, contratos (Protobuf/Avro) e recursos de infraestrutura para permitir builds seletivos e análise de impacto.
3. **Controle de Blueprint:** O campo `specification.docker.dockerfile` deve apontar obrigatoriamente para um blueprint validado em `architecture/blueprints/`.
4. **Conformidade (Compliance):** Definição estrita de thresholds de SonarQube e políticas de licenciamento diretamente no inventário.

### Regras de Ouro

* **Nenhum deploy sem registro:** O pipeline de CI/CD deve falhar se a solução que está sendo alterada não estiver cadastrada ou estiver inconsistente no inventário.
* **Imutabilidade de Namespace:** Uma vez definido o `systemNamespace`, ele não deve ser alterado sem um processo de migração de infraestrutura.

## 📐 Consequências

### Positivas

* **Visibilidade Total:** Facilidade em mapear o grafo de dependências da solução **TaxNexus**.
* **Automação Inteligente:** Possibilidade de criar scripts que autoconfiguram o ArgoCD e o Kestra baseados no JSON.
* **Segurança Nativa:** Imagens Distroless e compilação GraalVM garantidas via política de inventário.

### Negativas

* **Overhead de Manutenção:** Requer que o desenvolvedor atualize o JSON ao criar novos serviços ou mudar dependências.
* **Rigidez Inicial:** Mudanças na estrutura de pastas exigem atualização coordenada no inventário.

## 🔗 Referências

* Inventário: `architecture/governance/config/manager-solutions-inventory.json`
* Blueprints: `architecture/blueprints/`
* Estrutura Kestra: `ADR 0001`
