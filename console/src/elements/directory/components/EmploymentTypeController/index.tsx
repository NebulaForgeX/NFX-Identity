import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const EmploymentTypeController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Input
      label={t("occupation.employmentType.label")}
      placeholder={t("occupation.employmentType.placeholder")}
      error={errors.employmentType?.message as string | undefined}
      fullWidth
      {...register("employmentType")}
    />
  );
});

EmploymentTypeController.displayName = "EmploymentTypeController";

export default EmploymentTypeController;
