# Base build stage (cache dependencies)
FROM golang:alpine AS base-builder
WORKDIR /src
RUN apk update && apk upgrade && apk add --no-cache git ca-certificates tzdata
# Configure Go proxy for better network access in China
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn
COPY go.mod go.sum ./
RUN go mod download

# Production build stage (build binary for production)
FROM base-builder AS prod-build
COPY . ./
# Inject version information
ARG APP_VERSION=dev
ENV CGO_ENABLED=0 GOOS=linux
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go build -trimpath -ldflags="-s -w -X main.version=${APP_VERSION}" \
  -o /out/auth-api-service ./inputs/auth/api

# Production runtime stage (minimal image)
FROM alpine:latest AS prod
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
# Copy only the binary (rely on prod-build)
COPY --from=prod-build /out/auth-api-service /app/auth-api-service
COPY inputs/auth/config/prod.toml /app/config.toml
# Environment variable control（can be replaced with CONFIG_FILE）
ENV APP_ENV=prod
ENV CONFIG_FILE=/app/config.toml
EXPOSE 8080
ENTRYPOINT ["/app/auth-api-service"]

# Dev stage（air hot reload）
FROM base-builder AS dev
WORKDIR /app
ENV PATH="/go/bin:$PATH"
RUN go install github.com/air-verse/air@v1.61.7
# Copy all source files for initial setup (will be synced/overridden by docker compose watch)
COPY . ./
CMD ["sh", "-c", "cd /app/inputs/auth/api && /app/inputs/auth/air.sh"]

