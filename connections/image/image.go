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

// MoveImage 移动图片（从 tmp 移动到目标目录，如 avatar/background）
func (c *ImageClient) MoveImage(ctx context.Context, id string, targetType string) (*imagepb.Image, error) {
	req := &imagepb.MoveImageRequest{
		Id:         id,
		TargetType: targetType,
	}

	resp, err := c.client.MoveImage(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Image, nil
}

// DeleteImage 删除图片（如更换头像时删除旧头像文件）
func (c *ImageClient) DeleteImage(ctx context.Context, id string) error {
	req := &imagepb.DeleteImageRequest{
		Id: id,
	}

	_, err := c.client.DeleteImage(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC call failed: %w", err)
	}

	return nil
}
