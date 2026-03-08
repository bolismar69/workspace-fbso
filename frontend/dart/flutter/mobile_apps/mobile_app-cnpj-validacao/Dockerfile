# Build Flutter Web (HTML renderer) and serve via Nginx for browser-based tests
FROM ghcr.io/cirruslabs/flutter:stable AS build
WORKDIR /app

# Workaround for git safe.directory in this image
RUN git config --global --add safe.directory /sdks/flutter

COPY pubspec.yaml pubspec.lock ./
RUN flutter pub get

COPY . .

# Use HTML renderer so inputs are easier to drive in Playwright
RUN flutter build web --release --dart-define=FLUTTER_WEB_USE_SKIA=false

FROM nginx:1.27-alpine AS runtime
WORKDIR /usr/share/nginx/html

COPY --from=build /app/build/web ./

EXPOSE 80
