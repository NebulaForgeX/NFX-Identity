import type { BootstrapFormValues } from "../schemas/bootstrapSchema";
import type { FieldErrors } from "react-hook-form";

import { useCallback } from "react";
import { useTranslation } from "react-i18next";
import { useMutation } from "@tanstack/react-query";

import { InitializeSystemState } from "@/apis/system.api";
import { hideLoading, showError, showLoading, showSuccess } from "@/stores/modalStore";

export const useSubmitBootstrap = () => {
  const { t } = useTranslation("elements.bootstrap");

  const { mutateAsync, isPending } = useMutation({
    mutationFn: async (values: BootstrapFormValues) => {
      const params = {
        version: values.Version,
        admin_username: values.AdminUsername.trim(),
        admin_password: values.AdminPassword,
        admin_email: values.AdminEmail?.trim() || undefined,
        admin_phone: values.AdminPhone?.trim() || undefined,
        admin_country_code: values.AdminCountryCode?.trim() || undefined,
      };

      return await InitializeSystemState(params);
    },
    onMutate: () => {
      showLoading({ message: t("messages.initializing"), canClose: false });
    },
    onSuccess: () => {
      hideLoading();
      showSuccess({
        title: t("messages.init_success_title"),
        message: t("messages.init_success_message"),
        onClick: () => {
          window.location.reload();
        },
      });
    },
    onError: (error: Error) => {
      hideLoading();
      showError(error.message || t("messages.init_failed_unknown"), t("messages.init_failed_title"));
    },
  });

  const onSubmit = useCallback(
    async (values: BootstrapFormValues) => {
      await mutateAsync(values);
    },
    [mutateAsync],
  );

  const onSubmitError = useCallback(
    (errors: FieldErrors<BootstrapFormValues>) => {
      console.error("Form validation errors:", errors);
      const firstError = Object.values(errors)[0];
      if (firstError?.message) {
        showError(firstError.message as string, t("validation.form_validation_failed"));
      }
    },
    [t],
  );

  return {
    onSubmit,
    onSubmitError,
    isPending,
  };
};
