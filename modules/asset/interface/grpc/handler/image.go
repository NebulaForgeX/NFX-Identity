package handler

import (
	"context"

	imagesApp "nfxidentity/modules/asset/application/images"
	imgsQuery "nfxidentity/modules/asset/query/images"
	imagepb "nfxidentity/protos/gen/asset/image"

	"github.com/google/uuid"
)

type ImageHandler struct {
	imagepb.UnimplementedImageServiceServer
	svc *imagesApp.Service
}

func NewImageHandler(svc *imagesApp.Service) *ImageHandler {
	return &ImageHandler{svc: svc}
}

func toImage(row *imgsQuery.ImageVO) *imagepb.Image {
	if row == nil {
		return nil
	}
	return &imagepb.Image{
		Id: row.ID.String(), FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID.String(),
	}
}

func (h *ImageHandler) GetImageByID(ctx context.Context, req *imagepb.GetImageByIDRequest) (*imagepb.GetImageByIDResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, imagesApp.GetInput{ID: id})
	if err != nil {
		return nil, err
	}
	return &imagepb.GetImageByIDResponse{Image: toImage(row)}, nil
}

func (h *ImageHandler) BatchGetImages(ctx context.Context, req *imagepb.BatchGetImagesRequest) (*imagepb.BatchGetImagesResponse, error) {
	rows, err := h.svc.GetMany(ctx, req.GetIds())
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
	out, err := h.svc.PrepareUpload(ctx, imagesApp.PrepareUploadInput{
		AccountID: aid,
		FileName:  req.GetFileName(),
		MimeType:  req.GetMimeType(),
	})
	if err != nil {
		return nil, err
	}
	return &imagepb.PrepareImageUploadResponse{ImageId: out.ImageID.String(), UploadUrl: out.UploadURL, FilePath: out.ObjectKey}, nil
}

func (h *ImageHandler) PrepareImagesUpload(ctx context.Context, req *imagepb.PrepareImagesUploadRequest) (*imagepb.PrepareImagesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	items := make([]imagesApp.PrepareUploadItemInput, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		items = append(items, imagesApp.PrepareUploadItemInput{FileName: item.GetFileName(), MimeType: item.GetMimeType()})
	}
	out, err := h.svc.PrepareUploads(ctx, imagesApp.PrepareUploadsInput{AccountID: aid, Items: items})
	if err != nil {
		return nil, err
	}
	results := make([]*imagepb.PrepareImageUploadResult, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, &imagepb.PrepareImageUploadResult{ImageId: result.ImageID.String(), UploadUrl: result.UploadURL, FilePath: result.ObjectKey})
	}
	return &imagepb.PrepareImagesUploadResponse{Results: results}, nil
}

func (h *ImageHandler) ConfirmImageUpload(ctx context.Context, req *imagepb.ConfirmImageUploadRequest) (*imagepb.ConfirmImageUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	imageID, err := uuid.Parse(req.GetImageId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.ConfirmUpload(ctx, imagesApp.ConfirmUploadInput{AccountID: aid, ImageID: imageID})
	if err != nil {
		return nil, err
	}
	return &imagepb.ConfirmImageUploadResponse{Image: toImage(out.Image)}, nil
}

func (h *ImageHandler) ConfirmImagesUpload(ctx context.Context, req *imagepb.ConfirmImagesUploadRequest) (*imagepb.ConfirmImagesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	ids := make([]uuid.UUID, 0, len(req.GetImageIds()))
	for _, raw := range req.GetImageIds() {
		id, err := uuid.Parse(raw)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	out, err := h.svc.ConfirmUploads(ctx, imagesApp.ConfirmUploadsInput{AccountID: aid, ImageIDs: ids})
	if err != nil {
		return nil, err
	}
	rows := make([]*imagepb.Image, 0, len(out.Images))
	for _, vo := range out.Images {
		rows = append(rows, toImage(vo))
	}
	return &imagepb.ConfirmImagesUploadResponse{Images: rows}, nil
}

func (h *ImageHandler) DeleteImage(ctx context.Context, req *imagepb.DeleteImageRequest) (*imagepb.DeleteImageResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, imagesApp.DeleteInput{AccountID: aid, ImageID: id}); err != nil {
		return nil, err
	}
	return &imagepb.DeleteImageResponse{}, nil
}
