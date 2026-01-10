package pipeline

import (
	"nfxid/modules/clients/interfaces/pipeline/handler"
)

type Registry struct {
	Clients *handler.ClientsHandler
}
