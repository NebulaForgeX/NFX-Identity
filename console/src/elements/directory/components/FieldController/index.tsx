import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const FieldController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Input
      label={t("education.fieldOfStudy.label")}
      placeholder={t("education.fieldOfStudy.placeholder")}
      error={errors.fieldOfStudy?.message as string | undefined}
      fullWidth
      {...register("fieldOfStudy")}
    />
  );
});

FieldController.displayName = "FieldController";

export default FieldController;
