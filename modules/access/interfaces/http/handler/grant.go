package handler

import (
	grantApp "nfxid/modules/access/application/grants"
	grantAppCommands "nfxid/modules/access/application/grants/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type GrantHandler struct {
	appSvc *grantApp.Service
}

func NewGrantHandler(appSvc *grantApp.Service) *GrantHandler {
	return &GrantHandler{
		appSvc: appSvc,
	}
}

// Create 创建授权
func (h *GrantHandler) Create(c fiber.Ctx) error {
	var req reqdto.GrantCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	grantID, err := h.appSvc.CreateGrant(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created grant
	grantView, err := h.appSvc.GetGrant(c.Context(), grantID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Grant created successfully", httpx.SuccessOptions{Data: respdto.GrantROToDTO(&grantView)})
}

// GetByID 根据 ID 获取授权
func (h *GrantHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.GrantByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetGrant(c.Context(), req.GrantID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Grant retrieved successfully", httpx.SuccessOptions{Data: respdto.GrantROToDTO(&result)})
}

// Update 更新授权
func (h *GrantHandler) Update(c fiber.Ctx) error {
	var req reqdto.GrantUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateGrant(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Grant updated successfully")
}

// Revoke 撤销授权
func (h *GrantHandler) Revoke(c fiber.Ctx) error {
	var req reqdto.GrantRevokeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeGrant(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Grant revoked successfully")
}

// Delete 删除授权
func (h *GrantHandler) Delete(c fiber.Ctx) error {
	var req reqdto.GrantByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := grantAppCommands.DeleteGrantCmd{GrantID: req.GrantID}
	if err := h.appSvc.DeleteGrant(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Grant deleted successfully")
}

// GetBySubject 根据主体获取授权列表
func (h *GrantHandler) GetBySubject(c fiber.Ctx) error {
	var req reqdto.GrantBySubjectRequestDTO
	if err := c.Bind().Query(&req); err != nil {
		return errx.ErrInvalidQuery.WithCause(err)
	}

	// 使用 application 层的方法，避免 handler 直接依赖 domain 层
	grants, err := h.appSvc.GetGrantsBySubjectString(c.Context(), req.SubjectType, req.SubjectID)
	if err != nil {
		return err
	}

	dtos := make([]*respdto.GrantDTO, len(grants))
	for i, g := range grants {
		dtos[i] = respdto.GrantROToDTO(&g)
	}

	return fiberx.OK(c, "Grants retrieved successfully", httpx.SuccessOptions{Data: dtos})
}

// fiber:context-methods migrated
