package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"

	"nfxid/modules/permission/config"
	eventbusInterfaces "nfxid/modules/permission/interfaces/eventbus"
	grpcInterfaces "nfxid/modules/permission/interfaces/grpc"
	httpInterfaces "nfxid/modules/permission/interfaces/http"
	"nfxid/pkgs/logx"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// RunHTTP starts the HTTP server (used by api/main.go)
func RunHTTP(ctx context.Context, cfg *config.Config) error {
	// === Dependencies ===
	deps, err := NewDependencies(ctx, cfg)
	if err != nil {
		return err
	}
	defer deps.Cleanup()

	logx.S().Info("✅ HTTP Server: All dependencies initialized successfully (PostgreSQL, MongoDB, Redis, Kafka Publisher)")

	httpSrv := httpInterfaces.NewHTTPServer(deps)
	httpAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.HTTPPort))
	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		logx.S().Infof("✅ HTTP server listening on %s", httpAddr)
		if err := httpSrv.Listen(httpAddr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-gctx.Done()
		_ = httpSrv.Shutdown()
		return gctx.Err()
	})

	return g.Wait()
}

// RunGRPC starts the gRPC server (used by connection/main.go)
func RunGRPC(ctx context.Context, cfg *config.Config) error {
	// === Dependencies ===
	deps, err := NewDependencies(ctx, cfg)
	if err != nil {
		return err
	}
	defer deps.Cleanup()

	logx.S().Info("✅ gRPC Server: All dependencies initialized successfully (PostgreSQL, MongoDB, Redis, Kafka Publisher)")

	grpcSrv := grpcInterfaces.NewServer(deps)
	grpcAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.GRPCPort))

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}
	defer lis.Close()

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		logx.S().Infof("✅ gRPC server listening on %s", grpcAddr)
		if err := grpcSrv.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-gctx.Done()
		grpcSrv.GracefulStop()
		return gctx.Err()
	})

	return g.Wait()
}

// RunPipeline starts the Kafka eventbus server (used by pipeline/main.go)
func RunPipeline(ctx context.Context, cfg *config.Config) error {
	// === Dependencies ===
	deps, err := NewDependencies(ctx, cfg)
	if err != nil {
		return err
	}
	defer deps.Cleanup()
	logx.S().Info("✅ Kafka Subscriber: All dependencies initialized successfully (PostgreSQL, MongoDB, Redis, Kafka Publisher)")
	eventbusSrv, err := eventbusInterfaces.NewServer(deps)
	if err != nil {
		return err
	}

	logx.S().Info("✅ Kafka Subscriber initialized successfully")

	eventbusSrv.RegisterRoutes()

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		logx.S().Infof("✅ Eventbus (Kafka) server listening on %v", cfg.KafkaConfig.Brokers)
		return eventbusSrv.Run(ctx)
	})

	g.Go(func() error {
		<-gctx.Done()
		_ = eventbusSrv.Close()
		return gctx.Err()
	})

	return g.Wait()
}

// RunServer starts all three services (HTTP, gRPC, Pipeline) concurrently
// This is useful for development or when you want to run all services in a single process
func RunServer(ctx context.Context, cfg *config.Config) error {
	// === Dependencies ===
	deps, err := NewDependencies(ctx, cfg)
	if err != nil {
		return err
	}
	defer deps.Cleanup()

	logx.S().Info("✅ All-in-One Server: All dependencies initialized successfully (PostgreSQL, MongoDB, Redis, Kafka Publisher)")

	// === Initialize Servers ===
	httpSrv := httpInterfaces.NewHTTPServer(deps)
	grpcSrv := grpcInterfaces.NewServer(deps)
	eventbusSrv, err := eventbusInterfaces.NewServer(deps)
	if err != nil {
		return err
	}

	httpAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.HTTPPort))
	grpcAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.GRPCPort))

	grpcLis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}
	defer grpcLis.Close()

	eventbusSrv.RegisterRoutes()

	g, gctx := errgroup.WithContext(ctx)

	// HTTP Server
	g.Go(func() error {
		logx.S().Infof("✅ HTTP server listening on %s", httpAddr)
		if err := httpSrv.Listen(httpAddr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	// gRPC Server
	g.Go(func() error {
		logx.S().Infof("✅ gRPC server listening on %s", grpcAddr)
		if err := grpcSrv.Serve(grpcLis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			return err
		}
		return nil
	})

	// Pipeline (Kafka Eventbus) Server
	g.Go(func() error {
		logx.S().Infof("✅ Eventbus (Kafka) server listening on %v", cfg.KafkaConfig.Brokers)
		return eventbusSrv.Run(ctx)
	})

	// Graceful shutdown handler
	g.Go(func() error {
		<-gctx.Done()
		logx.S().Info("🛑 Shutting down all services...")
		_ = httpSrv.Shutdown()
		grpcSrv.GracefulStop()
		_ = eventbusSrv.Close()
		return gctx.Err()
	})

	return g.Wait()
}

