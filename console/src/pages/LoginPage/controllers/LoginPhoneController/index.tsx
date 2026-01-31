import { memo } from "react";
import { useTranslation } from "react-i18next";
import { Controller, useFormContext } from "react-hook-form";

import type { PhoneLoginFormValues } from "../../schemas/loginSchema";

import styles from "./styles.module.css";

const LoginPhoneController = memo(() => {
  const { t } = useTranslation("LoginPage");
  const {
    control,
    formState: { errors },
  } = useFormContext<PhoneLoginFormValues>();

  const codeError = errors.code;
  const phoneError = errors.phone;

  return (
    <div className={styles.formControl}>
      <div className={styles.phoneInputWrapper}>
        {/* 固定的 + 号 */}
        <span className={styles.plusSign}>+</span>
        {/* 代码输入框 */}
        <div className={styles.codeInputWrapper}>
          <Controller
            name="code"
            control={control}
            render={({ field }) => (
              <div className={styles.inputWrapper}>
                <input
                  {...field}
                  type="tel"
                  className={`${styles.codeInput} ${codeError ? styles.error : ""}`}
                  inputMode="numeric"
                  placeholder="86"
                  onChange={(e) => {
                    const value = e.target.value.replace(/\D/g, "");
                    field.onChange(value);
                  }}
                />
                <label className={styles.codeLabel}>{t("code")}</label>
                <div className={styles.textWrapper}>
                  <div className={styles.innerTextWrapper}>
                    {codeError && <span className={styles.errorText}>{codeError.message}</span>}
                  </div>
                </div>
              </div>
            )}
          />
        </div>
        {/* 手机号输入框 */}
        <div className={styles.phoneInputContainer}>
          <Controller
            name="phone"
            control={control}
            render={({ field }) => (
              <div className={styles.inputWrapper}>
                <input
                  {...field}
                  type="tel"
                  className={`${styles.input} ${phoneError ? styles.error : ""}`}
                  autoComplete="tel"
                  inputMode="numeric"
                  onChange={(e) => {
                    const value = e.target.value.replace(/\D/g, "");
                    field.onChange(value);
                  }}
                />
                <label className={styles.label}>{t("phone")}</label>
                <div className={styles.textWrapper}>
                  <div className={styles.innerTextWrapper}>
                    {phoneError && <span className={styles.errorText}>{phoneError.message}</span>}
                  </div>
                </div>
              </div>
            )}
          />
        </div>
      </div>
    </div>
  );
});

LoginPhoneController.displayName = "LoginPhoneController";

export default LoginPhoneController;
