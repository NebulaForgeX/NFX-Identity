package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"nfxid/modules/directory/config"
	"nfxid/modules/directory/server"
	"nfxid/pkgs/env"
	"nfxid/pkgs/logx"

	"go.uber.org/zap"
)

func main() {
	envStr := flag.String("env", "dev", "Environment (dev/prod)")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load(ctx, env.Env(*envStr))
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	if err := logx.Init(cfg.Logger, "directory-base-service", env.Env(*envStr)); err != nil {
		log.Fatalf("logger init failed: %v", err)
	}
	defer logx.Sync()

	if err := server.RunServer(ctx, cfg); err != nil && !errors.Is(err, context.Canceled) {
		logx.L().Fatal("directory server stopped with error", zap.Error(err))
	}

	logx.L().Info("directory server shutdown gracefully")
}
