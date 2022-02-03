package main

import (
	"Muffin/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Muffin REPL.\n", user.Username)
	fmt.Println("Type in commands.")
	repl.Start(os.Stdin, os.Stdout)
}
