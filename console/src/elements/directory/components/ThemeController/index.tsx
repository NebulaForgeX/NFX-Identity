import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const ThemeController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  return (
    <Input
      label={t("preference.theme.label")}
      placeholder={t("preference.theme.placeholder")}
      error={errors.theme?.message as string | undefined}
      fullWidth
      {...register("theme")}
    />
  );
});

ThemeController.displayName = "ThemeController";

export default ThemeController;
