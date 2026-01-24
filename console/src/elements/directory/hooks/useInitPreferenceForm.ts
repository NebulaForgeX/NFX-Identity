import { zodResolver } from "@hookform/resolvers/zod";
import { useTranslation } from "react-i18next";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { type PreferenceFormValues } from "../schemas/preferenceSchema";

export const useInitPreferenceForm = (defaultValues?: Partial<PreferenceFormValues>) => {
  const { t } = useTranslation("elements.directory");

  // 动态创建 schema，使用翻译
  const PreferenceFormSchema = z.object({
    key: z.string().trim().min(1, t("preference.key.required")),
    value: z.string().trim().min(1, t("preference.value.required")),
  });

  const form = useForm<PreferenceFormValues>({
    resolver: zodResolver(PreferenceFormSchema),
    mode: "onChange",
    defaultValues: {
      key: defaultValues?.key || "",
      value: defaultValues?.value || "",
    },
  });

  return form;
};
