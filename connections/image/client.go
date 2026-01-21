package image

import (
	imagepb "nfxid/protos/gen/image/image"
	imagetagpb "nfxid/protos/gen/image/image_tag"
	imagetypepb "nfxid/protos/gen/image/image_type"
	imagevariantpb "nfxid/protos/gen/image/image_variant"
)

// Client Image 服务客户端
type Client struct {
	Image       *ImageClient
	ImageType   *ImageTypeClient
	ImageVariant *ImageVariantClient
	ImageTag    *ImageTagClient
}

// NewClient 创建 Image 客户端
func NewClient(
	imageClient imagepb.ImageServiceClient,
	imageTypeClient imagetypepb.ImageTypeServiceClient,
	imageVariantClient imagevariantpb.ImageVariantServiceClient,
	imageTagClient imagetagpb.ImageTagServiceClient,
) *Client {
	return &Client{
		Image:       NewImageClient(imageClient),
		ImageType:   NewImageTypeClient(imageTypeClient),
		ImageVariant: NewImageVariantClient(imageVariantClient),
		ImageTag:    NewImageTagClient(imageTagClient),
	}
}
