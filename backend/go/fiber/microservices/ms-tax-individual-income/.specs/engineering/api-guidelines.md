# API Guidelines e Padrões Globais

## 🚀 Tratamento de Erros e Rastreabilidade

* **Trace ID:** Todas as requisições que passam pelo handler recebem o cabeçalho injetado pelo middleware `requestid` do Fiber. Ele deve ser propagado em logs para debug distribuído.
* **Estrutura Padrão de Erro:** Retorno consistente em `ErrorResponse` com chave unificada `error` em formato JSON.
* **Resiliência HTTP:** Falhas de comunicação com serviços terceiros devem gerar alertas sem abortar fluxos essenciais de processamento caso haja fallback.

