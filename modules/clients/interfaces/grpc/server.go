package grpc

import (
	apiKeyApp "nfxid/modules/clients/application/api_keys"
	appApp "nfxid/modules/clients/application/apps"
	clientCredentialApp "nfxid/modules/clients/application/client_credentials"
	clientScopeApp "nfxid/modules/clients/application/client_scopes"
	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	resourceApp "nfxid/modules/clients/application/resource"
	grpcHandler "nfxid/modules/clients/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"
	apppb "nfxid/protos/gen/clients/app"
	apikeypb "nfxid/protos/gen/clients/api_key"
	clientcredentialpb "nfxid/protos/gen/clients/client_credential"
	clientscopepb "nfxid/protos/gen/clients/client_scope"
	ipallowlistpb "nfxid/protos/gen/clients/ip_allowlist"
	ratelimitpb "nfxid/protos/gen/clients/rate_limit"

	"google.golang.org/grpc"
)

type Deps interface {
	AppAppSvc() *appApp.Service
	APIKeyAppSvc() *apiKeyApp.Service
	ClientCredentialAppSvc() *clientCredentialApp.Service
	ClientScopeAppSvc() *clientScopeApp.Service
	IPAllowlistAppSvc() *ipAllowlistApp.Service
	RateLimitAppSvc() *rateLimitApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	apppb.RegisterAppServiceServer(s, grpcHandler.NewAppHandler(d.AppAppSvc()))
	apikeypb.RegisterApiKeyServiceServer(s, grpcHandler.NewAPIKeyHandler(d.APIKeyAppSvc()))
	clientcredentialpb.RegisterClientCredentialServiceServer(s, grpcHandler.NewClientCredentialHandler(d.ClientCredentialAppSvc()))
	clientscopepb.RegisterClientScopeServiceServer(s, grpcHandler.NewClientScopeHandler(d.ClientScopeAppSvc()))
	ipallowlistpb.RegisterIpAllowlistServiceServer(s, grpcHandler.NewIPAllowlistHandler(d.IPAllowlistAppSvc()))
	ratelimitpb.RegisterRateLimitServiceServer(s, grpcHandler.NewRateLimitHandler(d.RateLimitAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "clients"))
	
	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "clients"))

	return s
}
