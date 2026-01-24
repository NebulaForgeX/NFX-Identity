import { useCallback } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import type { FieldErrors } from "react-hook-form";

import { useCreateUserOccupation, useUpdateUserOccupation } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";
import { ROUTES } from "@/types/navigation";
import { showError } from "@/stores/modalStore";
import type { UserOccupation } from "@/types";

import { type OccupationFormValues } from "../schemas/occupationSchema";

export const useSubmitOccupation = (occupation?: UserOccupation) => {
  const { t } = useTranslation("elements.directory");
  const navigate = useNavigate();
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const createOccupation = useCreateUserOccupation();
  const updateOccupation = useUpdateUserOccupation();

  const onSubmit = useCallback(
    async (values: OccupationFormValues) => {
      if (!currentUserId) {
        showError(t("occupation.messages.user_not_found"));
        return;
      }

      try {
        if (occupation) {
          // 更新现有工作
          await updateOccupation.mutateAsync({
            id: occupation.id,
            data: {
              company: values.company,
              position: values.position,
              department: values.department,
              industry: values.industry,
              location: values.location,
              employmentType: values.employmentType,
              startDate: values.startDate,
              endDate: values.endDate,
              isCurrent: values.isCurrent,
              description: values.description,
              responsibilities: values.responsibilities,
              achievements: values.achievements,
              skillsUsed: values.skillsUsed,
            },
          });
        } else {
          // 创建新工作
          await createOccupation.mutateAsync({
            userId: currentUserId,
            company: values.company,
            position: values.position,
            department: values.department,
            industry: values.industry,
            location: values.location,
            employmentType: values.employmentType,
            startDate: values.startDate,
            endDate: values.endDate,
            isCurrent: values.isCurrent,
            description: values.description,
            responsibilities: values.responsibilities,
            achievements: values.achievements,
            skillsUsed: values.skillsUsed,
          });
        }
        navigate(ROUTES.PROFILE);
      } catch (error) {
        console.error("Failed to save occupation:", error);
      }
    },
    [currentUserId, occupation, createOccupation, updateOccupation, navigate, t],
  );

  const onSubmitError = useCallback((errors: FieldErrors<OccupationFormValues>) => {
    console.error("Form validation errors:", errors);
  }, []);

  return {
    onSubmit,
    onSubmitError,
    isPending: createOccupation.isPending || updateOccupation.isPending,
  };
};
