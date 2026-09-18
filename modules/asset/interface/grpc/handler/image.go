package handler

import (
	"context"

	"nfxidentity/modules/asset/application/media"
	imagepb "nfxidentity/protos/gen/asset/image"

	"github.com/google/uuid"
)

type ImageHandler struct {
	imagepb.UnimplementedImageServiceServer
	svc *media.Service
}

func NewImageHandler(svc *media.Service) *ImageHandler {
	return &ImageHandler{svc: svc}
}

func toImage(row *media.ListItem) *imagepb.Image {
	if row == nil {
		return nil
	}
	return &imagepb.Image{
		Id: row.ID, FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID,
	}
}

func (h *ImageHandler) GetImageByID(ctx context.Context, req *imagepb.GetImageByIDRequest) (*imagepb.GetImageByIDResponse, error) {
	row, err := h.svc.Get(ctx, media.KindImages, req.GetId())
	if err != nil {
		return nil, err
	}
	return &imagepb.GetImageByIDResponse{Image: toImage(row)}, nil
}

func (h *ImageHandler) BatchGetImages(ctx context.Context, req *imagepb.BatchGetImagesRequest) (*imagepb.BatchGetImagesResponse, error) {
	rows, err := h.svc.GetMany(ctx, media.KindImages, req.GetIds())
	if err != nil {
		return nil, err
	}
	out := make([]*imagepb.Image, 0, len(rows))
	for i := range rows {
		out = append(out, toImage(&rows[i]))
	}
	return &imagepb.BatchGetImagesResponse{Images: out}, nil
}

func (h *ImageHandler) PrepareImageUpload(ctx context.Context, req *imagepb.PrepareImageUploadRequest) (*imagepb.PrepareImageUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.Prepare(ctx, aid, media.KindImages, req.GetFileName(), req.GetMimeType())
	if err != nil {
		return nil, err
	}
	return &imagepb.PrepareImageUploadResponse{ImageId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath}, nil
}

func (h *ImageHandler) PrepareImagesUpload(ctx context.Context, req *imagepb.PrepareImagesUploadRequest) (*imagepb.PrepareImagesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	results := make([]*imagepb.PrepareImageUploadResult, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		out, err := h.svc.Prepare(ctx, aid, media.KindImages, item.GetFileName(), item.GetMimeType())
		if err != nil {
			return nil, err
		}
		results = append(results, &imagepb.PrepareImageUploadResult{ImageId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath})
	}
	return &imagepb.PrepareImagesUploadResponse{Results: results}, nil
}

func (h *ImageHandler) ConfirmImageUpload(ctx context.Context, req *imagepb.ConfirmImageUploadRequest) (*imagepb.ConfirmImageUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Confirm(ctx, aid, media.KindImages, req.GetImageId()); err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, media.KindImages, req.GetImageId())
	if err != nil {
		return nil, err
	}
	return &imagepb.ConfirmImageUploadResponse{Image: toImage(row)}, nil
}

func (h *ImageHandler) ConfirmImagesUpload(ctx context.Context, req *imagepb.ConfirmImagesUploadRequest) (*imagepb.ConfirmImagesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out := make([]*imagepb.Image, 0, len(req.GetImageIds()))
	for _, id := range req.GetImageIds() {
		if err := h.svc.Confirm(ctx, aid, media.KindImages, id); err != nil {
			return nil, err
		}
		row, err := h.svc.Get(ctx, media.KindImages, id)
		if err != nil {
			return nil, err
		}
		out = append(out, toImage(row))
	}
	return &imagepb.ConfirmImagesUploadResponse{Images: out}, nil
}

func (h *ImageHandler) DeleteImage(ctx context.Context, req *imagepb.DeleteImageRequest) (*imagepb.DeleteImageResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, aid, media.KindImages, req.GetId()); err != nil {
		return nil, err
	}
	return &imagepb.DeleteImageResponse{}, nil
}
