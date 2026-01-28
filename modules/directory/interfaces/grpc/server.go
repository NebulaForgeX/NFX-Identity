package grpc

import (
	badgeApp "nfxid/modules/directory/application/badges"
	resourceApp "nfxid/modules/directory/application/resource"
	userAvatarApp "nfxid/modules/directory/application/user_avatars"
	userBadgeApp "nfxid/modules/directory/application/user_badges"
	userEducationApp "nfxid/modules/directory/application/user_educations"
	userEmailApp "nfxid/modules/directory/application/user_emails"
	userImageApp "nfxid/modules/directory/application/user_images"
	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	userPhoneApp "nfxid/modules/directory/application/user_phones"
	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	userProfileApp "nfxid/modules/directory/application/user_profiles"
	userApp "nfxid/modules/directory/application/users"
	grpcHandler "nfxid/modules/directory/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"
	badgepb "nfxid/protos/gen/directory/badge"
	userpb "nfxid/protos/gen/directory/user"
	useravatarpb "nfxid/protos/gen/directory/user_avatar"
	userbadgepb "nfxid/protos/gen/directory/user_badge"
	usereducationpb "nfxid/protos/gen/directory/user_education"
	useremailpb "nfxid/protos/gen/directory/user_email"
	userimagepb "nfxid/protos/gen/directory/user_image"
	useroccupationpb "nfxid/protos/gen/directory/user_occupation"
	userphonepb "nfxid/protos/gen/directory/user_phone"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"
	userprofilepb "nfxid/protos/gen/directory/user_profile"

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
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
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
