import Link from 'next/link';
import Image from 'next/image';
import { ShieldCheck } from 'lucide-react';
import {
  CONTACT_EMAIL,
  APP_STORE_URL,
  GOOGLE_PLAY_URL,
  INSTAGRAM_URL,
} from '@/lib/constants';

export default function Footer() {
  return (
    <footer className="border-t border-solar-border bg-solar-bg-alt">
      <div className="mx-auto max-w-7xl px-4 py-12">
        <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-4">
          {/* Brand */}
          <div>
            <Image src="/logo.svg" alt="Solar Fácil" width={120} height={24} />
            <p className="mt-3 text-sm text-solar-text-muted">
              Energia limpa compartilhada. Conectando produtores e consumidores dentro das normas ANEEL.
            </p>
          </div>

          {/* Links */}
          <div>
            <h4 className="mb-3 text-sm font-semibold text-solar-text">Navegação</h4>
            <ul className="space-y-2 text-sm text-solar-text-muted">
              <li><Link href="/" className="hover:text-solar-primary transition-colors">Home</Link></li>
              <li><a href="/planos" className="hover:text-solar-primary transition-colors">Planos</a></li>
              <li><a href="/contato" className="hover:text-solar-primary transition-colors">Contato</a></li>
            </ul>
          </div>

          {/* App Stores */}
          <div>
            <h4 className="mb-3 text-sm font-semibold text-solar-text">Baixe o App</h4>
            <ul className="space-y-2 text-sm text-solar-text-muted">
              <li><a href={APP_STORE_URL} className="hover:text-solar-primary transition-colors" target="_blank" rel="noopener">App Store</a></li>
              <li><a href={GOOGLE_PLAY_URL} className="hover:text-solar-primary transition-colors" target="_blank" rel="noopener">Google Play</a></li>
            </ul>
          </div>

          {/* Contact + Social */}
          <div>
            <h4 className="mb-3 text-sm font-semibold text-solar-text">Contato</h4>
            <ul className="space-y-2 text-sm text-solar-text-muted">
              <li><a href={`mailto:${CONTACT_EMAIL}`} className="hover:text-solar-primary transition-colors">{CONTACT_EMAIL}</a></li>
              <li><a href={INSTAGRAM_URL} className="hover:text-solar-primary transition-colors" target="_blank" rel="noopener">Instagram</a></li>
              <li><a href="https://linkedin.com/company/solarfacil" className="hover:text-solar-primary transition-colors" target="_blank" rel="noopener">LinkedIn</a></li>
            </ul>
          </div>
        </div>

        {/* ANEEL Seal + Legal */}
        <div className="mt-10 border-t border-solar-border pt-6">
          <div className="flex flex-col items-center gap-4 sm:flex-row sm:justify-between">
            <div className="flex items-center gap-2 text-xs text-solar-text-muted">
              <ShieldCheck className="h-4 w-4 text-solar-primary" />
              Plataforma regulamentada pela ANEEL (RN 687/2015)
            </div>
            <div className="flex gap-4 text-xs text-solar-text-muted">
              <a href="/termos" className="hover:text-solar-primary transition-colors">Termos de Uso</a>
              <a href="/privacidade" className="hover:text-solar-primary transition-colors">Privacidade</a>
              <span>© {new Date().getFullYear()} Solar Fácil</span>
            </div>
          </div>
        </div>
      </div>
    </footer>
  );
}
