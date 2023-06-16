package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"selfletter-backend/config"
)

// GenerateSecureToken returns a string token which consists of 128 random alphanumeric characters
//
// In case of error it returns an empty string
func GenerateSecureToken() string {
	cfg := config.GetConfig()
	byteSlice := make([]byte, cfg.TokenAndKeyLength/2)

	if _, err := rand.Read(byteSlice); err != nil {
		return ""
	}

	return hex.EncodeToString(byteSlice)
}
