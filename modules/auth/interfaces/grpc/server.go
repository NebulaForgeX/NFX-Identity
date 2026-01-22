package grpc

import (
	sessionApp "nfxid/modules/auth/application/sessions"
	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	passwordResetApp "nfxid/modules/auth/application/password_resets"
	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	resourceApp "nfxid/modules/auth/application/resource"
	grpcHandler "nfxid/modules/auth/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"
	sessionpb "nfxid/protos/gen/auth/session"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"
	mfafactorpb "nfxid/protos/gen/auth/mfa_factor"
	refreshtokenpb "nfxid/protos/gen/auth/refresh_token"
	passwordresetpb "nfxid/protos/gen/auth/password_reset"
	passwordhistorypb "nfxid/protos/gen/auth/password_history"
	loginattemptpb "nfxid/protos/gen/auth/login_attempt"
	accountlockoutpb "nfxid/protos/gen/auth/account_lockout"
	trusteddevicepb "nfxid/protos/gen/auth/trusted_device"

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
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
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
