package server

import (
	"context"
	"fmt"
	"time"

	domainVerificationApp "nfxid/modules/tenants/application/domain_verifications"
	groupApp "nfxid/modules/tenants/application/groups"
	invitationApp "nfxid/modules/tenants/application/invitations"
	memberAppRoleApp "nfxid/modules/tenants/application/member_app_roles"
	memberGroupApp "nfxid/modules/tenants/application/member_groups"
	memberRoleApp "nfxid/modules/tenants/application/member_roles"
	memberApp "nfxid/modules/tenants/application/members"
	tenantAppApp "nfxid/modules/tenants/application/tenant_apps"
	tenantSettingApp "nfxid/modules/tenants/application/tenant_settings"
	tenantApp "nfxid/modules/tenants/application/tenants"
	"nfxid/modules/tenants/config"
	domainVerificationRepo "nfxid/modules/tenants/infrastructure/repository/domain_verifications"
	groupRepo "nfxid/modules/tenants/infrastructure/repository/groups"
	invitationRepo "nfxid/modules/tenants/infrastructure/repository/invitations"
	memberAppRoleRepo "nfxid/modules/tenants/infrastructure/repository/member_app_roles"
	memberGroupRepo "nfxid/modules/tenants/infrastructure/repository/member_groups"
	memberRoleRepo "nfxid/modules/tenants/infrastructure/repository/member_roles"
	memberRepo "nfxid/modules/tenants/infrastructure/repository/members"
	tenantAppRepo "nfxid/modules/tenants/infrastructure/repository/tenant_apps"
	tenantSettingRepo "nfxid/modules/tenants/infrastructure/repository/tenant_settings"
	tenantRepo "nfxid/modules/tenants/infrastructure/repository/tenants"
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
	healthMgr                  *health.Manager
	cache                      *cache.Connection
	postgres                   *postgresqlx.Connection
	kafkaConfig                *kafkax.Config
	busPublisher               *eventbus.BusPublisher
	rabbitMQConfig             *rabbitmqx.Config
	tenantAppSvc               *tenantApp.Service
	groupAppSvc                *groupApp.Service
	memberAppSvc               *memberApp.Service
	invitationAppSvc           *invitationApp.Service
	tenantAppAppSvc            *tenantAppApp.Service
	tenantSettingAppSvc        *tenantSettingApp.Service
	domainVerificationAppSvc   *domainVerificationApp.Service
	memberRoleAppSvc           *memberRoleApp.Service
	memberGroupAppSvc          *memberGroupApp.Service
	memberAppRoleAppSvc        *memberAppRoleApp.Service
	userTokenVerifier          token.Verifier
	serverTokenVerifier        token.Verifier
	tokenxInstance             *tokenx.Tokenx
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
	// 使用默认 token 配置（tenants 模块可能不需要 token 生成，只需要验证）
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
	tenantRepoInstance := tenantRepo.NewRepo(postgres.DB())
	groupRepoInstance := groupRepo.NewRepo(postgres.DB())
	memberRepoInstance := memberRepo.NewRepo(postgres.DB())
	invitationRepoInstance := invitationRepo.NewRepo(postgres.DB())
	tenantAppRepoInstance := tenantAppRepo.NewRepo(postgres.DB())
	tenantSettingRepoInstance := tenantSettingRepo.NewRepo(postgres.DB())
	domainVerificationRepoInstance := domainVerificationRepo.NewRepo(postgres.DB())
	memberRoleRepoInstance := memberRoleRepo.NewRepo(postgres.DB())
	memberGroupRepoInstance := memberGroupRepo.NewRepo(postgres.DB())
	memberAppRoleRepoInstance := memberAppRoleRepo.NewRepo(postgres.DB())

	//! === Application Services ===
	tenantAppSvc := tenantApp.NewService(tenantRepoInstance)
	groupAppSvc := groupApp.NewService(groupRepoInstance)
	memberAppSvc := memberApp.NewService(memberRepoInstance)
	invitationAppSvc := invitationApp.NewService(invitationRepoInstance)
	tenantAppAppSvc := tenantAppApp.NewService(tenantAppRepoInstance)
	tenantSettingAppSvc := tenantSettingApp.NewService(tenantSettingRepoInstance)
	domainVerificationAppSvc := domainVerificationApp.NewService(domainVerificationRepoInstance)
	memberRoleAppSvc := memberRoleApp.NewService(memberRoleRepoInstance)
	memberGroupAppSvc := memberGroupApp.NewService(memberGroupRepoInstance)
	memberAppRoleAppSvc := memberAppRoleApp.NewService(memberAppRoleRepoInstance)

	return &Dependencies{
		healthMgr:                healthMgr,
		postgres:                 postgres,
		cache:                    cacheConn,
		kafkaConfig:              &kafkaConfig,
		busPublisher:             busPublisher,
		rabbitMQConfig:           &rabbitMQConfig,
		tenantAppSvc:             tenantAppSvc,
		groupAppSvc:              groupAppSvc,
		memberAppSvc:             memberAppSvc,
		invitationAppSvc:         invitationAppSvc,
		tenantAppAppSvc:          tenantAppAppSvc,
		tenantSettingAppSvc:     tenantSettingAppSvc,
		domainVerificationAppSvc: domainVerificationAppSvc,
		memberRoleAppSvc:         memberRoleAppSvc,
		memberGroupAppSvc:        memberGroupAppSvc,
		memberAppRoleAppSvc:      memberAppRoleAppSvc,
		userTokenVerifier:        userTokenVerifier,
		serverTokenVerifier:      serverTokenVerifier,
		tokenxInstance:           tokenxInstance,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

// Getter methods for interfaces
func (d *Dependencies) TenantAppSvc() *tenantApp.Service                    { return d.tenantAppSvc }
func (d *Dependencies) GroupAppSvc() *groupApp.Service                      { return d.groupAppSvc }
func (d *Dependencies) MemberAppSvc() *memberApp.Service                   { return d.memberAppSvc }
func (d *Dependencies) InvitationAppSvc() *invitationApp.Service             { return d.invitationAppSvc }
func (d *Dependencies) TenantAppAppSvc() *tenantAppApp.Service               { return d.tenantAppAppSvc }
func (d *Dependencies) TenantSettingAppSvc() *tenantSettingApp.Service       { return d.tenantSettingAppSvc }
func (d *Dependencies) DomainVerificationAppSvc() *domainVerificationApp.Service { return d.domainVerificationAppSvc }
func (d *Dependencies) MemberRoleAppSvc() *memberRoleApp.Service             { return d.memberRoleAppSvc }
func (d *Dependencies) MemberGroupAppSvc() *memberGroupApp.Service            { return d.memberGroupAppSvc }
func (d *Dependencies) MemberAppRoleAppSvc() *memberAppRoleApp.Service       { return d.memberAppRoleAppSvc }
func (d *Dependencies) UserTokenVerifier() token.Verifier                   { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier                   { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config                           { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher                 { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config                     { return d.rabbitMQConfig }

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
