package pipeline

import (
	"nfxidentity/modules/tenants/interfaces/pipeline/handler"
)

type Registry struct {
	TenantsHandler *handler.TenantsHandler
}
