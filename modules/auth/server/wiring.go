package server

import (
	"context"
	"fmt"
	"time"

	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	authApp "nfxid/modules/auth/application/auth"
	resourceApp "nfxid/modules/auth/application/resource"
	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	passwordResetApp "nfxid/modules/auth/application/password_resets"
	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	sessionApp "nfxid/modules/auth/application/sessions"
	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	"nfxid/modules/auth/config"
	authInfra "nfxid/modules/auth/infrastructure/auth"
	authGrpc "nfxid/modules/auth/infrastructure/grpc"
	accountLockoutRepo "nfxid/modules/auth/infrastructure/repository/account_lockouts"
	loginAttemptRepo "nfxid/modules/auth/infrastructure/repository/login_attempts"
	mfaFactorRepo "nfxid/modules/auth/infrastructure/repository/mfa_factors"
	passwordHistoryRepo "nfxid/modules/auth/infrastructure/repository/password_history"
	passwordResetRepo "nfxid/modules/auth/infrastructure/repository/password_resets"
	refreshTokenRepo "nfxid/modules/auth/infrastructure/repository/refresh_tokens"
	sessionRepo "nfxid/modules/auth/infrastructure/repository/sessions"
	trustedDeviceRepo "nfxid/modules/auth/infrastructure/repository/trusted_devices"
	userCredentialRepo "nfxid/modules/auth/infrastructure/repository/user_credentials"
	"nfxid/pkgs/cache"
	"nfxid/pkgs/health"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/kafkax/eventbus"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/rabbitmqx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"
	"nfxid/pkgs/email"
)

type Dependencies struct {
	healthMgr              *health.Manager
	cache                  *cache.Connection
	postgres               *postgresqlx.Connection
	kafkaConfig            *kafkax.Config
	busPublisher           *eventbus.BusPublisher
	rabbitMQConfig         *rabbitmqx.Config
	sessionAppSvc          *sessionApp.Service
	userCredentialAppSvc   *userCredentialApp.Service
	mfaFactorAppSvc        *mfaFactorApp.Service
	refreshTokenAppSvc     *refreshTokenApp.Service
	passwordResetAppSvc    *passwordResetApp.Service
	passwordHistoryAppSvc  *passwordHistoryApp.Service
	loginAttemptAppSvc     *loginAttemptApp.Service
	accountLockoutAppSvc   *accountLockoutApp.Service
	trustedDeviceAppSvc    *trustedDeviceApp.Service
	userTokenVerifier      token.Verifier // 用于 HTTP 中间件（用户 token）
	serverTokenVerifier    token.Verifier // 用于 gRPC 拦截器（服务间通信）
	resourceSvc            *resourceApp.Service
	tokenxInstance         *tokenx.Tokenx
	authAppSvc             *authApp.Service
	grpcClients            *authGrpc.GRPCClients
	emailService           *email.EmailService
}

