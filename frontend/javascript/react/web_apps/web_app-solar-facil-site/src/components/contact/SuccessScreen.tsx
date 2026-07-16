import { CheckCircle } from 'lucide-react';

export default function SuccessScreen() {
  return (
    <div className="flex flex-col items-center justify-center py-12 text-center">
      <CheckCircle size={64} className="text-solar-primary" />
      <h2 className="mt-6 text-2xl font-bold text-solar-text">
        Obrigado!
      </h2>
      <p className="mt-3 max-w-sm text-solar-text-muted">
        Entraremos em contato em até 24h. Enquanto isso, que tal baixar nosso app?
      </p>
      <a
        href="#"
        className="mt-6 inline-flex items-center justify-center rounded-lg bg-solar-primary px-6 py-2.5 font-bold text-white transition-colors hover:bg-solar-primary-dark"
      >
        Baixar App Solar Fácil
      </a>
    </div>
  );
}
