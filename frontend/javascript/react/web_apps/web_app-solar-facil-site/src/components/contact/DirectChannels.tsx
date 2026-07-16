import { MessageCircle, Mail, Globe, MapPin } from 'lucide-react';
import { WHATSAPP_NUMBER, CONTACT_EMAIL, INSTAGRAM_URL } from '@/lib/constants';

const channels = [
  {
    icon: MessageCircle,
    label: 'WhatsApp',
    info: '(11) 9xxxx-xxxx',
    action: `https://wa.me/${WHATSAPP_NUMBER}`,
    actionLabel: 'Abrir conversa',
    color: 'text-solar-primary',
  },
  {
    icon: Mail,
    label: 'E-mail',
    info: CONTACT_EMAIL,
    action: `mailto:${CONTACT_EMAIL}`,
    actionLabel: 'Enviar e-mail',
    color: 'text-solar-primary',
  },
  {
    icon: Globe,
    label: 'Redes Sociais',
    info: '@solarfacil',
    action: INSTAGRAM_URL,
    actionLabel: 'Seguir no Instagram',
    color: 'text-solar-secondary',
  },
];

export default function DirectChannels() {
  return (
    <div>
      <h2 className="mb-6 text-center text-xl font-bold text-solar-text">
        Ou fale direto com a gente
      </h2>
      <div className="grid gap-4 sm:grid-cols-3">
        {channels.map((ch) => {
          const Icon = ch.icon;
          return (
            <a
              key={ch.label}
              href={ch.action}
              target="_blank"
              rel="noopener noreferrer"
              className="flex flex-col items-center rounded-lg border border-solar-border p-4 shadow-sm text-center transition-shadow hover:shadow-md"
            >
              <Icon size={32} className={ch.color} />
              <h3 className="mt-3 font-semibold text-solar-text">{ch.label}</h3>
              <p className="mt-1 text-sm text-solar-text-muted">{ch.info}</p>
              <span className="mt-3 text-sm font-medium text-solar-primary">
                {ch.actionLabel} →
              </span>
            </a>
          );
        })}
      </div>
      <p className="mt-6 flex items-center justify-center gap-1.5 text-center text-sm text-solar-text-muted">
        <MapPin size={14} className="text-solar-primary" />
        São Paulo — SP
      </p>
    </div>
  );
}
