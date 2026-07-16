# Histórias de Usuário (User Stories) — Feature 02.2
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** M4 — Linha do Tempo e Auditoria (Entrega 2)
* **Feature Relacionada:** [02.2 — Trilha de Auditoria e Linha do Tempo](./04-FEATURES.md#feature-022-trilha-de-auditoria-e-linha-do-tempo) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Consulta à Linha do Tempo de Alterações

### 1. Descrição da História (Visão de Negócio)

Como Controller (perfil Auditor/Controller),
Quero consultar a linha do tempo completa de alterações realizadas nas tabelas fiscais, com filtros por período, usuário, entidade e tipo de operação,
Para rastrear rapidamente qualquer mudança e responder a questionamentos de auditoria interna ou externa sem depender do time técnico [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Filtros da Linha do Tempo): A consulta permite filtrar por: período (data inicial e final), usuário responsável, tipo de entidade (Alíquota, Classificação, Regime) e tipo de operação (Criação, Edição, Desativação).
* RN02 (Ordenação Cronológica): Os eventos são exibidos em ordem cronológica inversa (mais recente primeiro).
* RN03 (Acesso por Perfil): A Linha do Tempo é acessível a todos os perfis. Para o perfil Auditor/Controller, esta é a tela principal de trabalho — o sistema a exibe como dashboard padrão após o login.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Rastreamento de alteração específica

* Dado que o Controller precisa verificar quem alterou a alíquota de ICMS para SP na semana passada;
* Quando acessar a Linha do Tempo e filtrar por "Entidade = Alíquota", "Tributo = ICMS", "UF = SP" e período "últimos 7 dias";
* Então o sistema deve exibir todas as alterações que atendem aos filtros, cada uma com: data/hora, usuário, tipo de operação, entidade afetada e justificativa;
* E o Controller deve encontrar a alteração em menos de 2 minutos [INDEX].

#### Cenário 2: Auditor acessa a Linha do Tempo como tela inicial

* Dado que um usuário com perfil Auditor/Controller faz login no portal;
* Quando a sessão é iniciada;
* Então o sistema deve direcioná-lo para a tela da Linha do Tempo como dashboard padrão;
* E os botões de criação e edição não devem estar visíveis em nenhuma tela [INDEX].

------------------------------
## 📝 US-02: Comparação Visual entre Versões (Diff)

### 1. Descrição da História (Visão de Negócio)

Como Controller,
Quero visualizar lado a lado o estado anterior e o novo estado de uma entidade fiscal após uma alteração,
Para entender exatamente o que mudou sem precisar comparar manualmente dois registros [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Conteúdo do Diff): A comparação exibe todos os campos da entidade, destacando em vermelho os valores removidos e em verde os valores adicionados ou alterados.
* RN05 (Acesso via Linha do Tempo): A comparação é acessada clicando em qualquer evento na Linha do Tempo. O sistema carrega os snapshots do estado anterior e do novo estado registrados no momento da alteração.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Comparação de alteração de valor de alíquota

* Dado que a alíquota de CBS foi alterada de 0,9% para 9,0% em 15/07/2026 pelo Analista "João";
* Quando o Controller clica neste evento na Linha do Tempo;
* Então o sistema deve exibir a comparação lado a lado:
  * Lado esquerdo (ANTERIOR): CBS = 0,9%, destacado em vermelho
  * Lado direito (NOVO): CBS = 9,0%, destacado em verde
* E os campos que não foram alterados (ex: UF, Município) devem aparecer em cinza, idênticos nos dois lados [INDEX].

#### Cenário 2: Comparação de desativação de alíquota

* Dado que uma alíquota foi desativada;
* Quando o Controller clica no evento de desativação na Linha do Tempo;
* Então o lado esquerdo deve mostrar a alíquota com status "Vigente";
* E o lado direito deve mostrar a alíquota com status "Expirada" e a justificativa da desativação [INDEX].

------------------------------
## 📝 US-03: Garantia de Integridade e Imutabilidade da Trilha

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Controladoria e Compliance,
Quero ter a garantia de que os registros de auditoria são gerados automaticamente em 100% das alterações e são imutáveis — nenhum usuário, incluindo Administradores Fiscais, pode modificar ou excluir um registro de auditoria,
Para atestar a integridade da trilha em auditorias internas e externas e cumprir os requisitos da Lei das S.A. e do framework COSO [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN06 (Geração Automática): O registro de auditoria é gerado no exato momento em que a alteração é confirmada, sem nenhuma ação adicional do usuário.
* RN07 (Imutabilidade): Nenhum perfil de usuário possui permissão para modificar, excluir ou ocultar um registro de auditoria já gerado. Esta restrição é aplicada na camada de negócio do portal.
* RN08 (Registro de Tentativas de Violação): Qualquer tentativa de acessar ou modificar diretamente os registros de auditoria por meios não previstos deve ser registrada como evento de segurança e notificada ao Administrador Fiscal.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Verificação de geração automática

* Dado que um Analista Fiscal realiza qualquer operação de criação, edição ou desativação no portal;
* Quando a operação é confirmada;
* Então um registro de auditoria deve ser gerado automaticamente e aparecer imediatamente na Linha do Tempo;
* E o registro deve conter todos os campos obrigatórios: usuário, data/hora, entidade, valor anterior, valor novo, tipo de operação e justificativa [INDEX].

#### Cenário 2: Confirmação de imutabilidade para auditoria

* Dado que a Controladoria realiza a auditoria amostral mensal;
* Quando selecionar 50 alterações aleatórias do mês anterior e verificar cada uma na Linha do Tempo;
* Então 100% das alterações devem possuir registro de auditoria completo (KPI G1);
* E nenhum registro deve apresentar sinais de modificação posterior [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
