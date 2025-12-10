package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"

	"nebulaid/modules/image/config"
	eventbusInterfaces "nebulaid/modules/image/interfaces/eventbus"
	grpcInterfaces "nebulaid/modules/image/interfaces/grpc"
	httpInterfaces "nebulaid/modules/image/interfaces/http"
	"nebulaid/pkgs/logx"

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

	logx.S().Info("✅ HTTP Server: All dependencies initialized successfully (PostgreSQL, MongoDB, Kafka Publisher)")

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

	logx.S().Info("✅ gRPC Server: All dependencies initialized successfully (PostgreSQL, MongoDB, Kafka Publisher)")

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
	logx.S().Info("✅ Kafka Subscriber: All dependencies initialized successfully (PostgreSQL, MongoDB, Kafka Publisher)")
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
