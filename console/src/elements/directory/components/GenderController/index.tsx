import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const GenderController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.gender.label")}
      placeholder={t("profile.gender.placeholder")}
      error={errors.gender?.message as string | undefined}
      fullWidth
      {...register("gender")}
    />
  );
});

GenderController.displayName = "GenderController";

export default GenderController;
