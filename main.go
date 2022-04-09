package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/Laaman03/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
