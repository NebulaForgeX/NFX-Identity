/**
 * Dashboard background type definitions
 */
export type DashboardBackgroundType = "waves" | "squares" | "letterGlitch" | "pixelBlast" | "none";

/**
 * Default dashboard background value
 */
export const DEFAULT_DASHBOARD_BACKGROUND: DashboardBackgroundType = "none";

/**
 * Dashboard background values array (for zod enum)
 */
export const DASHBOARD_BACKGROUND_VALUES: readonly DashboardBackgroundType[] = [
  "waves",
  "squares",
  "letterGlitch",
  "pixelBlast",
  "none",
] as const;
