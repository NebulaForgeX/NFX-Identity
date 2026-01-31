import { z } from "zod";

// 基础 schema 用于类型推断（不包含翻译，仅用于类型）
const BaseOccupationFormSchema = z.object({
  company: z.string().trim().min(1),
  position: z.string().trim().min(1), // 后端要求 position 是必需的
  department: z.string().trim().optional(),
  industry: z.string().trim().optional(),
  location: z.string().trim().optional(),
  employmentType: z.string().trim().optional(),
  startDate: z.string().optional(),
  endDate: z.string().optional(),
  isCurrent: z.boolean().optional(),
  description: z.string().trim().optional(),
  responsibilities: z.string().trim().optional(),
  achievements: z.string().trim().optional(),
  skillsUsed: z.array(z.string()).optional(),
});

export type OccupationFormValues = z.input<typeof BaseOccupationFormSchema>;
