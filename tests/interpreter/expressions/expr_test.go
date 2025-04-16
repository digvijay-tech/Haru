package interpreter_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/interpreter"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestArithmeticExpressions(t *testing.T) {
	input := `
		--- ARITHMETIC EXPRESSIONS
		print "ARITHMETIC EXPRESSIONS";
		print 2 + 3 * 4;
		print (2 + 3) * 4;
		print 10 - 3 - 2;
		print 10 - (3 - 2);
		print (2 + 3) * (4 + 5);
		print 2 ** (3 ** 2);
		print (2 ** 3) ** 2;
		print 100 / (5 * (2 + 3));
		print (2 + 3) * (4 + 1) - 6;
		print ((20 - (4 * 2)) / 2) + 1;
		print 100 / (5 + 5) + (6 * 3 - 2);

		--- LOGICAL EXPRESSIONS
		print "LOGICAL EXPRESSIONS";
		print true && true;         	--- true
		print true && false;        	--- false
		print false || true;        	--- true
		print false || false;       	--- false
		print true && false || true; 	--- true (evaluates as (true && false) || true)
		print (true || false) && false; --- false
	`

	expected := `ARITHMETIC EXPRESSIONS
14
20
5
9
45
512
64
4
19
7
26
LOGICAL EXPRESSIONS
true
false
true
false
true
false
`

	// cleaning source input with custom preprocessor
	cleanIn := preprocessor.PreProcess(input)

	// setup ANTLR input and parsing
	is := antlr.NewInputStream(cleanIn)
	lexer := parser.NewharuLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewharuParser(stream)
	tree := p.Program()

	// capturing the output
	output := captureOutput(func() {
		visitor := interpreter.NewHaruVisitor()
		visitor.Visit(tree)
	})

	// normalizing out for Windows
	output = strings.ReplaceAll(output, "\r\n", "\n")

	if output != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, output)
	}
}
