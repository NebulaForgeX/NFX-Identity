package mapper

import (
	"nfxidentity/enums"
	accountpb "nfxidentity/protos/gen/auth/account"
)

func ForgerRoleToProto(role enums.AuthForgerRole) accountpb.ForgerRole {
	switch role {
	case enums.AuthForgerRoleForger:
		return accountpb.ForgerRole_FORGER_ROLE_FORGER
	default:
		return accountpb.ForgerRole_FORGER_ROLE_UNSPECIFIED
	}
}

func ForgerRolesToProto(roles []enums.AuthForgerRole) []accountpb.ForgerRole {
	out := make([]accountpb.ForgerRole, 0, len(roles))
	for _, role := range roles {
		out = append(out, ForgerRoleToProto(role))
	}
	return out
}

func ForgerRoleFromProto(role accountpb.ForgerRole) enums.AuthForgerRole {
	switch role {
	case accountpb.ForgerRole_FORGER_ROLE_FORGER:
		return enums.AuthForgerRoleForger
	default:
		return ""
	}
}
