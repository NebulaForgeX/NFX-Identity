package handler

import (
	"context"

	audiosApp "nfxidentity/modules/asset/application/audios"
	audsQuery "nfxidentity/modules/asset/query/audios"
	audiopb "nfxidentity/protos/gen/asset/audio"

	"github.com/google/uuid"
)

type AudioHandler struct {
	audiopb.UnimplementedAudioServiceServer
	svc *audiosApp.Service
}

func NewAudioHandler(svc *audiosApp.Service) *AudioHandler {
	return &AudioHandler{svc: svc}
}

func toAudio(row *audsQuery.AudioVO) *audiopb.Audio {
	if row == nil {
		return nil
	}
	return &audiopb.Audio{
		Id: row.ID.String(), FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID.String(),
	}
}

func (h *AudioHandler) GetAudioByID(ctx context.Context, req *audiopb.GetAudioByIDRequest) (*audiopb.GetAudioByIDResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, audiosApp.GetInput{ID: id})
	if err != nil {
		return nil, err
	}
	return &audiopb.GetAudioByIDResponse{Audio: toAudio(row)}, nil
}

func (h *AudioHandler) BatchGetAudios(ctx context.Context, req *audiopb.BatchGetAudiosRequest) (*audiopb.BatchGetAudiosResponse, error) {
	rows, err := h.svc.GetMany(ctx, req.GetIds())
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
	out, err := h.svc.PrepareUpload(ctx, audiosApp.PrepareUploadInput{
		AccountID: aid,
		FileName:  req.GetFileName(),
		MimeType:  req.GetMimeType(),
	})
	if err != nil {
		return nil, err
	}
	return &audiopb.PrepareAudioUploadResponse{AudioId: out.AudioID.String(), UploadUrl: out.UploadURL, FilePath: out.ObjectKey}, nil
}

func (h *AudioHandler) PrepareAudiosUpload(ctx context.Context, req *audiopb.PrepareAudiosUploadRequest) (*audiopb.PrepareAudiosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	items := make([]audiosApp.PrepareUploadItemInput, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		items = append(items, audiosApp.PrepareUploadItemInput{FileName: item.GetFileName(), MimeType: item.GetMimeType()})
	}
	out, err := h.svc.PrepareUploads(ctx, audiosApp.PrepareUploadsInput{AccountID: aid, Items: items})
	if err != nil {
		return nil, err
	}
	results := make([]*audiopb.PrepareAudioUploadResult, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, &audiopb.PrepareAudioUploadResult{AudioId: result.AudioID.String(), UploadUrl: result.UploadURL, FilePath: result.ObjectKey})
	}
	return &audiopb.PrepareAudiosUploadResponse{Results: results}, nil
}

func (h *AudioHandler) ConfirmAudioUpload(ctx context.Context, req *audiopb.ConfirmAudioUploadRequest) (*audiopb.ConfirmAudioUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	audioID, err := uuid.Parse(req.GetAudioId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.ConfirmUpload(ctx, audiosApp.ConfirmUploadInput{AccountID: aid, AudioID: audioID})
	if err != nil {
		return nil, err
	}
	return &audiopb.ConfirmAudioUploadResponse{Audio: toAudio(out.Audio)}, nil
}

func (h *AudioHandler) ConfirmAudiosUpload(ctx context.Context, req *audiopb.ConfirmAudiosUploadRequest) (*audiopb.ConfirmAudiosUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	ids := make([]uuid.UUID, 0, len(req.GetAudioIds()))
	for _, raw := range req.GetAudioIds() {
		id, err := uuid.Parse(raw)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	out, err := h.svc.ConfirmUploads(ctx, audiosApp.ConfirmUploadsInput{AccountID: aid, AudioIDs: ids})
	if err != nil {
		return nil, err
	}
	rows := make([]*audiopb.Audio, 0, len(out.Audios))
	for _, vo := range out.Audios {
		rows = append(rows, toAudio(vo))
	}
	return &audiopb.ConfirmAudiosUploadResponse{Audios: rows}, nil
}

func (h *AudioHandler) DeleteAudio(ctx context.Context, req *audiopb.DeleteAudioRequest) (*audiopb.DeleteAudioResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, audiosApp.DeleteInput{AccountID: aid, AudioID: id}); err != nil {
		return nil, err
	}
	return &audiopb.DeleteAudioResponse{}, nil
}
