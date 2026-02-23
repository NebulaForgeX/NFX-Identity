package handler

import (
	"context"

	userApp "nfxid/modules/directory/application/users"
	userAppCommands "nfxid/modules/directory/application/users/commands"
	"nfxid/modules/directory/interfaces/grpc/mapper"
	userpb "nfxid/protos/gen/directory/user"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	userAppSvc *userApp.Service
}

func NewUserHandler(userAppSvc *userApp.Service) *UserHandler {
	return &UserHandler{userAppSvc: userAppSvc}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	userStatus := mapper.ProtoStatusToDomain(req.Status)
	cmd := userAppCommands.CreateUserCmd{
		Username:   req.Username,
		Status:     userStatus,
		IsVerified: req.IsVerified,
	}
	userID, err := h.userAppSvc.CreateUser(ctx, cmd)
	if err != nil {
		return nil, err
	}
	userView, err := h.userAppSvc.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	user := mapper.UserROToProto(&userView)
	return &userpb.CreateUserResponse{User: user}, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}
	userView, err := h.userAppSvc.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	user := mapper.UserROToProto(&userView)
	return &userpb.GetUserByIDResponse{User: user}, nil
}

func (h *UserHandler) GetUserByUsername(ctx context.Context, req *userpb.GetUserByUsernameRequest) (*userpb.GetUserByUsernameResponse, error) {
	userView, err := h.userAppSvc.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	user := mapper.UserROToProto(&userView)
	return &userpb.GetUserByUsernameResponse{User: user}, nil
}

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
			continue
		}
		user := mapper.UserROToProto(&userView)
		users = append(users, user)
	}
	return &userpb.BatchGetUsersResponse{Users: users}, nil
}
