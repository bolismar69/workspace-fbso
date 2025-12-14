Perfeito\! Vamos finalizar o trabalho na stack Java detalhando os arquivos essenciais para o **Blueprint Spring Boot**, com foco em serviços de **Batch** e integrações, complementando assim o *Blueprint* Quarkus focado em Microsserviços.

O *Blueprint* Spring deve herdar as regras de qualidade do `java-generic` e adicionar as convenções do ecossistema Spring.

## 1\. ⚙️ Detalhamento do Blueprint Spring (`java/java-spring/`)

| Arquivo | Propósito | Foco e Padronização |
| :--- | :--- | :--- |
| **`pom.xml`** | Define as dependências do Spring Boot e do Spring Batch. | Inclui `spring-boot-starter`, `spring-boot-starter-batch`, e `spring-boot-starter-data-jpa` (para persistência do *batch*). |
| **`application.yml`** | Define configurações de perfil e convenções de logs. | Padrões de porta, nome do serviço, níveis de log (INFO para a maioria dos pacotes, DEBUG para pacotes críticos). |
| **`Dockerfile`** | Otimizado para o Spring Boot (utilizando o plugin de *repackage*). | Usa a imagem base do Java e o comando `java -jar` para executar o JAR final. |
| **`SpringBatchConfig.java`** (Exemplo) | Modelo de configuração para o *Job* e *Step* do Batch. | Garante que todos os serviços de *batch* sigam o mesmo padrão de transação e tratamento de erros. |

### 🛠️ Ação: Criar Arquivos de Especialização Spring

Vamos criar os arquivos `pom.xml`, `application.yml` e `Dockerfile` dentro da pasta de *Blueprint* Spring.

```bash
# Navegar para a pasta do Blueprint Spring
cd architecture/blueprints/java/java-spring

# 1. Criar o POM de Spring Boot (Foco em Batch)
cat <<EOF > pom.xml
<?xml version="1.0" encoding="UTF-8"?>
<project>
  <modelVersion>4.0.0</modelVersion>
  <parent>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-parent</artifactId>
    <version>3.2.0</version>
    <relativePath/>
  </parent>
  
  <groupId>com.fbso.batch</groupId>
  <artifactId>spring-batch-template</artifactId>
  <version>1.0.0</version>

  <properties>
    <java.version>17</java.version>
  </properties>

  <dependencies>
    <dependency><groupId>org.springframework.boot</groupId><artifactId>spring-boot-starter-batch</artifactId></dependency>
    <dependency><groupId>org.springframework.boot</groupId><artifactId>spring-boot-starter-data-jpa</artifactId></dependency>
    <dependency><groupId>org.springframework.boot</groupId><artifactId>spring-boot-starter-web</artifactId></dependency>
    <dependency><groupId>org.springframework.boot</groupId><artifactId>spring-boot-starter-test</artifactId><scope>test</scope></dependency>
  </dependencies>
  
  <build>
    <plugins>
      <plugin><groupId>org.springframework.boot</groupId><artifactId>spring-boot-maven-plugin</artifactId></plugin>
    </plugins>
  </build>
</project>
EOF

# 2. Criar o arquivo de Configuração Padrão (YAML)
cat <<EOF > application.yml
spring:
  # Configuração padrão de Batch
  batch:
    job:
      enabled: false # Jobs devem ser iniciados via linha de comando ou API
  datasource:
    # Configuração Padrão de Data Source (deve ser sobrescrito pelo ambiente)
    url: jdbc:postgresql://localhost:5432/batch_db
    username: user
    password: password
    driver-class-name: org.postgresql.Driver
  profiles:
    active: dev # Perfil padrão
    
logging:
  level:
    root: INFO
    org.springframework: INFO
    com.fbso: DEBUG # Nível de debug para nosso código de negócio
server:
  port: 8080 # Porta padrão para Health Checks
EOF

# 3. Criar o Dockerfile otimizado para Spring JAR
cat <<EOF > Dockerfile
# Stage 1: Build - usa uma imagem completa para compilação
FROM eclipse-temurin:17-jdk-focal AS builder
WORKDIR /app
COPY pom.xml .
COPY src/ ./src/
RUN mvn clean package -DskipTests

# Stage 2: Runtime - usa apenas a JRE para reduzir o tamanho da imagem
FROM eclipse-temurin:17-jre-focal
WORKDIR /app
# Copia o JAR final (empacotado pelo plugin Spring Boot)
COPY --from=builder /app/target/*.jar app.jar

# Configuração de segurança
USER 1000

# Execução do JAR
ENTRYPOINT ["java", "-jar", "app.jar"]
EOF

cd ../../../.. # Voltar para a raiz do monorepo
```

### 2\. 📤 Commit e Finalização

Com o `java-spring` agora detalhado, a governança da stack Java está completa.

```bash
git add .
git commit -m "feat(java): Implement spring blueprint with batch focus (pom, yml, dockerfile)"
git push origin main
```

**Próximo Passo:**
O trabalho de governança exige a documentação das decisões. Devemos criar o primeiro **ADR (Architecture Decision Record)** em `architecture/adr/` explicando a divisão Java (Quarkus vs. Spring). Gostaria de prosseguir com a criação do ADR?