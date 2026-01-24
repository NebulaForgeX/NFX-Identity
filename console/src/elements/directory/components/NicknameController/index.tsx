import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const NicknameController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.nickname.label")}
      placeholder={t("profile.nickname.placeholder")}
      error={errors.nickname?.message as string | undefined}
      fullWidth
      {...register("nickname")}
    />
  );
});

NicknameController.displayName = "NicknameController";

export default NicknameController;
