package handler

import (
	authconn "nfxidentity/connections/auth"
	asseterrs "nfxidentity/errors/src/asset"
	videosApp "nfxidentity/modules/asset/application/videos"
	"nfxidentity/modules/asset/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type VideoHandler struct {
	app  *videosApp.Service
	auth *authconn.Client
}

func NewVideoHandler(app *videosApp.Service, auth *authconn.Client) *VideoHandler {
	return &VideoHandler{app: app, auth: auth}
}

func (h *VideoHandler) ListVideos(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	rows, err := h.app.List(c.Context(), videosApp.ListInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: rows})
}

func (h *VideoHandler) PrepareVideoUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareVideoUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrVideoRequestBodyInvalid.WithCause(err)
	}
	out, err := h.app.PrepareUpload(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Upload URL created", httpx.SuccessOptions{Data: reqdto.PrepareVideoUploadDataFrom(out)})
}

func (h *VideoHandler) PrepareVideosUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareVideoUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrVideoRequestBodyInvalid.WithCause(err)
	}
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.app.PrepareUploads(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Upload URLs created", httpx.SuccessOptions{Data: map[string]any{"results": reqdto.PrepareVideoUploadsDataFrom(out)}})
}

func (h *VideoHandler) ConfirmVideoUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmVideoUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrVideoRequestBodyInvalid.WithCause(err)
	}
	in, err := req.ToInput(accountID)
	if err != nil {
		return err
	}
	if _, err := h.app.ConfirmUpload(c.Context(), in); err != nil {
		return err
	}
	return fiberx.OK(c, "Video confirmed", httpx.SuccessOptions{})
}

func (h *VideoHandler) ConfirmVideosUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmVideoUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrVideoRequestBodyInvalid.WithCause(err)
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
	return fiberx.OK(c, "Videos confirmed", httpx.SuccessOptions{})
}

func (h *VideoHandler) DeleteVideo(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return asseterrs.ErrInvalidAccountID
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	videoID, err := uuid.Parse(c.Params("id"))
	if err != nil || videoID == uuid.Nil {
		return asseterrs.ErrInvalidVideoID.WithCause(err)
	}
	if err := h.app.Delete(c.Context(), videosApp.DeleteInput{
		AccountID: accountID,
		VideoID:   videoID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Video deleted", httpx.SuccessOptions{})
}

func (h *VideoHandler) ServeVideoFile(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return asseterrs.ErrInvalidVideoID.WithCause(err)
	}
	vo, err := h.app.Get(c.Context(), videosApp.GetInput{ID: id})
	if err != nil {
		return err
	}
	url, err := h.app.FileURL(c.Context(), videosApp.FileURLInput{ObjectKey: vo.FilePath})
	if err != nil {
		return err
	}
	return c.Redirect().To(url)
}
