import Link from 'next/link';
import { ChevronRight } from 'lucide-react';

interface BreadcrumbProps {
  items: { label: string; href?: string }[];
}

export default function Breadcrumb({ items }: BreadcrumbProps) {
  return (
    <nav aria-label="Breadcrumb" className="mb-6">
      <ol className="flex items-center gap-1.5 text-sm">
        {items.map((item, i) => {
          const isLast = i === items.length - 1;
          return (
            <li key={i} className="flex items-center gap-1.5">
              {i > 0 && (
                <ChevronRight size={14} className="text-solar-text-muted/50" />
              )}
              {item.href && !isLast ? (
                <Link
                  href={item.href}
                  className="text-solar-text-muted transition-colors hover:text-solar-primary"
                >
                  {item.label}
                </Link>
              ) : (
                <span
                  className={
                    isLast
                      ? 'font-medium text-solar-text'
                      : 'text-solar-text-muted'
                  }
                >
                  {item.label}
                </span>
              )}
            </li>
          );
        })}
      </ol>
    </nav>
  );
}
