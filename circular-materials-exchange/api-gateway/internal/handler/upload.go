package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadHandler struct {
	minioURL string
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{
		minioURL: "http://localhost:9000",
	}
}

func (h *UploadHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Khong tim thay file"})
		return
	}

	// Kiem tra kich thuoc (toi da 5MB)
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "File qua lon (toi da 5MB)"})
		return
	}

	// Kiem tra dinh dang
	ext := strings.ToLower(getExt(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".webp" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Dinh dang khong ho tro (chi jpg, png, gif, webp)"})
		return
	}

	// Doc file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Loi doc file"})
		return
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Loi doc file"})
		return
	}

	// Tao ten file duy nhat
	filename := fmt.Sprintf("listings/%s_%d%s", uuid.New().String()[:8], time.Now().Unix(), ext)

	// Upload len MinIO via S3 API
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "image/jpeg"
	}

	url := fmt.Sprintf("%s/cme-images/%s", h.minioURL, filename)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Loi tao request"})
		return
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Loi upload len MinIO: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": fmt.Sprintf("MinIO tra ve loi: %d", resp.StatusCode)})
		return
	}

	// Tra ve URL
	imageURL := fmt.Sprintf("/images/%s", filename)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"url":      imageURL,
			"filename": filename,
		},
	})
}

func getExt(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return "." + parts[len(parts)-1]
	}
	return ""
}
