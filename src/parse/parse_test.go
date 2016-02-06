package parse

import (
	"flag"
	"fmt"
	"testing"
	"token"
)

var debug = flag.Bool("debug", false, "show the errors produced by the main tests")

type numberTest struct {
	text    string
	isInt   bool
	isFloat bool
}

var numberTests = []numberTest{
	// basics
	{"0", true, true},
	{"-0", true, true}, // check that -0 is a uint.
	{"73", true, true},
}

var builtins = map[string]interface{}{
	"printf": fmt.Sprintf,
}

func collectTokens(src, left, right string) (tokenList []string) {
	l := lex("testing", src, left, right)

	for {
		tok := l.nextToken()
		tokenList = append(tokenList, token.Tokens[tok.Type()])
		if tok.Type() == token.EOF || tok.Type() == token.ILLEGAL {
			break
		}
	}
	return
}

func TestLetExpr(t *testing.T) {
	src := "let a = 2.43 - 3.1 in 3 + 2"

	fmt.Println(collectTokens(src, "", ""))

	tmpl, err := New("let expression test").Parse(src, "", "", make(map[string]*Tree), builtins)

	if err != nil {
		fmt.Println("!!! Something went wrong! !!!\n", err.Error())
	}

	fmt.Println("parsed", tmpl.Root.String())
}