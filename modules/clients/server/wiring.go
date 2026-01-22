package server

import (
	"context"
	"fmt"
	"time"

	apiKeyApp "nfxid/modules/clients/application/api_keys"
	resourceApp "nfxid/modules/clients/application/resource"
	appApp "nfxid/modules/clients/application/apps"
	clientCredentialApp "nfxid/modules/clients/application/client_credentials"
	clientScopeApp "nfxid/modules/clients/application/client_scopes"
	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	"nfxid/modules/clients/config"
	apiKeyRepo "nfxid/modules/clients/infrastructure/repository/api_keys"
	appRepo "nfxid/modules/clients/infrastructure/repository/apps"
	clientCredentialRepo "nfxid/modules/clients/infrastructure/repository/client_credentials"
	clientScopeRepo "nfxid/modules/clients/infrastructure/repository/client_scopes"
	ipAllowlistRepo "nfxid/modules/clients/infrastructure/repository/ip_allowlist"
	rateLimitRepo "nfxid/modules/clients/infrastructure/repository/rate_limits"
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
	appAppSvc                *appApp.Service
	apiKeyAppSvc             *apiKeyApp.Service
	clientCredentialAppSvc   *clientCredentialApp.Service
	clientScopeAppSvc        *clientScopeApp.Service
	ipAllowlistAppSvc       *ipAllowlistApp.Service
	rateLimitAppSvc          *rateLimitApp.Service
	userTokenVerifier        token.Verifier
	serverTokenVerifier      token.Verifier
	resourceSvc         *resourceApp.Service
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
	tokenxInstance := tokenx.New(cfg.Token)

	//! === Token Verifiers ===
	// User Token Verifier (用于 HTTP 中间件 - 验证用户 token)
	userTokenVerifier := &tokenxVerifierAdapter{tokenx: tokenxInstance}

	// Server Token Verifier (用于 gRPC 拦截器 - 验证服务间通信 token)
	serverTokenVerifier := servertoken.NewVerifier(
		&servertoken.HMACSigner{Key: []byte(cfg.Token.SecretKey)},
		cfg.Token.Issuer,
		servertoken.WithAllowedSkew(5*time.Second),
	)

	//! === Repository ===
	appRepoInstance := appRepo.NewRepo(postgres.DB())
	apiKeyRepoInstance := apiKeyRepo.NewRepo(postgres.DB())
	clientCredentialRepoInstance := clientCredentialRepo.NewRepo(postgres.DB())
	clientScopeRepoInstance := clientScopeRepo.NewRepo(postgres.DB())
	ipAllowlistRepoInstance := ipAllowlistRepo.NewRepo(postgres.DB())
	rateLimitRepoInstance := rateLimitRepo.NewRepo(postgres.DB())

	//! === Application Services ===
	appAppSvc := appApp.NewService(appRepoInstance)
	apiKeyAppSvc := apiKeyApp.NewService(apiKeyRepoInstance)
	clientCredentialAppSvc := clientCredentialApp.NewService(clientCredentialRepoInstance)
	clientScopeAppSvc := clientScopeApp.NewService(clientScopeRepoInstance)
	ipAllowlistAppSvc := ipAllowlistApp.NewService(ipAllowlistRepoInstance)
	rateLimitAppSvc := rateLimitApp.NewService(rateLimitRepoInstance)

	resourceSvc := resourceApp.NewService(postgres, cacheConn, &kafkaConfig, &rabbitMQConfig)

	return &Dependencies{
		healthMgr:              healthMgr,
		postgres:               postgres,
		cache:                  cacheConn,
		kafkaConfig:            &kafkaConfig,
		busPublisher:           busPublisher,
		rabbitMQConfig:         &rabbitMQConfig,
		appAppSvc:              appAppSvc,
		apiKeyAppSvc:           apiKeyAppSvc,
		clientCredentialAppSvc: clientCredentialAppSvc,
		clientScopeAppSvc:      clientScopeAppSvc,
		ipAllowlistAppSvc:      ipAllowlistAppSvc,
		rateLimitAppSvc:        rateLimitAppSvc,
		userTokenVerifier:      userTokenVerifier,
		serverTokenVerifier:    serverTokenVerifier,
		resourceSvc:         resourceSvc,
		tokenxInstance:         tokenxInstance,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

// Getter methods for interfaces
func (d *Dependencies) HealthMgr() *health.Manager                   { return d.healthMgr }
func (d *Dependencies) ResourceSvc() *resourceApp.Service { return d.resourceSvc }
func (d *Dependencies) AppAppSvc() *appApp.Service                    { return d.appAppSvc }
func (d *Dependencies) APIKeyAppSvc() *apiKeyApp.Service            { return d.apiKeyAppSvc }
func (d *Dependencies) ClientCredentialAppSvc() *clientCredentialApp.Service { return d.clientCredentialAppSvc }
func (d *Dependencies) ClientScopeAppSvc() *clientScopeApp.Service    { return d.clientScopeAppSvc }
func (d *Dependencies) IPAllowlistAppSvc() *ipAllowlistApp.Service    { return d.ipAllowlistAppSvc }
func (d *Dependencies) RateLimitAppSvc() *rateLimitApp.Service        { return d.rateLimitAppSvc }
func (d *Dependencies) UserTokenVerifier() token.Verifier           { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier         { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config                 { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher        { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config            { return d.rabbitMQConfig }

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
