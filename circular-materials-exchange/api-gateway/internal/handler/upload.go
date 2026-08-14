package handler

import (
	materialpb "api-gateway/internal/pb/material"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	material materialpb.MaterialServiceClient
}

func NewUploadHandler(material materialpb.MaterialServiceClient) *UploadHandler {
	return &UploadHandler{material: material}
}

func (h *UploadHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Không tìm thấy tệp"})
		return
	}
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Tệp quá lớn (tối đa 5MB)"})
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
	default:
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Định dạng không hỗ trợ (chỉ jpg, jpeg, png, gif, webp)"})
		return
	}
	source, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Lỗi đọc tệp"})
		return
	}
	defer source.Close()
	content, err := io.ReadAll(source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Lỗi đọc tệp"})
		return
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.material.UploadImage(ctx, &materialpb.UploadImageRequest{
		Filename: file.Filename, ContentType: file.Header.Get("Content-Type"), Content: content,
	})
	if err != nil {
		writeRPCError(c, err, "Lỗi tải ảnh lên MinIO")
		return
	}
	filename := strings.TrimPrefix(response.GetImageUrl(), "/images/")
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"url": response.GetImageUrl(), "filename": filename,
	}})
}
