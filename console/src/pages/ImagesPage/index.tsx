import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Suspense } from "@/components";
import { useAuthStore } from "@/stores/authStore";

import { ImagesContent } from "./ImagesContent";
import { ImagesHeader } from "./components/ImagesHeader";
import styles from "./styles.module.css";

const ImagesPage = memo(function ImagesPage() {
  const { t } = useTranslation("ImagesPage");
  const currentUserId = useAuthStore((state) => state.currentUserId);

  if (!currentUserId) {
    return (
      <div className={styles.container}>
        <div className={styles.errorContainer}>
          <p>{t("userNotFound", "User not found")}</p>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <ImagesHeader />
      <Suspense loadingType="ecg" loadingText={t("loading", "Loading images...")}>
        <ImagesContent userId={currentUserId} />
      </Suspense>
    </div>
  );
});

ImagesPage.displayName = "ImagesPage";
export default ImagesPage;
