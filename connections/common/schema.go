package common

import (
	"context"

	schemapb "nfxid/protos/gen/common/schema"
)

// SchemaClient schema 清空客户端
type SchemaClient struct {
	client schemapb.SchemaServiceClient
}

// NewSchemaClient 创建 schema 清空客户端
func NewSchemaClient(client schemapb.SchemaServiceClient) *SchemaClient {
	return &SchemaClient{
		client: client,
	}
}

// ClearSchema 清空 schema
func (c *SchemaClient) ClearSchema(ctx context.Context) (*schemapb.ClearSchemaResponse, error) {
	return c.client.ClearSchema(ctx, &schemapb.ClearSchemaRequest{})
}
