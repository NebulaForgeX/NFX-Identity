package server

import (
	"context"
	"fmt"
	"time"

	imageApp "nebulaid/modules/image/application/image"
	imageTypeApp "nebulaid/modules/image/application/image_type"
	"nebulaid/modules/image/config"
	imageDomain "nebulaid/modules/image/domain/image"
	imageTypeDomain "nebulaid/modules/image/domain/image_type"
	"nebulaid/modules/image/infrastructure/grpcclient"
	"nebulaid/modules/image/infrastructure/query"
	"nebulaid/modules/image/infrastructure/repository"
	"nebulaid/pkgs/cleanup"
	"nebulaid/pkgs/eventbus"
	"nebulaid/pkgs/health"
	"nebulaid/pkgs/kafkax"
	"nebulaid/pkgs/logx"
	"nebulaid/pkgs/mongodbx"
	"nebulaid/pkgs/postgresqlx"
	"nebulaid/pkgs/security/token"
	"nebulaid/pkgs/security/token/servertoken"
)

type Dependencies struct {
	healthMgr       *health.Manager
	postgres        *postgresqlx.Connection
	mongo           *mongodbx.Client
	kafkaConfig     *kafkax.Config
	busPublisher    *eventbus.BusPublisher
	serverVerifier  token.Verifier
	authGRPCClient  *grpcclient.AuthGRPCClient
	imageRepo       imageDomain.Repo
	imageTypeRepo   imageTypeDomain.Repo
	imageAppSvc     *imageApp.Service
	imageTypeAppSvc *imageTypeApp.Service
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

	// === Health Manager ===
	healthMgr := health.NewManager(ctx, 30*time.Second)
	healthMgr.Register(postgres)
	healthMgr.Register(mongoClient)

	// === Token Verifier ===
	serverVerifier := servertoken.NewVerifier(
		&servertoken.HMACSigner{Key: []byte(cfg.Token.SecretKey)},
		cfg.Token.Issuer,
		servertoken.WithAllowedSkew(5*time.Second),
	)

	// === gRPC Client ===
	var authGRPCClient *grpcclient.AuthGRPCClient
	if cfg.GRPCClient.AuthAddr != "" {
		// 创建 server token provider（用于服务间 gRPC 认证）
		signer := &servertoken.HMACSigner{Key: []byte(cfg.Token.SecretKey)}
		serverTokenProvider := servertoken.NewProvider(
			signer,
			cfg.Token.Issuer,
			"image-service",
			servertoken.WithTTL(1*time.Hour),
			servertoken.WithMargin(10*time.Second),
		)

		authGRPCClient, err = grpcclient.NewAuthGRPCClient(cfg.GRPCClient.AuthAddr, serverTokenProvider)
		if err != nil {
			return nil, fmt.Errorf("init auth grpc client: %w", err)
		}
	}

	// === Kafka Service ===
	kafkaConfig := cfg.KafkaConfig
	busPublisher, err := kafkax.NewPublisher(&kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka publisher: %w", err)
	}

	// === Repository ===
	imageRepo := repository.NewImagePGRepo(postgres.DB())
	imageTypeRepo := repository.NewImageTypePGRepo(postgres.DB())

	// === Query ===
	imageQuery := query.NewImagePGQuery(postgres.DB())
	imageTypeQuery := query.NewImageTypePGQuery(postgres.DB())

	// === Application Services ===
	imageAppSvc := imageApp.NewService(
		imageRepo,
		imageQuery,
		busPublisher,
	)

	imageTypeAppSvc := imageTypeApp.NewService(
		imageTypeRepo,
		imageTypeQuery,
	)

	return &Dependencies{
		healthMgr:       healthMgr,
		postgres:        postgres,
		mongo:           mongoClient,
		kafkaConfig:     &kafkaConfig,
		busPublisher:    busPublisher,
		serverVerifier:  serverVerifier,
		authGRPCClient:  authGRPCClient,
		imageRepo:       imageRepo,
		imageTypeRepo:   imageTypeRepo,
		imageAppSvc:     imageAppSvc,
		imageTypeAppSvc: imageTypeAppSvc,
	}, nil
}

func (d *Dependencies) Cleanup() {
	var cleanupList []any
	cleanupList = append(cleanupList, d.healthMgr, d.postgres, d.mongo, d.busPublisher)
	if d.authGRPCClient != nil {
		cleanupList = append(cleanupList, d.authGRPCClient)
	}
	if err := cleanup.CleanupAll(cleanupList...); err != nil {
		logx.S().Errorf("cleanup image service: %v", err)
	}
}

func (d *Dependencies) ImageAppSvc() *imageApp.Service         { return d.imageAppSvc }
func (d *Dependencies) ImageTypeAppSvc() *imageTypeApp.Service { return d.imageTypeAppSvc }
func (d *Dependencies) KafkaConfig() *kafkax.Config            { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher   { return d.busPublisher }
func (d *Dependencies) ServerTokenVerifier() token.Verifier    { return d.serverVerifier }
