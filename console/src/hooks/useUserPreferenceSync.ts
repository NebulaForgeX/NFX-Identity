import { useEffect, useCallback, useRef } from "react";
import { useAuthStore } from "@/stores/authStore";
import { useUpdateUserPreference, useUserPreferenceNormal } from "@/hooks/useDirectory";
import { useTheme } from "@/providers/ThemeProvider/useTheme";
import { useLayout } from "@/providers/LayoutProvider/useLayout";
import { ChangeLanguage } from "@/assets/languages/i18n";
import { LANGUAGE } from "@/assets/languages/i18nResources";
import type { Language } from "@/assets/languages/i18nResources";

/**
 * Hook to sync theme preference with backend
 * Applies theme on initial load/login and syncs local changes to backend
 */
export const useThemeSync = () => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const { themeName, setTheme } = useTheme();
  const updatePreference = useUpdateUserPreference({ silent: true });

  const shouldFetch = !!currentUserId && isAuthValid;

  const { data: preference } = useUserPreferenceNormal({
    id: currentUserId || "00000000-0000-0000-0000-000000000000",
    options: {
      enabled: shouldFetch && !!currentUserId,
    },
  });

  const lastPreferenceId = useRef<string | null>(null);
  const isInitialized = useRef(false);

  // Apply theme preference ONLY on initial load/login
  useEffect(() => {
    if (!shouldFetch || !preference || !isAuthValid || !currentUserId) {
      isInitialized.current = false;
      lastPreferenceId.current = null;
      return;
    }

    // Only apply on initial load (when preference ID changes or first time)
    if (lastPreferenceId.current === preference.id && isInitialized.current) {
      return;
    }

    // Apply theme (only if different)
    if (preference.theme && preference.theme !== themeName) {
      setTheme(preference.theme as any);
    }

    isInitialized.current = true;
    lastPreferenceId.current = preference.id;
  }, [preference, isAuthValid, currentUserId, themeName, setTheme, shouldFetch]);

  // Sync theme changes to backend
  const syncTheme = useCallback(
    async (theme: string) => {
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
    [currentUserId, isAuthValid, updatePreference]
  );

  return { syncTheme };
};

/**
 * Hook to sync language preference with backend
 * Applies language on initial load/login and syncs local changes to backend
 */
export const useLanguageSync = () => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const updatePreference = useUpdateUserPreference({ silent: true });

  const shouldFetch = !!currentUserId && isAuthValid;

  const { data: preference } = useUserPreferenceNormal({
    id: currentUserId || "00000000-0000-0000-0000-000000000000",
    options: {
      enabled: shouldFetch && !!currentUserId,
    },
  });

  const lastPreferenceId = useRef<string | null>(null);
  const isInitialized = useRef(false);

  // Apply language preference ONLY on initial load/login
  useEffect(() => {
    if (!shouldFetch || !preference || !isAuthValid || !currentUserId) {
      isInitialized.current = false;
      lastPreferenceId.current = null;
      return;
    }

    // Only apply on initial load (when preference ID changes or first time)
    if (lastPreferenceId.current === preference.id && isInitialized.current) {
      return;
    }

    // Apply language (only if different)
    if (preference.language) {
      const langMap: Record<string, Language> = {
        en: LANGUAGE.EN,
        zh: LANGUAGE.ZH,
        fr: LANGUAGE.FR,
      };
      const targetLang = langMap[preference.language.toLowerCase()];
      if (targetLang) {
        const currentLang = localStorage.getItem("i18nextLng") || LANGUAGE.EN;
        if (currentLang !== targetLang) {
          ChangeLanguage(targetLang);
        }
      }
    }

    isInitialized.current = true;
    lastPreferenceId.current = preference.id;
  }, [preference, isAuthValid, currentUserId, shouldFetch]);

  // Sync language changes to backend
  const syncLanguage = useCallback(
    async (language: string) => {
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
    [currentUserId, isAuthValid, updatePreference]
  );

  return { syncLanguage };
};

/**
 * Hook to sync layout preference with backend
 * Applies layout on initial load/login and syncs local changes to backend
 */
export const useLayoutSync = () => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const { layoutMode, setLayoutMode } = useLayout();
  const updatePreference = useUpdateUserPreference({ silent: true });

  const shouldFetch = !!currentUserId && isAuthValid;

  const { data: preference } = useUserPreferenceNormal({
    id: currentUserId || "00000000-0000-0000-0000-000000000000",
    options: {
      enabled: shouldFetch && !!currentUserId,
    },
  });

  const lastPreferenceId = useRef<string | null>(null);
  const isInitialized = useRef(false);

  // Apply layout preference ONLY on initial load/login
  useEffect(() => {
    if (!shouldFetch || !preference || !isAuthValid || !currentUserId) {
      isInitialized.current = false;
      lastPreferenceId.current = null;
      return;
    }

    // Only apply on initial load (when preference ID changes or first time)
    // Once initialized, NEVER override local state from backend
    // This ensures user's click actions immediately update UI without backend interference
    if (lastPreferenceId.current === preference.id && isInitialized.current) {
      return;
    }

    // Apply layout mode ONLY on initial load (first time we get preference data)
    // After initialization, NEVER override local state from backend
    // User's local changes take precedence - backend sync is just for persistence
    if (preference.display && typeof preference.display === "object") {
      const display = preference.display as Record<string, unknown>;
      if (display.layoutMode && (display.layoutMode === "show" || display.layoutMode === "hide")) {
        // ONLY apply on initial load (when isInitialized is false)
        // After initialization, we NEVER touch layoutMode from backend
        if (!isInitialized.current && display.layoutMode !== layoutMode) {
          setLayoutMode(display.layoutMode);
        }
      }
    }

    isInitialized.current = true;
    lastPreferenceId.current = preference.id;
  }, [preference, isAuthValid, currentUserId, setLayoutMode, shouldFetch]);

  // Sync layout changes to backend (silent, fire-and-forget, failure doesn't matter)
  // This is called AFTER local state is already updated, just for persistence
  const syncLayout = useCallback(
    async (layoutMode: "show" | "hide") => {
      if (!currentUserId || !isAuthValid) return;

      // Fire and forget - don't wait for response, failure is OK
      // Local state is already updated, this is just for backend persistence
      const currentDisplay = preference?.display as Record<string, unknown> | undefined;
      updatePreference.mutateAsync({
        id: currentUserId,
        data: {
          display: {
            ...currentDisplay,
            layoutMode,
          },
        },
      }).catch((error) => {
        // Silent error - only log to console, don't show error toast
        // Failure doesn't matter, local state is already updated
        console.error("Failed to sync layout preference (non-critical):", error);
      });
    },
    [currentUserId, isAuthValid, updatePreference, preference]
  );

  return { syncLayout };
};