import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { Plus, Settings } from "@/assets/icons/lucide";
import { useAuthStore } from "@/stores/authStore";
import { ROUTES } from "@/types/navigation";

import {
  AccountInfoCard,
  BasicInfoCard,
  ProfileCard,
  SkillsCard,
  SocialLinksCard,
  UserEducationsCard,
  UserEmailsCard,
  UserOccupationsCard,
  UserPhonesCard,
} from "./components";
import styles from "./styles.module.css";

const ProfilePage = memo(() => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();
  const currentUserId = useAuthStore((state) => state.currentUserId);

  if (!currentUserId) {
    return (
      <div className={styles.container}>
        <div className={styles.errorContainer}>
          <p>{t("userNotFound")}</p>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1 className={styles.title}>{t("title")}</h1>
        <p className={styles.subtitle}>{t("subtitle")}</p>
      </div>

      <div className={styles.content}>
        {/* Profile Card */}
        <ProfileCard userId={currentUserId} />

        {/* Action Buttons */}
        <div className={styles.actionButtons}>
          <button
            className={styles.actionButton}
            onClick={() => navigate(ROUTES.ADD_EDUCATION)}
          >
            <Plus size={18} />
            <span>{t("addEducation")}</span>
          </button>
          <button
            className={styles.actionButton}
            onClick={() => navigate(ROUTES.ADD_OCCUPATION)}
          >
            <Plus size={18} />
            <span>{t("addOccupation")}</span>
          </button>
          <button
            className={styles.actionButton}
            onClick={() => navigate(ROUTES.EDIT_PREFERENCE)}
          >
            <Settings size={18} />
            <span>{t("editPreference")}</span>
          </button>
        </div>

        {/* Details Section */}
        <div className={styles.detailsSection}>
          <BasicInfoCard userId={currentUserId} />
          <AccountInfoCard userId={currentUserId} />
        </div>

        {/* User Data Cards */}
        <div className={styles.userDataSection}>
          {/* Social Links and Skills Section */}
          <div className={styles.socialSkillsSection}>
            <SocialLinksCard userId={currentUserId} />
            <SkillsCard userId={currentUserId} />
          </div>

          {/* Contact Section - Email and Phone side by side */}
          <div className={styles.contactSection}>
            <UserEmailsCard userId={currentUserId} />
            <UserPhonesCard userId={currentUserId} />
          </div>

          {/* Education and Occupation */}
          <UserEducationsCard userId={currentUserId} />
          <UserOccupationsCard userId={currentUserId} />
        </div>
      </div>
    </div>
  );
});

ProfilePage.displayName = "ProfilePage";

export default ProfilePage;
