'use client';

import { useSearchParams } from 'next/navigation';
import { useContactForm } from '@/hooks/useContactForm';
import { Home, Zap, Handshake, MessageCircle } from 'lucide-react';
import type { PersonaProfile } from '@/lib/types';
import FormField from './FormField';
import SuccessScreen from './SuccessScreen';
import Button from '@/components/shared/Button';
import { WHATSAPP_NUMBER } from '@/lib/constants';

export default function ContactForm() {
  const searchParams = useSearchParams();
  const initialProfile = (searchParams.get('perfil') || '') as PersonaProfile;
  const plano = searchParams.get('plano') || '';
  const economia = searchParams.get('economia') || '';
  const excedente = searchParams.get('excedente') || '';
  const ganho = searchParams.get('ganho') || '';

  const initialMessage = plano
    ? `Plano sugerido: ${plano} — Economia estimada: ~R$ ${economia}/mês`
    : excedente
      ? `Excedente informado: ${excedente} kWh/mês — Ganho estimado: ~R$ ${ganho}/mês`
      : '';

  const {
    values,
    errors,
    isSubmitting,
    isSuccess,
    submitError,
    setField,
    handleSubmit,
    honeypotField,
  } = useContactForm({
    initialProfile,
    initialMessage,
  });

  if (isSuccess) {
    return <SuccessScreen />;
  }

  return (
    <form onSubmit={handleSubmit} noValidate>
      {/* Honeypot — hidden from humans, visible to bots */}
      <FormField
        label="Website"
        name={honeypotField}
        value=""
        onChange={() => {}}
        hidden
      />

      <FormField
        label="Nome completo"
        name="name"
        value={values.name}
        onChange={(v) => setField('name', v)}
        error={errors.name}
        required
        placeholder="Seu nome"
      />

      <FormField
        label="E-mail"
        name="email"
        type="email"
        value={values.email}
        onChange={(v) => setField('email', v)}
        error={errors.email}
        required
        placeholder="seu@email.com"
      />

      <FormField
        label="Telefone (WhatsApp)"
        name="phone"
        type="tel"
        value={values.phone}
        onChange={(v) => setField('phone', v)}
        error={errors.phone}
        placeholder="(11) 99999-9999"
      />

      {/* Profile radio buttons */}
      <div className="mb-4">
        <label className="mb-2 block text-sm font-medium text-solar-text">
          Sou <span className="text-red-500">*</span>
        </label>
        <div className="flex flex-wrap gap-3">
          {(['consumidor', 'fornecedor', 'cooperativa'] as const).map((p) => (
            <label
              key={p}
              className={`cursor-pointer rounded-lg border px-4 py-3 text-sm font-medium transition-colors ${
                values.profile === p
                  ? 'border-solar-primary bg-solar-primary-light text-solar-primary-dark'
                  : 'border-solar-border text-solar-text-muted hover:border-solar-primary/50'
              }`}
            >
              <input
                type="radio"
                name="profile"
                value={p}
                checked={values.profile === p}
                onChange={(e) => setField('profile', e.target.value)}
                className="sr-only"
              />
              {p === 'consumidor' && <><Home size={14} className="inline" /> Consumidor (quero economizar)</>}
              {p === 'fornecedor' && <><Zap size={14} className="inline" /> Fornecedor (quero compartilhar)</>}
              {p === 'cooperativa' && <><Handshake size={14} className="inline" /> Cooperativa (quero parceria)</>}
            </label>
          ))}
        </div>
        {errors.profile && (
          <p className="mt-1 text-sm text-red-500">{errors.profile}</p>
        )}
      </div>

      <FormField
        label="Mensagem"
        name="message"
        type="textarea"
        value={values.message}
        onChange={(v) => setField('message', v)}
        error={errors.message}
        placeholder="Conte-nos como podemos ajudar..."
      />

      {/* Submit error fallback */}
      {submitError && (
        <div className="mb-4 rounded-lg border border-solar-error-border bg-solar-error-surface p-4">
          <p className="text-sm text-solar-error-text">{submitError}</p>
          <a
            href={`https://wa.me/${WHATSAPP_NUMBER}`}
            target="_blank"
            rel="noopener noreferrer"
            className="mt-2 inline-flex items-center gap-1 text-sm font-medium text-solar-primary hover:underline"
          >
            <MessageCircle size={14} className="inline" /> Falar pelo WhatsApp
          </a>
        </div>
      )}

      <Button
        variant="primary"
        size="lg"
        type="submit"
        disabled={isSubmitting}
        className="w-full"
      >
        {isSubmitting ? 'Enviando...' : 'Enviar mensagem'}
      </Button>
    </form>
  );
}
