package server

import (
	"context"
	"fmt"
	"time"

	badgeApp "nfxid/modules/auth/application/badge"
	badgeAppViews "nfxid/modules/auth/application/badge/views"
	educationApp "nfxid/modules/auth/application/education"
	educationAppViews "nfxid/modules/auth/application/education/views"
	occupationApp "nfxid/modules/auth/application/occupation"
	profileApp "nfxid/modules/auth/application/profile"
	profileBadgeApp "nfxid/modules/auth/application/profile_badge"
	roleApp "nfxid/modules/auth/application/role"
	userApp "nfxid/modules/auth/application/user"
	"nfxid/modules/auth/config"
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/infrastructure/grpcclient"
	"nfxid/modules/auth/infrastructure/query"
	"nfxid/modules/auth/infrastructure/repository"
	"nfxid/pkgs/cache"
	"nfxid/pkgs/cleanup"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/health"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/mongodbx"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"

	"github.com/google/uuid"
)

type Dependencies struct {
	healthMgr          *health.Manager
	cache              *cache.Connection
	postgres           *postgresqlx.Connection
	mongo              *mongodbx.Client
	kafkaConfig        *kafkax.Config
	busPublisher       *eventbus.BusPublisher
	serverVerifier     token.Verifier
	userRepo           userDomain.Repo
	userAppSvc         *userApp.Service
	profileAppSvc      *profileApp.Service
	roleAppSvc         *roleApp.Service
	badgeAppSvc        *badgeApp.Service
	educationAppSvc    *educationApp.Service
	occupationAppSvc   *occupationApp.Service
	profileBadgeAppSvc *profileBadgeApp.Service
	tokenx             *tokenx.Tokenx
	imageGRPCClient    *grpcclient.ImageGRPCClient
}

