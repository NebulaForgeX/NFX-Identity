package http

import (
	"nfxidentity/modules/access/interfaces/http/handler"
)

type Registry struct {
	TenantRole *handler.TenantRoleHandler
}
