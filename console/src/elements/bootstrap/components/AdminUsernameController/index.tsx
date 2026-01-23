import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import { Input } from "@/components";

const AdminUsernameController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <Input
      label={t("admin_username.label")}
      placeholder={t("admin_username.placeholder")}
      error={errors.AdminUsername?.message as string | undefined}
      fullWidth
      required
      autoComplete="username"
      {...register("AdminUsername")}
    />
  );
});

AdminUsernameController.displayName = "AdminUsernameController";

export default AdminUsernameController;
