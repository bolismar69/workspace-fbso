# Documento de Definição de Requisitos (02-BUSINESS-REQUIREMENTS.md)
* Programa: Adequação Corporativa à Reforma Tributária Nacional
* Status: Em Definição
* Responsáveis: Product Managers (PMs) e Product Owners (POs) Corporativos
* Público-Alvo: Comitê Fiscal, Lideranças de Negócios e Arquitetura de Sistemas

------------------------------

## 1. Objetivos do Documento

Este documento especifica os Requisitos de Negócio (Business Requirements) necessários para cumprir o objetivo estratégico do Project Charter: unificar a inteligência tributária do IVA Dual (CBS, IBS e Imposto Seletivo), garantindo governança de preços, proteção de margens no Lucro Real e conformidade de faturamento em todo o território nacional.

------------------------------

## 2. Princípios Macroscópicos de Negócio (Guardrails)

* Omnicanalidade Tributária: O canal de venda (seja e-commerce, portal B2B ou venda direta) não dita a regra do imposto; a regra é corporativa e agnóstica ao ponto de contato.
* Princípio do Destino Inviolável: Toda simulação ou faturamento deve considerar as regras vigentes do local exato de consumo/entrega do bem ou serviço.
* Convivência de Modelos: O ecossistema deve ser capaz de gerenciar a transição progressiva das alíquotas (modelo antigo decrescendo e modelo novo crescendo) de forma transparente.

------------------------------

## 3. Requisitos de Negócio (Business Requirements - BR)

As necessidades foram divididas de acordo com as ondas de entrega de valor definidas no planejamento estratégico.

## 📦 Bloco 1: Gestão de Dados, Alíquotas e Governança (Fundação de Negócio)

* BR-01: Centralização da Inteligência de Regras
   * Descrição: A companhia deve possuir uma única fonte da verdade para o cálculo de impostos, eliminando mecanismos de precificação locais ou isolados.
      * Justificativa de Negócio: Evitar desalinhamento de preços entre canais e garantir auditoria centralizada.
* BR-02: Autonomia do Time Fiscal (No-Code/Dynamic Update)
   * Descrição: O time fiscal/controladoria deve ser capaz de atualizar a matriz de alíquotas nacionais (gerais, reduzidas, isenções de IBS/CBS) sem depender de intervenção técnica ou desenvolvimento de software para cada alteração de lei.
      * Justificativa de Negócio: Agilidade para responder às flutuações e decretos do Comitê Gestor do IBS.
* BR-03: Qualificação Geográfica de Cadastro
   * Descrição: O processo corporativo deve exigir e validar a precisão da localização geográfica do tomador/comprador (código IBGE do destino) no início da transação comercial.
      * Justificativa de Negócio: Mitigar o risco de aplicar alíquotas incorretas geradas por cadastros desatualizados ou incompletos.

## 🛒 Bloco 2: Jornada Comercial e Experiência do Cliente (Onda 1 - Vendas)

* BR-04: Transparência e Cálculo "Por Fora"
   * Descrição: Os sistemas de vendas devem apresentar ao cliente o preço base do produto/serviço acrescido do impacto exato do IVA Dual calculado por fora da mercadoria.
      * Justificativa de Negócio: Atendimento à legislação de transparência fiscal e proteção da percepção de valor do cliente.
* BR-05: Proteção de Margem e Simulação Comercial
   * Descrição: A força de vendas deve ter visibilidade do preço líquido e da margem real antes de fechar contratos de longo prazo, considerando a variação do IBS do destino.
      * Justificativa de Negócio: Evitar que vendas interestaduais corroam a margem de lucro projetada.
* BR-06: Garantia de Preço Ofertado (Token de Validade)
   * Descrição: O valor tributário simulado na proposta comercial deve ser garantido por uma janela de tempo específica para que não haja flutuação de preço entre a intenção de compra e o faturamento.
      * Justificativa de Negócio: Evitar quebras de expectativa com o cliente e atritos comerciais.

## 🧾 Bloco 3: Operação Financeira e Faturamento (Onda 2 - Finanças e ERP)

* BR-07: Unicidade Matemática entre Pedido e Nota Fiscal
   * Descrição: O valor consolidado no momento da emissão do faturamento deve ser idêntico ao simulado em vendas, atingindo 100% de consistência.
      * Justificativa de Negócio: Garantir que nenhuma nota seja rejeitada pelos órgãos fiscalizadores e eliminar divergências de conciliação.
* BR-08: Rastreabilidade de Créditos no Lucro Real
   * Descrição: O processo de faturamento e compras deve mapear o potencial de apropriação de créditos de CBS/IBS de cada operação, separando o imposto recuperável do custo do produto.
      * Justificativa de Negócio: Maximizar a recuperação de créditos, essencial para a saúde financeira no regime de Lucro Real.
* BR-09: Viabilização do Mecanismo de Split Payment
   * Descrição: O processo de liquidação financeira deve discriminar a partição exata do valor da venda que pertence à receita da empresa e os percentuais que serão destinados imediatamente ao CBS e ao IBS.
      * Justificativa de Negócio: Preparar a companhia para as novas regras bancárias de recolhimento na fonte estabelecidas pela reforma.

------------------------------

## 4. Matriz de Rastreabilidade Comercial
Este mapeamento garante que nenhuma necessidade de negócio fique descoberta na fase de execução técnica:

| ID Requisito | Meta Relacionada no Charter | Canal Impactado | Área de Negócio Madrinha |
|---|---|---|---|
| BR-01 | Unificação da Inteligência | Todos | Controladoria / TI |
| BR-02 | Governança Ágil | Nenhum (Backoffice) | Comitê Fiscal |
| BR-03 | Precisão de Destino | Cadastro / CRM | Operações / Comercial |
| BR-04 | Transparência Fiscal | E-commerce / Canais | Marketing / Comercial |
| BR-05 | Preservação de Margem | Venda Direta / B2B | Controladoria / Comercial |
| BR-06 | Consistência Comercial | Todos os Canais | Comercial |
| BR-07 | Zero Rejeição Fiscal | Faturamento / ERP | Fiscal / Faturamento |
| BR-08 | Otimização de Crédito | Suprimentos / Compras | Controladoria |
| BR-09 | Fluxo de Caixa / Split | Tesouraria / Bancos | Financeiro |

------------------------------

## 5. Critérios de Homologação de Negócio (UAT)
A aceitação final deste bloco de requisitos pelos Business Owners ocorrerá quando:

1. O time comercial conseguir simular um pedido em Santana de Parnaíba com entrega em qualquer capital do país e visualizar o impacto correto do IBS na margem.
2. O time financeiro comprovar que os lançamentos contábeis de teste refletem as contas corretas de impostos a recuperar/recolher do Lucro Real.
3. A auditoria fiscal interna validar que a regra de transição de alíquotas preserva o caixa corporativo.

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: brd-creation, business-rules-analysis.*

