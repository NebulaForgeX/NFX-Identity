import { memo, useState } from "react";
import { useTranslation } from "react-i18next";
import LoginEmailForm from "../LoginEmailForm";
import LoginPhoneForm from "../LoginPhoneForm";
import LoginTypeSwitchController from "../LoginTypeSwitchController";
import styles from "./styles.module.css";

const LoginForm = memo(() => {
  const { t } = useTranslation("LoginPage");
  const [loginType, setLoginType] = useState<"email" | "phone">("email");

  return (
    <div className={styles.formWrapper}>
      <div className={styles.form}>
        <span className={styles.title}>{t("login")}</span>
        <LoginTypeSwitchController loginType={loginType} onLoginTypeChange={setLoginType} />
        {loginType === "email" ? <LoginEmailForm /> : <LoginPhoneForm />}
      </div>
    </div>
  );
});

LoginForm.displayName = "LoginForm";

export default LoginForm;
