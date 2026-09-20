package handler

import (
	authconn "nfxidentity/connections/auth"
	asseterrs "nfxidentity/errors/src/asset"
	imagesApp "nfxidentity/modules/asset/application/images"
	"nfxidentity/modules/asset/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ImageHandler struct {
	app  *imagesApp.Service
	auth *authconn.Client
}

func NewImageHandler(app *imagesApp.Service, auth *authconn.Client) *ImageHandler {
	return &ImageHandler{app: app, auth: auth}
}

func (h *ImageHandler) ListImages(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	rows, err := h.app.List(c.Context(), imagesApp.ListInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: rows})
}

func (h *ImageHandler) PrepareImageUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareImageUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrImageRequestBodyInvalid.WithCause(err)
	}
	out, err := h.app.PrepareUpload(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Upload URL created", httpx.SuccessOptions{Data: reqdto.PrepareImageUploadDataFrom(out)})
}

func (h *ImageHandler) PrepareImagesUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareImageUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrImageRequestBodyInvalid.WithCause(err)
	}
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.app.PrepareUploads(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Upload URLs created", httpx.SuccessOptions{Data: map[string]any{"results": reqdto.PrepareImageUploadsDataFrom(out)}})
}

func (h *ImageHandler) ConfirmImageUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmImageUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrImageRequestBodyInvalid.WithCause(err)
	}
	in, err := req.ToInput(accountID)
	if err != nil {
		return err
	}
	if _, err := h.app.ConfirmUpload(c.Context(), in); err != nil {
		return err
	}
	return fiberx.OK(c, "Image confirmed", httpx.SuccessOptions{})
}

func (h *ImageHandler) ConfirmImagesUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmImageUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrImageRequestBodyInvalid.WithCause(err)
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
	return fiberx.OK(c, "Images confirmed", httpx.SuccessOptions{})
}

func (h *ImageHandler) DeleteImage(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return asseterrs.ErrInvalidAccountID
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	imageID, err := uuid.Parse(c.Params("id"))
	if err != nil || imageID == uuid.Nil {
		return asseterrs.ErrInvalidImageID.WithCause(err)
	}
	if err := h.app.Delete(c.Context(), imagesApp.DeleteInput{
		AccountID: accountID,
		ImageID:   imageID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "Image deleted", httpx.SuccessOptions{})
}

func (h *ImageHandler) ServeImageFile(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return asseterrs.ErrInvalidImageID.WithCause(err)
	}
	vo, err := h.app.Get(c.Context(), imagesApp.GetInput{ID: id})
	if err != nil {
		return err
	}
	url, err := h.app.FileURL(c.Context(), imagesApp.FileURLInput{ObjectKey: vo.FilePath})
	if err != nil {
		return err
	}
	return c.Redirect().To(url)
}
