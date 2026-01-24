import { memo, useEffect } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider } from "react-hook-form";
import { useNavigate, useSearchParams } from "react-router-dom";

import { Button, IconButton, Suspense } from "@/components";
import { ArrowLeft as ArrowLeftIcon } from "@/assets/icons/lucide";
import { useInitPreferenceForm, useSubmitPreference } from "@/elements/directory";
import { PreferenceKeyController, PreferenceValueController } from "@/elements/directory";
import { useUserPreference } from "@/hooks/useDirectory";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

const EditPreferencePage = memo(() => {
  const { t } = useTranslation("EditPreferencePage");
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const preferenceId = searchParams.get("id");

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
        <h1 className={styles.title}>
          {preferenceId ? t("title") : t("addTitle")}
        </h1>
        <p className={styles.subtitle}>
          {preferenceId ? t("subtitle") : t("addSubtitle")}
        </p>
      </div>

      {preferenceId ? (
      <Suspense
        loadingType="ecg"
        loadingText={t("loading")}
        loadingSize="small"
        loadingContainerClassName={styles.loading}
      >
          <EditPreferenceContent preferenceId={preferenceId} />
        </Suspense>
      ) : (
        <AddPreferenceContent />
      )}
    </div>
  );
});

EditPreferencePage.displayName = "EditPreferencePage";

const EditPreferenceContent = memo(({ preferenceId }: { preferenceId: string }) => {
  const { t } = useTranslation("EditPreferencePage");
  const navigate = useNavigate();
  const { data: preference } = useUserPreference({ id: preferenceId });
  const form = useInitPreferenceForm(
    preference
      ? {
          key: preference.key,
          value: preference.value,
        }
      : undefined,
  );
  const { onSubmit, onSubmitError, isPending } = useSubmitPreference(preference);

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit, onSubmitError)} className={styles.form}>
        <PreferenceKeyController disabled={true} />
        <PreferenceValueController />

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

const AddPreferenceContent = memo(() => {
  const { t } = useTranslation("EditPreferencePage");
  const navigate = useNavigate();
  const form = useInitPreferenceForm();
  const { onSubmit, onSubmitError, isPending } = useSubmitPreference();

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit, onSubmitError)} className={styles.form}>
        <PreferenceKeyController disabled={false} />
        <PreferenceValueController />

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

AddPreferenceContent.displayName = "AddPreferenceContent";

export default EditPreferencePage;
