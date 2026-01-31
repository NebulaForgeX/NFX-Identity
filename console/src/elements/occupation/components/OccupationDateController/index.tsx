import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

import styles from "./styles.module.css";

const OccupationDateController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <div className={styles.dateGrid}>
      <Input
        type="date"
        label={t("occupation.startDate.label")}
        error={errors.startDate?.message as string | undefined}
        fullWidth
        {...register("startDate")}
      />
      <Input
        type="date"
        label={t("occupation.endDate.label")}
        error={errors.endDate?.message as string | undefined}
        fullWidth
        {...register("endDate")}
      />
    </div>
  );
});

OccupationDateController.displayName = "OccupationDateController";

export default OccupationDateController;
