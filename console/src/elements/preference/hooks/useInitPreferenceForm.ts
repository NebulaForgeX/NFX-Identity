import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

import { DEFAULT_DASHBOARD_BACKGROUND } from "nfx-ui/preference";

import { type PreferenceFormValues, PreferenceFormSchema } from "../schemas/preferenceSchema";

export const useInitPreferenceForm = (defaultValues?: Partial<PreferenceFormValues>) => {
  const form = useForm<PreferenceFormValues>({
    resolver: zodResolver(PreferenceFormSchema),
    mode: "onChange",
    defaultValues: {
      theme: defaultValues?.theme || "",
      base: defaultValues?.base || "",
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
