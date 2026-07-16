# Glossário do Projeto

- **Projeto:** Portal Corporativo de Gestão Tributária — Autonomia do Time de Finanças na Administração de Impostos
- **Código:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
- **Versão:** 1.0
- **Atualizado:** 2026-07-08
- **Objetivo:** Fonte única de verdade para a terminologia específica deste projeto, complementando o Glossário do programa pai PRJ-FIN-2026-0001.

---

## Nota Importante

Este glossário contém apenas os termos **específicos deste projeto**. Para a terminologia completa do domínio tributário (CBS, IBS, IS, IVA Dual, Lucro Real, Split Payment, NCM, NBS, CClassTrib, CFOP, DIFAL, Período Híbrido, Shadow Run, Princípio do Destino, etc.), consulte o [Glossário do Programa Pai](../PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/GLOSSARY.md).

---

## 1. Conceitos do Portal

### Portal Corporativo de Gestão Tributária

Plataforma web corporativa que concentra, valida e rastreia todas as tabelas de impostos da organização. É a ferramenta oficial de trabalho do time de Finanças para a administração do patrimônio fiscal da companhia.

- **Escopo:** Cadastro e manutenção de alíquotas, classificações fiscais, regimes tributários e benefícios; trilha de auditoria; importação e exportação em lote; dashboards gerenciais.
- **O que NÃO é:** Calculadora de impostos — o portal administra as tabelas que alimentam a calculadora corporativa, mas não realiza cálculos.
- **Sinônimos no contexto do projeto:** "o portal", "a plataforma", "a ferramenta de gestão fiscal"

### Fonte Única da Verdade Fiscal (Single Source of Truth)

Princípio de negócio que estabelece que toda alíquota, regime, classificação ou benefício fiscal praticado pela companhia deve estar registrado no Portal de Gestão Tributária — nenhuma tabela fiscal pode existir exclusivamente em planilhas, e-mails ou sistemas isolados.

- **Objetivo:** Eliminar a pulverização de informações fiscais e garantir que qualquer consulta ou auditoria encontre dados íntegros e atualizados em um único local.
- **Sinônimos no contexto do projeto:** "fonte única", "single source of truth", "base canônica"

---

## 2. Módulos do Portal

### Módulo 1 — Painel de Alíquotas Vigentes

Visão gerencial consolidada de todas as alíquotas ativas na companhia, com filtros por tributo, unidade federativa, município, período e status. Provê indicadores visuais de integridade, como sinalização de alíquotas prestes a expirar ou com potenciais conflitos.

### Módulo 2 — Cadastro e Manutenção de Alíquotas

Conjunto de funcionalidades para criação, edição e desativação de alíquotas e regras fiscais, com validações automáticas de negócio que impedem configurações inválidas (conflitos de vigência, referências inexistentes, transições inconsistentes).

### Módulo 3 — Gestão de Classificações e Regimes

Módulo de cadastro centralizado das classificações fiscais (NCM, NBS, CClassTrib, CFOP) e regimes tributários (Lucro Real, Lucro Presumido, Simples Nacional) que servem como base para aplicação das alíquotas.

### Módulo 4 — Linha do Tempo e Auditoria

Visualização cronológica de todas as alterações realizadas em qualquer tabela fiscal, com comparação entre versões sucessivas. Ferramenta primária do Controller e da Auditoria Interna para rastrear o histórico completo de mudanças.

### Módulo 5 — Importação e Exportação em Lote

Funcionalidade de carga massiva de alíquotas a partir de planilhas padronizadas, utilizada principalmente para absorver publicações oficiais (ex: alíquotas municipais de IBS divulgadas pelo Comitê Gestor). Também permite exportar as tabelas vigentes para relatórios e auditorias externas.

### Módulo 6 — Administração de Acessos

Gestão de usuários do portal e definição de perfis de acesso com privilégios segregados.

---

## 3. Perfis de Acesso e Governança

### Administrador Fiscal

Perfil de acesso com privilégios máximos no portal: pode cadastrar, editar e desativar quaisquer alíquotas e classificações; aprovar alterações propostas por Analistas Fiscais no fluxo de duas etapas; e gerenciar usuários e perfis de acesso.

- **Atribuído a:** Gerente Fiscal, Controller ou profissional sênior designado pelo Comitê Fiscal.
- **Restrição:** Não pode aprovar as próprias alterações (segregação de funções).

### Analista Fiscal

Perfil de acesso operacional: pode cadastrar, editar e desativar alíquotas e classificações dentro dos perímetros normais. Alterações de alto impacto (acima do patamar de materialidade) requerem aprovação de um Administrador Fiscal.

- **Atribuído a:** Analistas fiscais e contadores do time de Finanças.

### Auditor / Controller (Perfil de Leitura)

