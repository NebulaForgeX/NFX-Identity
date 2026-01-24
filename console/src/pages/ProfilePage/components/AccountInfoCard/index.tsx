import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Suspense } from "@/components";
import { useUser } from "@/hooks/useDirectory";
import { useStatus, useVerification } from "@/hooks/useStyles";

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

  const statusStyle = useStatus(user?.status);
  const verificationStyle = useVerification(user?.isVerified || false);

  return (
    <div className={styles.detailCard}>
      <h3 className={styles.detailTitle}>{t("accountInfo")}</h3>
      <div className={styles.detailList}>
        <div className={styles.detailItem}>
          <span className={styles.detailLabel}>{t("userId")}:</span>
          <span className={styles.detailValue}>{user.id}</span>
        </div>
        {user.status && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("status")}:</span>
            <span className={styles.detailValue}>
              <span
                style={{
                  padding: "0.25rem 0.5rem",
                  borderRadius: "0.25rem",
                  fontSize: "0.75rem",
                  fontWeight: 600,
                  backgroundColor: statusStyle.bgColor,
                  color: statusStyle.color,
                }}
              >
                {statusStyle.label}
              </span>
            </span>
          </div>
        )}
        <div className={styles.detailItem}>
          <span className={styles.detailLabel}>{t("isVerified")}:</span>
          <span className={styles.detailValue}>
            <span
              style={{
                padding: "0.25rem 0.5rem",
                borderRadius: "0.25rem",
                fontSize: "0.75rem",
                fontWeight: 600,
                backgroundColor: verificationStyle.bgColor,
                color: verificationStyle.color,
              }}
            >
              {verificationStyle.label}
            </span>
          </span>
        </div>
        {user.lastLoginAt && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("lastLoginAt")}:</span>
            <span className={styles.detailValue}>
              {new Date(user.lastLoginAt).toLocaleString()}
            </span>
          </div>
        )}
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
