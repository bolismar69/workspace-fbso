import SectionWrapper from '@/components/shared/SectionWrapper';
import { ShieldCheck, Zap, Cpu } from 'lucide-react';

export default function DifferentiatorsSection() {
  return (
    <SectionWrapper id="diferenciais">
      <h2 className="text-center text-3xl font-bold text-solar-text md:text-4xl [text-wrap:balance]">
        Por que Solar Fácil?
      </h2>

      {/* Asymmetric layout — each differentiator with its own voice */}
      <div className="mt-14 space-y-10">
        {/* Hero differentiator: R$ 0 Capex — the showstopper */}
        <div className="overflow-hidden rounded-2xl bg-solar-primary-light/60 px-6 py-10 text-center md:px-12">
          <p className="text-6xl font-extrabold tracking-tight text-solar-primary md:text-7xl">
            R$ 0
          </p>
          <p className="mt-3 text-lg font-semibold text-solar-primary-dark">
            Zero Capex — Sem investimento inicial
          </p>
          <p className="mx-auto mt-2 max-w-md text-sm text-solar-text-muted">
            Você começa a economizar sem precisar instalar nada. Zero custo de
            entrada, zero surpresas.
          </p>
        </div>

        {/* Three supporting differentiators — horizontal strip */}
        <div className="grid gap-6 md:grid-cols-3">
          {/* Pareamento Automático */}
          <div className="flex flex-col items-center rounded-xl border border-solar-secondary/30 bg-solar-secondary-light/50 p-6 text-center">
            <div className="flex h-14 w-14 items-center justify-center rounded-full bg-solar-secondary/20">
              <Zap size={26} className="text-solar-secondary-dark" />
            </div>
            <p className="mt-4 text-2xl font-extrabold text-solar-secondary-dark">
              100%
            </p>
            <p className="mt-1 text-sm font-semibold text-solar-text">
              Pareamento Automático
            </p>
            <p className="mt-2 text-xs text-solar-text-muted">
              Nossa plataforma conecta automaticamente produtores e
              consumidores dentro da mesma cooperativa.
            </p>
          </div>

          {/* Plataforma Legal ANEEL */}
          <div className="flex flex-col items-center rounded-xl border border-solar-border bg-solar-bg p-6 text-center shadow-sm">
            <div className="flex h-14 w-14 items-center justify-center rounded-full bg-solar-primary/10">
              <ShieldCheck size={26} className="text-solar-primary" />
            </div>
            <p className="mt-4 text-2xl font-extrabold text-solar-primary">
              100%
            </p>
            <p className="mt-1 text-sm font-semibold text-solar-text">
              Plataforma Legal (ANEEL)
            </p>
            <p className="mt-2 text-xs text-solar-text-muted">
              Operação 100% dentro das normas da ANEEL. Resolução Normativa
              687/2015.
            </p>
          </div>

          {/* API + ANEEL — escalável */}
          <div className="flex flex-col items-center rounded-xl border border-solar-border bg-solar-bg p-6 text-center shadow-sm">
            <div className="flex h-14 w-14 items-center justify-center rounded-full bg-solar-text/5">
              <Cpu size={26} className="text-solar-text" />
            </div>
            <p className="mt-4 text-2xl font-extrabold text-solar-text">
              API +
            </p>
            <p className="mt-1 text-sm font-semibold text-solar-text">
              ANEEL — Escalável
            </p>
            <p className="mt-2 text-xs text-solar-text-muted">
              APIs integradas diretamente à ANEEL. Transparente, escalável,
              pronto para crescer com você.
            </p>
          </div>
        </div>
      </div>
    </SectionWrapper>
  );
}
