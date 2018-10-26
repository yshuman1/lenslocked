package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenBytes = 32

// Bytes will help us generate n random bytes or will return an error. uses
// crypto/rand so its safe to use with remember tokens
func Bytes(n int) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// String will generate a byte slive of size nBytes and then return a string
// that is base64 url encoded
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken is helper func designed to generate remember tokens of a
// predetermined byte size.
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
