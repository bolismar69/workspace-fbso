## ⚠️ O Grande Alerta: Views Atualizáveis (Updatable Views)
O PostgreSQL permite fazer INSERT, UPDATE e DELETE diretamente em uma View, mas com restrições. Como os serviços antigos continuarão rodando CRUDs contra as Views (com nomes antigos), elas precisam ser automaticamente atualizáveis.
Para que a View aceite escrita sem precisar de gatilhos (triggers) complexos, ela deve seguir estas regras:

* Ter apenas uma tabela na cláusula FROM.
* Não conter GROUP BY, HAVING, DISTINCT, LIMIT ou OFFSET.
* Não conter operações de conjunto (UNION, INTERSECT, EXCEPT).
* Não conter funções de agregação (SUM, COUNT, AVG, etc.).

Exemplo Prático da Transição:

   1. Você renomeia a tabela real para o nome correto e padronizado.
   2. Você cria a View com o nome antigo apontando para a tabela nova.

-- 1. Renomeia a tabela antiga para o padrão novoALTER TABLE imp_tx_venda RENAME TO imposto_taxa_venda;
-- 2. Cria a View com o nome antigo para o sistema velho não quebrarCREATE VIEW imp_tx_venda AS SELECT * FROM imposto_taxa_venda;

------------------------------
## 🛠️ 4 Checklists Obrigatórios para a Produção
Para garantir que o sistema antigo continue funcionando perfeitamente através da View, valide os seguintes pontos:
## 1. Resolução de Conflitos de Escopo (Namespaces)
No PostgreSQL, tabelas e views compartilham o mesmo espaço de nomes (namespace) dentro de um Schema. Você não pode ter uma tabela e uma view com o mesmo nome no mesmo schema.

* Ação: O comando de renomear a tabela e criar a view deve ser feito dentro da mesma transação (BEGIN; ... COMMIT;) para evitar quedas no sistema.

## 2. Permissões de Acesso (Grants)
As permissões da tabela antiga não são copiadas automaticamente para a View. Se o serviço antigo usa um usuário específico de banco de dados (ex: user_legacy), ele perderá o acesso.

* Ação: Você precisa reaplicar os comandos GRANT SELECT, INSERT, UPDATE, DELETE diretamente na nova View.

## 3. Valores Padrão (Default Values e Sequences)
Se o serviço antigo der um INSERT na View omitindo a chave primária (esperando o SERIAL ou DEFAULT gen_random_uuid()), a View do PostgreSQL irá repassar isso corretamente para a tabela base.

* Ação: Garanta que todas as restrições (NOT NULL, CHECK) e valores padrões permaneçam na tabela nova mapeados corretamente.

## 4. Triggers e Regras de Negócio
Se a tabela antiga possuía Triggers (gatilhos para auditoria ou cálculo de impostos), eles continuam na tabela nova.

* Ação: Se o serviço antigo disparava algum gatilho que dependia estritamente do comando TG_TABLE_NAME (retornando o nome da tabela antiga), esse comportamento vai mudar para o nome novo. Verifique o código das suas funções/triggers atuais.

------------------------------
## 🚀 Plano de Execução Seguro (Script Exemplo)
Execute a migração em blocos transacionais por tabela durante uma janela de manutenção (ou baixa utilização), pois o 

```sql
ALTER TABLE gera um bloqueio exclusivo temporário:

BEGIN;

  -- 1. Renomeia a tabela para o padrão correto
  ALTER TABLE tb_imposto_antiga RENAME TO configuracao_imposto;
  
  -- 2. Renomeia a Sequence associada (opcional, para manter a padronização)
  ALTER SEQUENCE IF EXISTS tb_imposto_antiga_id_seq RENAME TO configuracao_imposto_id_seq;

  -- 3. Cria a View de compatibilidade com o nome antigo
  CREATE OR REPLACE VIEW tb_imposto_antiga AS 
  SELECT * FROM configuracao_imposto;

  -- 4. Garante os acessos para os sistemas legados na View
  GRANT SELECT, INSERT, UPDATE, DELETE ON tb_imposto_antiga TO usuario_servico_antigo;
  
  -- 5. Garante os acessos para os novos serviços na Tabela Nova
  GRANT SELECT, INSERT, UPDATE, DELETE ON configuracao_imposto TO usuario_servico_novo;
COMMIT;
```

Para avançarmos com segurança, você saberia me dizer se alguma dessas tabelas possui Triggers complexos ou se os serviços antigos utilizam ORMs específicos (como Hibernate, Entity Framework, Prisma)? Posso ajudar a mapear os riscos para cada tecnologia.


---

