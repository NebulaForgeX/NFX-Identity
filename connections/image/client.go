package image

import (
	imagepb "nfxidentity/protos/gen/image/image"
	imagetagpb "nfxidentity/protos/gen/image/image_tag"
	imagetypepb "nfxidentity/protos/gen/image/image_type"
	imagevariantpb "nfxidentity/protos/gen/image/image_variant"
)

// Client Image 服务客户端
type Client struct {
	Image        *ImageClient
	ImageType    *ImageTypeClient
	ImageVariant *ImageVariantClient
	ImageTag     *ImageTagClient
}

// NewClient 创建 Image 客户端
func NewClient(
	imageClient imagepb.ImageServiceClient,
	imageTypeClient imagetypepb.ImageTypeServiceClient,
	imageVariantClient imagevariantpb.ImageVariantServiceClient,
	imageTagClient imagetagpb.ImageTagServiceClient,
) *Client {
	return &Client{
		Image:        NewImageClient(imageClient),
		ImageType:    NewImageTypeClient(imageTypeClient),
		ImageVariant: NewImageVariantClient(imageVariantClient),
		ImageTag:     NewImageTagClient(imageTagClient),
	}
}
