import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const ResponsibilitiesController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Textarea
      label={t("occupation.responsibilities.label")}
      placeholder={t("occupation.responsibilities.placeholder")}
      error={errors.responsibilities?.message as string | undefined}
      fullWidth
      rows={4}
      {...register("responsibilities")}
    />
  );
});

ResponsibilitiesController.displayName = "ResponsibilitiesController";

export default ResponsibilitiesController;
