package main

import (
	"fmt"
	"os"

	"plutonium/lexer"
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

	lexerNew := lexer.New(string(content))

	for {
		tok := lexerNew.Consume()
		fmt.Printf("{Token Type: %v, Value: %v}\n", tok.Type, tok.Literal)
		if tok.Type == lexer.EOF {
			break
		}
	}
}
