package handler

import (
	"context"

	schemapb "nfxidentity/protos/gen/common/schema"
)

type SystemSchemaClearFn func(ctx context.Context) (tablesCleared int32, err error)

type SystemSchemaHandler struct {
	schemapb.UnimplementedSchemaServiceServer
	clear SystemSchemaClearFn
}

func NewSystemSchemaHandler(clear SystemSchemaClearFn) *SystemSchemaHandler {
	return &SystemSchemaHandler{clear: clear}
}

func (h *SystemSchemaHandler) ClearSchema(ctx context.Context, req *schemapb.ClearSchemaRequest) (*schemapb.ClearSchemaResponse, error) {
	tablesCleared, err := h.clear(ctx)
	if err != nil {
		errMsg := err.Error()
		return &schemapb.ClearSchemaResponse{
			Success:       false,
			ErrorMessage:  &errMsg,
			TablesCleared: 0,
		}, nil
	}
	return &schemapb.ClearSchemaResponse{
		Success:       true,
		TablesCleared: tablesCleared,
	}, nil
}
