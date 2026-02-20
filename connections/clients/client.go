package clients

import (
	applicationpb "nfxid/protos/gen/clients/application"
	clientcredentialpb "nfxid/protos/gen/clients/client_credential"
	ipallowlistpb "nfxid/protos/gen/clients/ip_allowlist"
	ratelimitpb "nfxid/protos/gen/clients/rate_limit"
)

// Client Clients 服务 gRPC 客户端聚合
type Client struct {
	Application       *ApplicationClient
	ClientCredential  *ClientCredentialClient
	IpAllowlist       *IpAllowlistClient
	RateLimit         *RateLimitClient
}

// NewClient 创建 Clients 客户端
func NewClient(
	applicationClient applicationpb.ApplicationServiceClient,
	clientCredentialClient clientcredentialpb.ClientCredentialServiceClient,
	ipAllowlistClient ipallowlistpb.IpAllowlistServiceClient,
	rateLimitClient ratelimitpb.RateLimitServiceClient,
) *Client {
	return &Client{
		Application:      NewApplicationClient(applicationClient),
		ClientCredential: NewClientCredentialClient(clientCredentialClient),
		IpAllowlist:      NewIpAllowlistClient(ipAllowlistClient),
		RateLimit:        NewRateLimitClient(rateLimitClient),
	}
}
