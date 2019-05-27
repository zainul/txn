package randomdigit

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

const (
	pkg       = "pkg.random"
	maxBigInt = 9223372036854775807
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	var (
		b   []byte
		err error
	)

	b, err = GenerateRandomBytes(s + 15)

	if err != nil {
		return "", err
	}

	str := strings.ToUpper(base64.URLEncoding.EncodeToString(b))

	str = strings.Replace(str, "=", "", -1)
	str = strings.Replace(str, "-", "", -1)
	str = strings.Replace(str, "_", "", -1)
	return str[0:s], nil

}
