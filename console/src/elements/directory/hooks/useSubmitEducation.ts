import { useCallback } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import type { FieldErrors } from "react-hook-form";

import { useCreateUserEducation, useUpdateUserEducation } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import { ROUTES } from "@/types/navigation";
import { showError } from "@/stores/modalStore";
import type { UserEducation } from "@/types";

import { type EducationFormValues } from "../schemas/educationSchema";

export const useSubmitEducation = (education?: UserEducation) => {
  const { t } = useTranslation("elements.directory");
  const navigate = useNavigate();
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const createEducation = useCreateUserEducation();
  const updateEducation = useUpdateUserEducation();

  const onSubmit = useCallback(
    async (values: EducationFormValues) => {
      if (!currentUserId) {
        showError(t("education.messages.user_not_found"));
        return;
      }

      try {
        if (education) {
          // 更新现有教育
          await updateEducation.mutateAsync({
            id: education.id,
            data: {
              school: values.school,
              degree: values.degree,
              major: values.major,
              fieldOfStudy: values.fieldOfStudy,
              startDate: values.startDate,
              endDate: values.endDate,
              isCurrent: values.isCurrent,
              description: values.description,
              grade: values.grade,
              activities: values.activities,
              achievements: values.achievements,
            },
          });
        } else {
          // 创建新教育
          await createEducation.mutateAsync({
            userId: currentUserId,
            school: values.school,
            degree: values.degree,
            major: values.major,
            fieldOfStudy: values.fieldOfStudy,
            startDate: values.startDate,
            endDate: values.endDate,
            isCurrent: values.isCurrent,
            description: values.description,
            grade: values.grade,
            activities: values.activities,
            achievements: values.achievements,
          });
        }
        navigate(ROUTES.PROFILE);
      } catch (error) {
        console.error("Failed to save education:", error);
      }
    },
    [currentUserId, education, createEducation, updateEducation, navigate, t],
  );

  const onSubmitError = useCallback((errors: FieldErrors<EducationFormValues>) => {
    console.error("Form validation errors:", errors);
  }, []);

  return {
    onSubmit,
    onSubmitError,
    isPending: createEducation.isPending || updateEducation.isPending,
  };
};
