package grpc

import (
	accountLockoutApp "nfxidentity/modules/auth/application/account_lockouts"
	loginAttemptApp "nfxidentity/modules/auth/application/login_attempts"
	mfaFactorApp "nfxidentity/modules/auth/application/mfa_factors"
	passwordHistoryApp "nfxidentity/modules/auth/application/password_history"
	passwordResetApp "nfxidentity/modules/auth/application/password_resets"
	refreshTokenApp "nfxidentity/modules/auth/application/refresh_tokens"
	resourceApp "nfxidentity/modules/auth/application/resource"
	sessionApp "nfxidentity/modules/auth/application/sessions"
	trustedDeviceApp "nfxidentity/modules/auth/application/trusted_devices"
	userCredentialApp "nfxidentity/modules/auth/application/user_credentials"
	grpcHandler "nfxidentity/modules/auth/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	accountlockoutpb "nfxidentity/protos/gen/auth/account_lockout"
	loginattemptpb "nfxidentity/protos/gen/auth/login_attempt"
	mfafactorpb "nfxidentity/protos/gen/auth/mfa_factor"
	passwordhistorypb "nfxidentity/protos/gen/auth/password_history"
	passwordresetpb "nfxidentity/protos/gen/auth/password_reset"
	refreshtokenpb "nfxidentity/protos/gen/auth/refresh_token"
	sessionpb "nfxidentity/protos/gen/auth/session"
	trusteddevicepb "nfxidentity/protos/gen/auth/trusted_device"
	usercredentialpb "nfxidentity/protos/gen/auth/user_credential"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"

	"google.golang.org/grpc"
)

type Deps interface {
	SessionAppSvc() *sessionApp.Service
	UserCredentialAppSvc() *userCredentialApp.Service
	MFAFactorAppSvc() *mfaFactorApp.Service
	RefreshTokenAppSvc() *refreshTokenApp.Service
	PasswordResetAppSvc() *passwordResetApp.Service
	PasswordHistoryAppSvc() *passwordHistoryApp.Service
	LoginAttemptAppSvc() *loginAttemptApp.Service
	AccountLockoutAppSvc() *accountLockoutApp.Service
	TrustedDeviceAppSvc() *trustedDeviceApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	// 创建 gRPC 服务器，添加认证拦截器（使用 ServerTokenVerifier 用于服务间通信）
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorHandler(),
			servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier()),
		),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	sessionpb.RegisterSessionServiceServer(s, grpcHandler.NewSessionHandler(
		d.SessionAppSvc(),
	))

	usercredentialpb.RegisterUserCredentialServiceServer(s, grpcHandler.NewUserCredentialHandler(
		d.UserCredentialAppSvc(),
	))

	mfafactorpb.RegisterMfaFactorServiceServer(s, grpcHandler.NewMFAFactorHandler(
		d.MFAFactorAppSvc(),
	))

	refreshtokenpb.RegisterRefreshTokenServiceServer(s, grpcHandler.NewRefreshTokenHandler(
		d.RefreshTokenAppSvc(),
	))

	passwordresetpb.RegisterPasswordResetServiceServer(s, grpcHandler.NewPasswordResetHandler(
		d.PasswordResetAppSvc(),
	))

	passwordhistorypb.RegisterPasswordHistoryServiceServer(s, grpcHandler.NewPasswordHistoryHandler(
		d.PasswordHistoryAppSvc(),
	))

	loginattemptpb.RegisterLoginAttemptServiceServer(s, grpcHandler.NewLoginAttemptHandler(
		d.LoginAttemptAppSvc(),
	))

	accountlockoutpb.RegisterAccountLockoutServiceServer(s, grpcHandler.NewAccountLockoutHandler(
		d.AccountLockoutAppSvc(),
	))

	trusteddevicepb.RegisterTrustedDeviceServiceServer(s, grpcHandler.NewTrustedDeviceHandler(
		d.TrustedDeviceAppSvc(),
	))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "auth"))

	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "auth"))

	return s
}
