import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider } from "react-hook-form";
import { useNavigate } from "react-router-dom";

import { Button, IconButton, Suspense } from "@/components";
import { ArrowLeft as ArrowLeftIcon } from "@/assets/icons/lucide";
import {
  useInitPreferenceForm,
  useSubmitPreference,
  ThemeController,
  LanguageController,
  TimezoneController,
  DashboardBackgroundController,
} from "@/elements/preference";
import { useUserPreference } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import type { DashboardBackgroundType } from "@/types";
import { DEFAULT_DASHBOARD_BACKGROUND } from "@/types";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

const EditPreferencePage = memo(() => {
  const { t } = useTranslation("EditPreferencePage");
  const navigate = useNavigate();
  const currentUserId = useAuthStore((state) => state.currentUserId);

  if (!currentUserId) {
    navigate(ROUTES.PROFILE);
    return null;
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <IconButton
          variant="ghost"
          leftIcon={<ArrowLeftIcon size={20} />}
          onClick={() => navigate(ROUTES.PROFILE)}
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
  const { t } = useTranslation("EditPreferencePage");
  const navigate = useNavigate();
  const { data: preference } = useUserPreference({ id: userId });
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
  const { onSubmit, onSubmitError, isPending } = useSubmitPreference(preference);

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit, onSubmitError)} className={styles.form}>
        <ThemeController />
        <LanguageController />
        <TimezoneController />
        <DashboardBackgroundController />

        <div className={styles.actions}>
          <Button
            type="button"
            variant="secondary"
            onClick={() => navigate(ROUTES.PROFILE)}
            disabled={isPending}
          >
            {t("cancel")}
          </Button>
          <Button type="submit" variant="primary" disabled={isPending}>
            {isPending ? t("submitting") : t("submit")}
          </Button>
        </div>
      </form>
    </FormProvider>
  );
});

EditPreferenceContent.displayName = "EditPreferenceContent";

export default EditPreferencePage;
