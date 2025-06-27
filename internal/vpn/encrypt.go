package vpn

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type Crypto struct {
	gcm cipher.AEAD
}

func NewCrypto(key []byte) (*Crypto, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &Crypto{gcm: gcm}, nil
}

func (c *Crypto) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, c.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return append(nonce, c.gcm.Seal(nil, nonce, plaintext, nil)...), nil
}

func (c *Crypto) Decrypt(ciphertext []byte) ([]byte, error) {
	size := c.gcm.NonceSize()
	if len(ciphertext) < size {
		return nil, errors.New("ciphertext too short")
	}
	nonce := ciphertext[:size]
	return c.gcm.Open(nil, nonce, ciphertext[size:], nil)
}
