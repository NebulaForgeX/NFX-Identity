import type { EducationFormValues } from "../../schemas/educationSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Textarea } from "@/components";

const AchievementsController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<EducationFormValues>();

  return (
    <Textarea
      label={t("education.achievements.label")}
      placeholder={t("education.achievements.placeholder")}
      error={errors.achievements?.message as string | undefined}
      fullWidth
      rows={3}
      {...register("achievements")}
    />
  );
});

AchievementsController.displayName = "AchievementsController";

export default AchievementsController;
