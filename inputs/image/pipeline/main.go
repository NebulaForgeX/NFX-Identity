package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"nebulaid/modules/image/config"
	"nebulaid/modules/image/server"
	"nebulaid/pkgs/env"
	"nebulaid/pkgs/logx"

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
	if err := logx.Init(cfg.Logger, "image-pipeline", env.Env(*envStr)); err != nil {
		log.Fatalf("logger init failed: %v", err)
	}
	defer logx.Sync()

	// === Run Pipeline Server (Kafka Eventbus) ===
	if err := server.RunPipeline(ctx, cfg); err != nil && !errors.Is(err, context.Canceled) {
		logx.L().Fatal("pipeline server stopped with error", zap.Error(err))
	}

	logx.L().Info("pipeline server shutdown gracefully")
}
