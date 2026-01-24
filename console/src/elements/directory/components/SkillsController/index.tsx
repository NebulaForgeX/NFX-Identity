import type { ProfileFormValues } from "../../schemas/profileSchema";
import type { KeyValuePair } from "@/components/KeyValueEditor";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { KeyValueEditor } from "@/components";

const SkillsController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    setValue,
    watch,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  const skills = watch("skills") || [];

  const handleChange = (pairs: KeyValuePair[]) => {
    setValue("skills", pairs, { shouldValidate: true });
  };

  return (
    <KeyValueEditor
      label={t("profile.skills.label")}
      pairs={skills}
      onChange={handleChange}
      valueType="array"
      keyPlaceholder={t("profile.skills.keyPlaceholder")}
      valuePlaceholder={t("profile.skills.valuePlaceholder")}
      error={errors.skills?.message as string | undefined}
    />
  );
});

SkillsController.displayName = "SkillsController";

export default SkillsController;
