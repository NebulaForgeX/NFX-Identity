package pipeline

import (
	"nfxid/modules/tenants/interfaces/pipeline/handler"
)

type Registry struct {
	TenantsHandler *handler.TenantsHandler
}
