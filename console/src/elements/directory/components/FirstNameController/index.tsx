import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const FirstNameController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.firstName.label")}
      placeholder={t("profile.firstName.placeholder")}
      error={errors.firstName?.message as string | undefined}
      fullWidth
      {...register("firstName")}
    />
  );
});

FirstNameController.displayName = "FirstNameController";

export default FirstNameController;
