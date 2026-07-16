// src/services/serviceFAQs.ts
import { FAQCategoryType } from "@/types/FAQType";
import FAQs from "@/mocks/mockFAQs.json";

export async function fetchFAQs(): Promise<FAQCategoryType[]> {
  // Simula fetch com delay artificial
  return new Promise((resolve) => {
    try{
    setTimeout(() => {
      resolve(FAQs as FAQCategoryType[]);
    }, 500);
    } catch (error) {
      console.error("Erro ao buscar dados de FAQs => ", error);
      resolve([]);
    }
  });
}
