package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// RandNick holds data for nickname generation
type RandNick struct {
	consonants []string
	vowels     []string
	endings    []string
}

// NewRandNick creates a new RandNick instance with default values
func NewRandNick() *RandNick {
	return &RandNick{
		consonants: []string{
			"b", "c", "d", "f", "g", "h", "j", "k", "l", "m",
			"n", "p", "r", "s", "t", "v", "w", "y", "z",
			"th", "sh", "ch", "ph", "qu",
		},
		vowels: []string{"a", "e", "i", "o", "u", "y", "ai", "ei", "ou", "ie", "oo"},
		endings: []string{"an", "el", "or", "in", "us", "ar", "on", "en", "is", "ion", "ius", "ara"},
	}
}

// contains checks if a string is in a slice
func (r *RandNick) contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// generateSyllable creates a syllable based on the last character
func (r *RandNick) generateSyllable(lastChar string) string {
	lastChar = strings.ToLower(lastChar)
	var syllable string
	if r.contains(r.consonants, lastChar) {
		// After consonant — choose vowel or diphthong
		for {
			syllable = r.vowels[rand.Intn(len(r.vowels))]
			if !strings.HasSuffix(lastChar, string(syllable[0])) {
				break
			}
		}
	} else {
		// After vowel — choose consonant or combination
		for {
			syllable = r.consonants[rand.Intn(len(r.consonants))]
			if !strings.HasSuffix(lastChar, string(syllable[0])) {
				break
			}
		}
	}
	return syllable
}

// Generate generates a random nickname of 8 characters
func (r *RandNick) Generate() string {
	nick := ""

	// First letter (capitalized)
	var first string
	if rand.Intn(2) == 0 {
		first = r.consonants[rand.Intn(len(r.consonants))]
	} else {
		first = r.vowels[rand.Intn(len(r.vowels))]
	}
	nick += strings.ToUpper(string(first[0])) + strings.ToLower(first[1:])

	// Generate syllables until length 6
	for len(nick) < 6 {
		lastChar := string(nick[len(nick)-1])
		next := r.generateSyllable(lastChar)

		// Ensure no consecutive identical letters
		for strings.HasSuffix(nick, string(next[0])) {
			next = r.generateSyllable(lastChar)
		}

		if len(nick)+len(next) > 6 {
			next = next[:6-len(nick)]
		}
		nick += strings.ToLower(next)
	}

	// Add realistic ending
	end := r.endings[rand.Intn(len(r.endings))]
	if len(nick)+len(end) > 8 {
		end = end[:8-len(nick)]
	}
	nick += strings.ToLower(end)

	return nick
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rng := NewRandNick()

	for i := 0; i < 10; i++ {
		fmt.Println(rng.Generate())
	}
}