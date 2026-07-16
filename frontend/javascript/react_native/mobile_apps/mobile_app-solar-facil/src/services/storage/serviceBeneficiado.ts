import { adicionarItem, carregarLista } from "@/services/storage/storageUtils";
import { BeneficiadoType } from "@/types/BeneficiadoType";

const FILENAME = "beneficiados.json";

export async function salvarBeneficiado(beneficiado: BeneficiadoType): Promise<string> {
  return new Promise((resolve, reject) => {
    console.log("Iniciando o processo de salvar beneficiado:", beneficiado);
    try {
      setTimeout(() => {
        // Carregar lista de beneficiados
        const beneficiados = carregarLista<BeneficiadoType>(FILENAME);
        beneficiados.then((lista) => {
          // Verificar se já existe um beneficiado com o mesmo CPF
          const existente = lista.find(b => b.cpf === beneficiado.cpf);
          if (existente) {
            reject(new Error("Beneficiado com este CPF já existe"));
            return; // Ensure no further execution
          }

          console.log("Adicionando beneficiado:", beneficiado);
          adicionarItem(FILENAME, beneficiado)
            .then(() => {
              console.log("Beneficiado salvo com sucesso:", beneficiado);
              resolve("Beneficiado salvo com sucesso");
            })
            .catch((error) => {
              console.error("Erro ao adicionar beneficiado:", error);
              reject(new Error("Erro ao adicionar beneficiado"));
            });
        }).catch((error) => {
          console.error("Erro ao carregar lista de beneficiados:", error);
          reject(new Error("Erro ao carregar lista de beneficiados"));
        });
      }, 500); // Simulação de delay
    } catch (error) {
      console.error("Erro inesperado:", error);
      reject(new Error("Erro inesperado"));
    }
  });
}

// Retorna uma lista de BeneficiadoType
export async function listarBeneficiados(): Promise<BeneficiadoType[]> {
  console.log("Iniciando o processo de listar beneficiados");
  try {
    const beneficiados = await carregarLista<BeneficiadoType>(FILENAME);
    console.log("Lista de beneficiados carregada com sucesso:", beneficiados.length, "=> ", beneficiados);
    return beneficiados;
  } catch (error) {
    console.error("Erro ao carregar lista de beneficiados:", error);
    throw new Error("Erro ao carregar lista de beneficiados");
  }
}

// Atualiza um beneficiado existente
export async function atualizarBeneficiado(data: BeneficiadoType) {
  const beneficiados = await listarBeneficiados();
  const index = beneficiados.findIndex(b => b.cpf === data.cpf);
  
  if (index !== -1) {
    beneficiados[index] = data;
    localStorage.setItem(FILENAME, JSON.stringify(beneficiados));
  } else {
    throw new Error("Beneficiado não encontrado");
  }
}

// Exclui um beneficiado pelo CPF
export async function excluirBeneficiado(cpf: string) {
  const beneficiados = await listarBeneficiados();
  const index = beneficiados.findIndex(b => b.cpf === cpf);
  
  if (index !== -1) {
    beneficiados.splice(index, 1);
    localStorage.setItem(FILENAME, JSON.stringify(beneficiados));
  } else {
    throw new Error("Beneficiado não encontrado");
  }
}
