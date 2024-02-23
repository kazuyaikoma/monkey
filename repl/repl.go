package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kazuyaikoma/monkey/lexer"
	"github.com/kazuyaikoma/monkey/token"
)

const PROMPT = ">> "
const EXIT = "exit"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

outer:
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			if tok.Literal == EXIT {
				break outer
			}
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
