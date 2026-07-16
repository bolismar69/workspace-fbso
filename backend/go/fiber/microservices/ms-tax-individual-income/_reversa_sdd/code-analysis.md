# Análise Técnica Consolidada — ms-tax-individual-income

Gerado pelo agente **Archaeologist** em 2026-06-08.

## 🏗️ Módulos e Componentes

### 1. Handlers (`handlers/`)
🟢 **CONFIRMADO**
- **Propósito:** Ponto de entrada da API, parse de requisições e injeção de Trace ID.
- **Destaque:** Utiliza o middleware `requestid` do Fiber para rastreabilidade.
- **Fluxo:** Recebe JSON -> Parse para `UniversalTaxRequest` -> Chama `CalculationService.Calculate` -> Retorna JSON.

### 2. Services (`services/`)
🟢 **CONFIRMADO**
- **Propósito:** Núcleo da lógica de negócio e integração externa.
- **Algoritmos Principais:**
    - **Cálculo Paralelo:** Dispara os cálculos 'Completa' e 'Simplificada' simultaneamente usando goroutines.
    - **Lógica de Recomendação:** Compara os resultados e marca o mais vantajoso para o contribuinte.
    - **Transição 2026:** Aplica regras de isenção e redução adicional baseadas em faixas de renda (`transition_2026_floor`, `factor_a`, etc.) para datas >= 01/01/2026.
- **Integração:** Cliente HTTP para o microserviço de INSS (`FetchINSS`).

### 3. Data (`data/`)
🟢 **CONFIRMADO**
- **Propósito:** Configuração do banco de dados e carga inicial.
- **Dicionário de Dados Resumido:**
    - `tax_definitions`: Cadastro de impostos (ex: IRPF).
    - `tax_rules_history`: Histórico de faixas, alíquotas e parcelas a deduzir.
    - `tax_configs`: Parâmetros de negócio (limite simplificado, dedução por dependente, limites de educação/saúde).

## 📊 Fluxos de Controle (Resumo)

1. **IRPF Completo:**
    - Busca INSS no MS externo.
    - Calcula dedução por dependentes (mensal ou anual).
    - Aplica limite de gastos com educação por pessoa (titular + dependentes).
    - Subtrai total de deduções da renda bruta.
    - Busca regra na tabela progressiva e aplica fórmula: `(Base * Alíquota) - Parcela a Deduzir`.
    - Aplica ajuste de transição 2026 se necessário.

2. **IRPF Simplificado:**
    - Aplica desconto padrão de 20% sobre a renda bruta.
    - Limita o desconto ao teto vigente (mensal ou anual).
    - Segue o mesmo cálculo de imposto sobre a base resultante.

## ⚙️ Configurações e Metadados

- **Constantes de Negócio:** Carregadas dinamicamente da tabela `tax_configs`.
- **Dependências Externas:** `INSS_SERVICE_URL`, `DATABASE_URL`, `REDIS_ADDR`.
