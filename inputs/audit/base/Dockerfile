# Base build stage (cache dependencies)
FROM golang:1.24.4-alpine AS base-builder
WORKDIR /src
RUN apk add --no-cache git ca-certificates tzdata
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
  -o /out/audit-base-service ./inputs/audit/base

# Production runtime stage (minimal image)
FROM alpine:latest AS prod
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
# Copy only the binary (rely on prod-build)
COPY --from=prod-build /out/audit-base-service /app/audit-base-service
COPY inputs/audit/configuration/prod.toml /app/config.toml
# Copy .env file for variable substitution
COPY .env /app/.env
# Environment variable control（can be replaced with CONFIG_FILE）
ENV APP_ENV=prod
ENV CONFIG_FILE=/app/config.toml
EXPOSE 8080
EXPOSE 10002
ENTRYPOINT ["/app/audit-base-service"]

# Dev stage（air hot reload）
FROM base-builder AS dev
WORKDIR /app
ENV PATH="/go/bin:$PATH"
RUN go install github.com/air-verse/air@v1.61.7
# Copy all source files for initial setup (will be synced/overridden by docker compose watch)
COPY . ./
CMD ["air", "-c", "inputs/audit/base/.air.toml"]
