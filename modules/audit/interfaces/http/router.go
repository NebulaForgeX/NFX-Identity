package http

import (
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/usertoken"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app           fiber.Router
	tokenVerifier token.Verifier
	handlers      *Registry
}

func NewRouter(app fiber.Router, tokenVerifier token.Verifier, handlers *Registry) *Router {
	return &Router{
		app:           app,
		tokenVerifier: tokenVerifier,
		handlers:      handlers,
	}
}

func (r *Router) RegisterRoutes() {
	audit := r.app.Group("/audit")

	// 需要认证的路由（需要token）
	auth := audit.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 事件相关
		auth.Post("/events", r.handlers.Event.Create)
		auth.Get("/events/:id", r.handlers.Event.GetByID)
		auth.Delete("/events/:id", r.handlers.Event.Delete)

		// Actor Snapshot 相关
		auth.Post("/actor-snapshots", r.handlers.ActorSnapshot.Create)
		auth.Get("/actor-snapshots/:id", r.handlers.ActorSnapshot.GetByID)
		auth.Delete("/actor-snapshots/:id", r.handlers.ActorSnapshot.Delete)

		// Event Retention Policy 相关
		auth.Post("/event-retention-policies", r.handlers.EventRetentionPolicy.Create)
		auth.Get("/event-retention-policies/:id", r.handlers.EventRetentionPolicy.GetByID)
		auth.Put("/event-retention-policies/:id", r.handlers.EventRetentionPolicy.Update)
		auth.Delete("/event-retention-policies/:id", r.handlers.EventRetentionPolicy.Delete)

		// Event Search Index 相关
		auth.Post("/event-search-index", r.handlers.EventSearchIndex.Create)
		auth.Get("/event-search-index/:id", r.handlers.EventSearchIndex.GetByID)
		auth.Delete("/event-search-index/:id", r.handlers.EventSearchIndex.Delete)

		// Hash Chain Checkpoint 相关
		auth.Post("/hash-chain-checkpoints", r.handlers.HashChainCheckpoint.Create)
		auth.Get("/hash-chain-checkpoints/:id", r.handlers.HashChainCheckpoint.GetByID)
		auth.Delete("/hash-chain-checkpoints/:id", r.handlers.HashChainCheckpoint.Delete)
	}
}
