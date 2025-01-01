package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"

	"company-name/pkg/file"
)

type FileHandler struct {
	FileService *file.FileService
}

func NewFileHandler(fileService *file.FileService) *FileHandler {
	return &FileHandler{
		FileService: fileService,
	}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form: " + err.Error()})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file: " + err.Error()})
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file: " + err.Error()})
		return
	}

	filename := filepath.Base(header.Filename) // Ensure a safe filename
	filePath, err := h.FileService.SaveFile(filename, fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"filePath": filePath})
}

func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("filePath")
	if filePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	if err := h.FileService.DeleteFile(filePath); err != nil {
		http.Error(w, "Failed to delete file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
