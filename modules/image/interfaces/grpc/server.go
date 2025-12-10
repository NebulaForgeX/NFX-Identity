package grpc

import (
	imageApp "nebulaid/modules/image/application/image"
	imageTypeApp "nebulaid/modules/image/application/image_type"
	grpcHandler "nebulaid/modules/image/interfaces/grpc/handler"
	"nebulaid/pkgs/grpcx"
	"nebulaid/pkgs/security/token"
	imagepb "nebulaid/protos/gen/image/image"
	imagetypepb "nebulaid/protos/gen/image/image_type"

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
