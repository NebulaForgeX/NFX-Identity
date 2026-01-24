import { z } from "zod";

// 基础 schema 用于类型推断（不包含翻译，仅用于类型）
const BaseBootstrapFormSchema = z
  .object({
    Version: z.string().trim().min(1).default("1.0.0"),
    AdminUsername: z.string().trim().min(1).min(3),
    AdminPassword: z.string().min(8),
    AdminPasswordConfirm: z.string().min(8),
    AdminEmail: z
      .string()
      .trim()
      .refine((val) => !val || z.string().email().safeParse(val).success)
      .optional(),
    AdminPhone: z.string().trim().optional(),
    AdminCountryCode: z.string().trim().optional(),
  })
  .refine((data) => data.AdminPassword === data.AdminPasswordConfirm, {
    path: ["AdminPasswordConfirm"],
  });

export type BootstrapFormValues = z.input<typeof BaseBootstrapFormSchema>;
