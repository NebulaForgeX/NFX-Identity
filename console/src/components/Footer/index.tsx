import { memo } from "react";
import { useTranslation } from "react-i18next";

import styles from "./styles.module.css";

interface FooterProps {
  className?: string;
}

const Footer = memo(({ className }: FooterProps) => {
  const { t } = useTranslation("components");
  const currentYear = new Date().getFullYear();

  return (
    <footer className={`${styles.footer} ${className || ""}`}>
      <div className={styles.footerContent}>
        <span className={styles.copyright}>© {currentYear} NebulaForgeX Identity System. {t("footer.allRightsReserved")}</span>
        <div className={styles.links}>
          <a href="#" className={styles.link}>
            {t("footer.aboutUs")}
          </a>
          <a href="#" className={styles.link}>
            {t("footer.privacyPolicy")}
          </a>
          <a href="#" className={styles.link}>
            {t("footer.termsOfService")}
          </a>
        </div>
      </div>
    </footer>
  );
});

Footer.displayName = "Footer";

export default Footer;
