package image

import (
	"context"
	"fmt"

	imagetypepb "nfxid/protos/gen/image/image_type"
)

// ImageTypeClient ImageType 客户端
type ImageTypeClient struct {
	client imagetypepb.ImageTypeServiceClient
}

// NewImageTypeClient 创建 ImageType 客户端
func NewImageTypeClient(client imagetypepb.ImageTypeServiceClient) *ImageTypeClient {
	return &ImageTypeClient{client: client}
}

// GetImageTypeByID 根据ID获取图片类型
func (c *ImageTypeClient) GetImageTypeByID(ctx context.Context, id string) (*imagetypepb.ImageType, error) {
	req := &imagetypepb.GetImageTypeByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetImageTypeByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageType, nil
}

// GetImageTypeByName 根据名称获取图片类型
func (c *ImageTypeClient) GetImageTypeByName(ctx context.Context, name string) (*imagetypepb.ImageType, error) {
	req := &imagetypepb.GetImageTypeByNameRequest{
		Name: name,
	}

	resp, err := c.client.GetImageTypeByName(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageType, nil
}

// GetAllImageTypes 获取所有图片类型列表
func (c *ImageTypeClient) GetAllImageTypes(ctx context.Context, isSystem *bool) ([]*imagetypepb.ImageType, error) {
	req := &imagetypepb.GetAllImageTypesRequest{
		IsSystem: isSystem,
	}

	resp, err := c.client.GetAllImageTypes(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageTypes, nil
}