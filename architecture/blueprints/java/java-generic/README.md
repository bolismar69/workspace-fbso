
## 1\. 🧬 Detalhamento do Blueprint Genérico (`java/java-generic/`)

O Blueprint Genérico define a qualidade de código, as regras de construção e a estrutura de alto nível, independentemente do framework.

| Arquivo | Propósito | Exemplo de Conteúdo (Simulado) |
| :--- | :--- | :--- |
| **`pom.xml`** (Base) | Define dependências comuns (JUnit, Lombok, Logs) e plugins de qualidade (Checkstyle, JaCoCo). | Versão padrão do JDK (17+), configurações para compilação. |
| **`checkstyle.xml`** | Define as regras de estilo e complexidade de código que **todos** os projetos Java devem seguir. | Limites de complexidade ciclomática, regras de *naming* (camelCase), cabeçalho de arquivos. |
| **`.editorconfig`** | Garante que todos os IDEs usem a mesma tabulação (spaces vs. tabs) e codificação (UTF-8). | `indent_style = space`, `indent_size = 4`. |
| **`Dockerfile.base`** | Define a imagem base (`Eclipse Temurin` ou `OpenJDK`) e a camada de segurança para a JVM. | Imagem base padronizada e usuário não-root. |

### 🛠️ Ação: Criar Arquivos Base

```bash
cd architecture/blueprints/java/java-generic

# Criar o arquivo base de configuração de IDEs
echo -e "[*]\nindent_style = space\nindent_size = 4\ncharset = utf-8" > .editorconfig

# Criar um POM base (simplificado)
cat <<EOF > pom-base.xml
<?xml version="1.0" encoding="UTF-8"?>
<project>
  <modelVersion>4.0.0</modelVersion>
  <properties>
    <java.version>17</java.version>
    <maven.compiler.source>\${java.version}</maven.compiler.source>
    <maven.compiler.target>\${java.version}</maven.compiler.target>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>
  <dependencies>
    <dependency><groupId>org.junit.jupiter</groupId><artifactId>junit-jupiter-engine</artifactId><version>5.9.2</version><scope>test</scope></dependency>
    </dependencies>
</project>
EOF

# Criar um Dockerfile base
cat <<EOF > Dockerfile.base
# Imagem Base (Java Runtime)
FROM eclipse-temurin:17-jre-focal

# Configuração de segurança
USER 1000

WORKDIR /app

# Argumentos de execução JVM padrão (segurança, memória)
ENTRYPOINT ["java", "-XX:InitialRAMPercentage=75.0", "-jar", "app.jar"]
EOF

cd ../../../.. # Voltar para a raiz do monorepo
```
