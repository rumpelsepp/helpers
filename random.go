package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	mathrand "math/rand"
	"time"
)

func init() {
	mathrand.Seed(time.Now().UTC().UnixNano())
}

// TODO: Remove this
func RandStr(length int) string {
	var (
		buf  = make([]byte, length)
		char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	)

	for i := 0; i < length; i++ {
		buf[i] = char[mathrand.Intn(len(char)-1)]
	}
	return string(buf)
}

// Adapted from:
//   * https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/
//   * https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

// RandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomBytes(n int) ([]byte, error) {
	assertAvailablePRNG()
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// RandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret = append(ret, letters[num.Int64()])
	}

	return string(ret), nil
}

// RandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func RandomStringURLSafe(n int) (string, error) {
	b, err := RandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
