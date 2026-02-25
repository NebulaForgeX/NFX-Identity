import type { PreferenceFormValues } from "../../schemas/preferenceSchema";

import { memo, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Controller, useFormContext } from "react-hook-form";

import { Dropdown } from "@/components";
import { useTheme } from "@/providers/ThemeProvider/useTheme";

export interface BaseControllerProps {
  /** 选择即改：选择新 base 后立即回调，用于保存并应用 */
  onApply?: (payload: { base: string }) => void;
}

const BaseController = memo(({ onApply }: BaseControllerProps) => {
  const { t } = useTranslation("elements.directory");
  const { t: tComponents } = useTranslation("components");
  const { availableBases } = useTheme();
  const {
    control,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  const baseOptions = useMemo(() => {
    return availableBases.map((base) => ({
      value: base,
      label: tComponents(`baseSwitcher.${base}`, { defaultValue: base }),
    }));
  }, [availableBases, tComponents]);

  return (
    <div style={{ marginBottom: "1rem" }}>
      <label style={{ display: "block", marginBottom: "0.5rem", fontSize: "0.875rem", fontWeight: 500, color: "var(--color-fg-text)" }}>
        {t("preference.base.label")}
      </label>
      <Controller
        control={control}
        name="base"
        render={({ field }) => (
          <Dropdown
            options={baseOptions}
            value={field.value || ""}
            onChange={(value) => {
              field.onChange(value);
              onApply?.({ base: value });
            }}
            placeholder={t("preference.base.placeholder")}
            error={!!errors.base}
          />
        )}
      />
      {errors.base && (
        <p style={{ fontSize: "0.75rem", color: "var(--color-danger)", margin: "0.25rem 0 0 0" }}>
          {errors.base.message as string}
        </p>
      )}
    </div>
  );
});

BaseController.displayName = "BaseController";

export default BaseController;
