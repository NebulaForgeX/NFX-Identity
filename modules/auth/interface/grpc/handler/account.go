package handler

import (
	"context"

	"nfxidentity/enums"
	"nfxidentity/modules/auth/application/account"
	"nfxidentity/modules/auth/interface/grpc/mapper"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/query"
	accountpb "nfxidentity/protos/gen/auth/account"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountHandler struct {
	accountpb.UnimplementedAccountServiceServer
	svc *account.Service
}

func NewAccountHandler(svc *account.Service) *AccountHandler {
	return &AccountHandler{svc: svc}
}

func (h *AccountHandler) GetFullAccountInformation(ctx context.Context, req *accountpb.GetFullAccountInformationRequest) (*accountpb.GetFullAccountInformationResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account_id: %v", err)
	}
	full, err := h.svc.GetFullInformationWithCommunityProfile(ctx, account.GetFullInformationWithCommunityProfileInput{AccountID: aid, ProfileID: uuid.Nil})
	if err != nil {
		return nil, err
	}
	return &accountpb.GetFullAccountInformationResponse{Full: mapper.CommunityFullToProto(full)}, nil
}

func (h *AccountHandler) GetAccountByID(ctx context.Context, req *accountpb.GetAccountByIDRequest) (*accountpb.GetAccountByIDResponse, error) {
	aid, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}
	full, err := h.svc.GetFullInformationWithCommunityProfile(ctx, account.GetFullInformationWithCommunityProfileInput{AccountID: aid, ProfileID: uuid.Nil})
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountByIDResponse{Items: []*accountpb.FullAccountInformation{mapper.CommunityFullToProto(full)}}, nil
}

func (h *AccountHandler) GetAccountByEmail(ctx context.Context, req *accountpb.GetAccountByEmailRequest) (*accountpb.GetAccountByEmailResponse, error) {
	aid, err := h.svc.AccountIDByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}
	full, err := h.svc.GetFullInformationWithCommunityProfile(ctx, account.GetFullInformationWithCommunityProfileInput{AccountID: aid, ProfileID: uuid.Nil})
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountByEmailResponse{Items: []*accountpb.FullAccountInformation{mapper.CommunityFullToProto(full)}}, nil
}

func (h *AccountHandler) GetAccountByPhone(ctx context.Context, req *accountpb.GetAccountByPhoneRequest) (*accountpb.GetAccountByPhoneResponse, error) {
	aid, err := h.svc.AccountIDByPhone(ctx, req.GetPhone())
	if err != nil {
		return nil, err
	}
	full, err := h.svc.GetFullInformationWithCommunityProfile(ctx, account.GetFullInformationWithCommunityProfileInput{AccountID: aid, ProfileID: uuid.Nil})
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountByPhoneResponse{Items: []*accountpb.FullAccountInformation{mapper.CommunityFullToProto(full)}}, nil
}

func (h *AccountHandler) GetPrimaryEmailByProfileID(ctx context.Context, req *accountpb.GetPrimaryEmailByProfileIDRequest) (*accountpb.GetPrimaryEmailByProfileIDResponse, error) {
	pid, err := uuid.Parse(req.GetProfileId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}
	vo, err := h.svc.GetPrimaryEmailByProfileID(ctx, pid)
	if err != nil {
		return nil, err
	}
	return &accountpb.GetPrimaryEmailByProfileIDResponse{Email: vo.Email, ProfileLanguage: vo.ProfileLanguage}, nil
}

func (h *AccountHandler) InvalidateFullInformationWithForgerProfile(ctx context.Context, req *accountpb.InvalidateFullInformationWithForgerProfileRequest) (*accountpb.InvalidateFullInformationWithForgerProfileResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account_id: %v", err)
	}
	pid, _ := uuid.Parse(req.GetProfileId())
	if err := h.svc.InvalidateFull(ctx, aid, pid); err != nil {
		return nil, err
	}
	return &accountpb.InvalidateFullInformationWithForgerProfileResponse{}, nil
}

func (h *AccountHandler) InvalidateFullInformationWithAuthorityProfile(ctx context.Context, req *accountpb.InvalidateFullInformationWithAuthorityProfileRequest) (*accountpb.InvalidateFullInformationWithAuthorityProfileResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account_id: %v", err)
	}
	pid, _ := uuid.Parse(req.GetProfileId())
	if err := h.svc.InvalidateFull(ctx, aid, pid); err != nil {
		return nil, err
	}
	return &accountpb.InvalidateFullInformationWithAuthorityProfileResponse{}, nil
}

func (h *AccountHandler) EnsureOwnedProfile(ctx context.Context, req *accountpb.EnsureOwnedProfileRequest) (*accountpb.EnsureOwnedProfileResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid account_id: %v", err)
	}
	pid, err := uuid.Parse(req.GetProfileId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid profile_id: %v", err)
	}
	allowed, err := h.svc.EnsureOwnedProfile(ctx, aid, pid, enums.AuthProfileScope(req.GetProfileScope()))
	if err != nil {
		return nil, err
	}
	return &accountpb.EnsureOwnedProfileResponse{Allowed: allowed}, nil
}

func (h *AccountHandler) BootstrapOwner(ctx context.Context, req *accountpb.BootstrapOwnerRequest) (*accountpb.BootstrapOwnerResponse, error) {
	out, err := h.svc.BootstrapOwner(ctx, account.BootstrapOwnerInput{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
		Phone:    req.GetPhone(),
	})
	if err != nil {
		return nil, err
	}
	return &accountpb.BootstrapOwnerResponse{
		AccountId:          out.AccountID,
		ForgerProfileId:    out.ForgerProfileID,
		AuthorityProfileId: out.AuthorityProfileID,
	}, nil
}

func (h *AccountHandler) ListProfilesInTable(ctx context.Context, req *accountpb.ListProfilesInTableRequest) (*accountpb.ListProfilesInTableResponse, error) {
	q := profileQuery.ListQuery{DomainPagination: query.DomainPagination{Limit: int(req.GetLimit()), Offset: int(req.GetOffset())}}
	if raw := req.GetQuery(); raw != "" {
		q.Search = &raw
	}
	q.Normalize()

	if enums.AuthProfileScope(req.GetTable()) == enums.AuthProfileScopeAuthority {
		page, err := h.svc.SearchAuthorityProfiles(ctx, account.SearchAuthorityProfilesInput{Query: q})
		if err != nil {
			return nil, err
		}
		out := make([]*accountpb.FullAccountInformation, 0, len(page.Items))
		for _, item := range page.Items {
			full, err := h.svc.GetFullInformationWithAuthorityProfile(ctx, account.GetFullInformationWithAuthorityProfileInput{AccountID: item.AccountID, ProfileID: item.ProfileID})
			if err != nil {
				continue
			}
			out = append(out, mapper.AuthorityFullToProto(full))
		}
		return &accountpb.ListProfilesInTableResponse{Items: out, Total: page.Total}, nil
	}

	page, err := h.svc.SearchCommunityProfiles(ctx, account.SearchCommunityProfilesInput{Query: q})
	if err != nil {
		return nil, err
	}
	out := make([]*accountpb.FullAccountInformation, 0, len(page.Items))
	for _, item := range page.Items {
		full, err := h.svc.GetFullInformationWithCommunityProfile(ctx, account.GetFullInformationWithCommunityProfileInput{AccountID: item.AccountID, ProfileID: item.ProfileID})
		if err != nil {
			continue
		}
		out = append(out, mapper.CommunityFullToProto(full))
	}
	return &accountpb.ListProfilesInTableResponse{Items: out, Total: page.Total}, nil
}
