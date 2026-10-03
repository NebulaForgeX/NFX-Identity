import { PenIcon } from "nfx-ui/icons";
import { Upload } from "lucide-react";
import type { ReactNode } from "react";

import { useEffect, useRef, useState } from "react";
import { CalendarIcon } from "@radix-ui/react-icons";
import { Avatar, Badge, Button, Flex, Grid, IconButton, Select, Text, TextArea, TextField } from "@radix-ui/themes";
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

import { ActionBar, EmptyState, LucideIcon, PageHeader, Surface } from "@/components";
import { showDateTimePickerModal } from "@/stores/modal";
import { PageFrame } from "@/layouts";
import { buildImageUrl, buildProfilePatch, compressImage, getApiErrorMessage, getCommandMessage, isEmptyPatch, minioUploadMessage, putToPresignedUrl, resolveAccountInitial, safeNullable } from "@/utils";

import BackgroundGallery from "../backgrounds/BackgroundGallery";
import { LedgerSection } from "../Ledger";

import styles from "./s.module.css";

function parseBirthday(value: string): boolean {
  const match = /^(\d{4})-(\d{2})-(\d{2})$/.exec(value);
  if (!match) return false;
  const year = Number(match[1]);
  const month = Number(match[2]);
  const day = Number(match[3]);
  const date = new Date(year, month - 1, day);
  return date.getFullYear() === year && date.getMonth() === month - 1 && date.getDate() === day;
}

