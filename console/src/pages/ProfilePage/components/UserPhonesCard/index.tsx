import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Suspense } from "@/components";
import { useUserPhonesByUserID } from "@/hooks/useDirectory";
import { useIsPrimary, useIsVerified } from "@/hooks/useStyles";

import styles from "./styles.module.css";

interface UserPhonesCardProps {
  userId: string;
}

const UserPhonesCard = memo(({ userId }: UserPhonesCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <UserPhonesCardContent userId={userId} />
    </Suspense>
  );
});

UserPhonesCard.displayName = "UserPhonesCard";

const UserPhonesCardContent = memo(({ userId }: UserPhonesCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const { data: userPhones } = useUserPhonesByUserID({ userId });

  if (!userPhones || userPhones.length === 0) {
    return null;
  }

  return (
    <div className={styles.card}>
      <h3 className={styles.title}>{t("phones")}</h3>
      <div className={styles.list}>
        {userPhones.map((phone) => {
          const primaryStyle = useIsPrimary(phone.isPrimary);
          const verifiedStyle = useIsVerified(phone.isVerified);

          return (
            <div key={phone.id} className={styles.item}>
              <span className={styles.phone}>
                {phone.countryCode && <span className={styles.countryCode}>+{phone.countryCode}</span>}
                {phone.phone}
              </span>
              {phone.isPrimary && (
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
                  {t("primary")}
                </span>
              )}
              {phone.isVerified && (
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
                  {t("verified")}
                </span>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
});

UserPhonesCardContent.displayName = "UserPhonesCardContent";

export default UserPhonesCard;
