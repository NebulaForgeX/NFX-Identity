import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const LanguageController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  return (
    <Input
      label={t("preference.language.label")}
      placeholder={t("preference.language.placeholder")}
      error={errors.language?.message as string | undefined}
      fullWidth
      {...register("language")}
    />
  );
});

LanguageController.displayName = "LanguageController";

export default LanguageController;
