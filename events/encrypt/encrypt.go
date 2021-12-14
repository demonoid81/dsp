package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(text string, keyEncrypt string) string {
	key, _ := hex.DecodeString(keyEncrypt)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "error"
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "error"
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return fmt.Sprintf("%x", ciphertext)
}

func Decrypt(text string, keyDecrypt string) string {
	key, _ := hex.DecodeString(keyDecrypt)
	ciphertext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "error"
	}

	if len(ciphertext) < aes.BlockSize {
		return "error"
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
