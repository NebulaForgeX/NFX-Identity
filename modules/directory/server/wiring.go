package server

import (
	"context"
	"fmt"
	"time"

	badgeApp "nfxid/modules/directory/application/badges"
	userApp "nfxid/modules/directory/application/users"
	userBadgeApp "nfxid/modules/directory/application/user_badges"
	userEducationApp "nfxid/modules/directory/application/user_educations"
	userEmailApp "nfxid/modules/directory/application/user_emails"
	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	userPhoneApp "nfxid/modules/directory/application/user_phones"
	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	userProfileApp "nfxid/modules/directory/application/user_profiles"
	"nfxid/modules/directory/config"
	badgeRepo "nfxid/modules/directory/infrastructure/repository/badges"
	userRepo "nfxid/modules/directory/infrastructure/repository/users"
	userBadgeRepo "nfxid/modules/directory/infrastructure/repository/user_badges"
	userEducationRepo "nfxid/modules/directory/infrastructure/repository/user_educations"
	userEmailRepo "nfxid/modules/directory/infrastructure/repository/user_emails"
	userOccupationRepo "nfxid/modules/directory/infrastructure/repository/user_occupations"
	userPhoneRepo "nfxid/modules/directory/infrastructure/repository/user_phones"
	userPreferenceRepo "nfxid/modules/directory/infrastructure/repository/user_preferences"
	userProfileRepo "nfxid/modules/directory/infrastructure/repository/user_profiles"
	"nfxid/pkgs/cache"
	"nfxid/pkgs/health"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/kafkax/eventbus"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/rabbitmqx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"
)

type Dependencies struct {
	healthMgr                *health.Manager
	cache                    *cache.Connection
	postgres                 *postgresqlx.Connection
	kafkaConfig              *kafkax.Config
	busPublisher             *eventbus.BusPublisher
	rabbitMQConfig           *rabbitmqx.Config
	userAppSvc               *userApp.Service
	badgeAppSvc              *badgeApp.Service
	userBadgeAppSvc          *userBadgeApp.Service
	userEducationAppSvc      *userEducationApp.Service
	userEmailAppSvc          *userEmailApp.Service
	userOccupationAppSvc     *userOccupationApp.Service
	userPhoneAppSvc          *userPhoneApp.Service
	userPreferenceAppSvc     *userPreferenceApp.Service
	userProfileAppSvc        *userProfileApp.Service
	userTokenVerifier        token.Verifier
	serverTokenVerifier      token.Verifier
	tokenxInstance           *tokenx.Tokenx
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
	// 使用默认 token 配置（directory 模块可能不需要 token 生成，只需要验证）
	tokenCfg := tokenx.DefaultConfig()
	tokenxInstance := tokenx.New(tokenCfg)

	//! === Token Verifiers ===
	// User Token Verifier (用于 HTTP 中间件 - 验证用户 token)
	userTokenVerifier := &tokenxVerifierAdapter{tokenx: tokenxInstance}

	// Server Token Verifier (用于 gRPC 拦截器 - 验证服务间通信 token)
	serverTokenVerifier := servertoken.NewVerifier(
		&servertoken.HMACSigner{Key: []byte(tokenCfg.SecretKey)},
		tokenCfg.Issuer,
		servertoken.WithAllowedSkew(5*time.Second),
	)

	//! === Repository ===
	userRepoInstance := userRepo.NewRepo(postgres.DB())
	badgeRepoInstance := badgeRepo.NewRepo(postgres.DB())
	userBadgeRepoInstance := userBadgeRepo.NewRepo(postgres.DB())
	userEducationRepoInstance := userEducationRepo.NewRepo(postgres.DB())
	userEmailRepoInstance := userEmailRepo.NewRepo(postgres.DB())
	userOccupationRepoInstance := userOccupationRepo.NewRepo(postgres.DB())
	userPhoneRepoInstance := userPhoneRepo.NewRepo(postgres.DB())
	userPreferenceRepoInstance := userPreferenceRepo.NewRepo(postgres.DB())
	userProfileRepoInstance := userProfileRepo.NewRepo(postgres.DB())

	//! === Application Services ===
	userAppSvc := userApp.NewService(userRepoInstance)
	badgeAppSvc := badgeApp.NewService(badgeRepoInstance)
	userBadgeAppSvc := userBadgeApp.NewService(userBadgeRepoInstance)
	userEducationAppSvc := userEducationApp.NewService(userEducationRepoInstance)
	userEmailAppSvc := userEmailApp.NewService(userEmailRepoInstance)
	userOccupationAppSvc := userOccupationApp.NewService(userOccupationRepoInstance)
	userPhoneAppSvc := userPhoneApp.NewService(userPhoneRepoInstance)
	userPreferenceAppSvc := userPreferenceApp.NewService(userPreferenceRepoInstance)
	userProfileAppSvc := userProfileApp.NewService(userProfileRepoInstance)

	return &Dependencies{
		healthMgr:            healthMgr,
		postgres:             postgres,
		cache:                cacheConn,
		kafkaConfig:          &kafkaConfig,
		busPublisher:         busPublisher,
		rabbitMQConfig:       &rabbitMQConfig,
		userAppSvc:           userAppSvc,
		badgeAppSvc:          badgeAppSvc,
		userBadgeAppSvc:      userBadgeAppSvc,
		userEducationAppSvc:  userEducationAppSvc,
		userEmailAppSvc:      userEmailAppSvc,
		userOccupationAppSvc: userOccupationAppSvc,
		userPhoneAppSvc:      userPhoneAppSvc,
		userPreferenceAppSvc: userPreferenceAppSvc,
		userProfileAppSvc:    userProfileAppSvc,
		userTokenVerifier:    userTokenVerifier,
		serverTokenVerifier:  serverTokenVerifier,
		tokenxInstance:       tokenxInstance,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

// Getter methods for interfaces
func (d *Dependencies) UserAppSvc() *userApp.Service                    { return d.userAppSvc }
func (d *Dependencies) BadgeAppSvc() *badgeApp.Service                 { return d.badgeAppSvc }
func (d *Dependencies) UserBadgeAppSvc() *userBadgeApp.Service          { return d.userBadgeAppSvc }
func (d *Dependencies) UserEducationAppSvc() *userEducationApp.Service  { return d.userEducationAppSvc }
func (d *Dependencies) UserEmailAppSvc() *userEmailApp.Service          { return d.userEmailAppSvc }
func (d *Dependencies) UserOccupationAppSvc() *userOccupationApp.Service { return d.userOccupationAppSvc }
func (d *Dependencies) UserPhoneAppSvc() *userPhoneApp.Service         { return d.userPhoneAppSvc }
func (d *Dependencies) UserPreferenceAppSvc() *userPreferenceApp.Service { return d.userPreferenceAppSvc }
func (d *Dependencies) UserProfileAppSvc() *userProfileApp.Service     { return d.userProfileAppSvc }
func (d *Dependencies) UserTokenVerifier() token.Verifier              { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier             { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config                    { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher          { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config             { return d.rabbitMQConfig }

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
			"user_id":  claims.UserID,
			"username": claims.Username,
			"email":    claims.Email,
			"phone":    claims.Phone,
			"role_id":  claims.RoleID,
			"type":     claims.Type,
		},
	}, nil
}
