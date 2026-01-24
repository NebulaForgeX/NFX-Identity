import type { BootstrapFormValues } from "../../schemas/bootstrapSchema";

import { memo, useState } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";
import { Eye, EyeOff } from "@/assets/icons/lucide";

import { Button, Input } from "@/components";

const AdminPasswordController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();
  const [showPassword, setShowPassword] = useState(false);

  return (
    <Input
      label={t("admin_password.label")}
      type={showPassword ? "text" : "password"}
      placeholder={t("admin_password.placeholder")}
      error={errors.AdminPassword?.message as string | undefined}
      fullWidth
      required
      autoComplete="new-password"
      rightIconInteractive
      rightIcon={
        <Button
          type="button"
          variant="ghost"
          size="small"
          aria-label={showPassword ? t("admin_password.hide_password") : t("admin_password.show_password")}
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
        </Button>
      }
      {...register("AdminPassword")}
    />
  );
});

AdminPasswordController.displayName = "AdminPasswordController";

export default AdminPasswordController;
