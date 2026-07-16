import SectionWrapper from '@/components/shared/SectionWrapper';
import { HOW_IT_WORKS_STEPS } from '@/lib/constants';
import { Sun, RefreshCw, Home } from 'lucide-react';

const iconMap: Record<string, typeof Sun> = {
  Sun,
  RefreshCw,
  Home,
};

export default function HowItWorksSection() {
  return (
    <SectionWrapper id="como-funciona" bg="alt">
      <h2 className="text-center text-3xl font-bold text-solar-text md:text-4xl [text-wrap:balance]">
        Como funciona a energia compartilhada?
      </h2>

      {/* Timeline */}
      <div className="mt-14">
        {/* Desktop: horizontal timeline */}
        <div className="relative hidden md:flex md:items-start md:justify-center md:gap-0">
          {/* Connecting line */}
          <div className="absolute left-[calc(16.67%+32px)] right-[calc(16.67%+32px)] top-8 h-0.5 bg-solar-primary/30" />

          {HOW_IT_WORKS_STEPS.map((step, i) => {
            const Icon = iconMap[step.icon] || Sun;
            return (
              <div
                key={i}
                className="relative flex flex-col items-center text-center"
                style={{ width: '33.333%' }}
              >
                {/* Circle with icon */}
                <div className="relative z-10 flex h-16 w-16 items-center justify-center rounded-full bg-solar-primary text-white shadow-lg ring-4 ring-solar-bg-alt">
                  <Icon size={28} />
                </div>

                {/* Step number */}
                <span className="mt-4 text-xs font-bold uppercase tracking-wider text-solar-primary/60">
                  Passo {i + 1}
                </span>

                {/* Content */}
                <h3 className="mt-2 text-lg font-semibold text-solar-text">
                  {step.title}
                </h3>
                <p className="mt-2 max-w-[240px] text-sm leading-relaxed text-solar-text-muted">
                  {step.description}
                </p>
              </div>
            );
          })}
        </div>

        {/* Mobile: vertical timeline */}
        <div className="relative flex flex-col gap-0 md:hidden">
          {HOW_IT_WORKS_STEPS.map((step, i) => {
            const Icon = iconMap[step.icon] || Sun;
            const isLast = i === HOW_IT_WORKS_STEPS.length - 1;

            return (
              <div key={i} className="relative flex gap-4 pb-8">
                {/* Vertical line + circle column */}
                <div className="relative flex flex-col items-center">
                  {/* Connecting line (above circle, except first) */}
                  {i > 0 && (
                    <div className="h-6 w-0.5 bg-solar-primary/30" />
                  )}
                  {/* Circle */}
                  <div className="relative z-10 flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-solar-primary text-white shadow-md ring-4 ring-solar-bg-alt">
                    <Icon size={22} />
                  </div>
                  {/* Connecting line (below circle, except last) */}
                  {!isLast && (
                    <div className="h-full w-0.5 bg-solar-primary/30" />
                  )}
                </div>

                {/* Content */}
                <div className="flex-1 pt-1">
                  <span className="text-xs font-bold uppercase tracking-wider text-solar-primary/60">
                    Passo {i + 1}
                  </span>
                  <h3 className="mt-1 text-lg font-semibold text-solar-text">
                    {step.title}
                  </h3>
                  <p className="mt-1 text-sm leading-relaxed text-solar-text-muted">
                    {step.description}
                  </p>
                </div>
              </div>
            );
          })}
        </div>
      </div>

      <p className="mt-10 text-center">
        <a
          href="/planos"
          className="text-sm font-medium text-solar-primary hover:underline"
        >
          Saiba mais: veja os planos disponíveis →
        </a>
      </p>
    </SectionWrapper>
  );
}
