package main

import (
	"fmt"
	"github.com/Laaman03/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
