import type { ProfileFormValues } from "../../schemas/profileSchema";
import type { KeyValuePair } from "@/components/KeyValueEditor";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { KeyValueEditor } from "@/components";

const SocialLinksController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    setValue,
    watch,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  const socialLinks = watch("socialLinks") || [];

  const handleChange = (pairs: KeyValuePair[]) => {
    setValue("socialLinks", pairs, { shouldValidate: true });
  };

  return (
    <KeyValueEditor
      label={t("profile.socialLinks.label")}
      pairs={socialLinks}
      onChange={handleChange}
      valueType="string"
      keyPlaceholder={t("profile.socialLinks.keyPlaceholder")}
      valuePlaceholder={t("profile.socialLinks.valuePlaceholder")}
      error={errors.socialLinks?.message as string | undefined}
    />
  );
});

SocialLinksController.displayName = "SocialLinksController";

export default SocialLinksController;
