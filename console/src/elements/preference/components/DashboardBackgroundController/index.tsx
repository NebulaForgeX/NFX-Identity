import type { PreferenceFormValues } from "../../schemas/preferenceSchema";
import type { DashboardBackgroundType } from "@/types";

import { memo, useMemo } from "react";
import { Controller, useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { Dropdown } from "@/components";
import { DASHBOARD_BACKGROUND_VALUES } from "@/types";

export interface DashboardBackgroundControllerProps {
  /** 选择即改：选择新背景后立即回调，用于保存 */
  onApply?: (payload: { dashboardBackground: DashboardBackgroundType }) => void;
}

const DashboardBackgroundController = memo(({ onApply }: DashboardBackgroundControllerProps) => {
  const { t } = useTranslation("elements.directory");
  const {
    control,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  const backgroundOptions = useMemo(() => {
    return DASHBOARD_BACKGROUND_VALUES.map((value) => ({
      value,
      label: t(`preference.dashboardBackground.${value}`),
    }));
  }, [t]);

  return (
    <div style={{ marginBottom: "1rem" }}>
      <label
        style={{
          display: "block",
          marginBottom: "0.5rem",
          fontSize: "0.875rem",
          fontWeight: 500,
          color: "var(--color-fg-text)",
        }}
      >
        {t("preference.dashboardBackground.label")}
      </label>
      <Controller
        control={control}
        name="dashboardBackground"
        render={({ field }) => (
          <Dropdown
            options={backgroundOptions}
            value={field.value || ""}
            onChange={(value) => {
              field.onChange(value);
              onApply?.({ dashboardBackground: value as DashboardBackgroundType });
            }}
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
