package server

import (
	"context"
	"fmt"
	"time"

	"nfxid/constants"
	imageTagApp "nfxid/modules/image/application/image_tags"
	imageTypeApp "nfxid/modules/image/application/image_types"
	imageVariantApp "nfxid/modules/image/application/image_variants"
	imageApp "nfxid/modules/image/application/images"
	resourceApp "nfxid/modules/image/application/resource"
	"nfxid/modules/image/config"
	imageTagRepo "nfxid/modules/image/infrastructure/repository/image_tags"
	imageTypeRepo "nfxid/modules/image/infrastructure/repository/image_types"
	imageVariantRepo "nfxid/modules/image/infrastructure/repository/image_variants"
	imageRepo "nfxid/modules/image/infrastructure/repository/images"
	"nfxid/pkgs/cachex"
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
	healthMgr           *health.Manager
	cache               *cachex.Connection
	postgres            *postgresqlx.Connection
	kafkaConfig         *kafkax.Config
	busPublisher        *eventbus.BusPublisher
	rabbitMQConfig      *rabbitmqx.Config
	userTokenVerifier   token.Verifier
	serverTokenVerifier token.Verifier
	resourceSvc         *resourceApp.Service
	tokenxInstance      *tokenx.Tokenx
	imageAppSvc         *imageApp.Service
	imageTypeAppSvc     *imageTypeApp.Service
	imageTagAppSvc      *imageTagApp.Service
	imageVariantAppSvc  *imageVariantApp.Service
	storagePath         string
}

func NewDeps(ctx context.Context, cfg *config.Config) (*Dependencies, error) {
	//! === Initialize Infrastructure ===

	// PostgreSQL Connection
	postgres, err := postgresqlx.Init(ctx, cfg.PostgreSQL)
	if err != nil {
		return nil, fmt.Errorf("init PostgreSQL: %w", err)
	}

	// Redis Cache
	cacheConn, err := cachex.InitConn(ctx, cfg.Cache)
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
	// 使用配置文件中的 token 配置（确保与其他服务一致）
	tokenCfg := cfg.Token
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
	imageRepoInstance := imageRepo.NewRepo(postgres.DB())
	imageTypeRepoInstance := imageTypeRepo.NewRepo(postgres.DB())
	imageTagRepoInstance := imageTagRepo.NewRepo(postgres.DB())
	imageVariantRepoInstance := imageVariantRepo.NewRepo(postgres.DB())

	//! === Application Services ===
	storageBasePath := cfg.Storage.BasePath
	if storageBasePath == "" {
		storageBasePath = constants.StorageBasePath
	}
	imageAppSvc := imageApp.NewService(storageBasePath, imageRepoInstance)
	imageTypeAppSvc := imageTypeApp.NewService(imageTypeRepoInstance)
	imageTagAppSvc := imageTagApp.NewService(imageTagRepoInstance)
	imageVariantAppSvc := imageVariantApp.NewService(imageVariantRepoInstance)
	resourceSvc := resourceApp.NewService(postgres, cacheConn, &kafkaConfig, &rabbitMQConfig)

	return &Dependencies{
		healthMgr:           healthMgr,
		postgres:            postgres,
		cache:               cacheConn,
		kafkaConfig:         &kafkaConfig,
		busPublisher:        busPublisher,
		rabbitMQConfig:      &rabbitMQConfig,
		userTokenVerifier:   userTokenVerifier,
		serverTokenVerifier: serverTokenVerifier,
		resourceSvc:         resourceSvc,
		tokenxInstance:      tokenxInstance,
		imageAppSvc:         imageAppSvc,
		imageTypeAppSvc:     imageTypeAppSvc,
		imageTagAppSvc:      imageTagAppSvc,
		imageVariantAppSvc:  imageVariantAppSvc,
		storagePath:         cfg.Storage.BasePath,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

// Getter methods for interfaces
func (d *Dependencies) HealthMgr() *health.Manager                   { return d.healthMgr }
func (d *Dependencies) ResourceSvc() *resourceApp.Service            { return d.resourceSvc }
func (d *Dependencies) Postgres() *postgresqlx.Connection            { return d.postgres }
func (d *Dependencies) UserTokenVerifier() token.Verifier            { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier          { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config                  { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher         { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config            { return d.rabbitMQConfig }
func (d *Dependencies) ImageAppSvc() *imageApp.Service               { return d.imageAppSvc }
func (d *Dependencies) ImageTypeAppSvc() *imageTypeApp.Service       { return d.imageTypeAppSvc }
func (d *Dependencies) ImageTagAppSvc() *imageTagApp.Service         { return d.imageTagAppSvc }
func (d *Dependencies) ImageVariantAppSvc() *imageVariantApp.Service { return d.imageVariantAppSvc }
func (d *Dependencies) StoragePath() string                          { return d.storagePath }

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
