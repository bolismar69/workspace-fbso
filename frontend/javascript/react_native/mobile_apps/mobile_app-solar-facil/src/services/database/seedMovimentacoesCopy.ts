// /src/services/database/seedMovimentacoes.ts
import * as SQLite from "expo-sqlite";
import { initializeDatabaseCopilot } from "./initializeDatabaseCopilot";

export async function seedMovimentacoesCopy(_associadoId: number, db: SQLite.SQLiteDatabase | null): Promise<void> {
  console.log("=== INICIO ========================================================================");
  console.log("seedMovimentacoes - Iniciando a inserção de registros de movimentação...",_associadoId);

  // Verifica se o banco de dados já foi inicializado
  try {
    if (!db) {
      db = await initializeDatabaseCopilot();
    }
  } catch (error) {
    console.error("seedMovimentacoes - openDatabaseAsync - Erro ao inicializar o banco de dados:", error);
    console.log("=== TERMINO ========================================================================");
    throw new Error("seedMovimentacoes - Erro ao abrir o banco de dados.");
  }

  // const associadoId = "00000000000000-01"; // ID fictício do associado
  const currentDate = new Date();

  for (let i = 0; i < 12; i++) {
    console.log("seedMovimentacoes - Inserindo registros de movimentação...", i + 1, "de 12");
    const movimentacaoDate = new Date(
      currentDate.getFullYear(),
      currentDate.getMonth() - i,
      1
    );

    const mes = movimentacaoDate.getMonth() + 1; // Mês (1-12)
    const ano = movimentacaoDate.getFullYear(); // Ano
    const valorTotal = parseFloat((Math.random() * 1000).toFixed(2)); // Valor total aleatório
    const dataCadastro = movimentacaoDate.toISOString(); // Data de cadastro no formato ISO
    const dataAtualizacao = dataCadastro;

    const dataVencimento = new Date(
      movimentacaoDate.getFullYear(),
      movimentacaoDate.getMonth(),
      Math.floor(Math.random() * 28) + 1
    ).toISOString(); // Data de vencimento aleatória
    const dataPagamento = Math.random() > 0.5 ? dataVencimento : null; // Data de pagamento aleatória ou null
    const statusPagamento = dataPagamento ? "Pago" : "Pendente"; // Status baseado na data de pagamento
    const observacoes = `Detalhe ${i + 1} da movimentação ${_associadoId}`; // Observações aleatórias

    const energiaRecebidaKwh = parseFloat((Math.random() * 100).toFixed(2)); // Energia recebida aleatória
    const valorEnergiaRecebida = parseFloat((energiaRecebidaKwh * 0.5).toFixed(2)); // Valor baseado na energia recebida
    const tarifaUnitariaKwh = 0.5; // Tarifa fixa
    const valorCobrado = parseFloat((valorEnergiaRecebida * 0.85).toFixed(2)); // Valor cobrado são 15% de desconto
    const valorEconomizado = parseFloat((valorEnergiaRecebida - valorCobrado).toFixed(2)); // Valor baseado na energia recebida();
    const percentualEconomizado = parseFloat(((1-(valorCobrado / valorEnergiaRecebida)) * 100).toFixed(2));

    console.log("seedMovimentacoes - Inserindo movimentação: ", 
      [ dataCadastro, dataAtualizacao, _associadoId, mes,
        ano, valorTotal, dataVencimento, dataPagamento, statusPagamento, observacoes,
        energiaRecebidaKwh, valorEnergiaRecebida, tarifaUnitariaKwh, valorCobrado
      ]);
    
    const movimentacaoId = (_associadoId + i +1 );

    // Inserir registro na tabela movimentacoes
    const result = await db?.runAsync(
      `INSERT INTO movimentacoes (id, dataCadastro, dataAtualizacao, associadoId, mes,
      ano, valorTotal, dataVencimento, dataPagamento, statusPagamento, observacoes,
      energiaRecebidaKwh, valorEnergiaRecebida, tarifaUnitariaKwh, valorCobrado,
      valorEconomizado, percentualEconomizado)
       VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );`,
      [ movimentacaoId, dataCadastro, dataAtualizacao, _associadoId, mes,
        ano, valorTotal, dataVencimento, dataPagamento, statusPagamento, observacoes,
        energiaRecebidaKwh, valorEnergiaRecebida, tarifaUnitariaKwh, valorCobrado,
        valorEconomizado, percentualEconomizado
      ]
    );

    console.log(`seedMovimentacoes - Registro de movimentação inserido:`, result);
  }

  console.log("seedMovimentacoes - Registros de movimentação e detalhes inseridos com sucesso.");
  console.log("=== TERMINO ========================================================================");
}
