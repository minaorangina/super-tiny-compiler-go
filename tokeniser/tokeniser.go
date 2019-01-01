package tokeniser

import (
	"errors"
	"fmt"
	"regexp"
)

// Token represents a unit of lisp-syntax code
type Token struct {
	tokenType string
	value     string
}

// Tokeniser breaks a lisp-syntax string into its tokens
func Tokeniser(input string) ([]Token, error) {
	var current int
	var tokens = []Token{}

	for current < len(input) {
		inputAsRunes := []rune(input)
		char := string(inputAsRunes[current])

		if char == "(" {
			tokens = append(tokens, Token{
				tokenType: "paren",
				value:     "(",
			})

			current++
			continue
		}

		if char == ")" {
			tokens = append(tokens, Token{
				tokenType: "paren",
				value:     ")",
			})

			current++
			continue
		}

		isWhitespace, err := regexp.MatchString("\\s", char)
		if err != nil {
			return []Token{}, errors.New("Could not compile regex")
		}
		if isWhitespace {
			current++
			continue
		}

		isNumber, err := regexp.MatchString("[0-9]", char)
		if err != nil {
			return []Token{}, errors.New("Could not compile regex")
		}
		if isNumber {
			var value string

			for isNumber {
				value += char
				current++
				char = string(inputAsRunes[current])
				isNumber, _ = regexp.MatchString("[0-9]", char)
			}

			tokens = append(tokens, Token{
				tokenType: "number",
				value:     value,
			})
			continue
		}

		if char == "\"" {
			var value string
			isBetweenQuotes := true

			for isBetweenQuotes {
				value += char
				current++
				char = string(inputAsRunes[current])
				isBetweenQuotes = char != "\""
			}

			tokens = append(tokens, Token{
				tokenType: "string",
				value:     value,
			})
			continue
		}

		isLetter, err := regexp.MatchString("[A-Za-z]", char)
		if err != nil {
			return []Token{}, errors.New("Could not compile regex")
		}
		if isLetter {
			var value string

			for isLetter {
				value += char
				current++
				char = string(inputAsRunes[current])
				isLetter, err = regexp.MatchString("[A-Za-z]", char)
			}

			tokens = append(tokens, Token{
				tokenType: "name",
				value:     value,
			})
			continue
		}

		return []Token{}, fmt.Errorf("I don't know what this is: %s", char)
	}

	return tokens, nil
}
