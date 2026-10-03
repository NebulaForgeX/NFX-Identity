import { ArrowLeft, Camera, ChevronRight, Save, Trash2 } from "lucide-react";
import type { Profile } from "nfx-ui/types";

import { useRef } from "react";
import { Badge, Box, Button, Flex, Grid, IconButton, Text } from "@radix-ui/themes";
import { CameraIcon } from "nfx-ui/icons";
import { useTranslation } from "react-i18next";

import { EmptyState, LucideIcon } from "@/components";

import { isUserProfileBackgroundDraftBusy } from "../drafts";
import styles from "./s.module.css";
import { useUserProfileBackgroundUpload } from "../useUserProfileBackgroundUpload";

const MAX_PROFILE_BACKGROUNDS = 6;

export default function BackgroundGallery({ profile }: { profile: Profile.Response.ProfileBase }) {
  const { t } = useTranslation("pages.User.Profile.Edit");
  const fileInputRef = useRef<HTMLInputElement>(null);
  const { drafts, uploading, confirming, dirty, imageError, uploadFiles, removeDraft, moveDraft, confirmDrafts } = useUserProfileBackgroundUpload(profile, MAX_PROFILE_BACKGROUNDS);

  const completedCount = drafts.filter((d) => !isUserProfileBackgroundDraftBusy(d) && d.status !== "failed").length;
  const atLimit = drafts.filter((d) => d.status !== "failed").length >= MAX_PROFILE_BACKGROUNDS;

  return (
    <Flex direction="column" gap="4">
      <Flex align="center" justify="between" gap="4" wrap="wrap">
        <Flex minWidth="0" flexGrow="1">
          <Text size="1" color="gray">
            {t("backgroundUpload.queueSummary", {
              done: completedCount,
              total: MAX_PROFILE_BACKGROUNDS,
            })}
          </Text>
        </Flex>
        <Flex gap="2" wrap="wrap" align="center">
          <Button type="button" size="2" variant="outline" disabled={uploading || confirming || atLimit} onClick={() => fileInputRef.current?.click()}>
            <LucideIcon icon={Camera} size={14} />
            {atLimit ? t("backgroundUpload.full") : t("backgroundUpload.add")}
          </Button>
          <Button type="button" size="2" disabled={!dirty || uploading || confirming} loading={confirming} onClick={() => void confirmDrafts()}>
            <LucideIcon icon={Save} size={14} />
            {confirming ? t("backgroundUpload.confirming") : t("backgroundUpload.confirm")}
          </Button>
        </Flex>
      </Flex>

      <input
        ref={fileInputRef}
        type="file"
        accept="image/*"
        multiple
        className={styles.hiddenInput}
        onChange={(event) => {
          const files = event.target.files;
          if (files?.length) void uploadFiles(files);
          event.target.value = "";
        }}
      />

      {drafts.length ? (
        <Grid columns={{ initial: "2", sm: "3", lg: "2" }} gap="3" width="100%">
          {drafts.map((draft, index) => {
            const busy = isUserProfileBackgroundDraftBusy(draft);
            const failed = draft.status === "failed";
            return (
              <Box key={draft.imageId} className={styles.tile} data-failed={failed ? "true" : undefined}>
                <img src={draft.previewUrl} alt="" className={styles.tileImage} draggable={false} />
                <Badge size="1" variant="solid" color="gray" highContrast className={styles.order}>
                  {draft.sortOrder + 1}
                </Badge>
                {busy ? (
                  <Flex position="absolute" inset="0" align="center" justify="center" className={styles.busy}>
                    <Text size="2" weight="bold">
                      {Math.round(draft.progress ?? 0)}%
                    </Text>
                  </Flex>
                ) : null}
                {failed ? (
                  <Flex position="absolute" inset="0" align="center" justify="center" className={styles.failed}>
                    <Text size="1" weight="bold">
                      {t("backgroundUpload.status.failed")}
                    </Text>
                  </Flex>
                ) : null}
                <Flex position="absolute" right="2" bottom="2" gap="1" className={styles.tileActions}>
                  <IconButton
                    type="button"
                    size="1"
                    variant="surface"
                    color="gray"
                    disabled={busy || failed || index === 0}
                    onClick={() => moveDraft(draft.imageId, -1)}
                    aria-label={t("backgroundUpload.moveLeft")}
                  >
                    <LucideIcon icon={ArrowLeft} size={12} />
                  </IconButton>
                  <IconButton
                    type="button"
                    size="1"
                    variant="surface"
                    color="gray"
                    disabled={busy || failed || index === drafts.length - 1}
                    onClick={() => moveDraft(draft.imageId, 1)}
                    aria-label={t("backgroundUpload.moveRight")}
                  >
                    <LucideIcon icon={ChevronRight} size={12} />
                  </IconButton>
                  <IconButton type="button" size="1" variant="solid" color="red" disabled={busy} onClick={() => removeDraft(draft.imageId)} aria-label={t("backgroundUpload.remove")}>
                    <LucideIcon icon={Trash2} size={12} />
                  </IconButton>
                </Flex>
              </Box>
            );
          })}
        </Grid>
      ) : (
        <EmptyState icon={CameraIcon} title={t("backgroundUpload.dropTitle")} description={t("backgroundUpload.hint")} />
      )}

      {drafts.length > 1 ? (
        <Text size="1" color="gray">
          {t("backgroundUpload.reorderHint")}
        </Text>
      ) : null}
      {imageError ? (
        <Text size="1" color="red">
          {imageError}
        </Text>
      ) : null}
    </Flex>
  );
}
