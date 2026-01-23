import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useFormContext } from "react-hook-form";

import styles from "./styles.module.css";

const VersionController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <div className={styles.formControl}>
      <label className={styles.label}>
        {t("version.label")} <span className={styles.required}>{t("required")}</span>
      </label>
      <input
        {...register("Version")}
        type="text"
        placeholder={t("version.placeholder")}
        className={`${styles.input} ${errors.Version ? styles.inputError : ""}`}
      />
      {errors.Version && <p className={styles.errorMessage}>{errors.Version.message}</p>}
    </div>
  );
});

VersionController.displayName = "VersionController";

export default VersionController;
