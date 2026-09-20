package http

import (
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
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
	system := r.app.Group("/system")
	r.RegisterLocalesGroup(system)
	r.RegisterSystemStateGroup(system)
}

func (r *Router) RegisterLocalesGroup(system fiber.Router) {
	locales := system.Group("/locales")
	locales.Get("/:lang", r.handlers.I18n.GetErrorTranslations)
	messages := system.Group("/messages")
	messages.Get("/:lang", r.handlers.I18n.GetMessageTranslations)
}

func (r *Router) RegisterSystemStateGroup(system fiber.Router) {
	state := system.Group("/system-state")
	state.Get("/latest", r.handlers.SystemState.GetLatest)
	state.Post("/initialize", r.handlers.SystemState.Initialize)
}
