package server

import (
	"context"
	"fmt"
	"time"

	actorSnapshotApp "nfxid/modules/audit/application/actor_snapshots"
	resourceApp "nfxid/modules/audit/application/resource"
	eventApp "nfxid/modules/audit/application/events"
	eventRetentionPolicyApp "nfxid/modules/audit/application/event_retention_policies"
	eventSearchIndexApp "nfxid/modules/audit/application/event_search_index"
	hashChainCheckpointApp "nfxid/modules/audit/application/hash_chain_checkpoints"
	"nfxid/modules/audit/config"
	actorSnapshotRepo "nfxid/modules/audit/infrastructure/repository/actor_snapshots"
	eventRepo "nfxid/modules/audit/infrastructure/repository/events"
	eventRetentionPolicyRepo "nfxid/modules/audit/infrastructure/repository/event_retention_policies"
	eventSearchIndexRepo "nfxid/modules/audit/infrastructure/repository/event_search_index"
	hashChainCheckpointRepo "nfxid/modules/audit/infrastructure/repository/hash_chain_checkpoints"
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
	healthMgr                    *health.Manager
	cache                        *cache.Connection
	postgres                     *postgresqlx.Connection
	kafkaConfig                  *kafkax.Config
	busPublisher                 *eventbus.BusPublisher
	rabbitMQConfig               *rabbitmqx.Config
	eventAppSvc                  *eventApp.Service
	actorSnapshotAppSvc          *actorSnapshotApp.Service
	eventRetentionPolicyAppSvc   *eventRetentionPolicyApp.Service
	eventSearchIndexAppSvc       *eventSearchIndexApp.Service
	hashChainCheckpointAppSvc    *hashChainCheckpointApp.Service
	userTokenVerifier            token.Verifier
	serverTokenVerifier          token.Verifier
	resourceSvc         *resourceApp.Service
	tokenxInstance               *tokenx.Tokenx
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
	eventRepoInstance := eventRepo.NewRepo(postgres.DB())
	actorSnapshotRepoInstance := actorSnapshotRepo.NewRepo(postgres.DB())
	eventRetentionPolicyRepoInstance := eventRetentionPolicyRepo.NewRepo(postgres.DB())
	eventSearchIndexRepoInstance := eventSearchIndexRepo.NewRepo(postgres.DB())
	hashChainCheckpointRepoInstance := hashChainCheckpointRepo.NewRepo(postgres.DB())

	//! === Application Services ===
	eventAppSvc := eventApp.NewService(eventRepoInstance)
	actorSnapshotAppSvc := actorSnapshotApp.NewService(actorSnapshotRepoInstance)
	eventRetentionPolicyAppSvc := eventRetentionPolicyApp.NewService(eventRetentionPolicyRepoInstance)
	eventSearchIndexAppSvc := eventSearchIndexApp.NewService(eventSearchIndexRepoInstance)
	hashChainCheckpointAppSvc := hashChainCheckpointApp.NewService(hashChainCheckpointRepoInstance)

	resourceSvc := resourceApp.NewService(postgres, cacheConn, &kafkaConfig, &rabbitMQConfig)

	return &Dependencies{
		healthMgr:                  healthMgr,
		postgres:                   postgres,
		cache:                      cacheConn,
		kafkaConfig:                &kafkaConfig,
		busPublisher:               busPublisher,
		rabbitMQConfig:             &rabbitMQConfig,
		eventAppSvc:                eventAppSvc,
		actorSnapshotAppSvc:        actorSnapshotAppSvc,
		eventRetentionPolicyAppSvc: eventRetentionPolicyAppSvc,
		eventSearchIndexAppSvc:     eventSearchIndexAppSvc,
		hashChainCheckpointAppSvc:  hashChainCheckpointAppSvc,
		userTokenVerifier:          userTokenVerifier,
		serverTokenVerifier:        serverTokenVerifier,
		resourceSvc:         resourceSvc,
		tokenxInstance:             tokenxInstance,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

// Getter methods for interfaces
func (d *Dependencies) HealthMgr() *health.Manager                           { return d.healthMgr }
func (d *Dependencies) ResourceSvc() *resourceApp.Service { return d.resourceSvc }
func (d *Dependencies) EventAppSvc() *eventApp.Service                        { return d.eventAppSvc }
func (d *Dependencies) ActorSnapshotAppSvc() *actorSnapshotApp.Service       { return d.actorSnapshotAppSvc }
func (d *Dependencies) EventRetentionPolicyAppSvc() *eventRetentionPolicyApp.Service { return d.eventRetentionPolicyAppSvc }
func (d *Dependencies) EventSearchIndexAppSvc() *eventSearchIndexApp.Service  { return d.eventSearchIndexAppSvc }
func (d *Dependencies) HashChainCheckpointAppSvc() *hashChainCheckpointApp.Service { return d.hashChainCheckpointAppSvc }
func (d *Dependencies) Postgres() *postgresqlx.Connection                    { return d.postgres }
func (d *Dependencies) UserTokenVerifier() token.Verifier                    { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier                  { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config                          { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher                 { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config                    { return d.rabbitMQConfig }

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
