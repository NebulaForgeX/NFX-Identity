package handler

import (
	"context"
	"fmt"

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
	// 执行清空
	tablesCleared, err := postgresqlx.ClearSchema(ctx, h.db, h.schemaName, nil)
	if err != nil {
		errMsg := err.Error()
		return &schemapb.ClearSchemaResponse{
			Success:       false,
			ErrorMessage:  &errMsg,
			TablesCleared: 0,
		}, nil
	}

	// 验证清空结果：检查 roles 表是否为空（如果存在）
	var rolesCount int64
	if err := h.db.WithContext(ctx).Table(`"access"."roles"`).Count(&rolesCount).Error; err == nil {
		if rolesCount > 0 {
			errMsg := fmt.Sprintf("roles table still has %d records after clear", rolesCount)
			return &schemapb.ClearSchemaResponse{
				Success:       false,
				ErrorMessage:  &errMsg,
				TablesCleared: int32(tablesCleared),
			}, nil
		}
	}

	return &schemapb.ClearSchemaResponse{
		Success:       true,
		TablesCleared: int32(tablesCleared),
	}, nil
}
