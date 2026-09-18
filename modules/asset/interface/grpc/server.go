package grpc

import (
	"context"

	"nfxidentity/modules/asset/application/media"
	"nfxidentity/modules/asset/application/resource"
	grpcHandler "nfxidentity/modules/asset/interface/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	audiopb "nfxidentity/protos/gen/asset/audio"
	filepb "nfxidentity/protos/gen/asset/file"
	imagepb "nfxidentity/protos/gen/asset/image"
	videopb "nfxidentity/protos/gen/asset/video"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type Deps interface {
	ResourceSvc() *resource.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
	MediaSvc() *media.Service
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorHandler(),
			servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier()),
		),
	}
	s := grpc.NewServer(opts...)
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "asset"))
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(func(ctx context.Context) (int32, error) {
		n, err := postgresqlx.ClearSchema(ctx, d.Postgres().DB(), "asset", nil)
		return int32(n), err
	}))
	imagepb.RegisterImageServiceServer(s, grpcHandler.NewImageHandler(d.MediaSvc()))
	filepb.RegisterFileServiceServer(s, grpcHandler.NewFileHandler(d.MediaSvc()))
	videopb.RegisterVideoServiceServer(s, grpcHandler.NewVideoHandler(d.MediaSvc()))
	audiopb.RegisterAudioServiceServer(s, grpcHandler.NewAudioHandler(d.MediaSvc()))
	return s
}
