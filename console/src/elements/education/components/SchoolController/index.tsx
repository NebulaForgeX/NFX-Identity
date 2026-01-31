import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const SchoolController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Input
      label={t("education.school.label")}
      placeholder={t("education.school.placeholder")}
      error={errors.school?.message as string | undefined}
      fullWidth
      required
      {...register("school")}
    />
  );
});

SchoolController.displayName = "SchoolController";

export default SchoolController;
