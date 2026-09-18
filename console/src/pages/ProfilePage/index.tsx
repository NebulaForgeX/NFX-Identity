import { memo, useCallback, useEffect, useState } from "react";
import { Avatar, Button, Flex, Heading, Select, Switch, Text, TextField } from "@radix-ui/themes";
import { User } from "lucide-react";
import { PageFrame } from "nfx-ui/layouts";
import { EmptyState, PageHeader } from "nfx-ui/components";
import { useAssetRepository, useAuthRepository } from "nfx-ui/apis";
import { useAuthStore } from "nfx-ui/stores";
import { AuthIdentityProviderEnum, LANGUAGE_VALUES, LanguageEnum, ProfileKindEnum } from "nfx-ui/enums";
import type { Profile } from "nfx-ui/types";
import { useTranslation } from "react-i18next";

const ProfilePage = memo(() => {
  const { t } = useTranslation("ProfilePage");
  const auth = useAuthRepository();
  const asset = useAssetRepository();
  const accountId = useAuthStore((s) => s.currentAccountId);
  const kind = useAuthStore((s) => s.currentProfileKind) || ProfileKindEnum.FORGER;
  const [info, setInfo] = useState<Profile.Response.FullAccountInformationWithForgerProfile | Profile.Response.FullAccountInformationWithAuthorityProfile | null>(null);
  const [loaded, setLoaded] = useState(false);
  const [emailInput, setEmailInput] = useState("");
  const [phoneInput, setPhoneInput] = useState("");
  const [verifyById, setVerifyById] = useState<Record<string, string>>({});
  const [emailEditById, setEmailEditById] = useState<Record<string, string>>({});
  const [displayName, setDisplayName] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [city, setCity] = useState("");
  const [country, setCountry] = useState("");
  const [website, setWebsite] = useState("");
  const [bio, setBio] = useState("");
  const [gender, setGender] = useState("");
  const [birthday, setBirthday] = useState("");
  const [timezone, setTimezone] = useState("");
  const [profileLanguage, setProfileLanguage] = useState<LanguageEnum>(LanguageEnum.ZH);

  const reload = useCallback(async () => {
    const next = await auth.GetCurrentFullAccountInformationWithProfile(kind);
    setInfo(next);
    const profile = "forgerProfile" in next ? next.forgerProfile : next.authorityProfile;
    setDisplayName(profile?.displayName ?? "");
    setFirstName(profile?.firstName ?? "");
    setLastName(profile?.lastName ?? "");
    setCity(profile?.city ?? "");
    setCountry(profile?.country ?? "");
    setWebsite(profile?.website ?? "");
    setBio(profile?.bio ?? "");
    setGender(profile?.gender ?? "");
    setBirthday(profile?.birthday ? String(profile.birthday).slice(0, 10) : "");
    setTimezone(profile?.timezone ?? "");
    setProfileLanguage((profile?.profileLanguage as LanguageEnum | undefined) ?? LanguageEnum.ZH);
    setLoaded(true);
  }, [auth, kind]);

  useEffect(() => {
    void reload();
  }, [reload]);

  const profile = info && "forgerProfile" in info ? info.forgerProfile : info && "authorityProfile" in info ? info.authorityProfile : null;
  const activeAvatar = profile?.avatars?.find((row: Profile.Response.ProfileAvatar) => row.isActive) ?? profile?.avatars?.[0];
  const background = profile?.backgrounds?.[0];
  const githubLinked = (info?.identities ?? []).some((row: Profile.Response.IdentityLink) => row.identityProvider === AuthIdentityProviderEnum.GITHUB);
  const loginNotification = profile?.settings?.loginNotification ?? false;

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
          <Heading size="3">{t("basicInfo")}</Heading>
          <TextField.Root value={displayName} onChange={(e) => setDisplayName(e.target.value)} placeholder={t("displayName")} />
          <Flex gap="2">
            <TextField.Root value={firstName} onChange={(e) => setFirstName(e.target.value)} placeholder={t("firstName")} />
            <TextField.Root value={lastName} onChange={(e) => setLastName(e.target.value)} placeholder={t("lastName")} />
          </Flex>
          <Flex gap="2">
            <TextField.Root value={city} onChange={(e) => setCity(e.target.value)} placeholder={t("city")} />
            <TextField.Root value={country} onChange={(e) => setCountry(e.target.value)} placeholder={t("country")} />
          </Flex>
          <TextField.Root value={website} onChange={(e) => setWebsite(e.target.value)} placeholder={t("website")} />
          <Flex gap="2">
            <TextField.Root value={gender} onChange={(e) => setGender(e.target.value)} placeholder={t("gender")} />
            <TextField.Root type="date" value={birthday} onChange={(e) => setBirthday(e.target.value)} />
          </Flex>
          <TextField.Root value={timezone} onChange={(e) => setTimezone(e.target.value)} placeholder={t("timezone")} />
          <Select.Root value={profileLanguage} onValueChange={(value) => setProfileLanguage(value as LanguageEnum)}>
            <Select.Trigger placeholder={t("profileLanguage")} />
            <Select.Content>
              {LANGUAGE_VALUES.map((lang) => (
                <Select.Item key={lang} value={lang}>
                  {lang}
                </Select.Item>
              ))}
            </Select.Content>
          </Select.Root>
          <TextField.Root value={bio} onChange={(e) => setBio(e.target.value)} placeholder={t("bio")} />
          <Button
            onClick={() => {
              void auth
                .PatchProfile(kind, {
                  profileLanguage,
                  displayName: displayName.trim() || null,
                  firstName: firstName.trim() || null,
                  lastName: lastName.trim() || null,
                  city: city.trim() || null,
                  country: country.trim() || null,
                  website: website.trim() || null,
                  bio: bio.trim() || null,
                  gender: gender.trim() || null,
                  birthday: birthday.trim() || null,
                  timezone: timezone.trim() || null,
                })
                .then(reload);
            }}
          >
            {t("submit")}
          </Button>
        </Flex>

        <Flex align="center" gap="2">
          <Switch
            checked={loginNotification}
            onCheckedChange={(checked) => {
              void auth.PatchProfileSettings(kind, { loginNotification: checked }).then(reload);
            }}
          />
          <Text size="2">{t("loginNotification", { defaultValue: "Login notification" })}</Text>
        </Flex>

        <Flex direction="column" gap="2">
          <Heading size="3">{t("emails")}</Heading>
          {(info.emails ?? []).map((email) => (
            <Flex key={email.id} gap="2" align="center" wrap="wrap">
              <Text size="2">
                {email.email}
                {email.isPrimary ? ` · ${t("primary")}` : ""}
                {email.verifiedAt ? ` · ${t("verified")}` : ` · ${t("unverified")}`}
              </Text>
              {!email.verifiedAt ? (
                <>
                  <TextField.Root
                    size="1"
                    value={verifyById[email.id] ?? ""}
                    onChange={(e) => setVerifyById((cur) => ({ ...cur, [email.id]: e.target.value }))}
                    placeholder={t("verificationCode", { defaultValue: "Code" })}
                  />
                  <Button
                    size="1"
                    variant="soft"
                    onClick={() => {
                      void auth.SendEmailVerificationCode(email.id, { lang: LanguageEnum.ZH }).then(() => undefined);
                    }}
                  >
                    {t("sendCode", { defaultValue: "Send code" })}
                  </Button>
                  <Button
                    size="1"
                    onClick={() => {
                      const code = (verifyById[email.id] ?? "").trim();
                      if (!code) return;
                      void auth.VerifyEmail(email.id, { verificationCode: code }).then(reload);
                    }}
                  >
                    {t("verified")}
                  </Button>
                </>
              ) : null}
              {!email.isPrimary ? (
                <Button size="1" variant="soft" onClick={() => void auth.SetPrimaryEmail(email.id).then(reload)}>
                  {t("primary")}
                </Button>
              ) : null}
              <Button size="1" color="red" variant="soft" onClick={() => void auth.DeleteEmail(email.id).then(reload)}>
                {t("cancel")}
              </Button>
              <TextField.Root
                size="1"
                value={emailEditById[email.id] ?? email.email}
                onChange={(e) => setEmailEditById((cur) => ({ ...cur, [email.id]: e.target.value }))}
              />
              <Button
                size="1"
                variant="soft"
                onClick={() => {
                  const next = (emailEditById[email.id] ?? "").trim();
                  if (!next || next === email.email) return;
                  void auth.UpdateEmail(email.id, { email: next }).then(reload);
                }}
              >
                {t("updateEmail")}
              </Button>
            </Flex>
          ))}
          <Flex gap="2">
            <TextField.Root value={emailInput} onChange={(e) => setEmailInput(e.target.value)} placeholder={t("email")} />
            <Button
              variant="soft"
              onClick={() => {
                const email = emailInput.trim();
                if (!email) return;
                void auth.CreateEmail({ email }).then(() => {
                  setEmailInput("");
                  return reload();
                });
              }}
            >
              {t("addEmail", { defaultValue: "Add email" })}
            </Button>
          </Flex>
        </Flex>

        <Flex direction="column" gap="2">
          <Heading size="3">{t("phones")}</Heading>
          {(info.phones ?? []).length === 0 ? (
            <Text size="2" color="gray">
              —
            </Text>
          ) : (
            (info.phones ?? []).map((phone) => (
              <Flex key={phone.id} gap="2" align="center" wrap="wrap">
                <Text size="2">
                  {phone.phone}
                  {phone.isPrimary ? ` · ${t("primary")}` : ""}
                  {phone.verifiedAt ? ` · ${t("verified")}` : ` · ${t("unverified")}`}
                </Text>
                {!phone.isPrimary ? (
                  <Button size="1" variant="soft" onClick={() => void auth.SetPrimaryPhone(phone.id).then(reload)}>
                    {t("primary")}
                  </Button>
                ) : null}
                <Button size="1" color="red" variant="soft" onClick={() => void auth.DeletePhone(phone.id).then(reload)}>
                  {t("cancel")}
                </Button>
              </Flex>
            ))
          )}
          <Text size="1" color="gray">
            {t("phoneSmsUnimplemented", { defaultValue: "SMS verification is not available yet. Numbers can still be stored as primary contact." })}
          </Text>
          <Flex gap="2">
            <TextField.Root value={phoneInput} onChange={(e) => setPhoneInput(e.target.value)} placeholder="+86138..." />
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
                  void uploadImage(file)
                    .then((id) => auth.ConfirmProfileAvatar(kind, { imageId: id }))
                    .then(reload);
                  e.target.value = "";
                }}
              />
            </label>
          </Button>
          {activeAvatar ? (
            <Button variant="soft" color="red" onClick={() => void auth.ClearProfileAvatar(kind).then(reload)}>
              {t("clearAvatar", { defaultValue: "Clear avatar" })}
            </Button>
          ) : null}
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
            <Button color="red" variant="soft" onClick={() => void auth.UnlinkGitHub().then(reload)}>
              {t("unlinkGitHub", { defaultValue: "Unlink GitHub" })}
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
              {t("linkGitHub", { defaultValue: "Link GitHub" })}
            </Button>
          )}
        </Flex>
        {background ? <img src={asset.FileURL("images", background.imageId)} alt="" style={{ maxWidth: 480, borderRadius: 12 }} /> : null}
      </Flex>
    </PageFrame>
  );
});

ProfilePage.displayName = "ProfilePage";
export default ProfilePage;
