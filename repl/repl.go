package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/amirhesham65/hera-lang/lexer"
	"github.com/amirhesham65/hera-lang/token"
)

func Start(userName string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("@%s>> ", userName)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			fmt.Println("exiting...")
			os.Exit(0)
		}

		l := lexer.New(line)

		// Output source code tokens
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
