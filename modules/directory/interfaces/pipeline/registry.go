package pipeline

import (
	"nfxid/modules/directory/interfaces/pipeline/handler"
)

type Registry struct {
	DirectoryHandler *handler.DirectoryHandler
}
