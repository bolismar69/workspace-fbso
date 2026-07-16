// Solar Fácil — Form validation (pure functions)
// Refs: BR-CON-001..006

import type { FormErrors, LeadForm } from './types';

export function validateName(name: string): string | null {
  if (!name || name.trim().length === 0) {
    return 'Nome é obrigatório';
  }
  if (name.trim().length < 2) {
    return 'Nome deve ter no mínimo 2 caracteres';
  }
  return null;
}

export function validateEmail(email: string): string | null {
  if (!email || email.trim().length === 0) {
    return 'E-mail é obrigatório';
  }
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email.trim())) {
    return 'Formato de e-mail inválido';
  }
  return null;
}

export function validatePhone(phone: string): string | null {
  if (!phone || phone.trim().length === 0) {
    return null; // Phone is optional
  }
  const digitsOnly = phone.replace(/\D/g, '');
  if (digitsOnly.length < 10 || digitsOnly.length > 11) {
    return 'Telefone deve ter 10 ou 11 dígitos (DDD + número)';
  }
  return null;
}

export function validateProfile(profile: string): string | null {
  if (!profile || profile.trim().length === 0) {
    return 'Selecione seu perfil';
  }
  return null;
}

export function validateMessage(message: string): string | null {
  if (message && message.length > 1000) {
    return 'Mensagem deve ter no máximo 1000 caracteres';
  }
  return null;
}

export function validateForm(data: LeadForm): FormErrors {
  return {
    name: validateName(data.name) ?? undefined,
    email: validateEmail(data.email) ?? undefined,
    phone: validatePhone(data.phone) ?? undefined,
    profile: validateProfile(data.profile) ?? undefined,
    message: validateMessage(data.message) ?? undefined,
  };
}

export function hasErrors(errors: FormErrors): boolean {
  return Object.values(errors).some((e) => e !== undefined);
}