Perfil de acesso exclusivamente consultivo: pode visualizar todas as alíquotas, classificações, regimes e a trilha de auditoria completa, mas não pode realizar nenhuma alteração.

- **Atribuído a:** Controller, Auditores Internos, Comitê Fiscal (membros que não operam o cadastro).

### Trilha de Auditoria

Registro imutável e automático de toda alteração realizada nas tabelas fiscais do portal, contendo: identificação do usuário, data e hora, entidade e identificador afetados, valor anterior completo, novo valor completo e justificativa de negócio.

- **Requisito legal:** Atende aos controles internos exigidos pela Lei das S.A. e pelo framework COSO.
- **Retenção:** Pelo prazo legal aplicável a documentos fiscais.

### Fluxo de Aprovação em Duas Etapas

Mecanismo de governança que exige que alterações de alto impacto (definidas por patamar de materialidade) sejam propostas por um Analista Fiscal e aprovadas por um Administrador Fiscal ou Controller antes de entrarem em vigor.

- **Patamar de materialidade:** Definido pelo Comitê Fiscal com base no faturamento estimado impactado pela alteração (ex: acima de R$ 100 mil mensais).

### Patamar de Materialidade

Limite financeiro a partir do qual uma alteração de alíquota é considerada de alto impacto e requer fluxo de aprovação em duas etapas. Definido e revisado periodicamente pelo Comitê Fiscal.

### Segregação de Funções

Princípio de controle interno pelo qual nenhum usuário pode acumular privilégios incompatíveis. No contexto do portal: o perfil que cria uma alíquota de alto impacto não pode ser o mesmo que a aprova; o perfil de auditoria tem visibilidade plena mas nenhuma capacidade de alteração.

---

## 4. Fases e Entregas do Projeto

### Fase 0 — Fundamentação

Fase preparatória do projeto: mapeamento das tabelas fiscais atualmente em uso, levantamento das necessidades do time de Finanças, definição do modelo conceitual e prototipação das telas principais do portal.

### Entrega 1 — Portal: Gestão Básica de Alíquotas

Primeira entrega do projeto: liberação do portal com os Módulos 1 (Painel), 2 (Cadastro de Alíquotas) e 3 (Classificações e Regimes), além da carga inicial das tabelas fiscais vigentes. O time de Finanças passa a utilizar o portal como ferramenta primária para consulta e cadastro de alíquotas.

### Entrega 2 — Governança e Auditoria Fiscal

Segunda entrega do projeto: implantação do Módulo 6 (Administração de Acessos e Perfis) e do Módulo 4 (Linha do Tempo e Auditoria), estabelecendo a camada completa de controle de acesso com segregação de funções e a trilha de auditoria imutável. O time de Finanças passa a operar exclusivamente via portal.

### Entrega 3 — Portal: Operações em Escala

Terceira entrega do projeto: implantação dos fluxos de aprovação em duas etapas para alterações de alto impacto e do Módulo 5 (Importação/Exportação em Lote). O portal adquire capacidade de processar alterações em volume e com governança proporcional ao risco financeiro.

### Entrega 4 — Portal Completo: Expansão Funcional

Quarta e última entrega do projeto: implantação dos relatórios gerenciais de governança, dashboards de KPIs fiscais e preparação completa para o Período Híbrido. O portal atinge seu escopo completo como plataforma única de gestão tributária corporativa.

---

## 5. Métricas e Indicadores

### KPIs de Autonomia (A)

Indicadores que medem a redução da dependência do time técnico: A1 (Redução de Chamados Técnicos) e A2 (Tempo Médio de Efetivação de Ajustes).

### KPIs de Governança (G)

Indicadores que medem a integridade dos controles fiscais: G1 (Cobertura de Trilha de Auditoria), G2 (Conflitos Prevenidos) e G3 (Conformidade em Auditorias).

### KPIs de Eficiência e Satisfação (E)

Indicadores que medem a efetividade operacional e a experiência do usuário: E1 (Cobertura Geográfica), E2 (NPS Interno) e E3 (Taxa de Adoção).

### NPS Interno (Net Promoter Score)

Métrica de satisfação aplicada trimestralmente ao time de Finanças para medir a probabilidade de recomendação do portal a colegas. Calculado a partir da pergunta: "Em uma escala de 0 a 10, quanto você recomendaria o Portal de Gestão Tributária para um colega de outra empresa?"

### SUS (System Usability Scale)

Escala padronizada de usabilidade aplicada juntamente com o NPS para medir a facilidade de uso percebida do portal. Pontuação de 0 a 100; meta ≥ 75.

---

## 6. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-08 | Criação inicial: conceitos do portal, módulos, perfis de acesso, fases, métricas | Time de Negócios |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: domain-modeling, agile-ba-practices.*
