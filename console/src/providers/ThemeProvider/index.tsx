import type { ReactNode } from "react";
import type { Theme, ThemeName } from "@/assets/themes/types";
import type { BaseName } from "@/assets/themes/types";

import { createContext, useMemo, useState } from "react";

import { BASE_VALUES, ThemeEnum, THEME_VALUES } from "@/assets/themes/types";
import { bases } from "@/assets/themes/bases";
import { themes } from "@/assets/themes/colors";

import useThemeVariables from "./hooks/useThemeVariables";

export interface ThemeContextType {
  currentTheme: Theme;
  themeName: ThemeName;
  baseName: BaseName;
  setTheme: (themeName: ThemeName) => void;
  setBase: (baseName: BaseName) => void;
  availableThemes: ThemeName[];
  availableBases: BaseName[];
}

export const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

interface ThemeProviderProps {
  children: ReactNode;
  defaultTheme?: ThemeName;
  defaultBase?: BaseName;
}

export function ThemeProvider({
  children,
  defaultTheme = "default",
  defaultBase = "default",
}: ThemeProviderProps) {
  const [themeName, setThemeName] = useState<ThemeName>(() => {
    const saved = localStorage.getItem("theme") as string | null;
    if (saved === "fresh") {
      localStorage.setItem("theme", ThemeEnum.FOREST);
      return ThemeEnum.FOREST;
    }
    return saved && saved in themes ? (saved as ThemeName) : defaultTheme;
  });
  const [baseName, setBaseName] = useState<BaseName>(() => {
    const saved = localStorage.getItem("base") as BaseName | null;
    return saved && BASE_VALUES.includes(saved) ? saved : defaultBase;
  });

  const currentTheme = useMemo<Theme>(
    () => ({ colors: themes[themeName].colors, base: bases[baseName] }),
    [themeName, baseName],
  );

  useThemeVariables(currentTheme, themeName, baseName);

  const setTheme = (newTheme: ThemeName) => setThemeName(newTheme);
  const setBase = (newBase: BaseName) => {
    setBaseName(newBase);
    localStorage.setItem("base", newBase);
  };

  return (
    <ThemeContext.Provider
      value={{
        currentTheme,
        themeName,
        baseName,
        setTheme,
        setBase,
        availableThemes: THEME_VALUES,
        availableBases: BASE_VALUES,
      }}
    >
      {children}
    </ThemeContext.Provider>
  );
}

export default ThemeProvider;
