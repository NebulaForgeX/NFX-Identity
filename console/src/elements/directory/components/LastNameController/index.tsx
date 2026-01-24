import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const LastNameController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.lastName.label")}
      placeholder={t("profile.lastName.placeholder")}
      error={errors.lastName?.message as string | undefined}
      fullWidth
      {...register("lastName")}
    />
  );
});

LastNameController.displayName = "LastNameController";

export default LastNameController;
