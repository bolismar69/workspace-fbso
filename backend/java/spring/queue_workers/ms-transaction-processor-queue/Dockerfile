# Stage 1: BUILDER - Usa o JDK para compilação
FROM eclipse-temurin:21-jdk-focal AS builder
WORKDIR /app
COPY pom.xml .
COPY src/ ./src/

# Comando de Build: Empacota o JAR executável
RUN mvn clean package -DskipTests

# Stage 2: RUNTIME - Usa apenas o JRE para execução (imagem menor)
FROM eclipse-temurin:21-jre-focal
WORKDIR /app

# Usuário de segurança
USER 1000

# Copia o JAR final (empacotado pelo plugin Spring Boot)
COPY --from=builder /app/target/*.jar app.jar

# Define o ponto de entrada (execução do JAR)
ENTRYPOINT ["java", "-jar", "app.jar"]