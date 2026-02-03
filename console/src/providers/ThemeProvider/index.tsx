import type { ReactNode } from "react";
import type { Theme, ThemeName } from "@/assets/themes/types";

import { createContext, useState } from "react";

import { themes } from "@/assets/themes/base";
import { THEME_VALUES } from "@/assets/themes/types";

import useThemeVariables from "./hooks/useThemeVariables";

export interface ThemeContextType {
  currentTheme: Theme;
  themeName: ThemeName;
  setTheme: (themeName: ThemeName) => void;
  availableThemes: ThemeName[];
}

export const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

interface ThemeProviderProps {
  children: ReactNode;
  defaultTheme?: ThemeName;
}

export function ThemeProvider({ children, defaultTheme = "default" }: ThemeProviderProps) {
  const [themeName, setThemeName] = useState<ThemeName>(() => {
    const saved = localStorage.getItem("theme") as ThemeName | null;
    return saved && saved in themes ? saved : defaultTheme;
  });

  const currentTheme = themes[themeName];

  // 将主题变量注入到 CSS Variables 并保存到 localStorage
  useThemeVariables(currentTheme, themeName);

  const setTheme = (newTheme: ThemeName) => {
    setThemeName(newTheme);
  };

  return (
    <ThemeContext.Provider value={{ currentTheme, themeName, setTheme, availableThemes: THEME_VALUES }}>
      {children}
    </ThemeContext.Provider>
  );
}
