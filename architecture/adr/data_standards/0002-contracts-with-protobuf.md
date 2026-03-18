## Proposta: ADR 0002 - Padronização de Contratos com Protobuf

### 📂 Caminho Sugerido

`architecture/adr/data_standards/0002-contracts-with-protobuf.md`

### 📝 Esboço do Conteúdo

**🎯 Contexto:** Com a stack poliglota (Go para cálculo e Java para persistência), o uso de JSON puro no Kafka pode gerar inconsistências de tipos e alto consumo de rede/processamento. Precisamos de um contrato "Typed" e independente de linguagem.

**✅ Decisão:** Adotar **Protocol Buffers (Protobuf)** como IDL (Interface Definition Language).

1. **Repositório Único:** Os contratos residirão em `architecture/data_standards/protobuf/`.
2. **Gerenciamento de Esquema:** Utilizar o **Confluent Schema Registry** (ou similar) no cluster para validar as mensagens em tempo real.
3. **Código Gerado:** O CI/CD gerará automaticamente os arquivos `.go` e `.java` a partir dos arquivos `.proto`.

**📐 Consequências:** * **Positiva:** Redução de 40% a 60% no tamanho das mensagens em comparação ao JSON.

* **Positiva:** Erros de contrato são pegos em tempo de compilação, não em runtime.

---

## Integração com sua Estrutura de Pastas

Para que isso ganhe vida dentro do seu monorepo, recomendo a criação desta subestrutura:

```text
architecture/
└── data_standards/
    └── protobuf/
        ├── tax/
        │   ├── calculation_request.proto  # Contrato p/ Motor em Go
        │   └── tax_result_event.proto     # Saída p/ Kafka (Lido pelo Java)
        └── common/
            └── person_identifiers.proto    # Tipos compartilhados (CNPJ/CPF)

```

### O papel do Kestra aqui

Com os contratos em Protobuf, o seu fluxo no **Kestra** (em `orchestration/kestra/flows/etl/`) poderá usar o plugin de Kafka para ler essas mensagens serializadas e transformá-las em carga para o **Oracle/PostgreSQL** com a garantia de que as colunas e tipos batem exatamente com o que o Motor de Cálculo gerou.

---
