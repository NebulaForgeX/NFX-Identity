import type { Theme } from "../types";
import { ThemeEnum } from "../types";
import { defaultBaseTheme } from "../bases/default";
import { coffeeColorTheme } from "./coffee";
import { corporateColorTheme } from "./corporate";
import { cosmicColorTheme } from "./cosmic";
import { darkColorTheme } from "./dark";
import { defaultColorTheme } from "./default";
import { forestColorTheme } from "./forest";
import { lightColorTheme } from "./light";
import { wineColorTheme } from "./wine";

const defaultTheme: Theme = { colors: defaultColorTheme, base: defaultBaseTheme };
const lightTheme: Theme = { colors: lightColorTheme, base: defaultBaseTheme };
const darkTheme: Theme = { colors: darkColorTheme, base: defaultBaseTheme };
const cosmicTheme: Theme = { colors: cosmicColorTheme, base: defaultBaseTheme };
const corporateTheme: Theme = { colors: corporateColorTheme, base: defaultBaseTheme };
const forestTheme: Theme = { colors: forestColorTheme, base: defaultBaseTheme };
const coffeeTheme: Theme = { colors: coffeeColorTheme, base: defaultBaseTheme };
const wineTheme: Theme = { colors: wineColorTheme, base: defaultBaseTheme };

export const themes: Record<(typeof ThemeEnum)[keyof typeof ThemeEnum], Theme> = {
  [ThemeEnum.DEFAULT]: defaultTheme,
  [ThemeEnum.LIGHT]: lightTheme,
  [ThemeEnum.DARK]: darkTheme,
  [ThemeEnum.COSMIC]: cosmicTheme,
  [ThemeEnum.CORPORATE]: corporateTheme,
  [ThemeEnum.FOREST]: forestTheme,
  [ThemeEnum.COFFEE]: coffeeTheme,
  [ThemeEnum.WINE]: wineTheme,
};

export { defaultTheme, lightTheme, darkTheme, cosmicTheme, corporateTheme, forestTheme, coffeeTheme, wineTheme };
