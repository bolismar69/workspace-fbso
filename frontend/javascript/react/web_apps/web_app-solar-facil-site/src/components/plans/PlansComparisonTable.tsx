import { PLANS } from '@/lib/constants';
import { Star } from 'lucide-react';
import Button from '@/components/shared/Button';

const FEATURE_ROWS: { label: string; get: (p: typeof PLANS[0]) => string }[] = [
  { label: 'Capacidade', get: (p) => p.capacity },
  { label: 'Monitoramento', get: (p) => p.features[0] ?? '—' },
  { label: 'Relatórios', get: (p) => p.features[1] ?? '—' },
  { label: 'Suporte', get: (p) => p.features[2] ?? '—' },
  { label: 'Extras', get: (p) => p.features.slice(3).join(' · ') || '—' },
];

export default function PlansComparisonTable() {
  return (
    <div className="overflow-x-auto">
      {/* Mobile: card stack */}
      <div className="grid gap-6 lg:hidden">
        {PLANS.map((plan) => (
          <div
            key={plan.name}
            className={`rounded-lg border p-4 shadow-sm ${
              plan.highlight
                ? 'border-solar-primary ring-2 ring-solar-primary/30'
                : 'border-solar-border'
            }`}
          >
            <div className="flex items-center justify-between">
              <h3 className="text-lg font-bold text-solar-text">{plan.name}</h3>
              {plan.highlight && (
                <span className="rounded-full bg-solar-primary px-3 py-0.5 text-xs font-bold text-white">
                  <Star size={12} fill="currentColor" className="inline" /> Popular
                </span>
              )}
            </div>
            <p className="mt-2 text-3xl font-extrabold text-solar-text">
              {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(plan.price)}
              <span className="text-sm font-normal text-solar-text-muted">/mês</span>
            </p>
            <p className="mt-1 text-sm text-solar-text-muted">{plan.capacity}</p>
            <ul className="mt-4 space-y-2">
              {plan.features.map((f) => (
                <li key={f} className="flex items-start gap-2 text-sm text-solar-text">
                  <span className="mt-0.5 text-solar-primary">✓</span> {f}
                </li>
              ))}
            </ul>
            <div className="mt-4">
              <Button
                variant={plan.highlight ? 'primary' : 'outline'}
                size="sm"
                href={`/contato?perfil=consumidor&plano=${plan.name}`}
                className="w-full"
              >
                Escolher {plan.name}
              </Button>
            </div>
          </div>
        ))}
      </div>

      {/* Desktop: comparison table */}
      <table className="hidden lg:table w-full border-collapse">
        <thead>
          <tr>
            <th className="p-4 text-left text-sm font-medium text-solar-text-muted w-1/4">Plano</th>
            {PLANS.map((plan) => (
              <th key={plan.name} className={`p-4 text-center ${plan.highlight ? 'bg-solar-primary-light rounded-t-2xl' : ''}`}>
                <div className="text-lg font-bold text-solar-text">{plan.name}</div>
                {plan.highlight && (
                  <span className="mt-1 inline-block rounded-full bg-solar-primary px-3 py-0.5 text-xs font-bold text-white">
                    <Star size={12} fill="currentColor" className="inline" /> Mais Popular
                  </span>
                )}
                <div className="mt-2 text-3xl font-extrabold text-solar-text">
                  {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(plan.price)}
                </div>
                <div className="text-sm text-solar-text-muted">/mês</div>
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          <tr className="border-b border-solar-border">
            <td className="p-4 text-sm font-medium text-solar-text">Capacidade</td>
            {PLANS.map((p) => (
              <td key={p.name} className={`p-4 text-center text-sm text-solar-text ${p.highlight ? 'bg-solar-primary-light/50' : ''}`}>
                {p.capacity}
              </td>
            ))}
          </tr>
          {FEATURE_ROWS.slice(1).map((row) => (
            <tr key={row.label} className="border-b border-solar-border">
              <td className="p-4 text-sm font-medium text-solar-text">{row.label}</td>
              {PLANS.map((p) => (
                <td key={p.name} className={`p-4 text-center text-sm text-solar-text ${p.highlight ? 'bg-solar-primary-light/50' : ''}`}>
                  {row.get(p)}
                </td>
              ))}
            </tr>
          ))}
          <tr>
            <td className="p-4" />
            {PLANS.map((p) => (
              <td key={p.name} className={`p-4 text-center ${p.highlight ? 'bg-solar-primary-light/50 rounded-b-2xl' : ''}`}>
                <Button
                  variant={p.highlight ? 'primary' : 'outline'}
                  size="sm"
                  href={`/contato?perfil=consumidor&plano=${p.name}`}
                >
                  Escolher {p.name}
                </Button>
              </td>
            ))}
          </tr>
        </tbody>
      </table>
    </div>
  );
}
