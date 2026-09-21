package handler

import (
	"context"

	"nfxidentity/modules/auth/application/account"
	"nfxidentity/modules/auth/interface/grpc/mapper"
	accountpb "nfxidentity/protos/gen/auth/account"
	forgerprofilepb "nfxidentity/protos/gen/auth/forger_profile"

	"github.com/google/uuid"
)

type ForgerHandler struct {
	forgerprofilepb.UnimplementedForgerProfileServiceServer
	svc *account.Service
}

func NewForgerHandler(svc *account.Service) *ForgerHandler {
	return &ForgerHandler{svc: svc}
}

func (h *ForgerHandler) ConfirmForgerProfileAvatar(ctx context.Context, req *forgerprofilepb.ConfirmForgerProfileAvatarRequest) (*forgerprofilepb.ConfirmForgerProfileAvatarResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	imageID, err := uuid.Parse(req.GetImageId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.ConfirmCommunityProfileAvatar(ctx, account.ConfirmCommunityProfileAvatarInput{
		AccountID: aid,
		ProfileID: pid,
		ImageID:   imageID,
	}); err != nil {
		return nil, err
	}
	return &forgerprofilepb.ConfirmForgerProfileAvatarResponse{}, nil
}

func (h *ForgerHandler) ConfirmForgerProfileBackgrounds(ctx context.Context, req *forgerprofilepb.ConfirmForgerProfileBackgroundsRequest) (*forgerprofilepb.ConfirmForgerProfileBackgroundsResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	images := make([]account.ConfirmCommunityProfileBackgroundItemInput, 0, len(req.GetImages()))
	for _, img := range req.GetImages() {
		imageID, err := uuid.Parse(img.GetImageId())
		if err != nil {
			return nil, err
		}
		images = append(images, account.ConfirmCommunityProfileBackgroundItemInput{
			ImageID:   imageID,
			SortOrder: int(img.GetSortOrder()),
		})
	}
	if _, err := h.svc.ConfirmCommunityProfileBackgrounds(ctx, account.ConfirmCommunityProfileBackgroundsInput{
		AccountID: aid,
		ProfileID: pid,
		Images:    images,
	}); err != nil {
		return nil, err
	}
	return &forgerprofilepb.ConfirmForgerProfileBackgroundsResponse{}, nil
}

func (h *ForgerHandler) BatchGetPublicProfileCards(ctx context.Context, req *forgerprofilepb.BatchGetPublicProfileCardsRequest) (*forgerprofilepb.BatchGetPublicProfileCardsResponse, error) {
	ids := make([]uuid.UUID, 0, len(req.GetProfileIds()))
	for _, raw := range req.GetProfileIds() {
		pid, err := uuid.Parse(raw)
		if err != nil {
			continue
		}
		ids = append(ids, pid)
	}
	out, err := h.svc.BatchGetCommunityPublicProfileCards(ctx, account.BatchGetCommunityPublicProfileCardsInput{ProfileIDs: ids})
	if err != nil {
		return nil, err
	}
	return &forgerprofilepb.BatchGetPublicProfileCardsResponse{Profiles: mapper.ForgerProfileItemsToPublicProfileCardProto(out.Profiles)}, nil
}

func (h *ForgerHandler) GetForgerRoles(ctx context.Context, req *forgerprofilepb.GetForgerRolesRequest) (*forgerprofilepb.GetForgerRolesResponse, error) {
	pid, err := uuid.Parse(req.GetProfileId())
	if err != nil {
		return nil, err
	}
	roles, err := h.svc.ForgerRoles(ctx, pid)
	if err != nil {
		return nil, err
	}
	out := make([]accountpb.ForgerRole, 0, len(roles))
	for _, r := range roles {
		if r == "forger" {
			out = append(out, accountpb.ForgerRole_FORGER_ROLE_FORGER)
		}
	}
	return &forgerprofilepb.GetForgerRolesResponse{ForgerRoles: out}, nil
}

func (h *ForgerHandler) HasForgerRole(ctx context.Context, req *forgerprofilepb.HasForgerRoleRequest) (*forgerprofilepb.HasForgerRoleResponse, error) {
	ok, err := h.svc.HasForgerRole(ctx, req.GetProfileId(), "forger")
	if err != nil {
		return nil, err
	}
	return &forgerprofilepb.HasForgerRoleResponse{Allowed: ok}, nil
}

func (h *ForgerHandler) IsForger(ctx context.Context, req *forgerprofilepb.HasForgerRoleRequest) (*forgerprofilepb.HasForgerRoleResponse, error) {
	return h.HasForgerRole(ctx, req)
}
