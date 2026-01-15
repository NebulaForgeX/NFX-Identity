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
  -o /out/tenants-connection-service ./inputs/tenants/connection

FROM alpine:latest AS prod
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
COPY --from=prod-build /out/tenants-connection-service /app/tenants-connection-service
COPY inputs/tenants/configuration/prod.toml /app/config.toml
ENV APP_ENV=prod
ENV CONFIG_FILE=/app/config.toml
EXPOSE 50051
ENTRYPOINT ["/app/tenants-connection-service"]

FROM base-builder AS dev
WORKDIR /app
ENV PATH="/go/bin:$PATH"
RUN go install github.com/air-verse/air@v1.61.7
COPY . ./
CMD ["air", "-c", "inputs/tenants/connection/.air.toml"]
