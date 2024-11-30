package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type Cryptography interface {
	Encrypt() (string, error)
	Decrypt() (string, error)
	GenerateKey() ([]byte, error)
}

type AES struct {
	Key       string
	PlainText string
}

type DES struct {
	Key       string
	PlainText string
}

func (m *AES) Encrypt() (cipherText string, err error) {
	keyBytes, err := base64.StdEncoding.DecodeString(m.Key)
	if err != nil {
		return
	}
	m.Key = string(keyBytes)
	if len(m.Key) != 16 && len(m.Key) != 24 && len(m.Key) != 32 {
		return
	}
	var plainTextBlock []byte
	length := len(m.PlainText)
	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}
	copy(plainTextBlock, []byte(m.PlainText))
	block, err := aes.NewCipher([]byte(m.Key))
	if err != nil {
		return
	}
	iv, err := GenerateRandomBytes(aes.BlockSize)
	if err != nil {
		return
	}
	cipherEncoded := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(cipherEncoded, plainTextBlock)
	cipherText = base64.StdEncoding.EncodeToString(cipherEncoded)
	return
}

func (m *AES) Decrypt() (cipherText string, err error) {
	return
}

func (m *AES) GenerateKey() (key []byte, err error) {
	return GenerateRandomBytes(32)
}

func (m *DES) Encrypt() (cipherText string, err error) {
	return
}

func (m *DES) Decrypt() (cipherText string, err error) {
	return
}

func (m *DES) GenerateKey() (key []byte, err error) {
	return GenerateRandomBytes(32)
}
