import { memo, useCallback, useEffect, useState } from "react";
import { Button, Flex, Heading, Tabs, Text } from "@radix-ui/themes";
import { Images } from "lucide-react";
import { PageHeader } from "@/components";
import { PageFrame } from "@/layouts";
import { useAssetRepository } from "nfx-ui/apis";
import type { Asset } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

const KINDS: Asset.Kind[] = ["images", "files", "videos", "audios"];

const AssetsPage = memo(function AssetsPage() {
  const { t } = useTranslation("ImagesPage");
  const asset = useAssetRepository();
  const [kind, setKind] = useState<Asset.Kind>("images");
  const [rows, setRows] = useState<Asset.Response.Detail[]>([]);
  const [error, setError] = useState<string | null>(null);

  const reload = useCallback(async () => {
    setError(null);
    try {
      setRows(await asset.List(kind));
    } catch (err) {
      setError((err as Error).message);
      setRows([]);
    }
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
      <PageHeader icon={Images} title={t("title")} description={t("subtitle")} />
      <Tabs.Root value={kind} onValueChange={(v) => setKind(v as Asset.Kind)}>
        <Tabs.List>
          {KINDS.map((k) => (
            <Tabs.Trigger key={k} value={k}>
              {t(`kinds.${k}`)}
            </Tabs.Trigger>
          ))}
        </Tabs.List>
      </Tabs.Root>
      <Flex mt="4" mb="4">
        <Button asChild>
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
      </Flex>
      {error ? (
        <Text size="2" color="red">
          {error}
        </Text>
      ) : null}
      <Flex direction="column" gap="2">
        {rows.length === 0 ? <Heading size="3">{t("noItems")}</Heading> : null}
        {rows.map((row) => (
          <Flex key={row.id} justify="between" align="center" gap="3">
            <Flex align="center" gap="3">
              {kind === "images" ? <img src={asset.FileURL("images", row.id)} alt="" style={{ width: 48, height: 48, objectFit: "cover", borderRadius: 8 }} /> : null}
              <Text size="2">
                <a href={asset.FileURL(kind, row.id)} target="_blank" rel="noreferrer">
                  {row.fileName}
                </a>
              </Text>
            </Flex>
            <Button
              variant="soft"
              color="red"
              onClick={() => {
                void asset.Delete(kind, row.id).then(reload);
              }}
            >
              {t("delete")}
            </Button>
          </Flex>
        ))}
      </Flex>
    </PageFrame>
  );
});

AssetsPage.displayName = "AssetsPage";
export default AssetsPage;
