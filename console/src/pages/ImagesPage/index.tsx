import { memo, useCallback, useEffect, useState } from "react";
import { Button, Flex, Heading, Tabs, Text } from "@radix-ui/themes";
import { Images } from "lucide-react";
import { PageFrame } from "nfx-ui/layouts";
import { PageHeader } from "nfx-ui/components";
import { useAssetRepository } from "nfx-ui/apis";
import type { Asset } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

const KINDS: Asset.Kind[] = ["images", "files", "videos", "audios"];

const AssetsPage = memo(function AssetsPage() {
  const { t } = useTranslation("ImagesPage");
  const asset = useAssetRepository();
  const [kind, setKind] = useState<Asset.Kind>("images");
  const [rows, setRows] = useState<Asset.Response.Detail[]>([]);

  const reload = useCallback(async () => {
    setRows(await asset.List(kind));
  }, [asset, kind]);

  useEffect(() => {
    void reload();
  }, [reload]);

  const onFile = async (file: File) => {
    const prepared = await asset.PrepareUpload(kind, { fileName: file.name, mimeType: file.type || "application/octet-stream" });
    await fetch(prepared.uploadUrl, { method: "PUT", body: file, headers: { "Content-Type": file.type || "application/octet-stream" } });
    await asset.ConfirmUpload(kind, { id: prepared.id });
    await reload();
  };

  return (
    <PageFrame>
      <PageHeader icon={Images} title={t("title", "Assets")} description={t("subtitle", "Images, files, videos, and audios in Stack MinIO")} />
      <Tabs.Root value={kind} onValueChange={(v) => setKind(v as Asset.Kind)}>
        <Tabs.List>
          {KINDS.map((k) => (
            <Tabs.Trigger key={k} value={k}>
              {k}
            </Tabs.Trigger>
          ))}
        </Tabs.List>
      </Tabs.Root>
      <Flex mt="4" mb="4">
        <Button asChild>
          <label>
            Upload
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
      </Flex>
      <Flex direction="column" gap="2">
        {rows.length === 0 ? <Heading size="3">No items</Heading> : null}
        {rows.map((row) => (
          <Flex key={row.id} justify="between" align="center">
            <Text size="2">{row.fileName}</Text>
            <Button
              variant="soft"
              color="red"
              onClick={() => {
                void asset.Delete(kind, row.id).then(reload);
              }}
            >
              Delete
            </Button>
          </Flex>
        ))}
      </Flex>
    </PageFrame>
  );
});

AssetsPage.displayName = "AssetsPage";
export default AssetsPage;
