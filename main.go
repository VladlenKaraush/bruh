package main

import (
	"bruh/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is bruh language tokenizer!\n Type commands:", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
