package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"

	"nfxidentity/modules/asset/config"
	grpcInterfaces "nfxidentity/modules/asset/interface/grpc"
	httpInterfaces "nfxidentity/modules/asset/interface/http"
	"nfxidentity/pkgs/logx"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func RunHTTP(ctx context.Context, cfg *config.Config) error {
	deps, err := NewDeps(ctx, cfg)
	if err != nil {
		return err
	}
	defer deps.Cleanup()
	httpSrv := httpInterfaces.NewHTTPServer(deps, cfg)
	httpAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.HTTPPort))
	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		logx.S().Infof("✅ Asset HTTP server listening on %s", httpAddr)
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

func RunGRPC(ctx context.Context, cfg *config.Config) error {
	deps, err := NewDeps(ctx, cfg)
	if err != nil {
		return err
	}
	defer deps.Cleanup()
	grpcSrv := grpcInterfaces.NewServer(deps)
	grpcAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.GRPCPort))
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}
	defer lis.Close()
	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		logx.S().Infof("✅ Asset gRPC server listening on %s", grpcAddr)
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

func RunServer(ctx context.Context, cfg *config.Config) error {
	deps, err := NewDeps(ctx, cfg)
	if err != nil {
		return err
	}
	defer deps.Cleanup()
	httpSrv := httpInterfaces.NewHTTPServer(deps, cfg)
	grpcSrv := grpcInterfaces.NewServer(deps)
	httpAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.HTTPPort))
	grpcAddr := net.JoinHostPort(cfg.Server.Host, strconv.Itoa(cfg.Server.GRPCPort))
	grpcLis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}
	defer grpcLis.Close()
	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		logx.S().Infof("✅ Asset HTTP server listening on %s", httpAddr)
		if err := httpSrv.Listen(httpAddr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})
	g.Go(func() error {
		logx.S().Infof("✅ Asset gRPC server listening on %s", grpcAddr)
		if err := grpcSrv.Serve(grpcLis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			return err
		}
		return nil
	})
	g.Go(func() error {
		<-gctx.Done()
		_ = httpSrv.Shutdown()
		grpcSrv.GracefulStop()
		return gctx.Err()
	})
	return g.Wait()
}
