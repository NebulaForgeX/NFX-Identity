package server

import (
	"context"
	"fmt"
	"time"

	authconn "nfxidentity/connections/auth"
	"nfxidentity/modules/asset/application/media"
	"nfxidentity/modules/asset/application/resource"
	"nfxidentity/modules/asset/config"
	"nfxidentity/modules/asset/infrastructure/objectstore"
	audiosQuery "nfxidentity/modules/asset/infrastructure/query/audios"
	filesQuery "nfxidentity/modules/asset/infrastructure/query/files"
	imagesQuery "nfxidentity/modules/asset/infrastructure/query/images"
	videosQuery "nfxidentity/modules/asset/infrastructure/query/videos"
	audiosRepo "nfxidentity/modules/asset/infrastructure/repository/audios"
	filesRepo "nfxidentity/modules/asset/infrastructure/repository/files"
	imagesRepo "nfxidentity/modules/asset/infrastructure/repository/images"
	videosRepo "nfxidentity/modules/asset/infrastructure/repository/videos"
	"nfxidentity/pkgs/cachex"
	"nfxidentity/pkgs/health"
	"nfxidentity/pkgs/kafkax"
	"nfxidentity/pkgs/kafkax/eventbus"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/transaction"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Dependencies struct {
	healthMgr           *health.Manager
	postgres            *postgresqlx.Connection
	cache               *cachex.Connection
	kafkaConfig         *kafkax.Config
	busPublisher        *eventbus.BusPublisher
	userTokenVerifier   token.Verifier
	serverTokenVerifier token.Verifier
	mediaSvc            *media.Service
	resourceSvc         *resource.Service
	authClient          *authconn.Client
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

	mc, err := minio.New(cfg.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIO.AccessKey, cfg.MinIO.SecretKey, ""),
		Secure: cfg.MinIO.UseSSL,
		Region: cfg.MinIO.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("init MinIO: %w", err)
	}
	db := postgres.DB()
	mediaSvc := media.NewService(
		transaction.NewGormTxManager(db),
		objectstore.New(mc, cfg.MinIO.Bucket),
		imagesRepo.NewRepo(db),
		filesRepo.NewRepo(db),
		videosRepo.NewRepo(db),
		audiosRepo.NewRepo(db),
		imagesQuery.NewQuery(db),
		filesQuery.NewQuery(db),
		videosQuery.NewQuery(db),
		audiosQuery.NewQuery(db),
	)
	if err := mediaSvc.EnsureBucket(ctx); err != nil {
		return nil, fmt.Errorf("ensure MinIO bucket: %w", err)
	}

	var authClient *authconn.Client
	if cfg.GRPCClient.AuthAddr != "" {
		authClient, err = authconn.Dial(authconn.GRPCConfig{
			Addr:           cfg.GRPCClient.AuthAddr,
			TokenSecretKey: cfg.Token.SecretKey,
			TokenIssuer:    cfg.Token.Issuer,
			CallerService:  "asset-service",
		})
		if err != nil {
			return nil, fmt.Errorf("dial auth: %w", err)
		}
	}

	tokenxInstance := tokenx.New(cfg.Token)
	resourceSvc := resource.NewService(postgres, cacheConn, &kafkaConfig)
	return &Dependencies{
		healthMgr:         healthMgr,
		postgres:          postgres,
		cache:             cacheConn,
		kafkaConfig:       &kafkaConfig,
		busPublisher:      busPublisher,
		userTokenVerifier: &tokenxVerifierAdapter{tokenx: tokenxInstance},
		serverTokenVerifier: servertoken.NewVerifier(
			&servertoken.HMACSigner{Key: []byte(cfg.Token.SecretKey)},
			cfg.Token.Issuer,
			servertoken.WithAllowedSkew(5*time.Second),
		),
		mediaSvc:    mediaSvc,
		resourceSvc: resourceSvc,
		authClient:  authClient,
	}, nil
}

func (d *Dependencies) Cleanup() {
	d.healthMgr.Stop()
	d.postgres.Close()
	d.cache.Close()
	if d.authClient != nil {
		_ = d.authClient.Close()
	}
}

func (d *Dependencies) MediaSvc() *media.Service             { return d.mediaSvc }
func (d *Dependencies) AuthClient() *authconn.Client         { return d.authClient }
func (d *Dependencies) UserTokenVerifier() token.Verifier    { return d.userTokenVerifier }
func (d *Dependencies) ServerTokenVerifier() token.Verifier  { return d.serverTokenVerifier }
func (d *Dependencies) HealthMgr() *health.Manager           { return d.healthMgr }
func (d *Dependencies) KafkaConfig() *kafkax.Config          { return d.kafkaConfig }
func (d *Dependencies) BusPublisher() *eventbus.BusPublisher { return d.busPublisher }
func (d *Dependencies) Postgres() *postgresqlx.Connection    { return d.postgres }
func (d *Dependencies) ResourceSvc() *resource.Service       { return d.resourceSvc }

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
