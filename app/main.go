package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  rune = '('
	RIGHT_PAREN rune = ')'
	LEFT_BRACE  rune = '{'
	RIGHT_BRACE rune = '}'
	STAR        rune = '*'
	COMMA       rune = ','
	DOT         rune = '.'
	MINUS       rune = '-'
	PLUS        rune = '+'
	SEMICOLON   rune = ';'
)

func unexpectedTokenError(line int, char rune) string {
	return fmt.Sprintf("[line %d] Error: Unexpected character: %c", line, char)
}

func tokenizeString(inp string) (tokens, errorTokens []string) {

	tokens, errorTokens = make([]string, 0), make([]string, 0)
	line := 1
	for _, char := range inp {
		switch char {
		case LEFT_PAREN:
			tokens = append(tokens, "LEFT_PAREN ( null")
		case RIGHT_PAREN:
			tokens = append(tokens, "RIGHT_PAREN ) null")
		case LEFT_BRACE:
			tokens = append(tokens, "LEFT_BRACE { null")
		case RIGHT_BRACE:
			tokens = append(tokens, "RIGHT_BRACE } null")
		case STAR:
			tokens = append(tokens, "STAR * null")
		case COMMA:
			tokens = append(tokens, "COMMA , null")
		case DOT:
			tokens = append(tokens, "DOT . null")
		case MINUS:
			tokens = append(tokens, "MINUS - null")
		case PLUS:
			tokens = append(tokens, "PLUS + null")
		case SEMICOLON:
			tokens = append(tokens, "SEMICOLON ; null")
		case '\n':
			line++
		default:
			errorTokens = append(errorTokens, unexpectedTokenError(line, char))
		}
	}
	tokens = append(tokens, "EOF  null")
	return
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		tokens, errors := tokenizeString(string(fileContents))
		exitCode := 0
		if len(errors) > 0 {
			exitCode = 65
			for _, err := range errors {
				fmt.Fprintln(os.Stderr, err)
			}
		}
		for _, token := range tokens {
			fmt.Fprintln(os.Stdout, token)
		}
		os.Exit(exitCode)
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
