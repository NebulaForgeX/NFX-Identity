import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Edit } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useUser, useUserProfile, useUserEmailsByUserID } from "@/hooks/useDirectory";
import { ROUTES } from "@/types/navigation";
import { buildImageUrl } from "@/utils/image";

import styles from "./styles.module.css";

interface ProfileCardProps {
  userId: string;
}

const ProfileCard = memo(({ userId }: ProfileCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <ProfileCardContent userId={userId} />
    </Suspense>
  );
});

ProfileCard.displayName = "ProfileCard";

const ProfileCardContent = memo(({ userId }: ProfileCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();

  const { data: user } = useUser({ id: userId });
  const { data: userProfile } = useUserProfile({ id: userId });
  const { data: userEmailsList } = useUserEmailsByUserID({ userId });

  // 从邮箱列表中提取主邮箱或第一个邮箱
  const userEmails =
    userEmailsList && userEmailsList.length > 0
      ? (userEmailsList.find((email) => email.isPrimary)?.email || userEmailsList[0]?.email || null)
      : null;

  const displayName =
    userProfile?.firstName && userProfile?.lastName
      ? `${userProfile.firstName} ${userProfile.lastName}`
      : userProfile?.displayName || userProfile?.firstName || user?.username || userEmails || t("user");

  return (
    <div className={styles.profileCard}>
      <button
        className={styles.editButton}
        onClick={() => navigate(ROUTES.EDIT_PROFILE)}
        title={t("edit")}
      >
        <Edit size={16} />
      </button>
      <div className={styles.avatarSection}>
        <img
          src={buildImageUrl(userProfile?.avatarId, "avatar") || "/default-avatar.png"}
          alt={displayName}
          className={styles.avatar}
        />
      </div>
      <div className={styles.infoSection}>
        <h2 className={styles.displayName}>{displayName}</h2>
        {userEmails && (
          <div className={styles.email}>
            <span className={styles.label}>{t("email")}:</span>
            <span className={styles.value}>{userEmails}</span>
          </div>
        )}
        {user?.username && (
          <div className={styles.username}>
            <span className={styles.label}>{t("username")}:</span>
            <span className={styles.value}>{user.username}</span>
          </div>
        )}
        {userProfile?.bio && (
          <div className={styles.bio}>
            <span className={styles.label}>{t("bio")}:</span>
            <span className={styles.value}>{userProfile.bio}</span>
          </div>
        )}
      </div>
    </div>
  );
});

ProfileCardContent.displayName = "ProfileCardContent";

export default ProfileCard;
