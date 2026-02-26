package http

import (
	"encoding/json"
	"time"

	actorSnapshotApp "nfxidentity/modules/audit/application/actor_snapshots"
	eventRetentionPolicyApp "nfxidentity/modules/audit/application/event_retention_policies"
	eventSearchIndexApp "nfxidentity/modules/audit/application/event_search_index"
	eventApp "nfxidentity/modules/audit/application/events"
	hashChainCheckpointApp "nfxidentity/modules/audit/application/hash_chain_checkpoints"
	"nfxidentity/modules/audit/interfaces/http/handler"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type httpDeps interface {
	EventAppSvc() *eventApp.Service
	ActorSnapshotAppSvc() *actorSnapshotApp.Service
	EventRetentionPolicyAppSvc() *eventRetentionPolicyApp.Service
	EventSearchIndexAppSvc() *eventSearchIndexApp.Service
	HashChainCheckpointAppSvc() *hashChainCheckpointApp.Service
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps, accessLog httpx.AccessLogConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: fiberx.ErrorHandler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	app.Use(middleware.Logger(), middleware.AccessLog(accessLog), middleware.Recover())

	reg := &Registry{
		Event:                handler.NewEventHandler(d.EventAppSvc()),
		ActorSnapshot:        handler.NewActorSnapshotHandler(d.ActorSnapshotAppSvc()),
		EventRetentionPolicy: handler.NewEventRetentionPolicyHandler(d.EventRetentionPolicyAppSvc()),
		EventSearchIndex:     handler.NewEventSearchIndexHandler(d.EventSearchIndexAppSvc()),
		HashChainCheckpoint:  handler.NewHashChainCheckpointHandler(d.HashChainCheckpointAppSvc()),
	}

	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
