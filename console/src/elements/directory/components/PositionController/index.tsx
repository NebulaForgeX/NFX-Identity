import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const PositionController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Input
      label={t("occupation.position.label")}
      placeholder={t("occupation.position.placeholder")}
      error={errors.position?.message as string | undefined}
      fullWidth
      {...register("position")}
    />
  );
});

PositionController.displayName = "PositionController";

export default PositionController;
