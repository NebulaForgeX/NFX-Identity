import { memo } from "react";
import { useTranslation } from "react-i18next";
import { Trash2 } from "@/assets/icons/lucide";

import { buildImageUrl } from "@/utils/image";
import { showConfirm } from "@/stores/modalStore";
import type { UserImage } from "@/types";

import styles from "../styles.module.css";

interface ImagesGalleryProps {
  userImages: UserImage[];
  onDelete: (userImageId: string) => void;
}

export const ImagesGallery = memo(function ImagesGallery({ userImages, onDelete }: ImagesGalleryProps) {
  const { t } = useTranslation("ImagesPage");

  const handleDeleteClick = (userImageId: string) => {
    showConfirm({
      title: t("deleteConfirmTitle", "Delete Image"),
      message: t("deleteConfirmMessage", "Are you sure you want to delete this image?"),
      onConfirm: () => onDelete(userImageId),
    });
  };

  if (userImages.length === 0) return null;

  return (
    <div className={styles.imagesGrid}>
      {userImages.map((userImage, index) => (
        <div key={userImage.id} className={styles.imageCard}>
          <div className={styles.imageWrapper}>
            <img
              src={buildImageUrl(userImage.imageId)}
              alt={`Background ${index + 1}`}
              className={styles.image}
            />
            {index === 0 && (
              <span className={styles.primaryBadge}>{t("primary", "Primary")}</span>
            )}
            <div className={styles.imageOverlay}>
              <button
                className={styles.deleteButton}
                onClick={() => handleDeleteClick(userImage.id)}
                title={t("delete", "Delete")}
              >
                <Trash2 size={20} />
              </button>
            </div>
          </div>
          <div className={styles.imageInfo}>
            <span className={styles.imageOrder}>#{index + 1}</span>
          </div>
        </div>
      ))}
    </div>
  );
});
