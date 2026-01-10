package http

import (
	"nfxid/modules/tenants/interfaces/http/handler"
)

type Registry struct {
	Tenant            *handler.TenantHandler
	Group             *handler.GroupHandler
	Member            *handler.MemberHandler
	Invitation        *handler.InvitationHandler
	TenantApp         *handler.TenantAppHandler
	TenantSetting     *handler.TenantSettingHandler
	DomainVerification *handler.DomainVerificationHandler
	MemberRole        *handler.MemberRoleHandler
	MemberGroup       *handler.MemberGroupHandler
	MemberAppRole     *handler.MemberAppRoleHandler
}
