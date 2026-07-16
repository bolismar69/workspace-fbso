// src/services/storage/storageUtils.ts
import * as FileSystem from "expo-file-system";
import AsyncStorage from "@react-native-async-storage/async-storage";

const DOC_DIR = (FileSystem.documentDirectory || "");
const BASE_DIR = FileSystem.documentDirectory + "storage/";

/**
 * Garante que o diretório base exista
 */
async function ensureDirExists() {
  console.log("=== INICIO ========================================================================");
  console.log("ensureDirExists - Verificando se o diretório base existe...", BASE_DIR);
  const dirInfo = await FileSystem.getInfoAsync(BASE_DIR);
  if (!dirInfo.exists) {
    console.log("ensureDirExists - Diretório base não encontrado. Criando diretório...", BASE_DIR);
    await FileSystem.makeDirectoryAsync(BASE_DIR, { intermediates: true });
    console.log("ensureDirExists - Diretório base criado com sucesso.", BASE_DIR);
  }
  console.log("=== TERMINO =======================================================================");
}

// Apaga o diretório base e recria, garantindo que o armazenamento esteja limpo
export async function resetStorage(): Promise<void> {
  console.log("=== INICIO ========================================================================");
  console.log("resetStorage - Iniciando o processo de reset do armazenamento...");
  try {
    console.log("resetStorage - Verificando se o diretório base existe...", BASE_DIR);
    const dirInfo = await FileSystem.getInfoAsync(BASE_DIR);
    if (dirInfo.exists) {
      console.log("resetStorage - Diretório base encontrado. Limpando armazenamento...", BASE_DIR);
      await FileSystem.deleteAsync(BASE_DIR, { idempotent: true });
    }
    console.log("resetStorage - Criando diretório base novamente...", BASE_DIR);
    await ensureDirExists();
    console.log("resetStorage - Armazenamento resetado com sucesso.", BASE_DIR);
  } catch (error) {
    console.error("resetStorage - Erro ao resetar armazenamento:", error);
  }
  console.log("=== TERMINO =======================================================================");
}

export async function clearStorageDocumentDirectory(): Promise<void> {
  console.log("=== INICIO ========================================================================");
  console.log("clearStorageDocumentDirectory - Iniciando o processo de reset do armazenamento...", " - DOC_DIR => ", DOC_DIR);
  try {
    console.log("clearStorageDocumentDirectory - Verificando se o diretório base existe...");
    const dirInfo = await FileSystem.getInfoAsync(DOC_DIR);
    if (dirInfo.exists) {

      await listarArquivosDocDir();

      console.log("clearStorageDocumentDirectory - Diretório base encontrado. Limpando armazenamento...");
      await FileSystem.deleteAsync(DOC_DIR, { idempotent: true });
      console.log("clearStorageDocumentDirectory - Armazenamento limpo com sucesso.");

      await listarArquivosDocDir();

    }
  } catch (error) {
    console.error("clearStorageDocumentDirectory - Erro ao resetar armazenamento:", error);
  }
  console.log("=== TERMINO =======================================================================");
}

// Funções de armazenamento
export async function limparArmazenamento(): Promise<void> {
  console.log("=== INICIO ========================================================================");
  console.log("limparArmazenamento - Iniciando o processo de limpeza do armazenamento...");
  try {
    console.log("limparArmazenamento - Verificando se o diretório base existe...", BASE_DIR);
    const dirInfo = await FileSystem.getInfoAsync(BASE_DIR);
    if (dirInfo.exists) {

      await listarArquivosBaseDir();

      console.log("limparArmazenamento - Diretório base encontrado. Limpando armazenamento...", BASE_DIR);
      await FileSystem.deleteAsync(BASE_DIR, { idempotent: true });
      console.log("limparArmazenamento - Armazenamento limpo com sucesso.", BASE_DIR);

      await listarArquivosBaseDir();

    }
    await ensureDirExists();
  } catch (error) {
    console.error("limparArmazenamento - Erro ao limpar armazenamento:", error);
  }
  console.log("=== TERMINO =======================================================================");
}

export async function salvarLista<T>(filename: string, data: T[]): Promise<void> {
  console.log("=== INICIO ========================================================================");
  console.log(`Iniciando o processo de salvar lista em ${filename}...`, " - BASE_DIR => ", BASE_DIR, "filename => ", filename, "data => ", "Dados:", data);
  await ensureDirExists();
  const fileUri = BASE_DIR + filename;
  const json = JSON.stringify(data, null, 2);
  console.log(`Salvando lista em ${fileUri}...`, "JSON => ", json);
  await FileSystem.writeAsStringAsync(fileUri, json);
  console.log("=== TERMINO =======================================================================");
}

