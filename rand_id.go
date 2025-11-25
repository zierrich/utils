package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandID generates IDs like XXXX-XXXX-XXXX-XXXX.
type RandID struct {
	letters []rune
	digits  []rune
}

// Initialize character sets.
func NewRandID() *RandID {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"

	return &RandID{
		letters: []rune(letters),
		digits:  []rune(digits),
	}
}

// Generate full ID of 4 groups.
func (r *RandID) Generate() string {
	parts := make([]string, 4)
	for i := 0; i < 4; i++ {
		parts[i] = r.buildGroup()
	}
	return fmt.Sprintf("%s-%s-%s-%s", parts[0], parts[1], parts[2], parts[3])
}

// Build one 4-char group (2 letters, 2 digits, alternating).
func (r *RandID) buildGroup() string {
	startWithLetter := r.randInt(2) == 0
	out := make([]rune, 4)

	for i := 0; i < 4; i++ {
		if (i%2 == 0 && startWithLetter) || (i%2 == 1 && !startWithLetter) {
			out[i] = r.randomLetter()
		} else {
			out[i] = r.randomDigit()
		}
	}
	return string(out)
}

// Pick random letter.
func (r *RandID) randomLetter() rune {
	return r.letters[r.randInt(len(r.letters))]
}

// Pick random digit.
func (r *RandID) randomDigit() rune {
	return r.digits[r.randInt(len(r.digits))]
}

// Secure random int [0, n).
func (r *RandID) randInt(n int) int {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0
	}
	return int(v.Int64())
}

func main() {
	r := NewRandID()
	for i := 0; i < 10; i++ {
		fmt.Println(r.Generate())
	}
}
