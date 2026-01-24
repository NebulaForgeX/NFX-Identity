import { z } from "zod";

// 基础 schema 用于类型推断（不包含翻译，仅用于类型）
const BasePreferenceFormSchema = z.object({
  key: z.string().trim().min(1),
  value: z.string().trim().min(1),
});

export type PreferenceFormValues = z.input<typeof BasePreferenceFormSchema>;
