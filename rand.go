package strings

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
)

const AlphaLower = "abcdefghijklmnopqrstuvwxyz"
const AlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Numbers = "0123456789"
const AlphaNumeric = AlphaLower + AlphaUpper + Numbers
const Token68 = AlphaNumeric + "-._~+/"

// RandomString generates a pseudo random string of the given length using the given characters.
// The distribution is not perfect, but works for general purposes.
func RandomString(source string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = source[rand.Intn(len(source))]
	}
	return string(b)
}

const PasswordLower = "abcdefghijkmnopqrstuvwxyz"
const PasswordUpper = "ABCDEFGHJKLMNPQRSTUVWXYZ"
const PasswordNum = "23456789"
const PasswordSym = "!@#%?+=_"
const PasswordBytes = PasswordLower + PasswordUpper + PasswordNum + PasswordSym

// PasswordString generates a pseudo random password with the given length using characters that are common in passwords.
// We leave out letters that are easily visually confused.
//
// Specific letters excluded are lowercase l, upper case I and the number 1, upper case O and the number 0.
// Also, only easily identifiable and describable symbols are used.
//
// It tries to protect against accidentally creating an easily guessed value by making sure the password has at
// least one lower-case letter, one upper-case letter, one number, and one symbol.
// n must be at least 4, or an empty string will be returned.
func PasswordString(n int) string {
	if n < 4 {
		return ""
	}
	b := make([]byte, n)
	b[0] = PasswordLower[rand.Int()%len(PasswordLower)]
	b[1] = PasswordUpper[rand.Int()%len(PasswordUpper)]
	b[2] = PasswordNum[rand.Int()%len(PasswordNum)]
	b[3] = PasswordSym[rand.Int()%len(PasswordSym)]
	for i := 4; i < len(b); i++ {
		b[i] = PasswordBytes[rand.Int()%len(PasswordBytes)]
	}
	rand.Shuffle(n, func(i, j int) {
		temp := b[i]
		b[i] = b[j]
		b[j] = temp
	})

	return string(b)
}

// CryptoString returns a cryptographically secure random string from the given source.
// Use AlphaNumeric, AlphaUpper, AlphaLower, or Numbers as shortcuts for source.
func CryptoString(source string, n int) string {
	ret := make([]byte, n)
	l := big.NewInt(int64(len(source)))
	for i := 0; i < n; i++ {
		num, _ := crand.Int(crand.Reader, l)
		ret[i] = source[num.Int64()]
	}

	return string(ret)
}
