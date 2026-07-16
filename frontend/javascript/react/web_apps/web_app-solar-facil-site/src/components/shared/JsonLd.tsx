import { SITE_URL } from '@/lib/constants';

export default function JsonLd() {
  const schema = {
    '@context': 'https://schema.org',
    '@type': 'Organization',
    name: 'Solar Fácil',
    url: SITE_URL,
    description:
      'Plataforma de energia limpa compartilhada. Conectamos produtores de energia solar a consumidores dentro das normas ANEEL.',
    logo: `${SITE_URL}/logo.svg`,
    sameAs: [
      'https://instagram.com/solarfacil',
      'https://linkedin.com/company/solarfacil',
    ],
    contactPoint: {
      '@type': 'ContactPoint',
      contactType: 'customer service',
      email: 'contato@solarfacil.com.br',
      areaServed: 'BR',
      availableLanguage: 'Portuguese',
    },
  };

  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(schema) }}
    />
  );
}
