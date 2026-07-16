interface SkeletonProps {
  variant?: 'text' | 'card' | 'circle' | 'input';
  className?: string;
}

export default function Skeleton({ variant = 'text', className = '' }: SkeletonProps) {
  const base = `rounded-lg bg-solar-text-muted/10 motion-safe:animate-pulse ${className}`;

  switch (variant) {
    case 'card':
      return (
        <div className={`${base} flex flex-col gap-3 p-4`} aria-hidden="true">
          <div className="h-4 w-2/3 rounded bg-solar-text-muted/8" />
          <div className="h-3 w-5/6 rounded bg-solar-text-muted/8" />
          <div className="h-3 w-3/4 rounded bg-solar-text-muted/8" />
        </div>
      );
    case 'input':
      return <div className={`${base} h-12`} aria-hidden="true" />;
    case 'circle':
      return <div className={`${base} rounded-full`} aria-hidden="true" style={{ width: 64, height: 64 }} />;
    case 'text':
    default:
      return <div className={`${base} h-4`} aria-hidden="true" />;
  }
}
