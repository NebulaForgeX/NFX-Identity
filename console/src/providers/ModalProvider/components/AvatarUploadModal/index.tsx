import { memo, useCallback, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { useMutation, useQueryClient } from "@tanstack/react-query";

import { Camera, Upload, X } from "@/assets/icons/lucide";
import { DIRECTORY_QUERY_KEY_PREFIXES } from "@/constants";
import { CreateOrUpdateUserAvatar } from "@/apis/directory.api";
import { UploadImage } from "@/apis/image.api";
import ModalStore, { useModalStore } from "@/stores/modalStore";

import styles from "./styles.module.css";

const AvatarUploadModal = memo(() => {
  const { t } = useTranslation("Modal");
  const dialogRef = useRef<HTMLDialogElement>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const queryClient = useQueryClient();

  const isOpen = useModalStore((state) => state.avatarUploadModal.isOpen);
  const userId = useModalStore((state) => state.avatarUploadModal.userId);
  const onSuccess = useModalStore((state) => state.avatarUploadModal.onSuccess);
  const onCancel = useModalStore((state) => state.avatarUploadModal.onCancel);
  const hideModal = ModalStore.getState().hideModal;

  const [preview, setPreview] = useState<string | null>(null);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploadedImageId, setUploadedImageId] = useState<string | null>(null);
  const [isDragging, setIsDragging] = useState(false);

  // 第一步：上传图片到 tmp
  const uploadMutation = useMutation({
    mutationFn: async (file: File) => {
      // 上传到 tmp 目录
      return await UploadImage(file);
    },
    onSuccess: (data) => {
      // 保存上传后的 image id，等待用户确认
      setUploadedImageId(data.id);
    },
  });

  // 第二步：用户确认后，调用 Directory 创建头像关联（后端会移动图片）
  const confirmMutation = useMutation({
    mutationFn: async (imageId: string) => {
      if (!userId) throw new Error("userId is required");
      // 调用 Directory 服务创建头像关联
      // Directory 服务会通过 gRPC 验证图片并移动到 avatar 目录
      await CreateOrUpdateUserAvatar({
        userId: userId,
        imageId: imageId,
      });
      return imageId;
    },
    onSuccess: (imageId) => {
      // 刷新用户头像数据
      queryClient.invalidateQueries({ queryKey: [...DIRECTORY_QUERY_KEY_PREFIXES.USER_AVATAR, userId] });
      onSuccess?.(imageId);
      handleClose();
    },
  });

  const handleClose = useCallback(() => {
    setPreview(null);
    setSelectedFile(null);
    setUploadedImageId(null);
    hideModal("avatarUpload");
    onCancel?.();
  }, [hideModal, onCancel]);

  useEffect(() => {
    const dialog = dialogRef.current;
    if (!dialog) return;
    if (isOpen && !dialog.open) {
      dialog.showModal();
    } else if (!isOpen && dialog.open) {
      dialog.close();
    }
  }, [isOpen]);

  const handleFileSelect = useCallback((file: File) => {
    // 验证文件类型
    if (!file.type.startsWith("image/")) {
      alert(t("avatarUpload.invalidFileType", "Please select an image file"));
      return;
    }
    // 验证文件大小 (最大 10MB)
    if (file.size > 10 * 1024 * 1024) {
      alert(t("avatarUpload.fileTooLarge", "File size must be less than 10MB"));
      return;
    }

    setSelectedFile(file);
    const reader = new FileReader();
    reader.onload = (e) => {
      setPreview(e.target?.result as string);
    };
    reader.readAsDataURL(file);
  }, [t]);

  const handleInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const file = e.target.files?.[0];
      if (file) {
        handleFileSelect(file);
      }
    },
    [handleFileSelect],
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
      if (file) {
        handleFileSelect(file);
      }
    },
    [handleFileSelect],
  );

  const handleUpload = useCallback(() => {
    if (selectedFile) {
      uploadMutation.mutate(selectedFile);
    }
  }, [selectedFile, uploadMutation]);

  const handleConfirm = useCallback(() => {
    if (uploadedImageId) {
      confirmMutation.mutate(uploadedImageId);
    }
  }, [uploadedImageId, confirmMutation]);

  const handleClickUploadArea = useCallback(() => {
    fileInputRef.current?.click();
  }, []);

  return (
    <dialog ref={dialogRef} className={styles.modal} onClose={handleClose}>
      <div className={styles.content}>
        <div className={styles.header}>
          <h3 className={styles.title}>{t("avatarUpload.title", "Upload Avatar")}</h3>
          <button className={styles.closeIcon} onClick={handleClose}>
            <X size={20} />
          </button>
        </div>

        <div
          className={`${styles.uploadArea} ${isDragging ? styles.dragging : ""} ${preview ? styles.hasPreview : ""}`}
          onClick={handleClickUploadArea}
          onDragOver={handleDragOver}
          onDragLeave={handleDragLeave}
          onDrop={handleDrop}
        >
          {preview ? (
            <div className={styles.previewContainer}>
              <img src={preview} alt="Preview" className={styles.preview} />
              <div className={styles.previewOverlay}>
                <Camera size={24} />
                <span>{t("avatarUpload.changeImage", "Click to change")}</span>
              </div>
            </div>
          ) : (
            <div className={styles.uploadPlaceholder}>
              <Upload size={48} />
              <p className={styles.uploadText}>
                {t("avatarUpload.dragDrop", "Drag and drop an image here")}
              </p>
              <p className={styles.uploadHint}>
                {t("avatarUpload.orClick", "or click to select a file")}
              </p>
              <p className={styles.uploadLimit}>
                {t("avatarUpload.sizeLimit", "Max size: 10MB. Formats: JPEG, PNG, GIF, WebP")}
              </p>
            </div>
          )}
          <input
            ref={fileInputRef}
            type="file"
            accept="image/*"
            onChange={handleInputChange}
            className={styles.fileInput}
          />
        </div>

        <div className={styles.actions}>
          <button className={styles.cancelButton} onClick={handleClose}>
            {t("avatarUpload.cancel", "Cancel")}
          </button>
          {!uploadedImageId ? (
            // 第一步：上传到 tmp
            <button
              className={styles.uploadButton}
              onClick={handleUpload}
              disabled={!selectedFile || uploadMutation.isPending}
            >
              {uploadMutation.isPending
                ? t("avatarUpload.uploading", "Uploading...")
                : t("avatarUpload.upload", "Upload")}
            </button>
          ) : (
            // 第二步：确认使用此图片
            <button
              className={styles.uploadButton}
              onClick={handleConfirm}
              disabled={confirmMutation.isPending}
            >
              {confirmMutation.isPending
                ? t("avatarUpload.confirming", "Confirming...")
                : t("avatarUpload.confirm", "Confirm")}
            </button>
          )}
        </div>

        {uploadMutation.isSuccess && !confirmMutation.isPending && !confirmMutation.isError && (
          <p className={styles.success}>
            {t("avatarUpload.uploadSuccess", "Image uploaded. Click Confirm to set as avatar.")}
          </p>
        )}

        {(uploadMutation.isError || confirmMutation.isError) && (
          <p className={styles.error}>
            {t("avatarUpload.error", "Failed to upload image. Please try again.")}
          </p>
        )}
      </div>
    </dialog>
  );
});

AvatarUploadModal.displayName = "AvatarUploadModal";

export default AvatarUploadModal;
