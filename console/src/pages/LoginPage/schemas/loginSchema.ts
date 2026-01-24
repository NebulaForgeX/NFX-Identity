import { z } from "zod";

// 创建邮箱登录 Schema 工厂函数
export const createEmailLoginSchema = (t: (key: string) => string) => {
  return z.object({
    loginType: z.literal("email"),
    email: z.email(t("validation.emailInvalid")),
    password: z
      .string()
      .min(1, t("validation.passwordRequired"))
      .min(6, t("validation.passwordMinLength")),
  });
};

// 创建手机号登录 Schema 工厂函数
export const createPhoneLoginSchema = (t: (key: string) => string) => {
  return z.object({
    loginType: z.literal("phone"),
    code: z
      .string()
      .trim()
      .min(1, t("validation.codeRequired"))
      .regex(/^\d+$/, t("validation.codeNumeric")),
    phone: z
      .string()
      .trim()
      .min(1, t("validation.phoneRequired"))
      .regex(/^\d+$/, t("validation.phoneNumeric"))
      .min(6, t("validation.phoneMinLength")),
    password: z
      .string()
      .min(1, t("validation.passwordRequired"))
      .min(6, t("validation.passwordMinLength")),
  });
};

// 基础 Schema（用于类型推断，不包含翻译）
const BaseEmailLoginSchema = z.object({
  loginType: z.literal("email"),
  email: z.email(),
  password: z.string().min(1).min(6),
});

const BasePhoneLoginSchema = z.object({
  loginType: z.literal("phone"),
  code: z.string().trim().min(1).regex(/^\d+$/),
  phone: z.string().trim().min(1).regex(/^\d+$/).min(6),
  password: z.string().min(1).min(6),
});

export const BaseLoginSchema = z.discriminatedUnion("loginType", [
  BaseEmailLoginSchema,
  BasePhoneLoginSchema,
]);

export type EmailLoginFormValues = z.infer<typeof BaseEmailLoginSchema>;
export type PhoneLoginFormValues = z.infer<typeof BasePhoneLoginSchema>;
export type LoginFormValues = z.infer<typeof BaseLoginSchema>;
