import React, { useState } from 'react';
import { BarChart, Bar, XAxis, YAxis, Tooltip, CartesianGrid, ResponsiveContainer, Legend } from 'recharts';
import { useTaxService } from '../hooks/useTaxService';

interface TaxSimulatorProps {
  cnpj: string;
}

const TaxSimulator: React.FC<TaxSimulatorProps> = ({ cnpj }) => {
  const { calculateTax, loading } = useTaxService();
  const [selectedState, setSelectedState] = useState('');
  const [selectedCity, setSelectedCity] = useState('');
  const [ncm, setNcm] = useState('62011100');
  const [saldo, setSaldo] = useState<number>(0);
  const [chartData, setChartData] = useState<any[] | null>(null);
  const [rawResponse, setRawResponse] = useState<any>(null);

  const states = [
    { id: '35', sigla: 'SP', nome: 'São Paulo' },
    { id: '33', sigla: 'RJ', nome: 'Rio de Janeiro' },
    { id: '13', sigla: 'AM', nome: 'Amazonas' }
  ];

  const citiesMap: Record<string, { ibge: string, nome: string }[]> = {
    '35': [{ ibge: '3550308', nome: 'São Paulo' }],
    '33': [{ ibge: '3304557', nome: 'Rio de Janeiro' }],
    '13': [{ ibge: '1302603', nome: 'Manaus' }]
  };

  const handleSimulate = async () => {
    const data = await calculateTax({
      cnpj,
      ncm,
      ibge: selectedCity,
      saldo_remanescente: saldo
    });

    if (data && data.calculation) {
      setRawResponse(data.calculation);

      const formatted = [
        {
          name: '2026 (Transição)',
          cbs: data.calculation.cbs_calculada,
          ibs: data.calculation.ibs_calculado
        },
        {
          name: '2027 (Projeção)',
          cbs: data.calculation.cbs_calculada * 1.02,
          ibs: data.calculation.ibs_calculado * 1.02
        }
      ];
      setChartData(formatted);
    }
  };

  return (
    <div className="p-6 max-w-4xl mx-auto bg-white rounded-lg shadow-lg space-y-6">
      {/* Cabeçalho */}
      <div className="mb-6 p-4 bg-gray-50 rounded border-l-4 border-blue-500">
        <p className="text-sm text-gray-600">Empresa Conectada: <span className="font-bold">{cnpj}</span></p>
      </div>

      {/* Seletores */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-2">
        <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Estado: </label>
          <select className="p-2 border rounded" onChange={(e) => setSelectedState(e.target.value)}>
            <option value="">Selecione o Estado</option>
            {states.map(s => <option key={s.id} value={s.id}>{s.nome}</option>)}
          </select>
        </div>

        <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Cidade: </label>
          <select className="p-2 border rounded" disabled={!selectedState} onChange={(e) => setSelectedCity(e.target.value)}>
            <option value="">Selecione a Cidade</option>
            {(citiesMap[selectedState] || []).map(c => <option key={c.ibge} value={c.ibge}>{c.nome}</option>)}
          </select>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">NCM: </label>
          <input type="text" value={ncm} onChange={(e) => setNcm(e.target.value)} className="p-2 border rounded" placeholder="NCM" />
        </div>
      </div>

      {/* Saldo Remanescente */}
      <div className="bg-yellow-50 p-4 rounded-md border border-yellow-100">
        <label className="block text-sm font-medium text-yellow-800">Saldo Remanescente de Créditos (R$)</label>
        <input type="number" value={saldo} onChange={(e) => setSaldo(Number(e.target.value))} className="mt-1 block w-full p-2 border border-yellow-300 rounded-md" />
      </div>

      <button
        onClick={handleSimulate}
        disabled={loading || !selectedCity}
        className="w-full bg-blue-600 text-white py-3 rounded font-bold hover:bg-blue-700 disabled:bg-gray-400"
      >
        {loading ? 'Processando...' : 'SIMULAR REFORMA TRIBUTÁRIA'}
      </button>

      {/* RESULTADOS - Só renderiza se rawResponse existir */}
      {rawResponse && (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6 animate-in slide-in-from-top-4 duration-700">
            {/* Sistema Atual */}
            <div className="bg-orange-50 p-5 rounded-xl border border-orange-200">
              <h3 className="font-bold text-orange-800 border-b pb-2 mb-3">Sistema Atual (Legado)</h3>
              <ul className="space-y-2 text-sm">
                <li className="flex justify-between"><span>PIS:</span> <b>{rawResponse.pis?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</b></li>
                <li className="flex justify-between"><span>COFINS:</span> <b>{rawResponse.cofins?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</b></li>
                <li className="flex justify-between"><span>ICMS:</span> <b>{rawResponse.icms?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</b></li>
                <li className="flex justify-between"><span>IPI:</span> <b>{rawResponse.ipi?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</b></li>
              </ul>
            </div>

            {/* Nova Reforma */}
            <div className="bg-blue-50 p-5 rounded-xl border border-blue-200">
              <h3 className="font-bold text-blue-800 border-b pb-2 mb-3">Pós-Reforma (IVA Dual)</h3>
              <ul className="space-y-2 text-sm">
                <li className="flex justify-between"><span>CBS:</span> <b>{rawResponse.cbs_calculada?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</b></li>
                <li className="flex justify-between"><span>IBS:</span> <b>{rawResponse.ibs_calculado?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</b></li>
                <li className="flex justify-between text-red-600 font-bold"><span>Seletivo:</span> <span>{rawResponse.imposto_seletivo?.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</span></li>
                <li className="flex justify-between text-purple-700 border-t pt-2"><span>Novos IPVA/ITCMD:</span> <b>{(rawResponse.ipva_novo + rawResponse.itcmd_novo).toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</b></li>
              </ul>
            </div>
          </div>

          {/* Gráfico */}
          {chartData && (
            <div className="h-80 w-full bg-white p-4 rounded-lg border shadow-inner mt-6">
              <h3 className="text-center text-gray-700 font-bold mb-4">Comparativo de Transição</h3>
              <ResponsiveContainer width="100%" height="100%">
                <BarChart data={chartData}>
                  <CartesianGrid strokeDasharray="3 3" vertical={false} />
                  <XAxis dataKey="name" />
                  <YAxis />
                  {/* <Tooltip formatter={(value: number) => value.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })} /> */}
                  <Tooltip
                    formatter={(value: any) => {
                      // Converte para número caso venha algo inesperado e formata
                      const num = Number(value);
                      return isNaN(num)
                        ? "R$ 0,00"
                        : num.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
                    }}
                  />
                  <Legend />
                  <Bar dataKey="cbs" name="CBS" fill="#2563eb" stackId="a" />
                  <Bar dataKey="ibs" name="IBS" fill="#059669" stackId="a" />
                </BarChart>
              </ResponsiveContainer>
            </div>
          )}
        </>
      )}
    </div>
  );
};

export default TaxSimulator;