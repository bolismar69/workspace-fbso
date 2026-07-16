'use client';

import { useState, useCallback } from 'react';
import type { LeadForm, FormErrors, PersonaProfile } from '@/lib/types';
import { validateForm, hasErrors } from '@/lib/validation';

const FORM_ENDPOINT = process.env.NEXT_PUBLIC_FORM_ENDPOINT || 'https://formspree.io/f/placeholder';
const MIN_SUBMIT_TIME_MS = 3000; // Anti-spam: 3s minimum (BR-CON-006)
const HONEYPOT_FIELD = 'website'; // Hidden field for bots

interface UseContactFormProps {
  initialProfile?: PersonaProfile;
  initialMessage?: string;
}

export function useContactForm(props?: UseContactFormProps) {
  const [values, setValues] = useState<LeadForm>({
    name: '',
    email: '',
    phone: '',
    profile: props?.initialProfile || '',
    message: props?.initialMessage || '',
  });
  const [errors, setErrors] = useState<FormErrors>({});
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isSuccess, setIsSuccess] = useState(false);
  const [submitError, setSubmitError] = useState<string | null>(null);
  const [pageLoadTime] = useState(() => Date.now());

  const setField = useCallback((field: keyof LeadForm, value: string) => {
    setValues((prev) => ({ ...prev, [field]: value }));
    setErrors((prev) => ({ ...prev, [field]: undefined }));
    setSubmitError(null); // Clear submit error on any field change
  }, []);

  const handleSubmit = useCallback(async (e?: React.FormEvent) => {
    e?.preventDefault();

    // Validate
    const validationErrors = validateForm(values);
    setErrors(validationErrors);

    if (hasErrors(validationErrors)) {
      return;
    }

    // Anti-spam: minimum time check
    const elapsed = Date.now() - pageLoadTime;
    if (elapsed < MIN_SUBMIT_TIME_MS) {
      // Block silently — show fake success
      setIsSuccess(true);
      return;
    }

    setIsSubmitting(true);
    setSubmitError(null);

    try {
      const formData = new FormData();
      formData.append('name', values.name);
      formData.append('email', values.email);
      formData.append('phone', values.phone);
      formData.append('profile', values.profile);
      formData.append('message', values.message);
      // Honeypot must be empty
      formData.append(HONEYPOT_FIELD, '');

      const response = await fetch(FORM_ENDPOINT, {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}`);
      }

      setIsSuccess(true);
    } catch {
      setSubmitError(
        'Não foi possível enviar. Verifique sua conexão ou fale conosco pelo WhatsApp.',
      );
    } finally {
      setIsSubmitting(false);
    }
  }, [values]);

  return {
    values,
    errors,
    isSubmitting,
    isSuccess,
    submitError,
    setField,
    handleSubmit,
    honeypotField: HONEYPOT_FIELD,
  };
}
