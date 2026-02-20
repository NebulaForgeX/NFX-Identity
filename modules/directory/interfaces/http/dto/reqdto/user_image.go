package reqdto

import (
	userImageAppCommands "nfxid/modules/directory/application/user_images/commands"

	"github.com/google/uuid"
)

type UserImageCreateRequestDTO struct {
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	ImageID      uuid.UUID `json:"image_id" validate:"required,uuid"`
	DisplayOrder int       `json:"display_order,omitempty"`
}

// UserImageByIDRequestDTO 用于 GetByID/SetPrimary/Update/Delete（路径 /user-images/:user_image_id）
type UserImageByIDRequestDTO struct {
	UserImageID uuid.UUID `uri:"user_image_id" validate:"required,uuid"`
}

type UserImageUpdateDisplayOrderRequestDTO struct {
	DisplayOrder int `json:"display_order" validate:"required"`
}

type UserImageUpdateImageIDRequestDTO struct {
	ImageID uuid.UUID `json:"image_id" validate:"required,uuid"`
}

// UserImagesDisplayOrderBatchItemRequestDTO 批量顺序项
type UserImagesDisplayOrderBatchItemRequestDTO struct {
	ID           uuid.UUID `json:"id" validate:"required,uuid"`
	DisplayOrder int       `json:"display_order" validate:"required"`
}

// UserImagesDisplayOrderBatchRequestDTO 批量更新用户图片显示顺序请求
type UserImagesDisplayOrderBatchRequestDTO struct {
	Order []UserImagesDisplayOrderBatchItemRequestDTO `json:"order" validate:"required,dive"`
}

func (r *UserImageCreateRequestDTO) ToCreateCmd() userImageAppCommands.CreateUserImageCmd {
	return userImageAppCommands.CreateUserImageCmd{
		UserID:       r.UserID,
		ImageID:      r.ImageID,
		DisplayOrder: r.DisplayOrder,
	}
}

func (r *UserImageUpdateDisplayOrderRequestDTO) ToUpdateDisplayOrderCmd(userImageID uuid.UUID) userImageAppCommands.UpdateUserImageDisplayOrderCmd {
	return userImageAppCommands.UpdateUserImageDisplayOrderCmd{
		UserImageID:  userImageID,
		DisplayOrder: r.DisplayOrder,
	}
}

func (r *UserImageUpdateImageIDRequestDTO) ToUpdateImageIDCmd(userImageID uuid.UUID) userImageAppCommands.UpdateUserImageImageIDCmd {
	return userImageAppCommands.UpdateUserImageImageIDCmd{
		UserImageID: userImageID,
		ImageID:     r.ImageID,
	}
}

func (r *UserImagesDisplayOrderBatchRequestDTO) ToBatchUpdateDisplayOrderCmd(userID uuid.UUID) userImageAppCommands.BatchUpdateDisplayOrderCmd {
	order := make([]userImageAppCommands.BatchUpdateDisplayOrderItem, len(r.Order))
	for i, item := range r.Order {
		order[i] = userImageAppCommands.BatchUpdateDisplayOrderItem{
			UserImageID:  item.ID,
			DisplayOrder: item.DisplayOrder,
		}
	}
	return userImageAppCommands.BatchUpdateDisplayOrderCmd{UserID: userID, Order: order}
}
