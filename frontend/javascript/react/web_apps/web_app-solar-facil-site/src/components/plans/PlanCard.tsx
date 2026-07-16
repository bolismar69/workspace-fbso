import type { Plan } from '@/lib/types';
import { Star } from 'lucide-react';
import Button from '@/components/shared/Button';

interface PlanCardProps {
  plan: Plan;
  variant?: 'compact' | 'full';
}

export default function PlanCard({ plan, variant = 'compact' }: PlanCardProps) {
  const featuresToShow = variant === 'full' ? plan.features : plan.features.slice(0, 3);

  return (
    <div
      className={`
        relative flex flex-col rounded-lg border p-4 shadow-sm transition-shadow hover:shadow-md
        ${plan.highlight
          ? 'border-solar-primary ring-2 ring-solar-primary/30 scale-[1.02]'
          : 'border-solar-border'
        }
      `}
    >
      {/* Popular badge */}
      {plan.highlight && (
        <span className="absolute -top-3 left-1/2 -translate-x-1/2 inline-flex items-center gap-1 rounded-full bg-solar-primary px-4 py-1 text-xs font-bold text-white">
          <Star size={12} fill="currentColor" />
          Mais Popular
        </span>
      )}

      {/* Name */}
      <h3 className="text-xl font-bold text-solar-text">{plan.name}</h3>

      {/* Price */}
      <div className="mt-3">
        <span className="text-3xl font-extrabold text-solar-text">
          {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(plan.price)}
        </span>
        <span className="text-sm text-solar-text-muted">/mês</span>
      </div>

      {/* Capacity */}
      <p className="mt-1 text-sm text-solar-text-muted">
        Capacidade: {plan.capacity}
      </p>

      {/* Features */}
      <ul className="mt-4 flex-1 space-y-2">
        {featuresToShow.map((feature) => (
          <li key={feature} className="flex items-start gap-2 text-sm text-solar-text">
            <span className="mt-0.5 flex-shrink-0 text-solar-primary">✓</span>
            {feature}
          </li>
        ))}
      </ul>

      {/* CTA */}
      <div className="mt-6">
        <Button
          variant={plan.highlight ? 'primary' : 'outline'}
          size="md"
          href={`/contato?perfil=consumidor&plano=${plan.name}`}
          className="w-full"
        >
          Escolher {plan.name}
        </Button>
      </div>
    </div>
  );
}
