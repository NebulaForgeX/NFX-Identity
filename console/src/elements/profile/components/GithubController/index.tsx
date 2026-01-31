import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const GithubController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.github.label")}
      placeholder={t("profile.github.placeholder")}
      error={errors.github?.message as string | undefined}
      fullWidth
      {...register("github")}
    />
  );
});

GithubController.displayName = "GithubController";

export default GithubController;
