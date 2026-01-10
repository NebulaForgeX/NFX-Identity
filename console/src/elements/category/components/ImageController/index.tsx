import type { CategoryFormValues } from "../../controllers/categorySchema";

import { memo, useCallback, useRef, useState } from "react";
import { Controller, useFormContext } from "react-hook-form";

import { Upload, X } from "@/assets/icons/lucide";
import { buildImageUrl } from "@/utils/image";

import styles from "./styles.module.css";

interface ImageControllerProps {
  existingImageUrl?: string;
}

const ImageController = memo(({ existingImageUrl }: ImageControllerProps) => {
  const { control, watch, setValue } = useFormContext<CategoryFormValues>();
  const fileInputRef = useRef<HTMLInputElement>(null);
  const imageFile = watch("Image");
  const [isImageRemoved, setIsImageRemoved] = useState(false);

  const handleFileChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const file = e.target.files?.[0];
      if (file) {
        if (!file.type.startsWith("image/")) return;
        setValue("Image", file, { shouldValidate: true });
        setIsImageRemoved(false); // 重新上传时重置删除标记
      }
      if (fileInputRef.current) fileInputRef.current.value = "";
    },
    [setValue],
  );

  const handleRemoveImage = useCallback(() => {
    setValue("Image", null, { shouldValidate: true });
    setIsImageRemoved(true); // 标记图片已删除
  }, [setValue]);

  return (
    <Controller<CategoryFormValues>
      control={control}
      name="Image"
      render={({ fieldState: { error } }) => (
        <div className={styles.container}>
          <label className={styles.label}>分类图片</label>

          <div className={styles.imageContainer}>
            {(imageFile || (existingImageUrl && !isImageRemoved)) && (
              <div className={styles.imagePreview}>
                <img
                  src={imageFile ? URL.createObjectURL(imageFile) : buildImageUrl(existingImageUrl, "category")}
                  alt="Category preview"
                  className={styles.image}
                />
                <button type="button" className={styles.deleteBtn} onClick={handleRemoveImage} aria-label="删除图片">
                  <X size={36} />
                </button>
              </div>
            )}

            {!imageFile && (!existingImageUrl || isImageRemoved) && (
              <button type="button" className={styles.uploadBtn} onClick={() => fileInputRef.current?.click()}>
                <Upload className={styles.uploadIcon} size={32} />
                <span className={styles.uploadText}>上传图片</span>
              </button>
            )}
          </div>

          <input
            ref={fileInputRef}
            type="file"
            accept="image/*"
            onChange={handleFileChange}
            className={styles.fileInput}
          />

          {error && <p className={styles.error}>{error.message}</p>}
          <p className={styles.hint}>💡 建议上传 16:9 比例的图片，大小不超过 10MB</p>
        </div>
      )}
    />
  );
});

ImageController.displayName = "ImageController";

export default ImageController;
