import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const EducationDescriptionController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Textarea
      label={t("education.description.label")}
      placeholder={t("education.description.placeholder")}
      error={errors.description?.message as string | undefined}
      fullWidth
      rows={4}
      {...register("description")}
    />
  );
});

EducationDescriptionController.displayName = "EducationDescriptionController";

export default EducationDescriptionController;
