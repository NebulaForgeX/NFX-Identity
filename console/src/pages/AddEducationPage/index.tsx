import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider } from "react-hook-form";
import { useNavigate } from "react-router-dom";

import { Button, IconButton } from "@/components";
import { ArrowLeft as ArrowLeftIcon } from "@/assets/icons/lucide";
import {
  useInitEducationForm,
  useSubmitEducation,
  SchoolController,
  DegreeController,
  MajorController,
  FieldController,
  EducationDateController,
  EducationDescriptionController,
  GradeController,
  ActivitiesController,
  AchievementsController,
  IsCurrentEducationController,
} from "@/elements/education";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

const AddEducationPage = memo(() => {
  const { t } = useTranslation("AddEducationPage");
  const navigate = useNavigate();
  const form = useInitEducationForm();
  const { onSubmit, onSubmitError, isPending } = useSubmitEducation();

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

      <FormProvider {...form}>
        <form onSubmit={form.handleSubmit(onSubmit, onSubmitError)} className={styles.form}>
          <SchoolController />
          <div className={styles.formGrid}>
            <DegreeController />
            <MajorController />
          </div>
          <FieldController />
          <EducationDateController />
          <IsCurrentEducationController />
          <EducationDescriptionController />
          <div className={styles.formGrid}>
            <GradeController />
            <div></div>
          </div>
          <ActivitiesController />
          <AchievementsController />

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
    </div>
  );
});

AddEducationPage.displayName = "AddEducationPage";

export default AddEducationPage;
