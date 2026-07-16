import type { ButtonHTMLAttributes, ReactNode } from 'react';

type ButtonVariant = 'primary' | 'secondary' | 'outline';
type ButtonSize = 'sm' | 'md' | 'lg';

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  size?: ButtonSize;
  children: ReactNode;
  href?: string;
}

const variantClasses: Record<ButtonVariant, string> = {
  primary:
    'bg-solar-primary text-white hover:bg-solar-primary-dark focus-visible:ring-solar-primary',
  secondary:
    'bg-solar-secondary text-white hover:bg-solar-secondary-dark focus-visible:ring-solar-secondary',
  outline:
    'border-2 border-solar-primary text-solar-primary bg-transparent hover:bg-solar-primary-light focus-visible:ring-solar-primary',
};

const sizeClasses: Record<ButtonSize, string> = {
  sm: 'px-4 py-2 text-sm rounded-lg',
  md: 'px-6 py-2.5 text-base rounded-lg',
  lg: 'px-6 py-2.5 text-base rounded-lg',
};

export default function Button({
  variant = 'primary',
  size = 'lg',
  children,
  className = '',
  href,
  ...props
}: ButtonProps) {
  const classes = `
    inline-flex items-center justify-center font-bold
    transition-colors duration-200 ease-in-out
    focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2
    cursor-pointer
    ${variantClasses[variant]}
    ${sizeClasses[size]}
    ${className}
  `;

  if (href) {
    return (
      <a href={href} className={classes}>
        {children}
      </a>
    );
  }

  return (
    <button className={classes} {...props}>
      {children}
    </button>
  );
}