func NewDependencies(ctx context.Context, cfg *config.Config) (*Dependencies, error) {
	// === Initialize Infrastructure ===

	// PostgreSQL Connection
	postgres, err := postgresqlx.Init(ctx, cfg.PostgreSQL, []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("init PostgreSQL: %w", err)
	}

	// MongoDB Connection
	mongoClient, err := mongodbx.Init(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("init MongoDB: %w", err)
	}

	// Redis Cache
	cacheConn, err := cache.InitConn(ctx, cfg.Cache)
	if err != nil {
		return nil, fmt.Errorf("init Redis: %w", err)
	}

	// === Health Manager ===
	healthMgr := health.NewManager(ctx, 30*time.Second)
	healthMgr.Register(postgres)
	healthMgr.Register(mongoClient)
	healthMgr.Register(cacheConn)

	// === Tokenx ===
	tokenxConfig, err := cfg.Token.ToTokenxConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to parse token config: %w", err)
	}
	tokenxInstance := tokenx.New(tokenxConfig)

	// === Token Verifier ===
	serverVerifier := servertoken.NewVerifier(
		&servertoken.HMACSigner{Key: []byte(tokenxConfig.SecretKey)},
		tokenxConfig.Issuer,
		servertoken.WithAllowedSkew(5*time.Second),
	)

	// === Kafka Service ===
	kafkaConfig := cfg.KafkaConfig
	busPublisher, err := kafkax.NewPublisher(&kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka publisher: %w", err)
	}

	// === Repository ===
	userRepo := repository.NewUserPGRepo(postgres.DB())
	profileRepo := repository.NewProfilePGRepo(postgres.DB())
	roleRepo := repository.NewRolePGRepo(postgres.DB())
	badgeRepo := repository.NewBadgePGRepo(postgres.DB())
	educationRepo := repository.NewEducationPGRepo(postgres.DB())
	occupationRepo := repository.NewOccupationPGRepo(postgres.DB())
	profileBadgeRepo := repository.NewProfileBadgePGRepo(postgres.DB())

	// === Query ===
	userQuery := query.NewUserPGQuery(postgres.DB())
	profileQuery := query.NewProfilePGQuery(postgres.DB())
	roleQuery := query.NewRolePGQuery(postgres.DB())
	badgeQuery := query.NewBadgePGQuery(postgres.DB())
	educationQuery := query.NewEducationPGQuery(postgres.DB())
	occupationQuery := query.NewOccupationPGQuery(postgres.DB())
	profileBadgeQuery := query.NewProfileBadgePGQuery(postgres.DB())

	// === Cache ===
	cacheNS := cfg.Env.String()
	baseCache := cache.NewBaseCache(cacheConn.Client(), &cache.JSONCodec{})
	badgeCacheSet := cache.NewCacheSet[badgeAppViews.BadgeView, uuid.UUID](
		baseCache,
		"auth",
		cacheNS,
		cache.WithEntity(cache.CacheConfig{TTL: 5 * time.Minute}),
	)
	educationCacheSet := cache.NewCacheSet[educationAppViews.EducationView, uuid.UUID](
		baseCache,
		"auth",
		cacheNS,
		cache.WithEntity(cache.CacheConfig{TTL: 5 * time.Minute}),
	)

	// === gRPC Client ===
	var imageGRPCClient *grpcclient.ImageGRPCClient
	if cfg.GRPCClient.ImageAddr != "" {
		// 创建 server token provider（用于服务间 gRPC 认证）
		tokenxConfig, _ := cfg.Token.ToTokenxConfig()
		signer := &servertoken.HMACSigner{Key: []byte(tokenxConfig.SecretKey)}
		serverTokenProvider := servertoken.NewProvider(
			signer,
			tokenxConfig.Issuer,
			"auth-service",
			servertoken.WithTTL(1*time.Hour),
			servertoken.WithMargin(10*time.Second),
		)

		imageGRPCClient, err = grpcclient.NewImageGRPCClient(cfg.GRPCClient.ImageAddr, serverTokenProvider)
		if err != nil {
			return nil, fmt.Errorf("init image grpc client: %w", err)
		}
	}

	// === Application Services ===
	userAppSvc := userApp.NewService(
		userRepo,
		profileRepo,
		userQuery,
		busPublisher,
		tokenxInstance,
	)

	profileAppSvc := profileApp.NewService(
		profileRepo,
		profileQuery,
		imageGRPCClient, // 注入 image gRPC client
	)

	roleAppSvc := roleApp.NewService(
		roleRepo,
		roleQuery,
	)

	badgeAppSvc := badgeApp.NewService(
		badgeRepo,
		badgeQuery,
		badgeCacheSet,
	)

	educationAppSvc := educationApp.NewService(
		educationRepo,
		educationQuery,
		busPublisher,
		educationCacheSet,
	)

	occupationAppSvc := occupationApp.NewService(
		occupationRepo,
		occupationQuery,
	)

	profileBadgeAppSvc := profileBadgeApp.NewService(
		profileBadgeRepo,
		profileBadgeQuery,
	)

	return &Dependencies{
		healthMgr:          healthMgr,
		postgres:           postgres,
		mongo:              mongoClient,
		cache:              cacheConn,
		kafkaConfig:        &kafkaConfig,
		busPublisher:       busPublisher,
		serverVerifier:     serverVerifier,
		userRepo:           userRepo,
		userAppSvc:         userAppSvc,
		profileAppSvc:      profileAppSvc,
		roleAppSvc:         roleAppSvc,
		badgeAppSvc:        badgeAppSvc,
		educationAppSvc:    educationAppSvc,
		occupationAppSvc:   occupationAppSvc,
		profileBadgeAppSvc: profileBadgeAppSvc,
		tokenx:             tokenxInstance,
		imageGRPCClient:    imageGRPCClient,
	}, nil
}

func (d *Dependencies) Cleanup() {
	var cleanupList []any
	cleanupList = append(cleanupList, d.healthMgr, d.postgres, d.mongo, d.cache, d.busPublisher)
	if d.imageGRPCClient != nil {
		cleanupList = append(cleanupList, d.imageGRPCClient)
	}
	if err := cleanup.CleanupAll(cleanupList...); err != nil {
		logx.S().Errorf("cleanup auth service: %v", err)
	}
}

func (d *Dependencies) UserAppSvc() *userApp.Service                 { return d.userAppSvc }
func (d *Dependencies) ProfileAppSvc() *profileApp.Service           { return d.profileAppSvc }
func (d *Dependencies) RoleAppSvc() *roleApp.Service                 { return d.roleAppSvc }
func (d *Dependencies) BadgeAppSvc() *badgeApp.Service               { return d.badgeAppSvc }
func (d *Dependencies) EducationAppSvc() *educationApp.Service       { return d.educationAppSvc }
func (d *Dependencies) OccupationAppSvc() *occupationApp.Service     { return d.occupationAppSvc }
func (d *Dependencies) ProfileBadgeAppSvc() *profileBadgeApp.Service { return d.profileBadgeAppSvc }
func (d *Dependencies) Tokenx() *tokenx.Tokenx                       { return d.tokenx }
func (d *Dependencies) KafkaConfig() *kafkax.Config                  { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher         { return d.busPublisher }
func (d *Dependencies) UserRepo() userDomain.Repo                    { return d.userRepo }
func (d *Dependencies) ServerTokenVerifier() token.Verifier          { return d.serverVerifier }
