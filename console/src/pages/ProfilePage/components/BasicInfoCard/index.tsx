import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Edit } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useUserProfile } from "@/hooks/useDirectory";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

interface BasicInfoCardProps {
  userId: string;
}

const BasicInfoCard = memo(({ userId }: BasicInfoCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <BasicInfoCardContent userId={userId} />
    </Suspense>
  );
});

BasicInfoCard.displayName = "BasicInfoCard";

const BasicInfoCardContent = memo(({ userId }: BasicInfoCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();
  const { data: userProfile } = useUserProfile({ id: userId });

  return (
    <div className={styles.detailCard}>
      <div className={styles.cardHeader}>
        <h3 className={styles.detailTitle}>{t("basicInfo")}</h3>
        <button
          className={styles.editButton}
          onClick={() => navigate(ROUTES.EDIT_PROFILE)}
          title={t("edit")}
        >
          <Edit size={16} />
        </button>
      </div>
      <div className={styles.detailList}>
        {userProfile?.role && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("role")}:</span>
            <span className={styles.detailValue}>{userProfile.role}</span>
          </div>
        )}
        {userProfile?.firstName && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("firstName")}:</span>
            <span className={styles.detailValue}>{userProfile.firstName}</span>
          </div>
        )}
        {userProfile?.lastName && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("lastName")}:</span>
            <span className={styles.detailValue}>{userProfile.lastName}</span>
          </div>
        )}
        {userProfile?.nickname && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("nickname")}:</span>
            <span className={styles.detailValue}>{userProfile.nickname}</span>
          </div>
        )}
        {userProfile?.displayName && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("displayName")}:</span>
            <span className={styles.detailValue}>{userProfile.displayName}</span>
          </div>
        )}
        {userProfile?.gender && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("gender")}:</span>
            <span className={styles.detailValue}>{userProfile.gender}</span>
          </div>
        )}
        {userProfile?.birthday && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("birthday")}:</span>
            <span className={styles.detailValue}>
              {new Date(userProfile.birthday).toLocaleDateString()}
            </span>
          </div>
        )}
        {userProfile?.age && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("age")}:</span>
            <span className={styles.detailValue}>{userProfile.age}</span>
          </div>
        )}
        {userProfile?.location && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("location")}:</span>
            <span className={styles.detailValue}>{userProfile.location}</span>
          </div>
        )}
        {userProfile?.website && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("website")}:</span>
            <span className={styles.detailValue}>
              <a href={userProfile.website} target="_blank" rel="noopener noreferrer">
                {userProfile.website}
              </a>
            </span>
          </div>
        )}
        {userProfile?.github && (
          <div className={styles.detailItem}>
            <span className={styles.detailLabel}>{t("github")}:</span>
            <span className={styles.detailValue}>
              <a
                href={`https://github.com/${userProfile.github}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                {userProfile.github}
              </a>
            </span>
          </div>
        )}
      </div>
    </div>
  );
});

BasicInfoCardContent.displayName = "BasicInfoCardContent";

export default BasicInfoCard;
