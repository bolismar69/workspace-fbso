# O arquivo DEPLOY.md (ou CICD.md)
A IA precisa saber exatamente como o código sai da máquina local e chega ao ambiente de produção. Se você não documentar isso, o agente pode sugerir ferramentas de deploy erradas ou quebrar o fluxo de Integração Contínua (CI/CD).

* A frequência: Ele é um arquivo global/estático por repositório. Você só o atualiza se mudar a infraestrutura ou o provedor de nuvem.
* O que deve conter:
* Variáveis de Ambiente: Lista de chaves necessárias (sem os valores secretos) para o sistema rodar na nuvem.
   * Comandos de Build: Ex: npm run build, docker build.
   * Fluxo de Branches: Como funciona o Git Flow (ex: commits na main disparam deploy para produção, commits na develop vão para staging).
   * Migrações de Banco: Instruções de como rodar as migrations em produção de forma segura. [1, 2] 

Importante para o SDD: Quando a IA cria uma nova rota ou funcionalidade, ela lê o DEPLOY.md para verificar se precisa atualizar o arquivo de configuração do GitHub Actions, Dockerfile ou adicionar novas variáveis de ambiente no painel da nuvem.

## O pulo do gato no SDD: Cruzando os dados
Embora DEPLOY.md sejam global, se uma feature específica (ex: Feature de Upload de Documentos) exigir uma regra de deploy única (ex: "Executar validação de existência de arquivo especifoc"), essa regra específica deve constar no SPEC.md daquela feature, enquanto o SECDEPLOYURITY.md global cuida das regras gerais do deploy.
