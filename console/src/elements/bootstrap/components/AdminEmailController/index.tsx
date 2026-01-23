import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import styles from "./styles.module.css";

const AdminEmailController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <div className={styles.formControl}>
      <label className={styles.label}>{t("admin_email.label")}</label>
      <input
        {...register("AdminEmail")}
        type="email"
        placeholder={t("admin_email.placeholder")}
        className={`${styles.input} ${errors.AdminEmail ? styles.inputError : ""}`}
        autoComplete="email"
      />
      {errors.AdminEmail && <p className={styles.errorMessage}>{errors.AdminEmail.message}</p>}
    </div>
  );
});

AdminEmailController.displayName = "AdminEmailController";

export default AdminEmailController;
