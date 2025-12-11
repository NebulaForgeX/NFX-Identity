package grpc

import (
	badgeApp "nfxid/modules/auth/application/badge"
	educationApp "nfxid/modules/auth/application/profile_education"
	occupationApp "nfxid/modules/auth/application/profile_occupation"
	profileApp "nfxid/modules/auth/application/profile"
	profileBadgeApp "nfxid/modules/auth/application/profile_badge"
	roleApp "nfxid/modules/auth/application/role"
	userApp "nfxid/modules/auth/application/user"
	userDomain "nfxid/modules/auth/domain/user"
	grpcHandler "nfxid/modules/auth/interfaces/grpc/handler"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token"
	authpb "nfxid/protos/gen/auth/auth"
	badgepb "nfxid/protos/gen/auth/badge"
	profileeducationpb "nfxid/protos/gen/auth/profile_education"
	profileoccupationpb "nfxid/protos/gen/auth/profile_occupation"
	profilepb "nfxid/protos/gen/auth/profile"
	rolepb "nfxid/protos/gen/auth/role"
	userpb "nfxid/protos/gen/auth/user"

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

	profileeducationpb.RegisterProfileEducationServiceServer(s, grpcHandler.NewEducationHandler(
		d.EducationAppSvc(),
	))

	profileoccupationpb.RegisterProfileOccupationServiceServer(s, grpcHandler.NewOccupationHandler(
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
