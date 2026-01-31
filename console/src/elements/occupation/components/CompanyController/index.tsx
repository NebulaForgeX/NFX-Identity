import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const CompanyController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Input
      label={t("occupation.company.label")}
      placeholder={t("occupation.company.placeholder")}
      error={errors.company?.message as string | undefined}
      fullWidth
      required
      {...register("company")}
    />
  );
});

CompanyController.displayName = "CompanyController";

export default CompanyController;
