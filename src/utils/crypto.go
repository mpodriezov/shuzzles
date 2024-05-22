package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash creates a hash from the input string using SHA256 algorithm.
func Hash(input string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash, nil
}

// CheckHash checks if the input string generates the given hash using SHA256 algorithm.
func CheckHash(input, hash string) bool {
	computedHash, err := Hash(input)
	if err != nil {
		return false
	}
	return computedHash == hash
}
