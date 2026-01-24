import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Controller, useFormContext } from "react-hook-form";

import { Dropdown } from "@/components";

export type DashboardBackgroundType = "waves" | "squares" | "letterGlitch" | "none";

const DashboardBackgroundController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    control,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  const backgroundOptions = useMemo(() => {
    return [
      { value: "waves", label: t("preference.dashboardBackground.waves") },
      { value: "squares", label: t("preference.dashboardBackground.squares") },
      { value: "letterGlitch", label: t("preference.dashboardBackground.letterGlitch") },
      { value: "none", label: t("preference.dashboardBackground.none") },
    ];
  }, [t]);

  return (
    <div style={{ marginBottom: "1rem" }}>
      <label style={{ display: "block", marginBottom: "0.5rem", fontSize: "0.875rem", fontWeight: 500, color: "var(--color-fg-text)" }}>
        {t("preference.dashboardBackground.label")}
      </label>
      <Controller
        control={control}
        name="dashboardBackground"
        render={({ field }) => (
          <Dropdown
            options={backgroundOptions}
            value={field.value || ""}
            onChange={field.onChange}
            placeholder={t("preference.dashboardBackground.placeholder")}
            error={!!errors.dashboardBackground}
          />
        )}
      />
      {errors.dashboardBackground && (
        <p style={{ fontSize: "0.75rem", color: "var(--color-danger)", margin: "0.25rem 0 0 0" }}>
          {errors.dashboardBackground.message as string}
        </p>
      )}
    </div>
  );
});

DashboardBackgroundController.displayName = "DashboardBackgroundController";

export default DashboardBackgroundController;
