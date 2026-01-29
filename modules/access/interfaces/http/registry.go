package http

import (
	"nfxid/modules/access/interfaces/http/handler"
)

type Registry struct {
	Role               *handler.RoleHandler
	Permission         *handler.PermissionHandler
	Scope              *handler.ScopeHandler
	Grant              *handler.GrantHandler
	RolePermission     *handler.RolePermissionHandler
	ScopePermission    *handler.ScopePermissionHandler
	Action             *handler.ActionHandler
	ActionRequirement  *handler.ActionRequirementHandler
}
