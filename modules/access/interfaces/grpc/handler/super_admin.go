package handler

import (
	"context"

	superadminpb "nfxid/protos/gen/access/super_admin"
	superadminsApp "nfxid/modules/access/application/super_admins"
	domain "nfxid/modules/access/domain/super_admins"
	"nfxid/modules/access/interfaces/grpc/mapper"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SuperAdminHandler 实现 SuperAdminServiceServer
type SuperAdminHandler struct {
	superadminpb.UnimplementedSuperAdminServiceServer
	svc *superadminsApp.Service
}

// NewSuperAdminHandler 创建 handler
func NewSuperAdminHandler(svc *superadminsApp.Service) *SuperAdminHandler {
	return &SuperAdminHandler{svc: svc}
}

func (h *SuperAdminHandler) GetSuperAdminByUserID(ctx context.Context, req *superadminpb.GetSuperAdminByUserIDRequest) (*superadminpb.GetSuperAdminByUserIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user_id")
	}
	s, err := h.svc.GetByUserID(ctx, userID)
	if err != nil {
		if err == domain.ErrSuperAdminNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.Internal, err.Error())
	}
	out := make([]*superadminpb.SuperAdmin, len(list))
	for i := range list {
		out[i] = mapper.SuperAdminDomainToProto(list[i])
	}
	return &superadminpb.ListSuperAdminsResponse{SuperAdmins: out}, nil
}

func (h *SuperAdminHandler) CreateSuperAdmin(ctx context.Context, req *superadminpb.CreateSuperAdminRequest) (*superadminpb.CreateSuperAdminResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id required")
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user_id")
	}
	if err := h.svc.Create(ctx, userID); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &superadminpb.CreateSuperAdminResponse{}, nil
}
