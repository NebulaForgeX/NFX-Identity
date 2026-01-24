import { useCallback } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import type { FieldErrors } from "react-hook-form";

import { useCreateUserPreference, useUpdateUserPreference } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import { ROUTES } from "@/types/navigation";
import { showError } from "@/stores/modalStore";
import type { UserPreference } from "@/types";

import { type PreferenceFormValues } from "../schemas/preferenceSchema";

export const useSubmitPreference = (preference?: UserPreference) => {
  const { t } = useTranslation("elements.directory");
  const navigate = useNavigate();
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const createPreference = useCreateUserPreference();
  const updatePreference = useUpdateUserPreference();

  const onSubmit = useCallback(
    async (values: PreferenceFormValues) => {
      if (!currentUserId) {
        showError(t("preference.messages.user_not_found"));
        return;
      }

      try {
        // 将 dashboardBackground 保存到 other 字段中，保留原有的 other 数据
        const existingOther = preference?.other as Record<string, unknown> || {};
        const otherData = {
          ...existingOther,
          ...values.other,
          dashboardBackground: values.dashboardBackground || "none",
        };

        if (preference) {
          // 更新现有偏好
          await updatePreference.mutateAsync({
            id: preference.id,
            data: {
              theme: values.theme,
              language: values.language,
              timezone: values.timezone,
              notifications: values.notifications,
              privacy: values.privacy,
              display: values.display,
              other: otherData,
            },
          });
        } else {
          // 创建新偏好
          await createPreference.mutateAsync({
            userId: currentUserId,
            theme: values.theme,
            language: values.language,
            timezone: values.timezone,
            notifications: values.notifications,
            privacy: values.privacy,
            display: values.display,
            other: otherData,
          });
        }
        navigate(ROUTES.PROFILE);
      } catch (error) {
        console.error("Failed to save preference:", error);
      }
    },
    [currentUserId, preference, createPreference, updatePreference, navigate, t],
  );

  const onSubmitError = useCallback((errors: FieldErrors<PreferenceFormValues>) => {
    console.error("Form validation errors:", errors);
  }, []);

  return {
    onSubmit,
    onSubmitError,
    isPending: createPreference.isPending || updatePreference.isPending,
  };
};
