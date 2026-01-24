import { z } from "zod";

// Key-value pair schema for form
const KeyValuePairSchema = z.object({
  key: z.string().trim().min(1),
  value: z.union([z.string(), z.array(z.string())]),
});

// 基础 schema 用于类型推断（不包含翻译，仅用于类型）
const BaseProfileFormSchema = z.object({
  role: z.string().trim().optional(),
  firstName: z.string().trim().optional(),
  lastName: z.string().trim().optional(),
  nickname: z.string().trim().optional(),
  displayName: z.string().trim().optional(),
  bio: z.string().trim().optional(),
  birthday: z.string().optional(),
  age: z.number().int().positive().optional(),
  gender: z.string().trim().optional(),
  location: z.string().trim().optional(),
  website: z.string().url().optional().or(z.literal("")),
  github: z.string().trim().optional(),
  socialLinks: z.array(KeyValuePairSchema).optional(),
  skills: z.array(KeyValuePairSchema).optional(),
});

export type ProfileFormValues = z.input<typeof BaseProfileFormSchema>;
