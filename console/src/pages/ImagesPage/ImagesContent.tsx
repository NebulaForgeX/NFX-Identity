import { memo, useCallback, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

import {
  useCreateUserImage,
  useDeleteUserImage,
  useUploadImage,
  useUserImagesByUserID,
} from "@/hooks";
import { showError } from "@/stores/modalStore";

import {
  ImagesConfirmSection,
  ImagesEmptyState,
  ImagesGallery,
  ImagesUploadArea,
} from "./components";
import styles from "./styles.module.css";

interface ImagesContentProps {
  userId: string;
}

export const ImagesContent = memo(function ImagesContent({ userId }: ImagesContentProps) {
  const { t } = useTranslation("ImagesPage");
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isDragging, setIsDragging] = useState(false);
  const [uploadedImageId, setUploadedImageId] = useState<string | null>(null);

  const { data: userImages = [] } = useUserImagesByUserID({ userId });
  const uploadMutation = useUploadImage();
  const confirmMutation = useCreateUserImage();
  const deleteMutation = useDeleteUserImage();

  const handleFileSelect = useCallback(
    (file: File) => {
      if (!file.type.startsWith("image/")) {
        showError(t("invalidFileType", "Please select an image file"));
        return;
      }
      if (file.size > 10 * 1024 * 1024) {
        showError(t("fileTooLarge", "File size must be less than 10MB"));
        return;
      }
      uploadMutation.mutate(file);
    },
    [uploadMutation, t]
  );

  const handleInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const file = e.target.files?.[0];
      if (file) handleFileSelect(file);
      e.target.value = "";
    },
    [handleFileSelect]
  );

  const handleDragOver = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    setIsDragging(true);
  }, []);

  const handleDragLeave = useCallback((e: React.DragEvent) => {
    e.preventDefault();
    setIsDragging(false);
  }, []);

  const handleDrop = useCallback(
    (e: React.DragEvent) => {
      e.preventDefault();
      setIsDragging(false);
      const file = e.dataTransfer.files?.[0];
      if (file) handleFileSelect(file);
    },
    [handleFileSelect]
  );

  const handleClickUploadArea = useCallback(() => {
    fileInputRef.current?.click();
  }, []);

  const handleConfirm = useCallback(() => {
    if (!uploadedImageId) return;
    confirmMutation.mutate(
      {
        userId,
        imageId: uploadedImageId,
        displayOrder: userImages.length,
      },
      {
        onSuccess: () => setUploadedImageId(null),
      }
    );
  }, [uploadedImageId, userId, userImages.length, confirmMutation]);

  const handleDelete = useCallback(
    (userImageId: string) => {
      deleteMutation.mutate(userImageId);
    },
    [deleteMutation]
  );

  useEffect(() => {
    if (uploadMutation.isSuccess && uploadMutation.data?.id) {
      setUploadedImageId(uploadMutation.data.id);
    }
  }, [uploadMutation.isSuccess, uploadMutation.data?.id]);

  return (
    <div className={styles.content}>
      <ImagesUploadArea
        fileInputRef={fileInputRef}
        isDragging={isDragging}
        isUploading={uploadMutation.isPending}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        onDrop={handleDrop}
        onClick={handleClickUploadArea}
        onInputChange={handleInputChange}
        disabled={uploadMutation.isPending || confirmMutation.isPending}
      />
      <ImagesConfirmSection
        uploadedImageId={uploadedImageId ?? ""}
        isConfirming={confirmMutation.isPending}
        onCancel={() => setUploadedImageId(null)}
        onConfirm={handleConfirm}
      />
      {userImages.length > 0 ? (
        <ImagesGallery userImages={userImages} onDelete={handleDelete} />
      ) : (
        <ImagesEmptyState />
      )}
    </div>
  );
});
