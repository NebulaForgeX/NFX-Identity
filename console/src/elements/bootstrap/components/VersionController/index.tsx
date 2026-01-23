import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const VersionController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <Input
      label={t("version.label")}
      placeholder={t("version.placeholder")}
      error={errors.Version?.message as string | undefined}
      fullWidth
      required
      {...register("Version")}
    />
  );
});

VersionController.displayName = "VersionController";

export default VersionController;
