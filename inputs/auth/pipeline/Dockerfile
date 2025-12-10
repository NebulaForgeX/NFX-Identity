# Base build stage (cache dependencies)
FROM golang:alpine AS base-builder
WORKDIR /src
RUN apk update && apk upgrade && apk add --no-cache git ca-certificates tzdata
# ENV GOPROXY=https://goproxy.io,https://mirrors.aliyun.com/goproxy/,direct
# ENV GOSUMDB=sum.golang.org
# # Configure Git for better network reliability and GitHub mirror
# RUN git config --global http.postBuffer 524288000 && \
#     git config --global http.lowSpeedLimit 0 && \
#     git config --global http.lowSpeedTime 0 && \
#     git config --global http.version HTTP/1.1 && \
#     git config --global url."https://ghproxy.com/https://github.com/".insteadOf "https://github.com/"
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
  -o /out/auth-pipeline-service ./inputs/auth/pipeline

# Production runtime stage (minimal image)
FROM alpine:latest AS prod
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata
# Copy only the binary (rely on prod-build)
COPY --from=prod-build /out/auth-pipeline-service /app/auth-pipeline-service
COPY inputs/auth/config/prod.toml /app/config.toml
# Environment variable control（can be replaced with CONFIG_FILE）
ENV APP_ENV=prod
ENV CONFIG_FILE=/app/config.toml
ENTRYPOINT ["/app/auth-pipeline-service"]

# Dev stage（air hot reload）
FROM base-builder AS dev
WORKDIR /app
ENV PATH="/go/bin:$PATH"
RUN go install github.com/air-verse/air@v1.61.7
# Copy all source files for initial setup (will be synced/overridden by docker compose watch)
COPY . ./
CMD ["air", "-c", "inputs/auth/pipeline/.air.toml"]

