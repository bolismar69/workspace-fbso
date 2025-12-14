
## 1\. 🚀 Detalhamento da Especialização - Quarkus

As especializações herdam o `checkstyle.xml` e `.editorconfig`, mas adicionam a lógica de *build* e *runtime* específica do framework.

### Quarkus (`java/java-quarkus/`) - Foco em Microservices

  * **`pom.xml`:** Inclui as dependências Quarkus (RESTEasy, Panache) e o plugin Quarkus para *build*.
  * **`Dockerfile.jvm`:** Focado em *fast startup* (JVM padrão).
  * **`Dockerfile.native`:** Focado na menor imagem possível usando GraalVM.

### B. Spring Boot (`java/java-spring/`) - Foco em Batch

  * **`pom.xml`:** Inclui Spring Boot Starters (Batch, Data JPA) e o plugin Spring Boot para *repackage*.
  * **`application.yml`:** Contém perfis de configuração padronizados (`dev`, `staging`, `prod`) com portas padrão e logs.

### 🛠️ Ação: Criar Arquivos de Especialização

Para manter o foco, vamos criar o `pom.xml` e o `Dockerfile.native` de Quarkus, que são os mais distintos:

```bash
# Criar o POM de Quarkus (Simplificado)
cat <<EOF > architecture/blueprints/java/java-quarkus/pom.xml
<project>
  <modelVersion>4.0.0</modelVersion>
  <parent><groupId>com.fbso</groupId><artifactId>java-generic</artifactId><version>1.0.0</version></parent>
  
  <dependencies>
    <dependency><groupId>io.quarkus</groupId><artifactId>quarkus-resteasy-reactive</artifactId></dependency>
    </dependencies>
  
  <build>
    <plugins>
      <plugin><groupId>io.quarkus</groupId><artifactId>quarkus-maven-plugin</artifactId></plugin>
    </plugins>
  </build>
</project>
EOF

# Criar o Dockerfile Native (Imagem Final para Microservices)
cat <<EOF > architecture/blueprints/java/java-quarkus/Dockerfile.native
# Stage 1: Build da imagem nativa (compilação)
FROM quay.io/quarkus/ubi-quarkus-mandrel-builder:jdk-17 AS build
COPY --chown=quarkus:quarkus src /home/quarkus/src
WORKDIR /home/quarkus/src
# Execução da compilação nativa
RUN mvn package -Pnative

# Stage 2: Imagem final de runtime (Menor possível)
FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /app
COPY --from=build /home/quarkus/src/target/*-runner /app/application
# Ponto de entrada nativo
CMD ["./application", "-Dquarkus.http.port=8080"]
EOF
```
