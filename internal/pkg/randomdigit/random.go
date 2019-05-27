package randomdigit

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
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
	b, err := GenerateRandomBytes(s)
	return strings.ToUpper(base64.URLEncoding.EncodeToString(b)), err
}

// GenerateRandomInt ...
func GenerateRandomInt() (int64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxBigInt))
	if err != nil {
		return time.Now().Unix(), nil
	}

	n := nBig.Int64()

	return n, nil

}

// GenerateUserID ...
func GenerateUserID() int64 {
	var ID int64
	uid, _ := GenerateRandomInt()

	nanoTime := time.Now().Nanosecond()

	strTime := strconv.Itoa(nanoTime)
	timeID := fmt.Sprintf("%s%s", strTime, fmt.Sprintf("%06d", uid))[0:17]
	idNotValidate, err := strconv.ParseInt(timeID, 10, 64)

	if err != nil {
		ID = idNotValidate
	}

	for {
		if ID >= maxBigInt {
			ID = GenerateUserID()
		} else {
			return ID
		}
	}
}
