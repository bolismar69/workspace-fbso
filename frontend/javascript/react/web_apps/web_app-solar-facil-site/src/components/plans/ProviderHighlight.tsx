import { Zap } from 'lucide-react';
import Button from '@/components/shared/Button';

export default function ProviderHighlight() {
  return (
    <div className="rounded-lg border border-solar-secondary/30 bg-solar-secondary-light p-4 shadow-sm text-center">
      <Zap size={32} className="mx-auto text-solar-secondary" />
      <h3 className="mt-3 text-xl font-bold text-solar-text">
        É um produtor de energia solar?
      </h3>
      <p className="mx-auto mt-3 max-w-md text-solar-text-muted">
        Receba{' '}
        <span className="text-lg font-extrabold text-solar-secondary-dark">
          R$ 0,40 por kWh
        </span>{' '}
        excedente compartilhado. Zero custo inicial. App automatizado. Dentro das normas ANEEL.
      </p>
      <div className="mt-5">
        <Button variant="secondary" size="md" href="/contato?perfil=fornecedor">
          Cadastrar meus painéis →
        </Button>
      </div>
    </div>
  );
}
