import type { Theme, ThemeName } from "../types";

import { coffeeTheme } from "./coffee";
import { corporateTheme } from "./corporate";
import { cosmicTheme } from "./cosmic";
import { darkTheme } from "./dark";
import { defaultTheme } from "./default";
import { freshTheme } from "./fresh";
import { lightTheme } from "./light";

export const themes: Record<ThemeName, Theme> = {
  default: defaultTheme,
  light: lightTheme,
  corporate: corporateTheme,
  fresh: freshTheme,
  dark: darkTheme,
  cosmic: cosmicTheme,
  coffee: coffeeTheme,
};

export { defaultTheme, lightTheme, darkTheme, cosmicTheme, corporateTheme, freshTheme, coffeeTheme };
