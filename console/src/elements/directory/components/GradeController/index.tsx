import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const GradeController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Input
      label={t("education.grade.label")}
      placeholder={t("education.grade.placeholder")}
      error={errors.grade?.message as string | undefined}
      fullWidth
      {...register("grade")}
    />
  );
});

GradeController.displayName = "GradeController";

export default GradeController;
