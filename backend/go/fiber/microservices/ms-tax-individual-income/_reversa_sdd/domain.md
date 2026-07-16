# Glossário e Regras de Domínio — ms-tax-individual-income

Gerado pelo agente **Detective** em 2026-06-08.

## 📖 Glossário

| Termo | Definição | Confiança |
|-------|-----------|-----------|
| **IRPF** | Imposto de Renda Pessoa Física. | 🟢 CONFIRMADO |
| **Modelo Completo** | Modalidade de declaração onde todas as deduções permitidas por lei são detalhadas e subtraídas da base de cálculo. | 🟢 CONFIRMADO |
| **Modelo Simplificado** | Modalidade que substitui as deduções detalhadas por um desconto padrão (geralmente 20% da renda bruta), limitado a um teto. | 🟢 CONFIRMADO |
| **Dedução por Dependente** | Valor fixo subtraído da base de cálculo para cada dependente informado. | 🟢 CONFIRMADO |
| **PGBL** | Plano Gerador de Benefício Livre. Contribuição previdenciária dedutível até um limite percentual da renda bruta. | 🟢 CONFIRMADO |
| **Tabela Progressiva** | Conjunto de faixas de renda com alíquotas crescentes e parcelas a deduzir correspondentes. | 🟢 CONFIRMADO |
| **Transição 2026** | Regras especiais de isenção ou redução de imposto vigentes a partir de 01/01/2026, possivelmente devido a uma reforma tributária. | 🟢 CONFIRMADO |

## ⚖️ Regras de Negócio Principais

### 1. Seleção do Modelo de Cálculo
- **Regra:** O sistema deve calcular o imposto devido em ambos os modelos (Completo e Simplificado) para a mesma entrada de dados.
- **Efeito:** Marcar como `IsRecommended` o modelo que resultar no menor `TaxAmount` para o contribuinte.
- **Implementação:** Realizado via goroutines em paralelo para otimização de performance.

### 2. Limites de Dedução (Modelo Completo)
- **Dependentes:** Valor multiplicado pela quantidade de dependentes. Diferencia valores para base mensal vs anual.
- **Educação:** Limitado ao teto individual multiplicado pelo número de pessoas (titular + dependentes). Aplica o menor valor entre o gasto real e o limite calculado.
- **Saúde:** Sem limite explícito de teto (usando valor de "infinito" no SQL: 999.999.999.999,99).
- **PGBL:** Limitado a um percentual (padrão 12%) da renda bruta total.

### 3. Desconto Simplificado
- **Regra:** 20% da Renda Bruta.
- **Trava:** O desconto não pode ultrapassar o teto configurado (`simplified_discount_monthly_limit` ou `simplified_discount_annual_limit`).

### 4. Integração de Previdência (INSS)
- **Regra:** O valor do INSS deve ser obtido de um microserviço externo.
- **Resiliência:** Se a chamada externa falhar, o cálculo de IRPF prossegue sem a dedução previdenciária (gera um aviso no log).

### 5. Mecanismo de Transição Reforma 2026
- **Gatilho:** Data de referência >= 01/01/2026.
- **Isenção Total:** Se a `BaseValue` for menor ou igual ao `transition_2026_floor`, o imposto é zerado.
- **Redução Adicional:** Se a `BaseValue` estiver entre o `floor` e o `ceiling`, aplica-se uma fórmula de redução: `fA - (fB * BaseValue)`. A redução não pode tornar o imposto negativo. Os fatores fA e fB são fornecidos dinamicamente por biblioteca externa. 🟢

## 🔴 Lacunas e Dúvidas
- 🟢 **RESOLVIDO:** A validação de tipos de gastos (Educação vs Saúde) é responsabilidade do sistema de origem. Este serviço foca apenas na aplicação das regras matemáticas de dedução.
- 🟢 **RESOLVIDO:** A origem e rational dos "Fatores A e B" são geridos externamente e consumidos via `GetTableConfigs`, permitindo atualizações sem deploy de código.
