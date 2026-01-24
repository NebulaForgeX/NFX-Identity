import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Edit, ExternalLink } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useUserProfile } from "@/hooks/useDirectory";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

interface SocialLinksCardProps {
  userId: string;
}

const SocialLinksCard = memo(({ userId }: SocialLinksCardProps) => {
  const { t } = useTranslation("ProfilePage");

  return (
    <Suspense
      loadingType="ecg"
      loadingText={t("loading")}
      loadingSize="small"
      loadingContainerClassName={styles.loading}
    >
      <SocialLinksCardContent userId={userId} />
    </Suspense>
  );
});

SocialLinksCard.displayName = "SocialLinksCard";

const SocialLinksCardContent = memo(({ userId }: SocialLinksCardProps) => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();
  const { data: userProfile } = useUserProfile({ id: userId });

  const socialLinks = userProfile?.socialLinks as Record<string, string> | undefined;

  if (!socialLinks || Object.keys(socialLinks).length === 0) {
    return null;
  }

  return (
    <div className={styles.card}>
      <div className={styles.cardHeader}>
        <h3 className={styles.title}>{t("socialLinks")}</h3>
        <button
          className={styles.editButton}
          onClick={() => navigate(ROUTES.EDIT_PROFILE)}
          title={t("edit")}
        >
          <Edit size={16} />
        </button>
      </div>
      <div className={styles.list}>
        {Object.entries(socialLinks).map(([key, value]) => (
          <div key={key} className={styles.item}>
            <span className={styles.label}>{key}:</span>
            <a
              href={typeof value === "string" ? value : String(value)}
              target="_blank"
              rel="noopener noreferrer"
              className={styles.link}
            >
              <span className={styles.value}>{typeof value === "string" ? value : String(value)}</span>
              <ExternalLink size={14} className={styles.icon} />
            </a>
          </div>
        ))}
      </div>
    </div>
  );
});

SocialLinksCardContent.displayName = "SocialLinksCardContent";

export default SocialLinksCard;
