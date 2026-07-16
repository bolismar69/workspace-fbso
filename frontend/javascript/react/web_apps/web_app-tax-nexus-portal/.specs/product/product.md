---
title: "Produto — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
tags: ["product", "vision", "explanation"]
---

# Visão do Produto — TaxNexus Portal (TaaS)

## 1. Propósito

O **TaxNexus Portal (TaaS — Tax as a Service)** é uma aplicação web que permite a empresas brasileiras simularem o impacto financeiro da **Reforma Tributária 2026** em suas operações. A ferramenta oferece uma comparação visual entre o sistema tributário atual (PIS, COFINS, ICMS, ISS, IPI) e o novo modelo de IVA Dual (CBS, IBS, Imposto Seletivo), auxiliando na tomada de decisão e planejamento fiscal.

## 2. Proposta de Valor

- **Simplicidade:** Interface minimalista — CNPJ + NCM + localidade → resultado comparativo
- **Visualização clara:** Cards lado a lado (Legado vs. Reforma) + gráfico de transição
- **Projeção de cenários:** Simulação para 2026 (transição) e 2027 (projeção)
- **Rastreabilidade:** Cada simulação recebe ID único para auditoria corporativa

## 3. Personas

| Persona | Necessidade | Contexto |
|---|---|---|
| **Controller Tributário** | Simular carga tributária para diferentes NCMs e jurisdições | Departamento fiscal de empresa |
| **CFO / Diretor Financeiro** | Projetar impacto da reforma no fluxo de caixa | Planejamento financeiro estratégico |
| **Consultor Tributário** | Gerar cenários comparativos para múltiplos clientes | Consultoria fiscal |
| **Desenvolvedor Integrador** | Consumir API de cálculo para integrar em ERPs | Integração de sistemas |

## 4. Escopo Atual (MVP — Fase 0)

- ✅ Autenticação simplificada por CNPJ
- ✅ Formulário de simulação: UF → Cidade → NCM → Saldo Remanescente
- ✅ Cálculo comparativo (chamada à API backend)
- ✅ Exibição de resultados em cards comparativos
- ✅ Gráfico de barras (2026 vs. 2027)
- ✅ Containerização Docker (build + deploy)

## 5. Fora do Escopo (MVP)

- ❌ Autenticação real (JWT/OAuth)
- ❌ Múltiplas simulações/histórico
- ❌ Persistência de preferências do usuário
- ❌ Multi-tenancy
- ❌ Exportação de relatórios (PDF, CSV)
- ❌ Comparação entre múltiplos cenários
- ❌ Design responsivo completo (mobile)
- ❌ Testes automatizados

## 6. Visão de Longo Prazo

TaxNexus evolui para uma plataforma SaaS completa de inteligência tributária:
- **Fase 1:** Multi-usuário com auth real, histórico de simulações
- **Fase 2:** Exportação de relatórios, comparação entre cenários
- **Fase 3:** Recomendações de otimização fiscal por IA
- **Fase 4:** Integração direta com ERPs (SAP, Totvs, Oracle)

🤖 *Documentação gerada por mineração reversa de especificações (spec-miner).*
