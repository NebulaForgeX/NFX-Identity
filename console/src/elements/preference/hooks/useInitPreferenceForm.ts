import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import type { DashboardBackgroundType } from "@/types";
import { DASHBOARD_BACKGROUND_VALUES, DEFAULT_DASHBOARD_BACKGROUND } from "@/types";
import { type PreferenceFormValues } from "../schemas/preferenceSchema";

export const useInitPreferenceForm = (defaultValues?: Partial<PreferenceFormValues>) => {
  // 动态创建 schema
  const PreferenceFormSchema = z.object({
    theme: z.string().optional(),
    language: z.string().optional(),
    timezone: z.string().optional(),
    dashboardBackground: z.enum(DASHBOARD_BACKGROUND_VALUES as [DashboardBackgroundType, ...DashboardBackgroundType[]]).optional(),
    notifications: z.record(z.string(), z.unknown()).optional(),
    privacy: z.record(z.string(), z.unknown()).optional(),
    display: z.record(z.string(), z.unknown()).optional(),
    other: z.record(z.string(), z.unknown()).optional(),
  });

  const form = useForm<PreferenceFormValues>({
    resolver: zodResolver(PreferenceFormSchema),
    mode: "onChange",
    defaultValues: {
      theme: defaultValues?.theme || "",
      language: defaultValues?.language || "",
      timezone: defaultValues?.timezone || "",
      dashboardBackground: defaultValues?.dashboardBackground || DEFAULT_DASHBOARD_BACKGROUND,
      notifications: defaultValues?.notifications || {},
      privacy: defaultValues?.privacy || {},
      display: defaultValues?.display || {},
      other: defaultValues?.other || {},
    },
  });

  return form;
};
