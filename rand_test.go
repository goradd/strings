package strings

import (
	"strings"
	"testing"
)

func TestPasswordString(t *testing.T) {
	if PasswordString(0) != "" {
		t.Errorf("PasswordString(0) did not return an empty string")
	}
	if PasswordString(3) != "" {
		t.Errorf("PasswordString(3) did not return an empty string")
	}

	s := PasswordString(10)
	if len(s) != 10 {
		t.Errorf("PasswordString(10) did not return string of length 10")
	}
	if !strings.ContainsAny(s, PasswordLower) {
		t.Errorf("PasswordString(10) did not contain a lower case letter")
	}
	if !strings.ContainsAny(s, PasswordUpper) {
		t.Errorf("PasswordString(10) did not contain an upper case letter")
	}
	if !strings.ContainsAny(s, PasswordNum) {
		t.Errorf("PasswordString(10) did not contain a numeric character")
	}
	if !strings.ContainsAny(s, PasswordSym) {
		t.Errorf("PasswordString(10) did not contain a symbol")
	}
}

func TestRandomString(t *testing.T) {
	if RandomString("a", 0) != "" {
		t.Errorf(`RandomString("a", 0) did not return an empty string`)
	}

	if RandomString("a", 1) != "a" {
		t.Errorf(`RandomString("a", 1) did not return "a"`)
	}

	if len(RandomString("abc", 1)) != 1 {
		t.Errorf(`RandomString("abc", 1) did not return a string of length 1`)
	}

	if len(RandomString("abc", 10)) != 10 {
		t.Errorf(`RandomString("abc", 10) did not return a string of length 10`)
	}
}

func TestCryptoString(t *testing.T) {
	if CryptoString("a", 0) != "" {
		t.Errorf(`RandomString("a", 0) did not return an empty string`)
	}

	if CryptoString("a", 1) != "a" {
		t.Errorf(`RandomString("a", 1) did not return "a"`)
	}

	if len(CryptoString("abc", 1)) != 1 {
		t.Errorf(`RandomString("abc", 1) did not return a string of length 1`)
	}

	if len(CryptoString("abc", 10)) != 10 {
		t.Errorf(`RandomString("abc", 10) did not return a string of length 10`)
	}
}
