import { PenIcon, UploadIcon } from "nfx-ui/icons";
import { useEffect, useRef, useState } from "react";
import { Avatar, Button, Flex, Select, Text, TextArea, TextField } from "@radix-ui/themes";
import { LanguageEnum } from "nfx-ui/enums";
import { systemEventEmitter } from "nfx-ui/events";
import {
  useClearProfileAvatar,
  useConfirmImageUpload,
  useConfirmProfileAvatar,
  useCurrentProfile,
  useDeleteImage,
  usePatchProfile,
  usePrepareImageUpload,
} from "nfx-ui/hooks";
import { buildUserProfileEditDefaults, useInitUserProfileEditForm } from "nfx-ui/schemas";
import type { Profile } from "nfx-ui/types";
import { Controller } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { LucideIcon, PageHeader } from "@/components";
import { PageFrame } from "@/layouts";
import { buildImageUrl, buildProfilePatch, compressImage, getApiErrorMessage, isEmptyPatch, resolveAccountInitial, safeNullable } from "@/utils";

import BackgroundGallery from "./backgrounds/BackgroundGallery";
import { FieldList, FieldRow, LedgerSection } from "./Ledger";

function AvatarSection({ profile, accountId }: { profile: Profile.Response.ProfileBase; accountId: Nullable<string> }) {
  const { t } = useTranslation("pages.User.Profile.Edit");
  const prepareUpload = usePrepareImageUpload();
  const confirmUpload = useConfirmImageUpload();
  const confirmAvatar = useConfirmProfileAvatar();
  const clearAvatar = useClearProfileAvatar();
  const deleteImage = useDeleteImage({ ifShowError: false });
  const fileRef = useRef<HTMLInputElement>(null);
  const [previewUrl, setPreviewUrl] = useState<Nullable<string>>(null);
  const [pendingImageId, setPendingImageId] = useState<Nullable<string>>(null);

  const initial = resolveAccountInitial(profile.displayName, accountId);
  const currentAvatarId = safeNullable(profile.avatars?.[0]?.imageId);
  const busy = prepareUpload.isPending || confirmUpload.isPending;

  const handleFile = async (file: File) => {
    if (!file.type.startsWith("image/")) {
      systemEventEmitter.showError(t("avatar.invalidType"));
      return;
    }
    if (pendingImageId) deleteImage.mutate(pendingImageId);
    if (previewUrl) URL.revokeObjectURL(previewUrl);
    setPreviewUrl(URL.createObjectURL(file));
    setPendingImageId(null);
    try {
      const compressed = await compressImage(file);
      const prep = await prepareUpload.mutateAsync({
        fileName: compressed.name,
        mimeType: compressed.type || "image/png",
      });
      const putRes = await fetch(prep.uploadUrl, {
        method: "PUT",
        body: compressed,
        headers: { "Content-Type": compressed.type || "image/png" },
      });
      if (!putRes.ok) throw new Error(`upload ${putRes.status}`);
      setPendingImageId(prep.id);
    } catch (err) {
      systemEventEmitter.showError(getApiErrorMessage(err, t("avatar.uploadFailed")));
      setPendingImageId(null);
    }
  };

  const handleConfirm = async () => {
    if (!pendingImageId) {
      systemEventEmitter.showError(t("avatar.noImage"));
      return;
    }
    try {
      await confirmUpload.mutateAsync({ id: pendingImageId });
      await confirmAvatar.mutateAsync({ imageId: pendingImageId });
      systemEventEmitter.showSuccess(t("avatar.success"));
      if (previewUrl) URL.revokeObjectURL(previewUrl);
      setPreviewUrl(null);
      setPendingImageId(null);
    } catch (err) {
      systemEventEmitter.showError(getApiErrorMessage(err, t("avatar.confirmFailed")));
    }
  };

  const src = previewUrl || (currentAvatarId ? buildImageUrl(currentAvatarId) : undefined);

  return (
    <LedgerSection title={t("avatar.title")} description={t("avatar.hint")}>
      <Flex align="center" justify="between" gap="3" wrap="wrap">
        <Flex align="center" gap="3" minWidth="0">
          <Avatar size="5" radius="none" src={src} fallback={initial} />
          <Text size="2" color="gray">
            {t("avatar.pickHint")}
          </Text>
        </Flex>
        <Flex gap="2" wrap="wrap">
          <Button size="2" variant="soft" disabled={busy} onClick={() => fileRef.current?.click()}>
            <LucideIcon icon={UploadIcon} size={14} />
            {busy ? t("avatar.uploading") : t("avatar.choose")}
          </Button>
          <Button size="2" disabled={!pendingImageId || busy} onClick={() => void handleConfirm()}>
            {confirmUpload.isPending ? t("avatar.confirming") : t("avatar.confirm")}
          </Button>
          <Button size="2" variant="soft" color="red" disabled={!currentAvatarId || busy} onClick={() => clearAvatar.mutate()}>
            {t("avatar.clear")}
          </Button>
        </Flex>
      </Flex>
      <input
        ref={fileRef}
        type="file"
        accept="image/*"
        hidden
        onChange={(e) => {
          const file = e.target.files?.[0];
          if (file) void handleFile(file);
          e.target.value = "";
        }}
      />
    </LedgerSection>
  );
}

