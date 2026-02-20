package http

import (
	"nfxid/modules/access/interfaces/http/handler"
)

type Registry struct {
	TenantRole *handler.TenantRoleHandler
}
