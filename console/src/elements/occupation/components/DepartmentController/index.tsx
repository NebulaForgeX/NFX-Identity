import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const DepartmentController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Input
      label={t("occupation.department.label")}
      placeholder={t("occupation.department.placeholder")}
      error={errors.department?.message as string | undefined}
      fullWidth
      {...register("department")}
    />
  );
});

DepartmentController.displayName = "DepartmentController";

export default DepartmentController;
