package pipeline

import (
	"nfxidentity/modules/audit/interfaces/pipeline/handler"
)

type Registry struct {
	AuditHandler *handler.AuditHandler
}
