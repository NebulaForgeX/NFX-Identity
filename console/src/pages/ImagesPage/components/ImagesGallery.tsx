import { memo, useCallback, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import {
  DndContext,
  DragEndEvent,
  closestCenter,
  PointerSensor,
  useSensor,
  useSensors,
} from "@dnd-kit/core";
import {
  SortableContext,
  useSortable,
  verticalListSortingStrategy,
  arrayMove,
} from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import { GripVertical, Star, Trash2 } from "@/assets/icons/lucide";
import Button from "@/components/Button";
import {
  useDeleteUserImage,
  useSetPrimaryUserImage,
  useUpdateUserImagesDisplayOrderBatch,
} from "@/hooks";

import { buildImageUrl } from "@/utils/image";
import { showConfirm } from "@/stores/modalStore";
import type { UserImage } from "@/types";

import styles from "../styles.module.css";

interface ImagesGalleryProps {
  userId: string;
  userImages: UserImage[];
}

function sortableId(item: UserImage) {
  return item.id;
}

function SortableCard({
  userImage,
  index,
  isPrimary,
  onDelete,
  onSetPrimary,
}: {
  userImage: UserImage;
  index: number;
  isPrimary: boolean;
  onDelete: (id: string) => void;
  onSetPrimary: (id: string) => void;
}) {
  const { t } = useTranslation("ImagesPage");
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id: userImage.id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  const handleDeleteClick = useCallback(() => {
    showConfirm({
      title: t("deleteConfirmTitle", "Delete Image"),
      message: t("deleteConfirmMessage", "Are you sure you want to delete this image?"),
      onConfirm: () => onDelete(userImage.id),
    });
  }, [userImage.id, onDelete, t]);

  return (
    <div
      ref={setNodeRef}
      style={style}
      className={`${styles.imageCard} ${isDragging ? styles.imageCardDragging : ""}`}
    >
      <div className={styles.sortableHandle} {...attributes} {...listeners} title={t("dragToReorder", "Drag to reorder")}>
        <GripVertical size={20} />
      </div>
      <div className={styles.imageWrapper}>
        <img
          src={buildImageUrl(userImage.imageId)}
          alt={`Background ${index + 1}`}
          className={styles.image}
        />
        {isPrimary && (
          <span className={styles.primaryBadge}>{t("primary", "Primary")}</span>
        )}
        <div className={styles.imageOverlay}>
          {!isPrimary && (
            <button
              type="button"
              className={styles.setPrimaryButton}
              onClick={() => onSetPrimary(userImage.id)}
              title={t("setAsPrimary", "Set as primary")}
            >
              <Star size={20} />
            </button>
          )}
          <button
            type="button"
            className={styles.deleteButton}
            onClick={handleDeleteClick}
            title={t("delete", "Delete")}
          >
            <Trash2 size={20} />
          </button>
        </div>
      </div>
      <div className={styles.imageInfo}>
        <span className={styles.imageOrder}>#{index + 1}</span>
      </div>
    </div>
  );
}

export const ImagesGallery = memo(function ImagesGallery({ userId, userImages }: ImagesGalleryProps) {
  const { t } = useTranslation("ImagesPage");
  const [orderedItems, setOrderedItems] = useState<UserImage[]>(userImages);

  const deleteMutation = useDeleteUserImage();
  const setPrimaryMutation = useSetPrimaryUserImage();
  const saveOrderMutation = useUpdateUserImagesDisplayOrderBatch();

  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: { distance: 8 },
    })
  );

  // Sync from server when userImages updates (refetch after set primary, save, delete)
  useEffect(() => {
    setOrderedItems(userImages);
  }, [userImages]);

  const handleDragEnd = useCallback((event: DragEndEvent) => {
    const { active, over } = event;
    if (!over || active.id === over.id) return;
    setOrderedItems((prev) => {
      const oldIndex = prev.findIndex((i) => i.id === active.id);
      const newIndex = prev.findIndex((i) => i.id === over.id);
      if (oldIndex === -1 || newIndex === -1) return prev;
      return arrayMove(prev, oldIndex, newIndex);
    });
  }, []);

  const hasOrderChange =
    orderedItems.length === userImages.length &&
    orderedItems.some((item, i) => userImages[i]?.id !== item.id);

  const handleDelete = useCallback(
    (userImageId: string) => {
      deleteMutation.mutate(userImageId);
    },
    [deleteMutation]
  );

  const handleSetPrimary = useCallback(
    (userImageId: string) => {
      setPrimaryMutation.mutate(userImageId);
    },
    [setPrimaryMutation]
  );

  const handleSaveOrder = useCallback(() => {
    const order = orderedItems.map((item, i) => ({ id: item.id, displayOrder: i }));
    saveOrderMutation.mutate({ userId, data: { order } });
  }, [orderedItems, userId, saveOrderMutation]);

  const isSavingOrder = saveOrderMutation.isPending;

  if (userImages.length === 0) return null;

  return (
    <div className={styles.gallerySection}>
      {hasOrderChange && (
        <div className={styles.saveOrderBar}>
          <span className={styles.saveOrderText}>{t("orderChanged", "Order changed")}</span>
          <Button
            type="button"
            variant="primary"
            size="small"
            className={styles.saveOrderButton}
            onClick={handleSaveOrder}
            disabled={isSavingOrder}
            loading={isSavingOrder}
          >
            {isSavingOrder ? t("saving", "Saving…") : t("saveOrder", "Save order")}
          </Button>
        </div>
      )}
      <DndContext sensors={sensors} collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
        <SortableContext
          items={orderedItems.map(sortableId)}
          strategy={verticalListSortingStrategy}
        >
          <div className={styles.imagesGrid}>
            {orderedItems.map((userImage, index) => (
              <SortableCard
                key={userImage.id}
                userImage={userImage}
                index={index}
                isPrimary={index === 0}
                onDelete={handleDelete}
                onSetPrimary={handleSetPrimary}
              />
            ))}
          </div>
        </SortableContext>
      </DndContext>
    </div>
  );
});
