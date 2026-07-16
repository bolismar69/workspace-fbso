import type { ReactNode } from 'react';

interface SectionWrapperProps {
  id?: string;
  children: ReactNode;
  className?: string;
  bg?: 'white' | 'alt';
}

export default function SectionWrapper({
  id,
  children,
  className = '',
  bg = 'white',
}: SectionWrapperProps) {
  return (
    <section
      id={id}
      className={`
        px-4 py-16 md:py-24
        ${bg === 'alt' ? 'bg-solar-bg-alt' : 'bg-solar-bg'}
        ${className}
      `}
    >
      <div className="mx-auto max-w-7xl">{children}</div>
    </section>
  );
}
