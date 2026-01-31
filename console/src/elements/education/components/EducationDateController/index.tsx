import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

import styles from "./styles.module.css";

const EducationDateController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <div className={styles.dateGrid}>
      <Input
        type="date"
        label={t("education.startDate.label")}
        error={errors.startDate?.message as string | undefined}
        fullWidth
        {...register("startDate")}
      />
      <Input
        type="date"
        label={t("education.endDate.label")}
        error={errors.endDate?.message as string | undefined}
        fullWidth
        {...register("endDate")}
      />
    </div>
  );
});

EducationDateController.displayName = "EducationDateController";

export default EducationDateController;
