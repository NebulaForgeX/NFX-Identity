import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const OccupationDescriptionController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Textarea
      label={t("occupation.description.label")}
      placeholder={t("occupation.description.placeholder")}
      error={errors.description?.message as string | undefined}
      fullWidth
      rows={4}
      {...register("description")}
    />
  );
});

OccupationDescriptionController.displayName = "OccupationDescriptionController";

export default OccupationDescriptionController;
