package main

import "strings"

func ProcessText(input string) string {
	tokens := strings.Fields(input)

	tokens = ApplyHex(tokens)
	tokens = ApplyBin(tokens)
	tokens = ApplyCase(tokens)

	result := strings.Join(tokens, " ")

	result = FixPunctuation(result)
	result = FixArticles(result)
	result = FixQuotes(result) 

	return result
}