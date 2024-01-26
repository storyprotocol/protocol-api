package encoder

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var key = validateKeySize("yourSecretKey1234567890123456")

func validateKeySize(key string) string {
	keyBytes := []byte(key)
	switch len(keyBytes) {
	case 16, 24, 32:
		// Valid key size
		return key
	case 0:
		// Empty key
		return ""
	default:
		// Adjust key size (truncate or pad)
		return string(keyBytes[:32])
	}
}

func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unpadPKCS7(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plaintext := []byte(plainText)
	plaintext = padPKCS7(plaintext, block.BlockSize())

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(encrypted string) (string, error) {
	ciphertext, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext = unpadPKCS7(ciphertext)

	return string(ciphertext), nil
}
