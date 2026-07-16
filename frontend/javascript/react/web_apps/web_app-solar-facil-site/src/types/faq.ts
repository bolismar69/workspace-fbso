// Tipos para FAQs — ≡ App (FAQCategoryType, FAQType)
export interface FAQType {
  pergunta: string;
  resposta: string;
}

export interface FAQCategoryType {
  titulo: string;
  faqs: FAQType[];
}
