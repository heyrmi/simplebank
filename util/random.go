package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var rng *rand.Rand

func init() {
	source := rand.NewSource(time.Now().Unix())
	rng = rand.New(source)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name of 6 char
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amout of money
func RandomMoney() int64 {
	return RandomInt(1, 1000)
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "INR", "CAD", "EUR"}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}
