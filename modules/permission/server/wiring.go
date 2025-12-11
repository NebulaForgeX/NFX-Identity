package server

import (
	"context"
	"fmt"
	"time"

	authApp "nfxid/modules/permission/application/auth"
	permissionApp "nfxid/modules/permission/application/permission"
	userPermissionApp "nfxid/modules/permission/application/user_permission"
	"nfxid/modules/permission/config"
	permissionDomain "nfxid/modules/permission/domain/permission"
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
	"nfxid/modules/permission/infrastructure/grpcclient"
	"nfxid/modules/permission/infrastructure/query"
	"nfxid/modules/permission/infrastructure/repository"
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
)

type Dependencies struct {
	healthMgr            *health.Manager
	cache                *cache.Connection
	postgres             *postgresqlx.Connection
	mongo                *mongodbx.Client
	kafkaConfig          *kafkax.Config
	busPublisher         *eventbus.BusPublisher
	serverVerifier       token.Verifier
	permissionRepo       permissionDomain.Repo
	userPermissionRepo   userPermissionDomain.Repo
	authAppSvc           *authApp.Service
	permissionAppSvc     *permissionApp.Service
	userPermissionAppSvc *userPermissionApp.Service
	tokenx               *tokenx.Tokenx
	authGRPCClient       *grpcclient.AuthGRPCClient
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
	permissionRepo := repository.NewPermissionPGRepo(postgres.DB())
	userPermissionRepo := repository.NewUserPermissionPGRepo(postgres.DB())

	// === Query ===
	permissionQuery := query.NewPermissionPGQuery(postgres.DB())
	userPermissionQuery := query.NewUserPermissionPGQuery(postgres.DB())

	// === gRPC Client ===
	var authGRPCClient *grpcclient.AuthGRPCClient
	if cfg.GRPCClient.AuthAddr != "" {
		// 创建 server token provider（用于服务间 gRPC 认证）
		tokenxConfig, _ := cfg.Token.ToTokenxConfig()
		signer := &servertoken.HMACSigner{Key: []byte(tokenxConfig.SecretKey)}
		serverTokenProvider := servertoken.NewProvider(
			signer,
			tokenxConfig.Issuer,
			"permission-service",
			servertoken.WithTTL(1*time.Hour),
			servertoken.WithMargin(10*time.Second),
		)

		authGRPCClient, err = grpcclient.NewAuthGRPCClient(cfg.GRPCClient.AuthAddr, serverTokenProvider)
		if err != nil {
			return nil, fmt.Errorf("init auth grpc client: %w", err)
		}
	}

	// === Application Services ===
	permissionAppSvc := permissionApp.NewService(
		permissionRepo,
		permissionQuery,
	)

	userPermissionAppSvc := userPermissionApp.NewService(
		userPermissionRepo,
		userPermissionQuery,
	)

	authAppSvc := authApp.NewService(
		authGRPCClient,
		userPermissionAppSvc,
		tokenxInstance,
	)

	return &Dependencies{
		healthMgr:            healthMgr,
		postgres:             postgres,
		mongo:                mongoClient,
		cache:                cacheConn,
		kafkaConfig:          &kafkaConfig,
		busPublisher:         busPublisher,
		serverVerifier:       serverVerifier,
		permissionRepo:       permissionRepo,
		userPermissionRepo:   userPermissionRepo,
		authAppSvc:           authAppSvc,
		permissionAppSvc:     permissionAppSvc,
		userPermissionAppSvc: userPermissionAppSvc,
		tokenx:               tokenxInstance,
		authGRPCClient:       authGRPCClient,
	}, nil
}

func (d *Dependencies) Cleanup() {
	var cleanupList []any
	cleanupList = append(cleanupList, d.healthMgr, d.postgres, d.mongo, d.cache, d.busPublisher)
	if d.authGRPCClient != nil {
		cleanupList = append(cleanupList, d.authGRPCClient)
	}
	if err := cleanup.CleanupAll(cleanupList...); err != nil {
		logx.S().Errorf("cleanup permission service: %v", err)
	}
}

func (d *Dependencies) AuthAppSvc() *authApp.Service                 { return d.authAppSvc }
func (d *Dependencies) PermissionAppSvc() *permissionApp.Service   { return d.permissionAppSvc }
func (d *Dependencies) UserPermissionAppSvc() *userPermissionApp.Service { return d.userPermissionAppSvc }
func (d *Dependencies) Tokenx() *tokenx.Tokenx                      { return d.tokenx }
func (d *Dependencies) KafkaConfig() *kafkax.Config                { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher       { return d.busPublisher }
func (d *Dependencies) PermissionRepo() permissionDomain.Repo       { return d.permissionRepo }
func (d *Dependencies) UserPermissionRepo() userPermissionDomain.Repo { return d.userPermissionRepo }
func (d *Dependencies) ServerTokenVerifier() token.Verifier         { return d.serverVerifier }

