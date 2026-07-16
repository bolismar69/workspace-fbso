import SectionWrapper from '@/components/shared/SectionWrapper';
import PlanCard from '@/components/plans/PlanCard';
import Button from '@/components/shared/Button';
import { PLANS } from '@/lib/constants';

export default function PlansSection() {
  return (
    <SectionWrapper id="planos" bg="alt">
      <h2 className="text-center text-3xl font-bold text-solar-text md:text-4xl [text-wrap:balance]">
        Planos para cada necessidade
      </h2>

      {/* Plan cards */}
      <div className="mt-12 grid gap-6 lg:grid-cols-3">
        {PLANS.map((plan) => (
          <PlanCard key={plan.name} plan={plan} variant="compact" />
        ))}
      </div>

      {/* Provider highlight */}
      <div className="mx-auto mt-10 max-w-lg rounded-lg border border-solar-secondary/30 bg-solar-secondary-light p-4 shadow-sm text-center">
        <p className="text-lg font-bold text-solar-secondary-dark">
          É um produtor de energia solar?
        </p>
        <p className="mt-2 text-sm text-solar-text-muted">
          Receba <span className="font-semibold text-solar-secondary">R$ 0,40 por kWh</span> excedente.
          Sem custo inicial. App automatizado. Dentro das normas ANEEL.
        </p>
        <div className="mt-4">
          <Button variant="secondary" size="sm" href="/contato?perfil=fornecedor">
            Cadastrar meus painéis →
          </Button>
        </div>
      </div>

      <p className="mt-6 text-center">
        <a href="/planos" className="text-sm font-medium text-solar-primary hover:underline">
          Comparar todos os planos em detalhes →
        </a>
      </p>
    </SectionWrapper>
  );
}
