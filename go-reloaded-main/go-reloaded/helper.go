package main

import (
	"strconv"
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.Fields(text)
}

func HexToDecimal(h string) string {
	d, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		return h
	}
	return strconv.FormatInt(d, 10)
}

func BinToDecimal(b string) string {
	d, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return b
	}
	return strconv.FormatInt(d, 10)
}

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func IsVowel(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
}