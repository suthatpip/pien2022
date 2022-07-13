package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const defaultKeyStore string = "580f6eba5513d42eec8e88f09f176026efcb8ffa19c5618a94a72adc7c7a3c8f"

func EncryptMessage(stringToEncrypt string, key ...string) (encryptedString string) {
	keyStore := defaultKeyStore
	if len(key) == 0 {
		keyStore = defaultKeyStore
	}

	hkey, _ := hex.DecodeString(keyStore)
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(hkey)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func DecryptMessage(encryptedString string, key ...string) (decryptedString string) {

	keyStore := defaultKeyStore
	if len(key) == 0 {
		keyStore = defaultKeyStore
	}

	hkey, _ := hex.DecodeString(keyStore)
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(hkey)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
