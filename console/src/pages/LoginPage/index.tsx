import { useState } from "react";
import { useTranslation } from "react-i18next";

import { TruckLoading } from "@/animations";
import LanguageSwitcher from "@/components/LanguageSwitcher";
import ThemeSwitcher from "@/components/ThemeSwitcher";

import { LoginForm, RegisterForm } from "./components";
import styles from "./styles.module.css";

export default function LoginPage() {
  const { t } = useTranslation("LoginPage");
  const [isRegister, setIsRegister] = useState(false);
  const isLoading = false;

  if (isLoading) {
    return (
      <div className={styles.loginPage}>
        <div className={styles.loadingContainer}>
          <TruckLoading size="medium" />
          <p className={styles.loadingText}>{t("loggingIn")}</p>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.loginPage}>
      {/* 顶部控制栏：语言和主题切换 */}
      <div className={styles.topControls}>
        <LanguageSwitcher status="default" />
        <ThemeSwitcher status="default" />
      </div>

      <div className={styles.container}>
        <input
          type="checkbox"
          id="register_toggle"
          checked={isRegister}
          onChange={(e) => setIsRegister(e.target.checked)}
          className={styles.registerToggle}
        />
        <div className={styles.slider}>
          <LoginForm />
          <RegisterForm />
        </div>
      </div>
    </div>
  );
}
