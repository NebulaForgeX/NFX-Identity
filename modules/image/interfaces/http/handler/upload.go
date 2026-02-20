package handler

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	imageApp "nfxid/modules/image/application/images"
	imageAppCommands "nfxid/modules/image/application/images/commands"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// 允许的图片 MIME 类型
var allowedMimeTypes = map[string]string{
	"image/jpeg":    ".jpg",
	"image/png":     ".png",
	"image/gif":     ".gif",
	"image/webp":    ".webp",
	"image/svg+xml": ".svg",
}

// 最大文件大小: 10MB
const maxFileSize = 10 * 1024 * 1024

type UploadHandler struct {
	appSvc      *imageApp.Service
	storagePath string
}

func NewUploadHandler(appSvc *imageApp.Service, storagePath string) *UploadHandler {
	return &UploadHandler{
		appSvc:      appSvc,
		storagePath: storagePath,
	}
}

// UploadImage 上传图片
// @Summary Upload an image
// @Description Upload an image file to the server
// @Tags Images
// @Accept multipart/form-data
// @Produce json
// @Param file formance file true "Image file"
// @Param type query string false "Image type: avatar, background, tmp" default(tmp)
// @Param is_public query bool false "Whether the image is public" default(true)
// @Success 201 {object} respdto.ImageDTO
// @Failure 400 {object} httpx.HTTPResp
// @Failure 500 {object} httpx.HTTPResp
// @Router /image/auth/upload [post]
func (h *UploadHandler) UploadImage(c fiber.Ctx) error {
	ctx := c.Context()
	var userID *uuid.UUID
	if uid, ok := fiberx.UserIDFromContext(ctx); ok {
		userID = &uid
	}

	// 图片类型固定为 tmp（所有上传都先到 tmp 目录，后续通过 Directory 服务创建时移动）
	imageType := "tmp"

	// 获取是否公开
	isPublic := fiber.Query[bool](c, "is_public", true)

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "No file uploaded").WithCause(err))
	}

	// 检查文件大小
	if file.Size > maxFileSize {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", fmt.Sprintf("File size exceeds maximum allowed size of %d bytes", maxFileSize)))
	}

	// 检查 MIME 类型
	contentType := file.Header.Get("Content-Type")
	ext, ok := allowedMimeTypes[contentType]
	if !ok {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "Invalid file type. Only JPEG, PNG, GIF, WebP, and SVG images are allowed"))
	}

	// 生成唯一的文件名
	imageID := uuid.New()
	filename := imageID.String() + ext

	// 构建存储路径: data/images/{type}/{filename}
	relativePath := filepath.Join("images", imageType, filename)
	fullPath := filepath.Join(h.storagePath, relativePath)

	// 确保目录存在
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标文件
	dst, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, src); err != nil {
		os.Remove(fullPath) // 清理失败的文件
		return err
	}

	// 创建图片记录
	cmd := imageAppCommands.CreateImageCmd{
		UserID:           userID,
		Filename:         filename,
		OriginalFilename: file.Filename,
		MimeType:         contentType,
		Size:             file.Size,
		StoragePath:      relativePath,
		IsPublic:         isPublic,
		Metadata: map[string]interface{}{
			"type":        imageType,
			"uploaded_at": time.Now().UTC().Format(time.RFC3339),
		},
	}

	createdID, err := h.appSvc.CreateImage(c.Context(), cmd)
	if err != nil {
		os.Remove(fullPath) // 清理文件
		return err
	}

	// 获取创建的图片
	imageView, err := h.appSvc.GetImage(c.Context(), createdID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Image uploaded successfully", httpx.SuccessOptions{Data: respdto.ImageROToDTO(&imageView)})
}

// MoveImage 移动图片（从 tmp 移动到正式目录）
// @Summary Move an image
// @Description Move an image from tmp to a permanent location
// @Tags Images
// @Accept json
// @Produce json
// @Param id path string true "Image ID"
// @Param type query string true "Target type: avatar, background"
// @Success 200 {object} respdto.ImageDTO
// @Failure 400 {object} httpx.HTTPResp
// @Failure 404 {object} httpx.HTTPResp
// @Failure 500 {object} httpx.HTTPResp
// @Router /image/auth/images/{id}/move [post]
func (h *UploadHandler) MoveImage(c fiber.Ctx) error {
	// 获取图片 ID
	imageIDStr := c.Params("id")
	imageID, err := uuid.Parse(imageIDStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// 获取目标类型
	targetType := c.Query("type", "")
	if targetType != "avatar" && targetType != "background" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "Invalid target type. Must be 'avatar' or 'background'"))
	}

	// 获取图片信息
	imageView, err := h.appSvc.GetImage(c.Context(), imageID)
	if err != nil {
		return err
	}

	// 检查是否在 tmp 目录
	if !strings.Contains(imageView.StoragePath, "/tmp/") && !strings.HasPrefix(imageView.StoragePath, "images/tmp/") {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "Image is not in tmp directory"))
	}

	// 构建新路径
	filename := filepath.Base(imageView.StoragePath)
	newRelativePath := filepath.Join("images", targetType, filename)
	oldFullPath := filepath.Join(h.storagePath, imageView.StoragePath)
	newFullPath := filepath.Join(h.storagePath, newRelativePath)

	// 确保目标目录存在
	dir := filepath.Dir(newFullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 移动文件
	if err := os.Rename(oldFullPath, newFullPath); err != nil {
		// 如果 rename 失败（可能跨文件系统），则使用复制+删除
		if err := copyFile(oldFullPath, newFullPath); err != nil {
			return err
		}
		os.Remove(oldFullPath)
	}

	// 更新数据库记录
	cmd := imageAppCommands.UpdateImageCmd{
		ImageID:     imageID,
		StoragePath: &newRelativePath,
		Metadata: map[string]interface{}{
			"type":     targetType,
			"moved_at": time.Now().UTC().Format(time.RFC3339),
		},
	}
	if err := h.appSvc.UpdateImage(c.Context(), cmd); err != nil {
		// 回滚文件移动
		os.Rename(newFullPath, oldFullPath)
		return err
	}

	// 获取更新后的图片
	updatedImage, err := h.appSvc.GetImage(c.Context(), imageID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image moved successfully", httpx.SuccessOptions{Data: respdto.ImageROToDTO(&updatedImage)})
}

// copyFile 复制文件
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// ServeImage 根据 image ID 返回图片文件（公开，供 <img src> 使用）
func (h *UploadHandler) ServeImage(c fiber.Ctx) error {
	idStr := c.Params("image_id")
	if idStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "image id required"))
	}
	imageID, err := uuid.Parse(idStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	ro, err := h.appSvc.GetImage(c.Context(), imageID)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(h.storagePath, ro.StoragePath)
	if _, err := os.Stat(fullPath); err != nil {
		return err
	}

	c.Set("Content-Type", ro.MimeType)
	return c.SendFile(fullPath)
}

// fiber:context-methods migrated
