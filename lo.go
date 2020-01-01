package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/abdil1234/go_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("halo %s! ini adalah bahasa program go_interpreter\n", user.Username)
	fmt.Printf("Tulis komando yang kamu suka \n")
	repl.Start(os.Stdin, os.Stdout)
}
