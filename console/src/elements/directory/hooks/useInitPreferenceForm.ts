import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { type PreferenceFormValues } from "../schemas/preferenceSchema";

export const useInitPreferenceForm = (defaultValues?: Partial<PreferenceFormValues>) => {
  // 动态创建 schema
  const PreferenceFormSchema = z.object({
    theme: z.string().optional(),
    language: z.string().optional(),
    timezone: z.string().optional(),
    dashboardBackground: z.enum(["waves", "squares", "letterGlitch", "pixelBlast", "none"]).optional(),
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
      dashboardBackground: defaultValues?.dashboardBackground || "none",
      notifications: defaultValues?.notifications || {},
      privacy: defaultValues?.privacy || {},
      display: defaultValues?.display || {},
      other: defaultValues?.other || {},
    },
  });

  return form;
};
