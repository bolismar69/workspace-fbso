import type { Metadata } from 'next';
import { Suspense } from 'react';
import SectionWrapper from '@/components/shared/SectionWrapper';
import Breadcrumb from '@/components/shared/Breadcrumb';
import Skeleton from '@/components/shared/Skeleton';
import ContactForm from '@/components/contact/ContactForm';
import JourneySummary from '@/components/contact/JourneySummary';
import DirectChannels from '@/components/contact/DirectChannels';

export const metadata: Metadata = {
  title: 'Contato — Solar Fácil',
  description:
    'Fale com a Solar Fácil. Preencha o formulário ou entre em contato direto por WhatsApp, e-mail ou redes sociais.',
};

function ContactFormFallback() {
  return (
    <div className="space-y-4" aria-label="Carregando formulário...">
      <Skeleton variant="input" />
      <Skeleton variant="input" />
      <Skeleton variant="input" />
      <Skeleton className="h-24" />
    </div>
  );
}

export default function ContatoPage() {
  return (
    <>
      <SectionWrapper id="contato-page">
        <Breadcrumb items={[{ label: 'Home', href: '/' }, { label: 'Contato' }]} />
        <h1 className="text-center text-3xl font-bold text-solar-text md:text-4xl [text-wrap:balance]">
          Fale com a Solar Fácil
        </h1>
        <p className="mt-4 text-center text-solar-text-muted">
          Estamos prontos para ajudar você a economizar ou compartilhar energia limpa.
        </p>

        <div className="mx-auto mt-12 max-w-lg">
          <div className="rounded-2xl border border-solar-border bg-solar-bg p-6 shadow-sm md:p-8">
            <Suspense fallback={<ContactFormFallback />}>
              <JourneySummary />
              <ContactForm />
            </Suspense>
          </div>
        </div>
      </SectionWrapper>

      <SectionWrapper id="canais-diretos" bg="alt">
        <DirectChannels />
      </SectionWrapper>
    </>
  );
}
