#!/bin/bash

# Configurações de cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# 1. Validação de Repositório Git
if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo -e "${RED}ERRO: Este diretório não é um repositório Git.${NC}"
    exit 1
fi

# Variáveis dos parâmetros
TIPO=$1
ESCOPO=$2
MENSAGEM=$3
TAG_INPUT=$4

exibir_help() {
    echo -e "${YELLOW}————————————————————————————————————————————————————————————————"
    echo -e "MODO DE USO:"
    echo -e "  ./gitcommit.sh <tipo> <escopo> <mensagem> <tag>"
    echo -e ""
    echo -e "PARÂMETROS:"
    echo -e "  1. TIPO:      feat, hotfix, chore, refactor"
    echo -e "  2. ESCOPO:    Texto curto (ex: pix, api)"
    echo -e "  3. MENSAGEM:  Descrição (entre aspas)"
    echo -e "  4. TAG:       wip ou final"
    echo -e ""
    echo -e "EXEMPLO:"
    echo -e "  ./gitcommit.sh feat pix 'integração banco central' final"
    echo -e "————————————————————————————————————————————————————————————————${NC}"
}

listar_historico() {
    echo -e "\n${CYAN}🕒 Últimos 5 commits na branch $BRANCH_ATUAL:${NC}"
    # Formata o log para mostrar hash curta, data relativa e a mensagem colorida
    git log -5 --pretty=format:"%C(yellow)%h%Creset %C(green)(%cr)%Creset %s"
    echo -e "\n"
}

# 2. Validação de quantidade de parâmetros
if [ $# -ne 4 ]; then
    echo -e "${RED}Erro: Quantidade de parâmetros inválida (recebidos: $# | esperados: 4).${NC}"
    exibir_help
    exit 1
fi

# 3. Bloqueio de Branches Protegidas
BRANCH_ATUAL=$(git rev-parse --abbrev-ref HEAD)
if [[ "$BRANCH_ATUAL" == "main" || "$BRANCH_ATUAL" == "master" || "$BRANCH_ATUAL" == "develop" || "$BRANCH_ATUAL" == release/* ]]; then
    echo -e "${RED}❌ OPERAÇÃO BLOQUEADA!${NC}"
    echo -e "Você está na branch de estado: ${YELLOW}$BRANCH_ATUAL${NC}"
    echo -e "Não é permitido commitar diretamente aqui (Regras de GitOps FBSO)."
    echo -e "Use uma branch de ${GREEN}feature/${NC} ou ${GREEN}hotfix/${NC}.${NC}"
    exit 1
fi

# 4. Validação do TIPO
case $TIPO in
    feat|hotfix|chore|refactor) ;;
    *) echo -e "${RED}Erro: Tipo '$TIPO' inválido.${NC}"; exibir_help; exit 1 ;;
esac

# 5. Validação e Formatação da TAG
TAG_LOWER=$(echo "$TAG_INPUT" | tr '[:upper:]' '[:lower:]')
if [[ "$TAG_LOWER" == "wip" ]]; then 
    TAG_FINAL="[WIP]"
elif [[ "$TAG_LOWER" == "final" ]]; then 
    TAG_FINAL="[FINAL]"
else 
    echo -e "${RED}Erro: Tag '$TAG_INPUT' inválida.${NC}"; exibir_help; exit 1
fi

# 6. Geração do Timestamp [YYYYMMDD-HHMMSS]
DATA_HORA=$(date +'%Y%m%d-%H%M%S')
TS="[$DATA_HORA]"

# 7. Verificação de Alterações (Staging)
echo -e "${BLUE}Analisando alterações...${NC}"
git add -A
if git diff --cached --quiet; then
    echo -e "${YELLOW}Nada para commitar. O diretório de trabalho está limpo.${NC}"
    listar_historico
    exit 0
fi

# 8. Confirmação e Execução
COMMIT_FULL_MSG="$TIPO($ESCOPO): $MENSAGEM $TAG_FINAL $TS"

echo -e "${YELLOW}————————————————————————————————————————————————————————————————"
echo -e "RESUMO DO COMMIT:"
echo -e "Branch:   $BRANCH_ATUAL"
echo -e "Mensagem: $COMMIT_FULL_MSG"
echo -e "————————————————————————————————————————————————————————————————${NC}"

read -p "Confirma o commit e push para origin? (s/n): " CONFIRMACAO

if [[ "$CONFIRMACAO" =~ ^[sS]$ ]]; then
    echo -e "${BLUE}Executando commit...${NC}"
    if git commit -m "$COMMIT_FULL_MSG"; then
        echo -e "${BLUE}Executando push para origin $BRANCH_ATUAL...${NC}"
        if git push origin "$BRANCH_ATUAL"; then
            echo -e "${GREEN}✅ Operação finalizada com sucesso!${NC}"
            listar_historico
        else
            echo -e "${RED}❌ Erro ao realizar o push.${NC}"
            exit 1
        fi
    else
        echo -e "${RED}❌ Erro ao realizar o commit.${NC}"
        exit 1
    fi
else
    echo -e "${RED}Operação cancelada.${NC}"
fi
