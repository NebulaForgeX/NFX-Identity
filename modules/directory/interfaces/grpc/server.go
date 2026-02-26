package grpc

import (
	badgeApp "nfxidentity/modules/directory/application/badges"
	resourceApp "nfxidentity/modules/directory/application/resource"
	userAvatarApp "nfxidentity/modules/directory/application/user_avatars"
	userBadgeApp "nfxidentity/modules/directory/application/user_badges"
	userEducationApp "nfxidentity/modules/directory/application/user_educations"
	userEmailApp "nfxidentity/modules/directory/application/user_emails"
	userImageApp "nfxidentity/modules/directory/application/user_images"
	userOccupationApp "nfxidentity/modules/directory/application/user_occupations"
	userPhoneApp "nfxidentity/modules/directory/application/user_phones"
	userPreferenceApp "nfxidentity/modules/directory/application/user_preferences"
	userProfileApp "nfxidentity/modules/directory/application/user_profiles"
	userApp "nfxidentity/modules/directory/application/users"
	grpcHandler "nfxidentity/modules/directory/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"
	badgepb "nfxidentity/protos/gen/directory/badge"
	userpb "nfxidentity/protos/gen/directory/user"
	useravatarpb "nfxidentity/protos/gen/directory/user_avatar"
	userbadgepb "nfxidentity/protos/gen/directory/user_badge"
	usereducationpb "nfxidentity/protos/gen/directory/user_education"
	useremailpb "nfxidentity/protos/gen/directory/user_email"
	userimagepb "nfxidentity/protos/gen/directory/user_image"
	useroccupationpb "nfxidentity/protos/gen/directory/user_occupation"
	userphonepb "nfxidentity/protos/gen/directory/user_phone"
	userpreferencepb "nfxidentity/protos/gen/directory/user_preference"
	userprofilepb "nfxidentity/protos/gen/directory/user_profile"

	"google.golang.org/grpc"
)

type Deps interface {
	UserAppSvc() *userApp.Service
	BadgeAppSvc() *badgeApp.Service
	UserEmailAppSvc() *userEmailApp.Service
	UserBadgeAppSvc() *userBadgeApp.Service
	UserEducationAppSvc() *userEducationApp.Service
	UserOccupationAppSvc() *userOccupationApp.Service
	UserPhoneAppSvc() *userPhoneApp.Service
	UserPreferenceAppSvc() *userPreferenceApp.Service
	UserProfileAppSvc() *userProfileApp.Service
	UserAvatarAppSvc() *userAvatarApp.Service
	UserImageAppSvc() *userImageApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorHandler(),
			servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier()),
		),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	userpb.RegisterUserServiceServer(s, grpcHandler.NewUserHandler(d.UserAppSvc()))
	badgepb.RegisterBadgeServiceServer(s, grpcHandler.NewBadgeHandler(d.BadgeAppSvc()))
	useremailpb.RegisterUserEmailServiceServer(s, grpcHandler.NewUserEmailHandler(d.UserEmailAppSvc()))
	useravatarpb.RegisterUserAvatarServiceServer(s, grpcHandler.NewUserAvatarHandler(d.UserAvatarAppSvc()))
	userimagepb.RegisterUserImageServiceServer(s, grpcHandler.NewUserImageHandler(d.UserImageAppSvc()))
	userbadgepb.RegisterUserBadgeServiceServer(s, grpcHandler.NewUserBadgeHandler(d.UserBadgeAppSvc()))
	usereducationpb.RegisterUserEducationServiceServer(s, grpcHandler.NewUserEducationHandler(d.UserEducationAppSvc()))
	useroccupationpb.RegisterUserOccupationServiceServer(s, grpcHandler.NewUserOccupationHandler(d.UserOccupationAppSvc()))
	userphonepb.RegisterUserPhoneServiceServer(s, grpcHandler.NewUserPhoneHandler(d.UserPhoneAppSvc()))
	userpreferencepb.RegisterUserPreferenceServiceServer(s, grpcHandler.NewUserPreferenceHandler(d.UserPreferenceAppSvc()))
	userprofilepb.RegisterUserProfileServiceServer(s, grpcHandler.NewUserProfileHandler(d.UserProfileAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "directory"))

	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "directory"))

	return s
}
