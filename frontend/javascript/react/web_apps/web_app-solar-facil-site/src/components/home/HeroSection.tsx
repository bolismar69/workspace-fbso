'use client';

import { trackCtaClick } from '@/lib/analytics';
import Button from '@/components/shared/Button';

export default function HeroSection() {
  const scrollTo = (id: string, ctaType: 'consumidor' | 'fornecedor') => {
    trackCtaClick({ cta_type: ctaType, location: 'hero' });
    document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' });
  };

  return (
    <section className="relative flex min-h-screen flex-col items-center justify-center px-4 text-center">
      {/* Background gradient */}
      <div className="pointer-events-none absolute inset-0 bg-gradient-to-b from-solar-primary-light/50 to-solar-bg" />

      <div className="relative z-10 max-w-3xl">
        <h1 className="text-4xl font-extrabold tracking-tight text-solar-text sm:text-5xl md:text-6xl [text-wrap:balance]">
          Energia Limpa
          <br />
          <span className="text-solar-primary">Compartilhada</span>
        </h1>

        <p className="mx-auto mt-6 max-w-xl text-base text-solar-text-muted sm:text-lg md:text-xl [text-wrap:pretty]">
          Conectamos quem gera energia solar com quem quer economizar na conta
          de luz. Sem investimento. 100% legal (ANEEL).
        </p>

        <div className="mt-10 flex flex-col items-center gap-4 sm:flex-row sm:justify-center">
          <Button
            variant="primary"
            size="lg"
            onClick={() => scrollTo('calculadora-consumidor', 'consumidor')}
          >
            Quero Economizar
          </Button>
          <Button
            variant="secondary"
            size="lg"
            onClick={() => scrollTo('calculadora-fornecedor', 'fornecedor')}
          >
            Quero Compartilhar Energia
          </Button>
        </div>
      </div>

      {/* Scroll indicator */}
      <div className="absolute bottom-8 left-1/2 -translate-x-1/2 motion-safe:animate-bounce text-solar-text-muted" aria-hidden="true">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" focusable="false">
          <path d="M7 13l5 5 5-5M7 6l5 5 5-5" />
        </svg>
      </div>
    </section>
  );
}
