package grpc

import (
	imageApp "nfxid/modules/image/application/images"
	resourceApp "nfxid/modules/image/application/resource"
	grpcHandler "nfxid/modules/image/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"
	imagepb "nfxid/protos/gen/image/image"
	schemapb "nfxid/protos/gen/common/schema"

	"google.golang.org/grpc"
)

type Deps interface {
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
	ImageAppSvc() *imageApp.Service
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	imagepb.RegisterImageServiceServer(s, grpcHandler.NewImageHandler(d.ImageAppSvc()))
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "image"))
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "image"))

	return s
}
