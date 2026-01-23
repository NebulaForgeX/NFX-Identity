import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import styles from "./styles.module.css";

const AdminUsernameController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <div className={styles.formControl}>
      <label className={styles.label}>
        {t("admin_username.label")} <span className={styles.required}>{t("required")}</span>
      </label>
      <input
        {...register("AdminUsername")}
        type="text"
        placeholder={t("admin_username.placeholder")}
        className={`${styles.input} ${errors.AdminUsername ? styles.inputError : ""}`}
        autoComplete="username"
      />
      {errors.AdminUsername && <p className={styles.errorMessage}>{errors.AdminUsername.message}</p>}
    </div>
  );
});

AdminUsernameController.displayName = "AdminUsernameController";

export default AdminUsernameController;
