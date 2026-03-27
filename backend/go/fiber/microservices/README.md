# 🗃️ Workspace Central FBSO - definitions

## IMPOSTOS A SEREM FOCO DE CONSTRUCAO DE MICROSERVICOS

### 1. Impostos para Pessoa Jurídica (Empresas)

As empresas possuem a maior carga de complexidade, variando conforme o regime tributário (Simples Nacional, Lucro Presumido ou Lucro Real).

| Esfera | Imposto / Contribuição | Descrição Curta |
| :--- | :--- | :--- |
| **Federal** | **IRPJ** | Imposto de Renda Pessoa Jurídica. |
| **Federal** | **CSLL** | Contribuição Social sobre o Lucro Líquido. |
| **Federal** | **PIS/Pasep** | Programa de Integração Social. |
| **Federal** | **COFINS** | Contribuição para o Financiamento da Seguridade Social. |
| **Federal** | **IPI** | Imposto sobre Produtos Industrializados. |
| **Federal** | **CPP** | Contribuição Patronal Previdenciária (INSS Patronal). |
| **Federal** | **FGTS** | Fundo de Garantia do Tempo de Serviço (encargo sobre folha). |
| **Estadual** | **ICMS** | Imposto sobre Circulação de Mercadorias e Serviços. |
| **Municipal** | **ISS (ou ISSQN)** | Imposto Sobre Serviços de Qualquer Natureza. |

---

### 2. Impostos para Pessoa Física (Indivíduos)

Focado em renda, propriedade e consumo direto.

| Esfera | Imposto / Contribuição | Descrição Curta |
| :--- | :--- | :--- |
| **Federal** | **IRPF** | Imposto de Renda Pessoa Física (Baseado na tabela progressiva). |
| **Federal** | **INSS** | Contribuição Previdenciária (desconto em folha ou autônomo). |
| **Federal** | **IOF** | Imposto sobre Operações Financeiras (cartão, câmbio, empréstimos). |
| **Federal** | **ITR** | Imposto sobre a Propriedade Territorial Rural. |
| **Estadual** | **IPVA** | Imposto sobre a Propriedade de Veículos Automotores. |
| **Estadual** | **ITCMD** | Imposto sobre Transmissão Causa Mortis e Doação (Heranças). |
| **Municipal** | **IPTU** | Imposto sobre a Propriedade Predial e Territorial Urbana. |
| **Municipal** | **ITBI** | Imposto sobre Transmissão de Bens Imóveis (Compra e Venda). |

---

### 3. O Cenário da Reforma Tributária (Transição 2026-2033)

Conforme os links de repositórios como o `FiscalNet` e `Calculadora RTC` indicam, o serviço deve estar preparado para o novo modelo de IVA (Imposto sobre Valor Agregado) Dual.

**Novos Tributos que substituirão os antigos:**

* **CBS (Contribuição sobre Bens e Serviços):** Federal. Substituirá PIS e COFINS.
* **IBS (Imposto sobre Bens e Serviços):** Subnacional (Estadual/Municipal). Substituirá ICMS e ISS.
* **IS (Imposto Seletivo):** Federal. Conhecido como "imposto do pecado", sobre produtos prejudiciais à saúde ou ambiente.

---

### 1. Impostos pagos por PJ: INSS e CadÚnico

* **INSS (CPP):** **Sim, aplica-se.** Toda empresa (PJ) que tem funcionários ou que paga pró-labore aos sócios deve recolher o INSS.
    * *Faz sentido na calculadora?* **Sim.** É um custo fixo de folha. Você deve calcular os 20% (Patronal) + RAT + Terceiros (para Lucro Presumido/Real) ou a alíquota unificada dentro do Simples Nacional.
* **CadÚnico:** **Não é um imposto.** O Cadastro Único é um registro do Governo Federal para identificar famílias de baixa renda para benefícios sociais (como o Bolsa Família).
    * *Faz sentido na calculadora?* **Não.** Uma empresa não "paga" CadÚnico. O que ocorre é que um MEI de baixa renda pode estar *inscrito* no CadÚnico para receber benefícios, mas não há cálculo tributário aqui.

---

### 2. Impostos pagos por PF: IPI, ICMS, ISS, FGTS, PIS, COFINS

Aqui precisamos separar o **Contribuinte de Direito** (quem emite a guia) do **Contribuinte de Fato** (você, quando compra um pão).

* **IPI, ICMS, ISS, PIS, COFINS:**
    * **Eles aplicam à PF?** Indiretamente sim, via **consumo**. Quando uma PF compra um celular, esses 5 impostos estão embutidos no preço. Mas a PF não "paga uma guia" de PIS.
    * *Faz sentido na calculadora?* **Depende do objetivo.**
        * Se a calculadora for de **Renda/Salário**, não faz sentido.
        * Se a calculadora for de **Consumo (Reforma Tributária)**, faz todo sentido, pois o objetivo da reforma (IBS/CBS) é justamente mostrar para a PF o quanto ela paga de imposto em cada produto.
* **FGTS:**
    * **Aplica-se?** O FGTS é um **direito** do trabalhador PF e um **custo** para a PJ. A Pessoa Física não paga FGTS; ela o recebe em uma conta vinculada.
    * *Faz sentido na calculadora?* Faz sentido apenas em um simulador de **"Salário Líquido"** ou **"Custo de Funcionário"**, para mostrar quanto está sendo depositado.

---

*********************************************************
*********************************************************
*********************************************************

## IMPLEMENTACAO

---

### Opção 1: O Modelo de "Engines" por Domínio (Recomendada)
Em vez de um serviço para cada imposto, agrupamos por **natureza do fato gerador**. Isso equilibra o isolamento com a facilidade de manutenção.

* **Micro-serviço `Tax-PF-Income`:** Cuida de tudo que incide sobre o "ganhar dinheiro" (**IRPF, INSS Retido**). Esses impostos compartilham a base de cálculos de folha e tabelas progressivas.
* **Micro-serviço `Tax-PF-Property`:** Cuida de bens (**IPVA, IPTU, ITR**). Eles dependem de tabelas de valores de mercado (FIPE, Planta Genérica de Valores).
* **Micro-serviço `Tax-PF-Transfer`:** Cuida de transações (**ITBI, ITCMD**). São cálculos eventuais disparados por eventos de compra/venda ou herança.

**Vantagem:** Se você alterar a regra do IRPF, os serviços de IPVA e ITBI continuam rodando 100% isolados, mas você não tem 10 serviços para gerenciar.

### SOLUÇÃO DE SEPARACAO DEFINIDA:

**Vamos de Opção 1 (Domínios)***, e a **pitada (2) será ter um biblioteca unica com as funções genéricas*** de arredondamento, regras para ler tabelas associada, acesso a banco de dados, compartilhar acesso ao banco/Redis. 

E para mitigar essa parte comum vamos implementar "futuramente" testes unitários gerais para PF, PJ. Sim, vamos executar os serviços com "(Kubernetes/Digital Ocean)" , vamos usar KEDA para controlar subida de PODs, KARPENTER para controlar provisionamento do cluster assim gerenciamos bem os custos de Cloud. 

