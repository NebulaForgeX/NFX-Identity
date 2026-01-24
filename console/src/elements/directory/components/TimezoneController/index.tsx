import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const TimezoneController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  return (
    <Input
      label={t("preference.timezone.label")}
      placeholder={t("preference.timezone.placeholder")}
      error={errors.timezone?.message as string | undefined}
      fullWidth
      {...register("timezone")}
    />
  );
});

TimezoneController.displayName = "TimezoneController";

export default TimezoneController;
