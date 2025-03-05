package storage

import (
	"crypto/sha256"
	"fmt"
)

// HashContent generates a SHA-256 hash of the given data.
func HashContent(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
