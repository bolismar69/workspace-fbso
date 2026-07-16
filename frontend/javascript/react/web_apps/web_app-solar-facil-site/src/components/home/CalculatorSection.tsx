import SectionWrapper from '@/components/shared/SectionWrapper';
import ConsumerCalculator from './ConsumerCalculator';
import ProviderCalculator from './ProviderCalculator';

export default function CalculatorSection() {
  return (
    <SectionWrapper id="calculadoras" bg="alt">
      <h2 className="text-center text-3xl font-bold text-solar-text md:text-4xl [text-wrap:balance]">
        Simule em segundos, sem compromisso
      </h2>

      <div className="mt-12 grid gap-8 lg:grid-cols-2">
        <div id="calculadora-consumidor">
          <ConsumerCalculator />
        </div>
        <div id="calculadora-fornecedor">
          <ProviderCalculator />
        </div>
      </div>
    </SectionWrapper>
  );
}
