package image

import (
	"context"
	"fmt"

	imagetagpb "nfxid/protos/gen/image/image_tag"
)

// ImageTagClient ImageTag 客户端
type ImageTagClient struct {
	client imagetagpb.ImageTagServiceClient
}

// NewImageTagClient 创建 ImageTag 客户端
func NewImageTagClient(client imagetagpb.ImageTagServiceClient) *ImageTagClient {
	return &ImageTagClient{client: client}
}

// GetImageTagByID 根据ID获取图片标签
func (c *ImageTagClient) GetImageTagByID(ctx context.Context, id string) (*imagetagpb.ImageTag, error) {
	req := &imagetagpb.GetImageTagByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetImageTagByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageTag, nil
}

// GetImageTagsByImageID 根据图片ID获取图片标签列表
func (c *ImageTagClient) GetImageTagsByImageID(ctx context.Context, imageID string) ([]*imagetagpb.ImageTag, error) {
	req := &imagetagpb.GetImageTagsByImageIDRequest{
		ImageId: imageID,
	}

	resp, err := c.client.GetImageTagsByImageID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageTags, nil
}

// GetImageTagsByTagName 根据标签名称获取图片标签列表
func (c *ImageTagClient) GetImageTagsByTagName(ctx context.Context, tagName string) ([]*imagetagpb.ImageTag, error) {
	req := &imagetagpb.GetImageTagsByTagNameRequest{
		TagName: tagName,
	}

	resp, err := c.client.GetImageTagsByTagName(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.ImageTags, nil
}