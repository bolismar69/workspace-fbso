'use client';

import { useState } from 'react';
import { usePathname } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import MobileMenu from './MobileMenu';
import { Menu } from 'lucide-react';

export default function Header() {
  const [mobileOpen, setMobileOpen] = useState(false);
  const pathname = usePathname();

  const linkClasses = (href: string) =>
    `text-sm font-medium transition-colors hover:text-solar-text ${
      pathname === href
        ? 'text-solar-primary font-semibold'
        : 'text-solar-text-muted'
    }`;

  return (
    <>
      <header className="fixed left-0 right-0 top-0 z-50 border-b border-solar-border bg-white/80 backdrop-blur-md">
        <div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-4">
          {/* Logo */}
          <Link href="/" className="flex items-center gap-2">
            <Image
              src="/logo.svg"
              alt="Solar Fácil"
              width={120}
              height={24}
              priority
            />
          </Link>

          {/* Desktop Nav */}
          <nav className="hidden items-center gap-6 md:flex">
            <Link href="/planos" className={linkClasses('/planos')}>
              Planos
            </Link>
            <Link href="/contato" className={linkClasses('/contato')}>
              Contato
            </Link>
            <a
              href="#"
              className="rounded-lg bg-solar-primary px-4 py-2 text-sm font-semibold text-white transition-colors hover:bg-solar-primary-dark"
            >
              Baixar App
            </a>
          </nav>

          {/* Mobile hamburger */}
          <button
            onClick={() => setMobileOpen(true)}
            className="rounded-lg p-3 text-solar-text transition-colors hover:bg-solar-bg-alt md:hidden"
            aria-label="Abrir menu"
          >
            <Menu size={24} />
          </button>
        </div>
      </header>

      {/* Spacer — compensates fixed header height */}
      <div className="h-16" />

      <MobileMenu open={mobileOpen} onClose={() => setMobileOpen(false)} />
    </>
  );
}