As triggers exigem um cuidado redobrado nessa estratégia. No PostgreSQL, quando você renomeia uma tabela, as triggers associadas a ela são movidas automaticamente para a tabela nova e continuam funcionando. [1, 2] 
No entanto, o comportamento das Views em relação a essas triggers exige atenção. Como os serviços antigos agora farão operações de CRUD na View (com o nome antigo) e não mais na tabela diretamente, você precisa avaliar como essas triggers estão escritas.
Aqui está o diagnóstico do que vai acontecer e como mitigar os riscos:
## 1. O problema da variável TG_TABLE_NAME
Se o código da sua trigger (a função em PL/pgSQL) utiliza a variável interna TG_TABLE_NAME para salvar o nome da tabela em logs de auditoria ou tomar decisões lógicas, o comportamento vai mudar.

* Antes: TG_TABLE_NAME retornava nome_tabela_antigo.
* Depois: TG_TABLE_NAME passará a retornar nome_tabela_novo.
* Solução: Se o seu sistema de auditoria depende estritamente do nome antigo, você precisará atualizar o código da função da trigger para tratar o novo nome ou injetar o nome correto manualmente.

## 2. Triggers disparadas por atualizações em colunas (UPDATE OF)
Se os serviços antigos atualizarem a View, o PostgreSQL transformará internamente essa operação em um UPDATE na tabela base.

* Se a trigger antiga continha uma cláusula como AFTER UPDATE OF nome_coluna ON tabela, e você não mudou o nome das colunas, a trigger irá disparar perfeitamente na tabela nova.
* Atenção: Se além de renomear as tabelas você também planeja renomear as colunas para fins de padronização, a View precisará fazer um "de/para" de colunas (SELECT coluna_nova AS coluna_antiga). Nesse cenário, triggers do tipo BEFORE/AFTER UPDATE na tabela base podem não identificar corretamente quais colunas foram alteradas através da View.

## 3. Triggers de View vs. Triggers de Tabela
Para operações de escrita em Views, o PostgreSQL segue o seguinte fluxo:

   1. O serviço antigo envia um INSERT para a View.
   2. O PostgreSQL reescreve a consulta para a tabela base (graças às regras de Updatable Views).
   3. As triggers de BEFORE INSERT ou AFTER INSERT cadastradas na tabela base são disparadas normalmente. [3] 

Onde mora o perigo? Triggers do tipo BEFORE que alteram valores (ex: preencher um campo de data de modificação) vão funcionar. Mas se a trigger original precisava validar algo do contexto da requisição que a View omitiu, pode haver falhas. [4] 
------------------------------
## 🛡️ O Plano de Ação para as Triggers
Antes de rodar a migração em produção, execute este script de validação no seu ambiente de homologação para listar todas as triggers e suas respectivas funções:

```sql
SELECT 
    trg.tgname AS nome_da_trigger,
    rel.relname AS nome_da_tabela,
    proc.proname AS nome_da_funcao_trigger,
    prosrc AS codigo_fonte_da_funcaoFROM pg_trigger trgJOIN pg_class rel ON trg.tgrelid = rel.oidJOIN pg_proc proc ON trg.tgfoid = proc.oidJOIN pg_namespace nsp ON rel.relnamespace = nsp.oidWHERE nsp.nspname = 'public' -- Altere para o seu schema, se necessário
  AND trg.tgisinternal = false;
```

O que procurar no resultado desse script:

* Busque por textos que façam referência explícita ao nome antigo da tabela em formato de string (ex: 'tb_imposto').
* Procure pelo uso de TG_TABLE_NAME. [5] 

Se as funções das triggers forem apenas para cálculos matemáticos de impostos ou preenchimento de campos padrões (sem depender do nome da tabela em formato de texto), a sua estratégia de View continuará 100% segura e funcional.
Você possui alguma trigger de auditoria que grava o nome da tabela em uma tabela central de logs? Se preferir, pode colar o código de uma das triggers mais complexas aqui para analisarmos o impacto juntos.

[1] [https://www.devmedia.com.br](https://www.devmedia.com.br/trabalhando-com-triggers-no-postgresql/33531)
[2] [https://medium.com](https://medium.com/prolog-app/salvando-logs-de-altera%C3%A7%C3%B5es-em-tabelas-com-postgresql-599c0efedd85)
[3] [https://www.devmedia.com.br](https://www.devmedia.com.br/mysql-triggers/8088)
[4] [https://www.devmedia.com.br](https://www.devmedia.com.br/trabalhando-com-triggers-dml-no-oracle/32329)
[5] [https://www.devmedia.com.br](https://www.devmedia.com.br/trabalhando-com-triggers-no-postgresql/33531)
