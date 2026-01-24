import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const PreferenceValueController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  return (
    <Textarea
      label={t("preference.value.label")}
      placeholder={t("preference.value.placeholder")}
      error={errors.value?.message as string | undefined}
      fullWidth
      required
      rows={4}
      {...register("value")}
    />
  );
});

PreferenceValueController.displayName = "PreferenceValueController";

export default PreferenceValueController;
