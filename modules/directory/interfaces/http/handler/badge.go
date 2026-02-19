package handler

import (
	badgeApp "nfxid/modules/directory/application/badges"
	badgeAppCommands "nfxid/modules/directory/application/badges/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type BadgeHandler struct {
	appSvc *badgeApp.Service
}

func NewBadgeHandler(appSvc *badgeApp.Service) *BadgeHandler {
	return &BadgeHandler{
		appSvc: appSvc,
	}
}

// Create 创建徽章
func (h *BadgeHandler) Create(c fiber.Ctx) error {
	var req reqdto.BadgeCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	badgeID, err := h.appSvc.CreateBadge(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created badge
	badgeView, err := h.appSvc.GetBadge(c.Context(), badgeID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Badge created successfully", httpx.SuccessOptions{Data: respdto.BadgeROToDTO(&badgeView)})
}

// GetByID 根据 ID 获取徽章
func (h *BadgeHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.BadgeByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetBadge(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Badge retrieved successfully", httpx.SuccessOptions{Data: respdto.BadgeROToDTO(&result)})
}

// GetByName 根据 Name 获取徽章
func (h *BadgeHandler) GetByName(c fiber.Ctx) error {
	var req reqdto.BadgeByNameRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetBadgeByName(c.Context(), req.Name)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Badge retrieved successfully", httpx.SuccessOptions{Data: respdto.BadgeROToDTO(&result)})
}

// Update 更新徽章
func (h *BadgeHandler) Update(c fiber.Ctx) error {
	var req reqdto.BadgeUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateBadge(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Badge updated successfully")
}

// Delete 删除徽章（软删除）
func (h *BadgeHandler) Delete(c fiber.Ctx) error {
	var req reqdto.BadgeByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := badgeAppCommands.DeleteBadgeCmd{BadgeID: req.ID}
	if err := h.appSvc.DeleteBadge(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Badge deleted successfully")
}

// fiber:context-methods migrated
