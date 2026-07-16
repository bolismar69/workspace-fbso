'use client';

import { useEffect } from 'react';
import { X } from 'lucide-react';

interface MobileMenuProps {
  open: boolean;
  onClose: () => void;
}

export default function MobileMenu({ open, onClose }: MobileMenuProps) {
  // Close on Escape key
  useEffect(() => {
    const handleKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose();
    };
    if (open) document.addEventListener('keydown', handleKey);
    return () => document.removeEventListener('keydown', handleKey);
  }, [open, onClose]);

  // Prevent body scroll when open
  useEffect(() => {
    if (open) {
      document.body.style.overflow = 'hidden';
    } else {
      document.body.style.overflow = '';
    }
    return () => { document.body.style.overflow = ''; };
  }, [open]);

  if (!open) return null;

  const links = [
    { href: '/', label: 'Home' },
    { href: '/planos', label: 'Planos' },
    { href: '/contato', label: 'Contato' },
  ];

  return (
    <div className="fixed inset-0 z-50 md:hidden">
      {/* Backdrop */}
      <div
        className="absolute inset-0 bg-black/50 backdrop-blur-sm"
        onClick={onClose}
      />

      {/* Panel */}
      <div className="absolute right-0 top-0 h-full w-64 motion-safe:animate-slide-in bg-white shadow-xl">
        <div className="flex items-center justify-between border-b border-solar-border p-4">
          <span className="font-bold text-solar-text">Menu</span>
          <button
            onClick={onClose}
            className="rounded-lg p-1 text-solar-text-muted hover:bg-solar-bg-alt"
            aria-label="Fechar menu"
          >
            <X size={24} />
          </button>
        </div>

        <nav className="flex flex-col p-4">
          {links.map((link) => (
            <a
              key={link.href}
              href={link.href}
              onClick={onClose}
              className="rounded-lg px-4 py-3 text-lg font-medium text-solar-text transition-colors hover:bg-solar-bg-alt"
            >
              {link.label}
            </a>
          ))}
          <a
            href="#"
            onClick={onClose}
            className="mt-4 rounded-lg bg-solar-primary px-4 py-3 text-center text-lg font-semibold text-white"
          >
            Baixar App
          </a>
        </nav>
      </div>
    </div>
  );
}