export async function carregarLista<T>(filename: string): Promise<T[]> {
  console.log("=== INICIO ========================================================================");
  console.log(`Iniciando o processo de carregar lista de ${filename}...`, " - BASE_DIR => ", BASE_DIR);
  try {
    const fileUri = BASE_DIR + filename;
    console.log(`Carregando lista de fileUri =>  ${fileUri}...`);
    const fileInfo = await FileSystem.getInfoAsync(fileUri);
    console.log(`Verificando se o arquivo existe: ${fileInfo}`);
    if (!fileInfo.exists) return [];
    console.log(`Arquivo encontrado. Lendo conteúdo de ${fileUri}...`);
    const content = await FileSystem.readAsStringAsync(fileUri);
    console.log(`Conteúdo lido de ${fileUri}:`, content);
    console.log("=== TERMINO =======================================================================");
    return JSON.parse(content);
  } catch (error) {
    console.warn(`Erro ao carregar ${filename}:`, error);
    console.log("=== TERMINO =======================================================================");
    return [];
  }
}

export async function adicionarItem<T>(filename: string, novoItem: T): Promise<void> {
  console.log("=== INICIO ========================================================================");
  const lista = await carregarLista<T>(filename);
  lista.push(novoItem);
  await salvarLista<T>(filename, lista);
  console.log("=== TERMINO =======================================================================");
}

export async function removerItem<T>(filename: string, predicate: (item: T) => boolean): Promise<void> {
  console.log("=== INICIO ========================================================================");
  const lista = await carregarLista<T>(filename);
  const novaLista = lista.filter((item) => !predicate(item));
  await salvarLista<T>(filename, novaLista);
  console.log("=== TERMINO =======================================================================");
}

export async function salvarListaV2<T>(filename: string, lista: T[]): Promise<void> {
  console.log("=== INICIO ========================================================================");
  const fileUri = BASE_DIR + filename;
  console.log(`Iniciando o processo de salvar lista V2 em ${filename}...`, " - BASE_DIR => ", BASE_DIR, "filename => ", filename, "data => ", lista);
  try {
    await ensureDirExists();
    const jsonValue = JSON.stringify(lista);
    await AsyncStorage.setItem(fileUri, jsonValue);
  } catch (error) {
    throw new Error("Erro ao salvar lista" + fileUri + ": " + error  );
  }
  console.log("=== TERMINO =======================================================================");
}

// listar arquivos no diretório base
export async function listarArquivosBaseDir(): Promise<string[]> {
  console.log("=== INICIO ========================================================================");
  console.log("listarArquivosBaseDir - listando arquivos no diretório base:", BASE_DIR);
  try {
    const dirInfo = await FileSystem.getInfoAsync(BASE_DIR);
    if (!dirInfo.exists) {
      console.warn("listarArquivosBaseDir - Diretório base não encontrado ou não possui arquivos:", BASE_DIR);
      console.log("=== TERMINO =======================================================================");
      return [];
    }
    const files = await FileSystem.readDirectoryAsync(BASE_DIR);
    console.log("listarArquivosBaseDir - Arquivos encontrados no diretório base:", files);
    console.log("=== TERMINO =======================================================================");
    return files;
  } catch (error) {
    console.error("listarArquivosBaseDir - Erro ao listar arquivos no diretório base:", error);
    console.log("=== TERMINO =======================================================================");
    return [];
  }
}

// listar arquivos no diretório document
export async function listarArquivosDocDir(): Promise<string[]> { 
  console.log("=== INICIO ========================================================================");
  console.log("listarArquivosDocDir - listando arquivos no diretório document:", DOC_DIR);
  try {
    const dirInfo = await FileSystem.getInfoAsync(DOC_DIR);
    if (!dirInfo.exists) {
      console.warn("listarArquivosDocDir - Diretório document não encontrado ou não possui arquivos:", DOC_DIR);
      console.log("=== TERMINO =======================================================================");
      return [];
    }
    const files = await FileSystem.readDirectoryAsync(DOC_DIR);
    console.log("listarArquivosDocDir - Arquivos encontrados no diretório document:", files);
    console.log("=== TERMINO =======================================================================");
    return files;
  } catch (error) {
    console.error("listarArquivosDocDir - Erro ao listar arquivos no diretório document:", error);
    console.log("=== TERMINO =======================================================================");
    return [];
  }
} 
