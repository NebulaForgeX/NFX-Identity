import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Controller, useFormContext } from "react-hook-form";

import { Dropdown } from "@/components";

// 常用时区列表
const TIMEZONES = [
  "UTC",
  "America/New_York",
  "America/Chicago",
  "America/Denver",
  "America/Los_Angeles",
  "Europe/London",
  "Europe/Paris",
  "Europe/Berlin",
  "Asia/Shanghai",
  "Asia/Tokyo",
  "Asia/Hong_Kong",
  "Asia/Singapore",
  "Australia/Sydney",
  "America/Toronto",
  "America/Mexico_City",
  "America/Sao_Paulo",
];

const TimezoneController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const {
    control,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  const timezoneOptions = useMemo(() => {
    return TIMEZONES.map((tz) => ({
      value: tz,
      label: tz,
    }));
  }, []);

  return (
    <div style={{ marginBottom: "1rem" }}>
      <label style={{ display: "block", marginBottom: "0.5rem", fontSize: "0.875rem", fontWeight: 500, color: "var(--color-fg-text)" }}>
        {t("preference.timezone.label")}
      </label>
      <Controller
        control={control}
        name="timezone"
        render={({ field }) => (
          <Dropdown
            options={timezoneOptions}
            value={field.value || ""}
            onChange={field.onChange}
            placeholder={t("preference.timezone.placeholder")}
            error={!!errors.timezone}
          />
        )}
      />
      {errors.timezone && (
        <p style={{ fontSize: "0.75rem", color: "var(--color-danger)", margin: "0.25rem 0 0 0" }}>
          {errors.timezone.message as string}
        </p>
      )}
    </div>
  );
});

TimezoneController.displayName = "TimezoneController";

export default TimezoneController;
