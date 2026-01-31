import type { PreferenceFormValues } from "../../schemas/preferenceSchema";
import type { Language } from "@/assets/languages/i18nResources";

import { memo, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Controller, useFormContext } from "react-hook-form";

import { Dropdown } from "@/components";
import { LANGUAGE } from "@/assets/languages/i18nResources";

export interface LanguageControllerProps {
  /** 选择即改：选择新语言后立即回调，用于保存并应用 */
  onApply?: (payload: { language: string }) => void;
}

const LanguageController = memo(({ onApply }: LanguageControllerProps) => {
  const { t } = useTranslation("elements.directory");
  const { t: tComponents } = useTranslation("components");
  const {
    control,
    formState: { errors },
  } = useFormContext<PreferenceFormValues>();

  const languageOptions = useMemo(() => {
    const availableLanguages: Language[] = [LANGUAGE.EN, LANGUAGE.ZH, LANGUAGE.FR];
    const keyMap: Record<Language, string> = {
      [LANGUAGE.EN]: "languageSwitcher.english",
      [LANGUAGE.ZH]: "languageSwitcher.chinese",
      [LANGUAGE.FR]: "languageSwitcher.french",
    };
    return availableLanguages.map((lang) => ({
      value: lang,
      label: tComponents(keyMap[lang], { defaultValue: lang }),
    }));
  }, [tComponents]);

  return (
    <div style={{ marginBottom: "1rem" }}>
      <label style={{ display: "block", marginBottom: "0.5rem", fontSize: "0.875rem", fontWeight: 500, color: "var(--color-fg-text)" }}>
        {t("preference.language.label")}
      </label>
      <Controller
        control={control}
        name="language"
        render={({ field }) => (
          <Dropdown
            options={languageOptions}
            value={field.value || ""}
            onChange={(value) => {
              field.onChange(value);
              onApply?.({ language: value });
            }}
            placeholder={t("preference.language.placeholder")}
            error={!!errors.language}
          />
        )}
      />
      {errors.language && (
        <p style={{ fontSize: "0.75rem", color: "var(--color-danger)", margin: "0.25rem 0 0 0" }}>
          {errors.language.message as string}
        </p>
      )}
    </div>
  );
});

LanguageController.displayName = "LanguageController";

export default LanguageController;
