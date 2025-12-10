package grpc

import (
	badgeApp "nebulaid/modules/auth/application/badge"
	educationApp "nebulaid/modules/auth/application/education"
	occupationApp "nebulaid/modules/auth/application/occupation"
	profileApp "nebulaid/modules/auth/application/profile"
	profileBadgeApp "nebulaid/modules/auth/application/profile_badge"
	roleApp "nebulaid/modules/auth/application/role"
	userApp "nebulaid/modules/auth/application/user"
	userDomain "nebulaid/modules/auth/domain/user"
	grpcHandler "nebulaid/modules/auth/interfaces/grpc/handler"
	"nebulaid/pkgs/grpcx"
	"nebulaid/pkgs/security/token"
	authpb "nebulaid/protos/gen/auth/auth"
	badgepb "nebulaid/protos/gen/auth/badge"
	educationpb "nebulaid/protos/gen/auth/education"
	occupationpb "nebulaid/protos/gen/auth/occupation"
	profilepb "nebulaid/protos/gen/auth/profile"
	rolepb "nebulaid/protos/gen/auth/role"
	userpb "nebulaid/protos/gen/auth/user"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

type Deps interface {
	ServerTokenVerifier() token.Verifier
	UserAppSvc() *userApp.Service
	ProfileAppSvc() *profileApp.Service
	RoleAppSvc() *roleApp.Service
	BadgeAppSvc() *badgeApp.Service
	EducationAppSvc() *educationApp.Service
	OccupationAppSvc() *occupationApp.Service
	ProfileBadgeAppSvc() *profileBadgeApp.Service
	UserRepo() userDomain.Repo
}

func NewServer(d Deps) *grpc.Server {
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)

	s := grpc.NewServer(grpcx.DefaultServerOptions(d.ServerTokenVerifier())...)

	// Register gRPC services
	authpb.RegisterAuthServiceServer(s, grpcHandler.NewAuthHandler(
		d.UserAppSvc(),
		d.ProfileAppSvc(),
		d.RoleAppSvc(),
		d.UserRepo(),
	))

	badgepb.RegisterBadgeServiceServer(s, grpcHandler.NewBadgeHandler(
		d.BadgeAppSvc(),
	))

	educationpb.RegisterEducationServiceServer(s, grpcHandler.NewEducationHandler(
		d.EducationAppSvc(),
	))

	occupationpb.RegisterOccupationServiceServer(s, grpcHandler.NewOccupationHandler(
		d.OccupationAppSvc(),
	))

	badgepb.RegisterProfileBadgeServiceServer(s, grpcHandler.NewProfileBadgeHandler(
		d.ProfileBadgeAppSvc(),
		d.BadgeAppSvc(),
	))

	// Register UserService
	userpb.RegisterUserServiceServer(s, grpcHandler.NewUserHandler(
		d.UserAppSvc(),
		d.UserRepo(),
	))

	// Register ProfileService
	profilepb.RegisterProfileServiceServer(s, grpcHandler.NewProfileHandler(
		d.ProfileAppSvc(),
		d.ProfileBadgeAppSvc(),
		d.BadgeAppSvc(),
	))

	// Register RoleService
	rolepb.RegisterRoleServiceServer(s, grpcHandler.NewRoleHandler(
		d.RoleAppSvc(),
	))

	return s
}
