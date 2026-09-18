package handler

import (
	"context"

	"nfxidentity/modules/auth/application/platform"
	accountpb "nfxidentity/protos/gen/auth/account"
	forgerprofilepb "nfxidentity/protos/gen/auth/forger_profile"

	"github.com/google/uuid"
)

type ForgerHandler struct {
	forgerprofilepb.UnimplementedForgerProfileServiceServer
	svc *platform.Service
}

func NewForgerHandler(svc *platform.Service) *ForgerHandler {
	return &ForgerHandler{svc: svc}
}

func (h *ForgerHandler) ConfirmForgerProfileAvatar(ctx context.Context, req *forgerprofilepb.ConfirmForgerProfileAvatarRequest) (*forgerprofilepb.ConfirmForgerProfileAvatarResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	if err := h.svc.ConfirmAvatar(ctx, aid, pid, "forger", req.GetImageId()); err != nil {
		return nil, err
	}
	return &forgerprofilepb.ConfirmForgerProfileAvatarResponse{}, nil
}

func (h *ForgerHandler) ConfirmForgerProfileBackgrounds(ctx context.Context, req *forgerprofilepb.ConfirmForgerProfileBackgroundsRequest) (*forgerprofilepb.ConfirmForgerProfileBackgroundsResponse, error) {
	aid, _ := uuid.Parse(req.GetAccountId())
	pid, _ := uuid.Parse(req.GetProfileId())
	items := make([]struct {
		ImageID   string
		SortOrder int
	}, 0, len(req.GetImages()))
	for _, img := range req.GetImages() {
		items = append(items, struct {
			ImageID   string
			SortOrder int
		}{ImageID: img.GetImageId(), SortOrder: int(img.GetSortOrder())})
	}
	if err := h.svc.ConfirmBackgrounds(ctx, aid, pid, "forger", items); err != nil {
		return nil, err
	}
	return &forgerprofilepb.ConfirmForgerProfileBackgroundsResponse{}, nil
}

func (h *ForgerHandler) BatchGetPublicProfileCards(ctx context.Context, req *forgerprofilepb.BatchGetPublicProfileCardsRequest) (*forgerprofilepb.BatchGetPublicProfileCardsResponse, error) {
	cards := h.svc.BatchPublicCards(ctx, req.GetProfileIds())
	out := make([]*forgerprofilepb.PublicProfileCard, 0, len(cards))
	for _, c := range cards {
		card := &forgerprofilepb.PublicProfileCard{ProfileId: asString(c["profile_id"])}
		if v := asStringPtr(c["display_name"]); v != nil {
			card.DisplayName = v
		}
		out = append(out, card)
	}
	return &forgerprofilepb.BatchGetPublicProfileCardsResponse{Profiles: out}, nil
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

func asStringPtr(v any) *string {
	if v == nil {
		return nil
	}
	switch t := v.(type) {
	case string:
		return &t
	case *string:
		return t
	default:
		return nil
	}
}
