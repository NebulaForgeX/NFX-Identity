package handler

import (
	"context"

	userApp "nfxid/modules/directory/application/users"
	userAppCommands "nfxid/modules/directory/application/users/commands"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	"nfxid/pkgs/logx"
	userpb "nfxid/protos/gen/directory/user"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	userAppSvc *userApp.Service
}

func NewUserHandler(userAppSvc *userApp.Service) *UserHandler {
	return &UserHandler{
		userAppSvc: userAppSvc,
	}
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// 转换 protobuf 状态到 domain 状态
	userStatus := mapper.ProtoStatusToDomain(req.Status)

	// 创建命令
	cmd := userAppCommands.CreateUserCmd{
		Username:   req.Username,
		Status:     userStatus,
		IsVerified: req.IsVerified,
	}

	// 调用应用服务创建用户
	userID, err := h.userAppSvc.CreateUser(ctx, cmd)
	if err != nil {
		logx.S().Errorf("failed to create user: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	// 获取创建的用户
	userView, err := h.userAppSvc.GetUser(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get created user: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get created user: %v", err)
	}

	// 转换为 protobuf 响应
	user := mapper.UserROToProto(&userView)
	return &userpb.CreateUserResponse{User: user}, nil
}

// GetUserByID 根据ID获取用户
func (h *UserHandler) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}

	userView, err := h.userAppSvc.GetUser(ctx, userID)
	if err != nil {
		logx.S().Errorf("failed to get user by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	user := mapper.UserROToProto(&userView)
	return &userpb.GetUserByIDResponse{User: user}, nil
}

// GetUserByUsername 根据用户名获取用户
func (h *UserHandler) GetUserByUsername(ctx context.Context, req *userpb.GetUserByUsernameRequest) (*userpb.GetUserByUsernameResponse, error) {
	userView, err := h.userAppSvc.GetUserByUsername(ctx, req.Username)
	if err != nil {
		logx.S().Errorf("failed to get user by username: %v", err)
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	user := mapper.UserROToProto(&userView)
	return &userpb.GetUserByUsernameResponse{User: user}, nil
}

// BatchGetUsers 批量获取用户
func (h *UserHandler) BatchGetUsers(ctx context.Context, req *userpb.BatchGetUsersRequest) (*userpb.BatchGetUsersResponse, error) {
	userIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		userIDs = append(userIDs, id)
	}

	users := make([]*userpb.User, 0, len(userIDs))
	for _, userID := range userIDs {
		userView, err := h.userAppSvc.GetUser(ctx, userID)
		if err != nil {
			logx.S().Warnf("failed to get user %s: %v", userID, err)
			continue
		}
		user := mapper.UserROToProto(&userView)
		users = append(users, user)
	}

	return &userpb.BatchGetUsersResponse{Users: users}, nil
}
