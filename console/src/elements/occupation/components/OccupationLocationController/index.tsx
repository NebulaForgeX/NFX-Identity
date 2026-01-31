import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const OccupationLocationController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Input
      label={t("occupation.location.label")}
      placeholder={t("occupation.location.placeholder")}
      error={errors.location?.message as string | undefined}
      fullWidth
      {...register("location")}
    />
  );
});

OccupationLocationController.displayName = "OccupationLocationController";

export default OccupationLocationController;
