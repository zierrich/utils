package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandPass struct represents the random password generator
type RandPass struct {
	lettersUpper string
	lettersLower string
	digits       string
	specials     string

	lettersCount  int
	digitsCount   int
	specialsCount int
}

// NewRandPass is the constructor for RandPass
func NewRandPass(lettersCount, digitsCount, specialsCount int) *RandPass {
	return &RandPass{
		lettersUpper:  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		lettersLower:  "abcdefghijklmnopqrstuvwxyz",
		digits:        "0123456789",
		specials:      "!@#$%^&*?",
		lettersCount:  lettersCount,
		digitsCount:   digitsCount,
		specialsCount: specialsCount,
	}
}

// randomChar selects a random character from the given string
func (rp *RandPass) randomChar(chars string) rune {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	return rune(chars[n.Int64()])
}

// randInt returns a random integer in [0,max)
func (rp *RandPass) randInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}

// isDigitOrSpecial checks if a rune is a digit or a special character
func (rp *RandPass) isDigitOrSpecial(r rune) bool {
	for _, c := range rp.digits + rp.specials {
		if r == c {
			return true
		}
	}
	return false
}

// isSpecial checks if a rune is a special character
func (rp *RandPass) isSpecial(r rune) bool {
	for _, c := range rp.specials {
		if r == c {
			return true
		}
	}
	return false
}

// Generate creates a single password following all rules
func (rp *RandPass) Generate() string {
	passwordLength := rp.lettersCount + rp.digitsCount + rp.specialsCount
	password := make([]rune, 0, passwordLength)
	lettersUsed, digitsUsed, specialsUsed := 0, 0, 0

	for len(password) < passwordLength {
		lastIsExtra := len(password) > 0 && rp.isDigitOrSpecial(password[len(password)-1])

		canLetter := lettersUsed < rp.lettersCount
		canDigit := digitsUsed < rp.digitsCount && !lastIsExtra
		canSpecial := specialsUsed < rp.specialsCount && !lastIsExtra && len(password) != 0 && len(password) != passwordLength-1

		options := []string{}
		// Первый символ всегда буква
		if len(password) == 0 && canLetter {
			options = append(options, "letter")
		} else if len(password) == passwordLength-1 && canLetter {
			// Последний символ всегда буква
			options = append(options, "letter")
		} else {
			if canLetter {
				options = append(options, "letter")
			}
			if canDigit {
				options = append(options, "digit")
			}
			if canSpecial {
				options = append(options, "special")
			}
		}

		if len(options) == 0 {
			options = append(options, "letter")
		}

		typ := options[rp.randInt(len(options))]

		var next rune
		switch typ {
		case "letter":
			if lettersUsed%2 == 0 {
				next = rp.randomChar(rp.lettersUpper)
			} else {
				next = rp.randomChar(rp.lettersLower)
			}
			lettersUsed++
		case "digit":
			next = rp.randomChar(rp.digits)
			digitsUsed++
		case "special":
			next = rp.randomChar(rp.specials)
			specialsUsed++
		}

		password = append(password, next)
	}

	return string(password)
}

func main() {
	rp := NewRandPass(10, 4, 2) // 10 letters, 4 digits, 2 special characters
	for i := 0; i < 10; i++ {
		fmt.Println(rp.Generate())
	}
}
