import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const IndustryController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Input
      label={t("occupation.industry.label")}
      placeholder={t("occupation.industry.placeholder")}
      error={errors.industry?.message as string | undefined}
      fullWidth
      {...register("industry")}
    />
  );
});

IndustryController.displayName = "IndustryController";

export default IndustryController;
