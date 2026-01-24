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
  // Use silent mode for preference sync (no success/error toasts)
  const updatePreference = useUpdateUserPreference({ silent: true });
  
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
  // Track if preferences have been initialized (only apply on first load/login)
  const isInitialized = useRef(false);

  // Apply preferences ONLY on initial load/login, not on every preference change
  // This prevents backend data from overriding local state changes
  useEffect(() => {
    if (!shouldFetch || !preference || !isAuthValid || !currentUserId) {
      hasAppliedPreferences.current = false;
      isInitialized.current = false;
      lastPreferenceId.current = null;
      return;
    }

    // Only apply preferences on initial load (when preference ID changes or first time)
    // Don't re-apply if we've already initialized and preference ID hasn't changed
    // This ensures we only apply preferences once on login, not on every preference update
    if (lastPreferenceId.current === preference.id && isInitialized.current) {
      return;
    }

    // Apply theme (only if different)
    if (preference.theme && preference.theme !== themeName) {
      setTheme(preference.theme as any);
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

    // Apply layout mode ONLY on initial load (first time we get preference data)
    // After initialization, NEVER override local state from backend
    // User's local changes take precedence - backend sync is just for persistence
    if (preference.display && typeof preference.display === "object") {
      const display = preference.display as Record<string, unknown>;
      if (display.layoutMode && (display.layoutMode === "show" || display.layoutMode === "hide")) {
        // ONLY apply on initial load (when isInitialized is false)
        // After initialization, we NEVER touch layoutMode from backend
        // This ensures user's click actions immediately update UI without backend interference
        if (!isInitialized.current) {
          LayoutStore.getState().setLayoutMode(display.layoutMode);
        }
        // If already initialized, completely ignore backend layoutMode changes
      }
    }

    // Mark as initialized after first application
    hasAppliedPreferences.current = true;
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

  return {
    syncTheme,
    syncLanguage,
    syncLayout,
  };
};
