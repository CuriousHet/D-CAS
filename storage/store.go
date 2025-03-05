package storage

import (
	"os"
)

// StoreFile saves the file content and returns its content-addressable hash.
func StoreFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := HashContent(data)
	storagePath := "storage/" + hash

	err = os.WriteFile(storagePath, data, 0644)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// RetrieveFile fetches file content using its hash.
func RetrieveFile(hash string) ([]byte, error) {
	storagePath := "storage/" + hash
	return os.ReadFile(storagePath)
}

// InitStorage ensures the storage directory exists.
func InitStorage() error {
	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		return os.Mkdir("storage", 0755)
	}
	return nil
}
