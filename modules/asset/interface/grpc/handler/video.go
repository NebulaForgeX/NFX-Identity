package handler

import (
	"context"

	"nfxidentity/modules/asset/application/media"
	videopb "nfxidentity/protos/gen/asset/video"

	"github.com/google/uuid"
)

type VideoHandler struct {
	videopb.UnimplementedVideoServiceServer
	svc *media.Service
}

func NewVideoHandler(svc *media.Service) *VideoHandler {
	return &VideoHandler{svc: svc}
}

func toVideo(row *media.ListItem) *videopb.Video {
	if row == nil {
		return nil
	}
	return &videopb.Video{
		Id: row.ID, FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID,
	}
}

func (h *VideoHandler) GetVideoByID(ctx context.Context, req *videopb.GetVideoByIDRequest) (*videopb.GetVideoByIDResponse, error) {
	row, err := h.svc.Get(ctx, media.KindVideos, req.GetId())
	if err != nil {
		return nil, err
	}
	return &videopb.GetVideoByIDResponse{Video: toVideo(row)}, nil
}

func (h *VideoHandler) BatchGetVideos(ctx context.Context, req *videopb.BatchGetVideosRequest) (*videopb.BatchGetVideosResponse, error) {
	rows, err := h.svc.GetMany(ctx, media.KindVideos, req.GetIds())
	if err != nil {
		return nil, err
	}
	out := make([]*videopb.Video, 0, len(rows))
	for i := range rows {
		out = append(out, toVideo(&rows[i]))
	}
	return &videopb.BatchGetVideosResponse{Videos: out}, nil
}

func (h *VideoHandler) PrepareVideoUpload(ctx context.Context, req *videopb.PrepareVideoUploadRequest) (*videopb.PrepareVideoUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.Prepare(ctx, aid, media.KindVideos, req.GetFileName(), req.GetMimeType())
	if err != nil {
		return nil, err
	}
	return &videopb.PrepareVideoUploadResponse{VideoId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath}, nil
}

func (h *VideoHandler) PrepareVideosUpload(ctx context.Context, req *videopb.PrepareVideosUploadRequest) (*videopb.PrepareVideosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	results := make([]*videopb.PrepareVideoUploadResult, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		out, err := h.svc.Prepare(ctx, aid, media.KindVideos, item.GetFileName(), item.GetMimeType())
		if err != nil {
			return nil, err
		}
		results = append(results, &videopb.PrepareVideoUploadResult{VideoId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath})
	}
	return &videopb.PrepareVideosUploadResponse{Results: results}, nil
}

func (h *VideoHandler) ConfirmVideoUpload(ctx context.Context, req *videopb.ConfirmVideoUploadRequest) (*videopb.ConfirmVideoUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Confirm(ctx, aid, media.KindVideos, req.GetVideoId()); err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, media.KindVideos, req.GetVideoId())
	if err != nil {
		return nil, err
	}
	return &videopb.ConfirmVideoUploadResponse{Video: toVideo(row)}, nil
}

func (h *VideoHandler) ConfirmVideosUpload(ctx context.Context, req *videopb.ConfirmVideosUploadRequest) (*videopb.ConfirmVideosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out := make([]*videopb.Video, 0, len(req.GetVideoIds()))
	for _, id := range req.GetVideoIds() {
		if err := h.svc.Confirm(ctx, aid, media.KindVideos, id); err != nil {
			return nil, err
		}
		row, err := h.svc.Get(ctx, media.KindVideos, id)
		if err != nil {
			return nil, err
		}
		out = append(out, toVideo(row))
	}
	return &videopb.ConfirmVideosUploadResponse{Videos: out}, nil
}

func (h *VideoHandler) DeleteVideo(ctx context.Context, req *videopb.DeleteVideoRequest) (*videopb.DeleteVideoResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, aid, media.KindVideos, req.GetId()); err != nil {
		return nil, err
	}
	return &videopb.DeleteVideoResponse{}, nil
}
