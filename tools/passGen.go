package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const usage = `
Generate Password

Generate a bcrypt encrypted password.
Useful for generating test passwords etc.

$ go run passGen.go password
`

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(args[1]), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(hash))
}
