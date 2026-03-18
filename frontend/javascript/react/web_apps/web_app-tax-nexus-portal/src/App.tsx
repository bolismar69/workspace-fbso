import React, { useState } from 'react';
import TaxSimulator from './components/TaxSimulator';

function App() {
  const [cnpj, setCnpj] = useState('');
  const [authenticated, setAuthenticated] = useState(false);

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    if (cnpj.length >= 14) {
      setAuthenticated(true);
    } else {
      alert("Por favor, insira um CNPJ válido.");
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 py-10">
      {!authenticated ? (
        <div className="max-w-md mx-auto bg-white p-8 rounded-lg shadow-md">
          <h1 className="text-2xl font-bold mb-6 text-center text-blue-900">TaxNexus TaaS</h1>
          <form onSubmit={handleLogin} className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">CNPJ do Contribuinte</label>
              <input 
                type="text" 
                className="mt-1 block w-full p-3 border rounded-md shadow-sm"
                placeholder="00000000000000"
                value={cnpj}
                onChange={(e) => setCnpj(e.target.value)}
              />
            </div>
            <button 
              type="submit"
              className="w-full bg-blue-900 text-white py-3 rounded-md font-semibold hover:bg-blue-800"
            >
              ACESSAR PORTAL
            </button>
          </form>
        </div>
      ) : (
        <div className="container mx-auto">
          <header className="flex justify-between items-center mb-8 px-6">
            <h1 className="text-xl font-bold text-blue-900">TaxNexus Simulator v1.0</h1>
            <button 
              onClick={() => setAuthenticated(false)}
              className="text-sm text-red-600 hover:underline"
            >
              Sair / Trocar CNPJ
            </button>
          </header>
          
          <TaxSimulator cnpj={cnpj} />
        </div>
      )}
    </div>
  );
}

export default App;
