package grpc

import (
	"context"
	"fmt"
	"sync"

	"nfxid/connections/access"
	"nfxid/connections/audit"
	"nfxid/connections/auth"
	"nfxid/connections/clients"
	"nfxid/connections/directory"
	"nfxid/connections/image"
	"nfxid/connections/tenants"
	"nfxid/modules/system/config"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"

	grantpb "nfxid/protos/gen/access/grant"
	permissionpb "nfxid/protos/gen/access/permission"
	rolepb "nfxid/protos/gen/access/role"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	scopepb "nfxid/protos/gen/access/scope"
	scopepermissionpb "nfxid/protos/gen/access/scope_permission"
	actorsnapshotpb "nfxid/protos/gen/audit/actor_snapshot"
	eventpb "nfxid/protos/gen/audit/event"
	eventretentionpolicypb "nfxid/protos/gen/audit/event_retention_policy"
	eventsearchindexpb "nfxid/protos/gen/audit/event_search_index"
	hashchaincheckpointpb "nfxid/protos/gen/audit/hash_chain_checkpoint"
	accountlockoutpb "nfxid/protos/gen/auth/account_lockout"
	loginattemptpb "nfxid/protos/gen/auth/login_attempt"
	mfafactorpb "nfxid/protos/gen/auth/mfa_factor"
	passwordhistorypb "nfxid/protos/gen/auth/password_history"
	passwordresetpb "nfxid/protos/gen/auth/password_reset"
	refreshtokenpb "nfxid/protos/gen/auth/refresh_token"
	sessionpb "nfxid/protos/gen/auth/session"
	trusteddevicepb "nfxid/protos/gen/auth/trusted_device"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"
	apikeypb "nfxid/protos/gen/clients/api_key"
	apppb "nfxid/protos/gen/clients/app"
	clientcredentialpb "nfxid/protos/gen/clients/client_credential"
	clientscopepb "nfxid/protos/gen/clients/client_scope"
	ipallowlistpb "nfxid/protos/gen/clients/ip_allowlist"
	ratelimitpb "nfxid/protos/gen/clients/rate_limit"
	badgepb "nfxid/protos/gen/directory/badge"
	userpb "nfxid/protos/gen/directory/user"
	userbadgepb "nfxid/protos/gen/directory/user_badge"
	usereducationpb "nfxid/protos/gen/directory/user_education"
	useremailpb "nfxid/protos/gen/directory/user_email"
	useroccupationpb "nfxid/protos/gen/directory/user_occupation"
	userphonepb "nfxid/protos/gen/directory/user_phone"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"
	userprofilepb "nfxid/protos/gen/directory/user_profile"
	imagepb "nfxid/protos/gen/image/image"
	imagetagpb "nfxid/protos/gen/image/image_tag"
	imagetypepb "nfxid/protos/gen/image/image_type"
	imagevariantpb "nfxid/protos/gen/image/image_variant"
	domainverificationpb "nfxid/protos/gen/tenants/domain_verification"
	grouppb "nfxid/protos/gen/tenants/group"
	invitationpb "nfxid/protos/gen/tenants/invitation"
	memberpb "nfxid/protos/gen/tenants/member"
	memberapprolepb "nfxid/protos/gen/tenants/member_app_role"
	membergrouppb "nfxid/protos/gen/tenants/member_group"
	memberrolepb "nfxid/protos/gen/tenants/member_role"
	tenantpb "nfxid/protos/gen/tenants/tenant"
	tenantapppb "nfxid/protos/gen/tenants/tenant_app"
	tenantsettingpb "nfxid/protos/gen/tenants/tenant_setting"

	"google.golang.org/grpc"
)

// GRPCClients gRPC 客户端集合
// 用于系统初始化时调用其他服务
type GRPCClients struct {
	DirectoryClient *directory.Client
	AccessClient    *access.Client
	AuthClient      *auth.Client
	AuditClient     *audit.Client
	ClientsClient   *clients.Client
	ImageClient     *image.Client
	TenantsClient   *tenants.Client

	conns []*grpc.ClientConn
	mu    sync.Mutex
}

