import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Controller, useFormContext } from "react-hook-form";

import { Dropdown } from "@/components";
import { useTheme } from "@/providers/ThemeProvider/useTheme";

const ThemeController = memo(() => {
  const { t } = useTranslation("elements.directory");
  const { t: tComponents } = useTranslation("components");
  const { availableThemes } = useTheme();
  const {
    control,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  const themeOptions = useMemo(() => {
    return availableThemes.map((theme) => ({
      value: theme,
      label: tComponents(`themeSwitcher.${theme}`, { defaultValue: theme }),
    }));
  }, [availableThemes, tComponents]);

  return (
    <div style={{ marginBottom: "1rem" }}>
      <label style={{ display: "block", marginBottom: "0.5rem", fontSize: "0.875rem", fontWeight: 500, color: "var(--color-fg-text)" }}>
        {t("preference.theme.label")}
      </label>
      <Controller
        control={control}
        name="theme"
        render={({ field }) => (
          <Dropdown
            options={themeOptions}
            value={field.value || ""}
            onChange={field.onChange}
            placeholder={t("preference.theme.placeholder")}
            error={!!errors.theme}
          />
        )}
      />
      {errors.theme && (
        <p style={{ fontSize: "0.75rem", color: "var(--color-danger)", margin: "0.25rem 0 0 0" }}>
          {errors.theme.message as string}
        </p>
      )}
    </div>
  );
});

ThemeController.displayName = "ThemeController";

export default ThemeController;
