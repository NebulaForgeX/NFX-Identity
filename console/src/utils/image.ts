import { API_ENDPOINTS } from "@/apis/ip";

/** 用 imageId 拼出后端按 ID 返回图片的 URL（头像/背景等） */
export const buildImageUrl = (imageId: string | null | undefined): string => {
  if (!imageId) return "";
  const base = API_ENDPOINTS.PURE.replace(/\/$/, "");
  return `${base}/image/public/images/${imageId}`;
};
