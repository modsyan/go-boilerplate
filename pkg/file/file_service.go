package file

import (
	"errors"
	"os"
	"path/filepath"
)

type FileService struct {
	BasePath string // Base directory for storing files
}

func NewFileService(basePath string) *FileService {
	return &FileService{BasePath: basePath}
}

func (s *FileService) SaveFile(filename string, data []byte) (string, error) {
	filePath := filepath.Join(s.BasePath, filename)

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return "", err
	}

	if err := os.WriteFile(filePath, data, os.ModePerm); err != nil {
		return "", err
	}

	return filePath, nil
}

func (s *FileService) DeleteFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil // Ignore error if the file doesn't exist
		}
		return err
	}
	return nil
}

func (s *FileService) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
