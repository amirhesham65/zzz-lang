package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/amirhesham65/zzz-lang/repl"
)

func main() {
	currUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("HERA LANG V0 - REPL")
	repl.Start(currUser.Username, os.Stdin, os.Stdout)
}
