import { useMemo, useState } from 'react';
import { formatCnpjMasked, isValidCnpj, normalizeCnpj } from './cnpj';

export function App() {
  const [cnpj, setCnpj] = useState('');

  const normalized = useMemo(() => normalizeCnpj(cnpj === '' ? null : cnpj), [cnpj]);
  const valid = useMemo(() => isValidCnpj(cnpj === '' ? null : cnpj), [cnpj]);

  function onChange(next: string) {
    setCnpj(formatCnpjMasked(next));
  }

  const inputOrNull = cnpj === '' ? null : cnpj;

  return (
    <div style={{ fontFamily: 'system-ui, -apple-system, Segoe UI, Roboto, Arial, sans-serif', margin: 24 }}>
      <h1>Validação de CNPJ</h1>

      <label htmlFor="cnpj" style={{ display: 'block', marginBottom: 8, fontWeight: 600 }}>
        CNPJ
      </label>
      <input
        id="cnpj"
        name="cnpj"
        type="text"
        placeholder="99.999.999/9999-99"
        value={cnpj}
        onChange={(e) => onChange(e.target.value)}
        autoComplete="off"
        data-testid="cnpj-input"
        style={{ fontSize: 16, padding: '10px 12px', width: 320, maxWidth: '100%' }}
      />
      <div style={{ color: '#555', marginTop: 6 }}>
        Máscara: <code style={{ background: '#f6f6f6', padding: '2px 6px' }}>99.999.999/9999-99</code>
      </div>

      <div
        data-testid="result"
        style={{ marginTop: 16, padding: 12, border: '1px solid #ddd', width: 520, maxWidth: '100%' }}
      >
        <div style={{ marginTop: 6 }}>
          input:{' '}
          <code data-testid="input-value" style={{ background: '#f6f6f6', padding: '2px 6px' }}>
            {String(inputOrNull)}
          </code>
        </div>
        <div style={{ marginTop: 6 }}>
          normalized:{' '}
          <code data-testid="normalized-value" style={{ background: '#f6f6f6', padding: '2px 6px' }}>
            {String(normalized)}
          </code>
        </div>
        <div style={{ marginTop: 6 }}>
          valid:{' '}
          <span
            data-testid="valid-value"
            style={{ color: valid ? '#0a7a0a' : '#b00020', fontWeight: 700 }}
          >
            {String(valid)}
          </span>
        </div>
        <div style={{ marginTop: 6 }}>
          <span
            data-testid="valid-label"
            style={{ color: valid ? '#0a7a0a' : '#b00020', fontWeight: 700 }}
          >
            {valid ? 'Válido' : 'Inválido'}
          </span>
        </div>
      </div>
    </div>
  );
}
