import { memo } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { ImagePlus, Plus, Settings } from "@/assets/icons/lucide";
import { Suspense } from "@/components";
import { useCurrentUserImageByUserID, useUserImagesByUserID } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import { ROUTES } from "@/types/navigation";
import { buildImageUrl } from "@/utils/image";

import AccountInfoCard from "./components/AccountInfoCard";
import BasicInfoCard from "./components/BasicInfoCard";
import ProfileCard from "./components/ProfileCard";
import SkillsCard from "./components/SkillsCard";
import SocialLinksCard from "./components/SocialLinksCard";
import UserEducationsCard from "./components/UserEducationsCard";
import UserEmailsCard from "./components/UserEmailsCard";
import UserOccupationsCard from "./components/UserOccupationsCard";
import UserPhonesCard from "./components/UserPhonesCard";
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

        {/* Background Banner */}
        <Suspense loadingType="ecg" loadingSize="small">
          <BackgroundBanner userId={currentUserId} />
        </Suspense>


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
          <button
            className={styles.actionButton}
            onClick={() => navigate(ROUTES.IMAGES)}
          >
            <ImagePlus size={18} />
            <span>{t("manageImages", "Manage Images")}</span>
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

        {/* Other Images Gallery */}
        <Suspense loadingType="ecg" loadingSize="small">
          <ImagesGallery userId={currentUserId} />
        </Suspense>
      </div>
    </div>
  );
});

ProfilePage.displayName = "ProfilePage";

// 背景横幅组件
interface BackgroundBannerProps {
  userId: string;
}

const BackgroundBanner = memo(({ userId }: BackgroundBannerProps) => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();
  const { data: currentImage } = useCurrentUserImageByUserID({ userId });

  if (!currentImage?.imageId) {
    return (
      <div className={styles.bannerPlaceholder} onClick={() => navigate(ROUTES.IMAGES)}>
        <ImagePlus size={32} />
        <span>{t("addBackground", "Add a background image")}</span>
      </div>
    );
  }

  return (
    <div className={styles.banner} onClick={() => navigate(ROUTES.IMAGES)}>
      <img
        src={buildImageUrl(currentImage.imageId)}
        alt="Profile Background"
        className={styles.bannerImage}
      />
      <div className={styles.bannerOverlay}>
        <span>{t("changeBackground", "Click to manage images")}</span>
      </div>
    </div>
  );
});

BackgroundBanner.displayName = "BackgroundBanner";

// 图片画廊组件
interface ImagesGalleryProps {
  userId: string;
}

const ImagesGallery = memo(({ userId }: ImagesGalleryProps) => {
  const { t } = useTranslation("ProfilePage");
  const navigate = useNavigate();
  const { data: userImages = [] } = useUserImagesByUserID({ userId });

  // 跳过第一张（primary/background）
  const otherImages = userImages.slice(1);

  if (otherImages.length === 0) {
    return null;
  }

  return (
    <div className={styles.gallerySection}>
      <div className={styles.galleryHeader}>
        <h3 className={styles.galleryTitle}>{t("otherImages", "Other Images")}</h3>
        <button className={styles.viewAllButton} onClick={() => navigate(ROUTES.IMAGES)}>
          {t("viewAll", "View All")}
        </button>
      </div>
      <div className={styles.galleryGrid}>
        {otherImages.slice(0, 6).map((userImage, index) => (
          <div key={userImage.id} className={styles.galleryItem}>
            <img
              src={buildImageUrl(userImage.imageId)}
              alt={`Image ${index + 2}`}
              className={styles.galleryImage}
            />
          </div>
        ))}
      </div>
      {otherImages.length > 6 && (
        <p className={styles.moreImages}>
          {t("moreImages", "+ {{count}} more images", { count: otherImages.length - 6 })}
        </p>
      )}
    </div>
  );
});

ImagesGallery.displayName = "ImagesGallery";

export default ProfilePage;
