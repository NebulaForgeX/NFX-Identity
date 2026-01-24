import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Edit } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useUserEducationsByUserID } from "@/hooks/useDirectory";
import { useIsCurrent } from "@/hooks/useStyles";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

interface UserEducationsCardProps {
  userId: string;
}

const UserEducationsCard = memo(({ userId }: UserEducationsCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <UserEducationsCardContent userId={userId} />
    </Suspense>
  );
});

UserEducationsCard.displayName = "UserEducationsCard";

const UserEducationsCardContent = memo(({ userId }: UserEducationsCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();
  const { data: userEducations } = useUserEducationsByUserID({ userId });

  if (!userEducations || userEducations.length === 0) {
    return null;
  }

  return (
    <div className={styles.card}>
      <h3 className={styles.title}>{t("educations")}</h3>
      <div className={styles.list}>
        {userEducations.map((education) => (
          <div key={education.id} className={styles.item}>
            <div className={styles.header}>
              <span className={styles.school}>{education.school}</span>
              <div className={styles.headerRight}>
                {(() => {
                  const currentStyle = useIsCurrent(education.isCurrent);
                  return currentStyle ? (
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
                  ) : null;
                })()}
                <button
                  className={styles.editButton}
                  onClick={() => navigate(`${ROUTES.EDIT_EDUCATION}?id=${education.id}`)}
                  title={t("edit")}
                >
                  <Edit size={16} />
                </button>
              </div>
            </div>
            {education.degree && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("degree")}:</span>
                <span className={styles.value}>{education.degree}</span>
              </div>
            )}
            {education.major && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("major")}:</span>
                <span className={styles.value}>{education.major}</span>
              </div>
            )}
            {education.fieldOfStudy && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("fieldOfStudy")}:</span>
                <span className={styles.value}>{education.fieldOfStudy}</span>
              </div>
            )}
            {education.startDate && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("startDate")}:</span>
                <span className={styles.value}>
                  {new Date(education.startDate).toLocaleDateString()}
                </span>
              </div>
            )}
            {education.endDate && !education.isCurrent && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("endDate")}:</span>
                <span className={styles.value}>
                  {new Date(education.endDate).toLocaleDateString()}
                </span>
              </div>
            )}
            {education.grade && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("grade")}:</span>
                <span className={styles.value}>{education.grade}</span>
              </div>
            )}
            {education.description && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("description")}:</span>
                <span className={styles.value}>{education.description}</span>
              </div>
            )}
            {education.activities && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("activities")}:</span>
                <span className={styles.value}>{education.activities}</span>
              </div>
            )}
            {education.achievements && (
              <div className={styles.detail}>
                <span className={styles.label}>{t("achievements")}:</span>
                <span className={styles.value}>{education.achievements}</span>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
});

UserEducationsCardContent.displayName = "UserEducationsCardContent";

export default UserEducationsCard;
