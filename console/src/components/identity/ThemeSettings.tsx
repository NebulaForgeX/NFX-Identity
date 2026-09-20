import { CheckedIcon, RefreshIcon, SaveIcon } from "nfx-ui/icons";
import { useEffect, useState } from "react";
import { Badge, Box, Button, Flex, Heading, RadioCards, SegmentedControl, Switch, Text, TextField, Theme } from "@radix-ui/themes";
import { APP_NAME } from "nfx-ui/config";
import { AccentColorEnum, AppearanceEnum, GrayColorEnum, LanguageEnum, PanelBackgroundEnum, RadiusEnum, ScalingEnum, ThemeFontFamilyEnum } from "nfx-ui/enums";
import { useBaseLabel, useSyncPreference } from "nfx-ui/hooks";
import { usePreferenceStore } from "nfx-ui/stores";
import {
  RADIX_ACCENT_VALUES,
  RADIX_GRAY_VALUES,
  RADIX_PANEL_BACKGROUND_VALUES,
  RADIX_RADIUS_TO_BASE,
  RADIX_RADIUS_VALUES,
  RADIX_SCALING_VALUES,
  ResolvedThemePreference,
  resolveRadixAppearance,
  THEME_APPEARANCE_VALUES,
  THEME_FONT_FAMILY_VALUES,
} from "nfx-ui/themes";
import { useTranslation } from "react-i18next";

import { LucideIcon } from "@/components";

import { FieldList, FieldRow, LedgerSection } from "./Ledger";
import styles from "./settings.module.css";

const FONT_LABEL_KEY: Record<ThemeFontFamilyEnum, string> = {
  [ThemeFontFamilyEnum.SYSTEM]: "labels.fontSystem",
  [ThemeFontFamilyEnum.IBM_PLEX]: "labels.fontIbmPlex",
  [ThemeFontFamilyEnum.NOTO]: "labels.fontNoto",
  [ThemeFontFamilyEnum.SOURCE_SANS]: "labels.fontSourceSans",
};

const LANG_CODE: Record<LanguageEnum, string> = {
  [LanguageEnum.EN]: "EN",
  [LanguageEnum.ZH]: "ZH",
  [LanguageEnum.FR]: "FR",
};

const LANG_NAME_KEY: Record<LanguageEnum, string> = {
  [LanguageEnum.EN]: "labels.langEn",
  [LanguageEnum.ZH]: "labels.langZh",
  [LanguageEnum.FR]: "labels.langFr",
};

function swatchVar(color: AccentColorEnum | GrayColorEnum): string {
  return color === "auto" ? "var(--gray-9)" : `var(--${color}-9)`;
}

function toDraft(pref: ResolvedThemePreference): ResolvedThemePreference {
  return { ...pref };
}

