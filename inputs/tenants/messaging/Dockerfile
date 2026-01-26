# Base build stage (cache dependencies)
FROM golang:1.24.4-alpine AS base-builder
WORKDIR /src
RUN apk add --no-cache git ca-certificates tzdata
COPY go.mod go.sum ./
RUN go mod download

FROM base-builder AS prod-build
COPY . ./
ARG APP_VERSION=dev
ENV CGO_ENABLED=0 GOOS=linux
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go build -trimpath -ldflags="-s -w -X main.version=${APP_VERSION}" \
  -o /out/tenants-messaging-service ./inputs/tenants/messaging

FROM alpine:latest AS prod
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
COPY --from=prod-build /out/tenants-messaging-service /app/tenants-messaging-service
COPY inputs/tenants/configuration/prod.toml /app/config.toml
# Copy .env file for variable substitution
COPY .env /app/.env
ENV APP_ENV=prod
ENV CONFIG_FILE=/app/config.toml
ENTRYPOINT ["/app/tenants-messaging-service"]

FROM base-builder AS dev
WORKDIR /app
ENV PATH="/go/bin:$PATH"
RUN go install github.com/air-verse/air@v1.61.7
COPY . ./
CMD ["air", "-c", "inputs/tenants/messaging/.air.toml"]
