import type { BootstrapFormValues } from "../../schemas/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const AdminEmailController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <Input
      label={t("admin_email.label")}
      type="email"
      placeholder={t("admin_email.placeholder")}
      error={errors.AdminEmail?.message as string | undefined}
      fullWidth
      autoComplete="email"
      {...register("AdminEmail")}
    />
  );
});

AdminEmailController.displayName = "AdminEmailController";

export default AdminEmailController;
