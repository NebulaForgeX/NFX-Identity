import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo, useState } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";
import { Eye, EyeOff } from "@/assets/icons/lucide";

import styles from "./styles.module.css";

const AdminPasswordConfirmController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();
  const [showPassword, setShowPassword] = useState(false);

  return (
    <div className={styles.formControl}>
      <label className={styles.label}>
        {t("admin_password_confirm.label")} <span className={styles.required}>{t("required")}</span>
      </label>
      <div className={styles.passwordWrapper}>
        <input
          {...register("AdminPasswordConfirm")}
          type={showPassword ? "text" : "password"}
          placeholder={t("admin_password_confirm.placeholder")}
          className={`${styles.input} ${errors.AdminPasswordConfirm ? styles.inputError : ""}`}
          autoComplete="new-password"
        />
        <button
          type="button"
          className={styles.toggleButton}
          onClick={() => setShowPassword(!showPassword)}
          aria-label={showPassword ? t("admin_password_confirm.hide_password") : t("admin_password_confirm.show_password")}
        >
          {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
        </button>
      </div>
      {errors.AdminPasswordConfirm && <p className={styles.errorMessage}>{errors.AdminPasswordConfirm.message}</p>}
    </div>
  );
});

AdminPasswordConfirmController.displayName = "AdminPasswordConfirmController";

export default AdminPasswordConfirmController;
