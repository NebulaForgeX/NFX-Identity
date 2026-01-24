import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const DisplayNameController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.displayName.label")}
      placeholder={t("profile.displayName.placeholder")}
      error={errors.displayName?.message as string | undefined}
      fullWidth
      {...register("displayName")}
    />
  );
});

DisplayNameController.displayName = "DisplayNameController";

export default DisplayNameController;
