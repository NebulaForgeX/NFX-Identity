import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const MajorController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Input
      label={t("education.major.label")}
      placeholder={t("education.major.placeholder")}
      error={errors.major?.message as string | undefined}
      fullWidth
      {...register("major")}
    />
  );
});

MajorController.displayName = "MajorController";

export default MajorController;
