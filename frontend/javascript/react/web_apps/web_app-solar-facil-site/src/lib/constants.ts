// Solar Fácil — Business logic constants
// Source: PPT SolarFacil v3 + 03-business-rules-analysis.md

import type { Plan, Metric, Step } from './types';

export const PLANS: Plan[] = [
  {
    name: 'Basic',
    price: 150,
    capacity: '100 a 200 kWh',
    capacityKwh: { min: 100, max: 200 },
    features: [
      'Monitoramento básico de consumo',
      'Relatórios mensais básicos',
      'Suporte via chat/e-mail',
    ],
    highlight: false,
  },
  {
    name: 'Special',
    price: 250,
    capacity: '200 a 350 kWh',
    capacityKwh: { min: 200, max: 350 },
    features: [
      'Monitoramento em tempo real',
      'Relatórios mensais básicos',
      'Suporte técnico prioritário',
      'Relatórios detalhados com insights',
      'Acesso a comunidades e fóruns',
    ],
    highlight: true,
  },
  {
    name: 'Premium',
    price: 400,
    capacity: '350 a 600 kWh',
    capacityKwh: { min: 350, max: 600 },
    features: [
      'Monitoramento avançado com alertas',
      'Relatórios personalizados',
      'Suporte técnico dedicado 24/7',
      'Acesso a comunidades e fóruns',
      'Descontos exclusivos',
    ],
    highlight: false,
  },
];

export const PROVIDER_RATE = 0.4; // R$ 0,40 por kWh excedente (BR-DER-004)

export const CONSUMER_DISCOUNT_RATE = 0.12; // 12% desconto médio (BR-DER-001)

export const CONSUMER_OUTLIER_MIN = 50; // BR-DER-005
export const CONSUMER_OUTLIER_MAX = 5000; // BR-DER-006
export const PROVIDER_OUTLIER_MIN = 50; // BR-DER-007
export const PROVIDER_OUTLIER_MAX = 10000; // BR-DER-008

export const METRICS: Metric[] = [
  { value: '500+', label: 'usuários beta' },
  { value: '4.8 ★', label: 'satisfação (NPS)' },
  { value: '12%', label: 'desconto médio' },
];

export const PROVIDER_METRICS: Metric[] = [
  { value: '3', label: 'cooperativas ativas' },
  { value: '1.000 kWh', label: 'compartilhados' },
  { value: 'R$ 0,40', label: 'por kWh excedente' },
];

export const HOW_IT_WORKS_STEPS: Step[] = [
  {
    icon: 'Sun',
    title: 'Produtor',
    description: 'Gera excedente de energia solar e compartilha via cooperativa',
  },
  {
    icon: 'RefreshCw',
    title: 'Cooperativa',
    description: 'Gerencia a distribuição da energia entre os participantes',
  },
  {
    icon: 'Home',
    title: 'Consumidor',
    description: 'Recebe energia limpa com desconto direto na conta de luz',
  },
];

export const DIFFERENTIATORS: Metric[] = [
  { value: 'R$ 0', label: 'Zero Capex — Sem investimento inicial' },
  { value: '100%', label: 'Pareamento Automático' },
  { value: '100%', label: 'Plataforma Legal (ANEEL)' },
  { value: 'API +', label: 'ANEEL — Escalável e transparente' },
];

export const FAQ_ITEMS = [
  {
    question: 'O que é energia solar compartilhada?',
    answer:
      'É um modelo regulamentado pela ANEEL onde o excedente de energia gerado por produtores solares é distribuído para consumidores dentro de uma mesma cooperativa, gerando desconto na conta de luz.',
  },
  {
    question: 'Preciso instalar painéis solares para participar?',
    answer:
      'Não! Se você é consumidor, basta escolher um plano e começar a receber energia compartilhada com desconto. A instalação é apenas para quem quer ser fornecedor.',
  },
  {
    question: 'Como funciona o desconto na conta de luz?',
    answer:
      'O desconto é calculado sobre o valor que você pagaria à concessionária. Nossos usuários têm em média 12% de desconto mensal.',
  },
  {
    question: 'É legalizado? Tem autorização da ANEEL?',
    answer:
      'Sim! A Solar Fácil opera dentro das normas da ANEEL (Resolução Normativa 687/2015).',
  },
  {
    question: 'Como recebo o dinheiro sendo fornecedor?',
    answer:
      'Você recebe R$ 0,40 por cada kWh excedente compartilhado. O pagamento é feito mensalmente via app Solar Fácil.',
  },
  {
    question: 'Tem fidelidade? Posso cancelar?',
    answer:
      'Não há período de fidelidade. Você pode cancelar a qualquer momento, sem multas.',
  },
];

export const SITE_URL = 'https://www.solarfacil.com.br';
export const CONTACT_EMAIL = 'contato@solarfacil.com.br';
export const WHATSAPP_NUMBER = '5511999999999'; // Placeholder — atualizar com número real
export const APP_STORE_URL = 'https://apps.apple.com/br/app/solar-facil/id0000000000'; // Placeholder
export const GOOGLE_PLAY_URL = 'https://play.google.com/store/apps/details?id=com.solarfacil'; // Placeholder
export const INSTAGRAM_URL = 'https://instagram.com/solarfacil';
export const LINKEDIN_URL = 'https://linkedin.com/company/solarfacil';
