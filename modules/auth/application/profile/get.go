package profile

import (
	"context"
	profileViews "nfxid/modules/auth/application/profile/views"
	profileDomain "nfxid/modules/auth/domain/profile"
	"nfxid/pkgs/logx"
	imagepb "nfxid/protos/gen/image/image"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetProfile(ctx context.Context, profileID uuid.UUID) (profileViews.ProfileView, error) {
	domainView, err := s.profileQuery.Single.ByID(ctx, profileID)
	if err != nil {
		return profileViews.ProfileView{}, err
	}
	view := profileViews.ProfileViewMapper(*domainView)

	// 通过 gRPC 调用 image 服务获取用户的所有图片
	if err := s.enrichWithImages(ctx, &view); err != nil {
		logx.S().Warnf("failed to enrich profile with images: %v", err)
		// 不返回错误，只是记录警告，继续返回 profile
	}

	return view, nil
}

func (s *Service) GetProfileByUserID(ctx context.Context, userID uuid.UUID) (profileViews.ProfileView, error) {
	domainView, err := s.profileQuery.Single.ByUserID(ctx, userID)
	if err != nil {
		return profileViews.ProfileView{}, err
	}
	view := profileViews.ProfileViewMapper(*domainView)

	// 通过 gRPC 调用 image 服务获取用户的所有图片
	if err := s.enrichWithImages(ctx, &view); err != nil {
		logx.S().Warnf("failed to enrich profile with images: %v", err)
		// 不返回错误，只是记录警告，继续返回 profile
	}

	return view, nil
}

// enrichWithImages 通过 gRPC 调用 image 服务获取用户的所有图片
func (s *Service) enrichWithImages(ctx context.Context, view *profileViews.ProfileView) error {
	if s.imageGRPCClient == nil {
		return nil // 如果没有配置 image client，直接返回
	}

	// 根据 user_id 获取该用户的所有图片
	resp, err := s.imageGRPCClient.ImageStub.GetImagesByUserID(ctx, &imagepb.GetImagesByUserIDRequest{
		UserId:   view.UserID.String(),
		Page:     1,
		PageSize: 100, // 获取前100张图片
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			// 没有图片，不报错
			return nil
		}
		return err
	}

	// 将图片列表添加到 view 中
	if resp.Images != nil && len(resp.Images) > 0 {
		images := make([]profileViews.ImageInfo, 0, len(resp.Images))
		for _, img := range resp.Images {
			imageInfo := profileViews.ImageInfo{
				ID:       img.Id,
				URL:      getStringPtr(img.Url),
				IsPublic: img.IsPublic,
			}
			if img.TypeId != nil {
				imageInfo.TypeID = *img.TypeId
			}
			images = append(images, imageInfo)
		}
		view.Images = images
	}

	return nil
}

func getStringPtr(s *string) *string {
	if s == nil {
		return nil
	}
	return s
}

type GetProfileListResult struct {
	Items []profileViews.ProfileView
	Total int64
}

func (s *Service) GetProfileList(ctx context.Context, q profileDomain.ListQuery) (GetProfileListResult, error) {
	q.Normalize()
	domainViews, total, err := s.profileQuery.List.Generic(ctx, q)
	if err != nil {
		return GetProfileListResult{}, err
	}
	items := make([]profileViews.ProfileView, len(domainViews))
	for i, v := range domainViews {
		items[i] = profileViews.ProfileViewMapper(*v)
	}
	return GetProfileListResult{
		Items: items,
		Total: total,
	}, nil
}
