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

	// running the code
	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestConstants(t *testing.T) {
	input := `
	--- EXPLICITE CONSTANT DECLARATIONS
	print "EXPLICITE CONSTANT DECLARATIONS";
	const a:i8 = 1;
	const b: f64 = 3.14;
	const c: bool = true;
	const d: string = "hello";

	print a;
	print b;
	print c;
	print d;
	print a + 4 * 2; --- expression with constant

	const e: bool = false;
	const f: bool = !e;
	const g: bool = !e == true && (e != f); --- true
	print e;
	print f;
	print e && f;
	print e || f;
	print g;
`

	expected := `EXPLICITE CONSTANT DECLARATIONS
1
3.14
true
hello
9
false
true
false
true
true
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
