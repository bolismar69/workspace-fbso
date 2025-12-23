# Stage 1: BUILDER - Usa o JDK para compilação
FROM maven:3.9.9-eclipse-temurin-21 AS builder
WORKDIR /app
COPY pom.xml .
COPY src/ ./src/

# Comando de Build: Empacota o JAR executável
RUN mvn clean package -DskipTests

# Stage 2: RUNTIME - Usa apenas o JRE para execução (imagem menor)
FROM eclipse-temurin:21-jre-jammy
WORKDIR /app

# Usuário de segurança
USER 1000

# Copia o JAR final (empacotado pelo plugin Spring Boot)
COPY --from=builder /app/target/*.jar app.jar

# Define o ponto de entrada (execução do JAR)
ENTRYPOINT ["java", "-jar", "app.jar"]
