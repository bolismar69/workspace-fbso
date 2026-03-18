# ADR 0004: Padronização de Stack Go (Gin vs Fiber)

## Data
2026-03-13

## Status
Aceito

## Contexto
A solução TaxNexus requer diferentes perfis de processamento:
1. **Perímetro (Gateway):** Exige alta compatibilidade com ecossistemas de segurança, middlewares padrão e estabilidade de protocolo.
2. **Core (Engine):** Exige latência ultra-baixa e alta vazão de processamento interno.

## Decisão
Adotaremos dois frameworks distintos baseados na responsabilidade do componente:

- **Gin (Baseado em net/http):** Utilizado para serviços de borda (Gateway/API). Justificativa: Maturidade, suporte extensivo a middlewares de auditoria e conformidade estrita com RFCs HTTP.
- **Fiber (Baseado em fasthttp):** Utilizado para motores de cálculo internos. Justificativa: Performance bruta, alocação zero de memória e otimização para tráfego interno de alta densidade.

## Consequências
- **Positivas:** Especialização técnica por perfil de carga; otimização de recursos de infraestrutura.
- **Negativas:** Necessidade de manter dois conjuntos de blueprints e middlewares específicos para cada framework no monorepo.