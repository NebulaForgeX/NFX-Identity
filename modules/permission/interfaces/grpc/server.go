package grpc

import (
	permissionApp "nfxid/modules/permission/application/permission"
	userPermissionApp "nfxid/modules/permission/application/user_permission"
	grpcHandler "nfxid/modules/permission/interfaces/grpc/handler"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token"
	permissionpb "nfxid/protos/gen/permission/permission"
	userpermissionpb "nfxid/protos/gen/permission/user_permission"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

type Deps interface {
	ServerTokenVerifier() token.Verifier
	PermissionAppSvc() *permissionApp.Service
	UserPermissionAppSvc() *userPermissionApp.Service
}

func NewServer(d Deps) *grpc.Server {
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)

	s := grpc.NewServer(grpcx.DefaultServerOptions(d.ServerTokenVerifier())...)

	// Register PermissionService
	permissionpb.RegisterPermissionServiceServer(s, grpcHandler.NewPermissionHandler(
		d.PermissionAppSvc(),
	))

	// Register UserPermissionService
	userpermissionpb.RegisterUserPermissionServiceServer(s, grpcHandler.NewUserPermissionHandler(
		d.UserPermissionAppSvc(),
	))

	return s
}

