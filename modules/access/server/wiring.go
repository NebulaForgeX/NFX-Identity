package server

import (
	"context"
	"fmt"
	"time"

	resourceApp "nfxid/modules/access/application/resource"
	superadminsApp "nfxid/modules/access/application/super_admins"
	tenantrolesApp "nfxid/modules/access/application/tenant_roles"
	"nfxid/modules/access/config"
	superadminsRepo "nfxid/modules/access/infrastructure/repository/super_admins"
	tenantrolesRepo "nfxid/modules/access/infrastructure/repository/tenant_roles"
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
	cache                *cachex.Connection
	postgres             *postgresqlx.Connection
	kafkaConfig          *kafkax.Config
	busPublisher         *eventbus.BusPublisher
	rabbitMQConfig       *rabbitmqx.Config
	tenantRoleAppSvc     *tenantrolesApp.Service
	superAdminAppSvc     *superadminsApp.Service
	resourceSvc          *resourceApp.Service
	userTokenVerifier    token.Verifier
	serverTokenVerifier  token.Verifier
	tokenxInstance       *tokenx.Tokenx
}

func NewDeps(ctx context.Context, cfg *config.Config) (*Dependencies, error) {
	postgres, err := postgresqlx.Init(ctx, cfg.PostgreSQL)
	if err != nil {
		return nil, fmt.Errorf("init PostgreSQL: %w", err)
	}

	cacheConn, err := cachex.InitConn(ctx, cfg.Cache)
	if err != nil {
		return nil, fmt.Errorf("init Redis: %w", err)
	}

	healthMgr := health.NewManager(ctx, 30*time.Second)
	healthMgr.Register(postgres)
	healthMgr.Register(cacheConn)

	kafkaConfig := cfg.KafkaConfig
	busPublisher, err := kafkax.NewPublisher(&kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka publisher: %w", err)
	}

	rabbitMQConfig := cfg.RabbitMQConfig
	tokenCfg := cfg.Token
	tokenxInstance := tokenx.New(tokenCfg)

	userTokenVerifier := &tokenxVerifierAdapter{tokenx: tokenxInstance}
	serverTokenVerifier := servertoken.NewVerifier(
		&servertoken.HMACSigner{Key: []byte(tokenCfg.SecretKey)},
		tokenCfg.Issuer,
		servertoken.WithAllowedSkew(5*time.Second),
	)

	tenantRoleRepoInstance := tenantrolesRepo.NewRepo(postgres.DB())
	tenantRoleAppSvc := tenantrolesApp.NewService(tenantRoleRepoInstance)
	superAdminRepoInstance := superadminsRepo.NewRepo(postgres.DB())
	superAdminAppSvc := superadminsApp.NewService(superAdminRepoInstance)
	resourceSvc := resourceApp.NewService(postgres, cacheConn, &kafkaConfig, &rabbitMQConfig)

	return &Dependencies{
		healthMgr:          healthMgr,
		postgres:           postgres,
		cache:              cacheConn,
		kafkaConfig:        &kafkaConfig,
		busPublisher:       busPublisher,
		rabbitMQConfig:     &rabbitMQConfig,
		tenantRoleAppSvc:   tenantRoleAppSvc,
		superAdminAppSvc:   superAdminAppSvc,
		resourceSvc:        resourceSvc,
		userTokenVerifier:  userTokenVerifier,
		serverTokenVerifier: serverTokenVerifier,
		tokenxInstance:     tokenxInstance,
	}, nil
}

func (d *Dependencies) HealthMgr() *health.Manager         { return d.healthMgr }
func (d *Dependencies) ResourceSvc() *resourceApp.Service  { return d.resourceSvc }
func (d *Dependencies) TenantRoleAppSvc() *tenantrolesApp.Service   { return d.tenantRoleAppSvc }
func (d *Dependencies) SuperAdminAppSvc() *superadminsApp.Service   { return d.superAdminAppSvc }
func (d *Dependencies) Postgres() *postgresqlx.Connection           { return d.postgres }
func (d *Dependencies) UserTokenVerifier() token.Verifier   { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config        { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher { return d.busPublisher }
func (d *Dependencies) RabbitMQConfig() *rabbitmqx.Config  { return d.rabbitMQConfig }

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

type tokenxVerifierAdapter struct {
	tokenx *tokenx.Tokenx
}

func (a *tokenxVerifierAdapter) Verify(ctx context.Context, tokenStr string) (*token.Claims, error) {
	claims, err := a.tokenx.VerifyAccessToken(tokenStr)
	if err != nil {
		return nil, err
	}
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
