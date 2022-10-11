package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

var (
	ErrInvalidBlockSize = errors.New("[!] Invalid block size")

	ErrInvalidPKCS7Data = errors.New("[!] Invalid PKCS7 Data (Empty or Not Padded)")

	ErrInvalidPKCS7Padding = errors.New("[!] Invalid padding on input")
)

func GetRandBuffer(size int) []byte {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func Pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func EncryptBuffer(buf []byte) ([]byte, []byte, []byte) {
	var rawbyte []byte
	key := GetRandBuffer(32)
	iv := GetRandBuffer(16)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	paddedInput, err := Pkcs7Pad([]byte(rawbyte), aes.BlockSize)
	if err != nil {
		log.Fatal(err)
	}

	encBuf := make([]byte, len(paddedInput))
	encMode := cipher.NewCBCEncrypter(block, iv)
	encMode.CryptBlocks(encBuf, paddedInput)

	return iv, key, encBuf
}

func Encrypt(scFile string) string {
	src, _ := ioutil.ReadFile(scFile)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	r := base64.StdEncoding.EncodeToString(dst)

	iv, key, encBuf := EncryptBuffer([]byte(r))
	b64EncBuf := base64.StdEncoding.EncodeToString(encBuf)
	b64Key := base64.StdEncoding.EncodeToString(key)
	b64IV := base64.StdEncoding.EncodeToString(iv)

	return strings.Join([]string{b64IV, b64Key, b64EncBuf}, ":")
}
