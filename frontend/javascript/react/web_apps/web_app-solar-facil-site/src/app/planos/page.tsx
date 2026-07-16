import type { Metadata } from 'next';
import SectionWrapper from '@/components/shared/SectionWrapper';
import Breadcrumb from '@/components/shared/Breadcrumb';
import PlansComparisonTable from '@/components/plans/PlansComparisonTable';
import ProviderHighlight from '@/components/plans/ProviderHighlight';
import FaqAccordion from '@/components/plans/FaqAccordion';

export const metadata: Metadata = {
  title: 'Planos — Solar Fácil',
  description:
    'Compare os planos Solar Fácil: Basic (R$150/mês), Special (R$250/mês) e Premium (R$400/mês). Energia limpa compartilhada para cada necessidade.',
};

export default function PlanosPage() {
  return (
    <>
      <SectionWrapper id="planos-page">
        <Breadcrumb items={[{ label: 'Home', href: '/' }, { label: 'Planos' }]} />
        <h1 className="text-center text-3xl font-bold text-solar-text md:text-4xl [text-wrap:balance]">
          Escolha seu plano
        </h1>
        <p className="mt-4 text-center text-solar-text-muted">
          Energia limpa para cada necessidade
        </p>
        <div className="mt-12">
          <PlansComparisonTable />
        </div>
      </SectionWrapper>

      <SectionWrapper id="provider-planos" bg="alt">
        <ProviderHighlight />
      </SectionWrapper>

      <SectionWrapper id="faq">
        <FaqAccordion />
      </SectionWrapper>
    </>
  );
}
