package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/amirhesham65/hera-lang/evaluator"
	"github.com/amirhesham65/hera-lang/lexer"
	"github.com/amirhesham65/hera-lang/object"
	"github.com/amirhesham65/hera-lang/parser"
)

func Start(userName string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
