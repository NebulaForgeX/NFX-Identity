/**
 * User preference sync - 与 Sjgz-Admin 一致：
 * 使用 nfx-ui parsePreferenceJson / Preference / ThemeEnum / LayoutModeEnum，无 as。
 */

import { useCallback, useEffect, useRef } from "react";

import { changeLanguage, LanguageEnum } from "nfx-ui/languages";
import { BaseEnum, ThemeEnum, useTheme } from "nfx-ui/themes";
import { useLayout, LayoutModeEnum } from "nfx-ui/layouts";
import { parsePreferenceJson } from "nfx-ui/preference";

import type { UserPreference } from "@/types";
import { useUpdateUserPreference, useUserPreferenceNormal } from "@/hooks/useDirectory";
import { useAuthStore } from "@/stores/authStore";

/** 将 API 返回的 UserPreference 转为 nfx-ui Preference（用于类型安全的 theme/layoutMode） */
function userPreferenceToPreference(raw: UserPreference | undefined): ReturnType<typeof parsePreferenceJson> {
  if (!raw) return null;
  const backendShape = {
    theme: raw.theme,
    base: raw.base,
    language: raw.language,
    layoutMode: raw.display?.layoutMode ?? "show",
    other: raw.other,
  };
  return parsePreferenceJson(JSON.stringify(backendShape));
}

export const useThemeSync = () => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const { themeName, baseName, setTheme, setBase } = useTheme();
  const updatePreference = useUpdateUserPreference({ silent: true });

  const shouldFetch = !!currentUserId && isAuthValid;

  const { data: rawPreference } = useUserPreferenceNormal({
    id: currentUserId,
    options: { enabled: shouldFetch && !!currentUserId },
  });
  const preference = userPreferenceToPreference(rawPreference);

  const lastPreferenceId = useRef<string | null>(null);
  const isInitialized = useRef(false);

  useEffect(() => {
    if (!shouldFetch || !preference || !currentUserId || !rawPreference) {
      isInitialized.current = false;
      lastPreferenceId.current = null;
      return;
    }
    if (lastPreferenceId.current === rawPreference.id && isInitialized.current) return;

    if (preference.theme && preference.theme !== themeName) {
      setTheme(preference.theme);
    }
    if (preference.base && preference.base !== baseName) {
      setBase(preference.base);
    }

    isInitialized.current = true;
    lastPreferenceId.current = rawPreference.id;
  }, [preference, rawPreference, currentUserId, themeName, baseName, setTheme, setBase, shouldFetch]);

  const syncTheme = useCallback(
    async (theme: ThemeEnum) => {
      if (!currentUserId || !isAuthValid) return;
      try {
        await updatePreference.mutateAsync({
          id: currentUserId,
          data: { theme },
        });
      } catch (error) {
        console.error("Failed to sync theme preference:", error);
      }
    },
    [currentUserId, isAuthValid, updatePreference],
  );

  const syncBase = useCallback(
    async (base: BaseEnum) => {
      if (!currentUserId || !isAuthValid) return;
      try {
        await updatePreference.mutateAsync({
          id: currentUserId,
          data: { base },
        });
      } catch (error) {
        console.error("Failed to sync base preference:", error);
      }
    },
    [currentUserId, isAuthValid, updatePreference],
  );

  return { syncTheme, syncBase };
};

export const useLanguageSync = () => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const updatePreference = useUpdateUserPreference({ silent: true });

  const shouldFetch = !!currentUserId && isAuthValid;

  const { data: rawPreference } = useUserPreferenceNormal({
    id: currentUserId,
    options: { enabled: shouldFetch && !!currentUserId },
  });
  const preference = userPreferenceToPreference(rawPreference);

  const lastPreferenceId = useRef<string | null>(null);
  const isInitialized = useRef(false);

  useEffect(() => {
    if (!shouldFetch || !preference || !currentUserId || !rawPreference) {
      isInitialized.current = false;
      lastPreferenceId.current = null;
      return;
    }
    if (lastPreferenceId.current === rawPreference.id && isInitialized.current) return;

    if (preference.language) {
      const currentLang = localStorage.getItem("i18nextLng") || LanguageEnum.ZH;
      if (currentLang !== preference.language) {
        changeLanguage(preference.language);
      }
    }

    isInitialized.current = true;
    lastPreferenceId.current = rawPreference.id;
  }, [preference, rawPreference, currentUserId, shouldFetch]);

  const syncLanguage = useCallback(
    async (language: LanguageEnum) => {
      if (!currentUserId || !isAuthValid) return;
      try {
        await updatePreference.mutateAsync({
          id: currentUserId,
          data: { language },
        });
      } catch (error) {
        console.error("Failed to sync language preference:", error);
      }
    },
    [currentUserId, isAuthValid, updatePreference],
  );

  return { syncLanguage };
};

export const useLayoutSync = () => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const { layoutMode, setLayoutMode } = useLayout();
  const updatePreference = useUpdateUserPreference({ silent: true });

  const shouldFetch = !!currentUserId && isAuthValid;

  const { data: rawPreference } = useUserPreferenceNormal({
    id: currentUserId,
    options: { enabled: shouldFetch && !!currentUserId },
  });
  const preference = userPreferenceToPreference(rawPreference);

  const lastPreferenceId = useRef<string | null>(null);
  const isInitialized = useRef(false);

  useEffect(() => {
    if (!shouldFetch || !preference || !currentUserId || !rawPreference) {
      isInitialized.current = false;
      lastPreferenceId.current = null;
      return;
    }
    if (lastPreferenceId.current === rawPreference.id && isInitialized.current) return;

    if (
      preference.layoutMode === LayoutModeEnum.SHOW ||
      preference.layoutMode === LayoutModeEnum.HIDE
    ) {
      if (!isInitialized.current && preference.layoutMode !== layoutMode) {
        setLayoutMode(preference.layoutMode);
      }
    }

    isInitialized.current = true;
    lastPreferenceId.current = rawPreference.id;
  }, [preference, rawPreference, currentUserId, layoutMode, setLayoutMode, shouldFetch]);

  const syncLayout = useCallback(
    async (mode: LayoutModeEnum) => {
      if (!currentUserId || !isAuthValid) return;
      const currentDisplay = (rawPreference?.display as Record<string, unknown> | undefined) ?? {};
      updatePreference
        .mutateAsync({
          id: currentUserId,
          data: { display: { ...currentDisplay, layoutMode: mode } },
        })
        .catch((error) => {
          console.error("Failed to sync layout preference (non-critical):", error);
        });
    },
    [currentUserId, isAuthValid, updatePreference, rawPreference?.display],
  );

  return { syncLayout };
};
