package handler

import (
	"context"

	"nfxidentity/modules/asset/application/media"
	audiopb "nfxidentity/protos/gen/asset/audio"

	"github.com/google/uuid"
)

type AudioHandler struct {
	audiopb.UnimplementedAudioServiceServer
	svc *media.Service
}

func NewAudioHandler(svc *media.Service) *AudioHandler {
	return &AudioHandler{svc: svc}
}

func toAudio(row *media.ListItem) *audiopb.Audio {
	if row == nil {
		return nil
	}
	return &audiopb.Audio{
		Id: row.ID, FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID,
	}
}

func (h *AudioHandler) GetAudioByID(ctx context.Context, req *audiopb.GetAudioByIDRequest) (*audiopb.GetAudioByIDResponse, error) {
	row, err := h.svc.Get(ctx, media.KindAudios, req.GetId())
	if err != nil {
		return nil, err
	}
	return &audiopb.GetAudioByIDResponse{Audio: toAudio(row)}, nil
}

func (h *AudioHandler) BatchGetAudios(ctx context.Context, req *audiopb.BatchGetAudiosRequest) (*audiopb.BatchGetAudiosResponse, error) {
	rows, err := h.svc.GetMany(ctx, media.KindAudios, req.GetIds())
	if err != nil {
		return nil, err
	}
	out := make([]*audiopb.Audio, 0, len(rows))
	for i := range rows {
		out = append(out, toAudio(&rows[i]))
	}
	return &audiopb.BatchGetAudiosResponse{Audios: out}, nil
}

func (h *AudioHandler) PrepareAudioUpload(ctx context.Context, req *audiopb.PrepareAudioUploadRequest) (*audiopb.PrepareAudioUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.Prepare(ctx, aid, media.KindAudios, req.GetFileName(), req.GetMimeType())
	if err != nil {
		return nil, err
	}
	return &audiopb.PrepareAudioUploadResponse{AudioId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath}, nil
}

func (h *AudioHandler) PrepareAudiosUpload(ctx context.Context, req *audiopb.PrepareAudiosUploadRequest) (*audiopb.PrepareAudiosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	results := make([]*audiopb.PrepareAudioUploadResult, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		out, err := h.svc.Prepare(ctx, aid, media.KindAudios, item.GetFileName(), item.GetMimeType())
		if err != nil {
			return nil, err
		}
		results = append(results, &audiopb.PrepareAudioUploadResult{AudioId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath})
	}
	return &audiopb.PrepareAudiosUploadResponse{Results: results}, nil
}

func (h *AudioHandler) ConfirmAudioUpload(ctx context.Context, req *audiopb.ConfirmAudioUploadRequest) (*audiopb.ConfirmAudioUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Confirm(ctx, aid, media.KindAudios, req.GetAudioId()); err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, media.KindAudios, req.GetAudioId())
	if err != nil {
		return nil, err
	}
	return &audiopb.ConfirmAudioUploadResponse{Audio: toAudio(row)}, nil
}

func (h *AudioHandler) ConfirmAudiosUpload(ctx context.Context, req *audiopb.ConfirmAudiosUploadRequest) (*audiopb.ConfirmAudiosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out := make([]*audiopb.Audio, 0, len(req.GetAudioIds()))
	for _, id := range req.GetAudioIds() {
		if err := h.svc.Confirm(ctx, aid, media.KindAudios, id); err != nil {
			return nil, err
		}
		row, err := h.svc.Get(ctx, media.KindAudios, id)
		if err != nil {
			return nil, err
		}
		out = append(out, toAudio(row))
	}
	return &audiopb.ConfirmAudiosUploadResponse{Audios: out}, nil
}

func (h *AudioHandler) DeleteAudio(ctx context.Context, req *audiopb.DeleteAudioRequest) (*audiopb.DeleteAudioResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, aid, media.KindAudios, req.GetId()); err != nil {
		return nil, err
	}
	return &audiopb.DeleteAudioResponse{}, nil
}
