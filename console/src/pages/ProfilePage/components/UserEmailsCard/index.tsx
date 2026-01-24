import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Suspense } from "@/components";
import { useUserEmailsByUserID } from "@/hooks/useDirectory";
import { useIsPrimary, useIsVerified } from "@/hooks/useStyles";

import styles from "./styles.module.css";

interface UserEmailsCardProps {
  userId: string;
}

const UserEmailsCard = memo(({ userId }: UserEmailsCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <UserEmailsCardContent userId={userId} />
    </Suspense>
  );
});

UserEmailsCard.displayName = "UserEmailsCard";

const UserEmailsCardContent = memo(({ userId }: UserEmailsCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const { data: userEmails } = useUserEmailsByUserID({ userId });

  if (!userEmails || userEmails.length === 0) {
    return null;
  }

  return (
    <div className={styles.card}>
      <h3 className={styles.title}>{t("emails")}</h3>
      <div className={styles.list}>
        {userEmails.map((email) => {
          const primaryStyle = useIsPrimary(email.isPrimary);
          const verifiedStyle = useIsVerified(email.isVerified);

          return (
            <div key={email.id} className={styles.item}>
              <span className={styles.email}>{email.email}</span>
              {primaryStyle && (
                <span
                  style={{
                    padding: "0.25rem 0.5rem",
                    borderRadius: "0.25rem",
                    fontSize: "0.75rem",
                    fontWeight: 600,
                    backgroundColor: primaryStyle.bgColor,
                    color: primaryStyle.color,
                  }}
                >
                  {primaryStyle.label}
                </span>
              )}
              {verifiedStyle && (
                <span
                  style={{
                    padding: "0.25rem 0.5rem",
                    borderRadius: "0.25rem",
                    fontSize: "0.75rem",
                    fontWeight: 600,
                    backgroundColor: verifiedStyle.bgColor,
                    color: verifiedStyle.color,
                  }}
                >
                  {verifiedStyle.label}
                </span>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
});

UserEmailsCardContent.displayName = "UserEmailsCardContent";

export default UserEmailsCard;
