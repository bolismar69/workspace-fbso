# STAGE 1: Build
FROM node:18-alpine AS builder

# Definindo diretório de trabalho
WORKDIR /app

# Copiando manifestos de dependências
COPY package.json package-lock.json ./

# Instalação limpa de dependências (frozen lockfile para governança)
RUN npm ci

# Copiando o código fonte
COPY . .

# Build de produção (gera a pasta /dist ou /build)
RUN npm run build

# STAGE 2: Runtime (Final Image)
FROM nginx:stable-alpine

# Compliance: Removendo configurações padrão do Nginx que podem expor info do server
RUN rm -rf /etc/nginx/conf.d/default.conf

# Copiando configuração otimizada do Nginx (ver nota abaixo)
# COPY nginx.conf /etc/nginx/conf.d/

# Copiando apenas os artefatos compilados do estágio anterior
COPY --from=builder /app/dist /usr/share/nginx/html

# Exposição da porta definida no seu inventário
EXPOSE 5173

# Ajuste para rodar o Nginx na porta 5173 (conforme seu JSON de conformidade)
RUN sed -i 's/listen  80;/listen  5173;/g' /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
