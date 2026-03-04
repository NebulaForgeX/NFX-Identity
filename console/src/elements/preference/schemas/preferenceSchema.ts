import { z } from "zod";

import { DashboardBackgroundEnum } from "nfx-ui/preference";

const BasePreferenceFormSchema = z.object({
  theme: z.string().optional(),
  base: z.string().optional(),
  language: z.string().optional(),
  timezone: z.string().optional(),
  dashboardBackground: z.nativeEnum(DashboardBackgroundEnum).optional(),
  notifications: z.record(z.string(), z.unknown()).optional(),
  privacy: z.record(z.string(), z.unknown()).optional(),
  display: z.record(z.string(), z.unknown()).optional(),
  other: z.record(z.string(), z.unknown()).optional(),
});

export type PreferenceFormValues = z.input<typeof BasePreferenceFormSchema>;
export { BasePreferenceFormSchema as PreferenceFormSchema };
