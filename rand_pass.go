package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type RandPass struct {
	lowerLetters string
	upperLetters string
	digits       string
	specials     string
	passwordLen  int
}

func NewRandPass() *RandPass {
	return &RandPass{
		lowerLetters: "abcdefghijklmnopqrstuvwxyz",
		upperLetters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		digits:       "0123456789",
		specials:     "!@#$%^&*?",
		passwordLen:  16,
	}
}

// Generate builds one password following all rules
func (rp *RandPass) Generate() string {
	letters := rp.pickLetters(10)
	digs := rp.pickDigits(4)
	specs := rp.pickSpecials(2)

	all := append(letters, digs...)
	all = append(all, specs...)

	shuffled := rp.shuffle(all)
	return string(shuffled)
}

// pickLetters returns 10 letters alternating case
func (rp *RandPass) pickLetters(n int) []rune {
	out := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			out = append(out, rp.randChar(rp.lowerLetters))
		} else {
			out = append(out, rp.randChar(rp.upperLetters))
		}
	}
	return out
}

// pickDigits returns required digits
func (rp *RandPass) pickDigits(n int) []rune {
	out := make([]rune, n)
	for i := 0; i < n; i++ {
		out[i] = rp.randChar(rp.digits)
	}
	return out
}

// pickSpecials returns required special characters
func (rp *RandPass) pickSpecials(n int) []rune {
	out := make([]rune, n)
	for i := 0; i < n; i++ {
		out[i] = rp.randChar(rp.specials)
	}
	return out
}

// randChar selects a cryptographically safe random rune
func (rp *RandPass) randChar(chars string) rune {
	i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	return rune(chars[i.Int64()])
}

// shuffle randomly shuffles until rules are satisfied
func (rp *RandPass) shuffle(chars []rune) []rune {
	for {
		rp.cryptoShuffle(chars)
		if rp.isValid(chars) {
			return chars
		}
	}
}

// cryptoShuffle performs Fisher–Yates using crypto/rand
func (rp *RandPass) cryptoShuffle(s []rune) {
	for i := len(s) - 1; i > 0; i-- {
		jb, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		j := int(jb.Int64())
		s[i], s[j] = s[j], s[i]
	}
}

// isValid checks all password constraints
func (rp *RandPass) isValid(p []rune) bool {
	if rp.isSpecial(p[0]) {
		return false
	}

	letters := 0
	digs := 0
	specs := 0

	for _, ch := range p {
		switch {
		case rp.isLetter(ch):
			letters++
			digs, specs = 0, 0
			if letters > 2 {
				return false
			}
		case rp.isDigit(ch):
			digs++
			letters, specs = 0, 0
			if digs > 1 {
				return false
			}
		case rp.isSpecial(ch):
			specs++
			letters, digs = 0, 0
			if specs > 1 {
				return false
			}
		}
	}
	return true
}

func (rp *RandPass) isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func (rp *RandPass) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (rp *RandPass) isSpecial(ch rune) bool {
	for _, s := range rp.specials {
		if ch == s {
			return true
		}
	}
	return false
}

func main() {
	rp := NewRandPass()
	for i := 0; i < 10; i++ {
		fmt.Println(rp.Generate())
	}
}