import { z } from "zod";
import type { DashboardBackgroundType } from "@/types";
import { DASHBOARD_BACKGROUND_VALUES } from "@/types";

// 基础 schema 用于类型推断（不包含翻译，仅用于类型）
const BasePreferenceFormSchema = z.object({
  theme: z.string().optional(),
  language: z.string().optional(),
  timezone: z.string().optional(),
  dashboardBackground: z.enum(DASHBOARD_BACKGROUND_VALUES as [DashboardBackgroundType, ...DashboardBackgroundType[]]).optional(),
  notifications: z.record(z.string(), z.unknown()).optional(),
  privacy: z.record(z.string(), z.unknown()).optional(),
  display: z.record(z.string(), z.unknown()).optional(),
  other: z.record(z.string(), z.unknown()).optional(),
});

export type PreferenceFormValues = z.input<typeof BasePreferenceFormSchema>;
