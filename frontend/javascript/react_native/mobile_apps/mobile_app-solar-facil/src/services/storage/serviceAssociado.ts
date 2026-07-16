// /src/services/storage/serviceAssociado.ts
import { adicionarItem, carregarLista, salvarListaV2 } from "@/services/storage/storageUtils";
import { AssociadoType } from "@/types/AssociadoType";

const FILENAME = "associados.json";

export async function salvarAssociado(associado: AssociadoType): Promise<string> {
  console.log("=== INICIO ========================================================================");
  return new Promise((resolve, reject) => {
    console.log("Iniciando o processo de salvar associado:", associado);
    try {
      setTimeout(() => {
        // Carregar lista de associados
        const associados = carregarLista<AssociadoType>(FILENAME);
        associados.then((lista) => {
          // Verificar se já existe um associado com o mesmo CPF
          const existente = lista.find(
            (b) =>  (b.cpf_cnpj === associado.cpf_cnpj));
          if (existente) {
            console.error("Associado já existe com o CPF/CNPJ:", associado.cpf_cnpj);
            console.log("=== TERMINO =======================================================================");
            reject(new Error("Associado já existe"));
            return; // Ensure no further execution
          }

          console.log("Adicionando associado:", associado);
          adicionarItem(FILENAME, associado)
            .then(() => {
              console.log("Associado salvo com sucesso:", associado);
              console.log("=== TERMINO =======================================================================");
              resolve("Associado salvo com sucesso");
            })
            .catch((error) => {
              console.error("Erro ao adicionar associado:", error);
              console.log("=== TERMINO =======================================================================");
              reject(new Error("Erro ao adicionar associado"));
            });
        }).catch((error) => {
          console.error("Erro ao carregar lista de associados:", error);
          console.log("=== TERMINO =======================================================================");
          reject(new Error("Erro ao carregar lista de associados"));
        });
      }, 500); // Simulação de delay
    } catch (error) {
      console.error("Erro inesperado:", error);
      console.log("=== TERMINO =======================================================================");
      reject(new Error("Erro inesperado"));
    }
  });
}

export async function listarAssociados(): Promise<AssociadoType[]> {
  console.log("=== INICIO ========================================================================");
  console.log("Iniciando o processo de listar associados");
  try {
    const associados = await carregarLista<AssociadoType>(FILENAME);
    console.log("Lista de associados carregada com sucesso:", associados.length, "=> ", associados);
    console.log("=== TERMINO ======================================================================="); 
    return associados;
  } catch (error) {
    console.error("Erro ao carregar lista de associados:", error);
    console.log("=== TERMINO =======================================================================");
    throw new Error("Erro ao carregar lista de associados");
  }
}

