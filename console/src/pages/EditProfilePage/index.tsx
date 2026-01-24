import { memo } from "react";
import { useTranslation } from "react-i18next";
import { FormProvider } from "react-hook-form";
import { useNavigate } from "react-router-dom";

import { Button, IconButton, Suspense } from "@/components";
import { ArrowLeft as ArrowLeftIcon } from "@/assets/icons/lucide";
import { useInitProfileForm, useSubmitProfile } from "@/elements/directory";
import {
  FirstNameController,
  LastNameController,
  NicknameController,
  DisplayNameController,
  BioController,
  BirthdayController,
  AgeController,
  GenderController,
  LocationController,
  WebsiteController,
  GithubController,
  SocialLinksController,
  SkillsController,
} from "@/elements/directory";
import { useUserProfile } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import { ROUTES } from "@/types/navigation";

import styles from "./styles.module.css";

const EditProfilePage = memo(() => {
  const { t } = useTranslation("EditProfilePage");
  const navigate = useNavigate();
  const currentUserId = useAuthStore((state) => state.currentUserId);

  if (!currentUserId) {
    navigate(ROUTES.PROFILE);
    return null;
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <IconButton
          variant="ghost"
          leftIcon={<ArrowLeftIcon size={20} />}
          onClick={() => navigate(ROUTES.PROFILE)}
          className={styles.backButton}
        >
          {t("back")}
        </IconButton>
        <h1 className={styles.title}>{t("title")}</h1>
        <p className={styles.subtitle}>{t("subtitle")}</p>
      </div>

      <Suspense
        loadingType="ecg"
        loadingText={t("loading")}
        loadingSize="small"
        loadingContainerClassName={styles.loading}
      >
        <EditProfileContent userId={currentUserId} />
      </Suspense>
    </div>
  );
});

EditProfilePage.displayName = "EditProfilePage";

const EditProfileContent = memo(({ userId }: { userId: string }) => {
  const { t } = useTranslation("EditProfilePage");
  const navigate = useNavigate();
  const { data: profile } = useUserProfile({ id: userId });
  const form = useInitProfileForm(profile);
  const profileId = profile?.id || userId;
  const { onSubmit, onSubmitError, isPending } = useSubmitProfile(profileId);

  return (
    <FormProvider {...form}>
      <form onSubmit={form.handleSubmit(onSubmit, onSubmitError)} className={styles.form}>
        <div className={styles.formGrid}>
          <FirstNameController />
          <LastNameController />
        </div>
        <NicknameController />
        <DisplayNameController />
        <BioController />
        <div className={styles.formGrid}>
          <BirthdayController />
          <AgeController />
        </div>
        <GenderController />
        <LocationController />
        <div className={styles.formGrid}>
          <WebsiteController />
          <GithubController />
        </div>
        <SocialLinksController />
        <SkillsController />

        <div className={styles.actions}>
          <Button
            type="button"
            variant="secondary"
            onClick={() => navigate(ROUTES.PROFILE)}
            disabled={isPending}
          >
            {t("cancel")}
          </Button>
          <Button type="submit" variant="primary" disabled={isPending}>
            {isPending ? t("submitting") : t("submit")}
          </Button>
        </div>
      </form>
    </FormProvider>
  );
});

EditProfileContent.displayName = "EditProfileContent";

export default EditProfilePage;
