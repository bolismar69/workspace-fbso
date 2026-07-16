// Solar Fácil — Core type definitions
// Refs: BR-STR-001..007, DT-001, DT-002, DT-003

export type PlanName = 'Basic' | 'Special' | 'Premium';

export type PersonaProfile = 'consumidor' | 'fornecedor' | 'cooperativa';

export interface ConsumerResult {
  economy: number;
  suggestedPlan: PlanName | null;
  isOutlier: boolean;
  message: string;
}

export interface ProviderResult {
  gain: number;
  rate: number;
  isOutlier: boolean;
  message: string;
}

export interface Plan {
  name: PlanName;
  price: number;
  capacity: string;
  capacityKwh: { min: number; max: number };
  features: string[];
  highlight: boolean;
}

export interface Metric {
  value: string;
  label: string;
  icon?: string;
}

export interface Step {
  icon: string;
  title: string;
  description: string;
}

export interface LeadForm {
  name: string;
  email: string;
  phone: string;
  profile: PersonaProfile | '';
  message: string;
}

export interface FormErrors {
  name?: string;
  email?: string;
  phone?: string;
  profile?: string;
  message?: string;
}
