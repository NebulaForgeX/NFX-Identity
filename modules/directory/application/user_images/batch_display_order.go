package user_images

import (
	"context"
	"fmt"
	userImageCommands "nfxid/modules/directory/application/user_images/commands"
)

// UpdateUserImagesDisplayOrderBatch 批量更新用户图片显示顺序
func (s *Service) UpdateUserImagesDisplayOrderBatch(ctx context.Context, cmd userImageCommands.BatchUpdateDisplayOrderCmd) error {
	list, err := s.userImageRepo.Get.ByUserID(ctx, cmd.UserID)
	if err != nil {
		return err
	}
	idToOrder := make(map[string]int)
	for _, item := range cmd.Order {
		idToOrder[item.UserImageID.String()] = item.DisplayOrder
	}
	for _, ui := range list {
		order, ok := idToOrder[ui.ID().String()]
		if !ok {
			return fmt.Errorf("user image %s not in order list", ui.ID())
		}
		if err := s.userImageRepo.Update.DisplayOrder(ctx, ui.ID(), order); err != nil {
			return err
		}
	}
	return nil
}