export default function ThemeSettings() {
  const { t } = useTranslation("pages.User.Setting");
  const themePreference = usePreferenceStore((s) => s.theme);
  const currentLanguage = usePreferenceStore((s) => s.language);
  const { syncPreference } = useSyncPreference();
  const { getBaseDisplayName } = useBaseLabel();
  const [draft, setDraft] = useState(() => toDraft(themePreference));
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    setDraft(toDraft(themePreference));
  }, [themePreference]);

  const dirty =
    draft.accent !== themePreference.accent ||
    draft.gray !== themePreference.gray ||
    draft.appearance !== themePreference.appearance ||
    draft.radius !== themePreference.radius ||
    draft.scaling !== themePreference.scaling ||
    draft.panelBackground !== themePreference.panelBackground ||
    draft.fontFamily !== themePreference.fontFamily;

  const setField = (patch: Partial<ResolvedThemePreference>) => setDraft((prev) => ({ ...prev, ...patch }));
  const previewAppearance = resolveRadixAppearance(draft.appearance);

  return (
    <>
      <LedgerSection
        title={t("sections.theme.title")}
        description={t("sections.theme.description")}
        actions={
          <Flex gap="2">
            <Button type="button" variant="soft" color="gray" size="2" onClick={() => setDraft(toDraft(themePreference))} disabled={!dirty || saving}>
              <LucideIcon icon={RefreshIcon} size={14} />
              {t("actions.reset")}
            </Button>
            <Button
              type="button"
              size="2"
              disabled={!dirty || saving}
              onClick={() => {
                setSaving(true);
                try {
                  syncPreference({ theme: { ...draft } });
                } finally {
                  setSaving(false);
                }
              }}
            >
              <LucideIcon icon={SaveIcon} size={14} />
              {t("actions.saveTheme")}
            </Button>
          </Flex>
        }
      >
        <FieldList>
          <FieldRow label={t("labels.appearance")}>
            <Box className={styles.controlFit}>
              <SegmentedControl.Root size="2" value={draft.appearance} onValueChange={(v) => setField({ appearance: v as AppearanceEnum })}>
                {THEME_APPEARANCE_VALUES.map((v) => (
                  <SegmentedControl.Item key={v} value={v}>
                    {v}
                  </SegmentedControl.Item>
                ))}
              </SegmentedControl.Root>
            </Box>
          </FieldRow>
          <FieldRow label={t("labels.accent")}>
            <Flex wrap="wrap" gap="2">
              {RADIX_ACCENT_VALUES.map((c) => {
                const active = draft.accent === c;
                return (
                  <button
                    key={c}
                    type="button"
                    aria-label={c}
                    aria-pressed={active}
                    onClick={() => setField({ accent: c })}
                    className={`${styles.swatch} ${active ? styles.swatchActive : ""}`}
                    style={{ background: swatchVar(c) }}
                  >
                    {active ? <LucideIcon icon={CheckedIcon} size={12} color="white" /> : null}
                  </button>
                );
              })}
            </Flex>
          </FieldRow>
          <FieldRow label={t("labels.gray")}>
            <Flex wrap="wrap" gap="2">
              {RADIX_GRAY_VALUES.map((c) => {
                const active = draft.gray === c;
                return (
                  <button
                    key={c}
                    type="button"
                    aria-label={c}
                    aria-pressed={active}
                    onClick={() => setField({ gray: c })}
                    className={`${styles.swatch} ${active ? styles.swatchActive : ""}`}
                    style={{ background: swatchVar(c) }}
                  >
                    {active ? <LucideIcon icon={CheckedIcon} size={12} color="white" /> : null}
                  </button>
                );
              })}
            </Flex>
          </FieldRow>
          <FieldRow label={t("labels.radius")}>
            <Box className={styles.controlFit}>
              <SegmentedControl.Root size="2" value={draft.radius} onValueChange={(v) => setField({ radius: v as RadiusEnum })}>
                {RADIX_RADIUS_VALUES.map((v) => (
                  <SegmentedControl.Item key={v} value={v}>
                    {getBaseDisplayName(RADIX_RADIUS_TO_BASE[v])}
                  </SegmentedControl.Item>
                ))}
              </SegmentedControl.Root>
            </Box>
          </FieldRow>
          <FieldRow label={t("labels.scaling")}>
            <Box className={styles.controlFit}>
              <SegmentedControl.Root size="2" value={draft.scaling} onValueChange={(v) => setField({ scaling: v as ScalingEnum })}>
                {RADIX_SCALING_VALUES.map((v) => (
                  <SegmentedControl.Item key={v} value={v}>
                    {v}
                  </SegmentedControl.Item>
                ))}
              </SegmentedControl.Root>
            </Box>
          </FieldRow>
          <FieldRow label={t("labels.panelBackground")}>
            <Box className={styles.controlFit}>
              <SegmentedControl.Root size="2" value={draft.panelBackground} onValueChange={(v) => setField({ panelBackground: v as PanelBackgroundEnum })}>
                {RADIX_PANEL_BACKGROUND_VALUES.map((v) => (
                  <SegmentedControl.Item key={v} value={v}>
                    {v}
                  </SegmentedControl.Item>
                ))}
              </SegmentedControl.Root>
            </Box>
          </FieldRow>
          <FieldRow label={t("labels.font")}>
            <Box className={styles.controlFit}>
              <SegmentedControl.Root size="2" value={draft.fontFamily} onValueChange={(v) => setField({ fontFamily: v as ThemeFontFamilyEnum })}>
                {THEME_FONT_FAMILY_VALUES.map((v) => (
                  <SegmentedControl.Item key={v} value={v}>
                    {t(FONT_LABEL_KEY[v])}
                  </SegmentedControl.Item>
                ))}
              </SegmentedControl.Root>
            </Box>
          </FieldRow>
          <FieldRow label={t("labels.livePreview")}>
            <Theme
              appearance={previewAppearance}
              accentColor={draft.accent}
              grayColor={draft.gray}
              radius={draft.radius}
              scaling={draft.scaling}
              panelBackground={draft.panelBackground}
              hasBackground
              className={styles.previewTheme}
            >
              <Flex direction="column" gap="3" p="3" style={{ border: "1px solid var(--gray-a5)" }}>
                <Flex align="center" justify="between">
                  <Heading size="4">{APP_NAME}</Heading>
                  <Badge size="1">{t("labels.previewBadge")}</Badge>
                </Flex>
                <Flex gap="2" wrap="wrap">
                  <Button size="2">{t("labels.previewSolid")}</Button>
                  <Button size="2" variant="soft">
                    {t("labels.previewSoft")}
                  </Button>
                </Flex>
                <TextField.Root size="2" placeholder={t("labels.sampleInput")} />
                <Flex align="center" gap="2">
                  <Switch size="2" defaultChecked />
                  <Text size="2">{t("labels.notifications")}</Text>
                </Flex>
              </Flex>
            </Theme>
          </FieldRow>
        </FieldList>
      </LedgerSection>
      <LedgerSection title={t("sections.language.title")} description={t("sections.language.description")}>
        <RadioCards.Root size="1" columns="3" gap="2" value={currentLanguage} onValueChange={(v) => syncPreference({ language: v as LanguageEnum })}>
          {(Object.values(LanguageEnum) as LanguageEnum[]).map((lang) => (
            <RadioCards.Item key={lang} value={lang}>
              <Flex direction="column" align="center" gap="1" width="100%">
                <Text size="3" weight="bold">
                  {LANG_CODE[lang]}
                </Text>
                <Text size="1" color="gray">
                  {t(LANG_NAME_KEY[lang])}
                </Text>
              </Flex>
            </RadioCards.Item>
          ))}
        </RadioCards.Root>
      </LedgerSection>
    </>
  );
}
