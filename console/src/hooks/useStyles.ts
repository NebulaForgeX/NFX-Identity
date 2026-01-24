import { useTranslation } from "react-i18next";

import {
  UserStatus,
  CredentialStatus,
  CredentialType,
  MFAType,
  TenantStatus,
  MemberStatus,
  AppStatus,
  ApiKeyStatus,
  VerificationStatus,
  InvitationStatus,
  ClientCredentialStatus,
  AllowlistStatus,
} from "@/types/domain/enums";

/**
 * 状态样式信息
 */
export interface StatusStyle {
  label: string;
  color: string;
  bgColor: string;
}

/**
 * 布尔状态样式信息
 */
export interface BooleanStatusStyle {
  label: string;
  color: string;
  bgColor: string;
}

// ========== Directory Schema ==========

/**
 * 获取用户状态的颜色和文本信息
 */
export const useUserStatus = (status: UserStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<UserStatus, StatusStyle> = {
    [UserStatus.ACTIVE]: {
      label: t("userStatus.active"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [UserStatus.PENDING]: {
      label: t("userStatus.pending"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [UserStatus.DEACTIVE]: {
      label: t("userStatus.deactive"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

// ========== Auth Schema ==========

/**
 * 获取凭证状态的颜色和文本信息
 */
export const useCredentialStatus = (status: CredentialStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<CredentialStatus, StatusStyle> = {
    [CredentialStatus.ACTIVE]: {
      label: t("credentialStatus.active"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [CredentialStatus.DISABLED]: {
      label: t("credentialStatus.disabled"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
    [CredentialStatus.EXPIRED]: {
      label: t("credentialStatus.expired"),
      color: "white",
      bgColor: "var(--color-danger)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

/**
 * 获取凭证类型的文本信息
 */
export const useCredentialType = (type: CredentialType | undefined | null): string => {
  const { t } = useTranslation("hooks.styles");

  const typeMap: Record<CredentialType, string> = {
    [CredentialType.PASSWORD]: t("credentialType.password"),
    [CredentialType.PASSKEY]: t("credentialType.passkey"),
    [CredentialType.OAUTH_LINK]: t("credentialType.oauth_link"),
    [CredentialType.SAML]: t("credentialType.saml"),
    [CredentialType.LDAP]: t("credentialType.ldap"),
  };

  if (!type) {
    return "";
  }

  return typeMap[type] || type;
};

/**
 * 获取 MFA 类型的文本信息
 */
export const useMFAType = (type: MFAType | undefined | null): string => {
  const { t } = useTranslation("hooks.styles");

  const typeMap: Record<MFAType, string> = {
    [MFAType.TOTP]: t("mfaType.totp"),
    [MFAType.SMS]: t("mfaType.sms"),
    [MFAType.EMAIL]: t("mfaType.email"),
    [MFAType.WEBAUTHN]: t("mfaType.webauthn"),
    [MFAType.BACKUP_CODE]: t("mfaType.backup_code"),
  };

  if (!type) {
    return "";
  }

  return typeMap[type] || type;
};

// ========== Tenants Schema ==========

/**
 * 获取租户状态的颜色和文本信息
 */
export const useTenantStatus = (status: TenantStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<TenantStatus, StatusStyle> = {
    [TenantStatus.ACTIVE]: {
      label: t("tenantStatus.ACTIVE"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [TenantStatus.SUSPENDED]: {
      label: t("tenantStatus.SUSPENDED"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [TenantStatus.CLOSED]: {
      label: t("tenantStatus.CLOSED"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
    [TenantStatus.PENDING]: {
      label: t("tenantStatus.PENDING"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

/**
 * 获取成员状态的颜色和文本信息
 */
export const useMemberStatus = (status: MemberStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<MemberStatus, StatusStyle> = {
    [MemberStatus.ACTIVE]: {
      label: t("memberStatus.ACTIVE"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [MemberStatus.INVITED]: {
      label: t("memberStatus.INVITED"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [MemberStatus.SUSPENDED]: {
      label: t("memberStatus.SUSPENDED"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [MemberStatus.REMOVED]: {
      label: t("memberStatus.REMOVED"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

/**
 * 获取验证状态的颜色和文本信息
 */
export const useVerificationStatus = (status: VerificationStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<VerificationStatus, StatusStyle> = {
    [VerificationStatus.VERIFIED]: {
      label: t("verificationStatus.VERIFIED"),
      color: "white",
      bgColor: "var(--color-success)",
    },
    [VerificationStatus.PENDING]: {
      label: t("verificationStatus.PENDING"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [VerificationStatus.FAILED]: {
      label: t("verificationStatus.FAILED"),
      color: "white",
      bgColor: "var(--color-danger)",
    },
    [VerificationStatus.EXPIRED]: {
      label: t("verificationStatus.EXPIRED"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

/**
 * 获取邀请状态的颜色和文本信息
 */
export const useInvitationStatus = (status: InvitationStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<InvitationStatus, StatusStyle> = {
    [InvitationStatus.ACCEPTED]: {
      label: t("invitationStatus.ACCEPTED"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [InvitationStatus.PENDING]: {
      label: t("invitationStatus.PENDING"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [InvitationStatus.EXPIRED]: {
      label: t("invitationStatus.EXPIRED"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
    [InvitationStatus.REVOKED]: {
      label: t("invitationStatus.REVOKED"),
      color: "white",
      bgColor: "var(--color-danger)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

// ========== Clients Schema ==========

/**
 * 获取应用状态的颜色和文本信息
 */
export const useAppStatus = (status: AppStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<AppStatus, StatusStyle> = {
    [AppStatus.ACTIVE]: {
      label: t("appStatus.active"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [AppStatus.DISABLED]: {
      label: t("appStatus.disabled"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
    [AppStatus.SUSPENDED]: {
      label: t("appStatus.suspended"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [AppStatus.PENDING]: {
      label: t("appStatus.pending"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

/**
 * 获取 API 密钥状态的颜色和文本信息
 */
export const useApiKeyStatus = (status: ApiKeyStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<ApiKeyStatus, StatusStyle> = {
    [ApiKeyStatus.ACTIVE]: {
      label: t("apiKeyStatus.active"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [ApiKeyStatus.REVOKED]: {
      label: t("apiKeyStatus.revoked"),
      color: "white",
      bgColor: "var(--color-danger)",
    },
    [ApiKeyStatus.EXPIRED]: {
      label: t("apiKeyStatus.expired"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

/**
 * 获取客户端凭证状态的颜色和文本信息
 */
export const useClientCredentialStatus = (status: ClientCredentialStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<ClientCredentialStatus, StatusStyle> = {
    [ClientCredentialStatus.ACTIVE]: {
      label: t("credentialStatus.active"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [ClientCredentialStatus.EXPIRED]: {
      label: t("credentialStatus.expired"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
    [ClientCredentialStatus.REVOKED]: {
      label: t("apiKeyStatus.revoked"),
      color: "white",
      bgColor: "var(--color-danger)",
    },
    [ClientCredentialStatus.ROTATING]: {
      label: t("credentialStatus.active"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

/**
 * 获取白名单状态的颜色和文本信息
 */
export const useAllowlistStatus = (status: AllowlistStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  const statusMap: Record<AllowlistStatus, StatusStyle> = {
    [AllowlistStatus.ACTIVE]: {
      label: t("appStatus.active"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [AllowlistStatus.DISABLED]: {
      label: t("appStatus.disabled"),
      color: "white",
      bgColor: "var(--color-fg-muted)",
    },
    [AllowlistStatus.REVOKED]: {
      label: t("apiKeyStatus.revoked"),
      color: "white",
      bgColor: "var(--color-danger)",
    },
  };

  if (!status) {
    return {
      label: "",
      color: "white",
      bgColor: "var(--color-fg-muted)",
    };
  }

  return statusMap[status] || {
    label: status,
    color: "white",
    bgColor: "var(--color-fg-muted)",
  };
};

// ========== Boolean Status Hooks ==========

/**
 * 获取 Primary 状态的颜色和文本信息
 * @param isPrimary 是否为 Primary
 * @returns 如果为 true 返回样式，否则返回 null
 */
export const useIsPrimary = (isPrimary: boolean): BooleanStatusStyle | null => {
  const { t } = useTranslation("hooks.styles");

  if (!isPrimary) {
    return null;
  }

  return {
    label: t("boolean.primary"),
    color: "var(--color-primary-fg)",
    bgColor: "var(--color-primary)",
  };
};

/**
 * 获取 Current 状态的颜色和文本信息
 * @param isCurrent 是否为 Current
 * @returns 如果为 true 返回样式，否则返回 null
 */
export const useIsCurrent = (isCurrent: boolean): BooleanStatusStyle | null => {
  const { t } = useTranslation("hooks.styles");

  if (!isCurrent) {
    return null;
  }

  return {
    label: t("boolean.current"),
    color: "var(--color-primary-fg)",
    bgColor: "var(--color-primary)",
  };
};

/**
 * 获取 Verified 状态的颜色和文本信息
 * @param isVerified 是否为 Verified
 * @returns 如果为 true 返回样式，否则返回 null
 */
export const useIsVerified = (isVerified: boolean): BooleanStatusStyle | null => {
  const { t } = useTranslation("hooks.styles");

  if (!isVerified) {
    return null;
  }

  return {
    label: t("boolean.verified"),
    color: "white",
    bgColor: "var(--color-success)",
  };
};

/**
 * 获取验证状态的颜色和文本信息（包含未验证状态）
 * @param isVerified 是否已验证
 * @returns 状态样式信息（标签、颜色、背景色）
 */
export const useVerification = (isVerified: boolean): StatusStyle => {
  const { t } = useTranslation("hooks.styles");

  return {
    label: isVerified ? t("boolean.verified") : t("boolean.unverified"),
    color: isVerified ? "white" : "white",
    bgColor: isVerified ? "var(--color-success)" : "var(--color-warning)",
  };
};

// ========== Legacy aliases for backward compatibility ==========

/**
 * @deprecated 使用 useUserStatus 代替
 */
export const useStatus = useUserStatus;
