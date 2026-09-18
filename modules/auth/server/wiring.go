package server

import (
	"context"
	"fmt"
	"time"

	"nfxidentity/modules/auth/application/platform"
	"nfxidentity/modules/auth/application/resource"
	"nfxidentity/modules/auth/config"
	emailQuery "nfxidentity/modules/auth/infrastructure/query/email"
	phoneQuery "nfxidentity/modules/auth/infrastructure/query/phone"
	profileQuery "nfxidentity/modules/auth/infrastructure/query/profile"
	repofactory "nfxidentity/modules/auth/infrastructure/repository/factory"
	"nfxidentity/pkgs/cachex"
	"nfxidentity/pkgs/email"
	"nfxidentity/pkgs/health"
	"nfxidentity/pkgs/kafkax"
	"nfxidentity/pkgs/kafkax/eventbus"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/transaction"
)

type Dependencies struct {
	healthMgr           *health.Manager
	cache               *cachex.Connection
	postgres            *postgresqlx.Connection
	kafkaConfig         *kafkax.Config
	busPublisher        *eventbus.BusPublisher
	userTokenVerifier   token.Verifier
	serverTokenVerifier token.Verifier
	tokenxInstance      *tokenx.Tokenx
	platformSvc         *platform.Service
	resourceSvc         *resource.Service
	mail                *email.EmailService
	repoFactory         *repofactory.TxRepoFactory
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

	tokenxInstance := tokenx.New(cfg.Token)
	userTokenVerifier := &tokenxVerifierAdapter{tokenx: tokenxInstance}
	serverTokenVerifier := servertoken.NewVerifier(
		&servertoken.HMACSigner{Key: []byte(cfg.Token.SecretKey)},
		cfg.Token.Issuer,
		servertoken.WithAllowedSkew(5*time.Second),
	)

	githubCfg := &platform.GitHubConfig{
		ClientID:     cfg.GitHub.ClientID,
		ClientSecret: cfg.GitHub.ClientSecret,
		RedirectURL:  cfg.GitHub.RedirectURL,
	}
	mail := email.NewEmailService(email.SMTPConfig{
		Host:     cfg.Email.SMTPHost,
		Port:     cfg.Email.SMTPPort,
		Username: cfg.Email.SMTPUser,
		Password: cfg.Email.SMTPPassword,
		From:     cfg.Email.SMTPFrom,
	})
	db := postgres.DB()
	factory := repofactory.NewTxRepoFactory(db)
	platformSvc := platform.NewService(
		transaction.NewGormTxManager(db),
		factory,
		emailQuery.NewQuery(db),
		phoneQuery.NewQuery(db),
		profileQuery.NewQuery(db),
		tokenxInstance,
		githubCfg,
		cacheConn.Client(),
		mail,
		busPublisher,
	)
	resourceSvc := resource.NewService(postgres, cacheConn, &kafkaConfig)

	return &Dependencies{
		healthMgr:           healthMgr,
		postgres:            postgres,
		cache:               cacheConn,
		kafkaConfig:         &kafkaConfig,
		busPublisher:        busPublisher,
		userTokenVerifier:   userTokenVerifier,
		serverTokenVerifier: serverTokenVerifier,
		tokenxInstance:      tokenxInstance,
		platformSvc:         platformSvc,
		resourceSvc:         resourceSvc,
		mail:                mail,
		repoFactory:         factory,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
}

func (d *Dependencies) HealthMgr() *health.Manager          { return d.healthMgr }
func (d *Dependencies) Postgres() *postgresqlx.Connection   { return d.postgres }
func (d *Dependencies) UserTokenVerifier() token.Verifier   { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier { return d.serverTokenVerifier }
func (d *Dependencies) KafkaConfig() *kafkax.Config         { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher {
	return d.busPublisher
}
func (d *Dependencies) PlatformSvc() *platform.Service { return d.platformSvc }
func (d *Dependencies) ResourceSvc() *resource.Service { return d.resourceSvc }
func (d *Dependencies) Cache() *cachex.Connection      { return d.cache }
func (d *Dependencies) Mail() *email.EmailService      { return d.mail }
func (d *Dependencies) RepoFactory() *repofactory.TxRepoFactory {
	return d.repoFactory
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
			"account_id":    claims.AccountID,
			"profile_id":    claims.ProfileID,
			"username":      claims.Username,
			"email":         claims.Email,
			"phone":         claims.Phone,
			"profile_scope": claims.ProfileScope,
			"type":          claims.Type,
		},
	}, nil
}
