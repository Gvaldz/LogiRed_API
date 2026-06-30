package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

var masterKey []byte

func init() {
	key := os.Getenv("APP_MASTER_KEY")
	if key == "" {
		// Default key for local dev if not set
		key = "local-dev-master-key-32bytes!!!"
	}
	// AES-256 requires a 32-byte key. Pad or truncate to ensure it's exactly 32 bytes.
	if len(key) < 32 {
		key = fmt.Sprintf("%-32s", key)
	} else if len(key) > 32 {
		key = key[:32]
	}
	masterKey = []byte(key)
}

// EncryptString encrypts a plaintext string using AES-GCM and returns a base64 encoded string.
func EncryptString(plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return "enc::v1::" + base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptString decrypts a base64 encoded string (prefixed with enc::v1::) back to plaintext.
func DecryptString(ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", nil
	}

	// Check if it's actually encrypted
	prefix := "enc::v1::"
	if len(ciphertext) < len(prefix) || ciphertext[:len(prefix)] != prefix {
		return ciphertext, nil // Return as is, it's not encrypted or it is legacy data
	}

	encodedData := ciphertext[len(prefix):]

	data, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]
	plaintextBytes, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}
