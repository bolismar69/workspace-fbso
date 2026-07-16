export type FAQItemType = {
  pergunta: string;
  resposta: string;
};

export type FAQCategoryType = {
  titulo: string;
  faqs: FAQItemType[];
};
