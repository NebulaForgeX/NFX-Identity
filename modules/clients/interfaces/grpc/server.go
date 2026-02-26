package grpc

import (
	apiKeyApp "nfxidentity/modules/clients/application/api_keys"
	appApp "nfxidentity/modules/clients/application/apps"
	clientCredentialApp "nfxidentity/modules/clients/application/client_credentials"
	clientScopeApp "nfxidentity/modules/clients/application/client_scopes"
	ipAllowlistApp "nfxidentity/modules/clients/application/ip_allowlist"
	rateLimitApp "nfxidentity/modules/clients/application/rate_limits"
	resourceApp "nfxidentity/modules/clients/application/resource"
	grpcHandler "nfxidentity/modules/clients/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	applicationpb "nfxidentity/protos/gen/clients/application"
	clientcredentialpb "nfxidentity/protos/gen/clients/client_credential"
	ipallowlistpb "nfxidentity/protos/gen/clients/ip_allowlist"
	ratelimitpb "nfxidentity/protos/gen/clients/rate_limit"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"

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
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorHandler(),
			servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier()),
		),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	applicationpb.RegisterApplicationServiceServer(s, grpcHandler.NewApplicationHandler(d.AppAppSvc()))
	clientcredentialpb.RegisterClientCredentialServiceServer(s, grpcHandler.NewClientCredentialHandler(d.ClientCredentialAppSvc()))
	ipallowlistpb.RegisterIpAllowlistServiceServer(s, grpcHandler.NewIPAllowlistHandler(d.IPAllowlistAppSvc()))
	ratelimitpb.RegisterRateLimitServiceServer(s, grpcHandler.NewRateLimitHandler(d.RateLimitAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "clients"))

	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "clients"))

	return s
}
