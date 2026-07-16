'use client';

import Button from '@/components/shared/Button';

export default function FinalCtaSection() {
  const scrollTo = (id: string) => {
    document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' });
  };

  return (
    <section id="cta-final" className="bg-solar-bg px-4 py-20 md:py-28">
      <div className="mx-auto max-w-xl text-center">
        <h2 className="text-3xl font-bold text-solar-text md:text-4xl [text-wrap:balance]">
          Pronto para fazer parte?
        </h2>
        <p className="mx-auto mt-4 max-w-md text-solar-text-muted [text-wrap:pretty]">
          Junte-se a 500+ pessoas que já estão economizando com energia limpa.
        </p>

        <div className="mt-10 flex flex-col items-center gap-4 sm:flex-row sm:justify-center">
          <Button
            variant="primary"
            size="lg"
            onClick={() => scrollTo('calculadora-consumidor')}
          >
            Quero Economizar
          </Button>
          <Button
            variant="secondary"
            size="lg"
            onClick={() => scrollTo('calculadora-fornecedor')}
          >
            Quero Compartilhar Energia
          </Button>
        </div>
      </div>
    </section>
  );
}
