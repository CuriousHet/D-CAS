package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// StoreFile stores the given file and returns its hash
func StoreFile(filePath string) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Generate SHA-256 hash
	hash := sha256.Sum256(content)
	hashStr := hex.EncodeToString(hash[:])

	// Save the file to the D-CAS storage (a simple local implementation for now)
	err = os.WriteFile("storage/"+hashStr, content, 0644)
	if err != nil {
		fmt.Println("Error storing file:", err)
		return
	}

	fmt.Printf("Stored %s with hash: %s\n", filePath, hashStr)
}

// RetrieveFile fetches a file from storage using its hash
func RetrieveFile(hash string) {
	// Read the file from the storage
	content, err := os.ReadFile("storage/" + hash)
	if err != nil {
		fmt.Println("File not found:", err)
		return
	}

	// Save it with a new name
	fileName := "retrieved_" + hash[:8] + ".txt"
	err = os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println("Error saving retrieved file:", err)
		return
	}

	fmt.Printf("Retrieved file stored as %s\n", fileName)
}
