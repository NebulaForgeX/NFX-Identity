package server

import (
	"context"
	"fmt"
	"time"

	grantApp "nfxid/modules/access/application/grants"
	permissionApp "nfxid/modules/access/application/permissions"
	resourceApp "nfxid/modules/access/application/resource"
	roleApp "nfxid/modules/access/application/roles"
	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	scopeApp "nfxid/modules/access/application/scopes"
	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	"nfxid/modules/access/config"
	grantRepo "nfxid/modules/access/infrastructure/repository/grants"
	permissionRepo "nfxid/modules/access/infrastructure/repository/permissions"
	roleRepo "nfxid/modules/access/infrastructure/repository/roles"
	rolePermissionRepo "nfxid/modules/access/infrastructure/repository/role_permissions"
	scopeRepo "nfxid/modules/access/infrastructure/repository/scopes"
	scopePermissionRepo "nfxid/modules/access/infrastructure/repository/scope_permissions"
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
	roleAppSvc               *roleApp.Service
	permissionAppSvc         *permissionApp.Service
	scopeAppSvc              *scopeApp.Service
	grantAppSvc              *grantApp.Service
	rolePermissionAppSvc     *rolePermissionApp.Service
	scopePermissionAppSvc    *scopePermissionApp.Service
	resourceSvc              *resourceApp.Service
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
	// 使用默认 token 配置（access 模块可能不需要 token 生成，只需要验证）
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
	roleRepoInstance := roleRepo.NewRepo(postgres.DB())
	permissionRepoInstance := permissionRepo.NewRepo(postgres.DB())
	scopeRepoInstance := scopeRepo.NewRepo(postgres.DB())
	grantRepoInstance := grantRepo.NewRepo(postgres.DB())
	rolePermissionRepoInstance := rolePermissionRepo.NewRepo(postgres.DB())
	scopePermissionRepoInstance := scopePermissionRepo.NewRepo(postgres.DB())

	//! === Application Services ===
	roleAppSvc := roleApp.NewService(roleRepoInstance)
	permissionAppSvc := permissionApp.NewService(permissionRepoInstance)
	scopeAppSvc := scopeApp.NewService(scopeRepoInstance)
	grantAppSvc := grantApp.NewService(grantRepoInstance)
	rolePermissionAppSvc := rolePermissionApp.NewService(rolePermissionRepoInstance)
	scopePermissionAppSvc := scopePermissionApp.NewService(scopePermissionRepoInstance)
	resourceSvc := resourceApp.NewService(postgres, cacheConn, &kafkaConfig, &rabbitMQConfig)

	return &Dependencies{
		healthMgr:             healthMgr,
		postgres:              postgres,
		cache:                 cacheConn,
		kafkaConfig:           &kafkaConfig,
		busPublisher:          busPublisher,
		rabbitMQConfig:        &rabbitMQConfig,
		roleAppSvc:            roleAppSvc,
		permissionAppSvc:      permissionAppSvc,
		scopeAppSvc:           scopeAppSvc,
		grantAppSvc:           grantAppSvc,
		rolePermissionAppSvc: rolePermissionAppSvc,
		scopePermissionAppSvc: scopePermissionAppSvc,
		resourceSvc:           resourceSvc,
		userTokenVerifier:     userTokenVerifier,
		serverTokenVerifier:   serverTokenVerifier,
		tokenxInstance:        tokenxInstance,
	}, nil
}

// Getter methods for interfaces
func (d *Dependencies) HealthMgr() *health.Manager { return d.healthMgr }
func (d *Dependencies) ResourceSvc() *resourceApp.Service { return d.resourceSvc }

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

// Getter methods for interfaces
func (d *Dependencies) RoleAppSvc() *roleApp.Service                    { return d.roleAppSvc }
func (d *Dependencies) PermissionAppSvc() *permissionApp.Service        { return d.permissionAppSvc }
func (d *Dependencies) ScopeAppSvc() *scopeApp.Service                  { return d.scopeAppSvc }
func (d *Dependencies) GrantAppSvc() *grantApp.Service                 { return d.grantAppSvc }
func (d *Dependencies) RolePermissionAppSvc() *rolePermissionApp.Service { return d.rolePermissionAppSvc }
func (d *Dependencies) ScopePermissionAppSvc() *scopePermissionApp.Service { return d.scopePermissionAppSvc }
func (d *Dependencies) UserTokenVerifier() token.Verifier              { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier            { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config                     { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher         { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config               { return d.rabbitMQConfig }

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
