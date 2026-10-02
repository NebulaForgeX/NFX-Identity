import { CameraIcon } from "nfx-ui/icons";
import { useState } from "react";
import { Button, Container, Flex, Grid, Section, Tabs, Text } from "@radix-ui/themes";
import { systemEventEmitter } from "nfx-ui/events";
import { useAssetFileURL, useConfirmUpload, useDeleteAsset, useListAssets, usePrepareUpload } from "nfx-ui/hooks";
import type { Asset } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

import { EmptyState, PageHeader } from "@/components";
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

  const upload = (
    <Button asChild disabled={busy}>
      <label>
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
      <Grid columns={{ initial: "1", lg: "minmax(0, 1fr) 16rem" }} gap="6" align="start">
        <Flex direction="column" gap="4" minWidth="0">
          <Tabs.Root value={kind} onValueChange={(v) => setKind(v as Asset.Kind)}>
            <Section size="1" py="0" className={styles.tabBar}>
              <Tabs.List className={styles.tabListGap}>
                {KINDS.map((k) => (
                  <Tabs.Trigger key={k} value={k}>
                    {t(`kinds.${k}`)}
                  </Tabs.Trigger>
                ))}
              </Tabs.List>
            </Section>
          </Tabs.Root>
          {list.isError ? (
            <Text size="2">
              {(list.error as Error).message}
            </Text>
          ) : null}
          {rows.length === 0 ? (
            <EmptyState icon={CameraIcon} title={t("noItems")} action={upload} />
          ) : kind === "images" ? (
            <Grid columns={{ initial: "2", sm: "3" }} gap="3">
              {rows.map((row) => (
                <Section key={row.id} size="1" py="3" className={styles.tile}>
                  <Container size="2" px="3">
                    <Flex direction="column" gap="3">
                      <img src={fileURL("images", row.id)} alt="" className={styles.tileImage} />
                      <Flex justify="between" align="center" gap="2">
                        <Text size="1" truncate>
                          <a href={fileURL(kind, row.id)} target="_blank" rel="noreferrer">
                            {row.fileName}
                          </a>
                        </Text>
                        <Button
                          size="1"
                          variant="outline"
                          onClick={() => {
                            void del.mutateAsync({ kind, id: row.id }).then(() => list.refetch());
                          }}
                        >
                          {t("delete")}
                        </Button>
                      </Flex>
                    </Flex>
                  </Container>
                </Section>
              ))}
            </Grid>
          ) : (
            <Flex direction="column" gap="2">
              {rows.map((row) => (
                <Section key={row.id} size="1" py="2" className={styles.hairline}>
                  <Flex justify="between" align="center" gap="3">
                    <Text size="2" truncate>
                      <a href={fileURL(kind, row.id)} target="_blank" rel="noreferrer">
                        {row.fileName}
                      </a>
                    </Text>
                    <Button
                      variant="outline"
                      onClick={() => {
                        void del.mutateAsync({ kind, id: row.id }).then(() => list.refetch());
                      }}
                    >
                      {t("delete")}
                    </Button>
                  </Flex>
                </Section>
              ))}
            </Flex>
          )}
        </Flex>
        <Section size="1" py="4" className={styles.side}>
          <Container size="2" px="4">
            <Flex direction="column" gap="3">
              <Text size="1" color="gray">
                {t(`kinds.${kind}`)}
              </Text>
              <Text size="7" className={styles.count}>
                {rows.length}
              </Text>
              {upload}
            </Flex>
          </Container>
        </Section>
      </Grid>
    </PageFrame>
  );
}
