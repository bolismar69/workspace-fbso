import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  // Standalone output for DigitalOcean / non-Vercel hosting
  output: 'standalone',

  // Image optimization — allow solarfacil.com.br domain
  images: {
    formats: ['image/avif', 'image/webp'],
  },

  // Security headers
  async headers() {
    return [
      {
        source: '/(.*)',
        headers: [
          { key: 'X-Frame-Options', value: 'DENY' },
          { key: 'X-Content-Type-Options', value: 'nosniff' },
          { key: 'Referrer-Policy', value: 'strict-origin-when-cross-origin' },
        ],
      },
    ];
  },
};

export default nextConfig;
