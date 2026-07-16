# ADR-0006 — Degradação graciosa do cache (erros de Redis não abortam a leitura)

> ADR **retroativo** inferido do código (sem histórico Git). · Detetive · 2026-06-10
> Confiança: 🟢 CONFIRMADO

## Status
Aceito (vigente).

## Contexto
O Redis é uma otimização, não a fonte da verdade — esta é o Postgres. Uma indisponibilidade de cache não deve impedir o cálculo de impostos.

## Decisão
Os erros de Redis são **deliberadamente engolidos** para preservar o caminho feliz:
- No `GET`: qualquer erro ou desserialização inválida cai para a consulta ao Postgres (não retorna erro). 🟢
- No `SET`: o erro de `json.Marshal` é ignorado (`cacheData, _ := json.Marshal(...)`) e a escrita no Redis não tem o retorno checado. 🟢
- Conexão Redis é **lazy** e não faz `Ping` (`cache/redis.go`) — falhas só aparecem no primeiro comando, e são absorvidas.
- Evidência: `tax_repository.go:66-70, 96-97, 108-113, 147-149`.

## Consequências
- 🟢 Resiliência: Redis fora do ar ⇒ sistema continua, mais lento, lendo do Postgres.
- 🟡 **Falhas de cache são silenciosas** — sem log/métrica de miss por erro. Dificulta diagnosticar degradação de performance (ex.: Redis intermitente passa despercebido). Recomendação para o Reviewer/Architect: instrumentar (log/contador) os erros engolidos.
- 🟡 Cache só é populado quando o resultado é não-vazio (`len > 0`); resultados vazios legítimos nunca são cacheados e sempre vão ao Postgres.
