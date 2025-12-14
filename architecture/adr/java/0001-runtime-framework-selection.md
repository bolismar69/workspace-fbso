# ADR 0001: Seleção de Frameworks Java para Runtime

## 📅 Data
[Data de Hoje: 2025-12-11]

## 💡 Status
Aceito

## 🎯 Contexto

O time de engenharia FBSO utiliza a linguagem Java para a maior parte da lógica de negócio. Para garantir a melhor performance, eficiência de recursos (custo de nuvem) e experiência do desenvolvedor, é necessário padronizar qual framework deve ser utilizado em diferentes contextos de execução (online vs. offline/batch).

Adotar um único framework para todos os casos (ex: Spring Boot para tudo) resultaria em alto consumo de memória e tempo de inicialização lento para microsserviços Cloud Native. Por outro lado, forçar o uso de um framework leve para tarefas complexas de integração pode aumentar a complexidade de desenvolvimento.

## ✅ Decisão

A organização adota uma estratégia de dois frameworks Java (Dual-Framework Strategy), selecionados com base no **Tipo de Execução (Runtime Context)**:

1.  **Quarkus:** Framework primário para serviços que exigem inicialização rápida e baixa pegada de memória (Memory Footprint).
2.  **Spring Boot:** Framework secundário para serviços que requerem integração complexa, ecossistema estabelecido e processamento em lote (Batch).

### ➡️ Regra de Seleção:

| Cenário de Uso | Framework Escolhido | Motivação Principal | Blueprint Recomendado |
| :--- | :--- | :--- | :--- |
| **Microsserviços/APIs HTTP/Serviços de Fila** | **Quarkus** | Cloud Native, compilação nativa (GraalVM), boot time < 1s, baixo custo de memória. |  |
| **Serviços de Processamento em Lote (Batch)** | **Spring Boot** | Ecossistema maduro para , transações complexas, alta capacidade de integração. |  |
| **Bibliotecas/Módulos de Reuso (Libs)** | **Nativo (Sem Framework)** | Padrão Clean Architecture (DNA Genérico). |  |

## 📐 Consequências

### Positivas:

* **Otimização de Custos:** Microsserviços Quarkus reduzem drasticamente o custo de memória na nuvem.
* **Performance:** Atingimos o melhor tempo de resposta com serviços leves de Quarkus para o runtime online.
* **Produtividade:** O time de Batch se beneficia da familiaridade e do vasto ecossistema do Spring Batch.

### Negativas:

* **Curva de Aprendizado:** Desenvolvedores precisam ter familiaridade com dois ecossistemas (Quarkus e Spring).
* **Governança Reforçada:** Requer manutenção contínua e atualização de dois conjuntos de *blueprints* e dependências.

## 🔗 Referências

* [Blueprint Quarkus](./../../blueprints/java/java-quarkus/README.md)
* [Blueprint Spring](./../../blueprints/java/java-spring/README.md)
* [Especificações de Dockerfile Native](./../../blueprints/java/java-quarkus/Dockerfile.native.v17)