function ProfileFields({ profile }: { profile: Profile.Response.ProfileBase }) {
  const { t } = useTranslation("pages.User.Profile.Edit");
  const form = useInitUserProfileEditForm(profile);
  const patch = usePatchProfile();

  useEffect(() => {
    form.reset(buildUserProfileEditDefaults(profile));
  }, [form, profile]);

  return (
    <LedgerSection
      title={t("sections.basics.title")}
      description={t("sections.basics.description")}
      actions={
        <Button
          size="2"
          loading={patch.isPending}
          onClick={form.handleSubmit((values) => {
            const body = buildProfilePatch(profile, values);
            if (isEmptyPatch(body)) return;
            patch.mutate(body);
          })}
        >
          {t("actions.saveChanges")}
        </Button>
      }
    >
      <FieldList>
        <FieldRow label={t("labels.displayName")}>
          <Controller
            name="displayName"
            control={form.control}
            render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />}
          />
        </FieldRow>
        <FieldRow label={t("labels.firstName")}>
          <Controller name="firstName" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.lastName")}>
          <Controller name="lastName" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.gender")}>
          <Controller name="gender" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.birthday")}>
          <Controller name="birthday" control={form.control} render={({ field }) => <TextField.Root size="2" type="date" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.city")}>
          <Controller name="city" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.country")}>
          <Controller name="country" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.website")}>
          <Controller name="website" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.timezone")}>
          <Controller name="timezone" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
        </FieldRow>
        <FieldRow label={t("labels.profileLanguage")}>
          <Controller
            name="profileLanguage"
            control={form.control}
            render={({ field }) => (
              <Select.Root value={field.value} onValueChange={field.onChange}>
                <Select.Trigger />
                <Select.Content>
                  <Select.Item value={LanguageEnum.EN}>{t("labels.langEn")}</Select.Item>
                  <Select.Item value={LanguageEnum.ZH}>{t("labels.langZh")}</Select.Item>
                  <Select.Item value={LanguageEnum.FR}>{t("labels.langFr")}</Select.Item>
                </Select.Content>
              </Select.Root>
            )}
          />
        </FieldRow>
        <FieldRow label={t("labels.bio")}>
          <Controller name="bio" control={form.control} render={({ field }) => <TextArea size="2" rows={5} value={field.value} onChange={field.onChange} />} />
        </FieldRow>
      </FieldList>
    </LedgerSection>
  );
}

export default function EditView() {
  const { t } = useTranslation("pages.User.Profile.Edit");
  const { profile, data } = useCurrentProfile();
  const accountId = safeNullable(data?.account.id);

  return (
    <PageFrame>
      <PageHeader icon={PenIcon} title={t("title")} description={t("description")} />
      {profile ? (
        <>
          <AvatarSection profile={profile} accountId={accountId} />
          <LedgerSection title={t("backgroundUpload.label")}>
            <BackgroundGallery profile={profile} />
          </LedgerSection>
          <ProfileFields profile={profile} />
        </>
      ) : (
        <Text size="2" color="gray">
          {t("empty.description")}
        </Text>
      )}
    </PageFrame>
  );
}
