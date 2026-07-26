Quando um incidente acontece em produção, o ecossistema de Spec-Driven Development (SDD) deve operar em modo de Engenharia Reversa e Contenção.
Você não deve criar uma nova estrutura de arquivos do zero. O grande segredo do SDD é que a correção de um bug de produção deve atualizar e reutilizar a mesma estrutura que você já possui (PRD, ARCHITECT, TASKS, TEST_PLAN), garantindo que a "fonte única da verdade" permaneça viva e protegida contra novos erros (regressões).
Abaixo, veja como estruturar o processo via agentes e como os arquivos Markdown existentes são impactados.

------------------------------
## 🕵️‍♂️ O Processo de Resolução via Agentes (Fluxo de 3 Passos)
> Em vez de deixar o desenvolvedor tentar adivinhar o erro, o processo de produção usa um agente especializado em diagnóstico (Triage/Forensics Agent).

## Passo 1: O Agente de Triagem e Causa Raiz (Root Cause Agent)
Esse agente recebe o log do erro de produção ou o report do usuário e varre o repositório existente para descobrir onde a especificação ou o código falharam.

### Prompt do Agente de Triagem:

```markdown
# PERSONA
Você é um Engenheiro de Confiabilidade de Sites (SRE) Sênior especializado em diagnóstico forense de sistemas.

# ENTRADAS DO PROCESSO
- Logs / Telemetria do Erro em Produção: {{LOGS_PRODUCAO}}
- Contexto Atual do Projeto: `PRD.md`, `ARCHITECT.md` e `TEST_PLAN.md` atuais.

# INSTRUÇÕES
1. Analise o log de erro e identifique a causa raiz técnica.
2. Identifique qual arquivo de código e linhas específicas estão gerando o incidente.
3. Compare o comportamento do bug com o `TEST_PLAN.md` atual. Explique por que este cenário não foi previsto ou por que o teste atual deixou o bug passar (Blindspot/Ponto Cego).
4. Explique se a correção violará alguma regra do `ARCHITECT.md` atual.

# FORMATO DE SAÍDA (INCIDENT_DIAGNOSIS.md)
## 🚨 Diagnóstico do Incidente em Produção
- **Causa Raiz:** [Explicação técnica clara]
- **Arquivos Afetados:** `caminho/do/arquivo.ext` (Linhas X-Y)
- **Ponto Cego Identificado:** [Por que o TEST_PLAN.md original falhou em pegar isso?]
- **Impacto Arquitetural:** [A correção exige mudar o ARCHITECT.md ou apenas o código?]
```

------------------------------
## 🔄 Como a Estrutura de Arquivos MD é Impactada?
Com o INCIDENT_DIAGNOSIS.md gerado, você não cria novos arquivos, mas sim executa uma atualização cirúrgica nos arquivos existentes através de um agente de atualização de contexto:
## 1. PRD.md (Geralmente não muda)

* Regra: Um bug de produção geralmente é um desvio técnica, não uma nova regra de negócio. O PRD.md permanece intacto, a menos que o incidente revele que o negócio operava sob uma premissa errada.

## 2. ARCHITECT.md (Raras alterações)

* Regra: Só muda se o incidente foi causado por uma falha estrutural (ex: vazamento de memória por má escolha de biblioteca ou estouro de conexões com o banco). Se mudar, a nova regra de infraestrutura/design é adicionada aqui para a IA nunca mais repetir o erro.

## 3. TEST_PLAN.md (⚠️ Alteração Obrigatória - O Passo Mais Importante)

* Regra: Antes de corrigir o código, o agente de QA deve ler o INCIDENT_DIAGNOSIS.md e adicionar o cenário do bug como um novo caso de teste no TEST_PLAN.md.
* Por que isso é vital? Isso se chama Regression Testing (Teste de Regressão). Ao escrever o teste que quebra com o bug atual antes de corrigi-lo, você garante que esse erro específico nunca mais voltará a acontecer em deploys futuros.

## 4. TASKS.md (Nova Entrada)

* Regra: O orquestrador adiciona uma tarefa específica ao final do arquivo (ou cria um HOTFIX_TASK.md temporário), cujo escopo é exclusivamente: "Implementar a correção descrita em INCIDENT_DIAGNOSIS.md e fazer o código passar no novo teste adicionado ao TEST_PLAN.md".

------------------------------
## 🚀 O Loop de Execução do Hotfix
Uma vez que o TEST_PLAN.md foi atualizado com o novo caso de teste que simula o bug e a tarefa foi inserida, o ecossistema roda exatamente o fluxo que você já desenhou anteriormente:

   1. O Agente Desenvolvedor lê a tarefa de hotfix e o código.
   2. Ele altera o código.
   3. O Agente de QA e o Agente de Segurança rodam as checagens (agora incluindo o novo teste do bug de produção).
   4. Passando no loop (com o limite de 3 tentativas), o código vai para produção via Pull Request de Hotfix.

Dessa forma, o seu ciclo se fecha de maneira perfeita: o incidente de produção gerou conhecimento que foi documentado diretamente no seu TEST_PLAN.md, tornando o seu software e suas especificações cada vez mais resilientes com o passar do tempo.
------------------------------
Para consolidar esse fluxo de tratamento de incidentes, o que gostaria de detalhar a seguir?

* O prompt do Agente de QA para atualizar o TEST_PLAN.md adicionando o caso de teste que simula o bug?
* O fluxo de como o orquestrador deve priorizar essa tarefa de hotfix na fila de desenvolvimento?