// NewClients 创建 gRPC 客户端连接
func NewGRPCClients(ctx context.Context, cfg *config.GRPCClientConfig, tokenCfg *tokenx.Config) (*GRPCClients, error) {
	// 创建 server token provider
	tokenProvider := servertoken.NewProvider(
		&servertoken.HMACSigner{Key: []byte(tokenCfg.SecretKey)},
		tokenCfg.Issuer,
		"system-service", // service ID
	)

	grpcClients := &GRPCClients{
		conns: make([]*grpc.ClientConn, 0),
	}

	// 连接 Directory 服务
	directoryConn, err := grpc.NewClient(
		cfg.DirectoryAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory client: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, directoryConn)
	grpcClients.DirectoryClient = directory.NewClient(
		userpb.NewUserServiceClient(directoryConn),
		userprofilepb.NewUserProfileServiceClient(directoryConn),
		useremailpb.NewUserEmailServiceClient(directoryConn),
		userphonepb.NewUserPhoneServiceClient(directoryConn),
		userpreferencepb.NewUserPreferenceServiceClient(directoryConn),
		usereducationpb.NewUserEducationServiceClient(directoryConn),
		useroccupationpb.NewUserOccupationServiceClient(directoryConn),
		badgepb.NewBadgeServiceClient(directoryConn),
		userbadgepb.NewUserBadgeServiceClient(directoryConn),
	)

	// 连接 Access 服务
	accessConn, err := grpc.NewClient(
		cfg.AccessAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create access client: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, accessConn)
	grpcClients.AccessClient = access.NewClient(
		rolepb.NewRoleServiceClient(accessConn),
		permissionpb.NewPermissionServiceClient(accessConn),
		grantpb.NewGrantServiceClient(accessConn),
		rolepermissionpb.NewRolePermissionServiceClient(accessConn),
		scopepb.NewScopeServiceClient(accessConn),
		scopepermissionpb.NewScopePermissionServiceClient(accessConn),
	)

	// 连接 Auth 服务
	authConn, err := grpc.NewClient(
		cfg.AuthAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create auth client: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, authConn)
	grpcClients.AuthClient = auth.NewClient(
		usercredentialpb.NewUserCredentialServiceClient(authConn),
		sessionpb.NewSessionServiceClient(authConn),
		trusteddevicepb.NewTrustedDeviceServiceClient(authConn),
		mfafactorpb.NewMfaFactorServiceClient(authConn),
		refreshtokenpb.NewRefreshTokenServiceClient(authConn),
		passwordresetpb.NewPasswordResetServiceClient(authConn),
		passwordhistorypb.NewPasswordHistoryServiceClient(authConn),
		loginattemptpb.NewLoginAttemptServiceClient(authConn),
		accountlockoutpb.NewAccountLockoutServiceClient(authConn),
	)

	// 连接 Audit 服务
	auditConn, err := grpc.NewClient(
		cfg.AuditAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create audit client: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, auditConn)
	grpcClients.AuditClient = audit.NewClient(
		eventpb.NewEventServiceClient(auditConn),
		eventsearchindexpb.NewEventSearchIndexServiceClient(auditConn),
		actorsnapshotpb.NewActorSnapshotServiceClient(auditConn),
		hashchaincheckpointpb.NewHashChainCheckpointServiceClient(auditConn),
		eventretentionpolicypb.NewEventRetentionPolicyServiceClient(auditConn),
	)

	// 连接 Clients 服务
	clientsConn, err := grpc.NewClient(
		cfg.ClientsAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create clients client: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, clientsConn)
	grpcClients.ClientsClient = clients.NewClient(
		apppb.NewAppServiceClient(clientsConn),
		apikeypb.NewApiKeyServiceClient(clientsConn),
		clientcredentialpb.NewClientCredentialServiceClient(clientsConn),
		clientscopepb.NewClientScopeServiceClient(clientsConn),
		ipallowlistpb.NewIpAllowlistServiceClient(clientsConn),
		ratelimitpb.NewRateLimitServiceClient(clientsConn),
	)

	// 连接 Image 服务
	imageConn, err := grpc.NewClient(
		cfg.ImageAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create image client: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, imageConn)
	grpcClients.ImageClient = image.NewClient(
		imagepb.NewImageServiceClient(imageConn),
		imagetypepb.NewImageTypeServiceClient(imageConn),
		imagevariantpb.NewImageVariantServiceClient(imageConn),
		imagetagpb.NewImageTagServiceClient(imageConn),
	)

	// 连接 Tenants 服务
	tenantsConn, err := grpc.NewClient(
		cfg.TenantsAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create tenants client: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, tenantsConn)
	grpcClients.TenantsClient = tenants.NewClient(
		tenantpb.NewTenantServiceClient(tenantsConn),
		tenantapppb.NewTenantAppServiceClient(tenantsConn),
		tenantsettingpb.NewTenantSettingServiceClient(tenantsConn),
		memberpb.NewMemberServiceClient(tenantsConn),
		grouppb.NewGroupServiceClient(tenantsConn),
		invitationpb.NewInvitationServiceClient(tenantsConn),
		membergrouppb.NewMemberGroupServiceClient(tenantsConn),
		memberrolepb.NewMemberRoleServiceClient(tenantsConn),
		memberapprolepb.NewMemberAppRoleServiceClient(tenantsConn),
		domainverificationpb.NewDomainVerificationServiceClient(tenantsConn),
	)

	return grpcClients, nil
}

// Close 关闭所有 gRPC 连接
func (c *GRPCClients) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var errs []error
	for _, conn := range c.conns {
		if err := conn.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing gRPC connections: %v", errs)
	}

	return nil
}
