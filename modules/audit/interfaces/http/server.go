package http

import (
	"encoding/json"
	"time"

	actorSnapshotApp "nfxid/modules/audit/application/actor_snapshots"
	eventRetentionPolicyApp "nfxid/modules/audit/application/event_retention_policies"
	eventSearchIndexApp "nfxid/modules/audit/application/event_search_index"
	eventApp "nfxid/modules/audit/application/events"
	hashChainCheckpointApp "nfxid/modules/audit/application/hash_chain_checkpoints"
	"nfxid/modules/audit/interfaces/http/handler"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

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

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  fiberx.ErrorHandler,
		ReadTimeout:   30 * time.Second,
		WriteTimeout:  30 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	app.Use(middleware.Logger(), middleware.Recover())

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
