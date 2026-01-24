import { memo } from "react";
import { useTranslation } from "react-i18next";
import styles from "./styles.module.css";

interface LoginTypeSwitchControllerProps {
  loginType: "email" | "phone";
  onLoginTypeChange: (type: "email" | "phone") => void;
}

const LoginTypeSwitchController = memo(({ loginType, onLoginTypeChange }: LoginTypeSwitchControllerProps) => {
  const { t } = useTranslation("LoginPage");

  const handleTypeChange = (type: "email" | "phone") => {
    onLoginTypeChange(type);
  };

  return (
    <div className={styles.loginTypeSwitch}>
      <button
        type="button"
        className={`${styles.loginTypeBtn} ${loginType === "email" ? styles.active : ""}`}
        onClick={() => handleTypeChange("email")}
      >
        {t("loginType.email")}
      </button>
      <button
        type="button"
        className={`${styles.loginTypeBtn} ${loginType === "phone" ? styles.active : ""}`}
        onClick={() => handleTypeChange("phone")}
      >
        {t("loginType.phone")}
      </button>
    </div>
  );
});

LoginTypeSwitchController.displayName = "LoginTypeSwitchController";

export default LoginTypeSwitchController;
