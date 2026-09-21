package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"nfxidentity/modules/asset/config"
	"nfxidentity/modules/asset/server"
	"nfxidentity/pkgs/connections/otelx"
	"nfxidentity/pkgs/env"
	"nfxidentity/pkgs/logx"

	"go.uber.org/zap"
)

func main() {
	envStr := flag.String("env", "dev", "Environment (dev/secure)")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// === Load Config ===
	cfg, err := config.Load(ctx, env.Env(*envStr))
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	// === Init Logger ===
	if err := logx.Init(cfg.Logger, "asset-base-service", env.Env(*envStr)); err != nil {
		log.Fatalf("logger init failed: %v", err)
	}
	defer logx.Sync()

	otelShutdown, err := otelx.Init(ctx, cfg.OTEL, "asset", env.Env(*envStr))
	if err != nil {
		log.Fatalf("otel init failed: %v", err)
	}
	defer func() { _ = otelShutdown(context.Background()) }()

	// === Run Server (all services in goroutines) ===
	if err := server.RunServer(ctx, cfg); err != nil && !errors.Is(err, context.Canceled) {
		logx.L().Fatal("asset server stopped with error", zap.Error(err))
	}

	logx.L().Info("asset server shutdown gracefully")
}
