package handler

import (
	"context"

	userApp "nfxid/modules/auth/application/user"
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userpb "nfxid/protos/gen/auth/user"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	userAppSvc *userApp.Service
	userRepo   *userDomain.Repo
}

func NewUserHandler(userAppSvc *userApp.Service, userRepo *userDomain.Repo) *UserHandler {
	return &UserHandler{
		userAppSvc: userAppSvc,
		userRepo:   userRepo,
	}
}

// GetUserByID 根据ID获取用户
func (h *UserHandler) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userView, err := h.userAppSvc.GetUser(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	user := mapper.UserViewToProto(&userView)
	return &userpb.GetUserByIDResponse{User: user}, nil
}

// GetUserByUsername 根据用户名获取用户
func (h *UserHandler) GetUserByUsername(ctx context.Context, req *userpb.GetUserByUsernameRequest) (*userpb.GetUserByUsernameResponse, error) {
	entity, err := h.userRepo.Get.ByUsername(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	userView, err := h.userAppSvc.GetUser(ctx, entity.ID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	user := mapper.UserViewToProto(&userView)
	return &userpb.GetUserByUsernameResponse{User: user}, nil
}

// GetUserByEmail 根据邮箱获取用户
func (h *UserHandler) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (*userpb.GetUserByEmailResponse, error) {
	entity, err := h.userRepo.Get.ByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	userView, err := h.userAppSvc.GetUser(ctx, entity.ID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	user := mapper.UserViewToProto(&userView)
	return &userpb.GetUserByEmailResponse{User: user}, nil
}

// GetUserByPhone 根据手机号获取用户
func (h *UserHandler) GetUserByPhone(ctx context.Context, req *userpb.GetUserByPhoneRequest) (*userpb.GetUserByPhoneResponse, error) {
	entity, err := h.userRepo.Get.ByPhone(ctx, req.Phone)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	userView, err := h.userAppSvc.GetUser(ctx, entity.ID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	user := mapper.UserViewToProto(&userView)
	return &userpb.GetUserByPhoneResponse{User: user}, nil
}

// BatchGetUsers 批量获取用户
func (h *UserHandler) BatchGetUsers(ctx context.Context, req *userpb.BatchGetUsersRequest) (*userpb.BatchGetUsersResponse, error) {
	userIDs := make([]uuid.UUID, 0, len(req.UserIds))
	for _, idStr := range req.UserIds {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		userIDs = append(userIDs, id)
	}

	users := make([]*userpb.User, 0, len(userIDs))
	errorById := make(map[string]string)

	for _, userID := range userIDs {
		userView, err := h.userAppSvc.GetUser(ctx, userID)
		if err != nil {
			errorById[userID.String()] = err.Error()
			continue
		}
		user := mapper.UserViewToProto(&userView)
		users = append(users, user)
	}

	return &userpb.BatchGetUsersResponse{
		Users:     users,
		ErrorById: errorById,
	}, nil
}

// CheckUserExists 检查用户是否存在
func (h *UserHandler) CheckUserExists(ctx context.Context, req *userpb.CheckUserExistsRequest) (*userpb.CheckUserExistsResponse, error) {
	if req.Username != nil && *req.Username != "" {
		exists, err := h.userRepo.Check.ByUsername(ctx, *req.Username)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to check: %v", err)
		}
		if exists {
			return &userpb.CheckUserExistsResponse{Exists: true, Field: "username"}, nil
		}
	}

	if req.Email != nil && *req.Email != "" {
		exists, err := h.userRepo.Check.ByEmail(ctx, *req.Email)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to check: %v", err)
		}
		if exists {
			return &userpb.CheckUserExistsResponse{Exists: true, Field: "email"}, nil
		}
	}

	if req.Phone != nil && *req.Phone != "" {
		exists, err := h.userRepo.Check.ByPhone(ctx, *req.Phone)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to check: %v", err)
		}
		if exists {
			return &userpb.CheckUserExistsResponse{Exists: true, Field: "phone"}, nil
		}
	}

	return &userpb.CheckUserExistsResponse{Exists: false, Field: ""}, nil
}
