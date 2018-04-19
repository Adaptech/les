package nodejs

import (
	"log"
	"strings"
)

func tokenize(text string) []string {
	return strings.Split(text, " ")
}

func getToken(text string, index int) string {
	tokens := tokenize(text)
	if index > (len(tokens) - 1) {
		log.Fatalln("getToken: Index out of range.")
	}
	return tokens[index]
}
