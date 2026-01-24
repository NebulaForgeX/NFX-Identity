import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider } from "react-hook-form";
import { useNavigate, useSearchParams } from "react-router-dom";

import { Button, IconButton, Suspense } from "@/components";
import { ArrowLeft as ArrowLeftIcon } from "@/assets/icons/lucide";
import { useInitOccupationForm, useSubmitOccupation } from "@/elements/directory";
import {
  CompanyController,
  PositionController,
  DepartmentController,
  IndustryController,
  OccupationLocationController,
  EmploymentTypeController,
  OccupationDateController,
  IsCurrentOccupationController,
  OccupationDescriptionController,
  ResponsibilitiesController,
  OccupationAchievementsController,
} from "@/elements/directory";
import { useUserOccupation } from "@/hooks/useDirectory";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

const EditOccupationPage = memo(() => {
  const { t } = useTranslation("EditOccupationPage");
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const occupationId = searchParams.get("id");

  if (!occupationId) {
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
        <EditOccupationContent occupationId={occupationId} />
      </Suspense>
    </div>
  );
});

EditOccupationPage.displayName = "EditOccupationPage";

const EditOccupationContent = memo(({ occupationId }: { occupationId: string }) => {
  const { t } = useTranslation("EditOccupationPage");
  const navigate = useNavigate();
  const { data: occupation } = useUserOccupation({ id: occupationId });
  const form = useInitOccupationForm(occupation);
  const { onSubmit, onSubmitError, isPending } = useSubmitOccupation(occupation);

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit, onSubmitError)} className={styles.form}>
        <div className={styles.formGrid}>
          <CompanyController />
          <PositionController />
        </div>
        <div className={styles.formGrid}>
          <DepartmentController />
          <IndustryController />
        </div>
        <div className={styles.formGrid}>
          <OccupationLocationController />
          <EmploymentTypeController />
        </div>
        <OccupationDateController />
        <IsCurrentOccupationController />
        <OccupationDescriptionController />
        <ResponsibilitiesController />
        <OccupationAchievementsController />

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

EditOccupationContent.displayName = "EditOccupationContent";

export default EditOccupationPage;
