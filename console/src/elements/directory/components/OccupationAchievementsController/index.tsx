import type { OccupationFormValues } from "../../schemas/occupationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const OccupationAchievementsController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<OccupationFormValues>();

  return (
    <Textarea
      label={t("occupation.achievements.label")}
      placeholder={t("occupation.achievements.placeholder")}
      error={errors.achievements?.message as string | undefined}
      fullWidth
      rows={3}
      {...register("achievements")}
    />
  );
});

OccupationAchievementsController.displayName = "OccupationAchievementsController";

export default OccupationAchievementsController;
