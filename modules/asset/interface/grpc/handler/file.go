package handler

import (
	"context"

	"nfxidentity/modules/asset/application/media"
	filepb "nfxidentity/protos/gen/asset/file"

	"github.com/google/uuid"
)

type FileHandler struct {
	filepb.UnimplementedFileServiceServer
	svc *media.Service
}

func NewFileHandler(svc *media.Service) *FileHandler {
	return &FileHandler{svc: svc}
}

func toFile(row *media.ListItem) *filepb.File {
	if row == nil {
		return nil
	}
	return &filepb.File{
		Id: row.ID, FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID,
	}
}

func (h *FileHandler) GetFileByID(ctx context.Context, req *filepb.GetFileByIDRequest) (*filepb.GetFileByIDResponse, error) {
	row, err := h.svc.Get(ctx, media.KindFiles, req.GetId())
	if err != nil {
		return nil, err
	}
	return &filepb.GetFileByIDResponse{File: toFile(row)}, nil
}

func (h *FileHandler) BatchGetFiles(ctx context.Context, req *filepb.BatchGetFilesRequest) (*filepb.BatchGetFilesResponse, error) {
	rows, err := h.svc.GetMany(ctx, media.KindFiles, req.GetIds())
	if err != nil {
		return nil, err
	}
	out := make([]*filepb.File, 0, len(rows))
	for i := range rows {
		out = append(out, toFile(&rows[i]))
	}
	return &filepb.BatchGetFilesResponse{Files: out}, nil
}

func (h *FileHandler) PrepareFileUpload(ctx context.Context, req *filepb.PrepareFileUploadRequest) (*filepb.PrepareFileUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.Prepare(ctx, aid, media.KindFiles, req.GetFileName(), req.GetMimeType())
	if err != nil {
		return nil, err
	}
	return &filepb.PrepareFileUploadResponse{FileId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath}, nil
}

func (h *FileHandler) PrepareFilesUpload(ctx context.Context, req *filepb.PrepareFilesUploadRequest) (*filepb.PrepareFilesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	results := make([]*filepb.PrepareFileUploadResult, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		out, err := h.svc.Prepare(ctx, aid, media.KindFiles, item.GetFileName(), item.GetMimeType())
		if err != nil {
			return nil, err
		}
		results = append(results, &filepb.PrepareFileUploadResult{FileId: out.ID, UploadUrl: out.UploadURL, FilePath: out.FilePath})
	}
	return &filepb.PrepareFilesUploadResponse{Results: results}, nil
}

func (h *FileHandler) ConfirmFileUpload(ctx context.Context, req *filepb.ConfirmFileUploadRequest) (*filepb.ConfirmFileUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Confirm(ctx, aid, media.KindFiles, req.GetFileId()); err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, media.KindFiles, req.GetFileId())
	if err != nil {
		return nil, err
	}
	return &filepb.ConfirmFileUploadResponse{File: toFile(row)}, nil
}

func (h *FileHandler) ConfirmFilesUpload(ctx context.Context, req *filepb.ConfirmFilesUploadRequest) (*filepb.ConfirmFilesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	out := make([]*filepb.File, 0, len(req.GetFileIds()))
	for _, id := range req.GetFileIds() {
		if err := h.svc.Confirm(ctx, aid, media.KindFiles, id); err != nil {
			return nil, err
		}
		row, err := h.svc.Get(ctx, media.KindFiles, id)
		if err != nil {
			return nil, err
		}
		out = append(out, toFile(row))
	}
	return &filepb.ConfirmFilesUploadResponse{Files: out}, nil
}

func (h *FileHandler) DeleteFile(ctx context.Context, req *filepb.DeleteFileRequest) (*filepb.DeleteFileResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, aid, media.KindFiles, req.GetId()); err != nil {
		return nil, err
	}
	return &filepb.DeleteFileResponse{}, nil
}
