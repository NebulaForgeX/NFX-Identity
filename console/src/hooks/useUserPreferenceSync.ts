import { useEffect, useCallback, useRef } from "react";
import { useAuthStore } from "@/stores/authStore";
import { useUpdateUserPreference, useUserPreferenceNormal } from "@/hooks/useDirectory";
import { useTheme } from "@/hooks/useTheme";
import { ChangeLanguage } from "@/assets/languages/i18n";
import { LANGUAGE } from "@/assets/languages/i18nResources";
import type { Language } from "@/assets/languages/i18nResources";
import LayoutStore from "@/stores/layoutStore";

/**
 * Hook to sync user preferences (theme, language, layout) with backend
 * and apply preferences on login
 */
export const useUserPreferenceSync = () => {
  const currentUserId = useAuthStore((state) => state.currentUserId);
  const isAuthValid = useAuthStore((state) => state.isAuthValid);
  const { themeName, setTheme } = useTheme();
  const updatePreference = useUpdateUserPreference();
  
  // Get user preference when authenticated
  // Only fetch when user is logged in
  const shouldFetch = !!currentUserId && isAuthValid;
  
  // Use normal mode (non-suspense) so we can use enabled option
  const { data: preference } = useUserPreferenceNormal({
    id: currentUserId || "00000000-0000-0000-0000-000000000000",
    options: {
      enabled: shouldFetch && !!currentUserId,
    },
  });

  // Track if we've already applied preferences to prevent loops
  const hasAppliedPreferences = useRef(false);
  const lastPreferenceId = useRef<string | null>(null);

  // Apply preferences on login or when preference data changes
  useEffect(() => {
    if (!shouldFetch || !preference || !isAuthValid || !currentUserId) {
      hasAppliedPreferences.current = false;
      return;
    }

    // Only apply if this is a new preference (different ID) or first time
    if (lastPreferenceId.current === preference.id && hasAppliedPreferences.current) {
      return;
    }

    // Apply theme
    if (preference.theme && preference.theme !== themeName) {
      setTheme(preference.theme as any);
    }

    // Apply language
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

    // Apply layout mode (store in display.other or use a specific field)
    // For now, we'll store layout in the display field
    if (preference.display && typeof preference.display === "object") {
      const display = preference.display as Record<string, unknown>;
      if (display.layoutMode && (display.layoutMode === "show" || display.layoutMode === "hide")) {
        LayoutStore.getState().setLayoutMode(display.layoutMode);
      }
    }

    hasAppliedPreferences.current = true;
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

  // Sync layout changes to backend
  const syncLayout = useCallback(
    async (layoutMode: "show" | "hide") => {
      if (!currentUserId || !isAuthValid) return;
      
      try {
        // Store layout in display field
        const currentDisplay = preference?.display as Record<string, unknown> | undefined;
        await updatePreference.mutateAsync({
          id: currentUserId,
          data: {
            display: {
              ...currentDisplay,
              layoutMode,
            },
          },
        });
      } catch (error) {
        console.error("Failed to sync layout preference:", error);
      }
    },
    [currentUserId, isAuthValid, updatePreference, preference]
  );

  return {
    syncTheme,
    syncLanguage,
    syncLayout,
  };
};
