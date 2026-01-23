import type { BootstrapFormValues } from "../../controllers/bootstrapSchema";

import { memo } from "react";
import { useTranslation } from "react-i18next";
import { Controller, useFormContext } from "react-hook-form";

import { Input } from "@/components";

import styles from "./styles.module.css";

const AdminPhoneController = memo(() => {
  const { t } = useTranslation("elements.bootstrap");
  const {
    control,
    register,
    formState: { errors },
  } = useFormContext<BootstrapFormValues>();

  return (
    <div className={styles.formControl}>
      <label className={styles.label}>{t("admin_phone.label")}</label>
      <div className={styles.phoneWrapper}>
        <Controller
          name="AdminCountryCode"
          control={control}
          render={({ field }) => (
            <Input
              {...field}
              onChange={(e) => {
                let value = e.target.value.replace(/\+/g, "").replace(/[^\d]/g, "");
                field.onChange(value);
              }}
              leftIcon={<span className={styles.plusSign}>+</span>}
              placeholder={t("admin_phone.country_code_placeholder")}
              error={errors.AdminCountryCode?.message as string | undefined}
              maxLength={5}
              inputMode="numeric"
              className={styles.countryCodeInput}
            />
          )}
        />
        <Input
          placeholder={t("admin_phone.placeholder")}
          type="tel"
          error={errors.AdminPhone?.message as string | undefined}
          autoComplete="tel"
          inputMode="numeric"
          fullWidth
          {...register("AdminPhone")}
        />
      </div>
    </div>
  );
});

AdminPhoneController.displayName = "AdminPhoneController";

export default AdminPhoneController;
