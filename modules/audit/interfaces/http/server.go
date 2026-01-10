package http

import (
	"encoding/json"

	actorSnapshotApp "nfxid/modules/audit/application/actor_snapshots"
	eventApp "nfxid/modules/audit/application/events"
	eventRetentionPolicyApp "nfxid/modules/audit/application/event_retention_policies"
	eventSearchIndexApp "nfxid/modules/audit/application/event_search_index"
	hashChainCheckpointApp "nfxid/modules/audit/application/hash_chain_checkpoints"
	"nfxid/modules/audit/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
	}))

	app.Use(recover.RecoverMiddleware(), logger.New())

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
