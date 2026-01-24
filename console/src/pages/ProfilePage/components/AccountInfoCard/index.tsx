import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Suspense } from "@/components";
import { useUser } from "@/hooks/useDirectory";

import styles from "./styles.module.css";

interface AccountInfoCardProps {
  userId: string;
}

const AccountInfoCard = memo(({ userId }: AccountInfoCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <AccountInfoCardContent userId={userId} />
    </Suspense>
  );
});

AccountInfoCard.displayName = "AccountInfoCard";

const AccountInfoCardContent = memo(({ userId }: AccountInfoCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const { data: user } = useUser({ id: userId });

  return (
    <div className={styles.detailCard}>
      <h3 className={styles.detailTitle}>{t("accountInfo")}</h3>
      <div className={styles.detailList}>
        <div className={styles.detailItem}>
          <span className={styles.detailLabel}>{t("userId")}:</span>
          <span className={styles.detailValue}>{user.id}</span>
        </div>
        {user.createdAt && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("createdAt")}:</span>
            <span className={styles.detailValue}>
              {new Date(user.createdAt).toLocaleString()}
            </span>
          </div>
        )}
        {user.updatedAt && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("updatedAt")}:</span>
            <span className={styles.detailValue}>
              {new Date(user.updatedAt).toLocaleString()}
            </span>
          </div>
        )}
      </div>
    </div>
  );
});

AccountInfoCardContent.displayName = "AccountInfoCardContent";

export default AccountInfoCard;
