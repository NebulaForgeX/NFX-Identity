package clients

import (
	apikeypb "nfxid/protos/gen/clients/api_key"
	apppb "nfxid/protos/gen/clients/app"
	clientcredentialpb "nfxid/protos/gen/clients/client_credential"
	clientscopepb "nfxid/protos/gen/clients/client_scope"
	ipallowlistpb "nfxid/protos/gen/clients/ip_allowlist"
	ratelimitpb "nfxid/protos/gen/clients/rate_limit"
)

// Client Clients 服务客户端
type Client struct {
	App              *AppClient
	ApiKey           *ApiKeyClient
	ClientCredential *ClientCredentialClient
	ClientScope      *ClientScopeClient
	IpAllowlist      *IpAllowlistClient
	RateLimit        *RateLimitClient
}

// NewClient 创建 Clients 客户端
func NewClient(
	appClient apppb.AppServiceClient,
	apiKeyClient apikeypb.ApiKeyServiceClient,
	clientCredentialClient clientcredentialpb.ClientCredentialServiceClient,
	clientScopeClient clientscopepb.ClientScopeServiceClient,
	ipAllowlistClient ipallowlistpb.IpAllowlistServiceClient,
	rateLimitClient ratelimitpb.RateLimitServiceClient,
) *Client {
	return &Client{
		App:              NewAppClient(appClient),
		ApiKey:           NewApiKeyClient(apiKeyClient),
		ClientCredential: NewClientCredentialClient(clientCredentialClient),
		ClientScope:      NewClientScopeClient(clientScopeClient),
		IpAllowlist:      NewIpAllowlistClient(ipAllowlistClient),
		RateLimit:        NewRateLimitClient(rateLimitClient),
	}
}
