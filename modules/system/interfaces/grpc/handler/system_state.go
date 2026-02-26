package handler

import (
	"context"

	systemStateApp "nfxidentity/modules/system/application/system_state"
	"nfxidentity/modules/system/interfaces/grpc/mapper"
	"nfxidentity/pkgs/errx"
	systemstatepb "nfxidentity/protos/gen/system/system_state"

	"github.com/google/uuid"
)

type SystemStateHandler struct {
	systemstatepb.UnimplementedSystemStateServiceServer
	appSvc *systemStateApp.Service
}

func NewSystemStateHandler(appSvc *systemStateApp.Service) *SystemStateHandler {
	return &SystemStateHandler{
		appSvc: appSvc,
	}
}

// GetSystemStateByID 根据ID获取系统状态
func (h *SystemStateHandler) GetSystemStateByID(
	ctx context.Context,
	req *systemstatepb.GetSystemStateByIDRequest,
) (*systemstatepb.GetSystemStateByIDResponse, error) {
	systemStateID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	systemStateView, err := h.appSvc.GetSystemState(ctx, systemStateID)
	if err != nil {
		return nil, err
	}

	systemState := mapper.SystemStateROToProto(&systemStateView)
	return &systemstatepb.GetSystemStateByIDResponse{SystemState: systemState}, nil
}

// GetSystemStateByKey 根据键获取系统状态（目前使用 Latest，因为 service 没有 ByKey 方法）
func (h *SystemStateHandler) GetSystemStateByKey(
	ctx context.Context,
	req *systemstatepb.GetSystemStateByKeyRequest,
) (*systemstatepb.GetSystemStateByKeyResponse, error) {
	// TODO: 如果 service 有 ByKey 方法，使用它；否则使用 Latest
	systemStateView, err := h.appSvc.GetLatestSystemState(ctx)
	if err != nil {
		return nil, err
	}

	systemState := mapper.SystemStateROToProto(&systemStateView)
	return &systemstatepb.GetSystemStateByKeyResponse{SystemState: systemState}, nil
}

// GetAllSystemStates 获取所有系统状态列表
func (h *SystemStateHandler) GetAllSystemStates(
	ctx context.Context,
	req *systemstatepb.GetAllSystemStatesRequest,
) (*systemstatepb.GetAllSystemStatesResponse, error) {
	// TODO: 如果 service 有 GetAll 方法，使用它；否则返回错误
	return nil, errx.FailedPrecond("UNIMPLEMENTED", "GetAllSystemStates not implemented yet")
}
