package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

var _ Public = (*rsaPub)(nil)
var _ Private = (*rsaPri)(nil)

type Public interface {
	i()
	Encrypt(encryptStr string) (string, error)
}

type Private interface {
	i()
	Decrypt(decryptStr string) (string, error)
}

type rsaPub struct {
	PublicKey string
}

type rsaPri struct {
	PrivateKey string
}

func NewPublic(publicKey string) Public {
	return &rsaPub{
		PublicKey: publicKey,
	}
}

func NewPrivate(privateKey string) Private {
	return &rsaPri{
		PrivateKey: privateKey,
	}
}

func (pub *rsaPub) i() {}

func (pub *rsaPub) Encrypt(encryptStr string) (string, error) {
	// pem
	block, _ := pem.Decode([]byte(pub.PublicKey))

	// x509
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)

	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(encryptStr))
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(encryptedStr), nil
}

func (pri *rsaPri) i() {}

func (pri *rsaPri) Decrypt(decryptStr string) (string, error) {
	// pem
	block, _ := pem.Decode([]byte(pri.PrivateKey))

	// X509
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	decrypted, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decryptBytes)

	return string(decrypted), nil
}
