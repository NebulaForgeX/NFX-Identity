import type { ProfileFormValues } from "../../schemas/profileSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const WebsiteController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    register,
    formState: { errors },
  } = useFormContext<ProfileFormValues>();

  return (
    <Input
      label={t("profile.website.label")}
      placeholder={t("profile.website.placeholder")}
      error={errors.website?.message as string | undefined}
      fullWidth
      type="url"
      {...register("website")}
    />
  );
});

WebsiteController.displayName = "WebsiteController";

export default WebsiteController;
