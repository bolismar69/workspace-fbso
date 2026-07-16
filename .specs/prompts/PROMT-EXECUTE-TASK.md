# Contexto:
  - Use as diretrize definidas nos arquivos informados
  - arquivo global @../../../../../.specs/security/SECURITY.md
  -  dos arquivos especficios dessa tarefa `ARCHITECTURE.md` , `PRD.md`
  - Atente-se que no arquivo `@README.md` vinculado a solução sistêmica foco da pasta atual, existem as instruções genéricas para execução e testes da aplicação.
  - Atente-se tambem que a documentação técnica da aplicação encontra-se na pasta `@.specs/` e ela deve ser usada como base de verdade atual da solução

# Missão: Implemente a FASE {solicitar qual fase de TASKS.md }
  - Para isso, siga estritamente as especificações contidas no arquivo `SPECS.md
  - e utilize o arquivo `TASKS.md` como o seu roteiro de execução passo a passo.

# Protocolo de Execução:
  1. Siga a Sequência: Execute uma tarefa de cada vez contida no checklist do @.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TASKS.md   . Não pule etapas e não tente implementar tarefas futuras antes de concluir as anteriores.
  2. Atualize o Progresso: À medida que você concluir com sucesso cada tarefa, marque a caixa de seleção correspondente mudando de [ ] para [x] diretamente no arquivo de tarefas.
  3. Desenvolvimento Orientado a Testes: Para cada trecho de código ou rota criada, consulte o arquivo @.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TEST_PLAN.md . Escreva e execute os testes unitários e de integração correspondentes imediatamente.
  4. Validação: Você só pode marcar uma tarefa como concluída no TASKS.md após rodar a suíte de testes localmente e garantir que ela passou com 100% de sucesso.
  5. Registro de Saída Imutável (TASK-EXECUTED): Assim que a Fase for concluída e validada por testes com 100% de sucesso, você deve OBRIGATORIAMENTE gerar um arquivo de documentação de saída na pasta correspondente da funcionalidade. O nome do arquivo deve seguir estritamente o padrão: `.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TASK-EXECUTED-AAAA-MM-DD-HHMMSS_[nome-da-feature-em-kebab-case].md` (Substitua AAAA-MM-DD-HHMMSS pelo carimbo de data e hora exato do momento da criação). O modelo para geração desse arquivo esta em @<pasta>

# Protocolo de Testes:
  * Se durante a execução do `TEST_PLAN.md` algum teste falhar, você deve seguir este fluxo obrigatoriamente:
  1. Auto-Correção Autônoma: Analise a saída do erro no terminal (stack trace), identifique se o problema está na lógica do código ou na estrutura do teste e faça as correções necessárias.
  2. Tratamento de Loops Infinitos: Se você tentar corrigir o mesmo erro por mais de 3 vezes seguidas sem sucesso, PARE a execução imediatamente.
  3. Registro de Impedimento (Fallback): Caso pare por erro repetido, crie um arquivo temporário chamado @.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/IMPEDIMENT-FASE-0.md   detalhando:
     - Qual teste quebrou e a mensagem de erro exata.
     - O que você tentou fazer para consertar.
     - Qual a sua suspeita do motivo de não estar funcionando (ex: limitação arquitetural, ambiguidade na SPEC,...).
     - Se existirem propostas adicionais de solução as adicione em uma sessão especifica.
  4. Geracao de arquivo com questionamentos (quando se aplicar): Caso se aplique crie arquivo com questionamentos ao Humano @.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/QUESTIONS.md  com perguntas que auxiliarão na retomada dos testes.
  5. Alerta ao Humano: Notifique o usuário no chat sobre o impedimento e aguarde instruções humanas antes de tentar alterar qualquer outro arquivo.

# Protocolo de Checagem Pós-Implementação (Sanity Check): 
Missão: Execute o protocolo de Checagem Pós-Implementação (Sanity Check) para as Fases 0, 1, 2.
Instruções passo a passo:
   1. Varredura de Limpeza: Procure e remova códigos comentados desnecessários, prints de debug (fmt.Println, console.log) e arquivos temporários criados durante os testes.
   2. Validação de Git Status: Execute um comando de status do Git e liste apenas os arquivos que foram modificados ou criados.
   3. Valide se a localização deles respeita o @.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/ARCHITECTURE.md e o @/.specs/architecture/architecture.md
   3. Travas do SECURITY.md: Faça uma última revisão nos arquivos alterados para garantir que nenhuma regra de segurança foi violada @../../../../../.specs/security/SECURITY.md
   4. Evidência de Sucesso: Confirme se o arquivo @.specs/business-projects/PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/TASKS.md está com todas as caixas da Fase 0 marcadas como [✅].
   5. Se tudo estiver 100% correto e limpo, prepare uma mensagem curta resumindo os arquivos alterados e informe que o projeto está pronto para o Commit/Pull Request.

# Protocolo para atualização da documentação do projeto na pasta @.specs/


# Comece agora lendo os arquivos mencionados e execute a primeira tarefa do checklist. Forneça o feedback de progresso a cada etapa concluída.
  