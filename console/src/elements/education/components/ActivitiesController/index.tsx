import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const ActivitiesController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Textarea
      label={t("education.activities.label")}
      placeholder={t("education.activities.placeholder")}
      error={errors.activities?.message as string | undefined}
      fullWidth
      rows={3}
      {...register("activities")}
    />
  );
});

ActivitiesController.displayName = "ActivitiesController";

export default ActivitiesController;
