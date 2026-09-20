package handler

import (
	authconn "nfxidentity/connections/auth"
	asseterrs "nfxidentity/errors/src/asset"
	filesApp "nfxidentity/modules/asset/application/files"
	"nfxidentity/modules/asset/interface/http/dto/reqdto"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type FileHandler struct {
	app  *filesApp.Service
	auth *authconn.Client
}

func NewFileHandler(app *filesApp.Service, auth *authconn.Client) *FileHandler {
	return &FileHandler{app: app, auth: auth}
}

func (h *FileHandler) ListFiles(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	rows, err := h.app.List(c.Context(), filesApp.ListInput{AccountID: accountID})
	if err != nil {
		return err
	}
	return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: rows})
}

func (h *FileHandler) PrepareFileUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareFileUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrFileRequestBodyInvalid.WithCause(err)
	}
	out, err := h.app.PrepareUpload(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Upload URL created", httpx.SuccessOptions{Data: reqdto.PrepareFileUploadDataFrom(out)})
}

func (h *FileHandler) PrepareFilesUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.PrepareFileUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrFileRequestBodyInvalid.WithCause(err)
	}
	if err := req.Validate(); err != nil {
		return err
	}
	out, err := h.app.PrepareUploads(c.Context(), req.ToInput(accountID))
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Upload URLs created", httpx.SuccessOptions{Data: map[string]any{"results": reqdto.PrepareFileUploadsDataFrom(out)}})
}

func (h *FileHandler) ConfirmFileUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmFileUpload
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrFileRequestBodyInvalid.WithCause(err)
	}
	in, err := req.ToInput(accountID)
	if err != nil {
		return err
	}
	if _, err := h.app.ConfirmUpload(c.Context(), in); err != nil {
		return err
	}
	return fiberx.OK(c, "File confirmed", httpx.SuccessOptions{})
}

func (h *FileHandler) ConfirmFilesUpload(c fiber.Ctx) error {
	accountID, _, _, err := uploadContext(c)
	if err != nil {
		return err
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	var req reqdto.ConfirmFileUploads
	if err := c.Bind().Body(&req); err != nil {
		return asseterrs.ErrFileRequestBodyInvalid.WithCause(err)
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
	return fiberx.OK(c, "Files confirmed", httpx.SuccessOptions{})
}

func (h *FileHandler) DeleteFile(c fiber.Ctx) error {
	accountID, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return asseterrs.ErrInvalidAccountID
	}
	if err := ensureOwnedProfile(c, h.auth, accountID); err != nil {
		return err
	}
	fileID, err := uuid.Parse(c.Params("id"))
	if err != nil || fileID == uuid.Nil {
		return asseterrs.ErrInvalidFileID.WithCause(err)
	}
	if err := h.app.Delete(c.Context(), filesApp.DeleteInput{
		AccountID: accountID,
		FileID:    fileID,
	}); err != nil {
		return err
	}
	return fiberx.OK(c, "File deleted", httpx.SuccessOptions{})
}

func (h *FileHandler) ServeFileFile(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return asseterrs.ErrInvalidFileID.WithCause(err)
	}
	vo, err := h.app.Get(c.Context(), filesApp.GetInput{ID: id})
	if err != nil {
		return err
	}
	url, err := h.app.FileURL(c.Context(), filesApp.FileURLInput{ObjectKey: vo.FilePath})
	if err != nil {
		return err
	}
	return c.Redirect().To(url)
}
