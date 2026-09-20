package handler

import (
	"context"

	videosApp "nfxidentity/modules/asset/application/videos"
	vidsQuery "nfxidentity/modules/asset/query/videos"
	videopb "nfxidentity/protos/gen/asset/video"

	"github.com/google/uuid"
)

type VideoHandler struct {
	videopb.UnimplementedVideoServiceServer
	svc *videosApp.Service
}

func NewVideoHandler(svc *videosApp.Service) *VideoHandler {
	return &VideoHandler{svc: svc}
}

func toVideo(row *vidsQuery.VideoVO) *videopb.Video {
	if row == nil {
		return nil
	}
	return &videopb.Video{
		Id: row.ID.String(), FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID.String(),
	}
}

func (h *VideoHandler) GetVideoByID(ctx context.Context, req *videopb.GetVideoByIDRequest) (*videopb.GetVideoByIDResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, videosApp.GetInput{ID: id})
	if err != nil {
		return nil, err
	}
	return &videopb.GetVideoByIDResponse{Video: toVideo(row)}, nil
}

func (h *VideoHandler) BatchGetVideos(ctx context.Context, req *videopb.BatchGetVideosRequest) (*videopb.BatchGetVideosResponse, error) {
	rows, err := h.svc.GetMany(ctx, req.GetIds())
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
	out, err := h.svc.PrepareUpload(ctx, videosApp.PrepareUploadInput{
		AccountID: aid,
		FileName:  req.GetFileName(),
		MimeType:  req.GetMimeType(),
	})
	if err != nil {
		return nil, err
	}
	return &videopb.PrepareVideoUploadResponse{VideoId: out.VideoID.String(), UploadUrl: out.UploadURL, FilePath: out.ObjectKey}, nil
}

func (h *VideoHandler) PrepareVideosUpload(ctx context.Context, req *videopb.PrepareVideosUploadRequest) (*videopb.PrepareVideosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	items := make([]videosApp.PrepareUploadItemInput, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		items = append(items, videosApp.PrepareUploadItemInput{FileName: item.GetFileName(), MimeType: item.GetMimeType()})
	}
	out, err := h.svc.PrepareUploads(ctx, videosApp.PrepareUploadsInput{AccountID: aid, Items: items})
	if err != nil {
		return nil, err
	}
	results := make([]*videopb.PrepareVideoUploadResult, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, &videopb.PrepareVideoUploadResult{VideoId: result.VideoID.String(), UploadUrl: result.UploadURL, FilePath: result.ObjectKey})
	}
	return &videopb.PrepareVideosUploadResponse{Results: results}, nil
}

func (h *VideoHandler) ConfirmVideoUpload(ctx context.Context, req *videopb.ConfirmVideoUploadRequest) (*videopb.ConfirmVideoUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	videoID, err := uuid.Parse(req.GetVideoId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.ConfirmUpload(ctx, videosApp.ConfirmUploadInput{AccountID: aid, VideoID: videoID})
	if err != nil {
		return nil, err
	}
	return &videopb.ConfirmVideoUploadResponse{Video: toVideo(out.Video)}, nil
}

func (h *VideoHandler) ConfirmVideosUpload(ctx context.Context, req *videopb.ConfirmVideosUploadRequest) (*videopb.ConfirmVideosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	ids := make([]uuid.UUID, 0, len(req.GetVideoIds()))
	for _, raw := range req.GetVideoIds() {
		id, err := uuid.Parse(raw)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	out, err := h.svc.ConfirmUploads(ctx, videosApp.ConfirmUploadsInput{AccountID: aid, VideoIDs: ids})
	if err != nil {
		return nil, err
	}
	rows := make([]*videopb.Video, 0, len(out.Videos))
	for _, vo := range out.Videos {
		rows = append(rows, toVideo(vo))
	}
	return &videopb.ConfirmVideosUploadResponse{Videos: rows}, nil
}

func (h *VideoHandler) DeleteVideo(ctx context.Context, req *videopb.DeleteVideoRequest) (*videopb.DeleteVideoResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, videosApp.DeleteInput{AccountID: aid, VideoID: id}); err != nil {
		return nil, err
	}
	return &videopb.DeleteVideoResponse{}, nil
}
