import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Suspense } from "@/components";
import { useUserEmailsByUserID } from "@/hooks/useDirectory";

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
        {userEmails.map((email) => (
          <div key={email.id} className={styles.item}>
            <span className={styles.email}>{email.email}</span>
            {email.isPrimary && <span className={styles.badge}>{t("primary")}</span>}
            {email.isVerified && <span className={styles.verified}>{t("verified")}</span>}
          </div>
        ))}
      </div>
    </div>
  );
});

UserEmailsCardContent.displayName = "UserEmailsCardContent";

export default UserEmailsCard;
