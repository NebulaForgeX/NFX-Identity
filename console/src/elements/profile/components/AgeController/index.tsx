import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const AgeController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.age.label")}
      placeholder={t("profile.age.placeholder")}
      error={errors.age?.message as string | undefined}
      fullWidth
      type="number"
      min={0}
      max={150}
      {...register("age", { valueAsNumber: true })}
    />
  );
});

AgeController.displayName = "AgeController";

export default AgeController;
