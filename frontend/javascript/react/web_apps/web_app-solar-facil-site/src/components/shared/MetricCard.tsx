interface MetricCardProps {
  value: string;
  label: string;
  highlighted?: boolean;
  highlightColor?: 'green' | 'amber';
  inverted?: boolean;
}

export default function MetricCard({
  value,
  label,
  highlighted = false,
  highlightColor = 'green',
  inverted = false,
}: MetricCardProps) {
  if (inverted) {
    return (
      <div
        className={`
          flex flex-col items-center justify-center rounded-lg p-4 text-center
          transition-colors duration-300
          ${highlighted
            ? 'bg-white/15 ring-2 ring-white/30 scale-105'
            : 'bg-white/5'
          }
        `}
      >
        <span className={`text-3xl md:text-4xl font-extrabold text-white`}>
          {value}
        </span>
        <span className="mt-2 text-sm text-white/70">{label}</span>
      </div>
    );
  }

  const highlightClasses = highlighted
    ? highlightColor === 'green'
      ? 'text-solar-primary font-extrabold'
      : 'text-solar-secondary font-extrabold'
    : 'text-solar-text';

  return (
    <div
      className={`
        flex flex-col items-center justify-center rounded-lg p-4 text-center shadow-sm
        transition-colors duration-300
        ${highlighted ? 'bg-solar-primary-light ring-2 ring-solar-primary/20 scale-105' : 'bg-solar-bg border border-solar-border'}
      `}
    >
      <span className={`text-3xl md:text-4xl ${highlightClasses}`}>{value}</span>
      <span className="mt-2 text-sm text-solar-text-muted">{label}</span>
    </div>
  );
}
