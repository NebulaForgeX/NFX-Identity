package handler

import (
	"context"
	"time"

	"nfxidentity/modules/auth/application/platform"
	accountpb "nfxidentity/protos/gen/auth/account"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountHandler struct {
	accountpb.UnimplementedAccountServiceServer
	svc *platform.Service
}

func NewAccountHandler(svc *platform.Service) *AccountHandler {
	return &AccountHandler{svc: svc}
}

func (h *AccountHandler) GetFullAccountInformation(ctx context.Context, req *accountpb.GetFullAccountInformationRequest) (*accountpb.GetFullAccountInformationResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	full, err := h.svc.FullAccountWithProfile(ctx, aid, uuid.Nil, "forger")
	if err != nil {
		return nil, err
	}
	return &accountpb.GetFullAccountInformationResponse{Full: mapFull(full)}, nil
}

func (h *AccountHandler) GetAccountByID(ctx context.Context, req *accountpb.GetAccountByIDRequest) (*accountpb.GetAccountByIDResponse, error) {
	aid, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	full, err := h.svc.FullAccountWithProfile(ctx, aid, uuid.Nil, "forger")
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountByIDResponse{Items: []*accountpb.FullAccountInformation{mapFull(full)}}, nil
}

func (h *AccountHandler) GetAccountByEmail(ctx context.Context, req *accountpb.GetAccountByEmailRequest) (*accountpb.GetAccountByEmailResponse, error) {
	aid, err := h.svc.AccountIDByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}
	full, err := h.svc.FullAccountWithProfile(ctx, aid, uuid.Nil, "forger")
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountByEmailResponse{Items: []*accountpb.FullAccountInformation{mapFull(full)}}, nil
}

func (h *AccountHandler) GetAccountByPhone(ctx context.Context, req *accountpb.GetAccountByPhoneRequest) (*accountpb.GetAccountByPhoneResponse, error) {
	aid, err := h.svc.AccountIDByPhone(ctx, req.GetPhone())
	if err != nil {
		return nil, err
	}
	full, err := h.svc.FullAccountWithProfile(ctx, aid, uuid.Nil, "forger")
	if err != nil {
		return nil, err
	}
	return &accountpb.GetAccountByPhoneResponse{Items: []*accountpb.FullAccountInformation{mapFull(full)}}, nil
}

func (h *AccountHandler) GetPrimaryEmailByProfileID(ctx context.Context, req *accountpb.GetPrimaryEmailByProfileIDRequest) (*accountpb.GetPrimaryEmailByProfileIDResponse, error) {
	pid, err := uuid.Parse(req.GetProfileId())
	if err != nil {
		return nil, err
	}
	card, err := h.svc.PublicProfileCard(ctx, pid)
	if err != nil {
		return nil, err
	}
	accountID, _ := uuid.Parse(asString(card["account_id"]))
	email, _ := h.svc.PrimaryContactsExport(ctx, accountID)
	lang := asString(card["profile_language"])
	return &accountpb.GetPrimaryEmailByProfileIDResponse{Email: email, ProfileLanguage: lang}, nil
}

func (h *AccountHandler) InvalidateFullInformationWithForgerProfile(ctx context.Context, req *accountpb.InvalidateFullInformationWithForgerProfileRequest) (*accountpb.InvalidateFullInformationWithForgerProfileResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
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
		return nil, err
	}
	pid, _ := uuid.Parse(req.GetProfileId())
	if err := h.svc.InvalidateFull(ctx, aid, pid); err != nil {
		return nil, err
	}
	return &accountpb.InvalidateFullInformationWithAuthorityProfileResponse{}, nil
}

func (h *AccountHandler) EnsureOwnedProfile(ctx context.Context, req *accountpb.EnsureOwnedProfileRequest) (*accountpb.EnsureOwnedProfileResponse, error) {
	err := h.svc.EnsureOwnedProfile(ctx, req.GetAccountId(), req.GetProfileId(), req.GetProfileScope())
	if err != nil {
		return &accountpb.EnsureOwnedProfileResponse{Allowed: false}, nil
	}
	return &accountpb.EnsureOwnedProfileResponse{Allowed: true}, nil
}

