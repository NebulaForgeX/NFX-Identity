package handler

import (
	"context"

	"nfxid/pkgs/postgresqlx"
	schemapb "nfxid/protos/gen/common/schema"
	"gorm.io/gorm"
)

// SchemaHandler schema 清空处理器
type SchemaHandler struct {
	schemapb.UnimplementedSchemaServiceServer
	db         *gorm.DB
	schemaName string
}

// NewSchemaHandler 创建 schema 处理器
func NewSchemaHandler(db *gorm.DB, schemaName string) *SchemaHandler {
	return &SchemaHandler{
		db:         db,
		schemaName: schemaName,
	}
}

// ClearSchema 清空 schema 中所有表的数据（不删除表）
func (h *SchemaHandler) ClearSchema(ctx context.Context, req *schemapb.ClearSchemaRequest) (*schemapb.ClearSchemaResponse, error) {
	tablesCleared, err := postgresqlx.ClearSchema(ctx, h.db, h.schemaName, nil)
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
		TablesCleared: int32(tablesCleared),
	}, nil
}
