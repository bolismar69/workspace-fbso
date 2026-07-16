# Mecanismo de Transição — Reforma Tributária 2026

## ⚖️ Regras de Redução Fiscais (Vigência >= 01/01/2026)

* **Isenção Total:** Se a `BaseValue` for menor ou igual ao `transition_2026_floor`, o imposto é automaticamente zerado.
* **Redução Adicional:** Se a `BaseValue` estiver entre o `floor` e o `ceiling`, aplica-se a fórmula de abatimento proporcional:

$$\text{Redução} = fA - (fB \times \text{BaseValue})$$



> **Regra Importante:** O cálculo de redução não pode, sob hipótese alguma, tornar o valor final do imposto negativo.

## 🔌 Origem dos Fatores Dinâmicos

Conforme validado nas sessões de lacunas, os valores de `transition_2026_factor_a` e `transition_2026_factor_b` (assim como os mapas de dados `configs`) são fornecidos dinamicamente de forma externa ao banco pela função **`GetTableConfigs`**, encapsulada na biblioteca local vinculada ao projeto.

