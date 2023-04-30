package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates random string of length 8
func RandomOwner() string {
	return RandomString(8)
}

// RandomMoney generates random int from 0 and 10000
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency return a random currency out of a given list
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "INR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
