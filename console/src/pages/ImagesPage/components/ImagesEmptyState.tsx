import { memo } from "react";
import { useTranslation } from "react-i18next";
import { ImagePlus } from "@/assets/icons/lucide";
import styles from "../styles.module.css";

export const ImagesEmptyState = memo(function ImagesEmptyState() {
  const { t } = useTranslation("ImagesPage");
  return (
    <div className={styles.emptyState}>
      <ImagePlus size={64} />
      <p>{t("noImages", "No images yet")}</p>
      <p className={styles.emptyHint}>{t("uploadFirst", "Upload your first background image")}</p>
    </div>
  );
});
