# VISAO STACK TECNOLOGICA

- AUTH: 
- IdP: 

| CAMADA TECNOLOGICA          | SOLUÇÃO SISTEMICA BACKEND    | SOLUÇÃO SISTEMICA FRONTEND |
|-----------------------------|------------------------------|----------------------------|
| SISTEMA OPERACIONAL         | Linux | Linux
| TIPO DE STACK               | BACKEND | FRONTEND
| LINGUAGEM                   | Java(TM) SE Runtime Environment Oracle GraalVM 25.0.3+9.1 | React |
| COMPILADOR                  | GraalVM (essencial para gerar Native Images) | Compilador padrão
| FRAMEWORKS                  | Spring Boot 3.x | React 
| AUTENTICAÇÃO E AUTORIZAÇÃO  | JWT (JSON Web Tokens) com Spring Security | SAML 
| IDENTITY PROVIDER (IdP)     | Keycloak | Keycloak
| BANCO DE DADOS              | Postgresql 17 | - |
| MENSAGERIA                  | RabbitMQ | - |
| FRAMEWORK TESTES            | Junit versão 5, Simulação de Dependências usando Mockito e Testcontainers, Cobertura de Código: JaCoCo |  |
| INFRAESTRUTURA              | CLOUD  | - |
| Containerização             | Docker (para empacotar o executável nativo gerado pelo GraalVM em uma imagem ultra-leve, usando distros como Distroless ou Alpine). | - |
| Orquestração Cloud          | Kubernetes (essencial para gerenciar microsserviços em nuvem). | - |
| Provedor Cloud              | A DEFINIR | - |
| Repositório Central         | GITHUB <url do repositorio> | - |
| Pipeline CI/CD              | Pipeline CI/CD: GitHub Actions ou GitLab CI/CD (configurados para rodar os testes com JaCoCo, compilar a imagem nativa com GraalVM, buildar o container e realizar o deploy na nuvem) |  - |


Sua tabela está muito bem estruturada, mas ela apresenta duas inconsistências técnicas importantes e pode ser otimizada no formato para dar mais clareza ao time de arquitetura.
## 🔍 Correções Críticas na sua Tabela

   1. Autenticação (JWT vs SAML): O seu Frontend (React) e o seu Backend (Spring Boot) precisam falar o mesmo protocolo de segurança. Se o Backend usa JWT (OAuth2), o Frontend deve usar JWT/OAuth2 (via fluxo Authorization Code com PKCE). O SAML é um protocolo diferente e mais antigo (baseado em XML), inadequado para SPAs modernas como React integradas a microsserviços.
   2. Ambiente do Frontend (React): O React roda no navegador do usuário final (Chrome, Safari, Edge) sobre sistemas operacionais de usuário (Windows, macOS, Android, iOS), e não diretamente em um servidor Linux. O servidor Linux entra apenas para servir os arquivos estáticos (HTML/JS/CSS) via Nginx ou CloudFront.
   3. Compilador Frontend: O React não tem um "compilador padrão". Ele usa ferramentas de build e empacotamento, atualmente dominadas pelo Vite (que usa Esbuild/Rollup).

------------------------------
## 📊 Nova Proposta de Tabela (Abordagem por Ciclo de Vida)
Em vez de separar rigidamente por "Backend vs Frontend" em colunas paralelas (já que itens como Nuvem, CI/CD e Repositório são compartilhados), a estrutura abaixo organiza a stack por Ciclo de Desenvolvimento, Execução e Infraestrutura, resolvendo as pontas soltas.

| CAMADA TECNOLÓGICA | COMPONENTE BACKEND | COMPONENTE FRONTEND |
|---|---|---|
| Tipo de Stack | API Microsserviço | Single Page Application (SPA) |
| Linguagem | Java 25 (Oracle GraalVM 25.0.3+9.1) | TypeScript / JavaScript |
| Framework Principal | Spring Boot 3.x (ou 4.x) | React 19 |
| Ferramenta de Build/Compilação | GraalVM (Native Image Ahead-of-Time) | Vite (Esbuild & Rollup) |
| Autenticação & Autorização | JWT + Spring Security (OAuth2 Resource Server) | JWT + OAuth2 (Authorization Code + PKCE) |
| Identity Provider (IdP) | Keycloak | Keycloak |
| Banco de Dados | PostgreSQL 17 (via Spring Data JPA) | Não se aplica (Acessa via API) |
| Mensageria | RabbitMQ | Não se aplica |
| Framework de Testes & Qualidade | JUnit 5, Mockito, Testcontainers, JaCoCo | Vitest (ou Jest) e Testing Library |
| Ambiente de Execução (Runtime) | Linux (Container Docker Distroless/Alpine) | Navegador Web do Usuário (Client-side) |
| Hospedagem & Servidor Web | Kubernetes (Pod rodando o binário nativo) | Nginx (Container Linux) ou CDN (Ex: AWS CloudFront) |
| Orquestração & Cloud Provider | Kubernetes (Provedor A DEFINIR) | Kubernetes ou Storage Estático (A DEFINIR) |
| Repositório Central | GitHub [URL_AQUI] | GitHub [URL_AQUI] (Mesmo mono-repo ou separado) |
| Pipeline CI/CD | GitHub Actions / GitLab CI | GitHub Actions / GitLab CI |

## Por que esta nova estrutura é melhor?

* Alinhamento de Segurança: Garante que ambos usam Keycloak + JWT, evitando falhas de comunicação.
* Clareza de Build: Define o Vite para o React, que é o padrão atual de mercado por ser extremamente rápido.
* Separação de Runtime: Deixa claro que o React morre no servidor assim que é baixado pelo navegador do cliente, enquanto o Java roda 100% do tempo no cluster Linux/Kubernetes.

Para fecharmos essa definição de arquitetura, o repositório do projeto será um Monorepo (Backend e Frontend juntos no mesmo repositório GitHub) ou vocês trabalharão com repositórios separados?

