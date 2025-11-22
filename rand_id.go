package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// RandID represents a random ID generator
type RandID struct {
	letters string
	digits  string
	parts   int
	partLen int
}

// NewRandID creates a new RandID generator
func NewRandID() *RandID {
	return &RandID{
		letters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		digits:  "0123456789",
		parts:   4,
		partLen: 4,
	}
}

// Generate produces a random ID in the format XXXX-XXXX-XXXX-XXXX
func (r *RandID) Generate() string {
	parts := make([]string, r.parts)
	for i := 0; i < r.parts; i++ {
		parts[i] = r.randomPart()
	}
	return strings.Join(parts, "-")
}

// randomPart generates a part of the ID alternating letters and digits
func (r *RandID) randomPart() string {
	firstIsDigit := r.randomBool()
	var part strings.Builder
	useDigit := firstIsDigit
	for i := 0; i < r.partLen; i++ {
		if useDigit {
			part.WriteRune(r.randomRune(r.digits))
		} else {
			part.WriteRune(r.randomRune(r.letters))
		}
		useDigit = !useDigit
	}
	return part.String()
}

// randomRune selects a random character from a string
func (r *RandID) randomRune(s string) rune {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(s))))
	return rune(s[n.Int64()])
}

// randomBool returns a random boolean
func (r *RandID) randomBool() bool {
	n, _ := rand.Int(rand.Reader, big.NewInt(2))
	return n.Int64() == 0
}

func main() {
	rid := NewRandID()
	for i := 0; i < 10; i++ {
		fmt.Println(rid.Generate())
	}
}
