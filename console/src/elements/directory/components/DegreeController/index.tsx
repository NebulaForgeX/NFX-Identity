import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const DegreeController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Input
      label={t("education.degree.label")}
      placeholder={t("education.degree.placeholder")}
      error={errors.degree?.message as string | undefined}
      fullWidth
      {...register("degree")}
    />
  );
});

DegreeController.displayName = "DegreeController";

export default DegreeController;
