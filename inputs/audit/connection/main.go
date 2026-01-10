package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"nfxid/modules/audit/config"
	"nfxid/modules/audit/server"
	"nfxid/pkgs/env"
	"nfxid/pkgs/logx"

	"go.uber.org/zap"
)

func main() {
	envStr := flag.String("env", "prod", "Environment (dev/prod)")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// === Load Config ===
	cfg, err := config.Load(ctx, env.Env(*envStr))
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	// === Init Logger ===
	if err := logx.Init(cfg.Logger, "audit-connection-service", env.Env(*envStr)); err != nil {
		log.Fatalf("logger init failed: %v", err)
	}
	defer logx.Sync()

	// === Run gRPC Server ===
	if err := server.RunGRPC(ctx, cfg); err != nil && !errors.Is(err, context.Canceled) {
		logx.L().Fatal("gRPC server stopped with error", zap.Error(err))
	}

	logx.L().Info("gRPC server shutdown gracefully")
}
