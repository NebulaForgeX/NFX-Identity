import { memo, type RefObject } from "react";
import { useTranslation } from "react-i18next";
import { ImagePlus, Upload } from "@/assets/icons/lucide";

import styles from "../styles.module.css";

export interface ImagesUploadAreaProps {
  fileInputRef: RefObject<HTMLInputElement | null>;
  isDragging: boolean;
  isUploading: boolean;
  onDragOver: (e: React.DragEvent) => void;
  onDragLeave: (e: React.DragEvent) => void;
  onDrop: (e: React.DragEvent) => void;
  onClick: () => void;
  onInputChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  disabled?: boolean;
}

export const ImagesUploadArea = memo(function ImagesUploadArea(
  props: ImagesUploadAreaProps
) {
  const { t } = useTranslation("ImagesPage");
  const {
    fileInputRef,
    isDragging,
    isUploading,
    onDragOver,
    onDragLeave,
    onDrop,
    onClick,
    onInputChange,
    disabled,
  } = props;

  return (
    <div
      className={`${styles.uploadArea} ${isDragging ? styles.dragging : ""}`}
      onClick={onClick}
      onDragOver={onDragOver}
      onDragLeave={onDragLeave}
      onDrop={onDrop}
    >
      <div className={styles.uploadPlaceholder}>
        {isUploading ? (
          <>
            <Upload size={48} className={styles.uploadingIcon} />
            <p className={styles.uploadText}>{t("uploading", "Uploading...")}</p>
          </>
        ) : (
          <>
            <ImagePlus size={48} />
            <p className={styles.uploadText}>{t("dragDrop", "Drag and drop an image here")}</p>
            <p className={styles.uploadHint}>{t("orClick", "or click to select a file")}</p>
            <p className={styles.uploadLimit}>
              {t("sizeLimit", "Max size: 10MB. Formats: JPEG, PNG, GIF, WebP")}
            </p>
          </>
        )}
      </div>
      <input
        ref={fileInputRef}
        type="file"
        accept="image/*"
        onChange={onInputChange}
        className={styles.fileInput}
        disabled={disabled}
      />
    </div>
  );
});
