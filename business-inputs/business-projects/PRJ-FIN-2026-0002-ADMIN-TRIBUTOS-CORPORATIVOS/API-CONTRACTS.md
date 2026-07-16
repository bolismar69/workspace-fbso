# Contrato de API — Portal de Gestão Tributária

- **Projeto:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
- **Programa Pai:** PRJ-FIN-2026-0001 — Adequação Corporativa à Reforma Tributária Nacional
- **Data de Criação:** 11 de Julho de 2026
- **Versão:** 1.0
- **Status:** Definição Inicial — norteará o desenvolvimento de DT-1 (backend) e DT-2 (frontend)
- **Referências:**
  - [01-PROJECT-CHARTER.md](./01-PROJECT-CHARTER.md) — escopo de negócio e módulos do portal
  - [02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md) — requisitos BR-01 a BR-10
  - [03-EPICS.md](./03-EPICS.md) — épicos e features
  - [04-FEATURES.md](./04-FEATURES.md) — regras de negócio RN-01 a RN-38
  - [INTEGRATION-MAP.md](./INTEGRATION-MAP.md) — mapa de integração entre componentes

---

## 1. Objetivo

Este documento define o contrato de API entre o **Portal de Gestão Tributária (frontend)** e o **microserviço de administração tributária (backend)**, estabelecendo a "fonte da verdade" que ambos os times consomem durante o desenvolvimento paralelo (Cenário C do plano técnico).

O contrato abrange:
- Relação completa dos recursos (endpoints) que o backend expõe
- Modelos de request e response para cada operação
- Regras de autenticação e autorização por endpoint
- Política de erros, códigos HTTP e versionamento
- Exemplos de fluxos de consumo pelo frontend

---

## 2. Visão Geral da API

### 2.1 Recursos e Agrupamentos

A API organiza-se em 8 grupos de recursos, correspondentes aos módulos do portal definidos no [Project Charter](./01-PROJECT-CHARTER.md):

| Grupo | Recurso Raiz | Módulo(s) do Portal | Épico | BRs Vinculados |
|:---|:---|:---|:---|:---|
| **Alíquotas** | `/api/v1/aliquotas` | M1 (Painel), M2 (Cadastro) | 01 — Motor de Cadastro Fiscal | BR-01, BR-02 |
| **Classificações Fiscais** | `/api/v1/classificacoes` | M3 (Classificações e Regimes) | 01 — Motor de Cadastro Fiscal | BR-03 |
| **Regimes Tributários** | `/api/v1/regimes` | M3 (Classificações e Regimes) | 01 — Motor de Cadastro Fiscal | BR-03 |
| **Usuários e Perfis** | `/api/v1/usuarios` | M6 (Administração de Acessos) | 02 — Controle de Acesso | BR-04 |
| **Auditoria** | `/api/v1/auditoria` | M4 (Linha do Tempo) | 02 — Controle de Acesso | BR-05 |
| **Carga em Lote** | `/api/v1/lotes` | M5 (Importação/Exportação) | 03 — Operações em Escala | BR-07 |
| **Aprovações** | `/api/v1/aprovacoes` | Transversal | 03 — Operações em Escala | BR-06 |
| **Relatórios e Dashboards** | `/api/v1/relatorios` | Transversal | 04 — Inteligência Fiscal | BR-08, BR-09 |
| **Entidades Corporativas** | `/api/v1/empresas` | Transversal (infraestrutura) | 01 — Motor de Cadastro Fiscal | BR-01 (suporte) |

### 2.2 Autenticação e Autorização

