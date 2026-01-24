import { zodResolver } from "@hookform/resolvers/zod";
import { useTranslation } from "react-i18next";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { type ProfileFormValues } from "../schemas/profileSchema";
import type { UserProfile } from "@/types";
import type { KeyValuePair } from "@/components/KeyValueEditor";

// Helper to convert map to key-value pairs
const mapToKeyValuePairs = (map?: Record<string, unknown>): KeyValuePair[] => {
  if (!map) return [];
  return Object.entries(map).map(([key, value]) => ({
    key,
    value: Array.isArray(value) ? value.map(String) : String(value),
  }));
};

export const useInitProfileForm = (profile?: UserProfile) => {
  const { t } = useTranslation("elements.directory");

  // 动态创建 schema，使用翻译
  const KeyValuePairSchema = z.object({
    key: z.string().trim().min(1),
    value: z.union([z.string(), z.array(z.string())]),
  });

  const ProfileFormSchema = z.object({
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

  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(ProfileFormSchema),
    mode: "onChange",
    defaultValues: {
      role: profile?.role || "",
      firstName: profile?.firstName || "",
      lastName: profile?.lastName || "",
      nickname: profile?.nickname || "",
      displayName: profile?.displayName || "",
      bio: profile?.bio || "",
      birthday: profile?.birthday ? new Date(profile.birthday).toISOString().split("T")[0] : "",
      age: profile?.age || undefined,
      gender: profile?.gender || "",
      location: profile?.location || "",
      website: profile?.website || "",
      github: profile?.github || "",
      socialLinks: mapToKeyValuePairs(profile?.socialLinks as Record<string, unknown> | undefined),
      skills: mapToKeyValuePairs(profile?.skills as Record<string, unknown> | undefined),
    },
  });

  return form;
};
