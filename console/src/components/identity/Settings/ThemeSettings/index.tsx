import { Check, RefreshCw, Save } from "lucide-react";
import { useEffect, useState } from "react";
import { Badge, Box, Button, Container, Flex, Grid, Heading, RadioCards, Section, SegmentedControl, Switch, Text, TextField, Theme } from "@radix-ui/themes";
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

import { LedgerSection } from "../../Ledger";
import styles from "./s.module.css";

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

function swatchBackground(color: AccentColorEnum | GrayColorEnum): string {
  if (color === "auto" || color === "gray") return "var(--gray-9)";
  return `var(--${color}-9)`;
}

function swatchInk(color: AccentColorEnum | GrayColorEnum): string {
  if (color === "auto" || color === "gray") return "var(--gray-contrast)";
  return `var(--${color}-contrast)`;
}

function toDraft(pref: ResolvedThemePreference): ResolvedThemePreference {
  return { ...pref };
}

export default function ThemeSettings() {
  const { t } = useTranslation("pages.User.Setting");
  const { t: tHeader } = useTranslation("language");
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
            <Button type="button" variant="outline" color="gray" size="2" onClick={() => setDraft(toDraft(themePreference))} disabled={!dirty || saving}>
              <LucideIcon icon={RefreshCw} size={14} />
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
              <LucideIcon icon={Save} size={14} />
              {t("actions.saveTheme")}
            </Button>
          </Flex>
        }
      >
        <Grid columns={{ initial: "1", md: "2" }} gap="6" width="100%">
          <Flex direction="column" gap="6">
            <Flex direction="column" gap="4">
              <Heading as="h3" size="2">
                {t("labels.appearance")}
              </Heading>
              <Grid columns={{ initial: "1", sm: "2" }} gap="4">
                <Flex direction="column" gap="2">
                  <Text size="1" color="gray">{t("labels.appearance")}</Text>
                  <Box className={styles.controlFit}>
                    <SegmentedControl.Root size="2" value={draft.appearance} onValueChange={(v) => setField({ appearance: v as AppearanceEnum })}>
                      {THEME_APPEARANCE_VALUES.map((v) => (
                        <SegmentedControl.Item key={v} value={v}>{tHeader(`header.appearanceMode.${v}`)}</SegmentedControl.Item>
                      ))}
                    </SegmentedControl.Root>
                  </Box>
                </Flex>
                <Flex direction="column" gap="2">
                  <Text size="1" color="gray">{t("labels.radius")}</Text>
                  <Box className={styles.controlFit}>
                    <SegmentedControl.Root size="2" value={draft.radius} onValueChange={(v) => setField({ radius: v as RadiusEnum })}>
                      {RADIX_RADIUS_VALUES.map((v) => (
                        <SegmentedControl.Item key={v} value={v}>{getBaseDisplayName(RADIX_RADIUS_TO_BASE[v])}</SegmentedControl.Item>
                      ))}
                    </SegmentedControl.Root>
                  </Box>
                </Flex>
                <Flex direction="column" gap="2">
                  <Text size="1" color="gray">{t("labels.scaling")}</Text>
                  <Box className={styles.controlFit}>
                    <SegmentedControl.Root size="2" value={draft.scaling} onValueChange={(v) => setField({ scaling: v as ScalingEnum })}>
                      {RADIX_SCALING_VALUES.map((v) => (
                        <SegmentedControl.Item key={v} value={v}>{v}</SegmentedControl.Item>
                      ))}
                    </SegmentedControl.Root>
                  </Box>
                </Flex>
                <Flex direction="column" gap="2">
                  <Text size="1" color="gray">{t("labels.panelBackground")}</Text>
                  <Box className={styles.controlFit}>
                    <SegmentedControl.Root size="2" value={draft.panelBackground} onValueChange={(v) => setField({ panelBackground: v as PanelBackgroundEnum })}>
                      {RADIX_PANEL_BACKGROUND_VALUES.map((v) => (
                        <SegmentedControl.Item key={v} value={v}>{tHeader(`header.panelBackgroundMode.${v}`)}</SegmentedControl.Item>
                      ))}
                    </SegmentedControl.Root>
                  </Box>
                </Flex>
              </Grid>
            </Flex>
            <Flex direction="column" gap="4">
              <Heading as="h3" size="2">
                {t("labels.themeColor")}
              </Heading>
              <Flex direction="column" gap="3">
                <Text size="1" color="gray">{t("labels.accent")}</Text>
                <Flex wrap="wrap" gap="2">
                  {RADIX_ACCENT_VALUES.map((c) => {
                    const active = draft.accent === c;
                    return (
                      <Button
                        key={c}
                        type="button"
                        variant="solid"
                        color={c}
                        aria-label={c}
                        aria-pressed={active}
                        onClick={() => setField({ accent: c })}
                        className={active ? `${styles.swatch} ${styles.swatchActive}` : styles.swatch}
                      >
                        {active ? <LucideIcon icon={Check} size={12} /> : null}
                      </Button>
                    );
                  })}
                </Flex>
                <Text size="1" color="gray">{t("labels.gray")}</Text>
                <Flex wrap="wrap" gap="2">
                  {RADIX_GRAY_VALUES.map((c) => {
                    const active = draft.gray === c;
                    return (
                      <Button
                        key={c}
                        type="button"
                        variant="solid"
                        color="gray"
                        aria-label={c}
                        aria-pressed={active}
                        onClick={() => setField({ gray: c })}
                        className={active ? `${styles.swatch} ${styles.swatchGray} ${styles.swatchActive}` : `${styles.swatch} ${styles.swatchGray}`}
                        style={c === "auto" || c === "gray" ? undefined : { background: swatchBackground(c), color: swatchInk(c) }}
                      >
                        {active ? <LucideIcon icon={Check} size={12} /> : null}
                      </Button>
                    );
                  })}
                </Flex>
              </Flex>
            </Flex>
            <Flex direction="column" gap="4">
              <Heading as="h3" size="2">
                {t("labels.font")}
              </Heading>
              <Box className={styles.controlFit}>
                <SegmentedControl.Root size="2" value={draft.fontFamily} onValueChange={(v) => setField({ fontFamily: v as ThemeFontFamilyEnum })}>
                  {THEME_FONT_FAMILY_VALUES.map((v) => (
                    <SegmentedControl.Item key={v} value={v}>{t(FONT_LABEL_KEY[v])}</SegmentedControl.Item>
                  ))}
                </SegmentedControl.Root>
              </Box>
            </Flex>
          </Flex>
          <Flex direction="column" gap="2">
            <Text size="1" color="gray">{t("labels.livePreview")}</Text>
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
              <Section size="1" py="3" className={styles.previewFrame}>
                <Container size="4" width="100%" px="3">
                  <Flex direction="column" gap="3">
                    <Flex align="center" justify="between">
                      <Heading as="h3" size="4">{APP_NAME}</Heading>
                      <Badge size="1">{t("labels.previewBadge")}</Badge>
                    </Flex>
                    <Flex gap="2" wrap="wrap">
                      <Button size="2">{t("labels.previewSolid")}</Button>
                      <Button size="2" variant="outline">{t("labels.previewSoft")}</Button>
                    </Flex>
                    <TextField.Root size="2" placeholder={t("labels.sampleInput")} />
                    <Flex align="center" gap="2">
                      <Switch size="2" defaultChecked />
                      <Text size="2">{t("labels.notifications")}</Text>
                    </Flex>
                  </Flex>
                </Container>
              </Section>
            </Theme>
          </Flex>
        </Grid>
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
