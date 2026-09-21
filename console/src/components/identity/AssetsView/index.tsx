import { CameraIcon } from "nfx-ui/icons";
import { useState } from "react";
import { Box, Button, Flex, Tabs, Text } from "@radix-ui/themes";
import { systemEventEmitter } from "nfx-ui/events";
import { useAssetFileURL, useConfirmUpload, useDeleteAsset, useListAssets, usePrepareUpload } from "nfx-ui/hooks";
import type { Asset } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

import { PageHeader } from "@/components";
import { PageFrame } from "@/layouts";
import { minioUploadMessage, putToPresignedUrl, safeArray } from "@/utils";

import { LedgerSection } from "../Ledger";
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

  return (
    <PageFrame>
      <PageHeader icon={CameraIcon} title={t("title")} description={t("subtitle")} />
      <Tabs.Root value={kind} onValueChange={(v) => setKind(v as Asset.Kind)}>
        <Tabs.List className={styles.tabList}>
          {KINDS.map((k) => (
            <Tabs.Trigger key={k} value={k}>
              {t(`kinds.${k}`)}
            </Tabs.Trigger>
          ))}
        </Tabs.List>
      </Tabs.Root>
      <LedgerSection
        title={t(`kinds.${kind}`)}
        actions={
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
        }
      >
        {list.isError ? (
          <Text size="2" color="red">
            {(list.error as Error).message}
          </Text>
        ) : null}
        <Flex direction="column" gap="2">
          {rows.length === 0 ? <Text size="2" color="gray">{t("noItems")}</Text> : null}
          {rows.map((row) => (
            <Box key={row.id} className={styles.hairline}>
              <Box py="2">
            <Flex justify="between" align="center" gap="3">
              <Flex align="center" gap="3" minWidth="0">
                {kind === "images" ? <img src={fileURL("images", row.id)} alt="" width={48} height={48} style={{ objectFit: "cover" }} /> : null}
                <Text size="2">
                  <a href={fileURL(kind, row.id)} target="_blank" rel="noreferrer">
                    {row.fileName}
                  </a>
                </Text>
              </Flex>
              <Button
                variant="outline"
                color="red"
                onClick={() => {
                  void del.mutateAsync({ kind, id: row.id }).then(() => list.refetch());
                }}
              >
                {t("delete")}
              </Button>
            </Flex>
              </Box>
            </Box>
          ))}
        </Flex>
      </LedgerSection>
    </PageFrame>
  );
}
