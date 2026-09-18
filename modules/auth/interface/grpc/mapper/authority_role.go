package mapper

import (
	"nfxidentity/enums"
	accountpb "nfxidentity/protos/gen/auth/account"
)

func AuthorityRoleToProto(role enums.AuthAuthorityRole) accountpb.AuthorityRole {
	switch role {
	case enums.AuthAuthorityRoleAuditor:
		return accountpb.AuthorityRole_AUTHORITY_ROLE_AUDITOR
	case enums.AuthAuthorityRoleAdministrator:
		return accountpb.AuthorityRole_AUTHORITY_ROLE_ADMINISTRATOR
	case enums.AuthAuthorityRoleOwner:
		return accountpb.AuthorityRole_AUTHORITY_ROLE_OWNER
	default:
		return accountpb.AuthorityRole_AUTHORITY_ROLE_UNSPECIFIED
	}
}

func AuthorityRolesToProto(roles []enums.AuthAuthorityRole) []accountpb.AuthorityRole {
	out := make([]accountpb.AuthorityRole, 0, len(roles))
	for _, role := range roles {
		out = append(out, AuthorityRoleToProto(role))
	}
	return out
}

func AuthorityRoleFromProto(role accountpb.AuthorityRole) enums.AuthAuthorityRole {
	switch role {
	case accountpb.AuthorityRole_AUTHORITY_ROLE_AUDITOR:
		return enums.AuthAuthorityRoleAuditor
	case accountpb.AuthorityRole_AUTHORITY_ROLE_ADMINISTRATOR:
		return enums.AuthAuthorityRoleAdministrator
	case accountpb.AuthorityRole_AUTHORITY_ROLE_OWNER:
		return enums.AuthAuthorityRoleOwner
	default:
		return ""
	}
}
