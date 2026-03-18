// /frontend/javascript/react/web_apps/web_app-tax-nexus-portal.ts
import { useState } from 'react';

// Interface de envio atualizada para incluir a regra de transição
export interface TaxRequest {
  cnpj: string;
  ncm: string;
  ibge: string;
  saldo_remanescente: number; // Adicionado para suportar o cálculo de 2027+
}

// Interface de resposta atualizada para incluir o Callback (Cadastro Único)
export interface TaxResponse {
  transaction_status: string;
  calculation: {
    municipio: string;
    uf: string;
    ncm: string;
    ano: number;
    // Tributos Legados
    pis: number;
    cofins: number;
    icms: number;
    iss: number;
    ipi: number;
    // Tributos Reforma
    cbs_calculada: number;
    ibs_calculado: number;
    imposto_seletivo: number;
    ipva_novo: number;
    itcmd_novo: number;
    total?: number;
  };
  callback: {
    id_cadastro_unico: string;
    status_integracao: string;
  };
}

export const useTaxService = () => {
  const [loading, setLoading] = useState(false);

  const calculateTax = async (payload: TaxRequest): Promise<TaxResponse | null> => {
    setLoading(true);
    try {
      const response = await fetch('http://localhost:8080/v1/tax/calculate', {
        method: 'POST',
        headers: { 
          'Content-Type': 'application/json',
          'Accept': 'application/json' 
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorData = await response.json();
        console.error("Erro do servidor:", errorData);
        return null;
      }
      
      const data: TaxResponse = await response.json();
      return data;
    } catch (error) {
      console.error("Erro na chamada Fetch:", error);
      return null;
    } finally {
      setLoading(false);
    }
  };

  return { calculateTax, loading };
};