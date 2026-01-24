import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

interface PreferenceKeyControllerProps {
  disabled?: boolean;
}

const PreferenceKeyController = memo(({ disabled = false }: PreferenceKeyControllerProps) => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  return (
    <Input
      label={t("preference.key.label")}
      placeholder={t("preference.key.placeholder")}
      error={errors.key?.message as string | undefined}
      fullWidth
      required
      disabled={disabled}
      {...register("key")}
    />
  );
});

PreferenceKeyController.displayName = "PreferenceKeyController";

export default PreferenceKeyController;
