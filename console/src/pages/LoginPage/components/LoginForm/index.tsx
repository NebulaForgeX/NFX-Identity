import { memo, useState } from "react";
import { useTranslation } from "react-i18next";

import { useLoginByEmail, useLoginByPhone } from "@/hooks/useAuth";

import styles from "./styles.module.css";

const LoginForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const [loginType, setLoginType] = useState<"email" | "phone">("email");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [countryCode, setCountryCode] = useState("+86");
  const [password, setPassword] = useState("");

  const { mutateAsync: loginByEmail, isPending: isEmailLoginPending } = useLoginByEmail();
  const { mutateAsync: loginByPhone, isPending: isPhoneLoginPending } = useLoginByPhone();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (loginType === "email") {
        await loginByEmail({ email, password });
      } else {
        await loginByPhone({ phone, password, countryCode });
      }
      // 跳转由 App.tsx 中的事件监听器统一处理
    } catch (error) {
      // 错误已在 useAuth hook 中处理
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <span className={styles.title}>{t("login")}</span>

      {/* Login Type Toggle */}
      <div className={styles.loginTypeSwitch}>
        <button
          type="button"
          className={`${styles.loginTypeBtn} ${loginType === "email" ? styles.active : ""}`}
          onClick={() => setLoginType("email")}
        >
          {t("loginType.email")}
        </button>
        <button
          type="button"
          className={`${styles.loginTypeBtn} ${loginType === "phone" ? styles.active : ""}`}
          onClick={() => setLoginType("phone")}
        >
          {t("loginType.phone")}
        </button>
      </div>

      {/* Email or Phone */}
      {loginType === "email" ? (
        <div className={styles.formControl}>
          <input
            type="email"
            className={styles.input}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            autoComplete="email"
            required
          />
          <label className={styles.label}>{t("email")}</label>
        </div>
      ) : (
        <div className={styles.formControl}>
          <div className={styles.phoneInputWrapper}>
            <input
              type="text"
              className={styles.countryCodeInput}
              value={countryCode}
              onChange={(e) => setCountryCode(e.target.value)}
              placeholder="+86"
            />
            <input
              type="tel"
              className={styles.input}
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
              autoComplete="tel"
              required
            />
            <label className={styles.label}>{t("phone")}</label>
          </div>
        </div>
      )}

      {/* Password */}
      <div className={styles.formControl}>
        <input
          type="password"
          className={styles.input}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          autoComplete="current-password"
          required
        />
        <label className={styles.label}>{t("password")}</label>
      </div>

      <button type="submit" className={styles.submitBtn} disabled={isEmailLoginPending || isPhoneLoginPending}>
        {isEmailLoginPending || isPhoneLoginPending ? t("loggingIn") : t("login")}
      </button>

      <span className={styles.bottomText}>
        {t("noAccount")}{" "}
        <label htmlFor="register_toggle" className={styles.switch}>
          {t("registerNow")}
        </label>
      </span>
    </form>
  );
});

LoginForm.displayName = "LoginForm";

export default LoginForm;
