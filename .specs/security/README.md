# O arquivo SECURITY.md
Em projetos grandes, as IAs têm uma tendência perigosa de priorizar a velocidade de entrega em detrimento da segurança (ex: esquecer de validar permissões de usuário em um novo endpoint ou expor dados sensíveis em logs). O SECURITY.md serve como a barreira de contenção de segurança da IA.

* A frequência: Também é um arquivo global, frequentemente atrelado às políticas da empresa ou do projeto.
* O que deve conter:
* Regras de Autenticação/Autorização: Ex: "Todo endpoint /admin precisa obrigatoriamente do middleware de validação de Role X".
   * Tratamento de Dados Sensíveis: Instruções claras para nunca expor senhas, cartões ou CPFs em logs do sistema.
   * Sanitização de Inputs: Padrões obrigatórios para evitar SQL Injection ou Cross-Site Scripting (XSS).
   * Política de Vulnerabilidades: Como reportar uma falha de segurança encontrada no repositório.

## O pulo do gato no SDD: Cruzando os dados
Embora SECURITY.md seja global, se uma feature específica (ex: Feature de Upload de Documentos) exigir uma regra de segurança única (ex: "Validar extensão e tamanho do arquivo para evitar malware"), essa regra específica deve constar no SPEC.md daquela feature, enquanto o SECURITY.md global cuida das regras gerais do sistema.
