import { memo } from "react";
import { useTranslation } from "react-i18next";

import Button from "@/components/Button";

import styles from "../styles.module.css";

interface ImagesConfirmSectionProps {
  uploadedImageId: string;
  isConfirming: boolean;
  onCancel: () => void;
  onConfirm: () => void;
}

export const ImagesConfirmSection = memo(function ImagesConfirmSection({
  uploadedImageId,
  isConfirming,
  onCancel,
  onConfirm,
}: ImagesConfirmSectionProps) {
  const { t } = useTranslation("ImagesPage");
  if (!uploadedImageId) return null;
  return (
    <div className={styles.confirmSection}>
      <p className={styles.confirmText}>
        {t("confirmUpload", "Image uploaded to temporary storage. Click confirm to add to your images.")}
      </p>
      <div className={styles.confirmActions}>
        <Button variant="outline" className={styles.cancelButton} onClick={onCancel} disabled={isConfirming}>
          {t("cancel", "Cancel")}
        </Button>
        <Button variant="primary" className={styles.confirmButton} onClick={onConfirm} loading={isConfirming} disabled={isConfirming}>
          {isConfirming ? t("confirming", "Confirming...") : t("confirm", "Confirm")}
        </Button>
      </div>
    </div>
  );
});