export async function excluirAssociado(associado: AssociadoType) {
  console.log("=== INICIO ========================================================================");
  const associados = await listarAssociados();
  const index = associados.findIndex(
    (b) =>  (b.cpf_cnpj === associado.cpf_cnpj ) );
  if (index !== -1) {
    associados.splice(index, 1); // Remove o associado da lista
    adicionarItem(FILENAME, associados)
      .then(() => {
        console.log("Associado excluído com sucesso");
        console.log("=== TERMINO =======================================================================");
      })
      .catch((error) => {
        console.error("Erro ao excluir associado:", error);
        console.log("=== TERMINO =======================================================================");
      });
  } 
  else {
    console.log("=== TERMINO =======================================================================");
    throw new Error("Beneficiado não encontrado");
  }
}
export async function buscarAssociadoPorCpfCnpj(cpf_cnpj: string): Promise<AssociadoType[]> {
  console.log("=== INICIO ========================================================================");
  console.log("Iniciando o processo de busca de associado por CPF:", cpf_cnpj);
  try {
    const associados = await carregarLista<AssociadoType>(FILENAME);
    const associado = associados.find((b) => b.cpf_cnpj === cpf_cnpj);
    if (associado) {
      console.log("Associado encontrado com o CPF/CNPJ:", cpf_cnpj);
      console.log("=== TERMINO =======================================================================");
      return [associado]; // Retorna o associado encontrado
    }
    // If not found, return an empty array
    console.log("Associado não encontrado com o CPF:", cpf_cnpj);
    console.log("=== TERMINO =======================================================================");
    return [];
  } catch (error) {
    console.error("Erro ao buscar associado por CPF:", error);
    console.log("=== TERMINO =======================================================================");
    throw new Error("Erro ao buscar associado por CPF");
  }
}
export async function buscarAssociadoPorId(id: string): Promise<AssociadoType[]> {
  console.log("=== INICIO ========================================================================");
  console.log("Iniciando o processo de busca de associado por ID:", id);
  try {
    const associados = await carregarLista<AssociadoType>(FILENAME);
    const associado = associados.find((b) => b.id === id);
    if (associado) {
      console.log("Associado encontrado com o ID:", id);
      console.log("=== TERMINO =======================================================================");
      return [associado]; // Retorna o associado encontrado
    }
    // If not found, return null
    console.log("Associado não encontrado com o ID:", id);
    console.log("=== TERMINO =======================================================================");
    return [];
  } catch (error) {
    console.error("Erro ao buscar associado por ID:", error);
    console.log("=== TERMINO =======================================================================");
    throw new Error("Erro ao buscar associado por ID");
  }
}
export async function buscarAssociadoPorCpfCnpjSenha(cpf_cnpj: string, senha: string): Promise<AssociadoType[]> {
  console.log("=== INICIO ========================================================================");
  console.log("Iniciando o processo de busca de associado por CPF:", cpf_cnpj);
  try {
    const associados = await carregarLista<AssociadoType>(FILENAME);
    const associado = associados.find((b) => b.cpf_cnpj === cpf_cnpj && b.senha === senha);
    if (associado) {
      console.log("Associado encontrado com o CPF/CNPJ:", cpf_cnpj);
      console.log("=== TERMINO =======================================================================");
      return [associado]; // Retorna o associado encontrado
    }
    // If not found, return an empty array
    console.log("Associado não encontrado com o CPF:", cpf_cnpj);
    console.log("=== TERMINO =======================================================================");
    return [];
  } catch (error) {
    console.error("Erro ao buscar associado por CPF:", error);
    console.log("=== TERMINO =======================================================================");
    throw new Error("Erro ao buscar associado por CPF");
  }
}

export async function atualizarAssociado(associado: AssociadoType): Promise<string> {
  console.log("=== INICIO ========================================================================");
  console.log("ATUALIZAR-ASSOCIADO: Iniciando o processo de atualizar associado:", associado);
  return new Promise((resolve, reject) => {
    try {
      console.log("ATUALIZAR-ASSOCIADO: Aguardando 500ms antes de atualizar associado...");
      setTimeout(() => {
        console.log("ATUALIZAR-ASSOCIADO: Carregando lista de associados para atualização... => ", FILENAME);
        carregarLista<AssociadoType>(FILENAME).then((lista) => {
          console.log("ATUALIZAR-ASSOCIADO: Lista de associados carregada com sucesso:", lista.length, "=> ", lista);
          const index = lista.findIndex((b) => b.cpf_cnpj === associado.cpf_cnpj);
          if (index === -1) {
            console.error("ATUALIZAR-ASSOCIADO: Associado não encontrado para atualização:", associado.cpf_cnpj);
            console.log("=== TERMINO =======================================================================");
            reject(new Error("Associado não encontrado"));
            return;
          }

          console.log("ATUALIZAR-ASSOCIADO: Atualizando associado:", index, " => ", associado);
          lista[index] = associado; // Atualiza o item
          console.log("ATUALIZAR-ASSOCIADO: Salvando lista - completa - atualizada de associados...");
          salvarListaV2(FILENAME, lista)
            .then(() => {
              console.log("ATUALIZAR-ASSOCIADO: Associado atualizado com sucesso:", associado);
              console.log("=== TERMINO =======================================================================");
              resolve("Associado atualizado com sucesso");
            })
            .catch((error) => {
              console.error("ATUALIZAR-ASSOCIADO: Erro ao salvar lista:", error);
              console.log("=== TERMINO =======================================================================");
              reject(new Error("Erro ao atualizar associado"));
            });
        }).catch((error) => {
          console.error("ATUALIZAR-ASSOCIADO: Erro ao carregar lista de associados:", error);
          console.log("=== TERMINO =======================================================================");
          reject(new Error("Erro ao carregar lista"));
        });
      }, 0);
    } catch (error) {
      console.error("ATUALIZAR-ASSOCIADO: Erro inesperado:", error);
      console.log("=== TERMINO =======================================================================");
      reject(new Error("Erro inesperado"));
    }
  });
}
