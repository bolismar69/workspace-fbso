'use client';

import { useFaqAccordion } from '@/hooks/useFaqAccordion';
import { FAQ_ITEMS } from '@/lib/constants';
import { ChevronDown } from 'lucide-react';

export default function FaqAccordion() {
  const { isOpen, toggle, showContactHint } = useFaqAccordion();

  return (
    <div className="mx-auto max-w-2xl">
      <h2 className="mb-8 text-center text-2xl font-bold text-solar-text">
        Perguntas Frequentes
      </h2>

      <div className="space-y-3">
        {FAQ_ITEMS.map((faq, i) => {
          const open = isOpen(i);
          return (
            <div
              key={i}
              className="overflow-hidden rounded-lg border border-solar-border bg-solar-bg transition-colors"
            >
              <button
                id={`faq-trigger-${i}`}
                onClick={() => toggle(i)}
                aria-expanded={open}
                aria-controls={`faq-panel-${i}`}
                className="flex w-full items-center justify-between px-5 py-4 text-left text-sm font-medium text-solar-text transition-colors hover:bg-solar-bg-alt"
              >
                <span>{faq.question}</span>
                <ChevronDown
                  size={18}
                  className={`flex-shrink-0 text-solar-text-muted transition-transform duration-300 ${
                    open ? 'rotate-180' : ''
                  }`}
                />
              </button>
              <div
                id={`faq-panel-${i}`}
                role="region"
                aria-labelledby={`faq-trigger-${i}`}
                className={`overflow-hidden transition-[max-height] duration-300 ${
                  open ? 'max-h-96' : 'max-h-0'
                }`}
              >
                <p className="px-5 pb-4 text-sm leading-relaxed text-solar-text-muted">
                  {faq.answer}
                </p>
              </div>
            </div>
          );
        })}
      </div>

      {/* Hint after 3+ FAQs without CTA */}
      {showContactHint && (
        <p className="mt-6 text-center text-sm text-solar-text-muted motion-safe:animate-fade-in">
          Não encontrou o que procurava?{' '}
          <a href="/contato" className="font-medium text-solar-primary hover:underline">
            Fale com a gente →
          </a>
        </p>
      )}
    </div>
  );
}
