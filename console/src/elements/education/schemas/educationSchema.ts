import { z } from "zod";

// 基础 schema 用于类型推断（不包含翻译，仅用于类型）
const BaseEducationFormSchema = z.object({
  school: z.string().trim().min(1),
  degree: z.string().trim().optional(),
  major: z.string().trim().optional(),
  fieldOfStudy: z.string().trim().optional(),
  startDate: z.string().optional(),
  endDate: z.string().optional(),
  isCurrent: z.boolean().optional(),
  description: z.string().optional(),
  grade: z.string().trim().optional(),
  activities: z.string().trim().optional(),
  achievements: z.string().trim().optional(),
});

export type EducationFormValues = z.input<typeof BaseEducationFormSchema>;
