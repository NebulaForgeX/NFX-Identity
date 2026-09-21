package handler

import (
	authconn "nfxidentity/connections/auth"
	asseterrs "nfxidentity/errors/src/asset"
	assetmsg "nfxidentity/messages/src/asset"
	audiosApp "nfxidentity/modules/asset/application/audios"
	"nfxidentity/modules/asset/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AudioHandler struct {
	app  *audiosApp.Service
	auth *authconn.Client
}

func NewAudioHandler(app *audiosApp.Service, auth *authconn.Client) *AudioHandler {
	return &AudioHandler{app: app, auth: auth}
}

func (h *AudioHandler) ListAudios(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	rows, err := h.app.List(c.Context(), audiosApp.ListInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: rows})
}

func (h *AudioHandler) PrepareAudioUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareAudioUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrAudioRequestBodyInvalid.WithCause(err)
	}
	out, err := h.app.PrepareUpload(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, assetmsg.UPLOAD_URL_CREATED, httpx.SuccessOptions{Data: reqdto.PrepareAudioUploadDataFrom(out)})
}

func (h *AudioHandler) PrepareAudiosUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareAudioUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrAudioRequestBodyInvalid.WithCause(err)
	}
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.app.PrepareUploads(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, assetmsg.UPLOAD_URLS_CREATED, httpx.SuccessOptions{Data: map[string]any{"results": reqdto.PrepareAudioUploadsDataFrom(out)}})
}

func (h *AudioHandler) ConfirmAudioUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmAudioUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrAudioRequestBodyInvalid.WithCause(err)
	}
	in, err := req.ToInput(accountID)
	if err != nil {
		return err
	}
	if _, err := h.app.ConfirmUpload(c.Context(), in); err != nil {
		return err
	}
	return fiberx.OK(c, assetmsg.AUDIO_CONFIRMED, httpx.SuccessOptions{})
}

func (h *AudioHandler) ConfirmAudiosUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmAudioUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrAudioRequestBodyInvalid.WithCause(err)
	}
	if err := req.Validate(); err != nil {
		return err
	}
	in, err := req.ToInput(accountID)
	if err != nil {
		return err
	}
	if _, err := h.app.ConfirmUploads(c.Context(), in); err != nil {
		return err
	}
	return fiberx.OK(c, assetmsg.AUDIOS_CONFIRMED, httpx.SuccessOptions{})
}

func (h *AudioHandler) DeleteAudio(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return asseterrs.ErrInvalidAccountID
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	audioID, err := uuid.Parse(c.Params("id"))
	if err != nil || audioID == uuid.Nil {
		return asseterrs.ErrInvalidAudioID.WithCause(err)
	}
	if err := h.app.Delete(c.Context(), audiosApp.DeleteInput{
		AccountID: accountID,
		AudioID:   audioID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, assetmsg.AUDIO_DELETED, httpx.SuccessOptions{})
}

func (h *AudioHandler) ServeAudioFile(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return asseterrs.ErrInvalidAudioID.WithCause(err)
	}
	vo, err := h.app.Get(c.Context(), audiosApp.GetInput{ID: id})
	if err != nil {
		return err
	}
	url, err := h.app.FileURL(c.Context(), audiosApp.FileURLInput{ObjectKey: vo.FilePath})
	if err != nil {
		return err
	}
	return c.Redirect().To(url)
}
