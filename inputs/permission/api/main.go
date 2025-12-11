package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"nfxid/modules/permission/config"
	"nfxid/modules/permission/server"
	"nfxid/pkgs/env"
	"nfxid/pkgs/logx"

	"go.uber.org/zap"
)

func main() {
	envStr := flag.String("env", "dev", "Environment (dev/prod)")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// === Load Config ===
	cfg, err := config.Load(ctx, env.Env(*envStr))
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	// === Init Logger ===
	if err := logx.Init(cfg.Logger, "permission-api", env.Env(*envStr)); err != nil {
		log.Fatalf("logger init failed: %v", err)
	}
	defer logx.Sync()

	// === Run API Server (HTTP) ===
	if err := server.RunHTTP(ctx, cfg); err != nil && !errors.Is(err, context.Canceled) {
		logx.L().Fatal("api server stopped with error", zap.Error(err))
	}

	logx.L().Info("api server shutdown gracefully")
}
