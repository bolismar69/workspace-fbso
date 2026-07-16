'use client';

import { useState, useCallback } from 'react';

export function useFaqAccordion() {
  const [openIndex, setOpenIndex] = useState<number | null>(null);
  const [totalOpened, setTotalOpened] = useState<number[]>([]);
  const [showContactHint, setShowContactHint] = useState(false);
  void totalOpened; // Used in toggle logic below via setTotalOpened

  const TOGGLE_THRESHOLD = 3;

  const toggle = useCallback((index: number) => {
    setOpenIndex((prev) => (prev === index ? null : index));
    setTotalOpened((prev) => {
      if (!prev.includes(index)) {
        const updated = [...prev, index];
        // Show hint after 3+ unique FAQs opened without CTA
        if (updated.length > TOGGLE_THRESHOLD) {
          setShowContactHint(true);
        }
        return updated;
      }
      return prev;
    });
  }, []);

  const isOpen = useCallback(
    (index: number) => openIndex === index,
    [openIndex],
  );

  return { openIndex, isOpen, toggle, showContactHint };
}
