package grpc

import (
	imageApp "nfxid/modules/image/application/image"
	imageTypeApp "nfxid/modules/image/application/image_type"
	grpcHandler "nfxid/modules/image/interfaces/grpc/handler"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token"
	imagepb "nfxid/protos/gen/image/image"
	imagetypepb "nfxid/protos/gen/image/image_type"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

type Deps interface {
	ServerTokenVerifier() token.Verifier
	ImageAppSvc() *imageApp.Service
	ImageTypeAppSvc() *imageTypeApp.Service
}

func NewServer(d Deps) *grpc.Server {
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)

	s := grpc.NewServer(grpcx.DefaultServerOptions(d.ServerTokenVerifier())...)

	// Register gRPC services
	imagepb.RegisterImageServiceServer(s, grpcHandler.NewImageHandler(
		d.ImageAppSvc(),
	))

	imagetypepb.RegisterImageTypeServiceServer(s, grpcHandler.NewImageTypeHandler(
		d.ImageTypeAppSvc(),
	))

	return s
}