func (h *AccountHandler) ListProfilesInTable(ctx context.Context, req *accountpb.ListProfilesInTableRequest) (*accountpb.ListProfilesInTableResponse, error) {
	items, total, err := h.svc.SearchProfiles(ctx, req.GetTable(), req.GetQuery(), int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, err
	}
	out := make([]*accountpb.FullAccountInformation, 0, len(items))
	for _, item := range items {
		aid, err := uuid.Parse(asString(item["account_id"]))
		if err != nil {
			continue
		}
		full, err := h.svc.FullAccountWithProfile(ctx, aid, uuid.Nil, req.GetTable())
		if err != nil {
			continue
		}
		out = append(out, mapFull(full))
	}
	return &accountpb.ListProfilesInTableResponse{Items: out, Total: total}, nil
}

func asString(v any) string {
	if v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}

func mapFull(in map[string]any) *accountpb.FullAccountInformation {
	full := &accountpb.FullAccountInformation{}
	if acc, ok := in["account"].(map[string]any); ok {
		full.Account = &accountpb.Account{
			Id:             asString(acc["id"]),
			SignupPlatform: asString(acc["signup_platform"]),
			AccountStatus:  accountStatus(asString(acc["account_status"])),
			CreatedAt:      ts(acc["created_at"]),
			UpdatedAt:      ts(acc["updated_at"]),
		}
	}
	if emails, ok := in["emails"].([]map[string]any); ok {
		for _, e := range emails {
			full.Emails = append(full.Emails, &accountpb.Email{
				Id:        asString(e["id"]),
				AccountId: asString(e["account_id"]),
				Email:     asString(e["email"]),
				IsPrimary: asBool(e["is_primary"]),
			})
		}
	}
	if phones, ok := in["phones"].([]map[string]any); ok {
		for _, p := range phones {
			full.Phones = append(full.Phones, &accountpb.Phone{
				Id:        asString(p["id"]),
				AccountId: asString(p["account_id"]),
				Phone:     asString(p["phone"]),
				IsPrimary: asBool(p["is_primary"]),
			})
		}
	}
	if fp, ok := in["forger_profile"].(map[string]any); ok && fp != nil {
		full.ForgerProfile = &accountpb.ForgerProfile{
			AccountId:       asString(fp["account_id"]),
			ProfileId:       asString(fp["profile_id"]),
			DisplayName:     strPtr(fp["display_name"]),
			ProfileLanguage: profileLang(asString(fp["profile_language"])),
		}
	}
	if ap, ok := in["authority_profile"].(map[string]any); ok && ap != nil {
		full.AuthorityProfile = &accountpb.AuthorityProfile{
			AccountId:       asString(ap["account_id"]),
			ProfileId:       asString(ap["profile_id"]),
			DisplayName:     strPtr(ap["display_name"]),
			ProfileLanguage: profileLang(asString(ap["profile_language"])),
		}
	}
	return full
}

func asBool(v any) bool {
	b, _ := v.(bool)
	return b
}

func strPtr(v any) *string {
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

func ts(v any) *timestamppb.Timestamp {
	switch t := v.(type) {
	case time.Time:
		return timestamppb.New(t)
	default:
		return nil
	}
}

func accountStatus(s string) accountpb.AccountStatus {
	switch s {
	case "active":
		return accountpb.AccountStatus_ACCOUNT_STATUS_ACTIVE
	case "suspended":
		return accountpb.AccountStatus_ACCOUNT_STATUS_SUSPENDED
	case "deleted":
		return accountpb.AccountStatus_ACCOUNT_STATUS_DELETED
	default:
		return accountpb.AccountStatus_ACCOUNT_STATUS_UNSPECIFIED
	}
}

func profileLang(s string) accountpb.ProfileLanguage {
	switch s {
	case "en":
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_EN
	case "zh":
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_ZH
	case "fr":
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_FR
	default:
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_UNSPECIFIED
	}
}
