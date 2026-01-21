package image

import (
	"context"
	"fmt"

	imagevariantpb "nfxid/protos/gen/image/image_variant"
)

// ImageVariantClient ImageVariant 客户端
type ImageVariantClient struct {
	client imagevariantpb.ImageVariantServiceClient
}

// NewImageVariantClient 创建 ImageVariant 客户端
func NewImageVariantClient(client imagevariantpb.ImageVariantServiceClient) *ImageVariantClient {
	return &ImageVariantClient{client: client}
}

// GetImageVariantByID 根据ID获取图片变体
func (c *ImageVariantClient) GetImageVariantByID(ctx context.Context, id string) (*imagevariantpb.ImageVariant, error) {
	req := &imagevariantpb.GetImageVariantByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetImageVariantByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageVariant, nil
}

// GetImageVariantsByImageID 根据图片ID获取图片变体列表
func (c *ImageVariantClient) GetImageVariantsByImageID(ctx context.Context, imageID string, variantName *string) ([]*imagevariantpb.ImageVariant, error) {
	req := &imagevariantpb.GetImageVariantsByImageIDRequest{
		ImageId:     imageID,
		VariantName: variantName,
	}

	resp, err := c.client.GetImageVariantsByImageID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageVariants, nil
}