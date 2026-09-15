import { memo, useCallback, useEffect, useState } from "react";
import { Avatar, Button, Flex, Heading, Text } from "@radix-ui/themes";
import { User } from "lucide-react";
import { PageFrame } from "nfx-ui/layouts";
import { EmptyState, PageHeader } from "nfx-ui/components";
import { useAssetRepository, useAuthRepository } from "nfx-ui/apis";
import { useAuthStore } from "nfx-ui/stores";
import { AuthIdentityProviderEnum, ProfileKindEnum } from "nfx-ui/enums";
import type { Profile } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

const ProfilePage = memo(() => {
  const { t } = useTranslation("ProfilePage");
  const auth = useAuthRepository();
  const asset = useAssetRepository();
  const accountId = useAuthStore((s) => s.currentAccountId);
  const kind = useAuthStore((s) => s.currentProfileKind) || ProfileKindEnum.FORGER;
  const [info, setInfo] = useState<Profile.Response.FullAccountInformationWithForgerProfile | Profile.Response.FullAccountInformationWithAuthorityProfile | null>(null);
  const [phoneInput, setPhoneInput] = useState("");

  const reload = useCallback(async () => {
    setInfo(await auth.GetCurrentFullAccountInformationWithProfile(kind));
    setLoaded(true);
  }, [auth, kind]);

  useEffect(() => {
    void reload();
  }, [reload]);

  const profile = info && "forgerProfile" in info ? info.forgerProfile : info && "authorityProfile" in info ? info.authorityProfile : null;
  const activeAvatar = profile?.avatars?.find((row: Profile.Response.ProfileAvatar) => row.isActive) ?? profile?.avatars?.[0];
  const background = profile?.backgrounds?.[0];
  const githubLinked = (info?.identities ?? []).some((row: Profile.Response.IdentityLink) => row.identityProvider === AuthIdentityProviderEnum.GITHUB);

  const uploadImage = async (file: File) => {
    const prepared = await asset.PrepareUpload("images", { fileName: file.name, mimeType: file.type || "image/png" });
    await fetch(prepared.uploadUrl, { method: "PUT", body: file, headers: { "Content-Type": file.type || "image/png" } });
    await asset.ConfirmUpload("images", { id: prepared.id });
    return prepared.id;
  };

  if (!accountId || !loaded || !info) {
    return (
      <PageFrame>
        <EmptyState icon={User} title={t("userNotFound")} />
      </PageFrame>
    );
  }

  return (
    <PageFrame>
      <PageHeader icon={User} title={t("title")} description={t("subtitle")} />
      <Flex direction="column" gap="5">
        <Flex align="center" gap="4">
          <Avatar size="5" src={activeAvatar ? asset.FileURL("images", activeAvatar.imageId) : undefined} fallback={profile?.displayName?.slice(0, 1) || "U"} />
          <Flex direction="column" gap="1">
            <Heading size="5">{profile?.displayName || t("user")}</Heading>
            <Text size="2" color="gray">
              {accountId}
            </Text>
          </Flex>
        </Flex>

        <Flex direction="column" gap="2">
          <Heading size="3">{t("emails")}</Heading>
          {(info.emails ?? []).map((email) => (
            <Text key={email.id} size="2">
              {email.email}
              {email.isPrimary ? ` · ${t("primary")}` : ""}
              {email.verifiedAt ? ` · ${t("verified")}` : ""}
            </Text>
          ))}
        </Flex>

        <Flex direction="column" gap="2">
          <Heading size="3">{t("phones")}</Heading>
          {(info.phones ?? []).length === 0 ? (
            <Text size="2" color="gray">
              —
            </Text>
          ) : (
            (info.phones ?? []).map((phone) => (
              <Text key={phone.id} size="2">
                {phone.phone}
                {phone.isPrimary ? ` · ${t("primary")}` : ""}
              </Text>
            ))
          )}
          <Flex gap="2">
            <input
              value={phoneInput}
              onChange={(e) => setPhoneInput(e.target.value)}
              placeholder="+86138..."
              style={{ flex: 1, minHeight: 32, borderRadius: 8, padding: "0 10px" }}
            />
            <Button
              variant="soft"
              onClick={() => {
                const phone = phoneInput.trim();
                if (!phone) return;
                void auth.CreatePhone({ phone }).then(() => {
                  setPhoneInput("");
                  return reload();
                });
              }}
            >
              {t("addPhone", { defaultValue: "Add phone" })}
            </Button>
          </Flex>
        </Flex>

        <Flex gap="3" wrap="wrap">
          <Button asChild>
            <label>
              {t("avatar", { defaultValue: "Avatar" })}
              <input
                type="file"
                accept="image/*"
                hidden
                onChange={(e) => {
                  const file = e.target.files?.[0];
                  if (!file) return;
                  void uploadImage(file).then((id) => auth.ConfirmProfileAvatar(kind, { imageId: id })).then(reload);
                  e.target.value = "";
                }}
              />
            </label>
          </Button>
          <Button asChild variant="soft">
            <label>
              {t("background", { defaultValue: "Background" })}
              <input
                type="file"
                accept="image/*"
                hidden
                onChange={(e) => {
                  const file = e.target.files?.[0];
                  if (!file) return;
                  void uploadImage(file)
                    .then((id) => auth.ConfirmProfileBackgrounds(kind, { images: [{ imageId: id, sortOrder: 0 }] }))
                    .then(reload);
                  e.target.value = "";
                }}
              />
            </label>
          </Button>
          {githubLinked ? (
            <Button
              color="red"
              variant="soft"
              onClick={() => {
                void auth.UnlinkGitHub().then(reload);
              }}
            >
              GitHub
            </Button>
          ) : (
            <Button
              variant="outline"
              onClick={() => {
                void auth.GetGitHubAuthorizeUrl().then((data) => {
                  if (data.authorizeUrl) window.location.href = data.authorizeUrl;
                });
              }}
            >
              GitHub
            </Button>
          )}
        </Flex>
        {background ? (
          <img src={asset.FileURL("images", background.imageId)} alt="" style={{ maxWidth: 480, borderRadius: 12 }} />
        ) : null}
      </Flex>
    </PageFrame>
  );
});

ProfilePage.displayName = "ProfilePage";
export default ProfilePage;
