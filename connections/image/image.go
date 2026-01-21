package image

import (
	"context"
	"fmt"

	imagepb "nfxid/protos/gen/image/image"
)

// ImageClient Image 客户端
type ImageClient struct {
	client imagepb.ImageServiceClient
}

// NewImageClient 创建 Image 客户端
func NewImageClient(client imagepb.ImageServiceClient) *ImageClient {
	return &ImageClient{client: client}
}

// GetImageByID 根据ID获取图片
func (c *ImageClient) GetImageByID(ctx context.Context, id string) (*imagepb.Image, error) {
	req := &imagepb.GetImageByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetImageByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Image, nil
}

// GetImageByImageID 根据图片ID获取图片
func (c *ImageClient) GetImageByImageID(ctx context.Context, imageID string) (*imagepb.Image, error) {
	req := &imagepb.GetImageByImageIDRequest{
		ImageId: imageID,
	}

	resp, err := c.client.GetImageByImageID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Image, nil
}

// BatchGetImages 批量获取图片
func (c *ImageClient) BatchGetImages(ctx context.Context, ids []string) ([]*imagepb.Image, error) {
	req := &imagepb.BatchGetImagesRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetImages(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Images, nil
}