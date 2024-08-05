package main

import (
	"fmt"
	"os"

	"PegasusScript/lexer"
	"PegasusScript/token"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lexer := lexer.New(string(content))

	for {
		tok := lexer.Consume()
		fmt.Printf("{Token Type: %v, Value: %v}\n", tok.Type, tok.Literal)
		if tok.Type == token.EOF {
			break
		}
	}
}
