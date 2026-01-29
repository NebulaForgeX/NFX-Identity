package user_images

import (
	"context"
	userImageCommands "nfxid/modules/directory/application/user_images/commands"
)

// SetPrimaryUserImage 将指定用户图片设为主图（display_order = 0），其余按原顺序重排为 1, 2, ...
func (s *Service) SetPrimaryUserImage(ctx context.Context, cmd userImageCommands.SetPrimaryUserImageCmd) error {
	ui, err := s.userImageRepo.Get.ByID(ctx, cmd.UserImageID)
	if err != nil {
		return err
	}
	userID := ui.UserID()

	list, err := s.userImageRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return err
	}

	var targetIndex int
	for i, item := range list {
		if item.ID() == cmd.UserImageID {
			targetIndex = i
			break
		}
	}

	// 新顺序：目标为 0，其余按原相对顺序为 1, 2, ...
	for i := range list {
		var newOrder int
		if i == targetIndex {
			newOrder = 0
		} else if i < targetIndex {
			newOrder = i + 1
		} else {
			newOrder = i
		}
		if err := s.userImageRepo.Update.DisplayOrder(ctx, list[i].ID(), newOrder); err != nil {
			return err
		}
	}
	return nil
}
