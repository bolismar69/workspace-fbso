# Lacunas de Conhecimento — ms-tax-individual-income

> Identificadas pelo agente **Reviewer** em 2026-06-10.
> Preencha o campo **Resposta** para cada item e me avise quando terminar.

## 🔴 Bloqueantes para Reimplementação

### 1. Racional dos Fatores A e B na Transição 2026
- **Contexto:** No arquivo `services/calculation_service.go:284`, os valores `transition_2026_factor_a` e `transition_2026_factor_b` são usados para calcular uma redução adicional.
- **Pergunta:** Qual é o embasamento legal ou a regra de negócio que define esses fatores? Eles variam por região ou outra condição não mapeada?
- **Sugestão de Resposta:** Confirmar se são valores fixos da legislação federal ou se dependem de alguma tabela extra.
- **Resposta:* os fatores A e B, assim como as demais informações existentes na variável/map de nome 'configs' esta sendo obtida de uma biblioteca externa pela função [GetTableConfigs].* 

---

## 🟡 Desejáveis (Melhoria de Qualidade)

### 2. Detalhamento de Gastos com Saúde/Educação
- **Contexto:** O sistema recebe apenas o valor total de gastos.
- **Pergunta:** Existe algum requisito para validar o tipo de procedimento de saúde ou instituição de ensino, ou a responsabilidade de validação é exclusiva da fonte dos dados?
- **Resposta:* a responsabilidade de validação é do sistema de origem da informação, esse serviço somente tem a função de aplicar as regras de calculo* 
