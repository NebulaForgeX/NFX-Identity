import { useCallback } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import type { FieldErrors } from "react-hook-form";

import { useUpdateUserProfile } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import { ROUTES } from "@/types/navigation";
import { showError } from "@/stores/modalStore";

import { type ProfileFormValues } from "../schemas/profileSchema";

export const useSubmitProfile = (profileId?: string) => {
  const { t } = useTranslation("elements.directory");
  const navigate = useNavigate();
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const updateProfile = useUpdateUserProfile();

  const onSubmit = useCallback(
    async (values: ProfileFormValues) => {
      const id = profileId || currentUserId;
      if (!id) {
        showError(t("profile.messages.user_not_found"));
        return;
      }

      try {
        // Convert key-value pairs to map
        const socialLinks =
          values.socialLinks && values.socialLinks.length > 0
            ? values.socialLinks.reduce((acc, pair) => {
                if (pair.key.trim()) {
                  acc[pair.key] = pair.value;
                }
                return acc;
              }, {} as Record<string, unknown>)
            : undefined;

        const skills =
          values.skills && values.skills.length > 0
            ? values.skills.reduce((acc, skill) => {
                if (skill.name && skill.name.trim()) {
                  acc[skill.name] = skill.score;
                }
                return acc;
              }, {} as Record<string, unknown>)
            : undefined;

        await updateProfile.mutateAsync({
          id,
          data: {
            role: values.role || undefined,
            firstName: values.firstName || undefined,
            lastName: values.lastName || undefined,
            nickname: values.nickname || undefined,
            displayName: values.displayName || undefined,
            bio: values.bio || undefined,
            birthday: values.birthday || undefined,
            age: values.age || undefined,
            gender: values.gender || undefined,
            location: values.location || undefined,
            website: values.website || undefined,
            github: values.github || undefined,
            socialLinks,
            skills,
          },
        });
        navigate(ROUTES.PROFILE);
      } catch (error) {
        console.error("Failed to update profile:", error);
      }
    },
    [profileId, currentUserId, updateProfile, navigate, t],
  );

  const onSubmitError = useCallback((errors: FieldErrors<ProfileFormValues>) => {
    console.error("Form validation errors:", errors);
  }, []);

  return {
    onSubmit,
    onSubmitError,
    isPending: updateProfile.isPending,
  };
};
