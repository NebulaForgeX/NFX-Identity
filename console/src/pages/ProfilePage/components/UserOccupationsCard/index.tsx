import { memo } from "react";
import { useTranslation } from "react-i18next";

import { Edit } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useUserOccupationsByUserID } from "@/hooks/useDirectory";
import { useIsCurrent } from "@/hooks/useStyles";
import { routerEventEmitter } from "@/events/router";
import { ROUTES } from "@/types/navigation";
import type { UserOccupation } from "@/types";

import styles from "./styles.module.css";

interface UserOccupationsCardProps {
  userId: string;
}

const UserOccupationsCard = memo(({ userId }: UserOccupationsCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <UserOccupationsCardContent userId={userId} />
    </Suspense>
  );
});

UserOccupationsCard.displayName = "UserOccupationsCard";

interface OccupationItemProps {
  occupation: UserOccupation;
  onEdit: (id: string) => void;
}

const OccupationItem = memo(({ occupation, onEdit }: OccupationItemProps) => {
  const { t } = useTranslation("ProfilePage");
  const currentStyle = useIsCurrent(occupation.isCurrent);

  return (
    <div className={styles.item}>
      <div className={styles.header}>
        <div>
          <span className={styles.position}>{occupation.position}</span>
          <span className={styles.company}> @ {occupation.company}</span>
        </div>
        <div className={styles.headerRight}>
          {currentStyle && (
            <span
              style={{
                padding: "0.25rem 0.5rem",
                borderRadius: "0.25rem",
                fontSize: "0.75rem",
                fontWeight: 600,
                backgroundColor: currentStyle.bgColor,
                color: currentStyle.color,
              }}
            >
              {currentStyle.label}
            </span>
          )}
          <button
            className={styles.editButton}
            onClick={() => onEdit(occupation.id)}
            title={t("edit")}
          >
            <Edit size={16} />
          </button>
        </div>
      </div>
      {occupation.department && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("department")}:</span>
          <span className={styles.value}>{occupation.department}</span>
        </div>
      )}
      {occupation.industry && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("industry")}:</span>
          <span className={styles.value}>{occupation.industry}</span>
        </div>
      )}
      {occupation.location && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("location")}:</span>
          <span className={styles.value}>{occupation.location}</span>
        </div>
      )}
      {occupation.employmentType && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("employmentType")}:</span>
          <span className={styles.value}>{occupation.employmentType}</span>
        </div>
      )}
      {occupation.startDate && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("startDate")}:</span>
          <span className={styles.value}>
            {new Date(occupation.startDate).toLocaleDateString()}
          </span>
        </div>
      )}
      {occupation.endDate && !occupation.isCurrent && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("endDate")}:</span>
          <span className={styles.value}>
            {new Date(occupation.endDate).toLocaleDateString()}
          </span>
        </div>
      )}
      {occupation.startDate && occupation.isCurrent && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("period")}:</span>
          <span className={styles.value}>
            {new Date(occupation.startDate).toLocaleDateString()} - {t("present")}
          </span>
        </div>
      )}
      {occupation.description && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("description")}:</span>
          <span className={styles.value}>{occupation.description}</span>
        </div>
      )}
      {occupation.responsibilities && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("responsibilities")}:</span>
          <span className={styles.value}>{occupation.responsibilities}</span>
        </div>
      )}
      {occupation.achievements && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("achievements")}:</span>
          <span className={styles.value}>{occupation.achievements}</span>
        </div>
      )}
      {occupation.skillsUsed && occupation.skillsUsed.length > 0 && (
        <div className={styles.detail}>
          <span className={styles.label}>{t("skillsUsed")}:</span>
          <span className={styles.value}>{occupation.skillsUsed.join(", ")}</span>
        </div>
      )}
    </div>
  );
});

OccupationItem.displayName = "OccupationItem";

const UserOccupationsCardContent = memo(({ userId }: UserOccupationsCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const { data: userOccupations } = useUserOccupationsByUserID({ userId });

  if (!userOccupations || userOccupations.length === 0) {
    return null;
  }

  const handleEdit = (id: string) => {
    routerEventEmitter.navigateToEditOccupation(id);
  };

  return (
    <div className={styles.card}>
      <h3 className={styles.title}>{t("occupations")}</h3>
      <div className={styles.list}>
        {userOccupations.map((occupation) => (
          <OccupationItem key={occupation.id} occupation={occupation} onEdit={handleEdit} />
        ))}
      </div>
    </div>
  );
});

UserOccupationsCardContent.displayName = "UserOccupationsCardContent";

export default UserOccupationsCard;
