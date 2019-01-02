package tinycompiler

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/minaorangina/super-tiny-compiler/types"
)

// Tokeniser breaks a lisp-syntax string into its tokens
func Tokeniser(input string) ([]types.Token, error) {
	var current int
	var tokens = []types.Token{}

	for current < len(input) {
		inputAsRunes := []rune(input)
		char := string(inputAsRunes[current])

		if char == "(" {
			tokens = append(tokens, types.Token{
				TokenType: "paren",
				Value:     "(",
			})

			current++
			continue
		}

		if char == ")" {
			tokens = append(tokens, types.Token{
				TokenType: "paren",
				Value:     ")",
			})

			current++
			continue
		}

		isWhitespace, err := regexp.MatchString("\\s", char)
		if err != nil {
			return []types.Token{}, errors.New("Could not compile regex")
		}
		if isWhitespace {
			current++
			continue
		}

		isNumber, err := regexp.MatchString("[0-9]", char)
		if err != nil {
			return []types.Token{}, errors.New("Could not compile regex")
		}
		if isNumber {
			var value string

			for isNumber {
				value += char
				current++
				char = string(inputAsRunes[current])
				isNumber, _ = regexp.MatchString("[0-9]", char)
			}

			tokens = append(tokens, types.Token{
				TokenType: "number",
				Value:     value,
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

			tokens = append(tokens, types.Token{
				TokenType: "string",
				Value:     value,
			})
			continue
		}

		isLetter, err := regexp.MatchString("[A-Za-z]", char)
		if err != nil {
			return []types.Token{}, errors.New("Could not compile regex")
		}
		if isLetter {
			var value string

			for isLetter {
				value += char
				current++
				char = string(inputAsRunes[current])
				isLetter, err = regexp.MatchString("[A-Za-z]", char)
			}

			tokens = append(tokens, types.Token{
				TokenType: "name",
				Value:     value,
			})
			continue
		}

		return []types.Token{}, fmt.Errorf("I don't know what this is: %s", char)
	}

	return tokens, nil
}
