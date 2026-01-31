import { memo, useCallback } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider } from "react-hook-form";

import { IconButton, Suspense } from "@/components";
import { ArrowLeft as ArrowLeftIcon } from "@/assets/icons/lucide";
import {
  useInitPreferenceForm,
  ThemeController,
  LanguageController,
  TimezoneController,
  DashboardBackgroundController,
} from "@/elements/preference";
import type { PreferenceFormValues } from "@/elements/preference";
import { useUserPreference, useCreateUserPreference, useUpdateUserPreference } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import type { DashboardBackgroundType } from "@/types";
import { DEFAULT_DASHBOARD_BACKGROUND } from "@/types";
import { ChangeLanguage } from "@/assets/languages/i18n";
import { useTheme } from "@/providers/ThemeProvider/useTheme";

import styles from "./styles.module.css";

const EditPreferencePage = memo(() => {
  const { t } = useTranslation("EditPreferencePage");
  const currentUserId = useAuthStore((state) => state.currentUserId);

  if (!currentUserId) {
    return null;
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <IconButton
          variant="ghost"
          leftIcon={<ArrowLeftIcon size={20} />}
          onClick={() => window.history.back()}
          className={styles.backButton}
        >
          {t("back")}
        </IconButton>
        <h1 className={styles.title}>{t("title")}</h1>
        <p className={styles.subtitle}>{t("subtitle")}</p>
      </div>

      <Suspense
        loadingType="ecg"
        loadingText={t("loading")}
        loadingSize="small"
        loadingContainerClassName={styles.loading}
      >
        <EditPreferenceContent userId={currentUserId} />
      </Suspense>
    </div>
  );
});

EditPreferencePage.displayName = "EditPreferencePage";

const EditPreferenceContent = memo(({ userId }: { userId: string }) => {
  const { data: preference } = useUserPreference({ id: userId });
  const { setTheme } = useTheme();
  const updatePreference = useUpdateUserPreference({ silent: true });
  const createPreference = useCreateUserPreference();

  const form = useInitPreferenceForm(
    preference
      ? {
          theme: preference.theme,
          language: preference.language,
          timezone: preference.timezone,
          dashboardBackground: ((preference.other as Record<string, unknown>)?.dashboardBackground as DashboardBackgroundType | undefined) || DEFAULT_DASHBOARD_BACKGROUND,
          notifications: preference.notifications,
          privacy: preference.privacy,
          display: preference.display,
          other: preference.other,
        }
      : undefined,
  );

  /** 选择即改：保存到后端并应用（主题/语言立即生效） */
  const handleApply = useCallback(
    async (payload: Partial<PreferenceFormValues>) => {
      try {
        if (preference) {
          const data: Record<string, unknown> = { ...payload };
          if (payload.dashboardBackground !== undefined) {
            const existingOther = (preference.other as Record<string, unknown>) || {};
            data.other = { ...existingOther, dashboardBackground: payload.dashboardBackground };
            delete data.dashboardBackground;
          }
          await updatePreference.mutateAsync({
            id: preference.id,
            data: data as Parameters<typeof updatePreference.mutateAsync>[0]["data"],
          });
        } else {
          const values = form.getValues();
          const existingOther = (values.other as Record<string, unknown>) || {};
          const otherData = { ...existingOther, dashboardBackground: (payload.dashboardBackground ?? values.dashboardBackground) || DEFAULT_DASHBOARD_BACKGROUND };
          await createPreference.mutateAsync({
            userId,
            theme: payload.theme ?? values.theme,
            language: payload.language ?? values.language,
            timezone: payload.timezone ?? values.timezone,
            notifications: values.notifications,
            privacy: values.privacy,
            display: values.display,
            other: otherData,
          });
        }
        if (payload.theme) setTheme(payload.theme as Parameters<typeof setTheme>[0]);
        if (payload.language) ChangeLanguage(payload.language as Parameters<typeof ChangeLanguage>[0]);
      } catch (e) {
        console.error("Failed to save preference:", e);
      }
    },
    [preference, form, userId, updatePreference, createPreference, setTheme],
  );

  return (
    <FormProvider {...form}>
      <div className={styles.form}>
        <ThemeController onApply={(p) => handleApply(p)} />
        <LanguageController onApply={(p) => handleApply(p)} />
        <TimezoneController onApply={(p) => handleApply(p)} />
        <DashboardBackgroundController onApply={(p) => handleApply(p)} />
      </div>
    </FormProvider>
  );
});

EditPreferenceContent.displayName = "EditPreferenceContent";

export default EditPreferencePage;
