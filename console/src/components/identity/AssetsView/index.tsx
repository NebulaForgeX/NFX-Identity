import { useState } from "react";
import { Box, Button, Container, Flex, Grid, IconButton, Section, Tabs, Text, Tooltip } from "@radix-ui/themes";
import { systemEventEmitter } from "nfx-ui/events";
import { useAssetFileURL, useConfirmUpload, useDeleteAsset, useListAssets, usePrepareUpload } from "nfx-ui/hooks";
import { AnimatedIcon, CameraIcon, TrashIcon, FileDescriptionIcon, UploadIcon } from "nfx-ui/icons";
import type { Asset } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

import { EmptyState, PageHeader, Surface } from "@/components";
import { PageFrame } from "@/layouts";
import { minioUploadMessage, putToPresignedUrl, safeArray } from "@/utils";

import styles from "./s.module.css";

const KINDS: Asset.Kind[] = ["images", "files", "videos", "audios"];

export default function AssetsView() {
  const { t } = useTranslation("pages.Assets");
  const [kind, setKind] = useState<Asset.Kind>("images");
  const list = useListAssets(kind);
  const prepare = usePrepareUpload();
  const confirm = useConfirmUpload();
  const del = useDeleteAsset();
  const fileURL = useAssetFileURL();
  const rows = safeArray(list.data);
  const busy = prepare.isPending || confirm.isPending;

  const onFile = async (file: File) => {
    try {
      const prepared = await prepare.mutateAsync({
        kind,
        params: { fileName: file.name, mimeType: file.type || "application/octet-stream" },
      });
      await putToPresignedUrl(prepared.uploadUrl, file, file.type || "application/octet-stream");
      await confirm.mutateAsync({ kind, params: { id: prepared.id } });
      await list.refetch();
    } catch (err) {
      systemEventEmitter.showError(minioUploadMessage(err, t("uploadFailedNetwork"), t("uploadFailed")));
    }
  };

  const remove = (id: string) => {
    void del.mutateAsync({ kind, id }).then(() => list.refetch());
  };

  const upload = (
    <Button asChild size="3" disabled={busy}>
      <label>
        <AnimatedIcon icon={UploadIcon} size={16} />
        {t("upload")}
        <input
          type="file"
          hidden
          onChange={(e) => {
            const file = e.target.files?.[0];
            if (file) void onFile(file);
            e.target.value = "";
          }}
        />
      </label>
    </Button>
  );

  return (
    <PageFrame>
      <PageHeader icon={CameraIcon} title={t("title")} description={t("subtitle")} />
      <Grid columns={{ initial: "1", lg: "minmax(0, 1fr) minmax(0, 18rem)" }} gap="5" align="start">
        <Surface>
          <Flex direction="column" gap="5">
            <Tabs.Root value={kind} onValueChange={(v) => setKind(v as Asset.Kind)}>
              <Tabs.List size="2">
                {KINDS.map((k) => (
                  <Tabs.Trigger key={k} value={k}>
                    {t(`kinds.${k}`)}
                  </Tabs.Trigger>
                ))}
              </Tabs.List>
            </Tabs.Root>
            {list.isError ? (
              <Text size="2" color="red">
                {(list.error as Error).message}
              </Text>
            ) : null}
            {rows.length === 0 ? (
              <EmptyState icon={kind === "images" ? CameraIcon : FileDescriptionIcon} title={t("noItems")} action={upload} />
            ) : kind === "images" ? (
              <Grid columns={{ initial: "2", sm: "3", xl: "4" }} gap="3">
                {rows.map((row) => (
                  <Box key={row.id} className={styles.tile}>
                    <img src={fileURL("images", row.id)} alt="" className={styles.tileImage} />
                    <Box className={styles.overlay}>
                      <Section size="1" py="2">
                        <Container size="4" width="100%" maxWidth="100%" px="3">
                          <Flex justify="between" align="center" gap="2">
                            <Text size="1" weight="medium" truncate className={styles.overlayText}>
                              <a href={fileURL(kind, row.id)} target="_blank" rel="noreferrer">
                                {row.fileName}
                              </a>
                            </Text>
                            <Tooltip content={t("delete")}>
                              <IconButton size="1" variant="solid" color="red" aria-label={t("delete")} onClick={() => remove(row.id)}>
                                <AnimatedIcon icon={TrashIcon} size={14} />
                              </IconButton>
                            </Tooltip>
                          </Flex>
                        </Container>
                      </Section>
                    </Box>
                  </Box>
                ))}
              </Grid>
            ) : (
              <Flex direction="column" gap="2">
                {rows.map((row) => (
                  <Surface key={row.id} tone="inset" py="3" px="4">
                    <Flex justify="between" align="center" gap="3">
                      <Flex align="center" gap="3" minWidth="0">
                        <Box className={styles.fileMark}>
                          <Flex align="center" justify="center" width="100%" height="100%">
                            <AnimatedIcon icon={FileDescriptionIcon} size={16} />
                          </Flex>
                        </Box>
                        <Text size="2" weight="medium" truncate>
                          <a href={fileURL(kind, row.id)} target="_blank" rel="noreferrer">
                            {row.fileName}
                          </a>
                        </Text>
                      </Flex>
                      <Button size="1" variant="ghost" color="red" onClick={() => remove(row.id)}>
                        {t("delete")}
                      </Button>
                    </Flex>
                  </Surface>
                ))}
              </Flex>
            )}
          </Flex>
        </Surface>
        <Surface tone="hero" sticky>
          <Flex direction="column" gap="4">
            <Text size="1" color="gray" weight="medium" className={styles.kindLabel}>
              {t(`kinds.${kind}`)}
            </Text>
            <Text size="9" weight="bold" className={styles.count}>
              {rows.length}
            </Text>
            {upload}
          </Flex>
        </Surface>
      </Grid>
    </PageFrame>
  );
}
