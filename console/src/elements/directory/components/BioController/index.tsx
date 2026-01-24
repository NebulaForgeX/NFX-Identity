import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const BioController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Textarea
      label={t("profile.bio.label")}
      placeholder={t("profile.bio.placeholder")}
      error={errors.bio?.message as string | undefined}
      fullWidth
      rows={4}
      {...register("bio")}
    />
  );
});

BioController.displayName = "BioController";

export default BioController;
