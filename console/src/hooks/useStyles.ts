import { useTranslation } from "react-i18next";

import { UserStatus } from "@/types/domain/enums";

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
  color: string;
  bgColor: string;
}

/**
 * 获取用户状态的颜色和文本信息
 * @param status 用户状态 (使用 UserStatus enum)
 * @returns 状态样式信息（标签、颜色、背景色）
 */
export const useStatus = (status: UserStatus | undefined | null): StatusStyle => {
  const { t } = useTranslation("ProfilePage");

  const statusMap: Record<UserStatus, StatusStyle> = {
    [UserStatus.ACTIVE]: {
      label: t("statusActive"),
      color: "var(--color-primary-fg)",
      bgColor: "var(--color-success)",
    },
    [UserStatus.PENDING]: {
      label: t("statusPending"),
      color: "white",
      bgColor: "var(--color-warning)",
    },
    [UserStatus.DEACTIVE]: {
      label: t("statusDeactive"),
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
 * 获取 Primary 状态的颜色信息
 * @param _isPrimary 是否为 Primary（未使用，保留用于类型推断）
 * @returns 颜色和背景色
 */
export const useIsPrimary = (_isPrimary: boolean): BooleanStatusStyle => {
  return {
    color: "var(--color-primary-fg)",
    bgColor: "var(--color-primary)",
  };
};

/**
 * 获取 Current 状态的颜色信息
 * @param _isCurrent 是否为 Current（未使用，保留用于类型推断）
 * @returns 颜色和背景色
 */
export const useIsCurrent = (_isCurrent: boolean): BooleanStatusStyle => {
  return {
    color: "var(--color-primary-fg)",
    bgColor: "var(--color-primary)",
  };
};

/**
 * 获取 Verified 状态的颜色信息
 * @param _isVerified 是否为 Verified（未使用，保留用于类型推断）
 * @returns 颜色和背景色
 */
export const useIsVerified = (_isVerified: boolean): BooleanStatusStyle => {
  return {
    color: "white",
    bgColor: "var(--color-success)",
  };
};

/**
 * 获取验证状态的颜色和文本信息
 * @param isVerified 是否已验证
 * @returns 状态样式信息（标签、颜色、背景色）
 */
export const useVerification = (isVerified: boolean): StatusStyle => {
  const { t } = useTranslation("ProfilePage");

  return {
    label: isVerified ? t("verified") : t("unverified"),
    color: isVerified ? "white" : "white",
    bgColor: isVerified ? "var(--color-success)" : "var(--color-warning)",
  };
};