func NewDeps(ctx context.Context, cfg *config.Config) (*Dependencies, error) {
	//! === Initialize Infrastructure ===

	// PostgreSQL Connection
	postgres, err := postgresqlx.Init(ctx, cfg.PostgreSQL)
	if err != nil {
		return nil, fmt.Errorf("init PostgreSQL: %w", err)
	}

	// Redis Cache
	cacheConn, err := cache.InitConn(ctx, cfg.Cache)
	if err != nil {
		return nil, fmt.Errorf("init Redis: %w", err)
	}

	//! === Health Manager ===
	healthMgr := health.NewManager(ctx, 30*time.Second)
	healthMgr.Register(postgres)
	healthMgr.Register(cacheConn)

	//! === Kafka Service ===
	kafkaConfig := cfg.KafkaConfig
	busPublisher, err := kafkax.NewPublisher(&kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka publisher: %w", err)
	}

	//! === RabbitMQ Config ===
	rabbitMQConfig := cfg.RabbitMQConfig

	//! === Tokenx ===
	tokenxInstance := tokenx.New(cfg.Token)

	//! === Email Service ===
	emailService := email.NewEmailService(email.SMTPConfig{
		Host:     cfg.Email.SMTPHost,
		Port:     cfg.Email.SMTPPort,
		Username: cfg.Email.SMTPUser,
		Password: cfg.Email.SMTPPassword,
		From:     cfg.Email.SMTPFrom,
	})

	//! === Token Verifiers ===
	// User Token Verifier (用于 HTTP 中间件 - 验证用户 token)
	// 使用 tokenx adapter 将 tokenx.Tokenx 适配为 token.Verifier 接口
	userTokenVerifier := &tokenxVerifierAdapter{tokenx: tokenxInstance}
	tokenIssuer := authInfra.NewTokenIssuer(tokenxInstance)
	// Server Token Verifier (用于 gRPC 拦截器 - 验证服务间通信 token)
	serverTokenVerifier := servertoken.NewVerifier(
		&servertoken.HMACSigner{Key: []byte(cfg.Token.SecretKey)},
		cfg.Token.Issuer,
		servertoken.WithAllowedSkew(5*time.Second),
	)

	//! === Repository ===
	sessionRepoInstance := sessionRepo.NewRepo(postgres.DB())
	userCredentialRepoInstance := userCredentialRepo.NewRepo(postgres.DB())
	mfaFactorRepoInstance := mfaFactorRepo.NewRepo(postgres.DB())
	refreshTokenRepoInstance := refreshTokenRepo.NewRepo(postgres.DB())
	passwordResetRepoInstance := passwordResetRepo.NewRepo(postgres.DB())
	passwordHistoryRepoInstance := passwordHistoryRepo.NewRepo(postgres.DB())
	loginAttemptRepoInstance := loginAttemptRepo.NewRepo(postgres.DB())
	accountLockoutRepoInstance := accountLockoutRepo.NewRepo(postgres.DB())
	trustedDeviceRepoInstance := trustedDeviceRepo.NewRepo(postgres.DB())



	//! === gRPC Clients ===
	grpcClientsInstance, err := authGrpc.NewGRPCClients(ctx, &cfg.GRPCClient, &cfg.Server, &cfg.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC clients: %w", err)
	}

	//! === Application Services ===
	sessionAppSvc := sessionApp.NewService(sessionRepoInstance)
	userCredentialAppSvc := userCredentialApp.NewService(userCredentialRepoInstance)
	mfaFactorAppSvc := mfaFactorApp.NewService(mfaFactorRepoInstance)
	refreshTokenAppSvc := refreshTokenApp.NewService(refreshTokenRepoInstance)
	passwordResetAppSvc := passwordResetApp.NewService(passwordResetRepoInstance)
	passwordHistoryAppSvc := passwordHistoryApp.NewService(passwordHistoryRepoInstance)
	loginAttemptAppSvc := loginAttemptApp.NewService(loginAttemptRepoInstance)
	accountLockoutAppSvc := accountLockoutApp.NewService(accountLockoutRepoInstance)
	trustedDeviceAppSvc := trustedDeviceApp.NewService(trustedDeviceRepoInstance)
	resourceSvc := resourceApp.NewService(postgres, cacheConn, &kafkaConfig, &rabbitMQConfig)
	authAppSvc := authApp.NewService(
			userCredentialRepoInstance,
			loginAttemptRepoInstance,
			accountLockoutRepoInstance,
			refreshTokenRepoInstance,
			grpcClientsInstance,
			tokenIssuer,
			int64(cfg.Token.AccessTokenTTL.Seconds()),
			int64(cfg.Token.RefreshTokenTTL.Seconds()),
			emailService,
			cacheConn,
			userCredentialAppSvc,
		)

	return &Dependencies{
		healthMgr:             healthMgr,
		postgres:              postgres,
		cache:                 cacheConn,
		kafkaConfig:           &kafkaConfig,
		busPublisher:          busPublisher,
		rabbitMQConfig:        &rabbitMQConfig,
		sessionAppSvc:         sessionAppSvc,
		userCredentialAppSvc:  userCredentialAppSvc,
		mfaFactorAppSvc:       mfaFactorAppSvc,
		refreshTokenAppSvc:    refreshTokenAppSvc,
		passwordResetAppSvc:   passwordResetAppSvc,
		passwordHistoryAppSvc: passwordHistoryAppSvc,
		loginAttemptAppSvc:    loginAttemptAppSvc,
		accountLockoutAppSvc:  accountLockoutAppSvc,
		trustedDeviceAppSvc:   trustedDeviceAppSvc,
		userTokenVerifier:     userTokenVerifier,
		serverTokenVerifier:    serverTokenVerifier,
		resourceSvc:         resourceSvc,
		tokenxInstance:      tokenxInstance,
		authAppSvc:          authAppSvc,
		grpcClients:         grpcClientsInstance,
		emailService:        emailService,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
	if d.grpcClients != nil {
		_ = d.grpcClients.Close()
	}
}

// Getter methods for interfaces
func (d *Dependencies) SessionAppSvc() *sessionApp.Service              { return d.sessionAppSvc }
func (d *Dependencies) UserCredentialAppSvc() *userCredentialApp.Service { return d.userCredentialAppSvc }
func (d *Dependencies) MFAFactorAppSvc() *mfaFactorApp.Service          { return d.mfaFactorAppSvc }
func (d *Dependencies) RefreshTokenAppSvc() *refreshTokenApp.Service   { return d.refreshTokenAppSvc }
func (d *Dependencies) PasswordResetAppSvc() *passwordResetApp.Service   { return d.passwordResetAppSvc }
func (d *Dependencies) PasswordHistoryAppSvc() *passwordHistoryApp.Service { return d.passwordHistoryAppSvc }
func (d *Dependencies) LoginAttemptAppSvc() *loginAttemptApp.Service     { return d.loginAttemptAppSvc }
func (d *Dependencies) AccountLockoutAppSvc() *accountLockoutApp.Service { return d.accountLockoutAppSvc }
func (d *Dependencies) TrustedDeviceAppSvc() *trustedDeviceApp.Service    { return d.trustedDeviceAppSvc }
func (d *Dependencies) HealthMgr() *health.Manager                       { return d.healthMgr }
func (d *Dependencies) ResourceSvc() *resourceApp.Service { return d.resourceSvc }
func (d *Dependencies) Postgres() *postgresqlx.Connection         { return d.postgres }
func (d *Dependencies) UserTokenVerifier() token.Verifier               { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier            { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config                       { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher             { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config                { return d.rabbitMQConfig }
func (d *Dependencies) AuthAppSvc() *authApp.Service                     { return d.authAppSvc }
func (d *Dependencies) EmailService() *email.EmailService                { return d.emailService }
func (d *Dependencies) Cache() *cache.Connection                         { return d.cache }

// tokenxVerifierAdapter 将 tokenx.Tokenx 适配为 token.Verifier 接口
type tokenxVerifierAdapter struct {
	tokenx *tokenx.Tokenx
}

func (a *tokenxVerifierAdapter) Verify(ctx context.Context, tokenStr string) (*token.Claims, error) {
	claims, err := a.tokenx.VerifyAccessToken(tokenStr)
	if err != nil {
		return nil, err
	}

	// 转换为 security/token.Claims
	return &token.Claims{
		Registered: claims.RegisteredClaims,
		Raw: map[string]any{
			"user_id":      claims.UserID,
			"username":     claims.Username,
			"email":        claims.Email,
			"phone":        claims.Phone,
			"country_code": claims.CountryCode,
			"role_id":      claims.RoleID,
			"type":         claims.Type,
		},
	}, nil
}
