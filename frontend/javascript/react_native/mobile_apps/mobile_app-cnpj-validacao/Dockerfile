# Build an Expo Web export and serve it via Nginx (for browser-based functional tests)
FROM node:20-bookworm-slim AS build
WORKDIR /app

ENV CI=true
ENV EXPO_NO_TELEMETRY=1

COPY package.json package-lock.json ./
RUN npm ci --no-audit --no-fund

COPY . .

# Export static web build
RUN npx expo export --platform web --output-dir dist

FROM nginx:1.27-alpine AS runtime
WORKDIR /usr/share/nginx/html

COPY --from=build /app/dist ./

EXPOSE 80
