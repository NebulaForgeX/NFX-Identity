# Base build stage (cache dependencies)
FROM golang:alpine AS base-builder
WORKDIR /src
RUN apk update && apk upgrade && apk add --no-cache git ca-certificates tzdata
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn
COPY go.mod go.sum ./
RUN go mod download

FROM base-builder AS prod-build
COPY . ./
ARG APP_VERSION=dev
ENV CGO_ENABLED=0 GOOS=linux
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go build -trimpath -ldflags="-s -w -X main.version=${APP_VERSION}" \
  -o /out/clients-connection-service ./inputs/clients/connection

FROM alpine:latest AS prod
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
COPY --from=prod-build /out/clients-connection-service /app/clients-connection-service
COPY inputs/clients/configuration/prod.toml /app/config.toml
ENV APP_ENV=prod
ENV CONFIG_FILE=/app/config.toml
EXPOSE 50051
ENTRYPOINT ["/app/clients-connection-service"]

FROM base-builder AS dev
WORKDIR /app
ENV PATH="/go/bin:$PATH"
RUN go install github.com/air-verse/air@v1.61.7
COPY . ./
CMD ["air", "-c", "inputs/clients/connection/.air.toml"]