| Aspecto | Definição |
|:---|:---|
| **Protocolo** | SAML 2.0 via Keycloak corporativo |
| **Transmissão de credenciais** | Todas as requisições exigem token de sessão via header `Authorization: Bearer <token>` |
| **Perfis de acesso** | `ADMINISTRADOR_FISCAL`, `ANALISTA_FISCAL`, `AUDITOR_CONTROLLER` — conforme RN-10 ([04-FEATURES.md](./04-FEATURES.md#feature-021-administração-de-acessos-e-perfis)) |
| **Segregação de funções** | Regra RN-11: um Administrador Fiscal não pode aprovar a própria alteração de alto impacto |
| **Rastreabilidade** | Toda operação de escrita registra o usuário autenticado na trilha de auditoria (RN-14) |

### 2.3 Versionamento

- **Estratégia:** Versionamento por URI (`/api/v1/`, `/api/v2/`)
- **Regra de compatibilidade:** Mudanças que quebram compatibilidade (remoção de campos, alteração de tipo) exigem nova versão. Adições de campos ou endpoints não quebram compatibilidade.
- **Depreciação:** Versões antigas são mantidas por 2 ciclos de entrega após o lançamento da nova versão, com header `Deprecation: true` e `Sunset: <data>` nas respostas.

### 2.4 Política de Erros

Toda resposta de erro segue o formato padronizado:

```json
{
  "codigo": "CONFLITO_VIGENCIA",
  "mensagem": "Já existe uma alíquota de IBS vigente para SP no período 01/01/2027 a 31/12/2027.",
  "detalhes": [
    {
      "campo": "inicio_validade",
      "mensagem": "Conflito com a alíquota ID 4521 (vigente de 01/01/2027 a 31/12/2027)"
    }
  ],
  "timestamp": "2026-09-15T14:30:00Z",
  "trace_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

| Código HTTP | Quando usar |
|:---|:---|
| `400` | Dados de entrada inválidos (validação de negócio ou formato) |
| `401` | Token ausente, expirado ou inválido |
| `403` | Token válido mas perfil sem permissão para o recurso |
| `404` | Recurso não encontrado (ex: alíquota inexistente) |
| `409` | Conflito de negócio (ex: vigência sobreposta, código duplicado) |
| `422` | Operação bloqueada por regra de negócio (ex: aprovar a própria alteração) |
| `429` | Limite de requisições excedido |
| `500` | Erro interno inesperado |

---

## 3. Recursos da API

### 3.1 Grupo: Alíquotas (`/api/v1/aliquotas`)

**Perfis com acesso:** `ADMINISTRADOR_FISCAL` (tudo), `ANALISTA_FISCAL` (criar/editar), `AUDITOR_CONTROLLER` (apenas consulta)

#### 3.1.1 Listar Alíquotas

```
GET /api/v1/aliquotas
```

**Filtros (query parameters):**

| Parâmetro | Tipo | Obrigatório | Descrição |
|:---|:---|:---|:---|
| `tributo` | string | Não | CBS, IBS, IS, ICMS, ISS, PIS, COFINS, IPI |
| `uf` | string | Não | Sigla da UF (2 caracteres) |
| `municipio_ibge` | string | Não | Código IBGE do município (7 dígitos) |
| `ncm` | string | Não | Código NCM (8 dígitos) ou grupo (4 dígitos) |
| `status` | string | Não | `vigente`, `programada`, `expirada` |
| `regime` | string | Não | `LEGADO` (ICMS/ISS/PIS/COFINS/IPI), `IVA_DUAL` (CBS/IBS/IS) |
| `empresa_id` | integer | Não | Identificador da empresa (multi-tenancy) |
| `pagina` | integer | Não | Número da página (default: 1) |
| `tamanho_pagina` | integer | Não | Registros por página (default: 50, max: 200) |

**Response (200):**
```json
{
  "dados": [
    {
      "id": 4521,
      "tributo": "IBS",
      "regime": "IVA_DUAL",
      "uf_destino": "SP",
      "municipio_destino_ibge": "3548708",
      "municipio_destino_nome": "São Bernardo do Campo",
      "ncm": "84713012",
      "aliquota_cbs": null,
      "aliquota_ibs_estadual": 5.20,
      "aliquota_ibs_municipal": 2.80,
      "aliquota_is": null,
      "inicio_validade": "2027-01-01",
      "final_validade": null,
      "status": "vigente",
      "origem_cadastro": "MANUAL",
      "lote_origem_id": null,
      "alertas": []
    }
  ],
  "paginacao": {
    "pagina": 1,
    "tamanho_pagina": 50,
    "total_registros": 2341,
    "total_paginas": 47
  }
}
```

#### 3.1.2 Obter Alíquota por ID

```
GET /api/v1/aliquotas/{id}
```

**Response (200):** Registro completo da alíquota com todos os campos, incluindo `criado_por`, `criado_em`, `atualizado_por`, `atualizado_em`.

#### 3.1.3 Criar Alíquota

```
POST /api/v1/aliquotas
```

**Request Body:**
```json
{
  "tributo": "IBS",
  "uf_destino": "SP",
  "municipio_destino_ibge": "3548708",
  "ncm": "84713012",
  "aliquota_ibs_estadual": 5.20,
  "aliquota_ibs_municipal": 2.80,
  "inicio_validade": "2027-01-01",
  "empresa_id": 1,
  "justificativa": "Publicação oficial Comitê Gestor IBS — Resolução CG-IBS 045/2026"
}
```

**Validações aplicadas antes da criação:**
- RN-01: Conflito de vigência para o mesmo tributo, mesma região e período sobreposto → `409 CONFLITO_VIGENCIA`
- RN-02: NCM deve existir na base de classificações → `400 NCM_INEXISTENTE`
- RN-04: `final_validade` deve ser posterior a `inicio_validade` → `400 DATA_INVALIDA`
- Se `tributo` do regime LEGADO e Período Híbrido, RN-03 emite alerta se não há alíquota substituta no regime IVA_DUAL

**Response (201):** Registro completo da alíquota criada.

#### 3.1.4 Editar Alíquota

```
PUT /api/v1/aliquotas/{id}
```

**Request Body:** Mesmos campos de criação. A edição gera um novo registro de auditoria com snapshot do estado anterior (RN-15).

**Response (200):** Registro atualizado.

#### 3.1.5 Desativar Alíquota (Encerrar Vigência)

```
PATCH /api/v1/aliquotas/{id}/desativar
```

**Request Body:**
```json
{
  "final_validade": "2026-12-31",
  "justificativa": "Encerramento do período de teste CBS 2026 conforme calendário constitucional"
}
```

**Validações:**
- RN-05: `justificativa` é obrigatória → `400 JUSTIFICATIVA_OBRIGATORIA`
- RN-03 (Período Híbrido): Alerta se alíquota LEGADO é desativada sem substituta IVA_DUAL → resposta inclui `alertas: [{codigo: "SUBSTITUTA_AUSENTE", ...}]`

**Response (200):** Registro atualizado com `final_validade` preenchido e `status: "expirada"`.

#### 3.1.6 Consultar Histórico de uma Alíquota

```
GET /api/v1/aliquotas/{id}/historico
```

**Response (200):** Lista cronológica de todas as versões anteriores da alíquota (snapshots da trilha de auditoria), com diff entre versões.

---

### 3.2 Grupo: Classificações Fiscais (`/api/v1/classificacoes`)

Gerencia os códigos de classificação fiscal: NCM, NBS, CClassTrib, CFOP.

| Método | Path | Descrição | Perfis |
|:---|:---|:---|:---|
| `GET` | `/api/v1/classificacoes` | Listar classificações (filtros: `tipo`, `codigo`, `status`) | Todos |
| `GET` | `/api/v1/classificacoes/{id}` | Obter classificação por ID | Todos |
| `POST` | `/api/v1/classificacoes` | Criar classificação (RN-07: código único; RN-08: formato por tipo) | Admin, Analista |
| `PUT` | `/api/v1/classificacoes/{id}` | Editar classificação | Admin, Analista |
| `PATCH` | `/api/v1/classificacoes/{id}/desativar` | Desativar classificação (RN-06: bloqueado se há alíquotas vinculadas) | Admin |
| `GET` | `/api/v1/classificacoes/{id}/aliquotas` | Listar alíquotas vinculadas a esta classificação | Todos |

**Campos comuns de request/response:**
```json
{
  "id": 342,
  "tipo": "NCM",
  "codigo": "84713012",
  "descricao": "Máquinas automáticas para processamento de dados — portáteis",
  "status": "ativo",
  "total_aliquotas_vinculadas": 12
}
```

---

### 3.3 Grupo: Regimes Tributários (`/api/v1/regimes`)

| Método | Path | Descrição | Perfis |
|:---|:---|:---|:---|
| `GET` | `/api/v1/regimes` | Listar regimes (Lucro Real, Lucro Presumido, Simples Nacional) | Todos |
| `POST` | `/api/v1/regimes` | Criar regime | Admin |
| `PUT` | `/api/v1/regimes/{id}` | Editar regime | Admin |
| `PATCH` | `/api/v1/regimes/{id}/padrao` | Marcar como regime padrão (RN-09: exatamente um deve ser padrão) | Admin |

---

### 3.4 Grupo: Usuários e Perfis (`/api/v1/usuarios`)

**Perfil com acesso:** Exclusivamente `ADMINISTRADOR_FISCAL`

| Método | Path | Descrição |
|:---|:---|:---|
| `GET` | `/api/v1/usuarios` | Listar usuários do portal |
| `GET` | `/api/v1/usuarios/{id}` | Obter usuário por ID |
| `POST` | `/api/v1/usuarios` | Cadastrar novo usuário e definir perfil |
| `PUT` | `/api/v1/usuarios/{id}` | Alterar perfil do usuário |
| `PATCH` | `/api/v1/usuarios/{id}/desativar` | Desativar usuário |

**Matriz de permissões por perfil** (RN-10):
| Ação | ADMINISTRADOR_FISCAL | ANALISTA_FISCAL | AUDITOR_CONTROLLER |
|:---|:---:|:---:|:---:|
| Consultar alíquotas, classificações e regimes | ✅ | ✅ | ✅ |
| Criar e editar alíquotas e classificações | ✅ | ✅ | ❌ |
| Aprovar alterações de alto impacto | ✅ (exceto as próprias — RN-11) | ❌ | ❌ |
| Gerenciar usuários e perfis | ✅ | ❌ | ❌ |
| Visualizar trilha de auditoria | ✅ | ✅ | ✅ |
| Exportar tabelas e relatórios | ✅ | ✅ | ✅ |

**Response de criação de usuário (201):**
```json
{
  "id": 27,
  "nome": "Maria Aparecida Silva",
  "email": "maria.silva@empresa.com.br",
  "perfil": "ANALISTA_FISCAL",
  "status": "ativo",
  "empresa_id": 1,
  "criado_em": "2026-09-15T10:30:00Z"
}
```

---

### 3.5 Grupo: Auditoria (`/api/v1/auditoria`)

**Perfis com acesso:** Todos (somente leitura — RN-17)

| Método | Path | Descrição |
|:---|:---|:---|
| `GET` | `/api/v1/auditoria` | Listar eventos de auditoria com filtros: `entidade_tipo`, `entidade_id`, `usuario_id`, `operacao`, `data_inicio`, `data_fim` |
| `GET` | `/api/v1/auditoria/{id}` | Obter registro de auditoria com diff completo (estado anterior × novo estado) |

**Response da listagem (200):**
```json
{
  "dados": [
    {
      "id": 98765,
      "entidade_tipo": "ALIQUOTA",
      "entidade_id": 4521,
      "operacao": "EDICAO",
      "usuario_nome": "João Carlos Pereira",
      "usuario_perfil": "ANALISTA_FISCAL",
      "justificativa": "Correção da alíquota IBS municipal conforme errata CG-IBS",
      "data_hora": "2026-09-15T14:22:00Z",
      "resumo": "aliquota_ibs_municipal: 2.50% → 2.80%"
    }
  ],
  "paginacao": { "...": "..." }
}
```

**Regras de auditoria:**
- RN-14: Registro automático, imutável — nenhum perfil pode alterar ou excluir
- RN-15: Cada registro contém: usuário, timestamp, entidade, estado anterior (snapshot completo), novo estado, operação, justificativa
- RN-16: Retenção mínima de 5 anos (tributos federais), podendo estender a 10 anos

---

### 3.6 Grupo: Carga em Lote (`/api/v1/lotes`)

**Perfis com acesso:** `ADMINISTRADOR_FISCAL` (aprovar/rejeitar), `ANALISTA_FISCAL` (enviar)

#### 3.6.1 Enviar Arquivo para Carga

```
POST /api/v1/lotes
Content-Type: multipart/form-data
```

**Request:**
- `arquivo`: planilha Excel (.xlsx) ou CSV seguindo o template padronizado
- `tributo`: tipo de tributo das alíquotas no arquivo
- `justificativa`: motivo da carga em lote
- `empresa_id`: empresa destinatária

**Response (202 — Accepted):**
```json
{
  "lote_id": 183,
  "status": "EM_VALIDACAO",
  "total_linhas": 5570,
  "linhas_processadas": 0,
  "linhas_aceitas": 0,
  "linhas_rejeitadas": 0,
  "mensagem": "Arquivo recebido e em processamento. Consulte GET /api/v1/lotes/183 para acompanhar."
}
```

#### 3.6.2 Consultar Status do Lote

```
GET /api/v1/lotes/{lote_id}
```

**Response (200):**
```json
{
  "lote_id": 183,
  "status": "AGUARDANDO_APROVACAO",
  "tributo": "IBS",
  "nome_arquivo": "ibs_municipios_sp_2027.xlsx",
  "enviado_por": "João Carlos Pereira",
  "enviado_em": "2026-09-15T14:00:00Z",
  "total_linhas": 5570,
  "linhas_processadas": 5570,
  "linhas_aceitas": 5498,
  "linhas_rejeitadas": 72,
  "linhas_com_alertas": 15,
  "empresa_id": 1
}
```

#### 3.6.3 Listar Itens do Lote

```
GET /api/v1/lotes/{lote_id}/itens?status=rejeitado
```

Retorna os itens individuais do lote com seu status (`ACEITO`, `REJEITADO`, `COM_ALERTA`) e, para rejeitados, o motivo específico (RN-24).

#### 3.6.4 Aprovar Lote

```
POST /api/v1/lotes/{lote_id}/aprovar
```

**Perfil:** `ADMINISTRADOR_FISCAL` (não pode aprovar o próprio lote — RN-11)

**Request Body:**
```json
{
  "justificativa": "Carga validada — Resolução CG-IBS 045/2026. 72 rejeições revisadas manualmente."
}
```

**Efeito:** Itens com status `ACEITO` são efetivados nas tabelas finais de alíquotas. Cada item aceito gera um registro de auditoria (RN-25). Itens com status `REJEITADO` permanecem no lote para correção e reenvio.

#### 3.6.5 Rejeitar Lote

```
POST /api/v1/lotes/{lote_id}/rejeitar
```

**Request Body:**
```json
{
  "justificativa": "Divergência de formato nas colunas de alíquota — rever template."
}
```

**Efeito:** Nenhum item é efetivado. O lote passa ao status `REJEITADO`. O analista pode corrigir e reenviar.

#### 3.6.6 Histórico de Lotes

```
GET /api/v1/lotes?status=aprovado&data_inicio=2026-09-01&data_fim=2026-09-30
```

Lista todos os lotes com filtros (RN-27).

---

### 3.7 Grupo: Aprovações (`/api/v1/aprovacoes`)

**Perfil com acesso:** `ADMINISTRADOR_FISCAL` (aprovar), `ANALISTA_FISCAL` (criar propostas que disparam aprovação)

#### 3.7.1 Listar Pendências de Aprovação

```
GET /api/v1/aprovacoes/pendentes
```

**Response (200):** Lista de alterações de alíquotas com status `PENDENTE_APROVACAO` que aguardam ação do Administrador Fiscal.

**Regras:**
- RN-18: O gatilho de aprovação é acionado quando o impacto estimado da alteração atinge o patamar de materialidade definido pelo Comitê Fiscal
- RN-19: Enquanto pendente, a alíquota anterior permanece vigente
- RN-22: Alterações pendentes há mais de 5 dias úteis geram alerta ao Gerente Fiscal

#### 3.7.2 Aprovar Alteração

```
POST /api/v1/aprovacoes/{id}/aprovar
```

**Validação RN-20:** O aprovador não pode ser o mesmo usuário que propôs a alteração → `422 SEGREGACAO_FUNCOES`

#### 3.7.3 Rejeitar Alteração

```
POST /api/v1/aprovacoes/{id}/rejeitar
```

**Request Body (justificativa obrigatória — RN-21):**
```json
{
  "justificativa": "Alíquota proposta conflita com benefício fiscal vigente para a Zona Franca de Manaus."
}
```

---

### 3.8 Grupo: Relatórios e Dashboards (`/api/v1/relatorios`)

**Perfis com acesso:** `ADMINISTRADOR_FISCAL` e `AUDITOR_CONTROLLER` (acesso completo), `ANALISTA_FISCAL` (apenas Dashboard de Vencimentos — RN-33)

#### 3.8.1 Relatório Mensal de Governança

```
GET /api/v1/relatorios/governanca?mes=2026-08&formato=pdf
```

RN-28: Gerado automaticamente no 1º dia útil. Disponível para download em PDF ou Excel (RN-31). Parâmetro `formato` aceita `pdf` ou `xlsx`.

#### 3.8.2 Dashboard de Cobertura Fiscal

```
GET /api/v1/relatorios/dashboard/cobertura
```

Retorna dados agregados para o mapa de cobertura geográfica: municípios com alíquotas cadastradas vs. total de municípios com operação (RN-32, RN-34).

#### 3.8.3 Dashboard de Atividade

```
GET /api/v1/relatorios/dashboard/atividade?mes=2026-08
```

Retorna volume de alterações por tipo de operação (criação/edição/desativação) e por tributo.

#### 3.8.4 Dashboard de Vencimentos

```
GET /api/v1/relatorios/dashboard/vencimentos
```

Retorna alíquotas com vigência a expirar nas próximas 12 semanas (RN-07, RN-08 das user stories). Acessível também ao perfil `ANALISTA_FISCAL`.

---

### 3.9 Grupo: Entidades Corporativas (`/api/v1/empresas`)

Recursos de infraestrutura multi-tenancy. Permitem segregar tabelas fiscais por empresa do grupo econômico.

| Método | Path | Descrição | Perfis |
|:---|:---|:---|:---|
| `GET` | `/api/v1/empresas` | Listar empresas | Admin, Auditor |
| `GET` | `/api/v1/empresas/{id}` | Obter empresa por ID | Admin, Auditor |
| `POST` | `/api/v1/empresas` | Cadastrar nova empresa (com CNPJ raiz) | Admin |
| `PUT` | `/api/v1/empresas/{id}` | Editar dados da empresa | Admin |
| `PATCH` | `/api/v1/empresas/{id}/desativar` | Desativar empresa (não remove dados fiscais) | Admin |
| `GET` | `/api/v1/empresas/{id}/tenants` | Listar tenants/estabelecimentos da empresa | Admin, Auditor |

---

## 4. Entidades Transversais

### 4.1 Multi-Tenancy e Rastreabilidade de Origem

Toda tabela de alíquotas e classificações inclui as seguintes colunas transversais:

| Coluna | Tipo | Descrição |
|:---|:---|:---|
| `empresa_id` | integer | Identificador da empresa/grupo econômico (FK → `empresas`) |
| `tenant_id` | integer | Identificador do estabelecimento dentro da empresa (FK → `tenants`) |
| `origem_cadastro` | varchar | `MANUAL` (cadastro via formulário do portal) ou `LOTE` (originado de carga em lote) |
| `lote_origem_id` | integer | Se `origem_cadastro = LOTE`, FK para o lote de origem |
| `lote_item_origem_id` | integer | Se `origem_cadastro = LOTE`, FK para o item específico dentro do lote |

Estas colunas garantem rastreabilidade completa: para qualquer alíquota, é possível determinar se ela foi criada manualmente (por qual usuário, em qual data) ou via carga em lote (de qual arquivo, em qual linha, aprovada por quem).

### 4.2 Modelo de Auditoria Unificada

Além da tabela `iva_dual_rules_log` já existente no motor de cálculo (que audita apenas o IVA Dual), o novo microserviço de administração implementa auditoria unificada para **todas** as entidades:

| Entidade Auditada | Eventos Registrados |
|:---|:---|
| Alíquotas (todos os tributos) | Criação, edição, desativação, aprovação, rejeição |
| Classificações fiscais | Criação, edição, desativação |
| Regimes tributários | Criação, edição, alteração de padrão |
| Usuários | Criação, alteração de perfil, desativação |
| Lotes | Envio, conclusão de validação, aprovação, rejeição |

Cada registro segue o padrão RN-15: snapshot completo do estado anterior + novo estado em formato estruturado.

### 4.3 Tabelas de Apoio — Carga em Lote

O fluxo de carga em lote utiliza duas tabelas intermediárias (staging):

| Tabela | Função |
|:---|:---|
| `lotes_carga` | Cabeçalho do lote: arquivo de origem, tributo, status (`EM_VALIDACAO`, `AGUARDANDO_APROVACAO`, `APROVADO`, `REJEITADO`), usuário que enviou, datas, contadores de linhas |
| `lotes_carga_itens` | Itens do lote: cada linha do arquivo com seu conteúdo original, resultado da validação (`ACEITO`, `REJEITADO`, `COM_ALERTA`), motivo da rejeição, e ID da entidade criada após aprovação |

Somente quando um lote é **aprovado**, os itens com status `ACEITO` são efetivados nas tabelas finais de alíquotas. Até lá, permanecem isolados nas tabelas de staging, sem impacto nas regras vigentes.

---

## 5. Fluxos de Consumo (Frontend → Backend)

### 5.1 Fluxo: Analista Fiscal cadastra uma nova alíquota

```
[Portal]──POST /api/v1/aliquotas──►[Backend]
                                      │
                                      ├── Valida RN-01 (conflito vigência)
                                      ├── Valida RN-02 (NCM existe?)
                                      ├── Valida RN-04 (datas)
                                      ├── Verifica impacto (RN-18)
                                      │     │
                                      │     ├── Abaixo do patamar: grava direto, retorna 201
                                      │     └── Acima do patamar: grava com status PENDENTE_APROVACAO,
                                      │       notifica Administrador Fiscal, retorna 202
                                      │
                                      └── Gera registro de auditoria (RN-14)
```

### 5.2 Fluxo: Carga em lote de alíquotas IBS municipais

```
[Portal]──POST /api/v1/lotes (multipart)──►[Backend]
                                              │
                                              ├── Valida template (RN-23)
                                              ├── Cria lote (status: EM_VALIDACAO)
                                              ├── Processa cada linha (RN-24)
                                              │     ├── Valida RN-01 a RN-05 por linha
                                              │     ├── ACEITO / REJEITADO / COM_ALERTA
                                              ├── Atualiza lote (status: AGUARDANDO_APROVACAO)
                                              └── Retorna 202 com lote_id

[Portal]──GET /api/v1/lotes/{id}/itens──►[Backend]  (analista confere rejeições)

[Portal]──POST /api/v1/lotes/{id}/aprovar──►[Backend]
                                              │
                                              ├── Valida RN-11 (aprovador ≠ enviador)
                                              ├── Efetiva itens ACEITOS nas tabelas finais
                                              ├── Gera registros de auditoria (RN-25)
                                              └── Retorna 200 com sumário
```

### 5.3 Fluxo: Controller audita alteração

```
[Portal]──GET /api/v1/auditoria?entidade_tipo=ALIQUOTA&data_inicio=2026-09-01──►[Backend]
                                                                                    │
                                                                                    └── Retorna lista de eventos

[Portal]──GET /api/v1/auditoria/{id}──►[Backend]
                                          │
                                          └── Retorna diff completo (antes × depois)
```

---

## 6. Matriz de Endpoints × Perfis

| Método | Path | ADMINISTRADOR_FISCAL | ANALISTA_FISCAL | AUDITOR_CONTROLLER |
|:---|:---|:---:|:---:|:---:|
| `GET` | `/api/v1/aliquotas` | ✅ | ✅ | ✅ |
| `GET` | `/api/v1/aliquotas/{id}` | ✅ | ✅ | ✅ |
| `POST` | `/api/v1/aliquotas` | ✅ | ✅ | ❌ |
| `PUT` | `/api/v1/aliquotas/{id}` | ✅ | ✅ | ❌ |
| `PATCH` | `/api/v1/aliquotas/{id}/desativar` | ✅ | ✅ | ❌ |
| `GET` | `/api/v1/aliquotas/{id}/historico` | ✅ | ✅ | ✅ |
| `GET` | `/api/v1/classificacoes` | ✅ | ✅ | ✅ |
| `POST` | `/api/v1/classificacoes` | ✅ | ✅ | ❌ |
| `PUT` | `/api/v1/classificacoes/{id}` | ✅ | ✅ | ❌ |
| `PATCH` | `/api/v1/classificacoes/{id}/desativar` | ✅ | ❌ | ❌ |
| `GET` | `/api/v1/regimes` | ✅ | ✅ | ✅ |
| `POST` `/PUT` | `/api/v1/regimes[/{id}]` | ✅ | ❌ | ❌ |
| `GET` `/POST` `/PUT` `/PATCH` | `/api/v1/usuarios[/{id}][/desativar]` | ✅ | ❌ | ❌ |
| `GET` | `/api/v1/auditoria` | ✅ | ✅ | ✅ |
| `POST` | `/api/v1/lotes` | ✅ | ✅ | ❌ |
| `GET` | `/api/v1/lotes[/{id}][/itens]` | ✅ | ✅ | ✅ |
| `POST` | `/api/v1/lotes/{id}/aprovar` | ✅ | ❌ | ❌ |
| `POST` | `/api/v1/lotes/{id}/rejeitar` | ✅ | ❌ | ❌ |
| `GET` | `/api/v1/aprovacoes/pendentes` | ✅ | ❌ | ✅ |
| `POST` | `/api/v1/aprovacoes/{id}/aprovar` | ✅ (exceto própria — RN-11) | ❌ | ❌ |
| `POST` | `/api/v1/aprovacoes/{id}/rejeitar` | ✅ | ❌ | ❌ |
| `GET` | `/api/v1/relatorios/*` | ✅ | Parcial (apenas vencimentos) | ✅ |
| `GET` | `/api/v1/empresas[/{id}][/tenants]` | ✅ | ❌ | ✅ |
| `POST` `/PUT` `/PATCH` | `/api/v1/empresas[/{id}][/desativar]` | ✅ | ❌ | ❌ |

---

## 7. Governança do Contrato

### 7.1 Processo de Alteração

```
Necessidade de mudança (negócio ou técnico)
        │
        ▼
[PM + PO] Avaliam impacto nos requisitos de negócio
        │
        ▼
[Tech Lead Backend + Tech Lead Frontend] Discutem viabilidade e impacto no contrato
        │
        ▼
[Arquiteto de Solução] Aprova a alteração no contrato
        │
        ▼
[API-CONTRACTS.md] Atualizado com a nova versão
        │
        ├──► [DT-1 Backend] OpenAPI YAML atualizado → código Java implementa
        │
        └──► [DT-2 Frontend] Tipos TypeScript regenerados → código React atualizado
```

### 7.2 Responsabilidades

| Atividade | Responsável |
|:---|:---|
| Manter o contrato atualizado | Tech Lead Backend (DT-1) |
| Revisar alterações propostas | Tech Lead Frontend (DT-2) |
| Aprovar mudanças no contrato | Arquiteto de Solução |
| Garantir aderência ao contrato no desenvolvimento | Ambos os Tech Leads (sincronização semanal) |
| Versionar o contrato (MAJOR.MINOR) | Tech Lead Backend |

### 7.3 Sincronização entre Times

- **Frequência:** Reunião semanal de 30 minutos entre Tech Lead Backend e Tech Lead Frontend
- **Pauta fixa:** (a) alterações no contrato na última semana, (b) divergências encontradas entre mock e implementação real, (c) endpoints pendentes de implementação
- **Ferramenta:** OpenAPI YAML versionado no repositório do DT-1, cópia sincronizada no repositório do DT-2

---

🤖 *Documento gerado com apoio de Claude Code (Anthropic), em 11 de Julho de 2026.*
