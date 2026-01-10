package http

import (
	"nfxid/modules/clients/interfaces/http/handler"
)

type Registry struct {
	App             *handler.AppHandler
	APIKey          *handler.APIKeyHandler
	ClientCredential *handler.ClientCredentialHandler
	ClientScope     *handler.ClientScopeHandler
	IPAllowlist     *handler.IPAllowlistHandler
	RateLimit       *handler.RateLimitHandler
}
