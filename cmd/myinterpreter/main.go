// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh <command> <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	filename := os.Args[2]

	rawFileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	logEnabled := false
	if command == "tokenize" {
		logEnabled = true
	}

	switch command {
	case "tokenize":
		scanner := NewLexer(string(rawFileContent), logEnabled)
		scanner.ScanTokens()
	case "parse":
		scanner := NewLexer(string(rawFileContent), logEnabled)
		scanner.ScanTokens() // Tokenize first
		parser := NewParser(scanner)
		ast := parser.Parse()
		fmt.Println(ast.String()) // Print the AST
	case "evaluate":
		scanner := NewLexer(string(rawFileContent), logEnabled)
		scanner.ScanTokens() // Tokenize first
		parser := NewParser(scanner)
		ast := parser.Parse()
		result := ast.Eval() // Evaluate the AST
		fmt.Println(result)  // Print the evaluated result
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}
