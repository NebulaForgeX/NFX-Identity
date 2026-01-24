import { zodResolver } from "@hookform/resolvers/zod";
import { useTranslation } from "react-i18next";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { type OccupationFormValues } from "../schemas/occupationSchema";
import type { UserOccupation } from "@/types";

export const useInitOccupationForm = (occupation?: UserOccupation) => {
  const { t } = useTranslation("elements.directory");

  // 动态创建 schema，使用翻译
  const OccupationFormSchema = z.object({
    company: z.string().trim().min(1, t("occupation.company.required")),
    position: z.string().trim().min(1, t("occupation.position.required") || "Position is required"), // 后端要求 position 是必需的
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

  const form = useForm<OccupationFormValues>({
    resolver: zodResolver(OccupationFormSchema),
    mode: "onChange",
    defaultValues: {
      company: occupation?.company || "",
      position: occupation?.position || "",
      department: occupation?.department || "",
      industry: occupation?.industry || "",
      location: occupation?.location || "",
      employmentType: occupation?.employmentType || "",
      startDate: occupation?.startDate ? new Date(occupation.startDate).toISOString().split("T")[0] : "",
      endDate: occupation?.endDate ? new Date(occupation.endDate).toISOString().split("T")[0] : "",
      isCurrent: occupation?.isCurrent || false,
      description: occupation?.description || "",
      responsibilities: occupation?.responsibilities || "",
      achievements: occupation?.achievements || "",
      skillsUsed: occupation?.skillsUsed || [],
    },
  });

  return form;
};
