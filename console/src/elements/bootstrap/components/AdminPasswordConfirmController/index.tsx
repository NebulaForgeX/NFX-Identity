import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo, useState } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";
import { Eye, EyeOff } from "@/assets/icons/lucide";

import { Button, Input } from "@/components";

const AdminPasswordConfirmController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();
  const [showPassword, setShowPassword] = useState(false);

  return (
    <Input
      label={t("admin_password_confirm.label")}
      type={showPassword ? "text" : "password"}
      placeholder={t("admin_password_confirm.placeholder")}
      error={errors.AdminPasswordConfirm?.message as string | undefined}
      fullWidth
      required
      autoComplete="new-password"
      rightIconInteractive
      rightIcon={
        <Button
          type="button"
          variant="ghost"
          size="small"
          aria-label={
            showPassword ? t("admin_password_confirm.hide_password") : t("admin_password_confirm.show_password")
          }
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
        </Button>
      }
      {...register("AdminPasswordConfirm")}
    />
  );
});

AdminPasswordConfirmController.displayName = "AdminPasswordConfirmController";

export default AdminPasswordConfirmController;
