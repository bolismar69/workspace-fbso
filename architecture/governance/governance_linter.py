import json
import os

def check_inventory_consistency(inventory_path):
    with open(inventory_path, 'r') as f:
        data = json.load(f)
    
    for sol in data['solutions']:
        name = sol['name']
        print(f"--- Verificando: {name} ---")
        
        # 1. Verifica se o PATH do projeto existe no monorepo
        if not os.path.exists(sol['path']):
            print(f"❌ ERRO: Caminho do projeto não encontrado: {sol['path']}")
            
        # 2. Verifica se o Dockerfile de blueprint existe
        dockerfile_path = sol['specification']['docker']['dockerfile']
        if not os.path.exists(dockerfile_path):
            print(f"❌ ERRO: Blueprint Dockerfile não encontrado: {dockerfile_path}")
            
        # 3. Valida se o systemNamespace segue o padrão (ex: letras minúsculas e hífens)
        ns = sol['systemNamespace']
        if not ns.islower() or " " in ns:
            print(f"⚠️ AVISO: systemNamespace '{ns}' deve seguir o padrão K8S (lowercase-kebab-case)")

        # 4. Verifica dependências de contratos
        for contract in sol['dependsOn'].get('contracts', []):
            if not os.path.exists(contract):
                print(f"❌ ERRO: Contrato Protobuf não encontrado no caminho: {contract}")

    print("\n--- Auditoria Concluída ---")

if __name__ == "__main__":
    check_inventory_consistency('architecture/governance/config/manager-solutions-inventory.json')
