package handler

import (
	"context"

	filesApp "nfxidentity/modules/asset/application/files"
	flsQuery "nfxidentity/modules/asset/query/files"
	filepb "nfxidentity/protos/gen/asset/file"

	"github.com/google/uuid"
)

type FileHandler struct {
	filepb.UnimplementedFileServiceServer
	svc *filesApp.Service
}

func NewFileHandler(svc *filesApp.Service) *FileHandler {
	return &FileHandler{svc: svc}
}

func toFile(row *flsQuery.FileVO) *filepb.File {
	if row == nil {
		return nil
	}
	return &filepb.File{
		Id: row.ID.String(), FilePath: row.FilePath, FileName: row.FileName,
		FileSize: row.FileSize, MimeType: row.MimeType, UploaderId: row.UploaderID.String(),
	}
}

func (h *FileHandler) GetFileByID(ctx context.Context, req *filepb.GetFileByIDRequest) (*filepb.GetFileByIDResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	row, err := h.svc.Get(ctx, filesApp.GetInput{ID: id})
	if err != nil {
		return nil, err
	}
	return &filepb.GetFileByIDResponse{File: toFile(row)}, nil
}

func (h *FileHandler) BatchGetFiles(ctx context.Context, req *filepb.BatchGetFilesRequest) (*filepb.BatchGetFilesResponse, error) {
	rows, err := h.svc.GetMany(ctx, req.GetIds())
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
	out, err := h.svc.PrepareUpload(ctx, filesApp.PrepareUploadInput{
		AccountID: aid,
		FileName:  req.GetFileName(),
		MimeType:  req.GetMimeType(),
	})
	if err != nil {
		return nil, err
	}
	return &filepb.PrepareFileUploadResponse{FileId: out.FileID.String(), UploadUrl: out.UploadURL, FilePath: out.ObjectKey}, nil
}

func (h *FileHandler) PrepareFilesUpload(ctx context.Context, req *filepb.PrepareFilesUploadRequest) (*filepb.PrepareFilesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	items := make([]filesApp.PrepareUploadItemInput, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		items = append(items, filesApp.PrepareUploadItemInput{FileName: item.GetFileName(), MimeType: item.GetMimeType()})
	}
	out, err := h.svc.PrepareUploads(ctx, filesApp.PrepareUploadsInput{AccountID: aid, Items: items})
	if err != nil {
		return nil, err
	}
	results := make([]*filepb.PrepareFileUploadResult, 0, len(out.Results))
	for _, result := range out.Results {
		results = append(results, &filepb.PrepareFileUploadResult{FileId: result.FileID.String(), UploadUrl: result.UploadURL, FilePath: result.ObjectKey})
	}
	return &filepb.PrepareFilesUploadResponse{Results: results}, nil
}

func (h *FileHandler) ConfirmFileUpload(ctx context.Context, req *filepb.ConfirmFileUploadRequest) (*filepb.ConfirmFileUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	fileID, err := uuid.Parse(req.GetFileId())
	if err != nil {
		return nil, err
	}
	out, err := h.svc.ConfirmUpload(ctx, filesApp.ConfirmUploadInput{AccountID: aid, FileID: fileID})
	if err != nil {
		return nil, err
	}
	return &filepb.ConfirmFileUploadResponse{File: toFile(out.File)}, nil
}

func (h *FileHandler) ConfirmFilesUpload(ctx context.Context, req *filepb.ConfirmFilesUploadRequest) (*filepb.ConfirmFilesUploadResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	ids := make([]uuid.UUID, 0, len(req.GetFileIds()))
	for _, raw := range req.GetFileIds() {
		id, err := uuid.Parse(raw)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	out, err := h.svc.ConfirmUploads(ctx, filesApp.ConfirmUploadsInput{AccountID: aid, FileIDs: ids})
	if err != nil {
		return nil, err
	}
	rows := make([]*filepb.File, 0, len(out.Files))
	for _, vo := range out.Files {
		rows = append(rows, toFile(vo))
	}
	return &filepb.ConfirmFilesUploadResponse{Files: rows}, nil
}

func (h *FileHandler) DeleteFile(ctx context.Context, req *filepb.DeleteFileRequest) (*filepb.DeleteFileResponse, error) {
	aid, err := uuid.Parse(req.GetAccountId())
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	if err := h.svc.Delete(ctx, filesApp.DeleteInput{AccountID: aid, FileID: id}); err != nil {
		return nil, err
	}
	return &filepb.DeleteFileResponse{}, nil
}
