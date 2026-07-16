import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import Header from '@/components/layout/Header';
import Footer from '@/components/layout/Footer';
import JsonLd from '@/components/shared/JsonLd';
import AnalyticsProvider from '@/components/shared/AnalyticsProvider';
import './globals.css';

const inter = Inter({
  subsets: ['latin'],
  variable: '--font-inter',
});

export const metadata: Metadata = {
  title: {
    default: 'Solar Fácil — Energia Limpa Compartilhada',
    template: '%s',
  },
  description:
    'Conectamos quem gera energia solar com quem quer economizar na conta de luz. Sem investimento. 100% legal (ANEEL).',
  metadataBase: new URL('https://www.solarfacil.com.br'),
  openGraph: {
    title: 'Solar Fácil — Energia Limpa Compartilhada',
    description:
      'Conectamos quem gera energia solar com quem quer economizar na conta de luz.',
    url: 'https://www.solarfacil.com.br',
    siteName: 'Solar Fácil',
    locale: 'pt_BR',
    type: 'website',
  },
  robots: {
    index: true,
    follow: true,
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-BR" className={`${inter.variable} h-full antialiased`}>
      <body className="flex min-h-full flex-col">
        <a
          href="#main-content"
          className="sr-only focus:not-sr-only focus:fixed focus:top-4 focus:left-4 focus:z-[100] focus:rounded-lg focus:bg-solar-primary focus:px-4 focus:py-3 focus:text-white focus:font-bold focus:outline-none focus:ring-2 focus:ring-solar-primary focus:ring-offset-2"
        >
          Pular para o conteúdo
        </a>
        <JsonLd />
        <AnalyticsProvider />
        <Header />
        <main id="main-content" className="flex-1">{children}</main>
        <Footer />
      </body>
    </html>
  );
}