function BirthdayField({ value, onChange, onBlur }: { value: string; onChange: (value: string) => void; onBlur: () => void }) {
  const { t } = useTranslation("pages.User.Profile.Edit");
  const displayValue = parseBirthday(value) ? value.slice(0, 10) : "";

  const openPicker = () => {
    const maxDate = new Date();
    maxDate.setHours(0, 0, 0, 0);
    showDateTimePickerModal({
      value: displayValue,
      title: t("datePicker.title"),
      minDate: new Date(maxDate.getFullYear() - 120, 0, 1),
      maxDate,
      allowClear: true,
      onConfirm: (next) => {
        onChange(next);
        onBlur();
      },
      onCancel: onBlur,
    });
  };

  return (
    <TextField.Root
      size="2"
      readOnly
      value={displayValue}
      placeholder={t("datePicker.placeholder")}
      className={styles.dateField}
      onClick={openPicker}
      onKeyDown={(event) => {
        if (event.key === "Enter" || event.key === " ") {
          event.preventDefault();
          openPicker();
        }
      }}
    >
      <TextField.Slot side="right">
        <IconButton
          type="button"
          size="1"
          variant="ghost"
          aria-label={t("datePicker.open")}
          onClick={(event) => {
            event.preventDefault();
            event.stopPropagation();
            openPicker();
          }}
        >
          <CalendarIcon width="14" height="14" />
        </IconButton>
      </TextField.Slot>
    </TextField.Root>
  );
}

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
      await putToPresignedUrl(prep.uploadUrl, compressed, compressed.type || "image/png");
      setPendingImageId(prep.id);
    } catch (err) {
      systemEventEmitter.showError(minioUploadMessage(err, t("avatar.uploadFailedNetwork"), getApiErrorMessage(err, t("avatar.uploadFailed"))));
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
      systemEventEmitter.showSuccess(getCommandMessage("USER_PROFILE_AVATAR_UPDATED", t("avatar.success")));
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
      <Flex direction="column" gap="4">
        <Surface tone="inset" py="4" px="4">
          <Flex align="center" gap="4" minWidth="0">
            <Avatar size="6" className={styles.avatar} data-pending={pendingImageId ? "true" : undefined} src={src} fallback={initial} />
            <Text size="2" color="gray">
              {t("avatar.pickHint")}
            </Text>
          </Flex>
        </Surface>
        <Flex gap="2" wrap="wrap">
          <Button size="2" variant="outline" disabled={busy} onClick={() => fileRef.current?.click()}>
            <LucideIcon icon={Upload} size={14} />
            {busy ? t("avatar.uploading") : t("avatar.choose")}
          </Button>
          <Button size="2" disabled={!pendingImageId || busy} onClick={() => void handleConfirm()}>
            {confirmUpload.isPending ? t("avatar.confirming") : t("avatar.confirm")}
          </Button>
          <Button size="2" variant="ghost" color="red" disabled={!currentAvatarId || busy} onClick={() => clearAvatar.mutate()}>
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

const GENDERS = ["female", "male", "nonbinary"] as const;
const TIMEZONES = ["UTC", "America/Vancouver", "America/Los_Angeles", "America/New_York", "Europe/London", "Europe/Paris", "Asia/Shanghai", "Asia/Tokyo"] as const;

function StackField({ label, children, error }: { label: string; children: ReactNode; error?: string }) {
  return (
    <Flex direction="column" gap="2">
      <Text as="label" size="2" weight="medium">
        {label}
      </Text>
      {children}
      {error ? (
        <Text size="1" color="red">
          {error}
        </Text>
      ) : null}
    </Flex>
  );
}

function genderLabel(t: (key: string) => string, value: string) {
  if (value === "female") return t("labels.genderFemale");
  if (value === "male") return t("labels.genderMale");
  if (value === "nonbinary") return t("labels.genderNonbinary");
  return value;
}

function ProfileFields({ profile }: { profile: Profile.Response.ProfileBase }) {
  const { t } = useTranslation("pages.User.Profile.Edit");
  const form = useInitUserProfileEditForm(profile);
  const patch = usePatchProfile();

  useEffect(() => {
    form.reset(buildUserProfileEditDefaults(profile));
  }, [form, profile]);

  const save = (
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
  );
  const genderValue = form.watch("gender");
  const timezoneValue = form.watch("timezone");
  const genderOptions = genderValue && !GENDERS.includes(genderValue as (typeof GENDERS)[number]) ? [genderValue, ...GENDERS] : [...GENDERS];
  const timezoneOptions = timezoneValue && !TIMEZONES.includes(timezoneValue as (typeof TIMEZONES)[number]) ? [timezoneValue, ...TIMEZONES] : [...TIMEZONES];

  return (
    <>
      <LedgerSection title={t("sections.identity.title")} description={t("sections.identity.description")}>
        <Grid columns={{ initial: "1", sm: "2" }} gap="4">
          <StackField label={t("labels.displayName")}>
            <Controller name="displayName" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
          </StackField>
          <StackField label={t("labels.profileLanguage")}>
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
          </StackField>
          <StackField label={t("labels.firstName")}>
            <Controller name="firstName" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
          </StackField>
          <StackField label={t("labels.lastName")}>
            <Controller name="lastName" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
          </StackField>
        </Grid>
      </LedgerSection>
      <LedgerSection title={t("sections.place.title")} description={t("sections.place.description")}>
        <Grid columns={{ initial: "1", sm: "2" }} gap="4">
          <StackField label={t("labels.city")}>
            <Controller name="city" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
          </StackField>
          <StackField label={t("labels.country")}>
            <Controller name="country" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
          </StackField>
          <StackField label={t("labels.timezone")}>
            <Controller
              name="timezone"
              control={form.control}
              render={({ field }) => (
                <Select.Root value={field.value || "UTC"} onValueChange={field.onChange}>
                  <Select.Trigger />
                  <Select.Content>
                    {timezoneOptions.map((zone) => (
                      <Select.Item key={zone} value={zone}>
                        {zone}
                      </Select.Item>
                    ))}
                  </Select.Content>
                </Select.Root>
              )}
            />
          </StackField>
          <StackField label={t("labels.website")} error={form.formState.errors.website?.message}>
            <Controller name="website" control={form.control} render={({ field }) => <TextField.Root size="2" value={field.value} onChange={field.onChange} />} />
          </StackField>
        </Grid>
      </LedgerSection>
      <LedgerSection title={t("sections.personal.title")} description={t("sections.personal.description")}>
        <Flex direction="column" gap="4">
          <Grid columns={{ initial: "1", sm: "2" }} gap="4">
            <StackField label={t("labels.gender")}>
              <Controller
                name="gender"
                control={form.control}
                render={({ field }) => (
                  <Select.Root value={field.value || "unspecified"} onValueChange={(value) => field.onChange(value === "unspecified" ? "" : value)}>
                    <Select.Trigger />
                    <Select.Content>
                      <Select.Item value="unspecified">{t("labels.genderUnspecified")}</Select.Item>
                      {genderOptions.map((value) => (
                        <Select.Item key={value} value={value}>
                          {genderLabel(t, value)}
                        </Select.Item>
                      ))}
                    </Select.Content>
                  </Select.Root>
                )}
              />
            </StackField>
            <StackField label={t("labels.birthday")}>
              <Controller name="birthday" control={form.control} render={({ field }) => <BirthdayField value={field.value} onChange={field.onChange} onBlur={field.onBlur} />} />
            </StackField>
          </Grid>
          <StackField label={t("labels.bio")}>
            <Controller name="bio" control={form.control} render={({ field }) => <TextArea size="2" rows={5} value={field.value} onChange={field.onChange} />} />
          </StackField>
        </Flex>
      </LedgerSection>
      <ActionBar
        status={
          <Badge size="2" variant="surface" color={form.formState.isDirty ? "amber" : "gray"}>
            {form.formState.isDirty ? t("status.dirty") : t("status.clean")}
          </Badge>
        }
      >
        <Button size="2" variant="ghost" color="gray" disabled={!form.formState.isDirty || patch.isPending} onClick={() => form.reset(buildUserProfileEditDefaults(profile))}>
          {t("actions.discard")}
        </Button>
        {save}
      </ActionBar>
    </>
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
        <Grid columns={{ initial: "1", lg: "minmax(0, 1.5fr) minmax(0, 1fr)" }} gap="5" align="start">
          <Flex direction="column" gap="5" minWidth="0">
            <ProfileFields profile={profile} />
          </Flex>
          <Flex direction="column" gap="5" minWidth="0">
            <AvatarSection profile={profile} accountId={accountId} />
            <LedgerSection title={t("backgroundUpload.label")} description={t("backgroundUpload.hint")}>
              <BackgroundGallery profile={profile} />
            </LedgerSection>
          </Flex>
        </Grid>
      ) : (
        <EmptyState icon={PenIcon} title={t("empty.title")} description={t("empty.description")} />
      )}
    </PageFrame>
  );
}
