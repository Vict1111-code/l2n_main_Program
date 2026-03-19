package main

import (
	"strconv"
	"strings"
)

func ApplyHex(tokens []string) []string {
	for i := 1; i < len(tokens); i++ {
		if tokens[i] == "(hex)" {
			tokens[i-1] = HexToDecimal(tokens[i-1])
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}
	}
	return tokens
}

func ApplyBin(tokens []string) []string {
	for i := 1; i < len(tokens); i++ {
		if tokens[i] == "(bin)" {
			tokens[i-1] = BinToDecimal(tokens[i-1])
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}
	}
	return tokens
}

func ApplyCase(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {

		// (up)
		if tokens[i] == "(up)" && i > 0 {
			tokens[i-1] = strings.ToUpper(tokens[i-1])
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}

		// (low)
		if tokens[i] == "(low)" && i > 0 {
			tokens[i-1] = strings.ToLower(tokens[i-1])
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}

		// (cap)
		if tokens[i] == "(cap)" && i > 0 {
			tokens[i-1] = Capitalize(tokens[i-1])
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}

		// (up, N)
		if tokens[i] == "(up," && i+1 < len(tokens) {
			nStr := strings.Trim(tokens[i+1], ")")
			n, _ := strconv.Atoi(nStr)

			for j := 1; j <= n && i-j >= 0; j++ {
				tokens[i-j] = strings.ToUpper(tokens[i-j])
			}

			// remove "(up," and "N)"
			tokens = append(tokens[:i], tokens[i+2:]...)
			i--
		}

		// (low, N)
		if tokens[i] == "(low," && i+1 < len(tokens) {
			nStr := strings.Trim(tokens[i+1], ")")
			n, _ := strconv.Atoi(nStr)

			for j := 1; j <= n && i-j >= 0; j++ {
				tokens[i-j] = strings.ToLower(tokens[i-j])
			}

			tokens = append(tokens[:i], tokens[i+2:]...)
			i--
		}

		// (cap, N)
		if tokens[i] == "(cap," && i+1 < len(tokens) {
			nStr := strings.Trim(tokens[i+1], ")")
			n, _ := strconv.Atoi(nStr)

			for j := 1; j <= n && i-j >= 0; j++ {
				tokens[i-j] = Capitalize(tokens[i-j])
			}

			tokens = append(tokens[:i], tokens[i+2:]...)
			i--
		}
	}

	return tokens
}

func FixPunctuation(s string) string {
	replacer := strings.NewReplacer(
		" ,", ", ",
		" .", ".",
		" !", "!",
		" ?", "?",
		" :", ":",
		" ;", ";",
	)
	return replacer.Replace(s)
}

func FixArticles(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			if IsVowel(rune(words[i+1][0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}

func FixQuotes(s string) string {
	words := strings.Fields(s)

	var result []string
	inQuote := false

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "'" {
			if !inQuote {
				// opening quote
				inQuote = true

				// attach to next word
				if i+1 < len(words) {
					words[i+1] = "'" + words[i+1]
				}
			} else {
				// closing quote
				inQuote = false

				// attach to previous word
				if len(result) > 0 {
					result[len(result)-1] = result[len(result)-1] + "'"
				}
			}
			continue
		}

		result = append(result, word)
	}

	return strings.Join(result, " ")
}