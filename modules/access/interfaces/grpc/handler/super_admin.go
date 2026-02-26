package handler

import (
	"context"

	superadminsApp "nfxidentity/modules/access/application/super_admins"
	"nfxidentity/modules/access/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	superadminpb "nfxidentity/protos/gen/access/super_admin"

	"github.com/google/uuid"
)

type SuperAdminHandler struct {
	superadminpb.UnimplementedSuperAdminServiceServer
	svc *superadminsApp.Service
}

func NewSuperAdminHandler(svc *superadminsApp.Service) *SuperAdminHandler {
	return &SuperAdminHandler{svc: svc}
}

func (h *SuperAdminHandler) GetSuperAdminByUserID(
	ctx context.Context,
	req *superadminpb.GetSuperAdminByUserIDRequest,
) (*superadminpb.GetSuperAdminByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	s, err := h.svc.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &superadminpb.GetSuperAdminByUserIDResponse{
		SuperAdmin: mapper.SuperAdminDomainToProto(s),
	}, nil
}

func (h *SuperAdminHandler) ListSuperAdmins(ctx context.Context, req *superadminpb.ListSuperAdminsRequest) (*superadminpb.ListSuperAdminsResponse, error) {
	limit, offset := 100, 0
	if req.Limit != nil {
		limit = int(*req.Limit)
	}
	if req.Offset != nil {
		offset = int(*req.Offset)
	}
	list, err := h.svc.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	out := make([]*superadminpb.SuperAdmin, len(list))
	for i := range list {
		out[i] = mapper.SuperAdminDomainToProto(list[i])
	}
	return &superadminpb.ListSuperAdminsResponse{SuperAdmins: out}, nil
}

func (h *SuperAdminHandler) CreateSuperAdmin(ctx context.Context, req *superadminpb.CreateSuperAdminRequest) (*superadminpb.CreateSuperAdminResponse, error) {
	if req.UserId == "" {
		return nil, errx.InvalidArg("USER_ID_REQUIRED", "user_id required")
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	if err := h.svc.Create(ctx, userID); err != nil {
		return nil, err
	}
	return &superadminpb.CreateSuperAdminResponse{}, nil
}
