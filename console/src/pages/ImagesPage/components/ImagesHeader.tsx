import { memo } from "react";
import { useTranslation } from "react-i18next";
import styles from "../styles.module.css";

const ImagesHeader = memo(() => {
  const { t } = useTranslation("ImagesPage");
  return (
    <div className={styles.header}>
      <h1 className={styles.title}>{t("title", "My Images")}</h1>
      <p className={styles.subtitle}>{t("subtitle", "Manage your background images")}</p>
    </div>
  );
});
ImagesHeader.displayName = "ImagesHeader";
export { ImagesHeader };
