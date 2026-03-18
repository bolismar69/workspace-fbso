import json
import os
import glob

# Configurações de caminhos base
INVENTORY_PATH = 'architecture/governance/config/manager-solutions-inventory.json'
# Pastas onde residem as implementações de código
CODE_ROOTS = ['backend', 'frontend', 'orchestration', 'data_engineering']

def get_inventory_paths():
    """Lê o JSON e retorna um conjunto de caminhos registrados."""
    if not os.path.exists(INVENTORY_PATH):
        print(f"❌ Erro: Inventário não encontrado em {INVENTORY_PATH}")
        return set()
    
    with open(INVENTORY_PATH, 'r') as f:
        data = json.load(f)
    
    return {sol['path'].rstrip('/') for sol in data.get('solutions', [])}

def get_actual_paths():
    """
    Varre o monorepo em busca de diretórios de projetos reais.
    Consideramos um projeto como um diretório que não contém subdiretórios 
    que também sejam projetos (as 'folhas' da estrutura de stacks).
    """
    actual_paths = set()
    for root_dir in CODE_ROOTS:
        for root, dirs, files in os.walk(root_dir):
            # Filtro simples: se houver arquivos de build ou src, é um projeto
            indicators = ['src', 'go.mod', 'package.json', 'Dockerfile', 'manage.py', 'pom.xml', '.csproj']
            if any(ind in files or ind in dirs for ind in indicators):
                actual_paths.add(root.rstrip('/'))
                # Evita entrar em subpastas de um projeto já identificado (ex: node_modules)
                dirs[:] = [] 
    return actual_paths

def run_gari():
    print(f"{'='*60}")
    print(f"🧹 GARI - AUDITOR DE GOVERNANÇA (TaxNexus TaaS)")
    print(f"{'='*60}\n")

    inventory_paths = get_inventory_paths()
    actual_paths = get_actual_paths()

    # 1. Pastas Órfãs (Existem no disco, mas não no JSON)
    orphans = actual_paths - inventory_paths
    if orphans:
        print(f"⚠️  DIRETÓRIOS ÓRFÃOS (Não registrados no inventário):")
        for p in sorted(orphans):
            print(f"   [ ] {p}")
    else:
        print(f"✅ Nenhum diretório órfão encontrado.")

    print("")

    # 2. Projetos Fantasmas (Estão no JSON, mas não existem no disco)
    ghosts = inventory_paths - actual_paths
    if ghosts:
        print(f"👻 PROJETOS FANTASMAS (Registrados, mas o caminho não existe):")
        for p in sorted(ghosts):
            print(f"   [ ] {p}")
    else:
        print(f"✅ Todos os caminhos do inventário são válidos.")

    print(f"\n{'='*60}")
    print(f"Dica: Utilize 'rm -rf' para órfãos e atualize o JSON para fantasmas.")
    print(f"{'='*60}")

if __name__ == "__main__":
    run_gari()
