import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const BirthdayController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.birthday.label")}
      placeholder={t("profile.birthday.placeholder")}
      error={errors.birthday?.message as string | undefined}
      fullWidth
      type="date"
      {...register("birthday")}
    />
  );
});

BirthdayController.displayName = "BirthdayController";

export default BirthdayController;
