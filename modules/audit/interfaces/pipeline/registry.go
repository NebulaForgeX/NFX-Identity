package pipeline

import (
	"nfxid/modules/audit/interfaces/pipeline/handler"
)

type Registry struct {
	AuditHandler *handler.AuditHandler
}
