package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//Encrypts some data and a secret to a HMAC encrypted message and returns it converted to a string
func encryptHMAC(data []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
