import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const LocationController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.location.label")}
      placeholder={t("profile.location.placeholder")}
      error={errors.location?.message as string | undefined}
      fullWidth
      {...register("location")}
    />
  );
});

LocationController.displayName = "LocationController";

export default LocationController;
