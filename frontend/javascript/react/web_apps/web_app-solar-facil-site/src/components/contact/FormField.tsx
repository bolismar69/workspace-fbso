import type { ReactNode } from 'react';

interface FormFieldProps {
  label: string;
  name: string;
  type?: 'text' | 'email' | 'tel' | 'textarea' | 'number';
  value: string;
  onChange: (value: string) => void;
  error?: string;
  required?: boolean;
  placeholder?: string;
  children?: ReactNode;
  hidden?: boolean;
}

export default function FormField({
  label,
  name,
  type = 'text',
  value,
  onChange,
  error,
  required = false,
  placeholder,
  children,
  hidden = false,
}: FormFieldProps) {
  if (hidden) {
    return (
      <input
        type="text"
        name={name}
        tabIndex={-1}
        autoComplete="off"
        style={{ position: 'absolute', left: '-9999px' }}
        value={value}
        onChange={(e) => onChange(e.target.value)}
      />
    );
  }

  const inputClasses = `w-full rounded-lg border px-3 py-2.5 text-base transition-colors
    focus:outline-none focus:ring-2 focus:ring-solar-primary
    ${error ? 'border-solar-error-border bg-solar-error-surface' : 'border-solar-border'}`;

  return (
    <div className="mb-4">
      <label htmlFor={name} className="mb-1 block text-base font-bold text-solar-text">
        {label}
        {required && <span className="ml-1 text-red-500">*</span>}
      </label>

      {type === 'textarea' ? (
        <textarea
          id={name}
          name={name}
          rows={3}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          placeholder={placeholder}
          className={inputClasses}
        />
      ) : children ? (
        <div>{children}</div>
      ) : (
        <input
          id={name}
          name={name}
          type={type}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          placeholder={placeholder}
          className={inputClasses}
        />
      )}

      {error && <p className="mt-1 text-sm text-solar-error-text">{error}</p>}
    </div>
  );
}
