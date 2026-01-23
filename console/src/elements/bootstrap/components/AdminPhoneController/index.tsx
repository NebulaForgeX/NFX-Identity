import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import styles from "./styles.module.css";

const AdminPhoneController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <div className={styles.formControl}>
      <label className={styles.label}>{t("admin_phone.label")}</label>
      <div className={styles.phoneWrapper}>
        <input
          {...register("AdminCountryCode")}
          type="text"
          placeholder={t("admin_phone.country_code_placeholder")}
          className={`${styles.countryCode} ${errors.AdminCountryCode ? styles.inputError : ""}`}
          maxLength={5}
        />
        <input
          {...register("AdminPhone")}
          type="tel"
          placeholder={t("admin_phone.placeholder")}
          className={`${styles.phone} ${errors.AdminPhone ? styles.inputError : ""}`}
          autoComplete="tel"
        />
      </div>
      {errors.AdminPhone && <p className={styles.errorMessage}>{errors.AdminPhone.message}</p>}
      {errors.AdminCountryCode && <p className={styles.errorMessage}>{errors.AdminCountryCode.message}</p>}
    </div>
  );
});

AdminPhoneController.displayName = "AdminPhoneController";

export default AdminPhoneController;
