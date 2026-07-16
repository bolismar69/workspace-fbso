import { Suspense } from 'react';
import HeroSection from '@/components/home/HeroSection';
import CalculatorSection from '@/components/home/CalculatorSection';
import ProofSection from '@/components/home/ProofSection';
import PlansSection from '@/components/home/PlansSection';
import HowItWorksSection from '@/components/home/HowItWorksSection';
import DifferentiatorsSection from '@/components/home/DifferentiatorsSection';
import FinalCtaSection from '@/components/home/FinalCtaSection';
import Skeleton from '@/components/shared/Skeleton';

function ProofSectionFallback() {
  return (
    <section className="px-4 py-20 md:py-28 bg-solar-primary">
      <div className="mx-auto max-w-5xl text-center">
        <h2 className="text-3xl font-bold text-white md:text-4xl">
          Quem já confia na Solar Fácil
        </h2>
        <div className="mt-12 grid gap-6 sm:grid-cols-3">
          {[1, 2, 3].map((i) => (
            <Skeleton key={i} variant="card" />
          ))}
        </div>
      </div>
    </section>
  );
}

export default function HomePage() {
  return (
    <>
      <HeroSection />
      <CalculatorSection />
      <Suspense fallback={<ProofSectionFallback />}>
        <ProofSection />
      </Suspense>
      <HowItWorksSection />
      <PlansSection />
      <DifferentiatorsSection />
      <FinalCtaSection />
    </>
  );
}
